package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"oan/x/oanwalletproto/types"
)

func (k msgServer) LockWallet(goCtx context.Context, msg *types.MsgLockWallet) (*types.MsgLockWalletResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := k.storeService.OpenKVStore(ctx)

	profileKey := fmt.Sprintf("wallet-profile-%s", msg.Creator)
	existing, _ := store.Get([]byte(profileKey))
	if existing == nil {
		return nil, fmt.Errorf("wallet profile not found")
	}
	if msg.Reason == "" {
		return nil, fmt.Errorf("reason required to lock wallet")
	}

	alreadyLocked, _ := store.Get([]byte(fmt.Sprintf("wallet-locked-%s", msg.Creator)))
	if alreadyLocked != nil {
		return nil, fmt.Errorf("wallet is already locked")
	}

	store.Set([]byte(fmt.Sprintf("wallet-locked-%s", msg.Creator)),
		[]byte(fmt.Sprintf("%s|%d", msg.Reason, ctx.BlockHeight())))

	ctx.EventManager().EmitEvent(sdk.NewEvent("wallet_locked",
		sdk.NewAttribute("wallet_id", msg.WalletId),
		sdk.NewAttribute("reason", msg.Reason),
		sdk.NewAttribute("owner", msg.Creator),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgLockWalletResponse{WalletId: msg.WalletId, Locked: true}, nil
}
