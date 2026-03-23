package keeper

import (
	"context"
	"fmt"
	"qcb/x/qcbqsec/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) EmergencyCryptoSwap(goCtx context.Context, msg *types.MsgEmergencyCryptoSwap) (*types.MsgEmergencyCryptoSwapResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	validUrgency := map[string]bool{
		"low": true, "medium": true, "high": true, "critical": true,
	}
	if !validUrgency[msg.Urgency] {
		return nil, fmt.Errorf("urgency must be low, medium, high, or critical")
	}
	if msg.Evidence == "" {
		return nil, fmt.Errorf("evidence required for emergency crypto swap")
	}
	if msg.FromAlgorithm == msg.ToAlgorithm {
		return nil, fmt.Errorf("from and to algorithms cannot be the same")
	}
	store := k.storeService.OpenKVStore(ctx)
	toAlgoKey := fmt.Sprintf("algorithm-%s", msg.ToAlgorithm)
	toAlgo, _ := store.Get([]byte(toAlgoKey))
	if toAlgo == nil {
		return nil, fmt.Errorf("target algorithm %s not registered — cannot swap to unknown algorithm", msg.ToAlgorithm)
	}
	deprecatedKey := fmt.Sprintf("deprecated-%s", msg.ToAlgorithm)
	isDeprecated, _ := store.Get([]byte(deprecatedKey))
	if isDeprecated != nil {
		return nil, fmt.Errorf("target algorithm %s is deprecated — cannot swap to deprecated algorithm", msg.ToAlgorithm)
	}
	store.Set([]byte("active-algorithm"), []byte(msg.ToAlgorithm))
	store.Set([]byte(fmt.Sprintf("deprecated-%s", msg.FromAlgorithm)), []byte("1"))
	if msg.Urgency == "critical" || msg.Urgency == "high" {
		store.Set([]byte("qs-only-mode"), []byte("1"))
	}
	activatedAt := int32(ctx.BlockTime().Unix())
	ctx.EventManager().EmitEvent(sdk.NewEvent("emergency_crypto_swap",
		sdk.NewAttribute("from_algorithm", msg.FromAlgorithm),
		sdk.NewAttribute("to_algorithm", msg.ToAlgorithm),
		sdk.NewAttribute("urgency", msg.Urgency),
		sdk.NewAttribute("evidence", msg.Evidence),
		sdk.NewAttribute("activated_by", msg.Creator),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgEmergencyCryptoSwapResponse{
		Success: true, ActivatedAt: activatedAt, NewAlgorithm: msg.ToAlgorithm,
	}, nil
}
