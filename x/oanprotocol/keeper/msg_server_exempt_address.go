package keeper

import (
	"context"
	"fmt"
	"oan/x/oanprotocol/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ExemptAddress(goCtx context.Context, msg *types.MsgExemptAddress) (*types.MsgExemptAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if msg.Creator != k.GetAuthority() {
		return nil, fmt.Errorf("unauthorized: only governance can exempt addresses")
	}
	k.Keeper.ExemptAddress(ctx, msg.Address)
	k.SetAddressTierStore(ctx, msg.Address, "exempt")
	ctx.EventManager().EmitEvent(sdk.NewEvent("address_exempted",
		sdk.NewAttribute("address", msg.Address),
		sdk.NewAttribute("reason", msg.Reason),
		sdk.NewAttribute("by", msg.Creator),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgExemptAddressResponse{}, nil
}
