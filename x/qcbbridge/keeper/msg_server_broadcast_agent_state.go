package keeper

import (
	"context"
	"fmt"
	"qcb/x/qcbbridge/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) BroadcastAgentState(goCtx context.Context, msg *types.MsgBroadcastAgentState) (*types.MsgBroadcastAgentStateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if msg.StateHash == "" {
		return nil, fmt.Errorf("stateHash cannot be empty")
	}
	if msg.TargetChains == "" {
		return nil, fmt.Errorf("targetChains cannot be empty — specify comma-separated chain IDs")
	}
	broadcastId := fmt.Sprintf("broadcast-%d-%s", ctx.BlockHeight(), msg.AgentId)
	chainCount := uint64(1)
	for _, c := range msg.TargetChains {
		if c == ',' {
			chainCount++
		}
	}
	store := k.storeService.OpenKVStore(ctx)
	store.Set([]byte(fmt.Sprintf("broadcast-%s-latest", msg.AgentId)),
		[]byte(fmt.Sprintf("%s|%s|%d", broadcastId, msg.StateHash, ctx.BlockHeight())))
	ctx.EventManager().EmitEvent(sdk.NewEvent("agent_state_broadcast",
		sdk.NewAttribute("broadcast_id", broadcastId),
		sdk.NewAttribute("agent_id", msg.AgentId),
		sdk.NewAttribute("state_hash", msg.StateHash),
		sdk.NewAttribute("target_chains", msg.TargetChains),
		sdk.NewAttribute("chain_count", fmt.Sprintf("%d", chainCount)),
	))
	return &types.MsgBroadcastAgentStateResponse{BroadcastId: broadcastId, ChainCount: chainCount}, nil
}
