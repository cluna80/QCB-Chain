package keeper

import (
	"context"
	"fmt"
	"oan/x/oanbridge/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RegisterChain(goCtx context.Context, msg *types.MsgRegisterChain) (*types.MsgRegisterChainResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	validBridgeTypes := map[string]bool{
		"ibc": true, "axelar": true, "wormhole": true, "native": true, "ftso": true,
	}
	if !validBridgeTypes[msg.BridgeType] {
		return nil, fmt.Errorf("bridgeType must be ibc, axelar, wormhole, native, or ftso")
	}
	if msg.Endpoint == "" {
		return nil, fmt.Errorf("endpoint cannot be empty")
	}
	store := k.storeService.OpenKVStore(ctx)
	chainKey := fmt.Sprintf("registered-chain-%s", msg.ChainId)
	existing, _ := store.Get([]byte(chainKey))
	if existing != nil {
		return nil, fmt.Errorf("chain %s already registered", msg.ChainId)
	}
	chainData := fmt.Sprintf("%s|%s|%s|%s|active|%d",
		msg.ChainId, msg.ChainName, msg.BridgeType, msg.Endpoint, ctx.BlockHeight())
	store.Set([]byte(chainKey), []byte(chainData))
	ctx.EventManager().EmitEvent(sdk.NewEvent("chain_registered",
		sdk.NewAttribute("chain_id", msg.ChainId),
		sdk.NewAttribute("chain_name", msg.ChainName),
		sdk.NewAttribute("bridge_type", msg.BridgeType),
		sdk.NewAttribute("registered_by", msg.Creator),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgRegisterChainResponse{ChainId: msg.ChainId, Registered: true}, nil
}
