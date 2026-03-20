package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"oan/x/oaneconomy/types"
)

func (k msgServer) StakeTokens(goCtx context.Context, msg *types.MsgStakeTokens) (*types.MsgStakeTokensResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	stakeId := fmt.Sprintf("stake-%d-%s", ctx.BlockHeight(), msg.Creator[:8])
	unlockBlock := int32(ctx.BlockHeight()) + msg.LockPeriod
	ctx.EventManager().EmitEvent(sdk.NewEvent("tokens_staked",
		sdk.NewAttribute("stake_id", stakeId),
		sdk.NewAttribute("staker", msg.Creator),
		sdk.NewAttribute("amount", fmt.Sprintf("%d", msg.Amount)),
		sdk.NewAttribute("unlock_block", fmt.Sprintf("%d", unlockBlock)),
	))
	return &types.MsgStakeTokensResponse{StakeId: stakeId, Amount: msg.Amount, UnlockBlock: unlockBlock}, nil
}
