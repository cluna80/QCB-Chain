package keeper

import (
	"context"
	"fmt"
	"oan/x/oannode/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) DeregisterNode(goCtx context.Context, msg *types.MsgDeregisterNode) (*types.MsgDeregisterNodeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := k.storeService.OpenKVStore(ctx)

	ownerBytes, _ := store.Get([]byte(fmt.Sprintf("nodeid-%s", msg.NodeId)))
	if ownerBytes == nil {
		return nil, fmt.Errorf("node %s not found", msg.NodeId)
	}
	if string(ownerBytes) != msg.Creator {
		return nil, fmt.Errorf("only the node operator can deregister their node")
	}

	// FAILSAFE — check if jailed
	jailed, _ := store.Get([]byte(fmt.Sprintf("node-jailed-%s", msg.NodeId)))
	if jailed != nil {
		return nil, fmt.Errorf("jailed nodes cannot deregister — submit appeal via report-node")
	}

	// FAILSAFE — 21 day unbonding period
	params := k.GetParams(ctx)
	unbondingBlocks := int64(params.UnbondingBlocks)
	if unbondingBlocks == 0 {
		unbondingBlocks = 302400
	} // ~21 days at 6s blocks

	unbondingEndsAt := int32(ctx.BlockHeight()) + int32(unbondingBlocks)

	// Mark as unbonding — not immediately released
	store.Set([]byte(fmt.Sprintf("node-unbonding-%s", msg.NodeId)),
		[]byte(fmt.Sprintf("%d", unbondingEndsAt)))
	store.Delete([]byte(fmt.Sprintf("nodeid-%s", msg.NodeId)))

	ctx.EventManager().EmitEvent(sdk.NewEvent("node_deregistered",
		sdk.NewAttribute("node_id", msg.NodeId),
		sdk.NewAttribute("operator", msg.Creator),
		sdk.NewAttribute("unbonding_ends_at", fmt.Sprintf("%d", unbondingEndsAt)),
		sdk.NewAttribute("reason", msg.Reason),
	))
	return &types.MsgDeregisterNodeResponse{NodeId: msg.NodeId, UnbondingEndsAt: unbondingEndsAt}, nil
}
