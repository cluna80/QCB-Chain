package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"oan/x/oanidentity/types"
)

func (k msgServer) VerifyIdentity(goCtx context.Context, msg *types.MsgVerifyIdentity) (*types.MsgVerifyIdentityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	identity, found := k.GetIdentity(ctx, msg.Did)
	if !found {
		return nil, fmt.Errorf("identity %s not found — register first", msg.Did)
	}
	if msg.ProofType == "" {
		return nil, fmt.Errorf("proofType cannot be empty")
	}

	identity.Verified = true
	k.SetIdentity(ctx, identity)

	// Write verified key to oanidentity store
	store := k.storeService.OpenKVStore(ctx)
	store.Set([]byte(fmt.Sprintf("verified-did-%s", msg.Creator)), []byte(msg.Did))
	store.Set([]byte(fmt.Sprintf("did-owner-%s", msg.Did)), []byte(msg.Creator))

	ctx.EventManager().EmitEvent(sdk.NewEvent("identity_verified",
		sdk.NewAttribute("did", msg.Did),
		sdk.NewAttribute("proof_type", msg.ProofType),
		sdk.NewAttribute("verifier", msg.Creator),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgVerifyIdentityResponse{Did: msg.Did, Verified: true}, nil
}
