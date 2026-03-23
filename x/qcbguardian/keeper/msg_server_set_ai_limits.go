package keeper

import (
	"context"
	"fmt"
	"qcb/x/qcbguardian/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SetAiLimits(goCtx context.Context, msg *types.MsgSetAiLimits) (*types.MsgSetAiLimitsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if msg.MaxJobsPerBlock == 0 {
		return nil, fmt.Errorf("maxJobsPerBlock must be greater than 0")
	}
	if msg.MaxFeePerJob == 0 {
		return nil, fmt.Errorf("maxFeePerJob must be greater than 0")
	}
	appliedAt := int32(ctx.BlockTime().Unix())
	ctx.EventManager().EmitEvent(sdk.NewEvent("ai_limits_set",
		sdk.NewAttribute("max_jobs_per_block", fmt.Sprintf("%d", msg.MaxJobsPerBlock)),
		sdk.NewAttribute("max_fee_per_job", fmt.Sprintf("%d", msg.MaxFeePerJob)),
		sdk.NewAttribute("allowed_model_types", msg.AllowedModelTypes),
		sdk.NewAttribute("banned_keywords", msg.BannedKeywords),
		sdk.NewAttribute("set_by", msg.Creator),
		sdk.NewAttribute("applied_at", fmt.Sprintf("%d", appliedAt)),
	))
	return &types.MsgSetAiLimitsResponse{Success: true, AppliedAt: appliedAt}, nil
}
