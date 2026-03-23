package qcbeconomy_test

import (
	"testing"

	keepertest "qcb/testutil/keeper"
	"qcb/testutil/nullify"
	qcbeconomy "qcb/x/qcbeconomy/module"
	"qcb/x/qcbeconomy/types"

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
	qcbeconomy.InitGenesis(ctx, k, genesisState)
	got := qcbeconomy.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.TaskList, got.TaskList)
	// this line is used by starport scaffolding # genesis/test/assert
}
