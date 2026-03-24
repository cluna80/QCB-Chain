package keeper

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// IntrinsicCharmEpochLength — how many blocks between UBI pool replenishments
// 14400 blocks = ~1 day at 6s block time
const IntrinsicCharmEpochLength = int64(14400)

// IntrinsicCharmBps — basis points of fee pool sent to UBI every epoch
// 50 bps = 0.5% (range: 10 = 0.1% to 100 = 1.0%)
const IntrinsicCharmBps = int64(50)

// RunIntrinsicCharm checks if an epoch has passed and if so
// sends a probabilistic fee slice to the UBI pool.
// Called from EndBlock every block.
func (k Keeper) RunIntrinsicCharm(ctx sdk.Context) {
	store := k.storeService.OpenKVStore(ctx)
	height := ctx.BlockHeight()

	// Read last epoch block
	epochBytes, _ := store.Get([]byte("intrinsic-charm-epoch"))
	lastEpoch := int64(0)
	if epochBytes != nil {
		fmt.Sscanf(string(epochBytes), "%d", &lastEpoch)
	}

	// Not yet time for next epoch
	if height-lastEpoch < IntrinsicCharmEpochLength {
		return
	}

	// Update epoch
	store.Set([]byte("intrinsic-charm-epoch"), []byte(fmt.Sprintf("%d", height)))

	// Read accumulated fee pool
	feePoolBytes, _ := store.Get([]byte("charm-fee-pool"))
	feePool := int64(0)
	if feePoolBytes != nil {
		fmt.Sscanf(string(feePoolBytes), "%d", &feePool)
	}

	if feePool <= 0 {
		k.Logger().Info("intrinsic charm: fee pool empty, skipping epoch",
			"block", height)
		return
	}

	// Calculate slice: IntrinsicCharmBps / 10000 of fee pool
	slice := (feePool * IntrinsicCharmBps) / 10000
	if slice <= 0 {
		return
	}

	// Read current UBI pool
	ubiBytes, _ := store.Get([]byte("ubi-pool-balance"))
	ubiPool := int64(0)
	if ubiBytes != nil {
		fmt.Sscanf(string(ubiBytes), "%d", &ubiPool)
	}

	// Add slice to UBI pool
	newUbi := ubiPool + slice
	newFeePool := feePool - slice

	store.Set([]byte("ubi-pool-balance"), []byte(fmt.Sprintf("%d", newUbi)))
	store.Set([]byte("charm-fee-pool"), []byte(fmt.Sprintf("%d", newFeePool)))

	ctx.EventManager().EmitEvent(sdk.NewEvent("intrinsic_charm",
		sdk.NewAttribute("epoch_block", fmt.Sprintf("%d", height)),
		sdk.NewAttribute("fee_pool_charmbits", fmt.Sprintf("%d", feePool)),
		sdk.NewAttribute("slice_charmbits", fmt.Sprintf("%d", slice)),
		sdk.NewAttribute("slice_bps", fmt.Sprintf("%d", IntrinsicCharmBps)),
		sdk.NewAttribute("ubi_pool_before", fmt.Sprintf("%d", ubiPool)),
		sdk.NewAttribute("ubi_pool_after", fmt.Sprintf("%d", newUbi)),
	))

	k.Logger().Info("⚛️  intrinsic charm fired",
		"block", height,
		"fee_slice_charmbits", slice,
		"ubi_pool_charmbits", newUbi,
	)
}

// AddToFeePool records fees collected — called when transactions are processed
func (k Keeper) AddToFeePool(ctx sdk.Context, amount int64) {
	if amount <= 0 {
		return
	}
	store := k.storeService.OpenKVStore(ctx)
	feePoolBytes, _ := store.Get([]byte("charm-fee-pool"))
	current := int64(0)
	if feePoolBytes != nil {
		fmt.Sscanf(string(feePoolBytes), "%d", &current)
	}
	store.Set([]byte("charm-fee-pool"), []byte(fmt.Sprintf("%d", current+amount)))
}

// GetIntrinsicCharmStats returns current pool stats for the explorer
func (k Keeper) GetIntrinsicCharmStats(ctx sdk.Context) map[string]int64 {
	store := k.storeService.OpenKVStore(ctx)

	feePoolBytes, _ := store.Get([]byte("charm-fee-pool"))
	feePool := int64(0)
	if feePoolBytes != nil {
		fmt.Sscanf(string(feePoolBytes), "%d", &feePool)
	}

	ubiBytes, _ := store.Get([]byte("ubi-pool-balance"))
	ubiPool := int64(0)
	if ubiBytes != nil {
		fmt.Sscanf(string(ubiBytes), "%d", &ubiPool)
	}

	epochBytes, _ := store.Get([]byte("intrinsic-charm-epoch"))
	lastEpoch := int64(0)
	if epochBytes != nil {
		fmt.Sscanf(string(epochBytes), "%d", &lastEpoch)
	}

	blocksUntilNext := IntrinsicCharmEpochLength - (ctx.BlockHeight() - lastEpoch)
	if blocksUntilNext < 0 {
		blocksUntilNext = 0
	}

	return map[string]int64{
		"fee_pool_charmbits":    feePool,
		"ubi_pool_charmbits":    ubiPool,
		"last_epoch_block":      lastEpoch,
		"blocks_until_next":     blocksUntilNext,
		"epoch_length_blocks":   IntrinsicCharmEpochLength,
		"slice_bps":             IntrinsicCharmBps,
	}
}
