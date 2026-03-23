package types

import paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

var _ paramtypes.ParamSet = (*Params)(nil)

func NewParams() Params {
	return Params{
		MaxAgents:            25,   // genesis tier max agents
		CircuitBreakerMaxTx:  1000, // max tx per window
		CircuitBreakerWindow: 100,  // blocks per window
	}
}

func DefaultParams() Params { return NewParams() }

func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{}
}

func (p Params) Validate() error { return nil }
