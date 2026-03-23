package qcbdao_test

import (
	"testing"

	keepertest "qcb/testutil/keeper"
	"qcb/testutil/nullify"
	qcbdao "qcb/x/qcbdao/module"
	"qcb/x/qcbdao/types"

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
	qcbdao.InitGenesis(ctx, k, genesisState)
	got := qcbdao.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.ProposalList, got.ProposalList)
	// this line is used by starport scaffolding # genesis/test/assert
}
