package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyMinGuardians = []byte("MinGuardians")
	// TODO: Determine the default value
	DefaultMinGuardians uint64 = 0
)

var (
	KeyMaxGuardians = []byte("MaxGuardians")
	// TODO: Determine the default value
	DefaultMaxGuardians uint64 = 0
)

var (
	KeyVetoThreshold = []byte("VetoThreshold")
	// TODO: Determine the default value
	DefaultVetoThreshold uint64 = 0
)

var (
	KeyPauseDuration = []byte("PauseDuration")
	// TODO: Determine the default value
	DefaultPauseDuration uint64 = 0
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	minGuardians uint64,
	maxGuardians uint64,
	vetoThreshold uint64,
	pauseDuration uint64,
) Params {
	return Params{
		MinGuardians:  minGuardians,
		MaxGuardians:  maxGuardians,
		VetoThreshold: vetoThreshold,
		PauseDuration: pauseDuration,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultMinGuardians,
		DefaultMaxGuardians,
		DefaultVetoThreshold,
		DefaultPauseDuration,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMinGuardians, &p.MinGuardians, validateMinGuardians),
		paramtypes.NewParamSetPair(KeyMaxGuardians, &p.MaxGuardians, validateMaxGuardians),
		paramtypes.NewParamSetPair(KeyVetoThreshold, &p.VetoThreshold, validateVetoThreshold),
		paramtypes.NewParamSetPair(KeyPauseDuration, &p.PauseDuration, validatePauseDuration),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateMinGuardians(p.MinGuardians); err != nil {
		return err
	}

	if err := validateMaxGuardians(p.MaxGuardians); err != nil {
		return err
	}

	if err := validateVetoThreshold(p.VetoThreshold); err != nil {
		return err
	}

	if err := validatePauseDuration(p.PauseDuration); err != nil {
		return err
	}

	return nil
}

// validateMinGuardians validates the MinGuardians param
func validateMinGuardians(v interface{}) error {
	minGuardians, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = minGuardians

	return nil
}

// validateMaxGuardians validates the MaxGuardians param
func validateMaxGuardians(v interface{}) error {
	maxGuardians, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = maxGuardians

	return nil
}

// validateVetoThreshold validates the VetoThreshold param
func validateVetoThreshold(v interface{}) error {
	vetoThreshold, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = vetoThreshold

	return nil
}

// validatePauseDuration validates the PauseDuration param
func validatePauseDuration(v interface{}) error {
	pauseDuration, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = pauseDuration

	return nil
}
