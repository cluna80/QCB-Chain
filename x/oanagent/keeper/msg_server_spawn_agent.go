package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"oan/x/oanagent/types"
)

func (k msgServer) SpawnAgent(goCtx context.Context, msg *types.MsgSpawnAgent) (*types.MsgSpawnAgentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// ── REQUIREMENT 1: parent must exist and be active ────────────────────────
	parent, found := k.GetAgent(ctx, msg.ParentId)
	if !found {
		return nil, fmt.Errorf("parent agent %s not found", msg.ParentId)
	}
	if !parent.Active {
		return nil, fmt.Errorf("parent agent %s is retired — only active agents can spawn", msg.ParentId)
	}

	// ── REQUIREMENT 2: caller must own the parent ─────────────────────────────
	if parent.Owner != msg.Creator {
		return nil, fmt.Errorf("only the owner of %s can spawn from it", msg.ParentId)
	}

	// ── REQUIREMENT 3: 1000 trade minimum ────────────────────────────────────
	minTrades := int32(1000)
	if parent.TotalTrades < minTrades {
		remaining := minTrades - parent.TotalTrades
		return nil, fmt.Errorf(
			"parent agent needs %d more trades before it can spawn — current: %d, required: %d",
			remaining, parent.TotalTrades, minTrades,
		)
	}

	// ── REQUIREMENT 4: 75% win rate minimum (7500 basis points) ──────────────
	minWinRateBps := int32(7500)
	if parent.WinRateBps < minWinRateBps {
		currentPct := float64(parent.WinRateBps) / 100.0
		return nil, fmt.Errorf(
			"parent agent win rate too low — current: %.2f%%, required: 75.00%% — keep trading to improve",
			currentPct,
		)
	}

	// ── REQUIREMENT 5: child ID must not already exist ────────────────────────
	if _, exists := k.GetAgent(ctx, msg.ChildId); exists {
		return nil, fmt.Errorf("agent %s already exists — choose a different child ID", msg.ChildId)
	}

	// ── REQUIREMENT 6: spawn cooldown — one spawn per 500 blocks ─────────────
	store := k.storeService.OpenKVStore(ctx)
	spawnCooldownKey := fmt.Sprintf("spawn-cooldown-%s", msg.ParentId)
	cooldownBytes, _ := store.Get([]byte(spawnCooldownKey))
	spawnCooldown := int32(500)
	if cooldownBytes != nil {
		lastSpawn := int32(0)
		for i, b := range cooldownBytes {
			lastSpawn |= int32(b) << (8 * i)
		}
		if int32(ctx.BlockHeight())-lastSpawn < spawnCooldown {
			blocksLeft := spawnCooldown - (int32(ctx.BlockHeight()) - lastSpawn)
			return nil, fmt.Errorf(
				"spawn cooldown active for %s — %d blocks remaining (~%d minutes)",
				msg.ParentId, blocksLeft, blocksLeft/10,
			)
		}
	}

	// ── REQUIREMENT 7: agent cap per staking tier ─────────────────────────────
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
	if cap == 0 { cap = 1 }
	if agentCount >= cap {
		return nil, fmt.Errorf(
			"agent cap reached (%s tier = max %d) — stake more OANT to spawn more agents",
			tier, cap,
		)
	}

	// ── SPAWN FEE: 20 OANT burned ─────────────────────────────────────────────
	spawnFeeKey := fmt.Sprintf("spawn-fee-%s-%s", msg.Creator, msg.ChildId)
	store.Set([]byte(spawnFeeKey), []byte("20000000")) // 20 OANT in uoan

	// ── INHERIT COGNITIVE STATE from parent ───────────────────────────────────
	// Child inherits parent's best traits with small mutation
	// Stronger than breed — single parent passes dominant traits
	rng := rand.New(rand.NewSource(ctx.BlockHeight() + int64(parent.TotalTrades)))

	inherit := func(base uint64) uint64 {
		// Spawn inherits more faithfully than breed — 90% of parent value
		// Small mutation +-5% to prevent exact clones
		mutation := int64(base) / 20
		d := rng.Int63n(mutation*2+1) - mutation
		result := int64(base) + d
		if result < 1   { result = 1 }
		if result > 100 { result = 100 }
		return uint64(result)
	}

	// Elite parent bonus — high win rate agents pass stronger traits
	winRateBonus := uint64(0)
	if parent.WinRateBps >= 9000 { winRateBonus = 5 } // 90%+ = +5 to all stats
	if parent.WinRateBps >= 9500 { winRateBonus = 10 } // 95%+ = +10 to all stats

	s  := inherit(parent.Strength + winRateBonus)
	a  := inherit(parent.Agility  + winRateBonus)
	st := inherit(parent.Stamina  + winRateBonus)
	sk := inherit(parent.Skill    + winRateBonus)

	// Cap at 100
	if s  > 100 { s  = 100 }
	if a  > 100 { a  = 100 }
	if st > 100 { st = 100 }
	if sk > 100 { sk = 100 }

	genomeScore := int32((s + a + st + sk) / 4)
	generation  := parent.Generation + 1
	generation32 := int32(generation)

	// ── DNA HASH includes parent cognitive history ────────────────────────────
	// Parent's win rate and trade count baked into child's DNA
	raw := fmt.Sprintf(
		"%s:spawn:%s:%d:%d:%d:%d:%d:%d:%d",
		msg.ChildId, msg.ParentId,
		s, a, st, sk,
		generation,
		parent.WinRateBps,
		parent.TotalTrades,
	)
	hash   := sha256.Sum256([]byte(raw))
	dnaHash := hex.EncodeToString(hash[:])

	// ── CREATE THE CHILD AGENT ────────────────────────────────────────────────
	child := types.Agent{
		Index:        msg.ChildId,
		NodeId:       msg.ChildId,
		Name:         msg.ChildName,
		AgentType:    msg.ChildType,
		Owner:        msg.Creator,
		DnaHash:      dnaHash,
		Strength:     s,
		Agility:      a,
		Stamina:      st,
		Skill:        sk,
		Generation:   generation,
		Active:       true,
		GenesisBlock: int32(ctx.BlockHeight()),
		ParentA:      msg.ParentId,
		ParentB:      "",          // spawn has one parent only
	}
	k.SetAgent(ctx, child)

	// ── UPDATE PARENT — record spawn ──────────────────────────────────────────
	store.Set([]byte(fmt.Sprintf("parent-spawned-%s-%s", msg.ParentId, msg.ChildId)),
		[]byte(fmt.Sprintf("%d", ctx.BlockHeight())))

	// Record cooldown
	heightBytes := make([]byte, 4)
	h := int32(ctx.BlockHeight())
	for i := 0; i < 4; i++ {
		heightBytes[i] = byte(h >> (8 * i))
	}
	store.Set([]byte(spawnCooldownKey), heightBytes)

	// Update agent count for wallet
	store.Set([]byte(agentCountKey), []byte(fmt.Sprintf("%d", agentCount+1)))

	// ── EMIT EVENT ────────────────────────────────────────────────────────────
	ctx.EventManager().EmitEvent(sdk.NewEvent("agent_spawned",
		sdk.NewAttribute("child_id",        msg.ChildId),
		sdk.NewAttribute("child_name",      msg.ChildName),
		sdk.NewAttribute("parent_id",       msg.ParentId),
		sdk.NewAttribute("parent_trades",   fmt.Sprintf("%d", parent.TotalTrades)),
		sdk.NewAttribute("parent_win_rate", fmt.Sprintf("%.2f%%", float64(parent.WinRateBps)/100.0)),
		sdk.NewAttribute("generation",      fmt.Sprintf("%d", generation)),
		sdk.NewAttribute("genome_score",    fmt.Sprintf("%d", genomeScore)),
		sdk.NewAttribute("dna_hash",        dnaHash),
		sdk.NewAttribute("win_rate_bonus",  fmt.Sprintf("%d", winRateBonus)),
		sdk.NewAttribute("spawn_fee_uoan",  "20000000"),
		sdk.NewAttribute("block",           fmt.Sprintf("%d", ctx.BlockHeight())),
	))

	return &types.MsgSpawnAgentResponse{
		ChildId:     msg.ChildId,
		DnaHash:     dnaHash,
		Generation:  generation32,
		GenomeScore: genomeScore,
	}, nil
}
