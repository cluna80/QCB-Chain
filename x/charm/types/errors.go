package types

import "cosmossdk.io/errors"

var (
	ErrConfinementViolation = errors.Register(ModuleName, 1, "charm confinement violation")
	ErrAlreadyPaired        = errors.Register(ModuleName, 2, "address already has a charmed pair")
	ErrPairNotFound         = errors.Register(ModuleName, 3, "charmed pair not found")
	ErrInvalidTier          = errors.Register(ModuleName, 4, "invalid address tier")
	ErrUnauthorized         = errors.Register(ModuleName, 5, "unauthorized")
)
