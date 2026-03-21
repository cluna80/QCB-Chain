package types

import paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

var _ paramtypes.ParamSet = (*Params)(nil)

func NewParams() Params {
	return Params{
		TaskFee:  1000000,  // 1 OANT to create a task
		UbiRate:  100000,   // 0.1 OANT per UBI claim
		MaxTasks: 10000,    // max 10,000 open tasks
	}
}

func DefaultParams() Params { return NewParams() }

func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{}
}

func (p Params) Validate() error { return nil }
