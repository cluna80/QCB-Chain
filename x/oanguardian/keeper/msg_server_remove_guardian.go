package keeper

import (
	"context"
	"fmt"
	"oan/x/oanguardian/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RemoveGuardian(goCtx context.Context, msg *types.MsgRemoveGuardian) (*types.MsgRemoveGuardianResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	ctx.EventManager().EmitEvent(sdk.NewEvent("guardian_removed",
		sdk.NewAttribute("guardian_addr", msg.GuardianAddr),
		sdk.NewAttribute("reason", msg.Reason),
		sdk.NewAttribute("removed_by", msg.Creator),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgRemoveGuardianResponse{Success: true}, nil
}
