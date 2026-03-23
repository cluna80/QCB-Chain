package qcbcomms

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"qcb/testutil/sample"
	qcbcommssimulation "qcb/x/qcbcomms/simulation"
	"qcb/x/qcbcomms/types"
)

// avoid unused import issue
var (
	_ = qcbcommssimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgRegisterMsgKey = "op_weight_msg_register_msg_key"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRegisterMsgKey int = 100

	opWeightMsgSendMsgHeader = "op_weight_msg_send_msg_header"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSendMsgHeader int = 100

	opWeightMsgAckMsg = "op_weight_msg_ack_msg"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAckMsg int = 100

	opWeightMsgRevokeMsgKey = "op_weight_msg_revoke_msg_key"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRevokeMsgKey int = 100

	opWeightMsgSetMsgPolicy = "op_weight_msg_set_msg_policy"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSetMsgPolicy int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	qcbcommsGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&qcbcommsGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgRegisterMsgKey int
	simState.AppParams.GetOrGenerate(opWeightMsgRegisterMsgKey, &weightMsgRegisterMsgKey, nil,
		func(_ *rand.Rand) {
			weightMsgRegisterMsgKey = defaultWeightMsgRegisterMsgKey
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRegisterMsgKey,
		qcbcommssimulation.SimulateMsgRegisterMsgKey(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSendMsgHeader int
	simState.AppParams.GetOrGenerate(opWeightMsgSendMsgHeader, &weightMsgSendMsgHeader, nil,
		func(_ *rand.Rand) {
			weightMsgSendMsgHeader = defaultWeightMsgSendMsgHeader
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSendMsgHeader,
		qcbcommssimulation.SimulateMsgSendMsgHeader(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAckMsg int
	simState.AppParams.GetOrGenerate(opWeightMsgAckMsg, &weightMsgAckMsg, nil,
		func(_ *rand.Rand) {
			weightMsgAckMsg = defaultWeightMsgAckMsg
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAckMsg,
		qcbcommssimulation.SimulateMsgAckMsg(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRevokeMsgKey int
	simState.AppParams.GetOrGenerate(opWeightMsgRevokeMsgKey, &weightMsgRevokeMsgKey, nil,
		func(_ *rand.Rand) {
			weightMsgRevokeMsgKey = defaultWeightMsgRevokeMsgKey
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRevokeMsgKey,
		qcbcommssimulation.SimulateMsgRevokeMsgKey(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSetMsgPolicy int
	simState.AppParams.GetOrGenerate(opWeightMsgSetMsgPolicy, &weightMsgSetMsgPolicy, nil,
		func(_ *rand.Rand) {
			weightMsgSetMsgPolicy = defaultWeightMsgSetMsgPolicy
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSetMsgPolicy,
		qcbcommssimulation.SimulateMsgSetMsgPolicy(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgRegisterMsgKey,
			defaultWeightMsgRegisterMsgKey,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				qcbcommssimulation.SimulateMsgRegisterMsgKey(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgSendMsgHeader,
			defaultWeightMsgSendMsgHeader,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				qcbcommssimulation.SimulateMsgSendMsgHeader(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgAckMsg,
			defaultWeightMsgAckMsg,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				qcbcommssimulation.SimulateMsgAckMsg(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgRevokeMsgKey,
			defaultWeightMsgRevokeMsgKey,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				qcbcommssimulation.SimulateMsgRevokeMsgKey(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgSetMsgPolicy,
			defaultWeightMsgSetMsgPolicy,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				qcbcommssimulation.SimulateMsgSetMsgPolicy(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
