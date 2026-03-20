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

func (k msgServer) BreedAgent(goCtx context.Context, msg *types.MsgBreedAgent) (*types.MsgBreedAgentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	pa, foundA := k.GetAgent(ctx, msg.ParentA)
	pb, foundB := k.GetAgent(ctx, msg.ParentB)
	if !foundA || !foundB {
		return nil, fmt.Errorf("one or both parent agents not found")
	}
	if _, found := k.GetAgent(ctx, msg.ChildId); found {
		return nil, fmt.Errorf("agent %s already exists", msg.ChildId)
	}
	inherit := func(a, b uint64) uint64 {
		base := (a + b) / 2
		m    := int64(base) / 8
		rng  := rand.New(rand.NewSource(ctx.BlockHeight() + int64(a+b)))
		d    := rng.Int63n(m*2+1) - m
		r    := int64(base) + d
		if r < 1   { r = 1 }
		if r > 100 { r = 100 }
		return uint64(r)
	}
	s  := inherit(pa.Strength, pb.Strength)
	a  := inherit(pa.Agility,  pb.Agility)
	st := inherit(pa.Stamina,  pb.Stamina)
	sk := inherit(pa.Skill,    pb.Skill)
	score := int32((s + a + st + sk) / 4)
	gen := pa.Generation
	if pb.Generation > gen { gen = pb.Generation }
	gen++
	raw     := fmt.Sprintf("%s:gen2:%d:%d:%d:%d:%d", msg.ChildId, s, a, st, sk, gen)
	hash    := sha256.Sum256([]byte(raw))
	dnaHash := hex.EncodeToString(hash[:])
	child := types.Agent{
		Index: msg.ChildId, NodeId: msg.ChildId, Name: msg.ChildName,
		AgentType: "gen2", Owner: msg.Creator, DnaHash: dnaHash,
		Strength: s, Agility: a, Stamina: st, Skill: sk,
		Generation: gen, Active: true,
		GenesisBlock: int32(ctx.BlockHeight()),
		ParentA: msg.ParentA, ParentB: msg.ParentB,
	}
	k.SetAgent(ctx, child)
	ctx.EventManager().EmitEvent(sdk.NewEvent("agent_bred",
		sdk.NewAttribute("child_id", msg.ChildId),
		sdk.NewAttribute("parent_a", msg.ParentA),
		sdk.NewAttribute("parent_b", msg.ParentB),
		sdk.NewAttribute("genome_score", fmt.Sprintf("%d", score)),
	))
	return &types.MsgBreedAgentResponse{ChildId: msg.ChildId, DnaHash: dnaHash, GenomeScore: score}, nil
}
