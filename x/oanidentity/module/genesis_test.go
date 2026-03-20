package oanidentity_test

import (
	"testing"

	keepertest "oan/testutil/keeper"
	"oan/testutil/nullify"
	oanidentity "oan/x/oanidentity/module"
	"oan/x/oanidentity/types"

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
	oanidentity.InitGenesis(ctx, k, genesisState)
	got := oanidentity.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.IdentityList, got.IdentityList)
	// this line is used by starport scaffolding # genesis/test/assert
}
