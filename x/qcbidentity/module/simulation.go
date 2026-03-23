package qcbidentity

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"qcb/testutil/sample"
	qcbidentitysimulation "qcb/x/qcbidentity/simulation"
	"qcb/x/qcbidentity/types"
)

// avoid unused import issue
var (
	_ = qcbidentitysimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgRegisterIdentity = "op_weight_msg_register_identity"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRegisterIdentity int = 100

	opWeightMsgVerifyIdentity = "op_weight_msg_verify_identity"
	// TODO: Determine the simulation weight value
	defaultWeightMsgVerifyIdentity int = 100

	opWeightMsgUpdateReputation = "op_weight_msg_update_reputation"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateReputation int = 100

	opWeightMsgLinkWallet = "op_weight_msg_link_wallet"
	// TODO: Determine the simulation weight value
	defaultWeightMsgLinkWallet int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	qcbidentityGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&qcbidentityGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgRegisterIdentity int
	simState.AppParams.GetOrGenerate(opWeightMsgRegisterIdentity, &weightMsgRegisterIdentity, nil,
		func(_ *rand.Rand) {
			weightMsgRegisterIdentity = defaultWeightMsgRegisterIdentity
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRegisterIdentity,
		qcbidentitysimulation.SimulateMsgRegisterIdentity(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgVerifyIdentity int
	simState.AppParams.GetOrGenerate(opWeightMsgVerifyIdentity, &weightMsgVerifyIdentity, nil,
		func(_ *rand.Rand) {
			weightMsgVerifyIdentity = defaultWeightMsgVerifyIdentity
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgVerifyIdentity,
		qcbidentitysimulation.SimulateMsgVerifyIdentity(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateReputation int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateReputation, &weightMsgUpdateReputation, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateReputation = defaultWeightMsgUpdateReputation
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateReputation,
		qcbidentitysimulation.SimulateMsgUpdateReputation(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgLinkWallet int
	simState.AppParams.GetOrGenerate(opWeightMsgLinkWallet, &weightMsgLinkWallet, nil,
		func(_ *rand.Rand) {
			weightMsgLinkWallet = defaultWeightMsgLinkWallet
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgLinkWallet,
		qcbidentitysimulation.SimulateMsgLinkWallet(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgRegisterIdentity,
			defaultWeightMsgRegisterIdentity,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				qcbidentitysimulation.SimulateMsgRegisterIdentity(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgVerifyIdentity,
			defaultWeightMsgVerifyIdentity,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				qcbidentitysimulation.SimulateMsgVerifyIdentity(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateReputation,
			defaultWeightMsgUpdateReputation,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				qcbidentitysimulation.SimulateMsgUpdateReputation(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgLinkWallet,
			defaultWeightMsgLinkWallet,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				qcbidentitysimulation.SimulateMsgLinkWallet(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
