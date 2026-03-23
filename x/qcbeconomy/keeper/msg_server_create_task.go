package keeper

import (
	"context"
	"fmt"

	"qcb/x/qcbeconomy/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateTask(goCtx context.Context, msg *types.MsgCreateTask) (*types.MsgCreateTaskResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	taskId := fmt.Sprintf("task-%d", ctx.BlockHeight())
	task := types.Task{
		Index: taskId, Title: msg.Title, Description: msg.Description,
		Creator: msg.Creator, Reward: msg.Reward,
		Status: "open", Deadline: msg.Deadline, Completed: false,
	}
	k.SetTask(ctx, task)
	ctx.EventManager().EmitEvent(sdk.NewEvent("task_created",
		sdk.NewAttribute("task_id", taskId),
		sdk.NewAttribute("title", msg.Title),
		sdk.NewAttribute("reward", fmt.Sprintf("%d", msg.Reward)),
	))
	return &types.MsgCreateTaskResponse{TaskId: taskId, Reward: msg.Reward}, nil
}
