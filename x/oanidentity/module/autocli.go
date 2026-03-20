package oanidentity

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "oan/api/oan/oanidentity"
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
					RpcMethod: "IdentityAll",
					Use:       "list-identity",
					Short:     "List all identity",
				},
				{
					RpcMethod:      "Identity",
					Use:            "show-identity [id]",
					Short:          "Shows a identity",
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
					RpcMethod:      "RegisterIdentity",
					Use:            "register-identity [did] [display-name] [identity-type]",
					Short:          "Send a register-identity tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "did"}, {ProtoField: "displayName"}, {ProtoField: "identityType"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
