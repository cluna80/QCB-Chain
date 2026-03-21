package keeper

import (
	"oan/x/oanwalletproto/types"
)

var _ types.QueryServer = Keeper{}
