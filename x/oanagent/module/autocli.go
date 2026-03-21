package oanagent

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "oan/api/oan/oanagent"
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
					RpcMethod: "AgentAll",
					Use:       "list-agent",
					Short:     "List all agent",
				},
				{
					RpcMethod:      "Agent",
					Use:            "show-agent [id]",
					Short:          "Shows a agent",
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
					RpcMethod:      "RegisterAgent",
					Use:            "register-agent [node-id] [name] [agent-type]",
					Short:          "Send a register-agent tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "nodeId"}, {ProtoField: "name"}, {ProtoField: "agentType"}},
				},
				{
					RpcMethod:      "RecordTrade",
					Use:            "record-trade [agent-id] [action] [amount] [result]",
					Short:          "Send a record-trade tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "agentId"}, {ProtoField: "action"}, {ProtoField: "amount"}, {ProtoField: "result"}},
				},
				{
					RpcMethod:      "BreedAgent",
					Use:            "breed-agent [parent-a] [parent-b] [child-id] [child-name]",
					Short:          "Send a breed-agent tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "parentA"}, {ProtoField: "parentB"}, {ProtoField: "childId"}, {ProtoField: "childName"}},
				},
				{
					RpcMethod:      "ChallengeAgent",
					Use:            "challenge-agent [target-id] [stake]",
					Short:          "Send a challenge-agent tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "targetId"}, {ProtoField: "stake"}},
				},
				{
					RpcMethod:      "RetireAgent",
					Use:            "retire-agent [node-id] [reason]",
					Short:          "Send a retire-agent tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "nodeId"}, {ProtoField: "reason"}},
				},
				{
					RpcMethod:      "SpawnAgent",
					Use:            "spawn-agent [parent-id] [child-id] [child-name] [child-type]",
					Short:          "Send a spawn-agent tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "parentId"}, {ProtoField: "childId"}, {ProtoField: "childName"}, {ProtoField: "childType"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
