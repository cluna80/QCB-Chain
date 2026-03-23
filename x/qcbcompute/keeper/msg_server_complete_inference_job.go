package keeper

import (
	"context"
	"fmt"
	"qcb/x/qcbcompute/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CompleteInferenceJob(goCtx context.Context, msg *types.MsgCompleteInferenceJob) (*types.MsgCompleteInferenceJobResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	reward := k.GetParams(ctx).InferenceReward
	ctx.EventManager().EmitEvent(sdk.NewEvent("inference_job_completed",
		sdk.NewAttribute("job_id", msg.JobId),
		sdk.NewAttribute("output_hash", msg.OutputHash),
		sdk.NewAttribute("validator", msg.Creator),
		sdk.NewAttribute("reward", fmt.Sprintf("%d", reward)),
	))
	return &types.MsgCompleteInferenceJobResponse{Success: true, Reward: reward}, nil
}
