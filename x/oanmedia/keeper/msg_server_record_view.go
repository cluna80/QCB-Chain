package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"oan/x/oanmedia/types"
)

func (k msgServer) RecordView(goCtx context.Context, msg *types.MsgRecordView) (*types.MsgRecordViewResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if msg.PaymentAmount == 0 {
		return nil, fmt.Errorf("payment amount must be greater than 0")
	}
	store := k.storeService.OpenKVStore(ctx)
	nftKey := fmt.Sprintf("nft-%s", msg.NftId)
	nftData, _ := store.Get([]byte(nftKey))
	if nftData == nil {
		return nil, fmt.Errorf("NFT %s not found", msg.NftId)
	}
	params := k.GetParams(ctx)
	royaltyRate := params.RoyaltyRate
	if royaltyRate == 0 { royaltyRate = 10 }
	creatorEarned := msg.PaymentAmount * royaltyRate / 100
	viewKey := fmt.Sprintf("views-%s", msg.NftId)
	viewBytes, _ := store.Get([]byte(viewKey))
	views := uint64(0)
	if viewBytes != nil {
		fmt.Sscanf(string(viewBytes), "%d", &views)
	}
	views++
	store.Set([]byte(viewKey), []byte(fmt.Sprintf("%d", views)))
	ctx.EventManager().EmitEvent(sdk.NewEvent("content_viewed",
		sdk.NewAttribute("nft_id", msg.NftId),
		sdk.NewAttribute("viewer", msg.ViewerAddr),
		sdk.NewAttribute("payment", fmt.Sprintf("%d", msg.PaymentAmount)),
		sdk.NewAttribute("creator_earned", fmt.Sprintf("%d", creatorEarned)),
		sdk.NewAttribute("total_views", fmt.Sprintf("%d", views)),
	))
	return &types.MsgRecordViewResponse{Success: true, CreatorEarned: creatorEarned}, nil
}
