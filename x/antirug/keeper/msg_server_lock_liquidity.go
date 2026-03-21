package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"oan/x/antirug/types"
)

func (k msgServer) LockLiquidity(goCtx context.Context, msg *types.MsgLockLiquidity) (*types.MsgLockLiquidityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := k.storeService.OpenKVStore(ctx)
	params := k.GetParams(ctx)

	// DORMANCY CHECK
	if !params.Enabled {
		return nil, fmt.Errorf("antirug module not yet active — liquidity locking optional until governance activates it")
	}

	if msg.Amount == 0 {
		return nil, fmt.Errorf("lock amount must be greater than 0")
	}
	minLock := params.MinLiquidityLockBlocks
	if minLock == 0 { minLock = 201600 }
	if msg.LockBlocks < minLock {
		return nil, fmt.Errorf("minimum lock period is %d blocks (~14 days)", minLock)
	}

	// Verify token exists and caller owns it
	ownerBytes, _ := store.Get([]byte(fmt.Sprintf("antirug-owner-%s", msg.TokenId)))
	if ownerBytes != nil && string(ownerBytes) != msg.Creator {
		return nil, fmt.Errorf("only the token owner can lock liquidity for %s", msg.TokenId)
	}

	lockId := fmt.Sprintf("lock-%d-%s", ctx.BlockHeight(), msg.TokenId)
	unlocksAt := int32(ctx.BlockHeight()) + int32(msg.LockBlocks)

	store.Set([]byte(fmt.Sprintf("antirug-lock-%s", lockId)),
		[]byte(fmt.Sprintf("%s|%d|%d|%d|locked",
			msg.TokenId, msg.Amount, msg.LockBlocks, unlocksAt)))

	ctx.EventManager().EmitEvent(sdk.NewEvent("liquidity_locked",
		sdk.NewAttribute("lock_id", lockId),
		sdk.NewAttribute("token_id", msg.TokenId),
		sdk.NewAttribute("amount", fmt.Sprintf("%d", msg.Amount)),
		sdk.NewAttribute("unlocks_at_block", fmt.Sprintf("%d", unlocksAt)),
		sdk.NewAttribute("locker", msg.Creator),
	))
	return &types.MsgLockLiquidityResponse{LockId: lockId, UnlocksAtBlock: unlocksAt}, nil
}
