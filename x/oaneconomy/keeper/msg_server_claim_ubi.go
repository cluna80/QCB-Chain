package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"oan/x/oaneconomy/types"
)

func (k msgServer) ClaimUbi(goCtx context.Context, msg *types.MsgClaimUbi) (*types.MsgClaimUbiResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	ubiAmount := uint64(1000)
	claimedAt := int32(ctx.BlockTime().Unix())
	ctx.EventManager().EmitEvent(sdk.NewEvent("ubi_claimed",
		sdk.NewAttribute("claimer", msg.Creator),
		sdk.NewAttribute("amount", fmt.Sprintf("%d", ubiAmount)),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgClaimUbiResponse{Amount: ubiAmount, ClaimedAt: claimedAt}, nil
}
