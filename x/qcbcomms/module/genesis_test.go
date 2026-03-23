package qcbcomms_test

import (
	"testing"

	keepertest "qcb/testutil/keeper"
	"qcb/testutil/nullify"
	qcbcomms "qcb/x/qcbcomms/module"
	"qcb/x/qcbcomms/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.OancommsKeeper(t)
	qcbcomms.InitGenesis(ctx, k, genesisState)
	got := qcbcomms.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
