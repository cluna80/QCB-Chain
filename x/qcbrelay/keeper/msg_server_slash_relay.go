package keeper

import (
	"context"
	"fmt"
	"qcb/x/qcbrelay/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SlashRelay(goCtx context.Context, msg *types.MsgSlashRelay) (*types.MsgSlashRelayResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := k.storeService.OpenKVStore(ctx)

	if msg.Evidence == "" {
		return nil, fmt.Errorf("evidence required to slash a relay")
	}
	validSlashTypes := map[string]bool{
		"offline": true, "censorship": true, "tampering": true, "malicious": true,
	}
	if !validSlashTypes[msg.SlashType] {
		return nil, fmt.Errorf("slashType must be offline, censorship, tampering, or malicious")
	}

	ownerBytes, _ := store.Get([]byte(fmt.Sprintf("relay-owner-%s", msg.RelayId)))
	if ownerBytes == nil {
		return nil, fmt.Errorf("relay %s not found", msg.RelayId)
	}

	slashRates := map[string]uint64{
		"offline": 5, "censorship": 15, "tampering": 25, "malicious": 50,
	}
	slashed := slashRates[msg.SlashType]

	// Reduce score
	scoreKey := fmt.Sprintf("relay-score-%s", msg.RelayId)
	scoreBytes, _ := store.Get([]byte(scoreKey))
	score := uint64(100)
	if scoreBytes != nil {
		fmt.Sscanf(string(scoreBytes), "%d", &score)
	}
	if score < slashed {
		score = 0
	} else {
		score -= slashed
	}
	store.Set([]byte(scoreKey), []byte(fmt.Sprintf("%d", score)))

	if msg.SlashType == "malicious" || msg.SlashType == "tampering" {
		store.Set([]byte(fmt.Sprintf("relay-jailed-%s", msg.RelayId)), []byte("1"))
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent("relay_slashed",
		sdk.NewAttribute("relay_id", msg.RelayId),
		sdk.NewAttribute("slash_type", msg.SlashType),
		sdk.NewAttribute("score_penalty", fmt.Sprintf("%d", slashed)),
		sdk.NewAttribute("new_score", fmt.Sprintf("%d", score)),
		sdk.NewAttribute("slashed_by", msg.Creator),
	))
	return &types.MsgSlashRelayResponse{RelayId: msg.RelayId, Slashed: slashed}, nil
}
