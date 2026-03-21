package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"oan/x/oancompute/types"
)

func (k msgServer) RegisterModel(goCtx context.Context, msg *types.MsgRegisterModel) (*types.MsgRegisterModelResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := k.storeService.OpenKVStore(ctx)

	if msg.ModelId == "" {
		return nil, fmt.Errorf("modelId cannot be empty")
	}
	if msg.ModelHash == "" {
		return nil, fmt.Errorf("modelHash cannot be empty")
	}

	// Write to store so submit-inference-job can find it
	modelKey := fmt.Sprintf("registered-model-%s", msg.ModelId)
	store.Set([]byte(modelKey),
		[]byte(fmt.Sprintf("%s|%s|%s|%d|%d",
			msg.ModelId, msg.ModelHash, msg.ModelType,
			msg.Parameters, ctx.BlockHeight())))

	ctx.EventManager().EmitEvent(sdk.NewEvent("model_registered",
		sdk.NewAttribute("model_id", msg.ModelId),
		sdk.NewAttribute("model_hash", msg.ModelHash),
		sdk.NewAttribute("model_type", msg.ModelType),
		sdk.NewAttribute("parameters", fmt.Sprintf("%d", msg.Parameters)),
		sdk.NewAttribute("registrar", msg.Creator),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgRegisterModelResponse{ModelId: msg.ModelId, Registered: true}, nil
}
