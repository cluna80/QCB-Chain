package simulation

import (
	"math/rand"

	"qcb/x/qcbqsec/keeper"
	"qcb/x/qcbqsec/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgDeprecateAlgorithm(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgDeprecateAlgorithm{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the DeprecateAlgorithm simulation

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "DeprecateAlgorithm simulation not implemented"), nil, nil
	}
}
