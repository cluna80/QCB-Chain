package keeper

import (
	"context"
	"fmt"
	"qcb/x/qcbprotocol/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SetAddressTier(goCtx context.Context, msg *types.MsgSetAddressTier) (*types.MsgSetAddressTierResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if msg.Creator != k.GetAuthority() {
		return nil, fmt.Errorf("unauthorized: only governance can set address tiers")
	}
	validTiers := map[string]bool{"unverified": true, "verified": true, "dao": true, "exempt": true}
	if !validTiers[msg.Tier] {
		return nil, fmt.Errorf("invalid tier: must be unverified, verified, dao, or exempt")
	}
	k.SetAddressTierStore(ctx, msg.Address, msg.Tier)
	if msg.Tier == "exempt" {
		k.Keeper.ExemptAddress(ctx, msg.Address)
	}
	ctx.EventManager().EmitEvent(sdk.NewEvent("tier_set",
		sdk.NewAttribute("address", msg.Address),
		sdk.NewAttribute("tier", msg.Tier),
		sdk.NewAttribute("set_by", msg.Creator),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgSetAddressTierResponse{}, nil
}
