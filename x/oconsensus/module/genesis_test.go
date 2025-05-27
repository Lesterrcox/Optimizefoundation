package oconsensus_test

import (
	"testing"

	keepertest "optimizeglobalcoin/testutil/keeper"
	"optimizeglobalcoin/testutil/nullify"
	oconsensus "optimizeglobalcoin/x/oconsensus/module"
	"optimizeglobalcoin/x/oconsensus/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.OconsensusKeeper(t)
	oconsensus.InitGenesis(ctx, k, genesisState)
	got := oconsensus.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
