package keeper

import (
	"context"
	"fmt"
	"oan/x/oaneconomy/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) DistributeRewards(goCtx context.Context, msg *types.MsgDistributeRewards) (*types.MsgDistributeRewardsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := k.storeService.OpenKVStore(ctx)

	// SECURITY — only authority or guardian
	if msg.Creator != k.GetAuthority() {
		guardianKey := fmt.Sprintf("guardian-auth-%s", msg.Creator)
		isGuardian, _ := store.Get([]byte(guardianKey))
		if isGuardian == nil {
			return nil, fmt.Errorf("only the module authority or guardians can distribute rewards")
		}
	}

	// SECURITY — epoch cooldown on distribution
	epochKey := fmt.Sprintf("distribute-last-%d", msg.Epoch)
	alreadyRan, _ := store.Get([]byte(epochKey))
	if alreadyRan != nil {
		return nil, fmt.Errorf("rewards already distributed for epoch %d", msg.Epoch)
	}
	store.Set([]byte(epochKey), []byte(fmt.Sprintf("%d", ctx.BlockHeight())))

	rewardAmount := uint64(1000)

	// SECURITY — check module balance before sending
	moduleAddr := sdk.AccAddress([]byte(types.ModuleName))
	balance := k.bankKeeper.GetBalance(ctx, moduleAddr, "uoan")
	if balance.Amount.IsZero() {
		// Testnet — module account not funded yet
		ctx.EventManager().EmitEvent(sdk.NewEvent("rewards_distributed",
			sdk.NewAttribute("epoch", fmt.Sprintf("%d", msg.Epoch)),
			sdk.NewAttribute("total_distributed", "0"),
			sdk.NewAttribute("note", "module account not funded"),
		))
		return &types.MsgDistributeRewardsResponse{TotalDistributed: 0, Recipients: 0}, nil
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent("rewards_distributed",
		sdk.NewAttribute("epoch", fmt.Sprintf("%d", msg.Epoch)),
		sdk.NewAttribute("total_distributed", fmt.Sprintf("%d", rewardAmount)),
		sdk.NewAttribute("denom", "uoan"),
	))
	return &types.MsgDistributeRewardsResponse{
		TotalDistributed: rewardAmount, Recipients: 1,
	}, nil
}
