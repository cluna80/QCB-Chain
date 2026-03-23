package qcbwalletproto_test

import (
	"testing"

	keepertest "qcb/testutil/keeper"
	"qcb/testutil/nullify"
	qcbwalletproto "qcb/x/qcbwalletproto/module"
	"qcb/x/qcbwalletproto/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.OanwalletprotoKeeper(t)
	qcbwalletproto.InitGenesis(ctx, k, genesisState)
	got := qcbwalletproto.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
