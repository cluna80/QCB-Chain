package qcbqsec

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "qcb/api/qcb/qcbqsec"
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
					RpcMethod:      "RegisterQsKey",
					Use:            "register-qs-key [wallet-addr] [key-type] [public-key-hash] [algorithm]",
					Short:          "Send a register-qs-key tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "walletAddr"}, {ProtoField: "keyType"}, {ProtoField: "publicKeyHash"}, {ProtoField: "algorithm"}},
				},
				{
					RpcMethod:      "RotateQsKey",
					Use:            "rotate-qs-key [old-key-id] [new-public-key-hash] [reason]",
					Short:          "Send a rotate-qs-key tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "oldKeyId"}, {ProtoField: "newPublicKeyHash"}, {ProtoField: "reason"}},
				},
				{
					RpcMethod:      "SubmitHybridSig",
					Use:            "submit-hybrid-sig [tx-hash] [classical-sig] [qs-sig] [key-id]",
					Short:          "Send a submit-hybrid-sig tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "txHash"}, {ProtoField: "classicalSig"}, {ProtoField: "qsSig"}, {ProtoField: "keyId"}},
				},
				{
					RpcMethod:      "SetThreatLevel",
					Use:            "set-threat-level [level] [evidence] [justification]",
					Short:          "Send a set-threat-level tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "level"}, {ProtoField: "evidence"}, {ProtoField: "justification"}},
				},
				{
					RpcMethod:      "RegisterAlgorithm",
					Use:            "register-algorithm [algorithm-id] [algorithm-type] [security-level] [specification]",
					Short:          "Send a register-algorithm tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "algorithmId"}, {ProtoField: "algorithmType"}, {ProtoField: "securityLevel"}, {ProtoField: "specification"}},
				},
				{
					RpcMethod:      "DeprecateAlgorithm",
					Use:            "deprecate-algorithm [algorithm-id] [reason] [replacement-id]",
					Short:          "Send a deprecate-algorithm tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "algorithmId"}, {ProtoField: "reason"}, {ProtoField: "replacementId"}},
				},
				{
					RpcMethod:      "VerifyQsSignature",
					Use:            "verify-qs-signature [tx-hash] [qs-sig] [key-id]",
					Short:          "Send a verify-qs-signature tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "txHash"}, {ProtoField: "qsSig"}, {ProtoField: "keyId"}},
				},
				{
					RpcMethod:      "EmergencyCryptoSwap",
					Use:            "emergency-crypto-swap [from-algorithm] [to-algorithm] [evidence] [urgency]",
					Short:          "Send a emergency-crypto-swap tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "fromAlgorithm"}, {ProtoField: "toAlgorithm"}, {ProtoField: "evidence"}, {ProtoField: "urgency"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
