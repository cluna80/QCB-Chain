package keeper

import (
	"context"
	"fmt"
	"oan/x/antirug/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) FreezeToken(goCtx context.Context, msg *types.MsgFreezeToken) (*types.MsgFreezeTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := k.storeService.OpenKVStore(ctx)

	// NO DORMANCY CHECK — circuit breaker always available for emergencies
	// But only guardians can freeze
	guardianKey := fmt.Sprintf("guardian-auth-%s", msg.Creator)
	isGuardian, _ := store.Get([]byte(guardianKey))
	isDaoModule := msg.Creator == "oan10d07y265gmmuvt4z0w9aw880jnsr700j6zn9kn"
	if isGuardian == nil && !isDaoModule {
		return nil, fmt.Errorf("only guardians can freeze tokens — protects against abuse of freeze power")
	}
	if msg.Evidence == "" {
		return nil, fmt.Errorf("evidence required to freeze a token")
	}

	frozenAt := int32(ctx.BlockTime().Unix())
	store.Set([]byte(fmt.Sprintf("antirug-frozen-%s", msg.TokenId)),
		[]byte(fmt.Sprintf("frozen|%s|%s|%d", msg.Reason, msg.Evidence, ctx.BlockHeight())))

	ctx.EventManager().EmitEvent(sdk.NewEvent("token_frozen",
		sdk.NewAttribute("token_id", msg.TokenId),
		sdk.NewAttribute("reason", msg.Reason),
		sdk.NewAttribute("frozen_by", msg.Creator),
		sdk.NewAttribute("frozen_at", fmt.Sprintf("%d", frozenAt)),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgFreezeTokenResponse{TokenId: msg.TokenId, Frozen: true, FrozenAt: frozenAt}, nil
}
