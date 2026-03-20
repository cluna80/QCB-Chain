package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"oan/x/oanagent/types"
)

func (k msgServer) ChallengeAgent(goCtx context.Context, msg *types.MsgChallengeAgent) (*types.MsgChallengeAgentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_, foundChallenger := k.GetAgent(ctx, msg.Creator)
	if !foundChallenger {
		return nil, fmt.Errorf("challenger must be a registered agent")
	}
	_, foundTarget := k.GetAgent(ctx, msg.TargetId)
	if !foundTarget {
		return nil, fmt.Errorf("target agent %s not found", msg.TargetId)
	}
	challengeId := fmt.Sprintf("challenge-%d-%s", ctx.BlockHeight(), msg.Creator[:8])
	ctx.EventManager().EmitEvent(sdk.NewEvent("agent_challenged",
		sdk.NewAttribute("challenge_id", challengeId),
		sdk.NewAttribute("challenger", msg.Creator),
		sdk.NewAttribute("target", msg.TargetId),
		sdk.NewAttribute("stake", fmt.Sprintf("%d", msg.Stake)),
	))
	return &types.MsgChallengeAgentResponse{ChallengeId: challengeId, Status: "pending"}, nil
}
