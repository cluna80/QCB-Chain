package qcbmarket

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "qcb/api/qcb/qcbmarket"
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
					RpcMethod:      "ListAgentForHire",
					Use:            "list-agent-for-hire [agent-id] [price-per-task] [skills]",
					Short:          "Send a list-agent-for-hire tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "agentId"}, {ProtoField: "pricePerTask"}, {ProtoField: "skills"}},
				},
				{
					RpcMethod:      "HireAgent",
					Use:            "hire-agent [listing-id] [task-description] [budget]",
					Short:          "Send a hire-agent tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "listingId"}, {ProtoField: "taskDescription"}, {ProtoField: "budget"}},
				},
				{
					RpcMethod:      "RateAgent",
					Use:            "rate-agent [agent-id] [contract-id] [rating] [review]",
					Short:          "Send a rate-agent tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "agentId"}, {ProtoField: "contractId"}, {ProtoField: "rating"}, {ProtoField: "review"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
