package types

import paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

var _ paramtypes.ParamSet = (*Params)(nil)

func NewParams() Params {
	return Params{
		MaxMsgSize:          65536,  // 64KB max message header
		MsgTtlBlocks:        1000,   // messages expire after 1000 blocks
		EncryptionRequired:  false,  // DAO activates when ready
	}
}

func DefaultParams() Params { return NewParams() }
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs { return paramtypes.ParamSetPairs{} }
func (p Params) Validate() error { return nil }
