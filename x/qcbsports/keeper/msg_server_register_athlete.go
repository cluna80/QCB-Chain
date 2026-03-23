package keeper

import (
	"context"
	"fmt"
	"qcb/x/qcbsports/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RegisterAthlete(goCtx context.Context, msg *types.MsgRegisterAthlete) (*types.MsgRegisterAthleteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if msg.Sport == "" {
		return nil, fmt.Errorf("sport cannot be empty")
	}
	if msg.AgentId == "" {
		return nil, fmt.Errorf("agentId cannot be empty — athlete must be linked to an OAN agent")
	}
	store := k.storeService.OpenKVStore(ctx)
	athleteKey := fmt.Sprintf("athlete-%s", msg.AthleteId)
	existing, _ := store.Get([]byte(athleteKey))
	if existing != nil {
		return nil, fmt.Errorf("athlete %s already registered", msg.AthleteId)
	}
	athleteData := fmt.Sprintf("%s|%s|%s|%s|%s|0|0|0|active",
		msg.AthleteId, msg.AgentId, msg.Sport, msg.Position, msg.Creator)
	store.Set([]byte(athleteKey), []byte(athleteData))
	ctx.EventManager().EmitEvent(sdk.NewEvent("athlete_registered",
		sdk.NewAttribute("athlete_id", msg.AthleteId),
		sdk.NewAttribute("agent_id", msg.AgentId),
		sdk.NewAttribute("sport", msg.Sport),
		sdk.NewAttribute("position", msg.Position),
		sdk.NewAttribute("owner", msg.Creator),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgRegisterAthleteResponse{AthleteId: msg.AthleteId, Registered: true}, nil
}
