package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"oan/x/antirug/types"
)

func (k msgServer) SetModuleActive(goCtx context.Context, msg *types.MsgSetModuleActive) (*types.MsgSetModuleActiveResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := k.storeService.OpenKVStore(ctx)

	guardianKey := fmt.Sprintf("guardian-auth-%s", msg.Creator)
	isGuardian, _ := store.Get([]byte(guardianKey))
	isDaoModule := msg.Creator == "oan10d07y265gmmuvt4z0w9aw880jnsr700j6zn9kn"
	if isGuardian == nil && !isDaoModule {
		return nil, fmt.Errorf("only guardians or DAO governance can activate antirug module")
	}
	if msg.Reason == "" {
		return nil, fmt.Errorf("reason required to change module state")
	}
	if msg.Active != "true" && msg.Active != "false" {
		return nil, fmt.Errorf("active must be 'true' or 'false'")
	}

	params := k.GetParams(ctx)
	params.Enabled = msg.Active == "true"
	k.SetParams(ctx, params)

	activatedAt := int32(ctx.BlockTime().Unix())
	store.Set([]byte("antirug-enabled"), []byte(fmt.Sprintf("%s|%d|%s", msg.Active, ctx.BlockHeight(), msg.Reason)))

	ctx.EventManager().EmitEvent(sdk.NewEvent("antirug_module_status_changed",
		sdk.NewAttribute("status", msg.Active),
		sdk.NewAttribute("reason", msg.Reason),
		sdk.NewAttribute("changed_by", msg.Creator),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))

	responseActive := false
	if msg.Active == "true" { responseActive = true }
	return &types.MsgSetModuleActiveResponse{Active: responseActive, ActivatedAt: activatedAt}, nil
}
