package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"oan/x/oanwalletproto/types"
)

func (k msgServer) SetWalletPermissions(goCtx context.Context, msg *types.MsgSetWalletPermissions) (*types.MsgSetWalletPermissionsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := k.storeService.OpenKVStore(ctx)

	profileKey := fmt.Sprintf("wallet-profile-%s", msg.Creator)
	existing, _ := store.Get([]byte(profileKey))
	if existing == nil {
		return nil, fmt.Errorf("wallet profile not found — register-wallet-profile first")
	}

	store.Set([]byte(fmt.Sprintf("wallet-perms-%s", msg.Creator)),
		[]byte(fmt.Sprintf("%v|%v|%v", msg.AllowMsgs, msg.AllowAiAgent, msg.AllowBridge)))

	ctx.EventManager().EmitEvent(sdk.NewEvent("wallet_permissions_set",
		sdk.NewAttribute("wallet_id", msg.WalletId),
		sdk.NewAttribute("allow_msgs", fmt.Sprintf("%v", msg.AllowMsgs)),
		sdk.NewAttribute("allow_ai_agent", fmt.Sprintf("%v", msg.AllowAiAgent)),
		sdk.NewAttribute("allow_bridge", fmt.Sprintf("%v", msg.AllowBridge)),
		sdk.NewAttribute("owner", msg.Creator),
	))
	return &types.MsgSetWalletPermissionsResponse{WalletId: msg.WalletId, PermissionsSet: true}, nil
}
