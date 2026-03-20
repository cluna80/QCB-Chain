package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"oan/x/oaneconomy/types"
)

func (k msgServer) AcceptTask(goCtx context.Context, msg *types.MsgAcceptTask) (*types.MsgAcceptTaskResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	task, found := k.GetTask(ctx, msg.TaskId)
	if !found {
		return nil, fmt.Errorf("task %s not found", msg.TaskId)
	}
	if task.Status != "open" {
		return nil, fmt.Errorf("task %s is not open", msg.TaskId)
	}
	task.Assignee = msg.Creator
	task.Status = "assigned"
	k.SetTask(ctx, task)
	ctx.EventManager().EmitEvent(sdk.NewEvent("task_accepted",
		sdk.NewAttribute("task_id", msg.TaskId),
		sdk.NewAttribute("assignee", msg.Creator),
	))
	return &types.MsgAcceptTaskResponse{TaskId: msg.TaskId, Assignee: msg.Creator}, nil
}
