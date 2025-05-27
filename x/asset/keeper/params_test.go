package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "optimizeglobalcoin/testutil/keeper"
	"optimizeglobalcoin/x/asset/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.AssetKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
