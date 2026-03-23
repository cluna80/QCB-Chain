package keeper

import (
	"context"
	"fmt"
	"qcb/x/qcbwalletproto/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SetEncryptionKey(goCtx context.Context, msg *types.MsgSetEncryptionKey) (*types.MsgSetEncryptionKeyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := k.storeService.OpenKVStore(ctx)

	validKeyTypes := map[string]bool{
		"x25519": true, "kyber-1024": true, "rsa-4096": true,
	}
	if !validKeyTypes[msg.KeyType] {
		return nil, fmt.Errorf("keyType must be x25519, kyber-1024, or rsa-4096")
	}
	if len(msg.EncKeyHash) < 32 {
		return nil, fmt.Errorf("encKeyHash must be at least 32 characters")
	}

	profileKey := fmt.Sprintf("wallet-profile-%s", msg.Creator)
	existing, _ := store.Get([]byte(profileKey))
	if existing == nil {
		return nil, fmt.Errorf("wallet profile not found — register-wallet-profile first")
	}

	// Check rotation cooldown
	params := k.GetParams(ctx)
	rotationBlocks := int64(params.KeyRotationBlocks)
	if rotationBlocks == 0 {
		rotationBlocks = 1000
	}
	lastRotKey := fmt.Sprintf("wallet-enc-key-rotated-%s", msg.Creator)
	lastRotBytes, _ := store.Get([]byte(lastRotKey))
	if lastRotBytes != nil {
		lastRot := int64(0)
		fmt.Sscanf(string(lastRotBytes), "%d", &lastRot)
		if ctx.BlockHeight()-lastRot < rotationBlocks {
			blocksLeft := rotationBlocks - (ctx.BlockHeight() - lastRot)
			return nil, fmt.Errorf("key rotation cooldown — %d blocks remaining", blocksLeft)
		}
	}

	store.Set([]byte(fmt.Sprintf("wallet-enc-key-%s", msg.Creator)),
		[]byte(fmt.Sprintf("%s|%s|%d", msg.EncKeyHash, msg.KeyType, ctx.BlockHeight())))
	store.Set([]byte(lastRotKey), []byte(fmt.Sprintf("%d", ctx.BlockHeight())))

	ctx.EventManager().EmitEvent(sdk.NewEvent("wallet_enc_key_set",
		sdk.NewAttribute("wallet_id", msg.WalletId),
		sdk.NewAttribute("key_type", msg.KeyType),
		sdk.NewAttribute("owner", msg.Creator),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgSetEncryptionKeyResponse{WalletId: msg.WalletId, KeySet: true}, nil
}
