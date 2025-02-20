package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/faucet module sentinel errors
var (
	ErrInvalidSigner = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrSample        = sdkerrors.Register(ModuleName, 1101, "sample error")
	ErrExceedsMaxRequest = sdkerrors.Register(ModuleName, 1, "request exceeds maximum amount per request")
    ErrExceedsMaxAddress = sdkerrors.Register(ModuleName, 2, "request exceeds maximum amount per address")

)
