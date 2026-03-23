package keeper

import (
	"qcb/x/qcbwalletproto/types"
)

var _ types.QueryServer = Keeper{}
