package qcbsports

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"qcb/testutil/sample"
	qcbsportssimulation "qcb/x/qcbsports/simulation"
	"qcb/x/qcbsports/types"
)

// avoid unused import issue
var (
	_ = qcbsportssimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgRegisterAthlete = "op_weight_msg_register_athlete"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRegisterAthlete int = 100

	opWeightMsgCreateStadium = "op_weight_msg_create_stadium"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateStadium int = 100

	opWeightMsgScheduleMatch = "op_weight_msg_schedule_match"
	// TODO: Determine the simulation weight value
	defaultWeightMsgScheduleMatch int = 100

	opWeightMsgRecordMatchResult = "op_weight_msg_record_match_result"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRecordMatchResult int = 100

	opWeightMsgPlacePrediction = "op_weight_msg_place_prediction"
	// TODO: Determine the simulation weight value
	defaultWeightMsgPlacePrediction int = 100

	opWeightMsgClaimPredictionReward = "op_weight_msg_claim_prediction_reward"
	// TODO: Determine the simulation weight value
	defaultWeightMsgClaimPredictionReward int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	qcbsportsGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&qcbsportsGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgRegisterAthlete int
	simState.AppParams.GetOrGenerate(opWeightMsgRegisterAthlete, &weightMsgRegisterAthlete, nil,
		func(_ *rand.Rand) {
			weightMsgRegisterAthlete = defaultWeightMsgRegisterAthlete
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRegisterAthlete,
		qcbsportssimulation.SimulateMsgRegisterAthlete(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateStadium int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateStadium, &weightMsgCreateStadium, nil,
		func(_ *rand.Rand) {
			weightMsgCreateStadium = defaultWeightMsgCreateStadium
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateStadium,
		qcbsportssimulation.SimulateMsgCreateStadium(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgScheduleMatch int
	simState.AppParams.GetOrGenerate(opWeightMsgScheduleMatch, &weightMsgScheduleMatch, nil,
		func(_ *rand.Rand) {
			weightMsgScheduleMatch = defaultWeightMsgScheduleMatch
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgScheduleMatch,
		qcbsportssimulation.SimulateMsgScheduleMatch(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRecordMatchResult int
	simState.AppParams.GetOrGenerate(opWeightMsgRecordMatchResult, &weightMsgRecordMatchResult, nil,
		func(_ *rand.Rand) {
			weightMsgRecordMatchResult = defaultWeightMsgRecordMatchResult
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRecordMatchResult,
		qcbsportssimulation.SimulateMsgRecordMatchResult(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgPlacePrediction int
	simState.AppParams.GetOrGenerate(opWeightMsgPlacePrediction, &weightMsgPlacePrediction, nil,
		func(_ *rand.Rand) {
			weightMsgPlacePrediction = defaultWeightMsgPlacePrediction
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgPlacePrediction,
		qcbsportssimulation.SimulateMsgPlacePrediction(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgClaimPredictionReward int
	simState.AppParams.GetOrGenerate(opWeightMsgClaimPredictionReward, &weightMsgClaimPredictionReward, nil,
		func(_ *rand.Rand) {
			weightMsgClaimPredictionReward = defaultWeightMsgClaimPredictionReward
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgClaimPredictionReward,
		qcbsportssimulation.SimulateMsgClaimPredictionReward(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgRegisterAthlete,
			defaultWeightMsgRegisterAthlete,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				qcbsportssimulation.SimulateMsgRegisterAthlete(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateStadium,
			defaultWeightMsgCreateStadium,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				qcbsportssimulation.SimulateMsgCreateStadium(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgScheduleMatch,
			defaultWeightMsgScheduleMatch,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				qcbsportssimulation.SimulateMsgScheduleMatch(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgRecordMatchResult,
			defaultWeightMsgRecordMatchResult,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				qcbsportssimulation.SimulateMsgRecordMatchResult(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgPlacePrediction,
			defaultWeightMsgPlacePrediction,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				qcbsportssimulation.SimulateMsgPlacePrediction(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgClaimPredictionReward,
			defaultWeightMsgClaimPredictionReward,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				qcbsportssimulation.SimulateMsgClaimPredictionReward(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
