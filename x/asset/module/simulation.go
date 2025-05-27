package asset

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"optimizeglobalcoin/testutil/sample"
	assetsimulation "optimizeglobalcoin/x/asset/simulation"
	"optimizeglobalcoin/x/asset/types"
)

// avoid unused import issue
var (
	_ = assetsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgSubmitAsset = "op_weight_msg_submit_asset"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSubmitAsset int = 100

	opWeightMsgSubmitValue = "op_weight_msg_submit_value"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSubmitValue int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	assetGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&assetGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgSubmitAsset int
	simState.AppParams.GetOrGenerate(opWeightMsgSubmitAsset, &weightMsgSubmitAsset, nil,
		func(_ *rand.Rand) {
			weightMsgSubmitAsset = defaultWeightMsgSubmitAsset
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSubmitAsset,
		assetsimulation.SimulateMsgSubmitAsset(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSubmitValue int
	simState.AppParams.GetOrGenerate(opWeightMsgSubmitValue, &weightMsgSubmitValue, nil,
		func(_ *rand.Rand) {
			weightMsgSubmitValue = defaultWeightMsgSubmitValue
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSubmitValue,
		assetsimulation.SimulateMsgSubmitValue(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgSubmitAsset,
			defaultWeightMsgSubmitAsset,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				assetsimulation.SimulateMsgSubmitAsset(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgSubmitValue,
			defaultWeightMsgSubmitValue,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				assetsimulation.SimulateMsgSubmitValue(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
