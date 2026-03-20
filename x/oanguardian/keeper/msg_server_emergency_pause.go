package keeper

import (
	"context"
	"fmt"
	"oan/x/oanguardian/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) EmergencyPause(goCtx context.Context, msg *types.MsgEmergencyPause) (*types.MsgEmergencyPauseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := k.GetParams(ctx)
	pauseDuration := int32(params.PauseDuration)
	if pauseDuration == 0 {
		pauseDuration = 86400
	}
	pauseId := fmt.Sprintf("pause-%d", ctx.BlockHeight())
	now := int32(ctx.BlockTime().Unix())
	expiresAt := now + pauseDuration
	ctx.EventManager().EmitEvent(sdk.NewEvent("emergency_pause_activated",
		sdk.NewAttribute("pause_id", pauseId),
		sdk.NewAttribute("reason", msg.Reason),
		sdk.NewAttribute("evidence", msg.Evidence),
		sdk.NewAttribute("guardian", msg.Creator),
		sdk.NewAttribute("expires_at", fmt.Sprintf("%d", expiresAt)),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgEmergencyPauseResponse{PauseId: pauseId, Active: true, ExpiresAt: expiresAt}, nil
}
