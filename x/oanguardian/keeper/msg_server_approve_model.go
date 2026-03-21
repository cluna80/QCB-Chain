package keeper

import (
	"context"
	"fmt"
	"oan/x/oanguardian/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ApproveModel(goCtx context.Context, msg *types.MsgApproveModel) (*types.MsgApproveModelResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := k.storeService.OpenKVStore(ctx)

	if msg.Verdict != "approved" && msg.Verdict != "rejected" && msg.Verdict != "restricted" {
		return nil, fmt.Errorf("verdict must be approved, rejected, or restricted")
	}

	approved := msg.Verdict == "approved"
	approvedKey := fmt.Sprintf("approved-model-%s", msg.ModelId)
	sharedKey := fmt.Sprintf("guardian-approved-model-%s", msg.ModelId)

	if approved {
		store.Set([]byte(approvedKey), []byte("1"))
		store.Set([]byte(sharedKey), []byte(fmt.Sprintf("approved|%d", ctx.BlockHeight())))
	} else {
		store.Delete([]byte(approvedKey))
		store.Delete([]byte(sharedKey))
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent("model_approval_decision",
		sdk.NewAttribute("model_id", msg.ModelId),
		sdk.NewAttribute("verdict", msg.Verdict),
		sdk.NewAttribute("decided_by", msg.Creator),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgApproveModelResponse{ModelId: msg.ModelId, Approved: approved}, nil
}
