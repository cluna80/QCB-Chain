package qcbnode

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "qcb/api/qcb/qcbnode"
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
					RpcMethod:      "RegisterNode",
					Use:            "register-node [node-type] [endpoint] [node-id]",
					Short:          "Send a register-node tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "nodeType"}, {ProtoField: "endpoint"}, {ProtoField: "nodeId"}},
				},
				{
					RpcMethod:      "UpdateNode",
					Use:            "update-node [node-id] [uptime-proof] [block-height]",
					Short:          "Send a update-node tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "nodeId"}, {ProtoField: "uptimeProof"}, {ProtoField: "blockHeight"}},
				},
				{
					RpcMethod:      "ClaimNodeReward",
					Use:            "claim-node-reward [node-id] [epoch]",
					Short:          "Send a claim-node-reward tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "nodeId"}, {ProtoField: "epoch"}},
				},
				{
					RpcMethod:      "SlashNode",
					Use:            "slash-node [node-id] [evidence] [slash-type]",
					Short:          "Send a slash-node tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "nodeId"}, {ProtoField: "evidence"}, {ProtoField: "slashType"}},
				},
				{
					RpcMethod:      "DeregisterNode",
					Use:            "deregister-node [node-id] [reason]",
					Short:          "Send a deregister-node tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "nodeId"}, {ProtoField: "reason"}},
				},
				{
					RpcMethod:      "SetNodeConfig",
					Use:            "set-node-config [node-id] [endpoint] [capabilities]",
					Short:          "Send a set-node-config tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "nodeId"}, {ProtoField: "endpoint"}, {ProtoField: "capabilities"}},
				},
				{
					RpcMethod:      "ReportNode",
					Use:            "report-node [node-id] [evidence] [violation-type]",
					Short:          "Send a report-node tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "nodeId"}, {ProtoField: "evidence"}, {ProtoField: "violationType"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
