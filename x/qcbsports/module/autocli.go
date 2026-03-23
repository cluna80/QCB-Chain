package qcbsports

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "qcb/api/qcb/qcbsports"
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
					RpcMethod:      "RegisterAthlete",
					Use:            "register-athlete [athlete-id] [agent-id] [sport] [position]",
					Short:          "Send a register-athlete tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "athleteId"}, {ProtoField: "agentId"}, {ProtoField: "sport"}, {ProtoField: "position"}},
				},
				{
					RpcMethod:      "CreateStadium",
					Use:            "create-stadium [stadium-id] [name] [capacity] [location]",
					Short:          "Send a create-stadium tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "stadiumId"}, {ProtoField: "name"}, {ProtoField: "capacity"}, {ProtoField: "location"}},
				},
				{
					RpcMethod:      "ScheduleMatch",
					Use:            "schedule-match [match-id] [athlete-a] [athlete-b] [stadium-id] [scheduled-at]",
					Short:          "Send a schedule-match tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "matchId"}, {ProtoField: "athleteA"}, {ProtoField: "athleteB"}, {ProtoField: "stadiumId"}, {ProtoField: "scheduledAt"}},
				},
				{
					RpcMethod:      "RecordMatchResult",
					Use:            "record-match-result [match-id] [winner] [loser] [score-a] [score-b] [stats-hash]",
					Short:          "Send a record-match-result tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "matchId"}, {ProtoField: "winner"}, {ProtoField: "loser"}, {ProtoField: "scoreA"}, {ProtoField: "scoreB"}, {ProtoField: "statsHash"}},
				},
				{
					RpcMethod:      "PlacePrediction",
					Use:            "place-prediction [match-id] [prediction] [stake]",
					Short:          "Send a place-prediction tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "matchId"}, {ProtoField: "prediction"}, {ProtoField: "stake"}},
				},
				{
					RpcMethod:      "ClaimPredictionReward",
					Use:            "claim-prediction-reward [prediction-id]",
					Short:          "Send a claim-prediction-reward tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "predictionId"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
