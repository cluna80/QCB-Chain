package types

import paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

var _ paramtypes.ParamSet = (*Params)(nil)

func NewParams() Params {
	return Params{
		MinValidatorStake: 10000000000, // 10,000 OANT in charmbits
		MinInferenceStake: 6000000000,  // 6,000 OANT in charmbits
		MinBridgeStake:    1600000000,  // 1,600 OANT in charmbits
		MinLightStake:     400000000,   // 400 OANT in charmbits
		UnbondingBlocks:   302400,      // 21 days at 6s blocks
		MaxWalletStakePct: 5,           // 5% whale cap
		EpochLength:       14400,       // 1 day at 6s blocks
	}
}

func DefaultParams() Params { return NewParams() }

func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{}
}

func (p Params) Validate() error { return nil }
