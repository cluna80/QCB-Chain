package qcbqsec

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"qcb/testutil/sample"
	qcbqsecsimulation "qcb/x/qcbqsec/simulation"
	"qcb/x/qcbqsec/types"
)

// avoid unused import issue
var (
	_ = qcbqsecsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgRegisterQsKey = "op_weight_msg_register_qs_key"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRegisterQsKey int = 100

	opWeightMsgRotateQsKey = "op_weight_msg_rotate_qs_key"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRotateQsKey int = 100

	opWeightMsgSubmitHybridSig = "op_weight_msg_submit_hybrid_sig"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSubmitHybridSig int = 100

	opWeightMsgSetThreatLevel = "op_weight_msg_set_threat_level"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSetThreatLevel int = 100

	opWeightMsgRegisterAlgorithm = "op_weight_msg_register_algorithm"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRegisterAlgorithm int = 100

	opWeightMsgDeprecateAlgorithm = "op_weight_msg_deprecate_algorithm"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeprecateAlgorithm int = 100

	opWeightMsgVerifyQsSignature = "op_weight_msg_verify_qs_signature"
	// TODO: Determine the simulation weight value
	defaultWeightMsgVerifyQsSignature int = 100

	opWeightMsgEmergencyCryptoSwap = "op_weight_msg_emergency_crypto_swap"
	// TODO: Determine the simulation weight value
	defaultWeightMsgEmergencyCryptoSwap int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	qcbqsecGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&qcbqsecGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgRegisterQsKey int
	simState.AppParams.GetOrGenerate(opWeightMsgRegisterQsKey, &weightMsgRegisterQsKey, nil,
		func(_ *rand.Rand) {
			weightMsgRegisterQsKey = defaultWeightMsgRegisterQsKey
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRegisterQsKey,
		qcbqsecsimulation.SimulateMsgRegisterQsKey(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRotateQsKey int
	simState.AppParams.GetOrGenerate(opWeightMsgRotateQsKey, &weightMsgRotateQsKey, nil,
		func(_ *rand.Rand) {
			weightMsgRotateQsKey = defaultWeightMsgRotateQsKey
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRotateQsKey,
		qcbqsecsimulation.SimulateMsgRotateQsKey(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSubmitHybridSig int
	simState.AppParams.GetOrGenerate(opWeightMsgSubmitHybridSig, &weightMsgSubmitHybridSig, nil,
		func(_ *rand.Rand) {
			weightMsgSubmitHybridSig = defaultWeightMsgSubmitHybridSig
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSubmitHybridSig,
		qcbqsecsimulation.SimulateMsgSubmitHybridSig(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSetThreatLevel int
	simState.AppParams.GetOrGenerate(opWeightMsgSetThreatLevel, &weightMsgSetThreatLevel, nil,
		func(_ *rand.Rand) {
			weightMsgSetThreatLevel = defaultWeightMsgSetThreatLevel
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSetThreatLevel,
		qcbqsecsimulation.SimulateMsgSetThreatLevel(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRegisterAlgorithm int
	simState.AppParams.GetOrGenerate(opWeightMsgRegisterAlgorithm, &weightMsgRegisterAlgorithm, nil,
		func(_ *rand.Rand) {
			weightMsgRegisterAlgorithm = defaultWeightMsgRegisterAlgorithm
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRegisterAlgorithm,
		qcbqsecsimulation.SimulateMsgRegisterAlgorithm(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeprecateAlgorithm int
	simState.AppParams.GetOrGenerate(opWeightMsgDeprecateAlgorithm, &weightMsgDeprecateAlgorithm, nil,
		func(_ *rand.Rand) {
			weightMsgDeprecateAlgorithm = defaultWeightMsgDeprecateAlgorithm
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeprecateAlgorithm,
		qcbqsecsimulation.SimulateMsgDeprecateAlgorithm(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgVerifyQsSignature int
	simState.AppParams.GetOrGenerate(opWeightMsgVerifyQsSignature, &weightMsgVerifyQsSignature, nil,
		func(_ *rand.Rand) {
			weightMsgVerifyQsSignature = defaultWeightMsgVerifyQsSignature
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgVerifyQsSignature,
		qcbqsecsimulation.SimulateMsgVerifyQsSignature(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgEmergencyCryptoSwap int
	simState.AppParams.GetOrGenerate(opWeightMsgEmergencyCryptoSwap, &weightMsgEmergencyCryptoSwap, nil,
		func(_ *rand.Rand) {
			weightMsgEmergencyCryptoSwap = defaultWeightMsgEmergencyCryptoSwap
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgEmergencyCryptoSwap,
		qcbqsecsimulation.SimulateMsgEmergencyCryptoSwap(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgRegisterQsKey,
			defaultWeightMsgRegisterQsKey,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				qcbqsecsimulation.SimulateMsgRegisterQsKey(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgRotateQsKey,
			defaultWeightMsgRotateQsKey,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				qcbqsecsimulation.SimulateMsgRotateQsKey(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgSubmitHybridSig,
			defaultWeightMsgSubmitHybridSig,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				qcbqsecsimulation.SimulateMsgSubmitHybridSig(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgSetThreatLevel,
			defaultWeightMsgSetThreatLevel,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				qcbqsecsimulation.SimulateMsgSetThreatLevel(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgRegisterAlgorithm,
			defaultWeightMsgRegisterAlgorithm,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				qcbqsecsimulation.SimulateMsgRegisterAlgorithm(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeprecateAlgorithm,
			defaultWeightMsgDeprecateAlgorithm,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				qcbqsecsimulation.SimulateMsgDeprecateAlgorithm(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgVerifyQsSignature,
			defaultWeightMsgVerifyQsSignature,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				qcbqsecsimulation.SimulateMsgVerifyQsSignature(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgEmergencyCryptoSwap,
			defaultWeightMsgEmergencyCryptoSwap,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				qcbqsecsimulation.SimulateMsgEmergencyCryptoSwap(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
