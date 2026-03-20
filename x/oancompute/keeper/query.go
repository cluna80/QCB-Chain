package keeper

import (
	"oan/x/oancompute/types"
)

var _ types.QueryServer = Keeper{}
