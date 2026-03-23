package keeper

import (
	"context"
	"fmt"
	"qcb/x/qcbprotocol/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateProtocolParams(goCtx context.Context, msg *types.MsgUpdateProtocolParams) (*types.MsgUpdateProtocolParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if msg.Creator != k.GetAuthority() {
		return nil, fmt.Errorf("unauthorized: only governance can update protocol params")
	}
	params := k.GetParams(ctx)
	k.SetParams(ctx, params)
	ctx.EventManager().EmitEvent(sdk.NewEvent("protocol_params_updated",
		sdk.NewAttribute("updated_by", msg.Creator),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgUpdateProtocolParamsResponse{}, nil
}
