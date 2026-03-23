package types

import paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

var _ paramtypes.ParamSet = (*Params)(nil)

func NewParams() Params {
	return Params{
		MinGuardians:  3,     // minimum 3 guardians at all times
		MaxGuardians:  9,     // maximum 9 guardians on council
		VetoThreshold: 2,     // 2-of-N required for emergency pause
		PauseDuration: 86400, // 24 hour default pause duration
	}
}

func DefaultParams() Params { return NewParams() }

func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{}
}

func (p Params) Validate() error { return nil }
