package keeper_test

import (
	"context"
	"testing"

	keepertest "optimizeglobalcoin/testutil/keeper"
	"optimizeglobalcoin/testutil/nullify"
	"optimizeglobalcoin/x/asset/keeper"
	"optimizeglobalcoin/x/asset/types"

	"github.com/stretchr/testify/require"
)

func createNValueVotes(keeper keeper.Keeper, ctx context.Context, n int) []types.ValueVotes {
	items := make([]types.ValueVotes, n)
	for i := range items {
		items[i].Id = keeper.AppendValueVotes(ctx, items[i])
	}
	return items
}

func TestValueVotesGet(t *testing.T) {
	keeper, ctx := keepertest.AssetKeeper(t)
	items := createNValueVotes(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetValueVotes(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestValueVotesRemove(t *testing.T) {
	keeper, ctx := keepertest.AssetKeeper(t)
	items := createNValueVotes(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveValueVotes(ctx, item.Id)
		_, found := keeper.GetValueVotes(ctx, item.Id)
		require.False(t, found)
	}
}

func TestValueVotesGetAll(t *testing.T) {
	keeper, ctx := keepertest.AssetKeeper(t)
	items := createNValueVotes(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllValueVotes(ctx)),
	)
}

func TestValueVotesCount(t *testing.T) {
	keeper, ctx := keepertest.AssetKeeper(t)
	items := createNValueVotes(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetValueVotesCount(ctx))
}
