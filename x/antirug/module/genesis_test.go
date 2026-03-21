package antirug_test

import (
	"testing"

	keepertest "oan/testutil/keeper"
	"oan/testutil/nullify"
	antirug "oan/x/antirug/module"
	"oan/x/antirug/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AntirugKeeper(t)
	antirug.InitGenesis(ctx, k, genesisState)
	got := antirug.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
