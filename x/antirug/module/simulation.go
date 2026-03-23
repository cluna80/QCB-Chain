package antirug

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"qcb/testutil/sample"
	antirugsimulation "qcb/x/antirug/simulation"
	"qcb/x/antirug/types"
)

// avoid unused import issue
var (
	_ = antirugsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgRegisterToken = "op_weight_msg_register_token"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRegisterToken int = 100

	opWeightMsgLockLiquidity = "op_weight_msg_lock_liquidity"
	// TODO: Determine the simulation weight value
	defaultWeightMsgLockLiquidity int = 100

	opWeightMsgDeclareMintLimit = "op_weight_msg_declare_mint_limit"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeclareMintLimit int = 100

	opWeightMsgFlagToken = "op_weight_msg_flag_token"
	// TODO: Determine the simulation weight value
	defaultWeightMsgFlagToken int = 100

	opWeightMsgFreezeToken = "op_weight_msg_freeze_token"
	// TODO: Determine the simulation weight value
	defaultWeightMsgFreezeToken int = 100

	opWeightMsgUnfreezeToken = "op_weight_msg_unfreeze_token"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUnfreezeToken int = 100

	opWeightMsgRequestUpgrade = "op_weight_msg_request_upgrade"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRequestUpgrade int = 100

	opWeightMsgApproveToken = "op_weight_msg_approve_token"
	// TODO: Determine the simulation weight value
	defaultWeightMsgApproveToken int = 100

	opWeightMsgSetModuleActive = "op_weight_msg_set_module_active"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSetModuleActive int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	antirugGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&antirugGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgRegisterToken int
	simState.AppParams.GetOrGenerate(opWeightMsgRegisterToken, &weightMsgRegisterToken, nil,
		func(_ *rand.Rand) {
			weightMsgRegisterToken = defaultWeightMsgRegisterToken
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRegisterToken,
		antirugsimulation.SimulateMsgRegisterToken(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgLockLiquidity int
	simState.AppParams.GetOrGenerate(opWeightMsgLockLiquidity, &weightMsgLockLiquidity, nil,
		func(_ *rand.Rand) {
			weightMsgLockLiquidity = defaultWeightMsgLockLiquidity
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgLockLiquidity,
		antirugsimulation.SimulateMsgLockLiquidity(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeclareMintLimit int
	simState.AppParams.GetOrGenerate(opWeightMsgDeclareMintLimit, &weightMsgDeclareMintLimit, nil,
		func(_ *rand.Rand) {
			weightMsgDeclareMintLimit = defaultWeightMsgDeclareMintLimit
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeclareMintLimit,
		antirugsimulation.SimulateMsgDeclareMintLimit(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgFlagToken int
	simState.AppParams.GetOrGenerate(opWeightMsgFlagToken, &weightMsgFlagToken, nil,
		func(_ *rand.Rand) {
			weightMsgFlagToken = defaultWeightMsgFlagToken
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgFlagToken,
		antirugsimulation.SimulateMsgFlagToken(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgFreezeToken int
	simState.AppParams.GetOrGenerate(opWeightMsgFreezeToken, &weightMsgFreezeToken, nil,
		func(_ *rand.Rand) {
			weightMsgFreezeToken = defaultWeightMsgFreezeToken
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgFreezeToken,
		antirugsimulation.SimulateMsgFreezeToken(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUnfreezeToken int
	simState.AppParams.GetOrGenerate(opWeightMsgUnfreezeToken, &weightMsgUnfreezeToken, nil,
		func(_ *rand.Rand) {
			weightMsgUnfreezeToken = defaultWeightMsgUnfreezeToken
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUnfreezeToken,
		antirugsimulation.SimulateMsgUnfreezeToken(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRequestUpgrade int
	simState.AppParams.GetOrGenerate(opWeightMsgRequestUpgrade, &weightMsgRequestUpgrade, nil,
		func(_ *rand.Rand) {
			weightMsgRequestUpgrade = defaultWeightMsgRequestUpgrade
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRequestUpgrade,
		antirugsimulation.SimulateMsgRequestUpgrade(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgApproveToken int
	simState.AppParams.GetOrGenerate(opWeightMsgApproveToken, &weightMsgApproveToken, nil,
		func(_ *rand.Rand) {
			weightMsgApproveToken = defaultWeightMsgApproveToken
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgApproveToken,
		antirugsimulation.SimulateMsgApproveToken(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSetModuleActive int
	simState.AppParams.GetOrGenerate(opWeightMsgSetModuleActive, &weightMsgSetModuleActive, nil,
		func(_ *rand.Rand) {
			weightMsgSetModuleActive = defaultWeightMsgSetModuleActive
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSetModuleActive,
		antirugsimulation.SimulateMsgSetModuleActive(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgRegisterToken,
			defaultWeightMsgRegisterToken,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				antirugsimulation.SimulateMsgRegisterToken(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgLockLiquidity,
			defaultWeightMsgLockLiquidity,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				antirugsimulation.SimulateMsgLockLiquidity(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeclareMintLimit,
			defaultWeightMsgDeclareMintLimit,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				antirugsimulation.SimulateMsgDeclareMintLimit(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgFlagToken,
			defaultWeightMsgFlagToken,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				antirugsimulation.SimulateMsgFlagToken(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgFreezeToken,
			defaultWeightMsgFreezeToken,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				antirugsimulation.SimulateMsgFreezeToken(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUnfreezeToken,
			defaultWeightMsgUnfreezeToken,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				antirugsimulation.SimulateMsgUnfreezeToken(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgRequestUpgrade,
			defaultWeightMsgRequestUpgrade,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				antirugsimulation.SimulateMsgRequestUpgrade(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgApproveToken,
			defaultWeightMsgApproveToken,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				antirugsimulation.SimulateMsgApproveToken(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgSetModuleActive,
			defaultWeightMsgSetModuleActive,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				antirugsimulation.SimulateMsgSetModuleActive(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
