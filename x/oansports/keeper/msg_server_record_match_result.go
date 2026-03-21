package keeper

import (
	"context"
	"fmt"
	"oan/x/oansports/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RecordMatchResult(goCtx context.Context, msg *types.MsgRecordMatchResult) (*types.MsgRecordMatchResultResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := k.storeService.OpenKVStore(ctx)
	matchKey := fmt.Sprintf("match-%s", msg.MatchId)
	matchData, _ := store.Get([]byte(matchKey))
	if matchData == nil {
		return nil, fmt.Errorf("match %s not found", msg.MatchId)
	}
	if msg.Winner == msg.Loser {
		return nil, fmt.Errorf("winner and loser cannot be the same athlete")
	}
	if msg.StatsHash == "" {
		return nil, fmt.Errorf("statsHash cannot be empty — must include full match stats")
	}
	resultData := fmt.Sprintf("%s|%s|%s|%d|%d|%s|completed",
		msg.MatchId, msg.Winner, msg.Loser, msg.ScoreA, msg.ScoreB, msg.StatsHash)
	store.Set([]byte(matchKey), []byte(resultData))
	winnerKey := fmt.Sprintf("athlete-wins-%s", msg.Winner)
	winsBytes, _ := store.Get([]byte(winnerKey))
	wins := uint64(0)
	if winsBytes != nil {
		fmt.Sscanf(string(winsBytes), "%d", &wins)
	}
	wins++
	store.Set([]byte(winnerKey), []byte(fmt.Sprintf("%d", wins)))
	ctx.EventManager().EmitEvent(sdk.NewEvent("match_result_recorded",
		sdk.NewAttribute("match_id", msg.MatchId),
		sdk.NewAttribute("winner", msg.Winner),
		sdk.NewAttribute("loser", msg.Loser),
		sdk.NewAttribute("score_a", fmt.Sprintf("%d", msg.ScoreA)),
		sdk.NewAttribute("score_b", fmt.Sprintf("%d", msg.ScoreB)),
		sdk.NewAttribute("stats_hash", msg.StatsHash),
	))
	return &types.MsgRecordMatchResultResponse{MatchId: msg.MatchId, Winner: msg.Winner, Recorded: true}, nil
}
