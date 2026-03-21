package antirug

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "oan/api/oan/antirug"
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
					RpcMethod:      "RegisterToken",
					Use:            "register-token [token-id] [token-name] [symbol] [max-supply] [liquidity-lock-blocks]",
					Short:          "Send a register-token tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "tokenId"}, {ProtoField: "tokenName"}, {ProtoField: "symbol"}, {ProtoField: "maxSupply"}, {ProtoField: "liquidityLockBlocks"}},
				},
				{
					RpcMethod:      "LockLiquidity",
					Use:            "lock-liquidity [token-id] [amount] [lock-blocks]",
					Short:          "Send a lock-liquidity tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "tokenId"}, {ProtoField: "amount"}, {ProtoField: "lockBlocks"}},
				},
				{
					RpcMethod:      "DeclareMintLimit",
					Use:            "declare-mint-limit [token-id] [max-supply] [max-per-block]",
					Short:          "Send a declare-mint-limit tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "tokenId"}, {ProtoField: "maxSupply"}, {ProtoField: "maxPerBlock"}},
				},
				{
					RpcMethod:      "FlagToken",
					Use:            "flag-token [token-id] [reason] [evidence] [severity]",
					Short:          "Send a flag-token tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "tokenId"}, {ProtoField: "reason"}, {ProtoField: "evidence"}, {ProtoField: "severity"}},
				},
				{
					RpcMethod:      "FreezeToken",
					Use:            "freeze-token [token-id] [reason] [evidence]",
					Short:          "Send a freeze-token tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "tokenId"}, {ProtoField: "reason"}, {ProtoField: "evidence"}},
				},
				{
					RpcMethod:      "UnfreezeToken",
					Use:            "unfreeze-token [token-id] [justification]",
					Short:          "Send a unfreeze-token tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "tokenId"}, {ProtoField: "justification"}},
				},
				{
					RpcMethod:      "RequestUpgrade",
					Use:            "request-upgrade [token-id] [upgrade-type] [description]",
					Short:          "Send a request-upgrade tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "tokenId"}, {ProtoField: "upgradeType"}, {ProtoField: "description"}},
				},
				{
					RpcMethod:      "ApproveToken",
					Use:            "approve-token [token-id] [verdict] [justification]",
					Short:          "Send a approve-token tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "tokenId"}, {ProtoField: "verdict"}, {ProtoField: "justification"}},
				},
				{
					RpcMethod:      "SetModuleActive",
					Use:            "set-module-active [active] [reason]",
					Short:          "Send a set-module-active tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "active"}, {ProtoField: "reason"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
