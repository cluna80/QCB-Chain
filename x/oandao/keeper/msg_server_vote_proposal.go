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
	if proposal.Status != "voting" {
		return nil, fmt.Errorf("proposal %s is not open for voting", msg.ProposalId)
	}
	if proposal.Executed {
		return nil, fmt.Errorf("proposal %s already executed", msg.ProposalId)
	}
	store := k.storeService.OpenKVStore(ctx)
	voteKey := fmt.Sprintf("vote-%s-%s", msg.ProposalId, msg.Creator)
	alreadyVoted, _ := store.Get([]byte(voteKey))
	if alreadyVoted != nil {
		return nil, fmt.Errorf("address %s already voted on proposal %s", msg.Creator, msg.ProposalId)
	}
	store.Set([]byte(voteKey), []byte(msg.Vote))
	switch msg.Vote {
	case "yes":
		proposal.VotesYes++
	case "no":
		proposal.VotesNo++
	default:
		return nil, fmt.Errorf("vote must be 'yes' or 'no'")
	}
	k.SetProposal(ctx, proposal)
	ctx.EventManager().EmitEvent(sdk.NewEvent("vote_cast",
		sdk.NewAttribute("proposal_id", msg.ProposalId),
		sdk.NewAttribute("voter", msg.Creator),
		sdk.NewAttribute("vote", msg.Vote),
		sdk.NewAttribute("votes_yes", fmt.Sprintf("%d", proposal.VotesYes)),
		sdk.NewAttribute("votes_no", fmt.Sprintf("%d", proposal.VotesNo)),
	))
	return &types.MsgVoteProposalResponse{
		Success: true, VotesYes: proposal.VotesYes, VotesNo: proposal.VotesNo,
	}, nil
}
