package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"oan/x/oanrelay/types"
)

func (k msgServer) UpdateRelayRegion(goCtx context.Context, msg *types.MsgUpdateRelayRegion) (*types.MsgUpdateRelayRegionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := k.storeService.OpenKVStore(ctx)

	ownerBytes, _ := store.Get([]byte(fmt.Sprintf("relay-owner-%s", msg.RelayId)))
	if ownerBytes == nil {
		return nil, fmt.Errorf("relay %s not found", msg.RelayId)
	}
	if string(ownerBytes) != msg.Creator {
		return nil, fmt.Errorf("only the relay operator can update region")
	}

	store.Set([]byte(fmt.Sprintf("relay-region-%s", msg.RelayId)),
		[]byte(fmt.Sprintf("%s|%s|%d", msg.NewRegion, msg.Endpoint, ctx.BlockHeight())))

	ctx.EventManager().EmitEvent(sdk.NewEvent("relay_region_updated",
		sdk.NewAttribute("relay_id", msg.RelayId),
		sdk.NewAttribute("new_region", msg.NewRegion),
		sdk.NewAttribute("endpoint", msg.Endpoint),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgUpdateRelayRegionResponse{RelayId: msg.RelayId, Updated: true}, nil
}
