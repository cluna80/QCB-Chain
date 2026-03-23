package qcbcompute

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "qcb/api/qcb/qcbcompute"
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
					RpcMethod:      "SubmitInferenceJob",
					Use:            "submit-inference-job [model-id] [input-hash] [max-fee]",
					Short:          "Send a submit-inference-job tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "modelId"}, {ProtoField: "inputHash"}, {ProtoField: "maxFee"}},
				},
				{
					RpcMethod:      "CompleteInferenceJob",
					Use:            "complete-inference-job [job-id] [output-hash] [proof]",
					Short:          "Send a complete-inference-job tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "jobId"}, {ProtoField: "outputHash"}, {ProtoField: "proof"}},
				},
				{
					RpcMethod:      "SlashBadInference",
					Use:            "slash-bad-inference [job-id] [validator-addr] [evidence]",
					Short:          "Send a slash-bad-inference tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "jobId"}, {ProtoField: "validatorAddr"}, {ProtoField: "evidence"}},
				},
				{
					RpcMethod:      "StakeCompute",
					Use:            "stake-compute [gpu-type] [capacity] [price-per-job]",
					Short:          "Send a stake-compute tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "gpuType"}, {ProtoField: "capacity"}, {ProtoField: "pricePerJob"}},
				},
				{
					RpcMethod:      "RegisterModel",
					Use:            "register-model [model-id] [model-hash] [model-type] [parameters]",
					Short:          "Send a register-model tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "modelId"}, {ProtoField: "modelHash"}, {ProtoField: "modelType"}, {ProtoField: "parameters"}},
				},
				{
					RpcMethod:      "VerifyInferenceProof",
					Use:            "verify-inference-proof [job-id] [proof-hash] [proof-type]",
					Short:          "Send a verify-inference-proof tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "jobId"}, {ProtoField: "proofHash"}, {ProtoField: "proofType"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
