package oancomms

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "oan/api/oan/oancomms"
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
					RpcMethod:      "RegisterMsgKey",
					Use:            "register-msg-key [key-id] [key-type] [public-key-hash] [algorithm]",
					Short:          "Send a register-msg-key tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "keyId"}, {ProtoField: "keyType"}, {ProtoField: "publicKeyHash"}, {ProtoField: "algorithm"}},
				},
				{
					RpcMethod:      "SendMsgHeader",
					Use:            "send-msg-header [to-addr] [msg-id] [enc-key-id] [payload-hash] [msg-type]",
					Short:          "Send a send-msg-header tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "toAddr"}, {ProtoField: "msgId"}, {ProtoField: "encKeyId"}, {ProtoField: "payloadHash"}, {ProtoField: "msgType"}},
				},
				{
					RpcMethod:      "AckMsg",
					Use:            "ack-msg [msg-id] [ack-type]",
					Short:          "Send a ack-msg tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "msgId"}, {ProtoField: "ackType"}},
				},
				{
					RpcMethod:      "RevokeMsgKey",
					Use:            "revoke-msg-key [key-id] [reason]",
					Short:          "Send a revoke-msg-key tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "keyId"}, {ProtoField: "reason"}},
				},
				{
					RpcMethod:      "SetMsgPolicy",
					Use:            "set-msg-policy [allow-list] [deny-list] [max-inbound] [require-pq-sig]",
					Short:          "Send a set-msg-policy tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "allowList"}, {ProtoField: "denyList"}, {ProtoField: "maxInbound"}, {ProtoField: "requirePqSig"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
