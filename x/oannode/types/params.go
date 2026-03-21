package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyMinValidatorStake = []byte("MinValidatorStake")
	// TODO: Determine the default value
	DefaultMinValidatorStake uint64 = 0
)

var (
	KeyMinInferenceStake = []byte("MinInferenceStake")
	// TODO: Determine the default value
	DefaultMinInferenceStake uint64 = 0
)

var (
	KeyMinBridgeStake = []byte("MinBridgeStake")
	// TODO: Determine the default value
	DefaultMinBridgeStake uint64 = 0
)

var (
	KeyMinLightStake = []byte("MinLightStake")
	// TODO: Determine the default value
	DefaultMinLightStake uint64 = 0
)

var (
	KeyUnbondingBlocks = []byte("UnbondingBlocks")
	// TODO: Determine the default value
	DefaultUnbondingBlocks uint64 = 0
)

var (
	KeyMaxWalletStakePct = []byte("MaxWalletStakePct")
	// TODO: Determine the default value
	DefaultMaxWalletStakePct uint64 = 0
)

var (
	KeyEpochLength = []byte("EpochLength")
	// TODO: Determine the default value
	DefaultEpochLength uint64 = 0
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	minValidatorStake uint64,
	minInferenceStake uint64,
	minBridgeStake uint64,
	minLightStake uint64,
	unbondingBlocks uint64,
	maxWalletStakePct uint64,
	epochLength uint64,
) Params {
	return Params{
		MinValidatorStake: minValidatorStake,
		MinInferenceStake: minInferenceStake,
		MinBridgeStake:    minBridgeStake,
		MinLightStake:     minLightStake,
		UnbondingBlocks:   unbondingBlocks,
		MaxWalletStakePct: maxWalletStakePct,
		EpochLength:       epochLength,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultMinValidatorStake,
		DefaultMinInferenceStake,
		DefaultMinBridgeStake,
		DefaultMinLightStake,
		DefaultUnbondingBlocks,
		DefaultMaxWalletStakePct,
		DefaultEpochLength,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMinValidatorStake, &p.MinValidatorStake, validateMinValidatorStake),
		paramtypes.NewParamSetPair(KeyMinInferenceStake, &p.MinInferenceStake, validateMinInferenceStake),
		paramtypes.NewParamSetPair(KeyMinBridgeStake, &p.MinBridgeStake, validateMinBridgeStake),
		paramtypes.NewParamSetPair(KeyMinLightStake, &p.MinLightStake, validateMinLightStake),
		paramtypes.NewParamSetPair(KeyUnbondingBlocks, &p.UnbondingBlocks, validateUnbondingBlocks),
		paramtypes.NewParamSetPair(KeyMaxWalletStakePct, &p.MaxWalletStakePct, validateMaxWalletStakePct),
		paramtypes.NewParamSetPair(KeyEpochLength, &p.EpochLength, validateEpochLength),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateMinValidatorStake(p.MinValidatorStake); err != nil {
		return err
	}

	if err := validateMinInferenceStake(p.MinInferenceStake); err != nil {
		return err
	}

	if err := validateMinBridgeStake(p.MinBridgeStake); err != nil {
		return err
	}

	if err := validateMinLightStake(p.MinLightStake); err != nil {
		return err
	}

	if err := validateUnbondingBlocks(p.UnbondingBlocks); err != nil {
		return err
	}

	if err := validateMaxWalletStakePct(p.MaxWalletStakePct); err != nil {
		return err
	}

	if err := validateEpochLength(p.EpochLength); err != nil {
		return err
	}

	return nil
}

// validateMinValidatorStake validates the MinValidatorStake param
func validateMinValidatorStake(v interface{}) error {
	minValidatorStake, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = minValidatorStake

	return nil
}

// validateMinInferenceStake validates the MinInferenceStake param
func validateMinInferenceStake(v interface{}) error {
	minInferenceStake, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = minInferenceStake

	return nil
}

// validateMinBridgeStake validates the MinBridgeStake param
func validateMinBridgeStake(v interface{}) error {
	minBridgeStake, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = minBridgeStake

	return nil
}

// validateMinLightStake validates the MinLightStake param
func validateMinLightStake(v interface{}) error {
	minLightStake, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = minLightStake

	return nil
}

// validateUnbondingBlocks validates the UnbondingBlocks param
func validateUnbondingBlocks(v interface{}) error {
	unbondingBlocks, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = unbondingBlocks

	return nil
}

// validateMaxWalletStakePct validates the MaxWalletStakePct param
func validateMaxWalletStakePct(v interface{}) error {
	maxWalletStakePct, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = maxWalletStakePct

	return nil
}

// validateEpochLength validates the EpochLength param
func validateEpochLength(v interface{}) error {
	epochLength, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = epochLength

	return nil
}
