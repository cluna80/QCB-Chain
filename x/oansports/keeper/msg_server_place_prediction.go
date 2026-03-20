package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"oan/x/oansports/types"
)

func (k msgServer) PlacePrediction(goCtx context.Context, msg *types.MsgPlacePrediction) (*types.MsgPlacePredictionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if msg.Stake == 0 {
		return nil, fmt.Errorf("stake must be greater than 0")
	}
	if msg.Prediction == "" {
		return nil, fmt.Errorf("prediction cannot be empty")
	}
	store := k.storeService.OpenKVStore(ctx)
	matchKey := fmt.Sprintf("match-%s", msg.MatchId)
	matchData, _ := store.Get([]byte(matchKey))
	if matchData == nil {
		return nil, fmt.Errorf("match %s not found", msg.MatchId)
	}
	dupKey := fmt.Sprintf("prediction-%s-%s", msg.MatchId, msg.Creator)
	existing, _ := store.Get([]byte(dupKey))
	if existing != nil {
		return nil, fmt.Errorf("you already placed a prediction on match %s", msg.MatchId)
	}
	predictionId := fmt.Sprintf("pred-%d-%s", ctx.BlockHeight(), msg.Creator[:8])
	odds := uint64(200)
	store.Set([]byte(dupKey), []byte(predictionId))
	store.Set([]byte(fmt.Sprintf("preddata-%s", predictionId)),
		[]byte(fmt.Sprintf("%s|%s|%s|%d|%d|pending", predictionId, msg.MatchId, msg.Prediction, msg.Stake, odds)))
	ctx.EventManager().EmitEvent(sdk.NewEvent("prediction_placed",
		sdk.NewAttribute("prediction_id", predictionId),
		sdk.NewAttribute("match_id", msg.MatchId),
		sdk.NewAttribute("prediction", msg.Prediction),
		sdk.NewAttribute("stake", fmt.Sprintf("%d", msg.Stake)),
		sdk.NewAttribute("predictor", msg.Creator),
	))
	return &types.MsgPlacePredictionResponse{PredictionId: predictionId, Stake: msg.Stake, Odds: odds}, nil
}
