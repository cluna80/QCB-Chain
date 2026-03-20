package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyMarketFee = []byte("MarketFee")
	// TODO: Determine the default value
	DefaultMarketFee uint64 = 0
)

var (
	KeyListingDuration = []byte("ListingDuration")
	// TODO: Determine the default value
	DefaultListingDuration uint64 = 0
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	marketFee uint64,
	listingDuration uint64,
) Params {
	return Params{
		MarketFee:       marketFee,
		ListingDuration: listingDuration,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultMarketFee,
		DefaultListingDuration,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMarketFee, &p.MarketFee, validateMarketFee),
		paramtypes.NewParamSetPair(KeyListingDuration, &p.ListingDuration, validateListingDuration),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateMarketFee(p.MarketFee); err != nil {
		return err
	}

	if err := validateListingDuration(p.ListingDuration); err != nil {
		return err
	}

	return nil
}

// validateMarketFee validates the MarketFee param
func validateMarketFee(v interface{}) error {
	marketFee, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = marketFee

	return nil
}

// validateListingDuration validates the ListingDuration param
func validateListingDuration(v interface{}) error {
	listingDuration, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = listingDuration

	return nil
}
