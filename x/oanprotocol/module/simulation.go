package oanprotocol

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"oan/testutil/sample"
	oanprotocolsimulation "oan/x/oanprotocol/simulation"
	"oan/x/oanprotocol/types"
)

// avoid unused import issue
var (
	_ = oanprotocolsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgSetAddressTier = "op_weight_msg_set_address_tier"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSetAddressTier int = 100

	opWeightMsgUpdateLaunchPhase = "op_weight_msg_update_launch_phase"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateLaunchPhase int = 100

	opWeightMsgUpdateProtocolParams = "op_weight_msg_update_protocol_params"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateProtocolParams int = 100

	opWeightMsgExemptAddress = "op_weight_msg_exempt_address"
	// TODO: Determine the simulation weight value
	defaultWeightMsgExemptAddress int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	oanprotocolGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&oanprotocolGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgSetAddressTier int
	simState.AppParams.GetOrGenerate(opWeightMsgSetAddressTier, &weightMsgSetAddressTier, nil,
		func(_ *rand.Rand) {
			weightMsgSetAddressTier = defaultWeightMsgSetAddressTier
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSetAddressTier,
		oanprotocolsimulation.SimulateMsgSetAddressTier(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateLaunchPhase int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateLaunchPhase, &weightMsgUpdateLaunchPhase, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateLaunchPhase = defaultWeightMsgUpdateLaunchPhase
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateLaunchPhase,
		oanprotocolsimulation.SimulateMsgUpdateLaunchPhase(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateProtocolParams int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateProtocolParams, &weightMsgUpdateProtocolParams, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateProtocolParams = defaultWeightMsgUpdateProtocolParams
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateProtocolParams,
		oanprotocolsimulation.SimulateMsgUpdateProtocolParams(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgExemptAddress int
	simState.AppParams.GetOrGenerate(opWeightMsgExemptAddress, &weightMsgExemptAddress, nil,
		func(_ *rand.Rand) {
			weightMsgExemptAddress = defaultWeightMsgExemptAddress
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgExemptAddress,
		oanprotocolsimulation.SimulateMsgExemptAddress(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgSetAddressTier,
			defaultWeightMsgSetAddressTier,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				oanprotocolsimulation.SimulateMsgSetAddressTier(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateLaunchPhase,
			defaultWeightMsgUpdateLaunchPhase,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				oanprotocolsimulation.SimulateMsgUpdateLaunchPhase(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateProtocolParams,
			defaultWeightMsgUpdateProtocolParams,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				oanprotocolsimulation.SimulateMsgUpdateProtocolParams(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgExemptAddress,
			defaultWeightMsgExemptAddress,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				oanprotocolsimulation.SimulateMsgExemptAddress(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
