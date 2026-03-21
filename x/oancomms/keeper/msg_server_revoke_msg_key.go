package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"oan/x/oancomms/types"
)

func (k msgServer) RevokeMsgKey(goCtx context.Context, msg *types.MsgRevokeMsgKey) (*types.MsgRevokeMsgKeyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := k.storeService.OpenKVStore(ctx)

	ownerBytes, _ := store.Get([]byte(fmt.Sprintf("msgkeyid-%s", msg.KeyId)))
	if ownerBytes == nil {
		return nil, fmt.Errorf("msg key %s not found", msg.KeyId)
	}
	if string(ownerBytes) != msg.Creator {
		return nil, fmt.Errorf("only the key owner can revoke it")
	}
	if msg.Reason == "" {
		return nil, fmt.Errorf("reason required to revoke a msg key")
	}

	store.Set([]byte(fmt.Sprintf("msgkeyid-%s", msg.KeyId)),
		[]byte(fmt.Sprintf("revoked|%s|%d", msg.Reason, ctx.BlockHeight())))

	ctx.EventManager().EmitEvent(sdk.NewEvent("msg_key_revoked",
		sdk.NewAttribute("key_id", msg.KeyId),
		sdk.NewAttribute("reason", msg.Reason),
		sdk.NewAttribute("revoked_by", msg.Creator),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgRevokeMsgKeyResponse{KeyId: msg.KeyId, Revoked: true}, nil
}
