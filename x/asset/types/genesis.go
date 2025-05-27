package types

import (
	fmt "fmt"
)

// this line is used by starport scaffolding # genesis/types/import

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ValueVotesList: []ValueVotes{},
		// this line is used by starport scaffolding # genesis/types/default
		Params:    DefaultParams(),
		AssetList: []Asset{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in asset
	assetIdMap := make(map[uint64]bool)
	assetCount := gs.GetAssetCount()
	for _, elem := range gs.AssetList {
		if _, ok := assetIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for asset")
		}
		if elem.Id >= assetCount {
			return fmt.Errorf("asset id should be lower or equal than the last id")
		}
		assetIdMap[elem.Id] = true
	}
	// Check for duplicated ID in valueVotes
	valueVotesIdMap := make(map[uint64]bool)
	valueVotesCount := gs.GetValueVotesCount()
	for _, elem := range gs.ValueVotesList {
		if _, ok := valueVotesIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for valueVotes")
		}
		if elem.Id >= valueVotesCount {
			return fmt.Errorf("valueVotes id should be lower or equal than the last id")
		}
		valueVotesIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
