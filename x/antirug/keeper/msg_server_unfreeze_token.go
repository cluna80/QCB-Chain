package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"oan/x/antirug/types"
)

func (k msgServer) UnfreezeToken(goCtx context.Context, msg *types.MsgUnfreezeToken) (*types.MsgUnfreezeTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := k.storeService.OpenKVStore(ctx)

	// Only guardians or DAO can unfreeze
	guardianKey := fmt.Sprintf("guardian-auth-%s", msg.Creator)
	isGuardian, _ := store.Get([]byte(guardianKey))
	isDaoModule := msg.Creator == "oan10d07y265gmmuvt4z0w9aw880jnsr700j6zn9kn"
	if isGuardian == nil && !isDaoModule {
		return nil, fmt.Errorf("only guardians or DAO governance can unfreeze tokens")
	}
	if msg.Justification == "" {
		return nil, fmt.Errorf("justification required to unfreeze a token")
	}

	frozenKey := fmt.Sprintf("antirug-frozen-%s", msg.TokenId)
	frozen, _ := store.Get([]byte(frozenKey))
	if frozen == nil {
		return nil, fmt.Errorf("token %s is not frozen", msg.TokenId)
	}
	store.Delete([]byte(frozenKey))

	ctx.EventManager().EmitEvent(sdk.NewEvent("token_unfrozen",
		sdk.NewAttribute("token_id", msg.TokenId),
		sdk.NewAttribute("justification", msg.Justification),
		sdk.NewAttribute("unfrozen_by", msg.Creator),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgUnfreezeTokenResponse{TokenId: msg.TokenId, Unfrozen: true}, nil
}
