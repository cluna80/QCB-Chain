package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyInferenceReward = []byte("InferenceReward")
	// TODO: Determine the default value
	DefaultInferenceReward uint64 = 0
)

var (
	KeySlashAmount = []byte("SlashAmount")
	// TODO: Determine the default value
	DefaultSlashAmount uint64 = 0
)

var (
	KeyComputeEpoch = []byte("ComputeEpoch")
	// TODO: Determine the default value
	DefaultComputeEpoch uint64 = 0
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	inferenceReward uint64,
	slashAmount uint64,
	computeEpoch uint64,
) Params {
	return Params{
		InferenceReward: inferenceReward,
		SlashAmount:     slashAmount,
		ComputeEpoch:    computeEpoch,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultInferenceReward,
		DefaultSlashAmount,
		DefaultComputeEpoch,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyInferenceReward, &p.InferenceReward, validateInferenceReward),
		paramtypes.NewParamSetPair(KeySlashAmount, &p.SlashAmount, validateSlashAmount),
		paramtypes.NewParamSetPair(KeyComputeEpoch, &p.ComputeEpoch, validateComputeEpoch),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateInferenceReward(p.InferenceReward); err != nil {
		return err
	}

	if err := validateSlashAmount(p.SlashAmount); err != nil {
		return err
	}

	if err := validateComputeEpoch(p.ComputeEpoch); err != nil {
		return err
	}

	return nil
}

// validateInferenceReward validates the InferenceReward param
func validateInferenceReward(v interface{}) error {
	inferenceReward, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = inferenceReward

	return nil
}

// validateSlashAmount validates the SlashAmount param
func validateSlashAmount(v interface{}) error {
	slashAmount, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = slashAmount

	return nil
}

// validateComputeEpoch validates the ComputeEpoch param
func validateComputeEpoch(v interface{}) error {
	computeEpoch, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = computeEpoch

	return nil
}
