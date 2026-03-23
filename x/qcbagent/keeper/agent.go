package keeper

import (
	"context"

	"qcb/x/qcbagent/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetAgent set a specific agent in the store from its index
func (k Keeper) SetAgent(ctx context.Context, agent types.Agent) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AgentKeyPrefix))
	b := k.cdc.MustMarshal(&agent)
	store.Set(types.AgentKey(
		agent.Index,
	), b)
}

// GetAgent returns a agent from its index
func (k Keeper) GetAgent(
	ctx context.Context,
	index string,

) (val types.Agent, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AgentKeyPrefix))

	b := store.Get(types.AgentKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAgent removes a agent from the store
func (k Keeper) RemoveAgent(
	ctx context.Context,
	index string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AgentKeyPrefix))
	store.Delete(types.AgentKey(
		index,
	))
}

// GetAllAgent returns all agent
func (k Keeper) GetAllAgent(ctx context.Context) (list []types.Agent) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AgentKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Agent
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
