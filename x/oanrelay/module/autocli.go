package oanrelay

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "oan/api/oan/oanrelay"
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
					RpcMethod:      "RegisterRelay",
					Use:            "register-relay [relay-id] [endpoint] [region] [pub-key-hash]",
					Short:          "Send a register-relay tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "relayId"}, {ProtoField: "endpoint"}, {ProtoField: "region"}, {ProtoField: "pubKeyHash"}},
				},
				{
					RpcMethod:      "RelayHeartbeat",
					Use:            "relay-heartbeat [relay-id] [proof-hash] [blocks-online]",
					Short:          "Send a relay-heartbeat tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "relayId"}, {ProtoField: "proofHash"}, {ProtoField: "blocksOnline"}},
				},
				{
					RpcMethod:      "RouteMsg",
					Use:            "route-msg [msg-id] [from-addr] [to-addr] [relay-id] [payload-ref]",
					Short:          "Send a route-msg tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "msgId"}, {ProtoField: "fromAddr"}, {ProtoField: "toAddr"}, {ProtoField: "relayId"}, {ProtoField: "payloadRef"}},
				},
				{
					RpcMethod:      "SlashRelay",
					Use:            "slash-relay [relay-id] [evidence] [slash-type]",
					Short:          "Send a slash-relay tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "relayId"}, {ProtoField: "evidence"}, {ProtoField: "slashType"}},
				},
				{
					RpcMethod:      "RemoveRelay",
					Use:            "remove-relay [relay-id] [reason]",
					Short:          "Send a remove-relay tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "relayId"}, {ProtoField: "reason"}},
				},
				{
					RpcMethod:      "UpdateRelayRegion",
					Use:            "update-relay-region [relay-id] [new-region] [endpoint]",
					Short:          "Send a update-relay-region tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "relayId"}, {ProtoField: "newRegion"}, {ProtoField: "endpoint"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
