package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyMatchFee = []byte("MatchFee")
	// TODO: Determine the default value
	DefaultMatchFee uint64 = 0
)

var (
	KeyStadiumMintFee = []byte("StadiumMintFee")
	// TODO: Determine the default value
	DefaultStadiumMintFee uint64 = 0
)

var (
	KeyPredictionCut = []byte("PredictionCut")
	// TODO: Determine the default value
	DefaultPredictionCut uint64 = 0
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	matchFee uint64,
	stadiumMintFee uint64,
	predictionCut uint64,
) Params {
	return Params{
		MatchFee:       matchFee,
		StadiumMintFee: stadiumMintFee,
		PredictionCut:  predictionCut,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultMatchFee,
		DefaultStadiumMintFee,
		DefaultPredictionCut,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMatchFee, &p.MatchFee, validateMatchFee),
		paramtypes.NewParamSetPair(KeyStadiumMintFee, &p.StadiumMintFee, validateStadiumMintFee),
		paramtypes.NewParamSetPair(KeyPredictionCut, &p.PredictionCut, validatePredictionCut),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateMatchFee(p.MatchFee); err != nil {
		return err
	}

	if err := validateStadiumMintFee(p.StadiumMintFee); err != nil {
		return err
	}

	if err := validatePredictionCut(p.PredictionCut); err != nil {
		return err
	}

	return nil
}

// validateMatchFee validates the MatchFee param
func validateMatchFee(v interface{}) error {
	matchFee, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = matchFee

	return nil
}

// validateStadiumMintFee validates the StadiumMintFee param
func validateStadiumMintFee(v interface{}) error {
	stadiumMintFee, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = stadiumMintFee

	return nil
}

// validatePredictionCut validates the PredictionCut param
func validatePredictionCut(v interface{}) error {
	predictionCut, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = predictionCut

	return nil
}
