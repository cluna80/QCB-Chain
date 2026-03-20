package keeper

import (
	"oan/x/oanidentity/types"
)

var _ types.QueryServer = Keeper{}
