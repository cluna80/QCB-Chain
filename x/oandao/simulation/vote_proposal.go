package simulation

import (
	"math/rand"

	"oan/x/oandao/keeper"
	"oan/x/oandao/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgVoteProposal(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgVoteProposal{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the VoteProposal simulation

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "VoteProposal simulation not implemented"), nil, nil
	}
}
