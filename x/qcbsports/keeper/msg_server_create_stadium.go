package keeper

import (
	"context"
	"fmt"
	"qcb/x/qcbsports/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateStadium(goCtx context.Context, msg *types.MsgCreateStadium) (*types.MsgCreateStadiumResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if msg.Capacity == 0 {
		return nil, fmt.Errorf("capacity must be greater than 0")
	}
	if msg.Capacity > 1000000 {
		return nil, fmt.Errorf("capacity cannot exceed 1,000,000")
	}
	store := k.storeService.OpenKVStore(ctx)
	stadiumKey := fmt.Sprintf("stadium-%s", msg.StadiumId)
	existing, _ := store.Get([]byte(stadiumKey))
	if existing != nil {
		return nil, fmt.Errorf("stadium %s already exists", msg.StadiumId)
	}
	stadiumData := fmt.Sprintf("%s|%s|%d|%s|%s|0|active",
		msg.StadiumId, msg.Name, msg.Capacity, msg.Location, msg.Creator)
	store.Set([]byte(stadiumKey), []byte(stadiumData))
	ctx.EventManager().EmitEvent(sdk.NewEvent("stadium_created",
		sdk.NewAttribute("stadium_id", msg.StadiumId),
		sdk.NewAttribute("name", msg.Name),
		sdk.NewAttribute("capacity", fmt.Sprintf("%d", msg.Capacity)),
		sdk.NewAttribute("location", msg.Location),
		sdk.NewAttribute("owner", msg.Creator),
	))
	return &types.MsgCreateStadiumResponse{StadiumId: msg.StadiumId, Capacity: msg.Capacity}, nil
}
