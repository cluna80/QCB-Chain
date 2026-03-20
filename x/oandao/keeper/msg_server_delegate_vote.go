package keeper

import (
	"context"
	"fmt"
	"oan/x/oandao/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) DelegateVote(goCtx context.Context, msg *types.MsgDelegateVote) (*types.MsgDelegateVoteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	ctx.EventManager().EmitEvent(sdk.NewEvent("vote_delegated",
		sdk.NewAttribute("delegator", msg.Creator),
		sdk.NewAttribute("delegate_to", msg.DelegateTo),
		sdk.NewAttribute("power", fmt.Sprintf("%d", msg.Power)),
	))
	return &types.MsgDelegateVoteResponse{Success: true, DelegatedPower: msg.Power}, nil
}
