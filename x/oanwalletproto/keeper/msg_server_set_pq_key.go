package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"oan/x/oanwalletproto/types"
)

func (k msgServer) SetPqKey(goCtx context.Context, msg *types.MsgSetPqKey) (*types.MsgSetPqKeyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := k.storeService.OpenKVStore(ctx)

	validAlgos := map[string]bool{
		"dilithium3": true, "sphincs-plus": true,
		"falcon-1024": true, "kyber-1024": true,
	}
	if !validAlgos[msg.Algorithm] {
		return nil, fmt.Errorf("algorithm must be dilithium3, sphincs-plus, falcon-1024, or kyber-1024")
	}
	if len(msg.PqKeyHash) < 32 {
		return nil, fmt.Errorf("pqKeyHash must be at least 32 characters")
	}

	profileKey := fmt.Sprintf("wallet-profile-%s", msg.Creator)
	existing, _ := store.Get([]byte(profileKey))
	if existing == nil {
		return nil, fmt.Errorf("wallet profile not found — register-wallet-profile first")
	}

	store.Set([]byte(fmt.Sprintf("wallet-pq-key-%s", msg.Creator)),
		[]byte(fmt.Sprintf("%s|%s|%d", msg.PqKeyHash, msg.Algorithm, ctx.BlockHeight())))

	ctx.EventManager().EmitEvent(sdk.NewEvent("wallet_pq_key_set",
		sdk.NewAttribute("wallet_id", msg.WalletId),
		sdk.NewAttribute("algorithm", msg.Algorithm),
		sdk.NewAttribute("owner", msg.Creator),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgSetPqKeyResponse{WalletId: msg.WalletId, PqKeySet: true}, nil
}
