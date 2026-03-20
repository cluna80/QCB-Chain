package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"oan/x/oanguardian/types"
)

func (k msgServer) GuardianVeto(goCtx context.Context, msg *types.MsgGuardianVeto) (*types.MsgGuardianVetoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if msg.Severity != "low" && msg.Severity != "medium" && msg.Severity != "high" && msg.Severity != "critical" {
		return nil, fmt.Errorf("severity must be low, medium, high, or critical")
	}
	vetoId := fmt.Sprintf("veto-%d-%s", ctx.BlockHeight(), msg.Creator[:8])
	ctx.EventManager().EmitEvent(sdk.NewEvent("guardian_veto_issued",
		sdk.NewAttribute("veto_id", vetoId),
		sdk.NewAttribute("job_id", msg.JobId),
		sdk.NewAttribute("reason", msg.Reason),
		sdk.NewAttribute("severity", msg.Severity),
		sdk.NewAttribute("guardian", msg.Creator),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgGuardianVetoResponse{VetoId: vetoId, JobId: msg.JobId, Halted: true}, nil
}
