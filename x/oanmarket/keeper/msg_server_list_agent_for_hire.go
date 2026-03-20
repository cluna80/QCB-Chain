package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"oan/x/oanmarket/types"
)

func (k msgServer) ListAgentForHire(goCtx context.Context, msg *types.MsgListAgentForHire) (*types.MsgListAgentForHireResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	listingId := fmt.Sprintf("listing-%d-%s", ctx.BlockHeight(), msg.AgentId)
	ctx.EventManager().EmitEvent(sdk.NewEvent("agent_listed_for_hire",
		sdk.NewAttribute("listing_id", listingId),
		sdk.NewAttribute("agent_id", msg.AgentId),
		sdk.NewAttribute("price_per_task", fmt.Sprintf("%d", msg.PricePerTask)),
		sdk.NewAttribute("skills", msg.Skills),
		sdk.NewAttribute("owner", msg.Creator),
	))
	return &types.MsgListAgentForHireResponse{ListingId: listingId, PricePerTask: msg.PricePerTask}, nil
}
