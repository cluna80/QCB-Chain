package keeper

import (
	"context"
	"fmt"
	"oan/x/oansports/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ClaimPredictionReward(goCtx context.Context, msg *types.MsgClaimPredictionReward) (*types.MsgClaimPredictionRewardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := k.storeService.OpenKVStore(ctx)
	predKey := fmt.Sprintf("preddata-%s", msg.PredictionId)
	predData, _ := store.Get([]byte(predKey))
	if predData == nil {
		return nil, fmt.Errorf("prediction %s not found", msg.PredictionId)
	}
	claimedKey := fmt.Sprintf("pred-claimed-%s", msg.PredictionId)
	alreadyClaimed, _ := store.Get([]byte(claimedKey))
	if alreadyClaimed != nil {
		return nil, fmt.Errorf("prediction %s reward already claimed", msg.PredictionId)
	}
	store.Set([]byte(claimedKey), []byte("1"))
	reward := uint64(1000)
	ctx.EventManager().EmitEvent(sdk.NewEvent("prediction_reward_claimed",
		sdk.NewAttribute("prediction_id", msg.PredictionId),
		sdk.NewAttribute("claimer", msg.Creator),
		sdk.NewAttribute("reward", fmt.Sprintf("%d", reward)),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgClaimPredictionRewardResponse{Success: true, Reward: reward}, nil
}
