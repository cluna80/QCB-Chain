package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"oan/x/antirug/types"
)

func (k msgServer) RegisterToken(goCtx context.Context, msg *types.MsgRegisterToken) (*types.MsgRegisterTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := k.storeService.OpenKVStore(ctx)
	params := k.GetParams(ctx)

	// DORMANCY CHECK
	if !params.Enabled {
		return nil, fmt.Errorf("antirug module not yet active — registration optional until governance activates it")
	}

	if msg.TokenId == "" || msg.Symbol == "" {
		return nil, fmt.Errorf("tokenId and symbol cannot be empty")
	}
	if msg.MaxSupply == 0 {
		return nil, fmt.Errorf("maxSupply must be declared — hidden unlimited minting is not allowed")
	}

	minLock := params.MinLiquidityLockBlocks
	if minLock == 0 { minLock = 201600 } // 14 days default
	if msg.LiquidityLockBlocks < minLock {
		return nil, fmt.Errorf("liquidity must be locked for at least %d blocks (~14 days) — got %d", minLock, msg.LiquidityLockBlocks)
	}

	// Check not already registered
	tokenKey := fmt.Sprintf("antirug-token-%s", msg.TokenId)
	existing, _ := store.Get([]byte(tokenKey))
	if existing != nil {
		return nil, fmt.Errorf("token %s already registered", msg.TokenId)
	}

	// Calculate safety score
	safetyScore := uint64(50) // base score
	if msg.LiquidityLockBlocks >= 403200 { safetyScore += 20 } // 28+ days
	if msg.LiquidityLockBlocks >= 864000 { safetyScore += 10 } // 60+ days
	if msg.MaxSupply <= 1000000000       { safetyScore += 10 } // reasonable supply
	if msg.MaxSupply <= 100000000        { safetyScore += 10 } // tight supply

	tokenData := fmt.Sprintf("%s|%s|%s|%d|%d|%d|pending|%d",
		msg.TokenId, msg.TokenName, msg.Symbol,
		msg.MaxSupply, msg.LiquidityLockBlocks,
		ctx.BlockHeight(), safetyScore)
	store.Set([]byte(tokenKey), []byte(tokenData))
	store.Set([]byte(fmt.Sprintf("antirug-owner-%s", msg.TokenId)), []byte(msg.Creator))

	ctx.EventManager().EmitEvent(sdk.NewEvent("token_registered_antirug",
		sdk.NewAttribute("token_id", msg.TokenId),
		sdk.NewAttribute("symbol", msg.Symbol),
		sdk.NewAttribute("max_supply", fmt.Sprintf("%d", msg.MaxSupply)),
		sdk.NewAttribute("liquidity_lock_blocks", fmt.Sprintf("%d", msg.LiquidityLockBlocks)),
		sdk.NewAttribute("safety_score", fmt.Sprintf("%d", safetyScore)),
		sdk.NewAttribute("registrant", msg.Creator),
	))
	return &types.MsgRegisterTokenResponse{TokenId: msg.TokenId, Registered: true, SafetyScore: safetyScore}, nil
}
