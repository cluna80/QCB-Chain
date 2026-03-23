package keeper

import (
	"context"
	"fmt"
	"oan/x/oannode/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ClaimNodeReward(goCtx context.Context, msg *types.MsgClaimNodeReward) (*types.MsgClaimNodeRewardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := k.storeService.OpenKVStore(ctx)

	ownerKey := fmt.Sprintf("nodeid-%s", msg.NodeId)
	ownerBytes, _ := store.Get([]byte(ownerKey))
	if ownerBytes == nil {
		return nil, fmt.Errorf("node %s not found", msg.NodeId)
	}
	if string(ownerBytes) != msg.Creator {
		return nil, fmt.Errorf("only the node operator can claim rewards")
	}

	// SECURITY — epoch cooldown written FIRST
	params := k.GetParams(ctx)
	epochLength := int64(params.EpochLength)
	if epochLength == 0 {
		epochLength = 1000
	}
	lastClaimKey := fmt.Sprintf("node-reward-claimed-%s", msg.NodeId)
	lastClaimBytes, _ := store.Get([]byte(lastClaimKey))
	if lastClaimBytes != nil {
		lastClaim := int64(0)
		fmt.Sscanf(string(lastClaimBytes), "%d", &lastClaim)
		if ctx.BlockHeight()-lastClaim < epochLength {
			blocksLeft := epochLength - (ctx.BlockHeight() - lastClaim)
			return nil, fmt.Errorf("reward already claimed this epoch — %d blocks remaining", blocksLeft)
		}
	}

	// SECURITY — must have recent heartbeat
	heartbeatKey := fmt.Sprintf("node-heartbeat-%s", msg.NodeId)
	heartbeat, _ := store.Get([]byte(heartbeatKey))
	if heartbeat == nil {
		return nil, fmt.Errorf("no uptime proof recorded — submit update-node heartbeat first")
	}
	lastBeat := int64(0)
	fmt.Sscanf(string(heartbeat), "%d", &lastBeat)
	if ctx.BlockHeight()-lastBeat > 500 {
		return nil, fmt.Errorf("node offline too long — submit heartbeat first")
	}

	reward := uint64(100)

	// SECURITY — write cooldown BEFORE transfer
	store.Set([]byte(lastClaimKey), []byte(fmt.Sprintf("%d", ctx.BlockHeight())))

	// REAL TOKEN TRANSFER
	operator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, fmt.Errorf("invalid operator address: %s", err)
	}
	moduleAddr := sdk.AccAddress([]byte(types.ModuleName))
	balance := k.bankKeeper.GetBalance(ctx, moduleAddr, "uoan")
	if !balance.Amount.IsZero() {
		coins := sdk.NewCoins(sdk.NewInt64Coin("uoan", int64(reward)))
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, operator, coins); err != nil {
			store.Delete([]byte(lastClaimKey))
			return nil, fmt.Errorf("reward transfer failed: %s", err)
		}
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent("node_reward_claimed",
		sdk.NewAttribute("node_id", msg.NodeId),
		sdk.NewAttribute("epoch", fmt.Sprintf("%d", msg.Epoch)),
		sdk.NewAttribute("reward", fmt.Sprintf("%d", reward)),
		sdk.NewAttribute("denom", "uoan"),
		sdk.NewAttribute("operator", msg.Creator),
	))
	return &types.MsgClaimNodeRewardResponse{
		NodeId: msg.NodeId, Reward: reward, Epoch: msg.Epoch,
	}, nil
}
