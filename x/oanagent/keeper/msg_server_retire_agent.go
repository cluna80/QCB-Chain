package keeper

import (
	"context"
	"fmt"
	"oan/x/oanagent/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RetireAgent(goCtx context.Context, msg *types.MsgRetireAgent) (*types.MsgRetireAgentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	agent, found := k.GetAgent(ctx, msg.NodeId)
	if !found {
		return nil, fmt.Errorf("agent %s not found", msg.NodeId)
	}
	if agent.Owner != msg.Creator {
		return nil, fmt.Errorf("only the owner can retire an agent")
	}
	agent.Active = false
	k.SetAgent(ctx, agent)
	finalScore := int32((agent.Strength + agent.Agility + agent.Stamina + agent.Skill) / 4)
	ctx.EventManager().EmitEvent(sdk.NewEvent("agent_retired",
		sdk.NewAttribute("node_id", msg.NodeId),
		sdk.NewAttribute("reason", msg.Reason),
		sdk.NewAttribute("final_score", fmt.Sprintf("%d", finalScore)),
	))
	return &types.MsgRetireAgentResponse{Success: true, FinalScore: finalScore}, nil
}
