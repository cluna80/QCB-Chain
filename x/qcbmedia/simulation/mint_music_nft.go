package simulation

import (
	"math/rand"

	"qcb/x/qcbmedia/keeper"
	"qcb/x/qcbmedia/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgMintMusicNft(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgMintMusicNft{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the MintMusicNft simulation

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "MintMusicNft simulation not implemented"), nil, nil
	}
}
