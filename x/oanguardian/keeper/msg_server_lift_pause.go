package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"oan/x/oanguardian/types"
)

func (k msgServer) LiftPause(goCtx context.Context, msg *types.MsgLiftPause) (*types.MsgLiftPauseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	ctx.EventManager().EmitEvent(sdk.NewEvent("pause_lifted",
		sdk.NewAttribute("pause_id", msg.PauseId),
		sdk.NewAttribute("justification", msg.Justification),
		sdk.NewAttribute("lifted_by", msg.Creator),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgLiftPauseResponse{Success: true}, nil
}
