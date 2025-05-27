package app

import (
	"cosmossdk.io/core/appmodule"
	storetypes "cosmossdk.io/store/types"

	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	oconsensusmodulekeeper "optimizeglobalcoin/x/oconsensus/keeper"
	oconsensusmodule "optimizeglobalcoin/x/oconsensus/module"
	oconsensusmoduletypes "optimizeglobalcoin/x/oconsensus/types"

	validatormodulekeeper "optimizeglobalcoin/x/validator/keeper"
	validatormodule "optimizeglobalcoin/x/validator/module"
	validatormoduletypes "optimizeglobalcoin/x/validator/types"

	assetmodulekeeper "optimizeglobalcoin/x/asset/keeper"
	assetmodule "optimizeglobalcoin/x/asset/module"
	assetmoduletypes "optimizeglobalcoin/x/asset/types"
)

func (app *App) registerCustomModules(_ servertypes.AppOptions) error {
	assetKey := storetypes.NewKVStoreKey(assetmoduletypes.StoreKey)
	oconsensusKey := storetypes.NewKVStoreKey(oconsensusmoduletypes.StoreKey)
	validatorKey := storetypes.NewKVStoreKey(validatormoduletypes.StoreKey)

	if err := app.RegisterStores(
		assetKey,
		oconsensusKey,
		validatorKey,
	); err != nil {
		return err
	}

	govAddress := authtypes.NewModuleAddress(govtypes.ModuleName).String()
	keyTable := oconsensusmoduletypes.ParamKeyTable()
	keyTable.RegisterParamSet(&oconsensusmoduletypes.Params{})
	app.ParamsKeeper.Subspace(oconsensusmoduletypes.ModuleName).WithKeyTable(keyTable)

	app.OconsensusKeeper = oconsensusmodulekeeper.NewKeeper(
		app.appCodec,
		runtime.NewKVStoreService(oconsensusKey),
		app.Logger(),
		app.StakingKeeper,
		govAddress,
	)

	keyTable.RegisterParamSet(&validatormoduletypes.Params{})
	app.ParamsKeeper.Subspace(validatormoduletypes.ModuleName).WithKeyTable(keyTable)
	app.ValidatorKeeper = validatormodulekeeper.NewKeeper(
		app.appCodec,
		runtime.NewKVStoreService(assetKey),
		app.Logger(),
		govAddress,
	)

	keyTable.RegisterParamSet(&assetmoduletypes.Params{})
	app.ParamsKeeper.Subspace(assetmoduletypes.ModuleName).WithKeyTable(keyTable)
	app.AssetKeeper = assetmodulekeeper.NewKeeper(
		app.appCodec,
		runtime.NewKVStoreService(assetKey),
		app.Logger(),
		app.OconsensusKeeper,
		app.ValidatorKeeper,
		govAddress,
	)

	if err := app.RegisterModules(
		assetmodule.NewAppModule(app.appCodec, app.AssetKeeper, app.AccountKeeper, app.BankKeeper),
		oconsensusmodule.NewAppModule(app.appCodec, app.OconsensusKeeper, app.AccountKeeper, app.BankKeeper),
		validatormodule.NewAppModule(app.appCodec, app.ValidatorKeeper, app.AccountKeeper, app.BankKeeper),
	); err != nil {
		return err
	}

	return nil
}

func RegisterOgcCustomModules(registry cdctypes.InterfaceRegistry) map[string]appmodule.AppModule {
	modules := map[string]appmodule.AppModule{
		assetmoduletypes.ModuleName:      assetmodule.AppModule{},
		oconsensusmoduletypes.ModuleName: oconsensusmodule.AppModule{},
		validatormoduletypes.ModuleName:  validatormodule.AppModule{},
	}

	for name, m := range modules {
		module.CoreAppModuleBasicAdaptor(name, m).RegisterInterfaces(registry)
	}

	return modules
}
