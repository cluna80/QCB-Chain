package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"oan/x/oandao/types"
)

func (k msgServer) SubmitProposal(goCtx context.Context, msg *types.MsgSubmitProposal) (*types.MsgSubmitProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if msg.Title == "" {
		return nil, fmt.Errorf("proposal title cannot be empty")
	}
	if msg.Description == "" {
		return nil, fmt.Errorf("proposal description cannot be empty")
	}
	if len(msg.Title) > 100 {
		return nil, fmt.Errorf("proposal title cannot exceed 100 characters")
	}
	if len(msg.Description) > 1000 {
		return nil, fmt.Errorf("proposal description cannot exceed 1000 characters")
	}
	store := k.storeService.OpenKVStore(ctx)
	spamKey := fmt.Sprintf("proposal-spam-%s", msg.Creator)
	lastBytes, _ := store.Get([]byte(spamKey))
	if lastBytes != nil {
		lastBlock := int64(0)
		for i, b := range lastBytes {
			lastBlock |= int64(b) << (8 * i)
		}
		cooldown := int64(50)
		if ctx.BlockHeight()-lastBlock < cooldown {
			blocksLeft := cooldown - (ctx.BlockHeight() - lastBlock)
			return nil, fmt.Errorf("proposal cooldown active — %d blocks remaining", blocksLeft)
		}
	}
	heightBytes := make([]byte, 8)
	h := ctx.BlockHeight()
	for i := 0; i < 8; i++ {
		heightBytes[i] = byte(h >> (8 * i))
	}
	store.Set([]byte(spamKey), heightBytes)
	propId   := fmt.Sprintf("prop-%d", ctx.BlockHeight())
	start    := int32(ctx.BlockTime().Unix())
	end      := start + int32(k.GetParams(ctx).VotingPeriod)
	timelock := end + 172800
	proposal := types.Proposal{
		Index: propId, Title: msg.Title, Description: msg.Description,
		Proposer: msg.Creator, VotesYes: 0, VotesNo: 0,
		Status: "voting", StartTime: start, EndTime: end,
		Executed: false, TimelockExpiry: int32(timelock),
	}
	k.SetProposal(ctx, proposal)
	ctx.EventManager().EmitEvent(sdk.NewEvent("proposal_submitted",
		sdk.NewAttribute("proposal_id", propId),
		sdk.NewAttribute("title", msg.Title),
		sdk.NewAttribute("proposer", msg.Creator),
	))
	return &types.MsgSubmitProposalResponse{ProposalId: propId, TimelockExpiry: int32(timelock)}, nil
}
