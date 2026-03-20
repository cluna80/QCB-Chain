package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"oan/x/oanqsec/types"
)

func (k msgServer) DeprecateAlgorithm(goCtx context.Context, msg *types.MsgDeprecateAlgorithm) (*types.MsgDeprecateAlgorithmResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if msg.Reason == "" {
		return nil, fmt.Errorf("deprecation reason cannot be empty")
	}
	if msg.AlgorithmId == msg.ReplacementId {
		return nil, fmt.Errorf("replacement algorithm cannot be the same as deprecated algorithm")
	}
	store := k.storeService.OpenKVStore(ctx)
	algoKey := fmt.Sprintf("algorithm-%s", msg.AlgorithmId)
	existing, _ := store.Get([]byte(algoKey))
	if existing == nil {
		return nil, fmt.Errorf("algorithm %s not found", msg.AlgorithmId)
	}
	replacementKey := fmt.Sprintf("algorithm-%s", msg.ReplacementId)
	replacement, _ := store.Get([]byte(replacementKey))
	if replacement == nil {
		return nil, fmt.Errorf("replacement algorithm %s not registered — register it first", msg.ReplacementId)
	}
	deprecatedData := fmt.Sprintf("%s|deprecated|%s|%s|%d",
		msg.AlgorithmId, msg.Reason, msg.ReplacementId, ctx.BlockHeight())
	store.Set([]byte(algoKey), []byte(deprecatedData))
	store.Set([]byte(fmt.Sprintf("deprecated-%s", msg.AlgorithmId)), []byte("1"))
	deprecatedAt := int32(ctx.BlockTime().Unix())
	ctx.EventManager().EmitEvent(sdk.NewEvent("algorithm_deprecated",
		sdk.NewAttribute("algorithm_id", msg.AlgorithmId),
		sdk.NewAttribute("reason", msg.Reason),
		sdk.NewAttribute("replacement", msg.ReplacementId),
		sdk.NewAttribute("deprecated_by", msg.Creator),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgDeprecateAlgorithmResponse{Success: true, DeprecatedAt: deprecatedAt}, nil
}
