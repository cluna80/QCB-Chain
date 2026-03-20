package keeper

import (
	"context"
	"fmt"
	"oan/x/oanmarket/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) HireAgent(goCtx context.Context, msg *types.MsgHireAgent) (*types.MsgHireAgentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	contractId := fmt.Sprintf("contract-%d-%s", ctx.BlockHeight(), msg.Creator[:8])
	agentId := fmt.Sprintf("agent-from-%s", msg.ListingId)
	ctx.EventManager().EmitEvent(sdk.NewEvent("agent_hired",
		sdk.NewAttribute("contract_id", contractId),
		sdk.NewAttribute("listing_id", msg.ListingId),
		sdk.NewAttribute("hirer", msg.Creator),
		sdk.NewAttribute("budget", fmt.Sprintf("%d", msg.Budget)),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgHireAgentResponse{ContractId: contractId, AgentId: agentId, Budget: msg.Budget}, nil
}
