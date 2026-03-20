package oandao

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"oan/testutil/sample"
	oandaosimulation "oan/x/oandao/simulation"
	"oan/x/oandao/types"
)

// avoid unused import issue
var (
	_ = oandaosimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgSubmitProposal = "op_weight_msg_submit_proposal"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSubmitProposal int = 100

	opWeightMsgVoteProposal = "op_weight_msg_vote_proposal"
	// TODO: Determine the simulation weight value
	defaultWeightMsgVoteProposal int = 100

	opWeightMsgExecuteProposal = "op_weight_msg_execute_proposal"
	// TODO: Determine the simulation weight value
	defaultWeightMsgExecuteProposal int = 100

	opWeightMsgVetoProposal = "op_weight_msg_veto_proposal"
	// TODO: Determine the simulation weight value
	defaultWeightMsgVetoProposal int = 100

	opWeightMsgDelegateVote = "op_weight_msg_delegate_vote"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDelegateVote int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	oandaoGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&oandaoGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgSubmitProposal int
	simState.AppParams.GetOrGenerate(opWeightMsgSubmitProposal, &weightMsgSubmitProposal, nil,
		func(_ *rand.Rand) {
			weightMsgSubmitProposal = defaultWeightMsgSubmitProposal
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSubmitProposal,
		oandaosimulation.SimulateMsgSubmitProposal(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgVoteProposal int
	simState.AppParams.GetOrGenerate(opWeightMsgVoteProposal, &weightMsgVoteProposal, nil,
		func(_ *rand.Rand) {
			weightMsgVoteProposal = defaultWeightMsgVoteProposal
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgVoteProposal,
		oandaosimulation.SimulateMsgVoteProposal(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgExecuteProposal int
	simState.AppParams.GetOrGenerate(opWeightMsgExecuteProposal, &weightMsgExecuteProposal, nil,
		func(_ *rand.Rand) {
			weightMsgExecuteProposal = defaultWeightMsgExecuteProposal
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgExecuteProposal,
		oandaosimulation.SimulateMsgExecuteProposal(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgVetoProposal int
	simState.AppParams.GetOrGenerate(opWeightMsgVetoProposal, &weightMsgVetoProposal, nil,
		func(_ *rand.Rand) {
			weightMsgVetoProposal = defaultWeightMsgVetoProposal
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgVetoProposal,
		oandaosimulation.SimulateMsgVetoProposal(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDelegateVote int
	simState.AppParams.GetOrGenerate(opWeightMsgDelegateVote, &weightMsgDelegateVote, nil,
		func(_ *rand.Rand) {
			weightMsgDelegateVote = defaultWeightMsgDelegateVote
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDelegateVote,
		oandaosimulation.SimulateMsgDelegateVote(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgSubmitProposal,
			defaultWeightMsgSubmitProposal,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				oandaosimulation.SimulateMsgSubmitProposal(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgVoteProposal,
			defaultWeightMsgVoteProposal,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				oandaosimulation.SimulateMsgVoteProposal(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgExecuteProposal,
			defaultWeightMsgExecuteProposal,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				oandaosimulation.SimulateMsgExecuteProposal(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgVetoProposal,
			defaultWeightMsgVetoProposal,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				oandaosimulation.SimulateMsgVetoProposal(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDelegateVote,
			defaultWeightMsgDelegateVote,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				oandaosimulation.SimulateMsgDelegateVote(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
