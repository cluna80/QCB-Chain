package keeper

import (
	"qcb/x/agent/types"
)

var _ types.QueryServer = Keeper{}
