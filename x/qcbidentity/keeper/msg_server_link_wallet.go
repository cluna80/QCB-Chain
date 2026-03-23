package keeper

import (
	"context"
	"fmt"
	"qcb/x/qcbidentity/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) LinkWallet(goCtx context.Context, msg *types.MsgLinkWallet) (*types.MsgLinkWalletResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	identity, found := k.GetIdentity(ctx, msg.Did)
	if !found {
		return nil, fmt.Errorf("identity %s not found", msg.Did)
	}
	identity.Owner = msg.WalletAddress
	identity.LastActive = int32(ctx.BlockTime().Unix())
	k.SetIdentity(ctx, identity)
	ctx.EventManager().EmitEvent(sdk.NewEvent("wallet_linked",
		sdk.NewAttribute("did", msg.Did),
		sdk.NewAttribute("wallet", msg.WalletAddress),
	))
	return &types.MsgLinkWalletResponse{Did: msg.Did, WalletAddress: msg.WalletAddress}, nil
}
