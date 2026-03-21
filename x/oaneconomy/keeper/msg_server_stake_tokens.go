package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"oan/x/oaneconomy/types"
)

func (k msgServer) StakeTokens(goCtx context.Context, msg *types.MsgStakeTokens) (*types.MsgStakeTokensResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := k.storeService.OpenKVStore(ctx)

	// SECURITY — amount validation
	if msg.Amount == 0 {
		return nil, fmt.Errorf("stake amount must be greater than 0")
	}
	if msg.LockPeriod <= 0 {
		return nil, fmt.Errorf("lock period must be positive")
	}
	// SECURITY — overflow check
	if msg.Amount > 9223372036854775807 {
		return nil, fmt.Errorf("stake amount overflow")
	}
	// SECURITY — minimum stake (anti-dust)
	if msg.Amount < 100 {
		return nil, fmt.Errorf("minimum stake is 100 uoan")
	}

	totalStakeBytes, _ := store.Get([]byte("total-staked"))
	totalStake := uint64(0)
	if totalStakeBytes != nil {
		fmt.Sscanf(string(totalStakeBytes), "%d", &totalStake)
	}
	walletStakeKey := fmt.Sprintf("staked-amount-%s", msg.Creator)
	walletStakeBytes, _ := store.Get([]byte(walletStakeKey))
	walletStake := uint64(0)
	if walletStakeBytes != nil {
		fmt.Sscanf(string(walletStakeBytes), "%d", &walletStake)
	}
	newWalletStake := walletStake + msg.Amount
	newTotal := totalStake + msg.Amount

	// SECURITY — whale cap
	if totalStake >= 10000 {
		walletPct := (newWalletStake * 100) / newTotal
		if walletPct > 5 {
			return nil, fmt.Errorf("staking this amount would exceed the 5%% whale cap")
		}
	}

	// REAL TOKEN TRANSFER — lock OANT in module account
	staker, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, fmt.Errorf("invalid staker address: %s", err)
	}
	coins := sdk.NewCoins(sdk.NewInt64Coin("uoan", int64(msg.Amount)))
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, staker, types.ModuleName, coins); err != nil {
		return nil, fmt.Errorf("failed to lock stake — insufficient balance: %s", err)
	}

	// Write state AFTER successful transfer
	store.Set([]byte(walletStakeKey), []byte(fmt.Sprintf("%d", newWalletStake)))
	store.Set([]byte("total-staked"), []byte(fmt.Sprintf("%d", newTotal)))

	tier := "none"
	switch {
	case newWalletStake >= 20000: tier = "genesis"
	case newWalletStake >= 6000:  tier = "sovereign"
	case newWalletStake >= 1600:  tier = "obsidian"
	case newWalletStake >= 400:   tier = "arcadian"
	}
	store.Set([]byte(fmt.Sprintf("stake-tier-%s", msg.Creator)), []byte(tier))
	store.Set([]byte(fmt.Sprintf("agent-staked-%s", msg.Creator)), []byte(fmt.Sprintf("%d", newWalletStake)))
	store.Set([]byte(fmt.Sprintf("agent-tier-%s", msg.Creator)), []byte(tier))
	store.Set([]byte(fmt.Sprintf("verified-did-%s", msg.Creator)), []byte("staked"))

	stakeId := fmt.Sprintf("stake-%d-%s", ctx.BlockHeight(), msg.Creator[:8])
	unlockBlock := int32(ctx.BlockHeight()) + msg.LockPeriod

	ctx.EventManager().EmitEvent(sdk.NewEvent("tokens_staked",
		sdk.NewAttribute("stake_id", stakeId),
		sdk.NewAttribute("staker", msg.Creator),
		sdk.NewAttribute("amount", fmt.Sprintf("%d", msg.Amount)),
		sdk.NewAttribute("denom", "uoan"),
		sdk.NewAttribute("tier", tier),
		sdk.NewAttribute("unlock_block", fmt.Sprintf("%d", unlockBlock)),
		sdk.NewAttribute("total_staked", fmt.Sprintf("%d", newTotal)),
	))
	return &types.MsgStakeTokensResponse{
		StakeId: stakeId, Amount: msg.Amount, UnlockBlock: unlockBlock,
	}, nil
}
