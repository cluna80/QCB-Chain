package keeper

import (
	"context"
	"fmt"
	"oan/x/oanmedia/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ClaimRoyalty(goCtx context.Context, msg *types.MsgClaimRoyalty) (*types.MsgClaimRoyaltyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := k.storeService.OpenKVStore(ctx)
	ownerKey := fmt.Sprintf("nft-owner-%s", msg.NftId)
	ownerBytes, _ := store.Get([]byte(ownerKey))
	if ownerBytes == nil {
		return nil, fmt.Errorf("NFT %s not found", msg.NftId)
	}
	if string(ownerBytes) != msg.Creator {
		return nil, fmt.Errorf("only the NFT owner can claim royalties")
	}
	if msg.PeriodEnd <= msg.PeriodStart {
		return nil, fmt.Errorf("periodEnd must be after periodStart")
	}
	cooldownKey := fmt.Sprintf("royalty-claimed-%s-%s", msg.NftId, msg.Creator)
	alreadyClaimed, _ := store.Get([]byte(cooldownKey))
	if alreadyClaimed != nil {
		return nil, fmt.Errorf("royalties already claimed for this period")
	}
	store.Set([]byte(cooldownKey), []byte(fmt.Sprintf("%d", ctx.BlockHeight())))
	royaltyAmount := uint64(500)
	ctx.EventManager().EmitEvent(sdk.NewEvent("royalty_claimed",
		sdk.NewAttribute("nft_id", msg.NftId),
		sdk.NewAttribute("claimer", msg.Creator),
		sdk.NewAttribute("amount", fmt.Sprintf("%d", royaltyAmount)),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgClaimRoyaltyResponse{Amount: royaltyAmount, NftId: msg.NftId}, nil
}
