package keeper

import (
	"context"
	"fmt"
	"oan/x/oancompute/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RegisterModel(goCtx context.Context, msg *types.MsgRegisterModel) (*types.MsgRegisterModelResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
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
