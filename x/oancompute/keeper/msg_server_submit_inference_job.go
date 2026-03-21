package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"oan/x/oancompute/types"
)

func (k msgServer) SubmitInferenceJob(goCtx context.Context, msg *types.MsgSubmitInferenceJob) (*types.MsgSubmitInferenceJobResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := k.storeService.OpenKVStore(ctx)

	if msg.InputHash == "" {
		return nil, fmt.Errorf("inputHash cannot be empty")
	}
	if msg.MaxFee == 0 {
		return nil, fmt.Errorf("maxFee must be greater than 0")
	}

	// Check model approved in oancompute's own store
	// Written by register-model which requires guardian approval first
	modelKey := fmt.Sprintf("registered-model-%s", msg.ModelId)
	modelBytes, _ := store.Get([]byte(modelKey))
	if modelBytes == nil {
		return nil, fmt.Errorf("model %s not registered in compute module — register-model first", msg.ModelId)
	}

	jobId := fmt.Sprintf("job-%d-%s", ctx.BlockHeight(), msg.Creator[:8])
	store.Set([]byte(fmt.Sprintf("job-%s", jobId)),
		[]byte(fmt.Sprintf("%s|%s|%s|%d|pending",
			jobId, msg.ModelId, msg.InputHash, ctx.BlockHeight())))

	ctx.EventManager().EmitEvent(sdk.NewEvent("inference_job_submitted",
		sdk.NewAttribute("job_id", jobId),
		sdk.NewAttribute("model_id", msg.ModelId),
		sdk.NewAttribute("input_hash", msg.InputHash),
		sdk.NewAttribute("submitter", msg.Creator),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgSubmitInferenceJobResponse{JobId: jobId, Status: "pending"}, nil
}
