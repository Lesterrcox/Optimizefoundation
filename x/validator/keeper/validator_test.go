package keeper_test

import (
	"context"
	"testing"

	keepertest "optimizeglobalcoin/testutil/keeper"
	"optimizeglobalcoin/testutil/nullify"
	"optimizeglobalcoin/x/validator/keeper"
	"optimizeglobalcoin/x/validator/types"

	"github.com/stretchr/testify/require"
)

func createNValidator(keeper keeper.Keeper, ctx context.Context, n int) []types.Validator {
	items := make([]types.Validator, n)
	for i := range items {
		items[i].Id = keeper.AppendValidator(ctx, items[i])
	}
	return items
}

func TestValidatorGet(t *testing.T) {
	keeper, ctx := keepertest.OptimizeglobalcoinKeeper(t)
	items := createNValidator(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetValidator(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestValidatorRemove(t *testing.T) {
	keeper, ctx := keepertest.OptimizeglobalcoinKeeper(t)
	items := createNValidator(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveValidator(ctx, item.Id)
		_, found := keeper.GetValidator(ctx, item.Id)
		require.False(t, found)
	}
}

func TestValidatorGetAll(t *testing.T) {
	keeper, ctx := keepertest.OptimizeglobalcoinKeeper(t)
	items := createNValidator(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllValidator(ctx)),
	)
}

func TestValidatorCount(t *testing.T) {
	keeper, ctx := keepertest.OptimizeglobalcoinKeeper(t)
	items := createNValidator(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetValidatorCount(ctx))
}
