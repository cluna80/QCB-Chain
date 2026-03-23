package keeper

import (
	"qcb/x/qcbcompute/types"
)

var _ types.QueryServer = Keeper{}
