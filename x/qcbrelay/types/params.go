package types

import paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

var _ paramtypes.ParamSet = (*Params)(nil)

func NewParams() Params {
	return Params{
		MinRelayStake:      1600000000, // 1,600 OANT min stake
		MaxRelaysPerRegion: 100,        // 100 relays per region max
		SlashRate:          5,          // 5% slash default
	}
}

func DefaultParams() Params                               { return NewParams() }
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs { return paramtypes.ParamSetPairs{} }
func (p Params) Validate() error                          { return nil }
