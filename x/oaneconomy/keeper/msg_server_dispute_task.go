package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"oan/x/oaneconomy/types"
)

func (k msgServer) DisputeTask(goCtx context.Context, msg *types.MsgDisputeTask) (*types.MsgDisputeTaskResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	task, found := k.GetTask(ctx, msg.TaskId)
	if !found {
		return nil, fmt.Errorf("task %s not found", msg.TaskId)
	}
	task.Status = "disputed"
	k.SetTask(ctx, task)
	disputeId := fmt.Sprintf("dispute-%d", ctx.BlockHeight())
	ctx.EventManager().EmitEvent(sdk.NewEvent("task_disputed",
		sdk.NewAttribute("dispute_id", disputeId),
		sdk.NewAttribute("task_id", msg.TaskId),
		sdk.NewAttribute("reason", msg.Reason),
	))
	return &types.MsgDisputeTaskResponse{DisputeId: disputeId, Status: "open"}, nil
}
