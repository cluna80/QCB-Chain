package keeper

import (
	"context"
	"fmt"
	"qcb/x/qcbeconomy/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ClaimUbi(goCtx context.Context, msg *types.MsgClaimUbi) (*types.MsgClaimUbiResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := k.storeService.OpenKVStore(ctx)

	// SECURITY 1 — verified human or staked
	didKey := fmt.Sprintf("verified-did-%s", msg.Creator)
	verified, _ := store.Get([]byte(didKey))
	stakedKey := fmt.Sprintf("staked-amount-%s", msg.Creator)
	staked, _ := store.Get([]byte(stakedKey))
	if verified == nil && staked == nil {
		return nil, fmt.Errorf("UBI requires a verified human identity or staked OANT")
	}

	// SECURITY 2 — epoch cooldown check
	epochBlocks := int64(100)
	lastClaimKey := fmt.Sprintf("ubi-last-claim-%s", msg.Creator)
	lastClaimBytes, _ := store.Get([]byte(lastClaimKey))
	if lastClaimBytes != nil {
		lastClaim := int64(0)
		for i, b := range lastClaimBytes {
			lastClaim |= int64(b) << (8 * i)
		}
		if ctx.BlockHeight()-lastClaim < epochBlocks {
			blocksLeft := epochBlocks - (ctx.BlockHeight() - lastClaim)
			return nil, fmt.Errorf("UBI already claimed this epoch — %d blocks remaining", blocksLeft)
		}
	}

	// SECURITY 3 — pool health
	poolKey := "ubi-pool-balance"
	poolBytes, _ := store.Get([]byte(poolKey))
	poolBalance := uint64(1000000)
	if poolBytes != nil {
		fmt.Sscanf(string(poolBytes), "%d", &poolBalance)
	}
	params := k.GetParams(ctx)
	ubiAmount := params.UbiRate
	if ubiAmount == 0 {
		ubiAmount = 100000
	}

	// SECURITY 4 — minimum amount check (anti-dust)
	if ubiAmount < 100 {
		return nil, fmt.Errorf("UBI amount too small — minimum 100 charmbits")
	}

	// SECURITY 5 — integer overflow check
	if ubiAmount > 9223372036854775807 {
		return nil, fmt.Errorf("UBI amount overflow — contact DAO")
	}

	if poolBalance < ubiAmount {
		return nil, fmt.Errorf("UBI pool is depleted — DAO must replenish via treasury")
	}

	// SECURITY — write cooldown BEFORE transfer (prevents reentrancy)
	newBalance := poolBalance - ubiAmount
	store.Set([]byte(poolKey), []byte(fmt.Sprintf("%d", newBalance)))
	heightBytes := make([]byte, 8)
	h := ctx.BlockHeight()
	for i := 0; i < 8; i++ {
		heightBytes[i] = byte(h >> (8 * i))
	}
	store.Set([]byte(lastClaimKey), heightBytes)

	// REAL TOKEN TRANSFER — after all state is written
	claimer, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, fmt.Errorf("invalid claimer address: %s", err)
	}

	// Check module balance first
	moduleAddr := sdk.AccAddress([]byte(types.ModuleName))
	balance := k.bankKeeper.GetBalance(ctx, moduleAddr, "charmbits")
	if !balance.Amount.IsZero() {
		coins := sdk.NewCoins(sdk.NewInt64Coin("charmbits", int64(ubiAmount)))
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, claimer, coins); err != nil {
			// Reverse state changes on failure
			store.Set([]byte(poolKey), []byte(fmt.Sprintf("%d", poolBalance)))
			return nil, fmt.Errorf("UBI transfer failed: %s", err)
		}
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent("ubi_claimed",
		sdk.NewAttribute("claimer", msg.Creator),
		sdk.NewAttribute("amount", fmt.Sprintf("%d", ubiAmount)),
		sdk.NewAttribute("denom", "charmbits"),
		sdk.NewAttribute("pool_remaining", fmt.Sprintf("%d", newBalance)),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgClaimUbiResponse{
		Amount: ubiAmount, ClaimedAt: int32(ctx.BlockTime().Unix()),
	}, nil
}
