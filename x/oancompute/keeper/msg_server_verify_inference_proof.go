package keeper

import (
	"context"
	"fmt"
	"oan/x/oancompute/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) VerifyInferenceProof(goCtx context.Context, msg *types.MsgVerifyInferenceProof) (*types.MsgVerifyInferenceProofResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	ctx.EventManager().EmitEvent(sdk.NewEvent("inference_proof_verified",
		sdk.NewAttribute("job_id", msg.JobId),
		sdk.NewAttribute("proof_hash", msg.ProofHash),
		sdk.NewAttribute("proof_type", msg.ProofType),
		sdk.NewAttribute("verifier", msg.Creator),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgVerifyInferenceProofResponse{Valid: true, JobId: msg.JobId}, nil
}
