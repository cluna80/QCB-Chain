package qcbguardian

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "qcb/api/qcb/qcbguardian"
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
					RpcMethod:      "AddGuardian",
					Use:            "add-guardian [guardian-addr] [display-name] [justification]",
					Short:          "Send a add-guardian tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "guardianAddr"}, {ProtoField: "displayName"}, {ProtoField: "justification"}},
				},
				{
					RpcMethod:      "RemoveGuardian",
					Use:            "remove-guardian [guardian-addr] [reason]",
					Short:          "Send a remove-guardian tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "guardianAddr"}, {ProtoField: "reason"}},
				},
				{
					RpcMethod:      "GuardianVeto",
					Use:            "guardian-veto [job-id] [reason] [severity]",
					Short:          "Send a guardian-veto tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "jobId"}, {ProtoField: "reason"}, {ProtoField: "severity"}},
				},
				{
					RpcMethod:      "SetAiLimits",
					Use:            "set-ai-limits [max-jobs-per-block] [max-fee-per-job] [allowed-model-types] [banned-keywords]",
					Short:          "Send a set-ai-limits tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "maxJobsPerBlock"}, {ProtoField: "maxFeePerJob"}, {ProtoField: "allowedModelTypes"}, {ProtoField: "bannedKeywords"}},
				},
				{
					RpcMethod:      "EmergencyPause",
					Use:            "emergency-pause [reason] [evidence]",
					Short:          "Send a emergency-pause tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "reason"}, {ProtoField: "evidence"}},
				},
				{
					RpcMethod:      "LiftPause",
					Use:            "lift-pause [pause-id] [justification]",
					Short:          "Send a lift-pause tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "pauseId"}, {ProtoField: "justification"}},
				},
				{
					RpcMethod:      "ApproveModel",
					Use:            "approve-model [model-id] [verdict] [justification]",
					Short:          "Send a approve-model tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "modelId"}, {ProtoField: "verdict"}, {ProtoField: "justification"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
