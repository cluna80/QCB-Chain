package keeper

import (
	"context"
	"fmt"
	"qcb/x/qcbidentity/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateReputation(goCtx context.Context, msg *types.MsgUpdateReputation) (*types.MsgUpdateReputationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := k.storeService.OpenKVStore(ctx)
	guardianKey := fmt.Sprintf("guardian-auth-%s", msg.Creator)
	guardianBytes, _ := store.Get([]byte(guardianKey))
	isDaoModule := msg.Creator == "oan10d07y265gmmuvt4z0w9aw880jnsr700j6zn9kn"
	if guardianBytes == nil && !isDaoModule {
		return nil, fmt.Errorf("only guardians or DAO module can update reputation")
	}
	identity, found := k.GetIdentity(ctx, msg.Did)
	if !found {
		return nil, fmt.Errorf("identity %s not found", msg.Did)
	}
	newScore := int64(identity.ReputationScore) + int64(msg.Delta)
	if newScore < 0 {
		newScore = 0
	}
	if newScore > 1000 {
		newScore = 1000
	}
	identity.ReputationScore = uint64(newScore)
	identity.LastActive = int32(ctx.BlockTime().Unix())
	k.SetIdentity(ctx, identity)
	ctx.EventManager().EmitEvent(sdk.NewEvent("reputation_updated",
		sdk.NewAttribute("did", msg.Did),
		sdk.NewAttribute("delta", fmt.Sprintf("%d", msg.Delta)),
		sdk.NewAttribute("new_score", fmt.Sprintf("%d", newScore)),
		sdk.NewAttribute("updated_by", msg.Creator),
	))
	return &types.MsgUpdateReputationResponse{Did: msg.Did, NewScore: int32(newScore)}, nil
}
