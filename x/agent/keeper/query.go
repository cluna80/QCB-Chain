package keeper

import (
	"oan/x/agent/types"
)

var _ types.QueryServer = Keeper{}
