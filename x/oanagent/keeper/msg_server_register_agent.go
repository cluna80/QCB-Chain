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

func (k msgServer) RegisterAgent(goCtx context.Context, msg *types.MsgRegisterAgent) (*types.MsgRegisterAgentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if _, found := k.GetAgent(ctx, msg.NodeId); found {
		return nil, fmt.Errorf("agent %s already registered", msg.NodeId)
	}
	seed := int64(0)
	for _, c := range msg.NodeId { seed += int64(c) }
	rng := rand.New(rand.NewSource(seed + ctx.BlockHeight()))
	s  := uint64(50 + rng.Intn(50))
	a  := uint64(50 + rng.Intn(50))
	st := uint64(50 + rng.Intn(50))
	sk := uint64(50 + rng.Intn(50))
	raw     := fmt.Sprintf("%s:%s:%d:%d:%d:%d:1", msg.NodeId, msg.AgentType, s, a, st, sk)
	hash    := sha256.Sum256([]byte(raw))
	dnaHash := hex.EncodeToString(hash[:])
	agent := types.Agent{
		Index: msg.NodeId, NodeId: msg.NodeId, Name: msg.Name,
		AgentType: msg.AgentType, Owner: msg.Creator, DnaHash: dnaHash,
		Strength: s, Agility: a, Stamina: st, Skill: sk,
		Generation: 1, Active: true,
		GenesisBlock: int32(ctx.BlockHeight()),
	}
	k.SetAgent(ctx, agent)
	ctx.EventManager().EmitEvent(sdk.NewEvent("agent_registered",
		sdk.NewAttribute("node_id", msg.NodeId),
		sdk.NewAttribute("dna_hash", dnaHash),
		sdk.NewAttribute("type", msg.AgentType),
	))
	return &types.MsgRegisterAgentResponse{NodeId: msg.NodeId, DnaHash: dnaHash}, nil
}
