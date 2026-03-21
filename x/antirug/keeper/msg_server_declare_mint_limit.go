package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"oan/x/antirug/types"
)

func (k msgServer) DeclareMintLimit(goCtx context.Context, msg *types.MsgDeclareMintLimit) (*types.MsgDeclareMintLimitResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := k.storeService.OpenKVStore(ctx)
	params := k.GetParams(ctx)

	// DORMANCY CHECK
	if !params.Enabled {
		return nil, fmt.Errorf("antirug module not yet active — mint limits optional until governance activates it")
	}

	if msg.MaxSupply == 0 {
		return nil, fmt.Errorf("maxSupply must be greater than 0 — unlimited minting not permitted")
	}

	maxPerBlock := params.MaxMintPerBlock
	if maxPerBlock == 0 { maxPerBlock = 1000000 }
	if msg.MaxPerBlock > maxPerBlock {
		return nil, fmt.Errorf("maxPerBlock exceeds protocol limit of %d — prevents flash minting attacks", maxPerBlock)
	}

	ownerBytes, _ := store.Get([]byte(fmt.Sprintf("antirug-owner-%s", msg.TokenId)))
	if ownerBytes != nil && string(ownerBytes) != msg.Creator {
		return nil, fmt.Errorf("only the token owner can declare mint limits for %s", msg.TokenId)
	}

	mintKey := fmt.Sprintf("antirug-mint-%s", msg.TokenId)
	store.Set([]byte(mintKey),
		[]byte(fmt.Sprintf("%d|%d|%d", msg.MaxSupply, msg.MaxPerBlock, ctx.BlockHeight())))

	ctx.EventManager().EmitEvent(sdk.NewEvent("mint_limit_declared",
		sdk.NewAttribute("token_id", msg.TokenId),
		sdk.NewAttribute("max_supply", fmt.Sprintf("%d", msg.MaxSupply)),
		sdk.NewAttribute("max_per_block", fmt.Sprintf("%d", msg.MaxPerBlock)),
		sdk.NewAttribute("declared_by", msg.Creator),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgDeclareMintLimitResponse{TokenId: msg.TokenId, LimitSet: true}, nil
}
