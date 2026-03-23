package keeper

import (
	"oan/x/oanprotocol/types"
)

var _ types.QueryServer = Keeper{}
