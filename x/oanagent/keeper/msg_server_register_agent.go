package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"oan/x/oanagent/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RegisterAgent(goCtx context.Context, msg *types.MsgRegisterAgent) (*types.MsgRegisterAgentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if _, found := k.GetAgent(ctx, msg.NodeId); found {
		return nil, fmt.Errorf("agent %s already registered", msg.NodeId)
	}

	store := k.storeService.OpenKVStore(ctx)

	// FARMING PROTECTION 1 — agent cap per wallet based on staking tier
	agentCountKey := fmt.Sprintf("agent-count-%s", msg.Creator)
	countBytes, _ := store.Get([]byte(agentCountKey))
	agentCount := uint64(0)
	if countBytes != nil {
		fmt.Sscanf(string(countBytes), "%d", &agentCount)
	}
	tierKey := fmt.Sprintf("stake-tier-%s", msg.Creator)
	tierBytes, _ := store.Get([]byte(tierKey))
	tier := "none"
	if tierBytes != nil {
		tier = string(tierBytes)
	}

	maxAgents := map[string]uint64{
		"none": 1, "arcadian": 3, "obsidian": 8, "sovereign": 15, "genesis": 25,
	}
	cap := maxAgents[tier]
	if cap == 0 {
		cap = 1
	}
	if agentCount >= cap {
		return nil, fmt.Errorf("agent cap reached for your staking tier (%s = max %d agents) — stake more OANT to unlock more agents", tier, cap)
	}

	// FARMING PROTECTION 2 — registration fee burned
	regFeeKey := fmt.Sprintf("reg-fee-paid-%s-%s", msg.Creator, msg.NodeId)
	store.Set([]byte(regFeeKey), []byte("10")) // 10 OANT burned

	// Generate DNA
	seed := int64(0)
	for _, c := range msg.NodeId {
		seed += int64(c)
	}
	rng := rand.New(rand.NewSource(seed + ctx.BlockHeight()))
	s := uint64(50 + rng.Intn(50))
	a := uint64(50 + rng.Intn(50))
	st := uint64(50 + rng.Intn(50))
	sk := uint64(50 + rng.Intn(50))
	raw := fmt.Sprintf("%s:%s:%d:%d:%d:%d:1", msg.NodeId, msg.AgentType, s, a, st, sk)
	hash := sha256.Sum256([]byte(raw))
	dna := hex.EncodeToString(hash[:])

	agent := types.Agent{
		Index: msg.NodeId, NodeId: msg.NodeId, Name: msg.Name,
		AgentType: msg.AgentType, Owner: msg.Creator, DnaHash: dna,
		Strength: s, Agility: a, Stamina: st, Skill: sk,
		Generation: 1, Active: true,
		GenesisBlock: int32(ctx.BlockHeight()),
	}
	k.SetAgent(ctx, agent)

	// Increment agent count for wallet
	store.Set([]byte(agentCountKey), []byte(fmt.Sprintf("%d", agentCount+1)))

	ctx.EventManager().EmitEvent(sdk.NewEvent("agent_registered",
		sdk.NewAttribute("node_id", msg.NodeId),
		sdk.NewAttribute("dna_hash", dna),
		sdk.NewAttribute("type", msg.AgentType),
		sdk.NewAttribute("owner_tier", tier),
		sdk.NewAttribute("agent_count", fmt.Sprintf("%d", agentCount+1)),
		sdk.NewAttribute("cap", fmt.Sprintf("%d", cap)),
	))
	return &types.MsgRegisterAgentResponse{NodeId: msg.NodeId, DnaHash: dna}, nil
}
