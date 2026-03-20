package oanagent_test

import (
	"testing"

	keepertest "oan/testutil/keeper"
	"oan/testutil/nullify"
	oanagent "oan/x/oanagent/module"
	"oan/x/oanagent/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		AgentList: []types.Agent{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.OanagentKeeper(t)
	oanagent.InitGenesis(ctx, k, genesisState)
	got := oanagent.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.AgentList, got.AgentList)
	// this line is used by starport scaffolding # genesis/test/assert
}
