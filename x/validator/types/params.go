package types

import (
	fmt "fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeySeatsForDao = []byte("SeatsForDao")
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(seatsForDao uint64) Params {
	return Params{
		SeatsForDao: seatsForDao,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(10)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeySeatsForDao, &p.SeatsForDao, validateSeatsForDao),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateSeatsForDao(p.SeatsForDao); err != nil {
		return err
	}

	return nil
}

func validateSeatsForDao(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if v < 0 {
		return fmt.Errorf("seats for dao must be positive")
	}
	return nil
}
