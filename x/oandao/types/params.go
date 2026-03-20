package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyVotingPeriod = []byte("VotingPeriod")
	// TODO: Determine the default value
	DefaultVotingPeriod uint64 = 0
)

var (
	KeyQuorum = []byte("Quorum")
	// TODO: Determine the default value
	DefaultQuorum uint64 = 0
)

var (
	KeyTimelockSeconds = []byte("TimelockSeconds")
	// TODO: Determine the default value
	DefaultTimelockSeconds uint64 = 0
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	votingPeriod uint64,
	quorum uint64,
	timelockSeconds uint64,
) Params {
	return Params{
		VotingPeriod:    votingPeriod,
		Quorum:          quorum,
		TimelockSeconds: timelockSeconds,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultVotingPeriod,
		DefaultQuorum,
		DefaultTimelockSeconds,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyVotingPeriod, &p.VotingPeriod, validateVotingPeriod),
		paramtypes.NewParamSetPair(KeyQuorum, &p.Quorum, validateQuorum),
		paramtypes.NewParamSetPair(KeyTimelockSeconds, &p.TimelockSeconds, validateTimelockSeconds),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateVotingPeriod(p.VotingPeriod); err != nil {
		return err
	}

	if err := validateQuorum(p.Quorum); err != nil {
		return err
	}

	if err := validateTimelockSeconds(p.TimelockSeconds); err != nil {
		return err
	}

	return nil
}

// validateVotingPeriod validates the VotingPeriod param
func validateVotingPeriod(v interface{}) error {
	votingPeriod, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = votingPeriod

	return nil
}

// validateQuorum validates the Quorum param
func validateQuorum(v interface{}) error {
	quorum, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = quorum

	return nil
}

// validateTimelockSeconds validates the TimelockSeconds param
func validateTimelockSeconds(v interface{}) error {
	timelockSeconds, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = timelockSeconds

	return nil
}
