package qcbwalletproto

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"qcb/testutil/sample"
	qcbwalletprotosimulation "qcb/x/qcbwalletproto/simulation"
	"qcb/x/qcbwalletproto/types"
)

// avoid unused import issue
var (
	_ = qcbwalletprotosimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgRegisterWalletProfile = "op_weight_msg_register_wallet_profile"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRegisterWalletProfile int = 100

	opWeightMsgSetEncryptionKey = "op_weight_msg_set_encryption_key"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSetEncryptionKey int = 100

	opWeightMsgSetPqKey = "op_weight_msg_set_pq_key"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSetPqKey int = 100

	opWeightMsgUpdateWalletProfile = "op_weight_msg_update_wallet_profile"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateWalletProfile int = 100

	opWeightMsgSetWalletPermissions = "op_weight_msg_set_wallet_permissions"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSetWalletPermissions int = 100

	opWeightMsgLockWallet = "op_weight_msg_lock_wallet"
	// TODO: Determine the simulation weight value
	defaultWeightMsgLockWallet int = 100

	opWeightMsgUnlockWallet = "op_weight_msg_unlock_wallet"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUnlockWallet int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	qcbwalletprotoGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&qcbwalletprotoGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgRegisterWalletProfile int
	simState.AppParams.GetOrGenerate(opWeightMsgRegisterWalletProfile, &weightMsgRegisterWalletProfile, nil,
		func(_ *rand.Rand) {
			weightMsgRegisterWalletProfile = defaultWeightMsgRegisterWalletProfile
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRegisterWalletProfile,
		qcbwalletprotosimulation.SimulateMsgRegisterWalletProfile(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSetEncryptionKey int
	simState.AppParams.GetOrGenerate(opWeightMsgSetEncryptionKey, &weightMsgSetEncryptionKey, nil,
		func(_ *rand.Rand) {
			weightMsgSetEncryptionKey = defaultWeightMsgSetEncryptionKey
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSetEncryptionKey,
		qcbwalletprotosimulation.SimulateMsgSetEncryptionKey(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSetPqKey int
	simState.AppParams.GetOrGenerate(opWeightMsgSetPqKey, &weightMsgSetPqKey, nil,
		func(_ *rand.Rand) {
			weightMsgSetPqKey = defaultWeightMsgSetPqKey
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSetPqKey,
		qcbwalletprotosimulation.SimulateMsgSetPqKey(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateWalletProfile int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateWalletProfile, &weightMsgUpdateWalletProfile, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateWalletProfile = defaultWeightMsgUpdateWalletProfile
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateWalletProfile,
		qcbwalletprotosimulation.SimulateMsgUpdateWalletProfile(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSetWalletPermissions int
	simState.AppParams.GetOrGenerate(opWeightMsgSetWalletPermissions, &weightMsgSetWalletPermissions, nil,
		func(_ *rand.Rand) {
			weightMsgSetWalletPermissions = defaultWeightMsgSetWalletPermissions
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSetWalletPermissions,
		qcbwalletprotosimulation.SimulateMsgSetWalletPermissions(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgLockWallet int
	simState.AppParams.GetOrGenerate(opWeightMsgLockWallet, &weightMsgLockWallet, nil,
		func(_ *rand.Rand) {
			weightMsgLockWallet = defaultWeightMsgLockWallet
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgLockWallet,
		qcbwalletprotosimulation.SimulateMsgLockWallet(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUnlockWallet int
	simState.AppParams.GetOrGenerate(opWeightMsgUnlockWallet, &weightMsgUnlockWallet, nil,
		func(_ *rand.Rand) {
			weightMsgUnlockWallet = defaultWeightMsgUnlockWallet
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUnlockWallet,
		qcbwalletprotosimulation.SimulateMsgUnlockWallet(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgRegisterWalletProfile,
			defaultWeightMsgRegisterWalletProfile,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				qcbwalletprotosimulation.SimulateMsgRegisterWalletProfile(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgSetEncryptionKey,
			defaultWeightMsgSetEncryptionKey,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				qcbwalletprotosimulation.SimulateMsgSetEncryptionKey(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgSetPqKey,
			defaultWeightMsgSetPqKey,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				qcbwalletprotosimulation.SimulateMsgSetPqKey(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateWalletProfile,
			defaultWeightMsgUpdateWalletProfile,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				qcbwalletprotosimulation.SimulateMsgUpdateWalletProfile(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgSetWalletPermissions,
			defaultWeightMsgSetWalletPermissions,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				qcbwalletprotosimulation.SimulateMsgSetWalletPermissions(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgLockWallet,
			defaultWeightMsgLockWallet,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				qcbwalletprotosimulation.SimulateMsgLockWallet(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUnlockWallet,
			defaultWeightMsgUnlockWallet,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				qcbwalletprotosimulation.SimulateMsgUnlockWallet(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
