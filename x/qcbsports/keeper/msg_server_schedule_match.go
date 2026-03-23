package keeper

import (
	"context"
	"fmt"
	"qcb/x/qcbsports/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ScheduleMatch(goCtx context.Context, msg *types.MsgScheduleMatch) (*types.MsgScheduleMatchResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if msg.AthleteA == msg.AthleteB {
		return nil, fmt.Errorf("an athlete cannot compete against themselves")
	}
	if msg.ScheduledAt <= int32(ctx.BlockTime().Unix()) {
		return nil, fmt.Errorf("match must be scheduled in the future")
	}
	store := k.storeService.OpenKVStore(ctx)
	matchKey := fmt.Sprintf("match-%s", msg.MatchId)
	existing, _ := store.Get([]byte(matchKey))
	if existing != nil {
		return nil, fmt.Errorf("match %s already scheduled", msg.MatchId)
	}
	matchData := fmt.Sprintf("%s|%s|%s|%s|%d|pending",
		msg.MatchId, msg.AthleteA, msg.AthleteB, msg.StadiumId, msg.ScheduledAt)
	store.Set([]byte(matchKey), []byte(matchData))
	ctx.EventManager().EmitEvent(sdk.NewEvent("match_scheduled",
		sdk.NewAttribute("match_id", msg.MatchId),
		sdk.NewAttribute("athlete_a", msg.AthleteA),
		sdk.NewAttribute("athlete_b", msg.AthleteB),
		sdk.NewAttribute("stadium_id", msg.StadiumId),
		sdk.NewAttribute("scheduled_at", fmt.Sprintf("%d", msg.ScheduledAt)),
	))
	return &types.MsgScheduleMatchResponse{MatchId: msg.MatchId, ScheduledAt: msg.ScheduledAt}, nil
}
