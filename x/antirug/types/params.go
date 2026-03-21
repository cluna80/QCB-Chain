package types

import paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

var _ paramtypes.ParamSet = (*Params)(nil)

func NewParams() Params {
	return Params{
		Enabled:                  false,    // DORMANT at launch — governance activates
		MinLiquidityLockBlocks:   201600,   // 14 days at 6s blocks
		MaxMintPerBlock:          1000000,  // max mint per block
		TimelockBlocks:           14400,    // 24h timelock on upgrades
		CircuitBreakerEnabled:    false,    // off until governance activates
		DaoApprovalThreshold:     1000000,  // 1M OANT market cap triggers DAO review
	}
}

func DefaultParams() Params { return NewParams() }

func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{}
}

func (p Params) Validate() error { return nil }
