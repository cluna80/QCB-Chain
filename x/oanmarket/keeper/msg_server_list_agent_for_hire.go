package keeper

import (
	"context"
	"fmt"
	"oan/x/oanmarket/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ListAgentForHire(goCtx context.Context, msg *types.MsgListAgentForHire) (*types.MsgListAgentForHireResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if msg.AgentId == "" {
		return nil, fmt.Errorf("agentId cannot be empty")
	}
	if msg.PricePerTask == 0 {
		return nil, fmt.Errorf("pricePerTask must be greater than 0")
	}
	if msg.Skills == "" {
		return nil, fmt.Errorf("skills cannot be empty")
	}
	store := k.storeService.OpenKVStore(ctx)
	ownerKey := fmt.Sprintf("agent-owner-%s", msg.AgentId)
	ownerBytes, _ := store.Get([]byte(ownerKey))
	if ownerBytes != nil {
		if string(ownerBytes) != msg.Creator {
			return nil, fmt.Errorf("only the agent owner can list agent %s for hire", msg.AgentId)
		}
	}
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
