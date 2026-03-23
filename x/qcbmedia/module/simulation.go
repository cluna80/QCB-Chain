package qcbmedia

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"qcb/testutil/sample"
	qcbmediasimulation "qcb/x/qcbmedia/simulation"
	"qcb/x/qcbmedia/types"
)

// avoid unused import issue
var (
	_ = qcbmediasimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateMediaNft = "op_weight_msg_create_media_nft"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateMediaNft int = 100

	opWeightMsgClaimRoyalty = "op_weight_msg_claim_royalty"
	// TODO: Determine the simulation weight value
	defaultWeightMsgClaimRoyalty int = 100

	opWeightMsgLicenseContent = "op_weight_msg_license_content"
	// TODO: Determine the simulation weight value
	defaultWeightMsgLicenseContent int = 100

	opWeightMsgRecordView = "op_weight_msg_record_view"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRecordView int = 100

	opWeightMsgMintMusicNft = "op_weight_msg_mint_music_nft"
	// TODO: Determine the simulation weight value
	defaultWeightMsgMintMusicNft int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	qcbmediaGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&qcbmediaGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateMediaNft int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateMediaNft, &weightMsgCreateMediaNft, nil,
		func(_ *rand.Rand) {
			weightMsgCreateMediaNft = defaultWeightMsgCreateMediaNft
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateMediaNft,
		qcbmediasimulation.SimulateMsgCreateMediaNft(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgClaimRoyalty int
	simState.AppParams.GetOrGenerate(opWeightMsgClaimRoyalty, &weightMsgClaimRoyalty, nil,
		func(_ *rand.Rand) {
			weightMsgClaimRoyalty = defaultWeightMsgClaimRoyalty
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgClaimRoyalty,
		qcbmediasimulation.SimulateMsgClaimRoyalty(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgLicenseContent int
	simState.AppParams.GetOrGenerate(opWeightMsgLicenseContent, &weightMsgLicenseContent, nil,
		func(_ *rand.Rand) {
			weightMsgLicenseContent = defaultWeightMsgLicenseContent
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgLicenseContent,
		qcbmediasimulation.SimulateMsgLicenseContent(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRecordView int
	simState.AppParams.GetOrGenerate(opWeightMsgRecordView, &weightMsgRecordView, nil,
		func(_ *rand.Rand) {
			weightMsgRecordView = defaultWeightMsgRecordView
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRecordView,
		qcbmediasimulation.SimulateMsgRecordView(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgMintMusicNft int
	simState.AppParams.GetOrGenerate(opWeightMsgMintMusicNft, &weightMsgMintMusicNft, nil,
		func(_ *rand.Rand) {
			weightMsgMintMusicNft = defaultWeightMsgMintMusicNft
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgMintMusicNft,
		qcbmediasimulation.SimulateMsgMintMusicNft(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateMediaNft,
			defaultWeightMsgCreateMediaNft,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				qcbmediasimulation.SimulateMsgCreateMediaNft(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgClaimRoyalty,
			defaultWeightMsgClaimRoyalty,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				qcbmediasimulation.SimulateMsgClaimRoyalty(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgLicenseContent,
			defaultWeightMsgLicenseContent,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				qcbmediasimulation.SimulateMsgLicenseContent(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgRecordView,
			defaultWeightMsgRecordView,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				qcbmediasimulation.SimulateMsgRecordView(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgMintMusicNft,
			defaultWeightMsgMintMusicNft,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				qcbmediasimulation.SimulateMsgMintMusicNft(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
