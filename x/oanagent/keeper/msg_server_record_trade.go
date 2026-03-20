package keeper

import (
	"context"
	"fmt"

	"oan/x/oanagent/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RecordTrade(goCtx context.Context, msg *types.MsgRecordTrade) (*types.MsgRecordTradeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	agent, found := k.GetAgent(ctx, msg.AgentId)
	if !found {
		return nil, fmt.Errorf("agent %s not found", msg.AgentId)
	}
	agent.TotalTrades++
	if msg.Result == "win" {
		agent.Wins++
	}
	agent.Experience++
	if agent.TotalTrades > 0 {
		agent.WinRateBps = (agent.Wins * 10000) / agent.TotalTrades
	}
	k.SetAgent(ctx, agent)
	ctx.EventManager().EmitEvent(sdk.NewEvent("trade_recorded",
		sdk.NewAttribute("agent_id", msg.AgentId),
		sdk.NewAttribute("action", msg.Action),
		sdk.NewAttribute("result", msg.Result),
	))
	return &types.MsgRecordTradeResponse{
		Success: true, TotalTrades: agent.TotalTrades, WinRateBps: agent.WinRateBps,
	}, nil
}
