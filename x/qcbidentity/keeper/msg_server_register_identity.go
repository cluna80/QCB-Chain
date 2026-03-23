package keeper

import (
	"context"
	"fmt"

	"qcb/x/qcbidentity/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RegisterIdentity(goCtx context.Context, msg *types.MsgRegisterIdentity) (*types.MsgRegisterIdentityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if _, found := k.GetIdentity(ctx, msg.Did); found {
		return nil, fmt.Errorf("identity %s already registered", msg.Did)
	}
	identity := types.Identity{
		Index: msg.Did, Did: msg.Did, Owner: msg.Creator,
		DisplayName: msg.DisplayName, IdentityType: msg.IdentityType,
		ReputationScore: 100, Verified: false,
		CreatedAt:  int32(ctx.BlockTime().Unix()),
		LastActive: int32(ctx.BlockTime().Unix()),
	}
	k.SetIdentity(ctx, identity)
	ctx.EventManager().EmitEvent(sdk.NewEvent("identity_registered",
		sdk.NewAttribute("did", msg.Did),
		sdk.NewAttribute("type", msg.IdentityType),
	))
	return &types.MsgRegisterIdentityResponse{Did: msg.Did, Verified: false}, nil
}
