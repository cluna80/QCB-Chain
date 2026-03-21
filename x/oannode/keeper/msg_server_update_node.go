package keeper

import (
	"context"
	"fmt"
	"oan/x/oannode/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateNode(goCtx context.Context, msg *types.MsgUpdateNode) (*types.MsgUpdateNodeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := k.storeService.OpenKVStore(ctx)
	ownerKey := fmt.Sprintf("nodeid-%s", msg.NodeId)
	ownerBytes, _ := store.Get([]byte(ownerKey))
	if ownerBytes == nil {
		return nil, fmt.Errorf("node %s not found", msg.NodeId)
	}
	if string(ownerBytes) != msg.Creator {
		return nil, fmt.Errorf("only the node operator can submit uptime proofs")
	}
	if msg.UptimeProof == "" {
		return nil, fmt.Errorf("uptimeProof cannot be empty")
	}
	// Record uptime heartbeat
	heartbeatKey := fmt.Sprintf("node-heartbeat-%s", msg.NodeId)
	store.Set([]byte(heartbeatKey), []byte(fmt.Sprintf("%d", ctx.BlockHeight())))

	// Track consecutive uptime blocks
	uptimeKey := fmt.Sprintf("node-uptime-%s", msg.NodeId)
	uptimeBytes, _ := store.Get([]byte(uptimeKey))
	uptime := uint64(0)
	if uptimeBytes != nil {
		fmt.Sscanf(string(uptimeBytes), "%d", &uptime)
	}
	uptime++
	store.Set([]byte(uptimeKey), []byte(fmt.Sprintf("%d", uptime)))

	ctx.EventManager().EmitEvent(sdk.NewEvent("node_heartbeat",
		sdk.NewAttribute("node_id", msg.NodeId),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
		sdk.NewAttribute("uptime_blocks", fmt.Sprintf("%d", uptime)),
	))
	return &types.MsgUpdateNodeResponse{NodeId: msg.NodeId, Verified: true}, nil
}
