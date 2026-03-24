package keeper

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"qcb/x/charm/types"
)

func (k Keeper) RunIntrinsicCharm(ctx sdk.Context) {
	store := k.storeService.OpenKVStore(ctx)
	height := ctx.BlockHeight()
	params := k.GetParams(ctx)

	epochBytes, _ := store.Get([]byte(types.KeyLastEpoch))
	lastEpoch := int64(0)
	if epochBytes != nil { fmt.Sscanf(string(epochBytes), "%d", &lastEpoch) }
	if height-lastEpoch < params.EpochBlocks { return }

	store.Set([]byte(types.KeyLastEpoch), []byte(fmt.Sprintf("%d", height)))

	feePoolBytes, _ := store.Get([]byte(types.KeyFeePool))
	feePool := int64(0)
	if feePoolBytes != nil { fmt.Sscanf(string(feePoolBytes), "%d", &feePool) }
	if feePool <= 0 { return }

	slice := (feePool * params.FeesliceBps) / 10000
	if slice <= 0 { return }

	ubiBytes, _ := store.Get([]byte(types.KeyUBIPool))
	ubiPool := int64(0)
	if ubiBytes != nil { fmt.Sscanf(string(ubiBytes), "%d", &ubiPool) }

	store.Set([]byte(types.KeyUBIPool), []byte(fmt.Sprintf("%d", ubiPool+slice)))
	store.Set([]byte(types.KeyFeePool), []byte(fmt.Sprintf("%d", feePool-slice)))

	ctx.EventManager().EmitEvent(sdk.NewEvent("intrinsic_charm",
		sdk.NewAttribute("epoch_block", fmt.Sprintf("%d", height)),
		sdk.NewAttribute("slice_charmbits", fmt.Sprintf("%d", slice)),
		sdk.NewAttribute("ubi_pool_after", fmt.Sprintf("%d", ubiPool+slice)),
	))
	k.Logger().Info("⚛️  intrinsic charm fired", "block", height, "slice", slice, "ubi_pool", ubiPool+slice)
}

func (k Keeper) AddToFeePool(ctx sdk.Context, amount int64) {
	if amount <= 0 { return }
	store := k.storeService.OpenKVStore(ctx)
	b, _ := store.Get([]byte(types.KeyFeePool))
	current := int64(0)
	if b != nil { fmt.Sscanf(string(b), "%d", &current) }
	store.Set([]byte(types.KeyFeePool), []byte(fmt.Sprintf("%d", current+amount)))
}

func (k Keeper) GetIntrinsicCharmStats(ctx sdk.Context) map[string]int64 {
	store := k.storeService.OpenKVStore(ctx)
	params := k.GetParams(ctx)
	get := func(key string) int64 {
		b, _ := store.Get([]byte(key))
		if b == nil { return 0 }
		var v int64; fmt.Sscanf(string(b), "%d", &v); return v
	}
	lastEpoch := get(types.KeyLastEpoch)
	next := params.EpochBlocks - (ctx.BlockHeight() - lastEpoch)
	if next < 0 { next = 0 }
	return map[string]int64{
		"fee_pool_charmbits":  get(types.KeyFeePool),
		"ubi_pool_charmbits":  get(types.KeyUBIPool),
		"last_epoch_block":    lastEpoch,
		"blocks_until_next":   next,
		"epoch_length_blocks": params.EpochBlocks,
		"slice_bps":           params.FeesliceBps,
	}
}
