package oandao

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "oan/api/oan/oandao"
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
					RpcMethod: "ProposalAll",
					Use:       "list-proposal",
					Short:     "List all proposal",
				},
				{
					RpcMethod:      "Proposal",
					Use:            "show-proposal [id]",
					Short:          "Shows a proposal",
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
					RpcMethod:      "SubmitProposal",
					Use:            "submit-proposal [title] [description]",
					Short:          "Send a submit-proposal tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "title"}, {ProtoField: "description"}},
				},
				{
					RpcMethod:      "VoteProposal",
					Use:            "vote-proposal [proposal-id] [vote]",
					Short:          "Send a vote-proposal tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "proposalId"}, {ProtoField: "vote"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
