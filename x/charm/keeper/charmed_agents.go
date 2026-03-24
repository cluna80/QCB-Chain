package keeper

import (
	"encoding/json"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"qcb/x/charm/types"
)

func (k Keeper) CreateCharmedPair(ctx sdk.Context, humanAddress string, agentID string) error {
	store := k.storeService.OpenKVStore(ctx)
	key := []byte(types.KeyCharmedPairPrefix + humanAddress)
	if existing, _ := store.Get(key); existing != nil { return types.ErrAlreadyPaired }
	pair := types.CharmedPair{HumanAddress: humanAddress, AgentID: agentID, BondedAt: ctx.BlockHeight(), YieldBonus: k.GetParams(ctx).CharmedPairBonus, Active: true}
	b, err := json.Marshal(pair)
	if err != nil { return err }
	store.Set(key, b)
	ctx.EventManager().EmitEvent(sdk.NewEvent("charmed_pair_created",
		sdk.NewAttribute("human_address", humanAddress),
		sdk.NewAttribute("agent_id", agentID),
		sdk.NewAttribute("yield_bonus_bps", fmt.Sprintf("%d", pair.YieldBonus)),
	))
	k.Logger().Info("⚛️  charmed pair bonded", "agent", agentID, "bonus_bps", pair.YieldBonus)
	return nil
}

func (k Keeper) GetCharmedPair(ctx sdk.Context, humanAddress string) (*types.CharmedPair, bool) {
	store := k.storeService.OpenKVStore(ctx)
	b, _ := store.Get([]byte(types.KeyCharmedPairPrefix + humanAddress))
	if b == nil { return nil, false }
	var pair types.CharmedPair
	if err := json.Unmarshal(b, &pair); err != nil { return nil, false }
	return &pair, true
}

func (k Keeper) DissolveCharmedPair(ctx sdk.Context, humanAddress string) error {
	store := k.storeService.OpenKVStore(ctx)
	key := []byte(types.KeyCharmedPairPrefix + humanAddress)
	if existing, _ := store.Get(key); existing == nil { return types.ErrPairNotFound }
	store.Delete(key)
	k.Logger().Info("charmed pair dissolved", "human", humanAddress)
	return nil
}

func (k Keeper) GetYieldBonus(ctx sdk.Context, address string) int64 {
	pair, found := k.GetCharmedPair(ctx, address)
	if !found || !pair.Active { return 0 }
	return pair.YieldBonus
}
