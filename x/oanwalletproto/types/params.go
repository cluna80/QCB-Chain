package types

import paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

var _ paramtypes.ParamSet = (*Params)(nil)

func NewParams() Params {
	return Params{
		MaxWalletsPerDid:  5,     // 5 wallets per DID
		KeyRotationBlocks: 1000,  // 1000 block cooldown on key rotation
		PqRequired:        false, // DAO activates when ready
	}
}

func DefaultParams() Params { return NewParams() }
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs { return paramtypes.ParamSetPairs{} }
func (p Params) Validate() error { return nil }
