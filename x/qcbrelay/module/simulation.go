package qcbrelay

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"qcb/testutil/sample"
	qcbrelaysimulation "qcb/x/qcbrelay/simulation"
	"qcb/x/qcbrelay/types"
)

// avoid unused import issue
var (
	_ = qcbrelaysimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgRegisterRelay = "op_weight_msg_register_relay"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRegisterRelay int = 100

	opWeightMsgRelayHeartbeat = "op_weight_msg_relay_heartbeat"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRelayHeartbeat int = 100

	opWeightMsgRouteMsg = "op_weight_msg_route_msg"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRouteMsg int = 100

	opWeightMsgSlashRelay = "op_weight_msg_slash_relay"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSlashRelay int = 100

	opWeightMsgRemoveRelay = "op_weight_msg_remove_relay"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRemoveRelay int = 100

	opWeightMsgUpdateRelayRegion = "op_weight_msg_update_relay_region"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateRelayRegion int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	qcbrelayGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&qcbrelayGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgRegisterRelay int
	simState.AppParams.GetOrGenerate(opWeightMsgRegisterRelay, &weightMsgRegisterRelay, nil,
		func(_ *rand.Rand) {
			weightMsgRegisterRelay = defaultWeightMsgRegisterRelay
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRegisterRelay,
		qcbrelaysimulation.SimulateMsgRegisterRelay(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRelayHeartbeat int
	simState.AppParams.GetOrGenerate(opWeightMsgRelayHeartbeat, &weightMsgRelayHeartbeat, nil,
		func(_ *rand.Rand) {
			weightMsgRelayHeartbeat = defaultWeightMsgRelayHeartbeat
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRelayHeartbeat,
		qcbrelaysimulation.SimulateMsgRelayHeartbeat(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRouteMsg int
	simState.AppParams.GetOrGenerate(opWeightMsgRouteMsg, &weightMsgRouteMsg, nil,
		func(_ *rand.Rand) {
			weightMsgRouteMsg = defaultWeightMsgRouteMsg
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRouteMsg,
		qcbrelaysimulation.SimulateMsgRouteMsg(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSlashRelay int
	simState.AppParams.GetOrGenerate(opWeightMsgSlashRelay, &weightMsgSlashRelay, nil,
		func(_ *rand.Rand) {
			weightMsgSlashRelay = defaultWeightMsgSlashRelay
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSlashRelay,
		qcbrelaysimulation.SimulateMsgSlashRelay(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRemoveRelay int
	simState.AppParams.GetOrGenerate(opWeightMsgRemoveRelay, &weightMsgRemoveRelay, nil,
		func(_ *rand.Rand) {
			weightMsgRemoveRelay = defaultWeightMsgRemoveRelay
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRemoveRelay,
		qcbrelaysimulation.SimulateMsgRemoveRelay(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateRelayRegion int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateRelayRegion, &weightMsgUpdateRelayRegion, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateRelayRegion = defaultWeightMsgUpdateRelayRegion
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateRelayRegion,
		qcbrelaysimulation.SimulateMsgUpdateRelayRegion(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgRegisterRelay,
			defaultWeightMsgRegisterRelay,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				qcbrelaysimulation.SimulateMsgRegisterRelay(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgRelayHeartbeat,
			defaultWeightMsgRelayHeartbeat,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				qcbrelaysimulation.SimulateMsgRelayHeartbeat(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgRouteMsg,
			defaultWeightMsgRouteMsg,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				qcbrelaysimulation.SimulateMsgRouteMsg(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgSlashRelay,
			defaultWeightMsgSlashRelay,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				qcbrelaysimulation.SimulateMsgSlashRelay(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgRemoveRelay,
			defaultWeightMsgRemoveRelay,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				qcbrelaysimulation.SimulateMsgRemoveRelay(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateRelayRegion,
			defaultWeightMsgUpdateRelayRegion,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				qcbrelaysimulation.SimulateMsgUpdateRelayRegion(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
