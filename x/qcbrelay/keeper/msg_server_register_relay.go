package keeper

import (
	"context"
	"fmt"
	"qcb/x/qcbrelay/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RegisterRelay(goCtx context.Context, msg *types.MsgRegisterRelay) (*types.MsgRegisterRelayResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := k.storeService.OpenKVStore(ctx)

	validRegions := map[string]bool{
		"us-east": true, "us-west": true, "eu-west": true, "eu-east": true,
		"asia-pacific": true, "africa": true, "latam": true, "global": true,
	}
	if !validRegions[msg.Region] {
		return nil, fmt.Errorf("region must be us-east, us-west, eu-west, eu-east, asia-pacific, africa, latam, or global")
	}
	if msg.Endpoint == "" {
		return nil, fmt.Errorf("endpoint cannot be empty")
	}
	if len(msg.PubKeyHash) < 32 {
		return nil, fmt.Errorf("pubKeyHash must be at least 32 characters")
	}

	// Check not already registered
	relayKey := fmt.Sprintf("relay-%s", msg.RelayId)
	existing, _ := store.Get([]byte(relayKey))
	if existing != nil {
		return nil, fmt.Errorf("relay %s already registered", msg.RelayId)
	}

	// Check region cap
	params := k.GetParams(ctx)
	maxPerRegion := params.MaxRelaysPerRegion
	if maxPerRegion == 0 {
		maxPerRegion = 100
	}
	regionCountKey := fmt.Sprintf("relay-region-count-%s", msg.Region)
	countBytes, _ := store.Get([]byte(regionCountKey))
	count := uint64(0)
	if countBytes != nil {
		fmt.Sscanf(string(countBytes), "%d", &count)
	}
	if count >= maxPerRegion {
		return nil, fmt.Errorf("region %s is at capacity (%d relays) — try another region", msg.Region, maxPerRegion)
	}

	store.Set([]byte(relayKey),
		[]byte(fmt.Sprintf("%s|%s|%s|%s|%s|0|active",
			msg.RelayId, msg.Endpoint, msg.Region,
			msg.PubKeyHash, msg.Creator)))
	store.Set([]byte(fmt.Sprintf("relay-owner-%s", msg.RelayId)), []byte(msg.Creator))
	store.Set([]byte(regionCountKey), []byte(fmt.Sprintf("%d", count+1)))

	ctx.EventManager().EmitEvent(sdk.NewEvent("relay_registered",
		sdk.NewAttribute("relay_id", msg.RelayId),
		sdk.NewAttribute("region", msg.Region),
		sdk.NewAttribute("endpoint", msg.Endpoint),
		sdk.NewAttribute("operator", msg.Creator),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgRegisterRelayResponse{RelayId: msg.RelayId, Registered: true}, nil
}
