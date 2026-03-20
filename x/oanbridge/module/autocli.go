package oanbridge

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "oan/api/oan/oanbridge"
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
					RpcMethod:      "SendIbcAgent",
					Use:            "send-ibc-agent [agent-id] [dest-chain] [dest-addr]",
					Short:          "Send a send-ibc-agent tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "agentId"}, {ProtoField: "destChain"}, {ProtoField: "destAddr"}},
				},
				{
					RpcMethod:      "BroadcastAgentState",
					Use:            "broadcast-agent-state [agent-id] [state-hash] [target-chains]",
					Short:          "Send a broadcast-agent-state tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "agentId"}, {ProtoField: "stateHash"}, {ProtoField: "targetChains"}},
				},
				{
					RpcMethod:      "PostStateRoot",
					Use:            "post-state-root [state-root] [block-height] [anchor-chain] [proof]",
					Short:          "Send a post-state-root tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "stateRoot"}, {ProtoField: "blockHeight"}, {ProtoField: "anchorChain"}, {ProtoField: "proof"}},
				},
				{
					RpcMethod:      "RegisterChain",
					Use:            "register-chain [chain-id] [chain-name] [bridge-type] [endpoint]",
					Short:          "Send a register-chain tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "chainId"}, {ProtoField: "chainName"}, {ProtoField: "bridgeType"}, {ProtoField: "endpoint"}},
				},
				{
					RpcMethod:      "TokenizeOutput",
					Use:            "tokenize-output [agent-id] [output-type] [content-hash] [value]",
					Short:          "Send a tokenize-output tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "agentId"}, {ProtoField: "outputType"}, {ProtoField: "contentHash"}, {ProtoField: "value"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
