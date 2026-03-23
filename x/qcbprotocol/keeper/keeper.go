package keeper

import (
	"fmt"
	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"qcb/x/qcbprotocol/types"
)

type (
	Keeper struct {
		cdc          codec.BinaryCodec
		storeService store.KVStoreService
		logger       log.Logger
		authority    string
		bankKeeper   types.BankKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	logger log.Logger,
	authority string,
	bankKeeper types.BankKeeper,
) Keeper {
	if _, err := sdk.AccAddressFromBech32(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address: %s", authority))
	}
	return Keeper{
		cdc:          cdc,
		storeService: storeService,
		authority:    authority,
		logger:       logger,
		bankKeeper:   bankKeeper,
	}
}

func (k Keeper) GetAuthority() string { return k.authority }

func (k Keeper) Logger() log.Logger {
	return k.logger.With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// ── TIER SYSTEM ──────────────────────────────────────────

func (k Keeper) GetAddressTier(ctx sdk.Context, address string) string {
	store := k.storeService.OpenKVStore(ctx)
	exemptKey := fmt.Sprintf("exempt-%s", address)
	if v, _ := store.Get([]byte(exemptKey)); v != nil {
		return "exempt"
	}
	daoKey := fmt.Sprintf("tier-%s", address)
	if t, _ := store.Get([]byte(daoKey)); t != nil {
		return string(t)
	}
	didKey := fmt.Sprintf("did-verified-%s", address)
	if v, _ := store.Get([]byte(didKey)); v != nil {
		return "verified"
	}
	return "unverified"
}

func (k Keeper) SetAddressTierStore(ctx sdk.Context, address string, tier string) {
	store := k.storeService.OpenKVStore(ctx)
	store.Set([]byte(fmt.Sprintf("tier-%s", address)), []byte(tier))
}

func (k Keeper) ExemptAddress(ctx sdk.Context, address string) {
	store := k.storeService.OpenKVStore(ctx)
	store.Set([]byte(fmt.Sprintf("exempt-%s", address)), []byte("1"))
}

// ── EPOCH & DAILY LIMITS ─────────────────────────────────

func (k Keeper) GetEpochStart(ctx sdk.Context) int64 {
	store := k.storeService.OpenKVStore(ctx)
	b, _ := store.Get([]byte("epoch-start"))
	if b == nil {
		return 0
	}
	var h int64
	fmt.Sscanf(string(b), "%d", &h)
	return h
}

func (k Keeper) ResetEpoch(ctx sdk.Context) {
	store := k.storeService.OpenKVStore(ctx)
	store.Set([]byte("epoch-start"), []byte(fmt.Sprintf("%d", ctx.BlockHeight())))
	k.Logger().Info("qcbprotocol: epoch reset", "block", ctx.BlockHeight())
}

func (k Keeper) GetDailyReceived(ctx sdk.Context, address string) uint64 {
	store := k.storeService.OpenKVStore(ctx)
	params := k.GetParams(ctx)
	epochStart := k.GetEpochStart(ctx)
	if ctx.BlockHeight()-epochStart >= int64(params.EpochLength) {
		k.ResetEpoch(ctx)
		return 0
	}
	b, _ := store.Get([]byte(fmt.Sprintf("daily-recv-%s", address)))
	if b == nil {
		return 0
	}
	var amount uint64
	fmt.Sscanf(string(b), "%d", &amount)
	return amount
}

func (k Keeper) AddDailyReceived(ctx sdk.Context, address string, amount uint64) {
	store := k.storeService.OpenKVStore(ctx)
	current := k.GetDailyReceived(ctx, address)
	store.Set(
		[]byte(fmt.Sprintf("daily-recv-%s", address)),
		[]byte(fmt.Sprintf("%d", current+amount)),
	)
}

// ── HARD CAP CHECK ───────────────────────────────────────

// CheckHardCap verifies minting or transferring amount won't exceed 50M OAN
func (k Keeper) CheckHardCap(ctx sdk.Context, additionalUoan uint64) error {
	params := k.GetParams(ctx)
	if !params.ProtectionsEnabled {
		return nil
	}

	// Get real total supply from bank module
	totalSupply := k.bankKeeper.GetSupply(ctx, "charmbits")
	currentSupply := totalSupply.Amount.Uint64()
	newSupply := currentSupply + additionalUoan

	if newSupply > params.HardCap {
		return fmt.Errorf(
			"hard cap exceeded: absolute max=%d OAN (50,000,000 OAN forever), current supply=%d OAN, attempted new supply=%d OAN — Arkadina hard cap enforced",
			params.HardCap/1_000_000,
			currentSupply/1_000_000,
			newSupply/1_000_000,
		)
	}

	k.Logger().Info("qcbprotocol: hard cap check passed",
		"current_supply_oan", currentSupply/1_000_000,
		"hard_cap_oan", params.HardCap/1_000_000,
		"remaining_oan", (params.HardCap-currentSupply)/1_000_000,
	)

	return nil
}

// ── CORE PROTECTION CHECK ─────────────────────────────────

// CheckTransferAllowed is the main gate — validates all protections
func (k Keeper) CheckTransferAllowed(ctx sdk.Context, toAddress string, amountUoan uint64) error {
	params := k.GetParams(ctx)
	if !params.ProtectionsEnabled {
		return nil
	}

	tier := k.GetAddressTier(ctx, toAddress)
	if tier == "exempt" {
		return nil
	}

	// ── HARD CAP ──
	if err := k.CheckHardCap(ctx, 0); err != nil {
		return err
	}

	// ── DAILY RECEIVE LIMIT ──
	dailyReceived := k.GetDailyReceived(ctx, toAddress)
	dailyLimit := params.MaxDailyReceive
	if tier == "unverified" {
		dailyLimit = params.UnverifiedDailyReceive
	}
	if dailyReceived+amountUoan > dailyLimit {
		remaining := uint64(0)
		if dailyLimit > dailyReceived {
			remaining = dailyLimit - dailyReceived
		}
		return fmt.Errorf(
			"daily receive limit exceeded: limit=%d OAN/day, received today=%d charmbits, remaining=%d charmbits — resets every 14400 blocks (~1 day)",
			dailyLimit/1_000_000, dailyReceived, remaining,
		)
	}

	// ── WALLET BALANCE CAP ──
	addr, err := sdk.AccAddressFromBech32(toAddress)
	if err != nil {
		return nil
	}
	currentBalance := k.bankKeeper.GetBalance(ctx, addr, "charmbits")
	newBalance := currentBalance.Amount.Uint64() + amountUoan

	switch tier {
	case "unverified":
		if newBalance > params.UnverifiedMaxBalance {
			return fmt.Errorf(
				"wallet cap exceeded: unverified max=%d OAN — register your DID identity to hold up to 400,000 OAN",
				params.UnverifiedMaxBalance/1_000_000,
			)
		}
	case "verified":
		if newBalance > params.VerifiedMaxBalance {
			return fmt.Errorf(
				"wallet cap exceeded: verified max=%d OAN — apply for DAO approval to hold up to 800,000 OAN",
				params.VerifiedMaxBalance/1_000_000,
			)
		}
	case "dao":
		if newBalance > params.MaxWalletBalance {
			return fmt.Errorf(
				"wallet cap exceeded: absolute max=%d OAN (2%% of total supply)",
				params.MaxWalletBalance/1_000_000,
			)
		}
	}

	// All checks passed — record the daily receive
	k.AddDailyReceived(ctx, toAddress, amountUoan)

	return nil
}

// ── SUPPLY INFO ───────────────────────────────────────────

// GetSupplyInfo returns current supply stats
func (k Keeper) GetSupplyInfo(ctx sdk.Context) map[string]uint64 {
	params := k.GetParams(ctx)
	totalSupply := k.bankKeeper.GetSupply(ctx, "charmbits")
	currentSupply := totalSupply.Amount.Uint64()
	remaining := uint64(0)
	if params.HardCap > currentSupply {
		remaining = params.HardCap - currentSupply
	}
	return map[string]uint64{
		"current_supply_charmbits": currentSupply,
		"hard_cap_charmbits":       params.HardCap,
		"remaining_charmbits":      remaining,
		"current_supply_oan":  currentSupply / 1_000_000,
		"hard_cap_oan":        params.HardCap / 1_000_000,
		"remaining_oan":       remaining / 1_000_000,
	}
}
