package keeper

import (
	"oan/x/oannode/types"
)

var _ types.QueryServer = Keeper{}
