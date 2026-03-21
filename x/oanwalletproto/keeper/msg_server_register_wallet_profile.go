package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"oan/x/oanwalletproto/types"
)

func (k msgServer) RegisterWalletProfile(goCtx context.Context, msg *types.MsgRegisterWalletProfile) (*types.MsgRegisterWalletProfileResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := k.storeService.OpenKVStore(ctx)

	if msg.WalletId == "" {
		return nil, fmt.Errorf("walletId cannot be empty")
	}
	if msg.Did == "" {
		return nil, fmt.Errorf("did cannot be empty — wallet must be linked to a DID")
	}
	if msg.DisplayName == "" {
		return nil, fmt.Errorf("displayName cannot be empty")
	}

	// One wallet profile per address
	profileKey := fmt.Sprintf("wallet-profile-%s", msg.Creator)
	existing, _ := store.Get([]byte(profileKey))
	if existing != nil {
		return nil, fmt.Errorf("wallet profile already registered for this address — use update-wallet-profile")
	}

	// Check max wallets per DID
	params := k.GetParams(ctx)
	maxPerDid := params.MaxWalletsPerDid
	if maxPerDid == 0 { maxPerDid = 5 }
	didCountKey := fmt.Sprintf("wallet-did-count-%s", msg.Did)
	countBytes, _ := store.Get([]byte(didCountKey))
	count := uint64(0)
	if countBytes != nil {
		fmt.Sscanf(string(countBytes), "%d", &count)
	}
	if count >= maxPerDid {
		return nil, fmt.Errorf("DID %s already has max wallets (%d) — cannot register more", msg.Did, maxPerDid)
	}

	store.Set([]byte(profileKey),
		[]byte(fmt.Sprintf("%s|%s|%s|%s|%d|active",
			msg.WalletId, msg.Did, msg.DisplayName,
			msg.AvatarHash, ctx.BlockHeight())))
	store.Set([]byte(didCountKey), []byte(fmt.Sprintf("%d", count+1)))

	ctx.EventManager().EmitEvent(sdk.NewEvent("wallet_profile_registered",
		sdk.NewAttribute("wallet_id", msg.WalletId),
		sdk.NewAttribute("did", msg.Did),
		sdk.NewAttribute("display_name", msg.DisplayName),
		sdk.NewAttribute("owner", msg.Creator),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgRegisterWalletProfileResponse{WalletId: msg.WalletId, Registered: true}, nil
}
