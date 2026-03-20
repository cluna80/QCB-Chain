package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"oan/x/oaneconomy/types"
)

func (k msgServer) DistributeRewards(goCtx context.Context, msg *types.MsgDistributeRewards) (*types.MsgDistributeRewardsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	totalDistributed := uint64(0)
	recipients := int32(0)
	ctx.EventManager().EmitEvent(sdk.NewEvent("rewards_distributed",
		sdk.NewAttribute("epoch", fmt.Sprintf("%d", msg.Epoch)),
		sdk.NewAttribute("total", fmt.Sprintf("%d", totalDistributed)),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgDistributeRewardsResponse{TotalDistributed: totalDistributed, Recipients: recipients}, nil
}
