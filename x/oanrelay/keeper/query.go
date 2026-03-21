package keeper

import (
	"oan/x/oanrelay/types"
)

var _ types.QueryServer = Keeper{}
