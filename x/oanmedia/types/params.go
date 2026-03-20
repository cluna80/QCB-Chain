package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyRoyaltyRate = []byte("RoyaltyRate")
	// TODO: Determine the default value
	DefaultRoyaltyRate uint64 = 0
)

var (
	KeyMintFee = []byte("MintFee")
	// TODO: Determine the default value
	DefaultMintFee uint64 = 0
)

var (
	KeyMaxRoyaltyChain = []byte("MaxRoyaltyChain")
	// TODO: Determine the default value
	DefaultMaxRoyaltyChain uint64 = 0
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	royaltyRate uint64,
	mintFee uint64,
	maxRoyaltyChain uint64,
) Params {
	return Params{
		RoyaltyRate:     royaltyRate,
		MintFee:         mintFee,
		MaxRoyaltyChain: maxRoyaltyChain,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultRoyaltyRate,
		DefaultMintFee,
		DefaultMaxRoyaltyChain,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyRoyaltyRate, &p.RoyaltyRate, validateRoyaltyRate),
		paramtypes.NewParamSetPair(KeyMintFee, &p.MintFee, validateMintFee),
		paramtypes.NewParamSetPair(KeyMaxRoyaltyChain, &p.MaxRoyaltyChain, validateMaxRoyaltyChain),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateRoyaltyRate(p.RoyaltyRate); err != nil {
		return err
	}

	if err := validateMintFee(p.MintFee); err != nil {
		return err
	}

	if err := validateMaxRoyaltyChain(p.MaxRoyaltyChain); err != nil {
		return err
	}

	return nil
}

// validateRoyaltyRate validates the RoyaltyRate param
func validateRoyaltyRate(v interface{}) error {
	royaltyRate, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = royaltyRate

	return nil
}

// validateMintFee validates the MintFee param
func validateMintFee(v interface{}) error {
	mintFee, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = mintFee

	return nil
}

// validateMaxRoyaltyChain validates the MaxRoyaltyChain param
func validateMaxRoyaltyChain(v interface{}) error {
	maxRoyaltyChain, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = maxRoyaltyChain

	return nil
}
