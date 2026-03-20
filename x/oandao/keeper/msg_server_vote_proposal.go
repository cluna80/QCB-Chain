package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"oan/x/oandao/types"
)

func (k msgServer) VoteProposal(goCtx context.Context, msg *types.MsgVoteProposal) (*types.MsgVoteProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	proposal, found := k.GetProposal(ctx, msg.ProposalId)
	if !found {
		return nil, fmt.Errorf("proposal %s not found", msg.ProposalId)
	}
	switch msg.Vote {
	case "yes":
		proposal.VotesYes++
	case "no":
		proposal.VotesNo++
	default:
		return nil, fmt.Errorf("vote must be 'yes' or 'no'")
	}
	k.SetProposal(ctx, proposal)
	return &types.MsgVoteProposalResponse{
		Success: true, VotesYes: proposal.VotesYes, VotesNo: proposal.VotesNo,
	}, nil
}
