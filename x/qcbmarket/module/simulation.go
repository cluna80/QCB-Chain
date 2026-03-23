package qcbmarket

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"qcb/testutil/sample"
	qcbmarketsimulation "qcb/x/qcbmarket/simulation"
	"qcb/x/qcbmarket/types"
)

// avoid unused import issue
var (
	_ = qcbmarketsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgListAgentForHire = "op_weight_msg_list_agent_for_hire"
	// TODO: Determine the simulation weight value
	defaultWeightMsgListAgentForHire int = 100

	opWeightMsgHireAgent = "op_weight_msg_hire_agent"
	// TODO: Determine the simulation weight value
	defaultWeightMsgHireAgent int = 100

	opWeightMsgRateAgent = "op_weight_msg_rate_agent"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRateAgent int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	qcbmarketGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&qcbmarketGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgListAgentForHire int
	simState.AppParams.GetOrGenerate(opWeightMsgListAgentForHire, &weightMsgListAgentForHire, nil,
		func(_ *rand.Rand) {
			weightMsgListAgentForHire = defaultWeightMsgListAgentForHire
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgListAgentForHire,
		qcbmarketsimulation.SimulateMsgListAgentForHire(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgHireAgent int
	simState.AppParams.GetOrGenerate(opWeightMsgHireAgent, &weightMsgHireAgent, nil,
		func(_ *rand.Rand) {
			weightMsgHireAgent = defaultWeightMsgHireAgent
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgHireAgent,
		qcbmarketsimulation.SimulateMsgHireAgent(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRateAgent int
	simState.AppParams.GetOrGenerate(opWeightMsgRateAgent, &weightMsgRateAgent, nil,
		func(_ *rand.Rand) {
			weightMsgRateAgent = defaultWeightMsgRateAgent
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRateAgent,
		qcbmarketsimulation.SimulateMsgRateAgent(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgListAgentForHire,
			defaultWeightMsgListAgentForHire,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				qcbmarketsimulation.SimulateMsgListAgentForHire(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgHireAgent,
			defaultWeightMsgHireAgent,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				qcbmarketsimulation.SimulateMsgHireAgent(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgRateAgent,
			defaultWeightMsgRateAgent,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				qcbmarketsimulation.SimulateMsgRateAgent(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
