package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "qcb/testutil/keeper"
	"qcb/x/qcbqsec/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.OanqsecKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
