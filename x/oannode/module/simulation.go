package oannode

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"oan/testutil/sample"
	oannodesimulation "oan/x/oannode/simulation"
	"oan/x/oannode/types"
)

// avoid unused import issue
var (
	_ = oannodesimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgRegisterNode = "op_weight_msg_register_node"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRegisterNode int = 100

	opWeightMsgUpdateNode = "op_weight_msg_update_node"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateNode int = 100

	opWeightMsgClaimNodeReward = "op_weight_msg_claim_node_reward"
	// TODO: Determine the simulation weight value
	defaultWeightMsgClaimNodeReward int = 100

	opWeightMsgSlashNode = "op_weight_msg_slash_node"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSlashNode int = 100

	opWeightMsgDeregisterNode = "op_weight_msg_deregister_node"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeregisterNode int = 100

	opWeightMsgSetNodeConfig = "op_weight_msg_set_node_config"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSetNodeConfig int = 100

	opWeightMsgReportNode = "op_weight_msg_report_node"
	// TODO: Determine the simulation weight value
	defaultWeightMsgReportNode int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	oannodeGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&oannodeGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgRegisterNode int
	simState.AppParams.GetOrGenerate(opWeightMsgRegisterNode, &weightMsgRegisterNode, nil,
		func(_ *rand.Rand) {
			weightMsgRegisterNode = defaultWeightMsgRegisterNode
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRegisterNode,
		oannodesimulation.SimulateMsgRegisterNode(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateNode int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateNode, &weightMsgUpdateNode, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateNode = defaultWeightMsgUpdateNode
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateNode,
		oannodesimulation.SimulateMsgUpdateNode(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgClaimNodeReward int
	simState.AppParams.GetOrGenerate(opWeightMsgClaimNodeReward, &weightMsgClaimNodeReward, nil,
		func(_ *rand.Rand) {
			weightMsgClaimNodeReward = defaultWeightMsgClaimNodeReward
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgClaimNodeReward,
		oannodesimulation.SimulateMsgClaimNodeReward(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSlashNode int
	simState.AppParams.GetOrGenerate(opWeightMsgSlashNode, &weightMsgSlashNode, nil,
		func(_ *rand.Rand) {
			weightMsgSlashNode = defaultWeightMsgSlashNode
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSlashNode,
		oannodesimulation.SimulateMsgSlashNode(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeregisterNode int
	simState.AppParams.GetOrGenerate(opWeightMsgDeregisterNode, &weightMsgDeregisterNode, nil,
		func(_ *rand.Rand) {
			weightMsgDeregisterNode = defaultWeightMsgDeregisterNode
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeregisterNode,
		oannodesimulation.SimulateMsgDeregisterNode(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSetNodeConfig int
	simState.AppParams.GetOrGenerate(opWeightMsgSetNodeConfig, &weightMsgSetNodeConfig, nil,
		func(_ *rand.Rand) {
			weightMsgSetNodeConfig = defaultWeightMsgSetNodeConfig
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSetNodeConfig,
		oannodesimulation.SimulateMsgSetNodeConfig(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgReportNode int
	simState.AppParams.GetOrGenerate(opWeightMsgReportNode, &weightMsgReportNode, nil,
		func(_ *rand.Rand) {
			weightMsgReportNode = defaultWeightMsgReportNode
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgReportNode,
		oannodesimulation.SimulateMsgReportNode(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgRegisterNode,
			defaultWeightMsgRegisterNode,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				oannodesimulation.SimulateMsgRegisterNode(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateNode,
			defaultWeightMsgUpdateNode,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				oannodesimulation.SimulateMsgUpdateNode(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgClaimNodeReward,
			defaultWeightMsgClaimNodeReward,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				oannodesimulation.SimulateMsgClaimNodeReward(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgSlashNode,
			defaultWeightMsgSlashNode,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				oannodesimulation.SimulateMsgSlashNode(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeregisterNode,
			defaultWeightMsgDeregisterNode,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				oannodesimulation.SimulateMsgDeregisterNode(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgSetNodeConfig,
			defaultWeightMsgSetNodeConfig,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				oannodesimulation.SimulateMsgSetNodeConfig(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgReportNode,
			defaultWeightMsgReportNode,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				oannodesimulation.SimulateMsgReportNode(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
