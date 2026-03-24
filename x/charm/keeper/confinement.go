package keeper

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"qcb/x/charm/types"
)

func (k Keeper) GetAddressTier(ctx sdk.Context, address string) string {
	store := k.storeService.OpenKVStore(ctx)
	b, _ := store.Get([]byte(types.KeyBalanceTierPrefix + address))
	if b == nil { return types.TierUnverified }
	return string(b)
}

func (k Keeper) SetAddressTier(ctx sdk.Context, address string, tier string) {
	store := k.storeService.OpenKVStore(ctx)
	store.Set([]byte(types.KeyBalanceTierPrefix+address), []byte(tier))
}

func (k Keeper) GetDailyReceive(ctx sdk.Context, address string) int64 {
	store := k.storeService.OpenKVStore(ctx)
	b, _ := store.Get([]byte(types.KeyDailyReceivePrefix + address))
	if b == nil { return 0 }
	var v int64; fmt.Sscanf(string(b), "%d", &v); return v
}

func (k Keeper) RecordDailyReceive(ctx sdk.Context, address string, amount int64) {
	store := k.storeService.OpenKVStore(ctx)
	current := k.GetDailyReceive(ctx, address)
	store.Set([]byte(types.KeyDailyReceivePrefix+address), []byte(fmt.Sprintf("%d", current+amount)))
}

func (k Keeper) CheckConfinement(ctx sdk.Context, toAddress string, amount int64) (bool, string) {
	tier := k.GetAddressTier(ctx, toAddress)
	if tier == types.TierExempt { return true, "" }
	params := k.GetParams(ctx)
	daily := k.GetDailyReceive(ctx, toAddress)
	var limit int64
	switch tier {
	case types.TierVerified:    limit = params.VerifiedDailyLimit
	case types.TierDAOApproved: limit = params.VerifiedDailyLimit
	default:                    limit = params.UnverifiedDailyLimit
	}
	if daily+amount > limit {
		return false, fmt.Sprintf("charm confinement: daily limit exceeded (tier=%s, limit=%d, received=%d+%d)", tier, limit, daily, amount)
	}
	return true, ""
}
