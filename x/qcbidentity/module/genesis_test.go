package qcbidentity_test

import (
	"testing"

	keepertest "qcb/testutil/keeper"
	"qcb/testutil/nullify"
	qcbidentity "qcb/x/qcbidentity/module"
	"qcb/x/qcbidentity/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		IdentityList: []types.Identity{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.OanidentityKeeper(t)
	qcbidentity.InitGenesis(ctx, k, genesisState)
	got := qcbidentity.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.IdentityList, got.IdentityList)
	// this line is used by starport scaffolding # genesis/test/assert
}
