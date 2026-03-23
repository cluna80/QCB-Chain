package keeper

import (
	"context"
	"fmt"
	"qcb/x/qcbrelay/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RelayHeartbeat(goCtx context.Context, msg *types.MsgRelayHeartbeat) (*types.MsgRelayHeartbeatResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := k.storeService.OpenKVStore(ctx)

	ownerBytes, _ := store.Get([]byte(fmt.Sprintf("relay-owner-%s", msg.RelayId)))
	if ownerBytes == nil {
		return nil, fmt.Errorf("relay %s not found", msg.RelayId)
	}
	if string(ownerBytes) != msg.Creator {
		return nil, fmt.Errorf("only the relay operator can submit heartbeats")
	}
	if msg.ProofHash == "" {
		return nil, fmt.Errorf("proofHash cannot be empty")
	}

	// Update score based on uptime
	scoreKey := fmt.Sprintf("relay-score-%s", msg.RelayId)
	scoreBytes, _ := store.Get([]byte(scoreKey))
	score := uint64(0)
	if scoreBytes != nil {
		fmt.Sscanf(string(scoreBytes), "%d", &score)
	}
	score += 1
	if score > 1000 {
		score = 1000
	}
	store.Set([]byte(scoreKey), []byte(fmt.Sprintf("%d", score)))
	store.Set([]byte(fmt.Sprintf("relay-heartbeat-%s", msg.RelayId)),
		[]byte(fmt.Sprintf("%d|%s", ctx.BlockHeight(), msg.ProofHash)))

	ctx.EventManager().EmitEvent(sdk.NewEvent("relay_heartbeat",
		sdk.NewAttribute("relay_id", msg.RelayId),
		sdk.NewAttribute("score", fmt.Sprintf("%d", score)),
		sdk.NewAttribute("blocks_online", fmt.Sprintf("%d", msg.BlocksOnline)),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgRelayHeartbeatResponse{RelayId: msg.RelayId, Score: score}, nil
}
