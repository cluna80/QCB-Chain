package keeper

import (
	"context"
	"fmt"
	"oan/x/oaneconomy/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ClaimUbi(goCtx context.Context, msg *types.MsgClaimUbi) (*types.MsgClaimUbiResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := k.storeService.OpenKVStore(ctx)

	// PROTECTION 1 — must be verified human OR staked
	// verified-did written by stake-tokens (same store)
	// In production oanidentity writes cross-chain proof
	didKey := fmt.Sprintf("verified-did-%s", msg.Creator)
	verified, _ := store.Get([]byte(didKey))
	stakedKey := fmt.Sprintf("staked-amount-%s", msg.Creator)
	staked, _ := store.Get([]byte(stakedKey))
	if verified == nil && staked == nil {
		return nil, fmt.Errorf("UBI requires a verified human identity or staked OANT — register and verify your DID first")
	}

	// PROTECTION 2 — epoch cooldown
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

	// PROTECTION 3 — pool health
	poolKey := "ubi-pool-balance"
	poolBytes, _ := store.Get([]byte(poolKey))
	poolBalance := uint64(1000000)
	if poolBytes != nil {
		fmt.Sscanf(string(poolBytes), "%d", &poolBalance)
	}
	ubiAmount := uint64(1000)
	if poolBalance < poolBalance/5 {
		ubiAmount = ubiAmount / 10
	}
	if poolBalance < 1000 {
		return nil, fmt.Errorf("UBI pool is depleted — DAO must replenish via treasury")
	}
	newBalance := poolBalance - ubiAmount
	store.Set([]byte(poolKey), []byte(fmt.Sprintf("%d", newBalance)))

	heightBytes := make([]byte, 8)
	h := ctx.BlockHeight()
	for i := 0; i < 8; i++ {
		heightBytes[i] = byte(h >> (8 * i))
	}
	store.Set([]byte(lastClaimKey), heightBytes)

	ctx.EventManager().EmitEvent(sdk.NewEvent("ubi_claimed",
		sdk.NewAttribute("claimer", msg.Creator),
		sdk.NewAttribute("amount", fmt.Sprintf("%d", ubiAmount)),
		sdk.NewAttribute("pool_remaining", fmt.Sprintf("%d", newBalance)),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgClaimUbiResponse{
		Amount: ubiAmount, ClaimedAt: int32(ctx.BlockTime().Unix()),
	}, nil
}
