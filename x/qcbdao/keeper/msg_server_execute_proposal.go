package keeper

import (
	"context"
	"fmt"
	"qcb/x/qcbdao/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ExecuteProposal(goCtx context.Context, msg *types.MsgExecuteProposal) (*types.MsgExecuteProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	proposal, found := k.GetProposal(ctx, msg.ProposalId)
	if !found {
		return nil, fmt.Errorf("proposal %s not found", msg.ProposalId)
	}
	if proposal.Status != "voting" {
		return nil, fmt.Errorf("proposal %s is not in voting state", msg.ProposalId)
	}
	now := int32(ctx.BlockTime().Unix())
	if now < proposal.TimelockExpiry {
		return nil, fmt.Errorf("timelock has not expired yet")
	}
	if proposal.VotesYes > proposal.VotesNo {
		proposal.Status = "passed"
	} else {
		proposal.Status = "rejected"
	}
	proposal.Executed = true
	k.SetProposal(ctx, proposal)
	ctx.EventManager().EmitEvent(sdk.NewEvent("proposal_executed",
		sdk.NewAttribute("proposal_id", msg.ProposalId),
		sdk.NewAttribute("status", proposal.Status),
	))
	return &types.MsgExecuteProposalResponse{Success: true, ExecutedAt: now}, nil
}
