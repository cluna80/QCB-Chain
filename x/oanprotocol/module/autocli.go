package oanprotocol

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "oan/api/oan/oanprotocol"
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
					RpcMethod:      "AddressStatus",
					Use:            "address-status [address]",
					Short:          "Query AddressStatus",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "address"}},
				},

				{
					RpcMethod:      "ProtocolParams",
					Use:            "protocol-params",
					Short:          "Query ProtocolParams",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
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
					RpcMethod:      "SetAddressTier",
					Use:            "set-address-tier [address] [tier]",
					Short:          "Send a SetAddressTier tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "address"}, {ProtoField: "tier"}},
				},
				{
					RpcMethod:      "UpdateLaunchPhase",
					Use:            "update-launch-phase [phase]",
					Short:          "Send a UpdateLaunchPhase tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "phase"}},
				},
				{
					RpcMethod:      "UpdateProtocolParams",
					Use:            "update-protocol-params",
					Short:          "Send a UpdateProtocolParams tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},
				{
					RpcMethod:      "ExemptAddress",
					Use:            "exempt-address [address] [reason]",
					Short:          "Send a ExemptAddress tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "address"}, {ProtoField: "reason"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
