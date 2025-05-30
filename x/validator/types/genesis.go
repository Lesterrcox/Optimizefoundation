package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ValidatorList: []Validator{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in validator
	validatorIdMap := make(map[uint64]bool)
	validatorCount := gs.GetValidatorCount()
	for _, elem := range gs.ValidatorList {
		if _, ok := validatorIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for validator")
		}
		if elem.Id >= validatorCount {
			return fmt.Errorf("validator id should be lower or equal than the last id")
		}
		validatorIdMap[elem.Id] = true
	}

	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
