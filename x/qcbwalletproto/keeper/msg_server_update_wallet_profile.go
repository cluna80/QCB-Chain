package keeper

import (
	"context"
	"fmt"
	"qcb/x/qcbwalletproto/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateWalletProfile(goCtx context.Context, msg *types.MsgUpdateWalletProfile) (*types.MsgUpdateWalletProfileResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := k.storeService.OpenKVStore(ctx)

	profileKey := fmt.Sprintf("wallet-profile-%s", msg.Creator)
	existing, _ := store.Get([]byte(profileKey))
	if existing == nil {
		return nil, fmt.Errorf("wallet profile not found — register-wallet-profile first")
	}
	if msg.DisplayName == "" {
		return nil, fmt.Errorf("displayName cannot be empty")
	}

	store.Set([]byte(profileKey),
		[]byte(fmt.Sprintf("%s|updated|%s|%s|%d",
			msg.WalletId, msg.DisplayName, msg.AvatarHash, ctx.BlockHeight())))

	ctx.EventManager().EmitEvent(sdk.NewEvent("wallet_profile_updated",
		sdk.NewAttribute("wallet_id", msg.WalletId),
		sdk.NewAttribute("display_name", msg.DisplayName),
		sdk.NewAttribute("owner", msg.Creator),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgUpdateWalletProfileResponse{WalletId: msg.WalletId, Updated: true}, nil
}
