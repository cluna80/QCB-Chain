package oandao_test

import (
	"testing"

	keepertest "oan/testutil/keeper"
	"oan/testutil/nullify"
	oandao "oan/x/oandao/module"
	"oan/x/oandao/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		ProposalList: []types.Proposal{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.OandaoKeeper(t)
	oandao.InitGenesis(ctx, k, genesisState)
	got := oandao.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.ProposalList, got.ProposalList)
	// this line is used by starport scaffolding # genesis/test/assert
}
