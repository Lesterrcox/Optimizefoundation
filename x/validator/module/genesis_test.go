package optimizeglobalcoin_test

import (
	"testing"

	keepertest "optimizeglobalcoin/testutil/keeper"
	"optimizeglobalcoin/testutil/nullify"
	optimizeglobalcoin "optimizeglobalcoin/x/validator/module"
	"optimizeglobalcoin/x/validator/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		ValidatorList: []types.Validator{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		ValidatorCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.OptimizeglobalcoinKeeper(t)
	optimizeglobalcoin.InitGenesis(ctx, k, genesisState)
	got := optimizeglobalcoin.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.ValidatorList, got.ValidatorList)
	require.Equal(t, genesisState.ValidatorCount, got.ValidatorCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
