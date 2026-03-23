package types

import paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

var _ paramtypes.ParamSet = (*Params)(nil)

func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams() Params {
	return Params{
		MaxWalletBalance:       800_000_000_000_000,
		VerifiedMaxBalance:     400_000_000_000_000,
		UnverifiedMaxBalance:   4_000_000_000_000,
		MaxDailyReceive:        2_000_000_000_000,
		UnverifiedDailyReceive: 40_000_000_000,
		EpochLength:            14400,
		HardCap:                50_000_000_000_000_000,
		LaunchPhase:            "genesis",
		InflationBps:           100,
		ProtectionsEnabled:     true,
	}
}

func DefaultParams() Params { return NewParams() }

func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{}
}

func (p Params) Validate() error { return nil }
