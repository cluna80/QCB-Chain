package simulation

import (
	"math/rand"

	"oan/x/oanmarket/keeper"
	"oan/x/oanmarket/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgHireAgent(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgHireAgent{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the HireAgent simulation

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "HireAgent simulation not implemented"), nil, nil
	}
}
