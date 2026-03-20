package oaneconomy

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "oan/api/oan/oaneconomy"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "TaskAll",
					Use:       "list-task",
					Short:     "List all task",
				},
				{
					RpcMethod:      "Task",
					Use:            "show-task [id]",
					Short:          "Shows a task",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateTask",
					Use:            "create-task [title] [description] [reward] [deadline]",
					Short:          "Send a create-task tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "title"}, {ProtoField: "description"}, {ProtoField: "reward"}, {ProtoField: "deadline"}},
				},
				{
					RpcMethod:      "CompleteTask",
					Use:            "complete-task [task-id] [result-hash]",
					Short:          "Send a complete-task tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "taskId"}, {ProtoField: "resultHash"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
