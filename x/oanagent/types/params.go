package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyMaxAgents = []byte("MaxAgents")
	// TODO: Determine the default value
	DefaultMaxAgents uint64 = 0
)

var (
	KeyCircuitBreakerMaxTx = []byte("CircuitBreakerMaxTx")
	// TODO: Determine the default value
	DefaultCircuitBreakerMaxTx uint64 = 0
)

var (
	KeyCircuitBreakerWindow = []byte("CircuitBreakerWindow")
	// TODO: Determine the default value
	DefaultCircuitBreakerWindow uint64 = 0
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	maxAgents uint64,
	circuitBreakerMaxTx uint64,
	circuitBreakerWindow uint64,
) Params {
	return Params{
		MaxAgents:            maxAgents,
		CircuitBreakerMaxTx:  circuitBreakerMaxTx,
		CircuitBreakerWindow: circuitBreakerWindow,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultMaxAgents,
		DefaultCircuitBreakerMaxTx,
		DefaultCircuitBreakerWindow,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMaxAgents, &p.MaxAgents, validateMaxAgents),
		paramtypes.NewParamSetPair(KeyCircuitBreakerMaxTx, &p.CircuitBreakerMaxTx, validateCircuitBreakerMaxTx),
		paramtypes.NewParamSetPair(KeyCircuitBreakerWindow, &p.CircuitBreakerWindow, validateCircuitBreakerWindow),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateMaxAgents(p.MaxAgents); err != nil {
		return err
	}

	if err := validateCircuitBreakerMaxTx(p.CircuitBreakerMaxTx); err != nil {
		return err
	}

	if err := validateCircuitBreakerWindow(p.CircuitBreakerWindow); err != nil {
		return err
	}

	return nil
}

// validateMaxAgents validates the MaxAgents param
func validateMaxAgents(v interface{}) error {
	maxAgents, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = maxAgents

	return nil
}

// validateCircuitBreakerMaxTx validates the CircuitBreakerMaxTx param
func validateCircuitBreakerMaxTx(v interface{}) error {
	circuitBreakerMaxTx, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = circuitBreakerMaxTx

	return nil
}

// validateCircuitBreakerWindow validates the CircuitBreakerWindow param
func validateCircuitBreakerWindow(v interface{}) error {
	circuitBreakerWindow, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = circuitBreakerWindow

	return nil
}
