package oaneconomy_test

import (
	"testing"

	keepertest "oan/testutil/keeper"
	"oan/testutil/nullify"
	oaneconomy "oan/x/oaneconomy/module"
	"oan/x/oaneconomy/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		TaskList: []types.Task{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.OaneconomyKeeper(t)
	oaneconomy.InitGenesis(ctx, k, genesisState)
	got := oaneconomy.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.TaskList, got.TaskList)
	// this line is used by starport scaffolding # genesis/test/assert
}
