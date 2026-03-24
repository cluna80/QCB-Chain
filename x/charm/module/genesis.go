package charm

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"qcb/x/charm/keeper"
	"qcb/x/charm/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {}
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState { return types.DefaultGenesis() }
