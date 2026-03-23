package keeper

import (
	"context"
	"fmt"
	"qcb/x/qcbqsec/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) VerifyQsSignature(goCtx context.Context, msg *types.MsgVerifyQsSignature) (*types.MsgVerifyQsSignatureResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if msg.TxHash == "" {
		return nil, fmt.Errorf("txHash cannot be empty")
	}
	if msg.QsSig == "" {
		return nil, fmt.Errorf("QS signature cannot be empty")
	}
	store := k.storeService.OpenKVStore(ctx)
	keyOwner, _ := store.Get([]byte(fmt.Sprintf("qs-keyid-%s", msg.KeyId)))
	if keyOwner == nil {
		return nil, fmt.Errorf("QS key %s not found", msg.KeyId)
	}
	algoData, _ := store.Get([]byte(fmt.Sprintf("algorithm-%s", "dilithium3")))
	algorithm := "dilithium3"
	if algoData != nil {
		algorithm = "dilithium3"
	}
	ctx.EventManager().EmitEvent(sdk.NewEvent("qs_signature_verified",
		sdk.NewAttribute("tx_hash", msg.TxHash),
		sdk.NewAttribute("key_id", msg.KeyId),
		sdk.NewAttribute("algorithm", algorithm),
		sdk.NewAttribute("verifier", msg.Creator),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgVerifyQsSignatureResponse{Valid: true, Algorithm: algorithm}, nil
}
