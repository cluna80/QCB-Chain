package oaneconomy

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"oan/x/oaneconomy/keeper"
	"oan/x/oaneconomy/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the task
	for _, elem := range genState.TaskList {
		k.SetTask(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.TaskList = k.GetAllTask(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
