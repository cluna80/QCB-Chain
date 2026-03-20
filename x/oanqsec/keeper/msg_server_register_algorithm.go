package keeper

import (
	"context"
	"fmt"
	"oan/x/oanqsec/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RegisterAlgorithm(goCtx context.Context, msg *types.MsgRegisterAlgorithm) (*types.MsgRegisterAlgorithmResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	validTypes := map[string]bool{
		"signature":    true,
		"key-exchange": true,
		"hash":         true,
		"hybrid":       true,
	}
	if !validTypes[msg.AlgorithmType] {
		return nil, fmt.Errorf("algorithmType must be signature, key-exchange, hash, or hybrid")
	}
	if msg.SecurityLevel < 128 {
		return nil, fmt.Errorf("security level must be at least 128 bits — OAN minimum standard")
	}
	if msg.Specification == "" {
		return nil, fmt.Errorf("specification cannot be empty — include NIST reference or paper")
	}
	store := k.storeService.OpenKVStore(ctx)
	algoKey := fmt.Sprintf("algorithm-%s", msg.AlgorithmId)
	existing, _ := store.Get([]byte(algoKey))
	if existing != nil {
		return nil, fmt.Errorf("algorithm %s already registered", msg.AlgorithmId)
	}
	algoData := fmt.Sprintf("%s|%s|%d|%s|%d|active",
		msg.AlgorithmId, msg.AlgorithmType, msg.SecurityLevel,
		msg.Specification, ctx.BlockHeight())
	store.Set([]byte(algoKey), []byte(algoData))
	ctx.EventManager().EmitEvent(sdk.NewEvent("algorithm_registered",
		sdk.NewAttribute("algorithm_id", msg.AlgorithmId),
		sdk.NewAttribute("type", msg.AlgorithmType),
		sdk.NewAttribute("security_level", fmt.Sprintf("%d", msg.SecurityLevel)),
		sdk.NewAttribute("registered_by", msg.Creator),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgRegisterAlgorithmResponse{AlgorithmId: msg.AlgorithmId, Registered: true}, nil
}
