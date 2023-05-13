package types

// DONTCOVER

import (
	"cosmossdk.io/errors"
)

const (
	codeErrIbcElyDenomEmpty = uint32(iota) + 2 // NOTE: code 1 is reserved for internal errors
	codeErrIbcElyDenomInvalid
)

// x/elysium module sentinel errors
var (
	ErrIbcElyDenomEmpty   = errors.Register(ModuleName, codeErrIbcElyDenomEmpty, "ibc ely denom is not set")
	ErrIbcElyDenomInvalid = errors.Register(ModuleName, codeErrIbcElyDenomInvalid, "ibc ely denom is invalid")
	// this line is used by starport scaffolding # ibc/errors
)
