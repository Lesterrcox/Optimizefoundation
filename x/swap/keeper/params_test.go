package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "optimizeglobalcoin/testutil/keeper"
	"optimizeglobalcoin/x/swap/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.SwapKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
