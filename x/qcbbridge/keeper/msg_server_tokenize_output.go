package keeper

import (
	"context"
	"fmt"
	"qcb/x/qcbbridge/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) TokenizeOutput(goCtx context.Context, msg *types.MsgTokenizeOutput) (*types.MsgTokenizeOutputResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	validOutputTypes := map[string]bool{
		"trade": true, "research": true, "strategy": true,
		"code": true, "analysis": true, "prediction": true,
	}
	if !validOutputTypes[msg.OutputType] {
		return nil, fmt.Errorf("outputType must be trade, research, strategy, code, analysis, or prediction")
	}
	if msg.ContentHash == "" {
		return nil, fmt.Errorf("contentHash cannot be empty")
	}
	if msg.Value == 0 {
		return nil, fmt.Errorf("value must be greater than 0")
	}
	tokenId := fmt.Sprintf("token-%d-%s-%s", ctx.BlockHeight(), msg.AgentId, msg.OutputType)
	store := k.storeService.OpenKVStore(ctx)
	store.Set([]byte(fmt.Sprintf("token-%s", tokenId)),
		[]byte(fmt.Sprintf("%s|%s|%s|%s|%d|%s|active",
			tokenId, msg.AgentId, msg.OutputType, msg.ContentHash, msg.Value, msg.Creator)))
	store.Set([]byte(fmt.Sprintf("token-owner-%s", tokenId)), []byte(msg.Creator))
	ctx.EventManager().EmitEvent(sdk.NewEvent("output_tokenized",
		sdk.NewAttribute("token_id", tokenId),
		sdk.NewAttribute("agent_id", msg.AgentId),
		sdk.NewAttribute("output_type", msg.OutputType),
		sdk.NewAttribute("content_hash", msg.ContentHash),
		sdk.NewAttribute("value", fmt.Sprintf("%d", msg.Value)),
		sdk.NewAttribute("owner", msg.Creator),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgTokenizeOutputResponse{TokenId: tokenId, ContentHash: msg.ContentHash, Value: msg.Value}, nil
}
