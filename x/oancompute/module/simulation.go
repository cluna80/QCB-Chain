package oancompute

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"oan/testutil/sample"
	oancomputesimulation "oan/x/oancompute/simulation"
	"oan/x/oancompute/types"
)

// avoid unused import issue
var (
	_ = oancomputesimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgSubmitInferenceJob = "op_weight_msg_submit_inference_job"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSubmitInferenceJob int = 100

	opWeightMsgCompleteInferenceJob = "op_weight_msg_complete_inference_job"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCompleteInferenceJob int = 100

	opWeightMsgSlashBadInference = "op_weight_msg_slash_bad_inference"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSlashBadInference int = 100

	opWeightMsgStakeCompute = "op_weight_msg_stake_compute"
	// TODO: Determine the simulation weight value
	defaultWeightMsgStakeCompute int = 100

	opWeightMsgRegisterModel = "op_weight_msg_register_model"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRegisterModel int = 100

	opWeightMsgVerifyInferenceProof = "op_weight_msg_verify_inference_proof"
	// TODO: Determine the simulation weight value
	defaultWeightMsgVerifyInferenceProof int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	oancomputeGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&oancomputeGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgSubmitInferenceJob int
	simState.AppParams.GetOrGenerate(opWeightMsgSubmitInferenceJob, &weightMsgSubmitInferenceJob, nil,
		func(_ *rand.Rand) {
			weightMsgSubmitInferenceJob = defaultWeightMsgSubmitInferenceJob
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSubmitInferenceJob,
		oancomputesimulation.SimulateMsgSubmitInferenceJob(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCompleteInferenceJob int
	simState.AppParams.GetOrGenerate(opWeightMsgCompleteInferenceJob, &weightMsgCompleteInferenceJob, nil,
		func(_ *rand.Rand) {
			weightMsgCompleteInferenceJob = defaultWeightMsgCompleteInferenceJob
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCompleteInferenceJob,
		oancomputesimulation.SimulateMsgCompleteInferenceJob(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSlashBadInference int
	simState.AppParams.GetOrGenerate(opWeightMsgSlashBadInference, &weightMsgSlashBadInference, nil,
		func(_ *rand.Rand) {
			weightMsgSlashBadInference = defaultWeightMsgSlashBadInference
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSlashBadInference,
		oancomputesimulation.SimulateMsgSlashBadInference(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgStakeCompute int
	simState.AppParams.GetOrGenerate(opWeightMsgStakeCompute, &weightMsgStakeCompute, nil,
		func(_ *rand.Rand) {
			weightMsgStakeCompute = defaultWeightMsgStakeCompute
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgStakeCompute,
		oancomputesimulation.SimulateMsgStakeCompute(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRegisterModel int
	simState.AppParams.GetOrGenerate(opWeightMsgRegisterModel, &weightMsgRegisterModel, nil,
		func(_ *rand.Rand) {
			weightMsgRegisterModel = defaultWeightMsgRegisterModel
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRegisterModel,
		oancomputesimulation.SimulateMsgRegisterModel(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgVerifyInferenceProof int
	simState.AppParams.GetOrGenerate(opWeightMsgVerifyInferenceProof, &weightMsgVerifyInferenceProof, nil,
		func(_ *rand.Rand) {
			weightMsgVerifyInferenceProof = defaultWeightMsgVerifyInferenceProof
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgVerifyInferenceProof,
		oancomputesimulation.SimulateMsgVerifyInferenceProof(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgSubmitInferenceJob,
			defaultWeightMsgSubmitInferenceJob,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				oancomputesimulation.SimulateMsgSubmitInferenceJob(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCompleteInferenceJob,
			defaultWeightMsgCompleteInferenceJob,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				oancomputesimulation.SimulateMsgCompleteInferenceJob(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgSlashBadInference,
			defaultWeightMsgSlashBadInference,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				oancomputesimulation.SimulateMsgSlashBadInference(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgStakeCompute,
			defaultWeightMsgStakeCompute,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				oancomputesimulation.SimulateMsgStakeCompute(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgRegisterModel,
			defaultWeightMsgRegisterModel,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				oancomputesimulation.SimulateMsgRegisterModel(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgVerifyInferenceProof,
			defaultWeightMsgVerifyInferenceProof,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				oancomputesimulation.SimulateMsgVerifyInferenceProof(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
