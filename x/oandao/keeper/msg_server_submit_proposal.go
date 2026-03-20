package keeper

import (
	"context"
	"fmt"

	"oan/x/oandao/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SubmitProposal(goCtx context.Context, msg *types.MsgSubmitProposal) (*types.MsgSubmitProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	proposalId := fmt.Sprintf("prop-%d", ctx.BlockHeight())
	start := int32(ctx.BlockTime().Unix())
	end := start + int32(k.GetParams(ctx).VotingPeriod)
	timelock := end + 172800
	proposal := types.Proposal{
		Index: proposalId, Title: msg.Title, Description: msg.Description,
		Proposer: msg.Creator, VotesYes: 0, VotesNo: 0,
		Status: "voting", StartTime: start, EndTime: end,
		Executed: false, TimelockExpiry: timelock,
	}
	k.SetProposal(ctx, proposal)
	ctx.EventManager().EmitEvent(sdk.NewEvent("proposal_submitted",
		sdk.NewAttribute("proposal_id", proposalId),
		sdk.NewAttribute("title", msg.Title),
	))
	return &types.MsgSubmitProposalResponse{
		ProposalId:     proposalId,
		TimelockExpiry: timelock,
	}, nil
}
