package keeper

import (
	"context"
	"fmt"
	"qcb/x/qcbwalletproto/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UnlockWallet(goCtx context.Context, msg *types.MsgUnlockWallet) (*types.MsgUnlockWalletResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := k.storeService.OpenKVStore(ctx)

	profileKey := fmt.Sprintf("wallet-profile-%s", msg.Creator)
	existing, _ := store.Get([]byte(profileKey))
	if existing == nil {
		return nil, fmt.Errorf("wallet profile not found")
	}
	if msg.Proof == "" {
		return nil, fmt.Errorf("proof required to unlock wallet")
	}

	lockedKey := fmt.Sprintf("wallet-locked-%s", msg.Creator)
	locked, _ := store.Get([]byte(lockedKey))
	if locked == nil {
		return nil, fmt.Errorf("wallet is not locked")
	}

	store.Delete([]byte(lockedKey))

	ctx.EventManager().EmitEvent(sdk.NewEvent("wallet_unlocked",
		sdk.NewAttribute("wallet_id", msg.WalletId),
		sdk.NewAttribute("owner", msg.Creator),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgUnlockWalletResponse{WalletId: msg.WalletId, Unlocked: true}, nil
}
