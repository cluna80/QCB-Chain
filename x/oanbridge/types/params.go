package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyBridgeFee = []byte("BridgeFee")
	// TODO: Determine the default value
	DefaultBridgeFee uint64 = 0
)

var (
	KeyMaxPacketSize = []byte("MaxPacketSize")
	// TODO: Determine the default value
	DefaultMaxPacketSize uint64 = 0
)

var (
	KeyTimeoutBlocks = []byte("TimeoutBlocks")
	// TODO: Determine the default value
	DefaultTimeoutBlocks uint64 = 0
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	bridgeFee uint64,
	maxPacketSize uint64,
	timeoutBlocks uint64,
) Params {
	return Params{
		BridgeFee:     bridgeFee,
		MaxPacketSize: maxPacketSize,
		TimeoutBlocks: timeoutBlocks,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultBridgeFee,
		DefaultMaxPacketSize,
		DefaultTimeoutBlocks,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyBridgeFee, &p.BridgeFee, validateBridgeFee),
		paramtypes.NewParamSetPair(KeyMaxPacketSize, &p.MaxPacketSize, validateMaxPacketSize),
		paramtypes.NewParamSetPair(KeyTimeoutBlocks, &p.TimeoutBlocks, validateTimeoutBlocks),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateBridgeFee(p.BridgeFee); err != nil {
		return err
	}

	if err := validateMaxPacketSize(p.MaxPacketSize); err != nil {
		return err
	}

	if err := validateTimeoutBlocks(p.TimeoutBlocks); err != nil {
		return err
	}

	return nil
}

// validateBridgeFee validates the BridgeFee param
func validateBridgeFee(v interface{}) error {
	bridgeFee, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = bridgeFee

	return nil
}

// validateMaxPacketSize validates the MaxPacketSize param
func validateMaxPacketSize(v interface{}) error {
	maxPacketSize, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = maxPacketSize

	return nil
}

// validateTimeoutBlocks validates the TimeoutBlocks param
func validateTimeoutBlocks(v interface{}) error {
	timeoutBlocks, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = timeoutBlocks

	return nil
}
