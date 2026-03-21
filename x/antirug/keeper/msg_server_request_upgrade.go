package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"oan/x/antirug/types"
)

func (k msgServer) RequestUpgrade(goCtx context.Context, msg *types.MsgRequestUpgrade) (*types.MsgRequestUpgradeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := k.storeService.OpenKVStore(ctx)
	params := k.GetParams(ctx)

	// DORMANCY CHECK
	if !params.Enabled {
		return nil, fmt.Errorf("antirug module not yet active — upgrade requests optional until governance activates it")
	}

	ownerBytes, _ := store.Get([]byte(fmt.Sprintf("antirug-owner-%s", msg.TokenId)))
	if ownerBytes == nil {
		return nil, fmt.Errorf("token %s not registered in antirug system", msg.TokenId)
	}
	if string(ownerBytes) != msg.Creator {
		return nil, fmt.Errorf("only the token owner can request upgrades for %s", msg.TokenId)
	}

	// Check not frozen
	frozenKey := fmt.Sprintf("antirug-frozen-%s", msg.TokenId)
	frozen, _ := store.Get([]byte(frozenKey))
	if frozen != nil {
		return nil, fmt.Errorf("token %s is frozen — cannot request upgrades", msg.TokenId)
	}

	timelockBlocks := int64(params.TimelockBlocks)
	if timelockBlocks == 0 { timelockBlocks = 14400 } // 24h default
	timelockEndsAt := int32(ctx.BlockHeight()) + int32(timelockBlocks)

	requestId := fmt.Sprintf("upgrade-%d-%s", ctx.BlockHeight(), msg.TokenId)
	store.Set([]byte(fmt.Sprintf("antirug-upgrade-%s", requestId)),
		[]byte(fmt.Sprintf("%s|%s|%s|%d|pending",
			msg.TokenId, msg.UpgradeType, msg.Description, timelockEndsAt)))

	ctx.EventManager().EmitEvent(sdk.NewEvent("upgrade_requested",
		sdk.NewAttribute("request_id", requestId),
		sdk.NewAttribute("token_id", msg.TokenId),
		sdk.NewAttribute("upgrade_type", msg.UpgradeType),
		sdk.NewAttribute("timelock_ends_at", fmt.Sprintf("%d", timelockEndsAt)),
		sdk.NewAttribute("requester", msg.Creator),
	))
	return &types.MsgRequestUpgradeResponse{RequestId: requestId, TimelockEndsAt: timelockEndsAt}, nil
}
