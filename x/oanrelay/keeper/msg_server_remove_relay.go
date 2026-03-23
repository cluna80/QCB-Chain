package keeper

import (
	"context"
	"fmt"
	"oan/x/oanrelay/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RemoveRelay(goCtx context.Context, msg *types.MsgRemoveRelay) (*types.MsgRemoveRelayResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := k.storeService.OpenKVStore(ctx)

	ownerBytes, _ := store.Get([]byte(fmt.Sprintf("relay-owner-%s", msg.RelayId)))
	if ownerBytes == nil {
		return nil, fmt.Errorf("relay %s not found", msg.RelayId)
	}
	if string(ownerBytes) != msg.Creator {
		return nil, fmt.Errorf("only the relay operator can remove it")
	}
	if msg.Reason == "" {
		return nil, fmt.Errorf("reason required to remove a relay")
	}

	jailed, _ := store.Get([]byte(fmt.Sprintf("relay-jailed-%s", msg.RelayId)))
	if jailed != nil {
		return nil, fmt.Errorf("jailed relay cannot be removed — contact guardians")
	}

	store.Delete([]byte(fmt.Sprintf("relay-%s", msg.RelayId)))
	store.Delete([]byte(fmt.Sprintf("relay-owner-%s", msg.RelayId)))

	ctx.EventManager().EmitEvent(sdk.NewEvent("relay_removed",
		sdk.NewAttribute("relay_id", msg.RelayId),
		sdk.NewAttribute("reason", msg.Reason),
		sdk.NewAttribute("removed_by", msg.Creator),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgRemoveRelayResponse{RelayId: msg.RelayId, Removed: true}, nil
}
