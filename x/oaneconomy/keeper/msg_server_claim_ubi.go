package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"oan/x/oaneconomy/types"
)

func (k msgServer) ClaimUbi(goCtx context.Context, msg *types.MsgClaimUbi) (*types.MsgClaimUbiResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	epochBlocks := int64(100)
	store := k.storeService.OpenKVStore(ctx)
	lastClaimKey := fmt.Sprintf("ubi-last-claim-%s", msg.Creator)
	lastClaimBytes, _ := store.Get([]byte(lastClaimKey))
	if lastClaimBytes != nil {
		lastClaim := int64(0)
		for i, b := range lastClaimBytes {
			lastClaim |= int64(b) << (8 * i)
		}
		if ctx.BlockHeight()-lastClaim < epochBlocks {
			blocksLeft := epochBlocks - (ctx.BlockHeight() - lastClaim)
			return nil, fmt.Errorf("UBI already claimed this epoch — %d blocks remaining", blocksLeft)
		}
	}
	heightBytes := make([]byte, 8)
	h := ctx.BlockHeight()
	for i := 0; i < 8; i++ {
		heightBytes[i] = byte(h >> (8 * i))
	}
	store.Set([]byte(lastClaimKey), heightBytes)
	ubiAmount := uint64(1000)
	claimedAt := int32(ctx.BlockTime().Unix())
	ctx.EventManager().EmitEvent(sdk.NewEvent("ubi_claimed",
		sdk.NewAttribute("claimer", msg.Creator),
		sdk.NewAttribute("amount", fmt.Sprintf("%d", ubiAmount)),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgClaimUbiResponse{Amount: ubiAmount, ClaimedAt: claimedAt}, nil
}
