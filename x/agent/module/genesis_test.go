package agent_test

import (
	"testing"

	keepertest "qcb/testutil/keeper"
	"qcb/testutil/nullify"
	agent "qcb/x/agent/module"
	"qcb/x/agent/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AgentKeeper(t)
	agent.InitGenesis(ctx, k, genesisState)
	got := agent.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
