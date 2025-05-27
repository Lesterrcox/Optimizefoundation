package optimizeglobalcoin

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"optimizeglobalcoin/x/validator/keeper"
	"optimizeglobalcoin/x/validator/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the validator
	for _, elem := range genState.ValidatorList {
		k.SetValidator(ctx, elem)
	}

	// Set validator count
	k.SetValidatorCount(ctx, genState.ValidatorCount)

	// this line is used by starport scaffolding # genesis/module/init
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.ValidatorList = k.GetAllValidator(ctx)
	genesis.ValidatorCount = k.GetValidatorCount(ctx)

	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
