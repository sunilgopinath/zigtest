package types

import (
	fmt "fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = (*Params)(nil)

const (
    DefaultMaxPerRequest = uint64(100_000_000) // 100 tokens assuming 6 decimals
    DefaultMaxPerAddress = uint64(500_000_000) // 500 tokens assuming 6 decimals
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams() Params {
	return Params{
        MaxPerRequest: DefaultMaxPerRequest,
        MaxPerAddress: DefaultMaxPerAddress,
    }
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams()
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if p.MaxPerRequest == 0 {
        return fmt.Errorf("max per request cannot be 0")
    }

    if p.MaxPerAddress == 0 {
        return fmt.Errorf("max per address cannot be 0")
    }

    if p.MaxPerAddress < p.MaxPerRequest {
        return fmt.Errorf("max per address (%d) cannot be less than max per request (%d)", 
            p.MaxPerAddress, p.MaxPerRequest)
    }

    return nil
}
