package oaneconomy

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"oan/testutil/sample"
	oaneconomysimulation "oan/x/oaneconomy/simulation"
	"oan/x/oaneconomy/types"
)

// avoid unused import issue
var (
	_ = oaneconomysimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateTask = "op_weight_msg_create_task"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateTask int = 100

	opWeightMsgCompleteTask = "op_weight_msg_complete_task"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCompleteTask int = 100

	opWeightMsgAcceptTask = "op_weight_msg_accept_task"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAcceptTask int = 100

	opWeightMsgDisputeTask = "op_weight_msg_dispute_task"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDisputeTask int = 100

	opWeightMsgClaimUbi = "op_weight_msg_claim_ubi"
	// TODO: Determine the simulation weight value
	defaultWeightMsgClaimUbi int = 100

	opWeightMsgStakeTokens = "op_weight_msg_stake_tokens"
	// TODO: Determine the simulation weight value
	defaultWeightMsgStakeTokens int = 100

	opWeightMsgDistributeRewards = "op_weight_msg_distribute_rewards"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDistributeRewards int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	oaneconomyGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&oaneconomyGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateTask int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateTask, &weightMsgCreateTask, nil,
		func(_ *rand.Rand) {
			weightMsgCreateTask = defaultWeightMsgCreateTask
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateTask,
		oaneconomysimulation.SimulateMsgCreateTask(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCompleteTask int
	simState.AppParams.GetOrGenerate(opWeightMsgCompleteTask, &weightMsgCompleteTask, nil,
		func(_ *rand.Rand) {
			weightMsgCompleteTask = defaultWeightMsgCompleteTask
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCompleteTask,
		oaneconomysimulation.SimulateMsgCompleteTask(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAcceptTask int
	simState.AppParams.GetOrGenerate(opWeightMsgAcceptTask, &weightMsgAcceptTask, nil,
		func(_ *rand.Rand) {
			weightMsgAcceptTask = defaultWeightMsgAcceptTask
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAcceptTask,
		oaneconomysimulation.SimulateMsgAcceptTask(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDisputeTask int
	simState.AppParams.GetOrGenerate(opWeightMsgDisputeTask, &weightMsgDisputeTask, nil,
		func(_ *rand.Rand) {
			weightMsgDisputeTask = defaultWeightMsgDisputeTask
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDisputeTask,
		oaneconomysimulation.SimulateMsgDisputeTask(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgClaimUbi int
	simState.AppParams.GetOrGenerate(opWeightMsgClaimUbi, &weightMsgClaimUbi, nil,
		func(_ *rand.Rand) {
			weightMsgClaimUbi = defaultWeightMsgClaimUbi
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgClaimUbi,
		oaneconomysimulation.SimulateMsgClaimUbi(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgStakeTokens int
	simState.AppParams.GetOrGenerate(opWeightMsgStakeTokens, &weightMsgStakeTokens, nil,
		func(_ *rand.Rand) {
			weightMsgStakeTokens = defaultWeightMsgStakeTokens
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgStakeTokens,
		oaneconomysimulation.SimulateMsgStakeTokens(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDistributeRewards int
	simState.AppParams.GetOrGenerate(opWeightMsgDistributeRewards, &weightMsgDistributeRewards, nil,
		func(_ *rand.Rand) {
			weightMsgDistributeRewards = defaultWeightMsgDistributeRewards
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDistributeRewards,
		oaneconomysimulation.SimulateMsgDistributeRewards(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateTask,
			defaultWeightMsgCreateTask,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				oaneconomysimulation.SimulateMsgCreateTask(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCompleteTask,
			defaultWeightMsgCompleteTask,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				oaneconomysimulation.SimulateMsgCompleteTask(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgAcceptTask,
			defaultWeightMsgAcceptTask,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				oaneconomysimulation.SimulateMsgAcceptTask(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDisputeTask,
			defaultWeightMsgDisputeTask,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				oaneconomysimulation.SimulateMsgDisputeTask(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgClaimUbi,
			defaultWeightMsgClaimUbi,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				oaneconomysimulation.SimulateMsgClaimUbi(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgStakeTokens,
			defaultWeightMsgStakeTokens,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				oaneconomysimulation.SimulateMsgStakeTokens(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDistributeRewards,
			defaultWeightMsgDistributeRewards,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				oaneconomysimulation.SimulateMsgDistributeRewards(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
