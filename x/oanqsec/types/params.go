package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyThreatLevel = []byte("ThreatLevel")
	// TODO: Determine the default value
	DefaultThreatLevel uint64 = 0
)

var (
	KeyHybridRequired = []byte("HybridRequired")
	// TODO: Determine the default value
	DefaultHybridRequired bool = false
)

var (
	KeyRotationPeriod = []byte("RotationPeriod")
	// TODO: Determine the default value
	DefaultRotationPeriod uint64 = 0
)

var (
	KeyActiveAlgorithm = []byte("ActiveAlgorithm")
	// TODO: Determine the default value
	DefaultActiveAlgorithm string = "active_algorithm"
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	threatLevel uint64,
	hybridRequired bool,
	rotationPeriod uint64,
	activeAlgorithm string,
) Params {
	return Params{
		ThreatLevel:     threatLevel,
		HybridRequired:  hybridRequired,
		RotationPeriod:  rotationPeriod,
		ActiveAlgorithm: activeAlgorithm,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultThreatLevel,
		DefaultHybridRequired,
		DefaultRotationPeriod,
		DefaultActiveAlgorithm,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyThreatLevel, &p.ThreatLevel, validateThreatLevel),
		paramtypes.NewParamSetPair(KeyHybridRequired, &p.HybridRequired, validateHybridRequired),
		paramtypes.NewParamSetPair(KeyRotationPeriod, &p.RotationPeriod, validateRotationPeriod),
		paramtypes.NewParamSetPair(KeyActiveAlgorithm, &p.ActiveAlgorithm, validateActiveAlgorithm),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateThreatLevel(p.ThreatLevel); err != nil {
		return err
	}

	if err := validateHybridRequired(p.HybridRequired); err != nil {
		return err
	}

	if err := validateRotationPeriod(p.RotationPeriod); err != nil {
		return err
	}

	if err := validateActiveAlgorithm(p.ActiveAlgorithm); err != nil {
		return err
	}

	return nil
}

// validateThreatLevel validates the ThreatLevel param
func validateThreatLevel(v interface{}) error {
	threatLevel, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = threatLevel

	return nil
}

// validateHybridRequired validates the HybridRequired param
func validateHybridRequired(v interface{}) error {
	hybridRequired, ok := v.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = hybridRequired

	return nil
}

// validateRotationPeriod validates the RotationPeriod param
func validateRotationPeriod(v interface{}) error {
	rotationPeriod, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = rotationPeriod

	return nil
}

// validateActiveAlgorithm validates the ActiveAlgorithm param
func validateActiveAlgorithm(v interface{}) error {
	activeAlgorithm, ok := v.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = activeAlgorithm

	return nil
}
