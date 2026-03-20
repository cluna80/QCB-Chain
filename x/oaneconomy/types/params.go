package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyTaskFee = []byte("TaskFee")
	// TODO: Determine the default value
	DefaultTaskFee uint64 = 0
)

var (
	KeyUbiRate = []byte("UbiRate")
	// TODO: Determine the default value
	DefaultUbiRate uint64 = 0
)

var (
	KeyMaxTasks = []byte("MaxTasks")
	// TODO: Determine the default value
	DefaultMaxTasks uint64 = 0
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	taskFee uint64,
	ubiRate uint64,
	maxTasks uint64,
) Params {
	return Params{
		TaskFee:  taskFee,
		UbiRate:  ubiRate,
		MaxTasks: maxTasks,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultTaskFee,
		DefaultUbiRate,
		DefaultMaxTasks,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyTaskFee, &p.TaskFee, validateTaskFee),
		paramtypes.NewParamSetPair(KeyUbiRate, &p.UbiRate, validateUbiRate),
		paramtypes.NewParamSetPair(KeyMaxTasks, &p.MaxTasks, validateMaxTasks),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateTaskFee(p.TaskFee); err != nil {
		return err
	}

	if err := validateUbiRate(p.UbiRate); err != nil {
		return err
	}

	if err := validateMaxTasks(p.MaxTasks); err != nil {
		return err
	}

	return nil
}

// validateTaskFee validates the TaskFee param
func validateTaskFee(v interface{}) error {
	taskFee, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = taskFee

	return nil
}

// validateUbiRate validates the UbiRate param
func validateUbiRate(v interface{}) error {
	ubiRate, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = ubiRate

	return nil
}

// validateMaxTasks validates the MaxTasks param
func validateMaxTasks(v interface{}) error {
	maxTasks, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = maxTasks

	return nil
}
