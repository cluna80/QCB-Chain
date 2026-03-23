package keeper

import (
	"context"
	"fmt"
	"oan/x/oancomms/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RegisterMsgKey(goCtx context.Context, msg *types.MsgRegisterMsgKey) (*types.MsgRegisterMsgKeyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := k.storeService.OpenKVStore(ctx)

	validTypes := map[string]bool{
		"ed25519": true, "secp256k1": true, "dilithium3": true,
		"sphincs-plus": true, "falcon-1024": true,
	}
	if !validTypes[msg.KeyType] {
		return nil, fmt.Errorf("keyType must be ed25519, secp256k1, dilithium3, sphincs-plus, or falcon-1024")
	}
	if len(msg.PublicKeyHash) < 32 {
		return nil, fmt.Errorf("publicKeyHash must be at least 32 characters")
	}

	// One msg key per wallet per algorithm
	keyStoreKey := fmt.Sprintf("msgkey-%s-%s", msg.Creator, msg.KeyType)
	existing, _ := store.Get([]byte(keyStoreKey))
	if existing != nil {
		return nil, fmt.Errorf("msg key already registered for %s — revoke first", msg.KeyType)
	}

	store.Set([]byte(keyStoreKey),
		[]byte(fmt.Sprintf("%s|%s|%s|%s|%d|active",
			msg.KeyId, msg.KeyType, msg.PublicKeyHash, msg.Creator, ctx.BlockHeight())))
	store.Set([]byte(fmt.Sprintf("msgkeyid-%s", msg.KeyId)), []byte(msg.Creator))

	ctx.EventManager().EmitEvent(sdk.NewEvent("msg_key_registered",
		sdk.NewAttribute("key_id", msg.KeyId),
		sdk.NewAttribute("key_type", msg.KeyType),
		sdk.NewAttribute("owner", msg.Creator),
		sdk.NewAttribute("algorithm", msg.Algorithm),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgRegisterMsgKeyResponse{KeyId: msg.KeyId, Registered: true}, nil
}
