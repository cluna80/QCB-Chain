package keeper

import (
	"context"
	"fmt"
	"oan/x/oanguardian/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AddGuardian(goCtx context.Context, msg *types.MsgAddGuardian) (*types.MsgAddGuardianResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := k.GetParams(ctx)
	if params.MaxGuardians == 0 {
		params.MaxGuardians = 9
	}
	ctx.EventManager().EmitEvent(sdk.NewEvent("guardian_added",
		sdk.NewAttribute("guardian_addr", msg.GuardianAddr),
		sdk.NewAttribute("display_name", msg.DisplayName),
		sdk.NewAttribute("justification", msg.Justification),
		sdk.NewAttribute("added_by", msg.Creator),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgAddGuardianResponse{GuardianAddr: msg.GuardianAddr, Active: true}, nil
}
