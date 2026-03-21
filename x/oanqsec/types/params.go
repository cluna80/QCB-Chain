package types

import paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

var _ paramtypes.ParamSet = (*Params)(nil)

func NewParams() Params {
	return Params{
		ThreatLevel:     0,            // 0=none, 1=monitor, 3=elevated, 5=critical
		HybridRequired:  false,        // not required at launch — DAO activates
		RotationPeriod:  10000,        // rotate QS keys every 10,000 blocks
		ActiveAlgorithm: "dilithium3", // NIST FIPS 204 primary
	}
}

func DefaultParams() Params { return NewParams() }

func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{}
}

func (p Params) Validate() error { return nil }
