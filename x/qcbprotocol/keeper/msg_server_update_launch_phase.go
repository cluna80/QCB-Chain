package keeper

import (
	"context"
	"fmt"
	"qcb/x/qcbprotocol/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateLaunchPhase(goCtx context.Context, msg *types.MsgUpdateLaunchPhase) (*types.MsgUpdateLaunchPhaseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if msg.Creator != k.GetAuthority() {
		return nil, fmt.Errorf("unauthorized: only governance can update launch phase")
	}
	validPhases := map[string]bool{
		"genesis": true, "validator": true, "ubi": true,
		"ecosystem": true, "dex": true, "open": true,
	}
	if !validPhases[msg.Phase] {
		return nil, fmt.Errorf("invalid phase: must be genesis/validator/ubi/ecosystem/dex/open")
	}
	params := k.GetParams(ctx)
	oldPhase := params.LaunchPhase
	params.LaunchPhase = msg.Phase
	k.SetParams(ctx, params)
	ctx.EventManager().EmitEvent(sdk.NewEvent("launch_phase_updated",
		sdk.NewAttribute("old_phase", oldPhase),
		sdk.NewAttribute("new_phase", msg.Phase),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	k.Logger().Info("qcbprotocol: launch phase advanced", "from", oldPhase, "to", msg.Phase)
	return &types.MsgUpdateLaunchPhaseResponse{}, nil
}
