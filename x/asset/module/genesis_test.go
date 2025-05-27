package asset_test

import (
	"testing"

	keepertest "optimizeglobalcoin/testutil/keeper"
	"optimizeglobalcoin/testutil/nullify"
	asset "optimizeglobalcoin/x/asset/module"
	"optimizeglobalcoin/x/asset/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		ValueVotesList: []types.ValueVotes{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		ValueVotesCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AssetKeeper(t)
	asset.InitGenesis(ctx, k, genesisState)
	got := asset.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.ValueVotesList, got.ValueVotesList)
	require.Equal(t, genesisState.ValueVotesCount, got.ValueVotesCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
