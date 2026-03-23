package keeper

import (
	"context"
	"fmt"
	"qcb/x/qcbcomms/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AckMsg(goCtx context.Context, msg *types.MsgAckMsg) (*types.MsgAckMsgResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := k.storeService.OpenKVStore(ctx)

	validAckTypes := map[string]bool{
		"delivered": true, "read": true, "rejected": true, "expired": true,
	}
	if !validAckTypes[msg.AckType] {
		return nil, fmt.Errorf("ackType must be delivered, read, rejected, or expired")
	}

	msgKey := fmt.Sprintf("msgheader-%s", msg.MsgId)
	msgData, _ := store.Get([]byte(msgKey))
	if msgData == nil {
		return nil, fmt.Errorf("message %s not found", msg.MsgId)
	}

	// Update status
	store.Set([]byte(fmt.Sprintf("msgack-%s", msg.MsgId)),
		[]byte(fmt.Sprintf("%s|%s|%d", msg.AckType, msg.Creator, ctx.BlockHeight())))

	ctx.EventManager().EmitEvent(sdk.NewEvent("msg_acked",
		sdk.NewAttribute("msg_id", msg.MsgId),
		sdk.NewAttribute("ack_type", msg.AckType),
		sdk.NewAttribute("acker", msg.Creator),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgAckMsgResponse{MsgId: msg.MsgId, Acked: true}, nil
}
