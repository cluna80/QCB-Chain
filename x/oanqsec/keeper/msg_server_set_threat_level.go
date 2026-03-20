package keeper

import (
	"context"
	"fmt"
	"oan/x/oanqsec/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SetThreatLevel(goCtx context.Context, msg *types.MsgSetThreatLevel) (*types.MsgSetThreatLevelResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if msg.Level > 5 {
		return nil, fmt.Errorf("threat level must be 0-5 (0=none, 3=elevated, 5=critical)")
	}
	if msg.Evidence == "" {
		return nil, fmt.Errorf("evidence required to set threat level")
	}
	if msg.Justification == "" {
		return nil, fmt.Errorf("justification required to set threat level")
	}
	store := k.storeService.OpenKVStore(ctx)
	store.Set([]byte("current-threat-level"), []byte(fmt.Sprintf("%d", msg.Level)))
	if msg.Level >= 3 {
		store.Set([]byte("qs-only-mode"), []byte("1"))
		ctx.EventManager().EmitEvent(sdk.NewEvent("qs_only_mode_activated",
			sdk.NewAttribute("threat_level", fmt.Sprintf("%d", msg.Level)),
			sdk.NewAttribute("activated_by", msg.Creator),
			sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
		))
	} else {
		store.Delete([]byte("qs-only-mode"))
	}
	appliedAt := int32(ctx.BlockTime().Unix())
	ctx.EventManager().EmitEvent(sdk.NewEvent("threat_level_set",
		sdk.NewAttribute("level", fmt.Sprintf("%d", msg.Level)),
		sdk.NewAttribute("evidence", msg.Evidence),
		sdk.NewAttribute("set_by", msg.Creator),
		sdk.NewAttribute("applied_at", fmt.Sprintf("%d", appliedAt)),
	))
	return &types.MsgSetThreatLevelResponse{Level: msg.Level, AppliedAt: appliedAt}, nil
}
