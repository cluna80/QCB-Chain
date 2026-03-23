package keeper

import (
	"context"
	"fmt"
	"qcb/x/qcbcompute/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) StakeCompute(goCtx context.Context, msg *types.MsgStakeCompute) (*types.MsgStakeComputeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	stakeId := fmt.Sprintf("compute-%d-%s", ctx.BlockHeight(), msg.Creator[:8])
	ctx.EventManager().EmitEvent(sdk.NewEvent("compute_staked",
		sdk.NewAttribute("stake_id", stakeId),
		sdk.NewAttribute("provider", msg.Creator),
		sdk.NewAttribute("gpu_type", msg.GpuType),
		sdk.NewAttribute("capacity", fmt.Sprintf("%d", msg.Capacity)),
		sdk.NewAttribute("price_per_job", fmt.Sprintf("%d", msg.PricePerJob)),
	))
	return &types.MsgStakeComputeResponse{StakeId: stakeId, Capacity: msg.Capacity}, nil
}
