package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyMaxIdentities = []byte("MaxIdentities")
	// TODO: Determine the default value
	DefaultMaxIdentities uint64 = 0
)

var (
	KeyRegistrationFee = []byte("RegistrationFee")
	// TODO: Determine the default value
	DefaultRegistrationFee uint64 = 0
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	maxIdentities uint64,
	registrationFee uint64,
) Params {
	return Params{
		MaxIdentities:   maxIdentities,
		RegistrationFee: registrationFee,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultMaxIdentities,
		DefaultRegistrationFee,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMaxIdentities, &p.MaxIdentities, validateMaxIdentities),
		paramtypes.NewParamSetPair(KeyRegistrationFee, &p.RegistrationFee, validateRegistrationFee),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateMaxIdentities(p.MaxIdentities); err != nil {
		return err
	}

	if err := validateRegistrationFee(p.RegistrationFee); err != nil {
		return err
	}

	return nil
}

// validateMaxIdentities validates the MaxIdentities param
func validateMaxIdentities(v interface{}) error {
	maxIdentities, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = maxIdentities

	return nil
}

// validateRegistrationFee validates the RegistrationFee param
func validateRegistrationFee(v interface{}) error {
	registrationFee, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = registrationFee

	return nil
}
