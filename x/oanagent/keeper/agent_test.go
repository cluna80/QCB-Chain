package keeper_test

import (
	"context"
	"strconv"
	"testing"

	keepertest "oan/testutil/keeper"
	"oan/testutil/nullify"
	"oan/x/oanagent/keeper"
	"oan/x/oanagent/types"

	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNAgent(keeper keeper.Keeper, ctx context.Context, n int) []types.Agent {
	items := make([]types.Agent, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetAgent(ctx, items[i])
	}
	return items
}

func TestAgentGet(t *testing.T) {
	keeper, ctx := keepertest.OanagentKeeper(t)
	items := createNAgent(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetAgent(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestAgentRemove(t *testing.T) {
	keeper, ctx := keepertest.OanagentKeeper(t)
	items := createNAgent(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveAgent(ctx,
			item.Index,
		)
		_, found := keeper.GetAgent(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestAgentGetAll(t *testing.T) {
	keeper, ctx := keepertest.OanagentKeeper(t)
	items := createNAgent(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllAgent(ctx)),
	)
}
