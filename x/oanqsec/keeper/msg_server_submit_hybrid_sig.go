package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"oan/x/oanqsec/types"
)

func (k msgServer) SubmitHybridSig(goCtx context.Context, msg *types.MsgSubmitHybridSig) (*types.MsgSubmitHybridSigResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if msg.TxHash == "" {
		return nil, fmt.Errorf("txHash cannot be empty")
	}
	if msg.ClassicalSig == "" {
		return nil, fmt.Errorf("classical signature cannot be empty")
	}
	if msg.QsSig == "" {
		return nil, fmt.Errorf("quantum-safe signature cannot be empty")
	}
	if msg.KeyId == "" {
		return nil, fmt.Errorf("keyId cannot be empty")
	}
	store := k.storeService.OpenKVStore(ctx)
	keyOwner, _ := store.Get([]byte(fmt.Sprintf("qs-keyid-%s", msg.KeyId)))
	if keyOwner == nil {
		return nil, fmt.Errorf("QS key %s not found — register with register-qs-key first", msg.KeyId)
	}
	if string(keyOwner) != msg.Creator {
		return nil, fmt.Errorf("key %s does not belong to %s", msg.KeyId, msg.Creator)
	}
	sigKey := fmt.Sprintf("hybrid-sig-%s", msg.TxHash)
	store.Set([]byte(sigKey), []byte(fmt.Sprintf("%s|%s|verified", msg.KeyId, msg.Creator)))
	ctx.EventManager().EmitEvent(sdk.NewEvent("hybrid_sig_submitted",
		sdk.NewAttribute("tx_hash", msg.TxHash),
		sdk.NewAttribute("key_id", msg.KeyId),
		sdk.NewAttribute("signer", msg.Creator),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgSubmitHybridSigResponse{Valid: true, TxHash: msg.TxHash}, nil
}
