package qcbguardian

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"qcb/testutil/sample"
	qcbguardiansimulation "qcb/x/qcbguardian/simulation"
	"qcb/x/qcbguardian/types"
)

// avoid unused import issue
var (
	_ = qcbguardiansimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgAddGuardian = "op_weight_msg_add_guardian"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAddGuardian int = 100

	opWeightMsgRemoveGuardian = "op_weight_msg_remove_guardian"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRemoveGuardian int = 100

	opWeightMsgGuardianVeto = "op_weight_msg_guardian_veto"
	// TODO: Determine the simulation weight value
	defaultWeightMsgGuardianVeto int = 100

	opWeightMsgSetAiLimits = "op_weight_msg_set_ai_limits"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSetAiLimits int = 100

	opWeightMsgEmergencyPause = "op_weight_msg_emergency_pause"
	// TODO: Determine the simulation weight value
	defaultWeightMsgEmergencyPause int = 100

	opWeightMsgLiftPause = "op_weight_msg_lift_pause"
	// TODO: Determine the simulation weight value
	defaultWeightMsgLiftPause int = 100

	opWeightMsgApproveModel = "op_weight_msg_approve_model"
	// TODO: Determine the simulation weight value
	defaultWeightMsgApproveModel int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	qcbguardianGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&qcbguardianGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgAddGuardian int
	simState.AppParams.GetOrGenerate(opWeightMsgAddGuardian, &weightMsgAddGuardian, nil,
		func(_ *rand.Rand) {
			weightMsgAddGuardian = defaultWeightMsgAddGuardian
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddGuardian,
		qcbguardiansimulation.SimulateMsgAddGuardian(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRemoveGuardian int
	simState.AppParams.GetOrGenerate(opWeightMsgRemoveGuardian, &weightMsgRemoveGuardian, nil,
		func(_ *rand.Rand) {
			weightMsgRemoveGuardian = defaultWeightMsgRemoveGuardian
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRemoveGuardian,
		qcbguardiansimulation.SimulateMsgRemoveGuardian(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgGuardianVeto int
	simState.AppParams.GetOrGenerate(opWeightMsgGuardianVeto, &weightMsgGuardianVeto, nil,
		func(_ *rand.Rand) {
			weightMsgGuardianVeto = defaultWeightMsgGuardianVeto
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgGuardianVeto,
		qcbguardiansimulation.SimulateMsgGuardianVeto(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSetAiLimits int
	simState.AppParams.GetOrGenerate(opWeightMsgSetAiLimits, &weightMsgSetAiLimits, nil,
		func(_ *rand.Rand) {
			weightMsgSetAiLimits = defaultWeightMsgSetAiLimits
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSetAiLimits,
		qcbguardiansimulation.SimulateMsgSetAiLimits(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgEmergencyPause int
	simState.AppParams.GetOrGenerate(opWeightMsgEmergencyPause, &weightMsgEmergencyPause, nil,
		func(_ *rand.Rand) {
			weightMsgEmergencyPause = defaultWeightMsgEmergencyPause
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgEmergencyPause,
		qcbguardiansimulation.SimulateMsgEmergencyPause(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgLiftPause int
	simState.AppParams.GetOrGenerate(opWeightMsgLiftPause, &weightMsgLiftPause, nil,
		func(_ *rand.Rand) {
			weightMsgLiftPause = defaultWeightMsgLiftPause
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgLiftPause,
		qcbguardiansimulation.SimulateMsgLiftPause(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgApproveModel int
	simState.AppParams.GetOrGenerate(opWeightMsgApproveModel, &weightMsgApproveModel, nil,
		func(_ *rand.Rand) {
			weightMsgApproveModel = defaultWeightMsgApproveModel
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgApproveModel,
		qcbguardiansimulation.SimulateMsgApproveModel(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgAddGuardian,
			defaultWeightMsgAddGuardian,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				qcbguardiansimulation.SimulateMsgAddGuardian(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgRemoveGuardian,
			defaultWeightMsgRemoveGuardian,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				qcbguardiansimulation.SimulateMsgRemoveGuardian(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgGuardianVeto,
			defaultWeightMsgGuardianVeto,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				qcbguardiansimulation.SimulateMsgGuardianVeto(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgSetAiLimits,
			defaultWeightMsgSetAiLimits,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				qcbguardiansimulation.SimulateMsgSetAiLimits(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgEmergencyPause,
			defaultWeightMsgEmergencyPause,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				qcbguardiansimulation.SimulateMsgEmergencyPause(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgLiftPause,
			defaultWeightMsgLiftPause,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				qcbguardiansimulation.SimulateMsgLiftPause(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgApproveModel,
			defaultWeightMsgApproveModel,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				qcbguardiansimulation.SimulateMsgApproveModel(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
