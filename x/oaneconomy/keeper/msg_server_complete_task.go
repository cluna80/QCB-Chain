package keeper

import (
	"context"
	"fmt"
	"oan/x/oaneconomy/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CompleteTask(goCtx context.Context, msg *types.MsgCompleteTask) (*types.MsgCompleteTaskResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	task, found := k.GetTask(ctx, msg.TaskId)
	if !found {
		return nil, fmt.Errorf("task %s not found", msg.TaskId)
	}
	if task.Completed {
		return nil, fmt.Errorf("task %s already completed", msg.TaskId)
	}
	if task.Status == "disputed" {
		return nil, fmt.Errorf("task %s is under dispute", msg.TaskId)
	}
	if task.Assignee != "" && task.Assignee != msg.Creator {
		return nil, fmt.Errorf("only the assigned worker can complete this task")
	}
	if task.Creator == msg.Creator && task.Assignee == "" {
		return nil, fmt.Errorf("task creator cannot self-complete an unassigned task")
	}
	task.Completed = true
	task.Status = "completed"
	task.ResultHash = msg.ResultHash
	task.Assignee = msg.Creator
	k.SetTask(ctx, task)
	ctx.EventManager().EmitEvent(sdk.NewEvent("task_completed",
		sdk.NewAttribute("task_id", msg.TaskId),
		sdk.NewAttribute("completed_by", msg.Creator),
		sdk.NewAttribute("result_hash", msg.ResultHash),
		sdk.NewAttribute("reward", fmt.Sprintf("%d", task.Reward)),
	))
	return &types.MsgCompleteTaskResponse{Success: true, Reward: task.Reward}, nil
}
