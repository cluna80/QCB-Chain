package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"oan/x/oanqsec/types"
)

func (k msgServer) RotateQsKey(goCtx context.Context, msg *types.MsgRotateQsKey) (*types.MsgRotateQsKeyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if len(msg.NewPublicKeyHash) < 32 {
		return nil, fmt.Errorf("new public key hash too short — minimum 32 characters")
	}
	if msg.Reason == "" {
		return nil, fmt.Errorf("rotation reason cannot be empty")
	}
	store := k.storeService.OpenKVStore(ctx)
	rotationCooldownKey := fmt.Sprintf("qs-rotation-cooldown-%s", msg.Creator)
	lastRotation, _ := store.Get([]byte(rotationCooldownKey))
	if lastRotation != nil {
		lastBlock := int64(0)
		for i, b := range lastRotation {
			lastBlock |= int64(b) << (8 * i)
		}
		if ctx.BlockHeight()-lastBlock < 100 {
			return nil, fmt.Errorf("key rotation cooldown active — wait %d blocks", 100-(ctx.BlockHeight()-lastBlock))
		}
	}
	heightBytes := make([]byte, 8)
	h := ctx.BlockHeight()
	for i := 0; i < 8; i++ {
		heightBytes[i] = byte(h >> (8 * i))
	}
	store.Set([]byte(rotationCooldownKey), heightBytes)
	newKeyId := fmt.Sprintf("qskey-%d-%s-rotated", ctx.BlockHeight(), msg.Creator[:8])
	store.Set([]byte(fmt.Sprintf("qs-keyid-%s", newKeyId)), []byte(msg.Creator))
	store.Delete([]byte(fmt.Sprintf("qs-keyid-%s", msg.OldKeyId)))
	rotatedAt := int32(ctx.BlockTime().Unix())
	ctx.EventManager().EmitEvent(sdk.NewEvent("qs_key_rotated",
		sdk.NewAttribute("old_key_id", msg.OldKeyId),
		sdk.NewAttribute("new_key_id", newKeyId),
		sdk.NewAttribute("reason", msg.Reason),
		sdk.NewAttribute("rotated_by", msg.Creator),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgRotateQsKeyResponse{NewKeyId: newKeyId, RotatedAt: rotatedAt}, nil
}
