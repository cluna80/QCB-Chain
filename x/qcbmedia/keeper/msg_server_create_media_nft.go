package keeper

import (
	"context"
	"fmt"
	"qcb/x/qcbmedia/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateMediaNft(goCtx context.Context, msg *types.MsgCreateMediaNft) (*types.MsgCreateMediaNftResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	validTypes := map[string]bool{"film": true, "music": true, "art": true, "research": true, "clip": true}
	if !validTypes[msg.MediaType] {
		return nil, fmt.Errorf("mediaType must be film, music, art, research, or clip")
	}
	if msg.CreatorShare > 100 {
		return nil, fmt.Errorf("creatorShare cannot exceed 100 percent")
	}
	if msg.ContentHash == "" {
		return nil, fmt.Errorf("contentHash cannot be empty")
	}
	params := k.GetParams(ctx)
	royaltyRate := params.RoyaltyRate
	if royaltyRate == 0 {
		royaltyRate = 10
	}
	nftId := fmt.Sprintf("media-%d-%s", ctx.BlockHeight(), msg.Creator[:8])
	store := k.storeService.OpenKVStore(ctx)
	nftKey := fmt.Sprintf("nft-%s", nftId)
	nftData := fmt.Sprintf("%s|%s|%s|%d|%d|%s|active",
		nftId, msg.Title, msg.MediaType, msg.CreatorShare, royaltyRate, msg.Creator)
	store.Set([]byte(nftKey), []byte(nftData))
	store.Set([]byte(fmt.Sprintf("nft-owner-%s", nftId)), []byte(msg.Creator))
	ctx.EventManager().EmitEvent(sdk.NewEvent("media_nft_created",
		sdk.NewAttribute("nft_id", nftId),
		sdk.NewAttribute("title", msg.Title),
		sdk.NewAttribute("media_type", msg.MediaType),
		sdk.NewAttribute("content_hash", msg.ContentHash),
		sdk.NewAttribute("creator", msg.Creator),
		sdk.NewAttribute("creator_share", fmt.Sprintf("%d", msg.CreatorShare)),
		sdk.NewAttribute("royalty_rate", fmt.Sprintf("%d", royaltyRate)),
	))
	return &types.MsgCreateMediaNftResponse{NftId: nftId, ContentHash: msg.ContentHash, RoyaltyRate: royaltyRate}, nil
}
