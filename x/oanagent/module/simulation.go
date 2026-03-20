package oanagent

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"oan/testutil/sample"
	oanagentsimulation "oan/x/oanagent/simulation"
	"oan/x/oanagent/types"
)

// avoid unused import issue
var (
	_ = oanagentsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgRegisterAgent = "op_weight_msg_register_agent"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRegisterAgent int = 100

	opWeightMsgRecordTrade = "op_weight_msg_record_trade"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRecordTrade int = 100

	opWeightMsgBreedAgent = "op_weight_msg_breed_agent"
	// TODO: Determine the simulation weight value
	defaultWeightMsgBreedAgent int = 100

	opWeightMsgChallengeAgent = "op_weight_msg_challenge_agent"
	// TODO: Determine the simulation weight value
	defaultWeightMsgChallengeAgent int = 100

	opWeightMsgRetireAgent = "op_weight_msg_retire_agent"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRetireAgent int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	oanagentGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&oanagentGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgRegisterAgent int
	simState.AppParams.GetOrGenerate(opWeightMsgRegisterAgent, &weightMsgRegisterAgent, nil,
		func(_ *rand.Rand) {
			weightMsgRegisterAgent = defaultWeightMsgRegisterAgent
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRegisterAgent,
		oanagentsimulation.SimulateMsgRegisterAgent(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRecordTrade int
	simState.AppParams.GetOrGenerate(opWeightMsgRecordTrade, &weightMsgRecordTrade, nil,
		func(_ *rand.Rand) {
			weightMsgRecordTrade = defaultWeightMsgRecordTrade
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRecordTrade,
		oanagentsimulation.SimulateMsgRecordTrade(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgBreedAgent int
	simState.AppParams.GetOrGenerate(opWeightMsgBreedAgent, &weightMsgBreedAgent, nil,
		func(_ *rand.Rand) {
			weightMsgBreedAgent = defaultWeightMsgBreedAgent
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgBreedAgent,
		oanagentsimulation.SimulateMsgBreedAgent(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgChallengeAgent int
	simState.AppParams.GetOrGenerate(opWeightMsgChallengeAgent, &weightMsgChallengeAgent, nil,
		func(_ *rand.Rand) {
			weightMsgChallengeAgent = defaultWeightMsgChallengeAgent
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgChallengeAgent,
		oanagentsimulation.SimulateMsgChallengeAgent(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRetireAgent int
	simState.AppParams.GetOrGenerate(opWeightMsgRetireAgent, &weightMsgRetireAgent, nil,
		func(_ *rand.Rand) {
			weightMsgRetireAgent = defaultWeightMsgRetireAgent
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRetireAgent,
		oanagentsimulation.SimulateMsgRetireAgent(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgRegisterAgent,
			defaultWeightMsgRegisterAgent,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				oanagentsimulation.SimulateMsgRegisterAgent(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgRecordTrade,
			defaultWeightMsgRecordTrade,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				oanagentsimulation.SimulateMsgRecordTrade(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgBreedAgent,
			defaultWeightMsgBreedAgent,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				oanagentsimulation.SimulateMsgBreedAgent(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgChallengeAgent,
			defaultWeightMsgChallengeAgent,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				oanagentsimulation.SimulateMsgChallengeAgent(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgRetireAgent,
			defaultWeightMsgRetireAgent,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				oanagentsimulation.SimulateMsgRetireAgent(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
