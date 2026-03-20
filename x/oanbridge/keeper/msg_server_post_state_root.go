package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"oan/x/oanbridge/types"
)

func (k msgServer) PostStateRoot(goCtx context.Context, msg *types.MsgPostStateRoot) (*types.MsgPostStateRootResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if msg.StateRoot == "" {
		return nil, fmt.Errorf("stateRoot cannot be empty")
	}
	if msg.Proof == "" {
		return nil, fmt.Errorf("proof cannot be empty — must include merkle proof")
	}
	validAnchors := map[string]bool{"bitcoin": true, "ethereum": true, "cosmos": true}
	if !validAnchors[msg.AnchorChain] {
		return nil, fmt.Errorf("anchorChain must be bitcoin, ethereum, or cosmos")
	}
	store := k.storeService.OpenKVStore(ctx)
	anchorId := fmt.Sprintf("anchor-%d", ctx.BlockHeight())
	anchorData := fmt.Sprintf("%s|%s|%d|%s|%d",
		anchorId, msg.StateRoot, msg.BlockHeight, msg.AnchorChain, ctx.BlockHeight())
	store.Set([]byte(fmt.Sprintf("state-root-%d", ctx.BlockHeight())), []byte(anchorData))
	store.Set([]byte("latest-state-root"), []byte(msg.StateRoot))
	ctx.EventManager().EmitEvent(sdk.NewEvent("state_root_posted",
		sdk.NewAttribute("anchor_id", anchorId),
		sdk.NewAttribute("state_root", msg.StateRoot),
		sdk.NewAttribute("anchor_chain", msg.AnchorChain),
		sdk.NewAttribute("block_height", fmt.Sprintf("%d", msg.BlockHeight)),
		sdk.NewAttribute("posted_by", msg.Creator),
		sdk.NewAttribute("oan_block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgPostStateRootResponse{AnchorId: anchorId, StateRoot: msg.StateRoot, Anchored: true}, nil
}
