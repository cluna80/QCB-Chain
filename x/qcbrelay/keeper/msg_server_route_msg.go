package keeper

import (
	"context"
	"fmt"
	"qcb/x/qcbrelay/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RouteMsg(goCtx context.Context, msg *types.MsgRouteMsg) (*types.MsgRouteMsgResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := k.storeService.OpenKVStore(ctx)

	// Relay must exist and be active
	relayKey := fmt.Sprintf("relay-%s", msg.RelayId)
	relayData, _ := store.Get([]byte(relayKey))
	if relayData == nil {
		return nil, fmt.Errorf("relay %s not found", msg.RelayId)
	}

	// PayloadRef is off-chain reference — IPFS hash or encrypted blob ref
	if msg.PayloadRef == "" {
		return nil, fmt.Errorf("payloadRef cannot be empty — must reference off-chain encrypted payload")
	}
	if msg.FromAddr == "" || msg.ToAddr == "" {
		return nil, fmt.Errorf("fromAddr and toAddr cannot be empty")
	}

	routeId := fmt.Sprintf("route-%d-%s", ctx.BlockHeight(), msg.MsgId)
	store.Set([]byte(fmt.Sprintf("route-%s", routeId)),
		[]byte(fmt.Sprintf("%s|%s|%s|%s|%s|%d|routed",
			msg.MsgId, msg.FromAddr, msg.ToAddr,
			msg.RelayId, msg.PayloadRef, ctx.BlockHeight())))

	ctx.EventManager().EmitEvent(sdk.NewEvent("msg_routed",
		sdk.NewAttribute("route_id", routeId),
		sdk.NewAttribute("msg_id", msg.MsgId),
		sdk.NewAttribute("relay_id", msg.RelayId),
		sdk.NewAttribute("from", msg.FromAddr),
		sdk.NewAttribute("to", msg.ToAddr),
		sdk.NewAttribute("payload_ref", msg.PayloadRef),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgRouteMsgResponse{
		RouteId: routeId, RelayId: msg.RelayId, Status: "routed",
	}, nil
}
