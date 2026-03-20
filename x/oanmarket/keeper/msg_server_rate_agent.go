package keeper

import (
	"context"
	"fmt"
	"oan/x/oanmarket/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RateAgent(goCtx context.Context, msg *types.MsgRateAgent) (*types.MsgRateAgentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if msg.Rating > 5 {
		return nil, fmt.Errorf("rating must be between 0 and 5")
	}
	ctx.EventManager().EmitEvent(sdk.NewEvent("agent_rated",
		sdk.NewAttribute("agent_id", msg.AgentId),
		sdk.NewAttribute("contract_id", msg.ContractId),
		sdk.NewAttribute("rating", fmt.Sprintf("%d", msg.Rating)),
		sdk.NewAttribute("rater", msg.Creator),
	))
	return &types.MsgRateAgentResponse{Success: true, NewRating: msg.Rating}, nil
}
