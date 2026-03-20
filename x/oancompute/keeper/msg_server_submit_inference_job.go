package keeper

import (
	"context"
	"fmt"
	"oan/x/oancompute/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SubmitInferenceJob(goCtx context.Context, msg *types.MsgSubmitInferenceJob) (*types.MsgSubmitInferenceJobResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if msg.MaxFee == 0 {
		return nil, fmt.Errorf("maxFee must be greater than 0")
	}
	if msg.InputHash == "" {
		return nil, fmt.Errorf("inputHash cannot be empty")
	}
	store := k.storeService.OpenKVStore(ctx)
	approvedKey := fmt.Sprintf("approved-model-%s", msg.ModelId)
	approvedBytes, _ := store.Get([]byte(approvedKey))
	if approvedBytes == nil {
		return nil, fmt.Errorf("model %s is not guardian-approved — submit approve-model first", msg.ModelId)
	}
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
