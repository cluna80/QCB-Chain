package keeper

import (
	"context"
	"fmt"
	"qcb/x/qcbmedia/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) LicenseContent(goCtx context.Context, msg *types.MsgLicenseContent) (*types.MsgLicenseContentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	validTypes := map[string]bool{"exclusive": true, "non-exclusive": true, "limited": true, "perpetual": true}
	if !validTypes[msg.LicenseType] {
		return nil, fmt.Errorf("licenseType must be exclusive, non-exclusive, limited, or perpetual")
	}
	if msg.Fee == 0 {
		return nil, fmt.Errorf("license fee must be greater than 0")
	}
	if msg.Duration <= 0 && msg.LicenseType != "perpetual" {
		return nil, fmt.Errorf("duration must be positive for non-perpetual licenses")
	}
	store := k.storeService.OpenKVStore(ctx)
	nftKey := fmt.Sprintf("nft-%s", msg.NftId)
	nftData, _ := store.Get([]byte(nftKey))
	if nftData == nil {
		return nil, fmt.Errorf("NFT %s not found", msg.NftId)
	}
	licenseId := fmt.Sprintf("license-%d-%s", ctx.BlockHeight(), msg.NftId)
	now := int32(ctx.BlockTime().Unix())
	expiresAt := now + msg.Duration
	if msg.LicenseType == "perpetual" {
		expiresAt = 0
	}
	store.Set([]byte(fmt.Sprintf("license-%s", licenseId)),
		[]byte(fmt.Sprintf("%s|%s|%s|%d|%d", msg.NftId, msg.LicenseeAddr, msg.LicenseType, msg.Fee, expiresAt)))
	ctx.EventManager().EmitEvent(sdk.NewEvent("content_licensed",
		sdk.NewAttribute("license_id", licenseId),
		sdk.NewAttribute("nft_id", msg.NftId),
		sdk.NewAttribute("licensee", msg.LicenseeAddr),
		sdk.NewAttribute("type", msg.LicenseType),
		sdk.NewAttribute("fee", fmt.Sprintf("%d", msg.Fee)),
	))
	return &types.MsgLicenseContentResponse{LicenseId: licenseId, ExpiresAt: expiresAt}, nil
}
