package keeper

import (
	"context"
	"fmt"
	"oan/x/oanidentity/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) VerifyIdentity(goCtx context.Context, msg *types.MsgVerifyIdentity) (*types.MsgVerifyIdentityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	identity, found := k.GetIdentity(ctx, msg.Did)
	if !found {
		return nil, fmt.Errorf("identity %s not found", msg.Did)
	}
	identity.Verified = true
	identity.LastActive = int32(ctx.BlockTime().Unix())
	k.SetIdentity(ctx, identity)
	ctx.EventManager().EmitEvent(sdk.NewEvent("identity_verified",
		sdk.NewAttribute("did", msg.Did),
		sdk.NewAttribute("proof_type", msg.ProofType),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgVerifyIdentityResponse{Did: msg.Did, Verified: true}, nil
}
