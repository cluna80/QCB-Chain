package keeper

import (
	"context"
	"fmt"
	"oan/x/oancompute/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SlashBadInference(goCtx context.Context, msg *types.MsgSlashBadInference) (*types.MsgSlashBadInferenceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	slashed := k.GetParams(ctx).SlashAmount
	ctx.EventManager().EmitEvent(sdk.NewEvent("bad_inference_slashed",
		sdk.NewAttribute("job_id", msg.JobId),
		sdk.NewAttribute("validator", msg.ValidatorAddr),
		sdk.NewAttribute("evidence", msg.Evidence),
		sdk.NewAttribute("slashed", fmt.Sprintf("%d", slashed)),
	))
	return &types.MsgSlashBadInferenceResponse{Success: true, Slashed: slashed}, nil
}
