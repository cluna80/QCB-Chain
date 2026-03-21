package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"oan/x/oannode/types"
)

func (k msgServer) SetNodeConfig(goCtx context.Context, msg *types.MsgSetNodeConfig) (*types.MsgSetNodeConfigResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := k.storeService.OpenKVStore(ctx)

	ownerBytes, _ := store.Get([]byte(fmt.Sprintf("nodeid-%s", msg.NodeId)))
	if ownerBytes == nil {
		return nil, fmt.Errorf("node %s not found", msg.NodeId)
	}
	if string(ownerBytes) != msg.Creator {
		return nil, fmt.Errorf("only the node operator can update config")
	}
	if msg.Endpoint == "" {
		return nil, fmt.Errorf("endpoint cannot be empty")
	}

	configKey := fmt.Sprintf("node-config-%s", msg.NodeId)
	store.Set([]byte(configKey),
		[]byte(fmt.Sprintf("%s|%s|%d", msg.Endpoint, msg.Capabilities, ctx.BlockHeight())))

	ctx.EventManager().EmitEvent(sdk.NewEvent("node_config_updated",
		sdk.NewAttribute("node_id", msg.NodeId),
		sdk.NewAttribute("endpoint", msg.Endpoint),
		sdk.NewAttribute("capabilities", msg.Capabilities),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgSetNodeConfigResponse{NodeId: msg.NodeId, Updated: true}, nil
}
