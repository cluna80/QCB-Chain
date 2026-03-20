package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"oan/x/oanqsec/types"
)

func (k msgServer) RegisterQsKey(goCtx context.Context, msg *types.MsgRegisterQsKey) (*types.MsgRegisterQsKeyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	validAlgorithms := map[string]bool{
		"dilithium3": true,
		"sphincs+":   true,
		"kyber-1024": true,
		"sha3-512":   true,
	}
	if !validAlgorithms[msg.Algorithm] {
		return nil, fmt.Errorf("algorithm %s not supported — use dilithium3, sphincs+, kyber-1024, or sha3-512", msg.Algorithm)
	}
	if len(msg.PublicKeyHash) < 32 {
		return nil, fmt.Errorf("public key hash too short — minimum 32 characters")
	}
	store := k.storeService.OpenKVStore(ctx)
	existingKey := fmt.Sprintf("qs-key-%s-%s", msg.WalletAddr, msg.Algorithm)
	existing, _ := store.Get([]byte(existingKey))
	if existing != nil {
		return nil, fmt.Errorf("QS key already registered for %s — use rotate-qs-key to update", msg.WalletAddr)
	}
	keyId := fmt.Sprintf("qskey-%d-%s", ctx.BlockHeight(), msg.WalletAddr[:8])
	keyData := fmt.Sprintf("%s|%s|%s|%d", keyId, msg.PublicKeyHash, msg.Algorithm, ctx.BlockHeight())
	store.Set([]byte(existingKey), []byte(keyData))
	store.Set([]byte(fmt.Sprintf("qs-keyid-%s", keyId)), []byte(msg.WalletAddr))
	ctx.EventManager().EmitEvent(sdk.NewEvent("qs_key_registered",
		sdk.NewAttribute("key_id", keyId),
		sdk.NewAttribute("wallet", msg.WalletAddr),
		sdk.NewAttribute("algorithm", msg.Algorithm),
		sdk.NewAttribute("key_type", msg.KeyType),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgRegisterQsKeyResponse{KeyId: keyId, Registered: true, Algorithm: msg.Algorithm}, nil
}
