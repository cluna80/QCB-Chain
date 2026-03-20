package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"oan/x/oaneconomy/types"
)

func (k msgServer) CompleteTask(goCtx context.Context, msg *types.MsgCompleteTask) (*types.MsgCompleteTaskResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	task, found := k.GetTask(ctx, msg.TaskId)
	if !found {
		return nil, fmt.Errorf("task %s not found", msg.TaskId)
	}
	task.Completed  = true
	task.Status     = "completed"
	task.ResultHash = msg.ResultHash
	task.Assignee   = msg.Creator
	k.SetTask(ctx, task)
	ctx.EventManager().EmitEvent(sdk.NewEvent("task_completed",
		sdk.NewAttribute("task_id", msg.TaskId),
		sdk.NewAttribute("result_hash", msg.ResultHash),
	))
	return &types.MsgCompleteTaskResponse{Success: true, Reward: task.Reward}, nil
}
