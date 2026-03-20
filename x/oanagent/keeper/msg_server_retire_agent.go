package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"oan/x/oanagent/types"
)

func (k msgServer) RetireAgent(goCtx context.Context, msg *types.MsgRetireAgent) (*types.MsgRetireAgentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	agent, found := k.GetAgent(ctx, msg.NodeId)
	if !found {
		return nil, fmt.Errorf("agent %s not found", msg.NodeId)
	}
	if agent.Owner != msg.Creator {
		return nil, fmt.Errorf("only the owner can retire agent %s", msg.NodeId)
	}
	if !agent.Active {
		return nil, fmt.Errorf("agent %s is already retired", msg.NodeId)
	}
	if msg.Reason == "" {
		return nil, fmt.Errorf("retirement reason cannot be empty")
	}
	agent.Active = false
	k.SetAgent(ctx, agent)
	finalScore := int32((agent.Strength + agent.Agility + agent.Stamina + agent.Skill) / 4)
	ctx.EventManager().EmitEvent(sdk.NewEvent("agent_retired",
		sdk.NewAttribute("node_id", msg.NodeId),
		sdk.NewAttribute("owner", msg.Creator),
		sdk.NewAttribute("reason", msg.Reason),
		sdk.NewAttribute("final_score", fmt.Sprintf("%d", finalScore)),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgRetireAgentResponse{Success: true, FinalScore: finalScore}, nil
}
