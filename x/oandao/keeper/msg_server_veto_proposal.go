package keeper

import (
	"context"
	"fmt"
	"oan/x/oandao/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) VetoProposal(goCtx context.Context, msg *types.MsgVetoProposal) (*types.MsgVetoProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	proposal, found := k.GetProposal(ctx, msg.ProposalId)
	if !found {
		return nil, fmt.Errorf("proposal %s not found", msg.ProposalId)
	}
	proposal.Status = "vetoed"
	proposal.Executed = true
	k.SetProposal(ctx, proposal)
	ctx.EventManager().EmitEvent(sdk.NewEvent("proposal_vetoed",
		sdk.NewAttribute("proposal_id", msg.ProposalId),
		sdk.NewAttribute("reason", msg.Reason),
		sdk.NewAttribute("vetoed_by", msg.Creator),
	))
	return &types.MsgVetoProposalResponse{Success: true}, nil
}
