package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"oan/x/oancompute/types"
)

func (k msgServer) SubmitInferenceJob(goCtx context.Context, msg *types.MsgSubmitInferenceJob) (*types.MsgSubmitInferenceJobResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	jobId := fmt.Sprintf("job-%d-%s", ctx.BlockHeight(), msg.Creator[:8])
	ctx.EventManager().EmitEvent(sdk.NewEvent("inference_job_submitted",
		sdk.NewAttribute("job_id", jobId),
		sdk.NewAttribute("model_id", msg.ModelId),
		sdk.NewAttribute("input_hash", msg.InputHash),
		sdk.NewAttribute("max_fee", fmt.Sprintf("%d", msg.MaxFee)),
		sdk.NewAttribute("submitter", msg.Creator),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgSubmitInferenceJobResponse{JobId: jobId, Status: "pending"}, nil
}
