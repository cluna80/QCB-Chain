package oanbridge

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"oan/testutil/sample"
	oanbridgesimulation "oan/x/oanbridge/simulation"
	"oan/x/oanbridge/types"
)

// avoid unused import issue
var (
	_ = oanbridgesimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgSendIbcAgent = "op_weight_msg_send_ibc_agent"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSendIbcAgent int = 100

	opWeightMsgBroadcastAgentState = "op_weight_msg_broadcast_agent_state"
	// TODO: Determine the simulation weight value
	defaultWeightMsgBroadcastAgentState int = 100

	opWeightMsgPostStateRoot = "op_weight_msg_post_state_root"
	// TODO: Determine the simulation weight value
	defaultWeightMsgPostStateRoot int = 100

	opWeightMsgRegisterChain = "op_weight_msg_register_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRegisterChain int = 100

	opWeightMsgTokenizeOutput = "op_weight_msg_tokenize_output"
	// TODO: Determine the simulation weight value
	defaultWeightMsgTokenizeOutput int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	oanbridgeGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&oanbridgeGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgSendIbcAgent int
	simState.AppParams.GetOrGenerate(opWeightMsgSendIbcAgent, &weightMsgSendIbcAgent, nil,
		func(_ *rand.Rand) {
			weightMsgSendIbcAgent = defaultWeightMsgSendIbcAgent
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSendIbcAgent,
		oanbridgesimulation.SimulateMsgSendIbcAgent(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgBroadcastAgentState int
	simState.AppParams.GetOrGenerate(opWeightMsgBroadcastAgentState, &weightMsgBroadcastAgentState, nil,
		func(_ *rand.Rand) {
			weightMsgBroadcastAgentState = defaultWeightMsgBroadcastAgentState
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgBroadcastAgentState,
		oanbridgesimulation.SimulateMsgBroadcastAgentState(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgPostStateRoot int
	simState.AppParams.GetOrGenerate(opWeightMsgPostStateRoot, &weightMsgPostStateRoot, nil,
		func(_ *rand.Rand) {
			weightMsgPostStateRoot = defaultWeightMsgPostStateRoot
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgPostStateRoot,
		oanbridgesimulation.SimulateMsgPostStateRoot(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRegisterChain int
	simState.AppParams.GetOrGenerate(opWeightMsgRegisterChain, &weightMsgRegisterChain, nil,
		func(_ *rand.Rand) {
			weightMsgRegisterChain = defaultWeightMsgRegisterChain
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRegisterChain,
		oanbridgesimulation.SimulateMsgRegisterChain(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgTokenizeOutput int
	simState.AppParams.GetOrGenerate(opWeightMsgTokenizeOutput, &weightMsgTokenizeOutput, nil,
		func(_ *rand.Rand) {
			weightMsgTokenizeOutput = defaultWeightMsgTokenizeOutput
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgTokenizeOutput,
		oanbridgesimulation.SimulateMsgTokenizeOutput(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgSendIbcAgent,
			defaultWeightMsgSendIbcAgent,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				oanbridgesimulation.SimulateMsgSendIbcAgent(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgBroadcastAgentState,
			defaultWeightMsgBroadcastAgentState,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				oanbridgesimulation.SimulateMsgBroadcastAgentState(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgPostStateRoot,
			defaultWeightMsgPostStateRoot,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				oanbridgesimulation.SimulateMsgPostStateRoot(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgRegisterChain,
			defaultWeightMsgRegisterChain,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				oanbridgesimulation.SimulateMsgRegisterChain(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgTokenizeOutput,
			defaultWeightMsgTokenizeOutput,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				oanbridgesimulation.SimulateMsgTokenizeOutput(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
