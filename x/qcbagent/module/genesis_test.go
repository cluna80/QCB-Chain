package qcbagent_test

import (
	"testing"

	keepertest "qcb/testutil/keeper"
	"qcb/testutil/nullify"
	qcbagent "qcb/x/qcbagent/module"
	"qcb/x/qcbagent/types"

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
	qcbagent.InitGenesis(ctx, k, genesisState)
	got := qcbagent.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.AgentList, got.AgentList)
	// this line is used by starport scaffolding # genesis/test/assert
}
