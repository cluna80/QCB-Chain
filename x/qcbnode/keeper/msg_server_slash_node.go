package keeper

import (
	"context"
	"fmt"
	"qcb/x/qcbnode/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SlashNode(goCtx context.Context, msg *types.MsgSlashNode) (*types.MsgSlashNodeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := k.storeService.OpenKVStore(ctx)

	ownerBytes, _ := store.Get([]byte(fmt.Sprintf("nodeid-%s", msg.NodeId)))
	if ownerBytes == nil {
		return nil, fmt.Errorf("node %s not found", msg.NodeId)
	}
	if msg.Evidence == "" {
		return nil, fmt.Errorf("evidence required to slash a node")
	}

	// Slash rates by violation type
	slashRates := map[string]uint64{
		"offline":       1,  // 1% for going offline
		"double-sign":   5,  // 5% for double signing — hard slash
		"bad-inference": 2,  // 2% for wrong AI output
		"bad-vote":      1,  // 0.5% for voting on losing fork — rounded to 1
		"malicious":     10, // 10% for confirmed malicious behavior
	}
	slashPct, valid := slashRates[msg.SlashType]
	if !valid {
		return nil, fmt.Errorf("slashType must be offline, double-sign, bad-inference, bad-vote, or malicious")
	}

	// Get current stake
	stakeKey := fmt.Sprintf("staked-amount-%s", string(ownerBytes))
	stakeBytes, _ := store.Get([]byte(stakeKey))
	stakedAmount := uint64(0)
	if stakeBytes != nil {
		fmt.Sscanf(string(stakeBytes), "%d", &stakedAmount)
	}

	slashed := stakedAmount * slashPct / 100
	newStake := stakedAmount - slashed
	store.Set([]byte(stakeKey), []byte(fmt.Sprintf("%d", newStake)))

	// Mark node as slashed
	store.Set([]byte(fmt.Sprintf("node-slashed-%s", msg.NodeId)),
		[]byte(fmt.Sprintf("%s|%d|%d", msg.SlashType, slashed, ctx.BlockHeight())))

	// If double-sign — immediately jail the node
	if msg.SlashType == "double-sign" || msg.SlashType == "malicious" {
		store.Set([]byte(fmt.Sprintf("node-jailed-%s", msg.NodeId)), []byte("1"))
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent("node_slashed",
		sdk.NewAttribute("node_id", msg.NodeId),
		sdk.NewAttribute("slash_type", msg.SlashType),
		sdk.NewAttribute("slashed_amount", fmt.Sprintf("%d", slashed)),
		sdk.NewAttribute("remaining_stake", fmt.Sprintf("%d", newStake)),
		sdk.NewAttribute("slashed_by", msg.Creator),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgSlashNodeResponse{NodeId: msg.NodeId, Slashed: slashed, Reason: msg.SlashType}, nil
}
