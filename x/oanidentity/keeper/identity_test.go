package keeper_test

import (
	"context"
	"strconv"
	"testing"

	keepertest "oan/testutil/keeper"
	"oan/testutil/nullify"
	"oan/x/oanidentity/keeper"
	"oan/x/oanidentity/types"

	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNIdentity(keeper keeper.Keeper, ctx context.Context, n int) []types.Identity {
	items := make([]types.Identity, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetIdentity(ctx, items[i])
	}
	return items
}

func TestIdentityGet(t *testing.T) {
	keeper, ctx := keepertest.OanidentityKeeper(t)
	items := createNIdentity(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetIdentity(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestIdentityRemove(t *testing.T) {
	keeper, ctx := keepertest.OanidentityKeeper(t)
	items := createNIdentity(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveIdentity(ctx,
			item.Index,
		)
		_, found := keeper.GetIdentity(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestIdentityGetAll(t *testing.T) {
	keeper, ctx := keepertest.OanidentityKeeper(t)
	items := createNIdentity(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllIdentity(ctx)),
	)
}
