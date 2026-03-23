package qcbwalletproto

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "qcb/api/qcb/qcbwalletproto"
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
					RpcMethod:      "RegisterWalletProfile",
					Use:            "register-wallet-profile [wallet-id] [did] [display-name] [avatar-hash]",
					Short:          "Send a register-wallet-profile tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "walletId"}, {ProtoField: "did"}, {ProtoField: "displayName"}, {ProtoField: "avatarHash"}},
				},
				{
					RpcMethod:      "SetEncryptionKey",
					Use:            "set-encryption-key [wallet-id] [enc-key-hash] [key-type]",
					Short:          "Send a set-encryption-key tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "walletId"}, {ProtoField: "encKeyHash"}, {ProtoField: "keyType"}},
				},
				{
					RpcMethod:      "SetPqKey",
					Use:            "set-pq-key [wallet-id] [pq-key-hash] [algorithm]",
					Short:          "Send a set-pq-key tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "walletId"}, {ProtoField: "pqKeyHash"}, {ProtoField: "algorithm"}},
				},
				{
					RpcMethod:      "UpdateWalletProfile",
					Use:            "update-wallet-profile [wallet-id] [display-name] [avatar-hash]",
					Short:          "Send a update-wallet-profile tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "walletId"}, {ProtoField: "displayName"}, {ProtoField: "avatarHash"}},
				},
				{
					RpcMethod:      "SetWalletPermissions",
					Use:            "set-wallet-permissions [wallet-id] [allow-msgs] [allow-ai-agent] [allow-bridge]",
					Short:          "Send a set-wallet-permissions tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "walletId"}, {ProtoField: "allowMsgs"}, {ProtoField: "allowAiAgent"}, {ProtoField: "allowBridge"}},
				},
				{
					RpcMethod:      "LockWallet",
					Use:            "lock-wallet [wallet-id] [reason]",
					Short:          "Send a lock-wallet tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "walletId"}, {ProtoField: "reason"}},
				},
				{
					RpcMethod:      "UnlockWallet",
					Use:            "unlock-wallet [wallet-id] [proof]",
					Short:          "Send a unlock-wallet tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "walletId"}, {ProtoField: "proof"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
