#!/bin/bash
# Run this from ~/oan
# Writes all oanprotocol business logic directly — no downloads needed

echo "Writing oanprotocol business logic..."

# ── params.go ──────────────────────────────────────────────
cat > ~/oan/x/oanprotocol/types/params.go << 'GOEOF'
package types

import paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

var _ paramtypes.ParamSet = (*Params)(nil)

func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams() Params {
	return Params{
		MaxWalletBalance:       800_000_000_000_000,
		VerifiedMaxBalance:     400_000_000_000_000,
		UnverifiedMaxBalance:   4_000_000_000_000,
		MaxDailyReceive:        2_000_000_000_000,
		UnverifiedDailyReceive: 40_000_000_000,
		EpochLength:            14400,
		HardCap:                50_000_000_000_000_000,
		LaunchPhase:            "genesis",
		InflationBps:           100,
		ProtectionsEnabled:     true,
	}
}

func DefaultParams() Params { return NewParams() }

func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{}
}

func (p Params) Validate() error { return nil }
GOEOF

echo "✅ params.go"

# ── expected_keepers.go ────────────────────────────────────
cat > ~/oan/x/oanprotocol/types/expected_keepers.go << 'GOEOF'
package types

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type AccountKeeper interface {
	GetAccount(context.Context, sdk.AccAddress) sdk.AccountI
}

type BankKeeper interface {
	SpendableCoins(context.Context, sdk.AccAddress) sdk.Coins
	GetBalance(context.Context, sdk.AccAddress, string) sdk.Coin
}

type ParamSubspace interface {
	Get(context.Context, []byte, interface{})
	Set(context.Context, []byte, interface{})
}
GOEOF

echo "✅ expected_keepers.go"

# ── keeper.go ──────────────────────────────────────────────
cat > ~/oan/x/oanprotocol/keeper/keeper.go << 'GOEOF'
package keeper

import (
	"fmt"
	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"oan/x/oanprotocol/types"
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
	k.Logger().Info("oanprotocol: epoch reset", "block", ctx.BlockHeight())
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

func (k Keeper) CheckTransferAllowed(ctx sdk.Context, toAddress string, amountUoan uint64) error {
	params := k.GetParams(ctx)
	if !params.ProtectionsEnabled {
		return nil
	}
	tier := k.GetAddressTier(ctx, toAddress)
	if tier == "exempt" {
		return nil
	}

	// Daily limit
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
		return fmt.Errorf("daily receive limit exceeded: limit=%d uoan (%d OAN/day), received=%d uoan, remaining=%d uoan — resets every 14400 blocks",
			dailyLimit, dailyLimit/1_000_000, dailyReceived, remaining)
	}

	// Wallet balance cap
	addr, err := sdk.AccAddressFromBech32(toAddress)
	if err != nil {
		return nil
	}
	currentBalance := k.bankKeeper.GetBalance(ctx, addr, "uoan")
	newBalance := currentBalance.Amount.Uint64() + amountUoan

	switch tier {
	case "unverified":
		if newBalance > params.UnverifiedMaxBalance {
			return fmt.Errorf("wallet cap exceeded: unverified max=%d OAN — register your DID to hold up to 400,000 OAN",
				params.UnverifiedMaxBalance/1_000_000)
		}
	case "verified":
		if newBalance > params.VerifiedMaxBalance {
			return fmt.Errorf("wallet cap exceeded: verified max=%d OAN — apply for DAO approval to hold up to 800,000 OAN",
				params.VerifiedMaxBalance/1_000_000)
		}
	case "dao":
		if newBalance > params.MaxWalletBalance {
			return fmt.Errorf("wallet cap exceeded: absolute max=%d OAN (2%% of supply)",
				params.MaxWalletBalance/1_000_000)
		}
	}

	k.AddDailyReceived(ctx, toAddress, amountUoan)
	return nil
}
GOEOF

echo "✅ keeper.go"

# ── msg_server_set_address_tier.go ────────────────────────
cat > ~/oan/x/oanprotocol/keeper/msg_server_set_address_tier.go << 'GOEOF'
package keeper

import (
	"context"
	"fmt"
	"oan/x/oanprotocol/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SetAddressTier(goCtx context.Context, msg *types.MsgSetAddressTier) (*types.MsgSetAddressTierResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if msg.Creator != k.GetAuthority() {
		return nil, fmt.Errorf("unauthorized: only governance can set address tiers")
	}
	validTiers := map[string]bool{"unverified": true, "verified": true, "dao": true, "exempt": true}
	if !validTiers[msg.Tier] {
		return nil, fmt.Errorf("invalid tier: must be unverified, verified, dao, or exempt")
	}
	k.SetAddressTierStore(ctx, msg.Address, msg.Tier)
	if msg.Tier == "exempt" {
		k.ExemptAddress(ctx, msg.Address)
	}
	ctx.EventManager().EmitEvent(sdk.NewEvent("tier_set",
		sdk.NewAttribute("address", msg.Address),
		sdk.NewAttribute("tier", msg.Tier),
		sdk.NewAttribute("set_by", msg.Creator),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgSetAddressTierResponse{}, nil
}
GOEOF

echo "✅ msg_server_set_address_tier.go"

# ── msg_server_update_launch_phase.go ─────────────────────
cat > ~/oan/x/oanprotocol/keeper/msg_server_update_launch_phase.go << 'GOEOF'
package keeper

import (
	"context"
	"fmt"
	"oan/x/oanprotocol/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateLaunchPhase(goCtx context.Context, msg *types.MsgUpdateLaunchPhase) (*types.MsgUpdateLaunchPhaseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if msg.Creator != k.GetAuthority() {
		return nil, fmt.Errorf("unauthorized: only governance can update launch phase")
	}
	validPhases := map[string]bool{
		"genesis": true, "validator": true, "ubi": true,
		"ecosystem": true, "dex": true, "open": true,
	}
	if !validPhases[msg.Phase] {
		return nil, fmt.Errorf("invalid phase: must be genesis/validator/ubi/ecosystem/dex/open")
	}
	params := k.GetParams(ctx)
	oldPhase := params.LaunchPhase
	params.LaunchPhase = msg.Phase
	k.SetParams(ctx, params)
	ctx.EventManager().EmitEvent(sdk.NewEvent("launch_phase_updated",
		sdk.NewAttribute("old_phase", oldPhase),
		sdk.NewAttribute("new_phase", msg.Phase),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	k.Logger().Info("oanprotocol: launch phase advanced", "from", oldPhase, "to", msg.Phase)
	return &types.MsgUpdateLaunchPhaseResponse{}, nil
}
GOEOF

echo "✅ msg_server_update_launch_phase.go"

# ── msg_server_exempt_address.go ──────────────────────────
cat > ~/oan/x/oanprotocol/keeper/msg_server_exempt_address.go << 'GOEOF'
package keeper

import (
	"context"
	"fmt"
	"oan/x/oanprotocol/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ExemptAddress(goCtx context.Context, msg *types.MsgExemptAddress) (*types.MsgExemptAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if msg.Creator != k.GetAuthority() {
		return nil, fmt.Errorf("unauthorized: only governance can exempt addresses")
	}
	k.Keeper.ExemptAddress(ctx, msg.Address)
	k.SetAddressTierStore(ctx, msg.Address, "exempt")
	ctx.EventManager().EmitEvent(sdk.NewEvent("address_exempted",
		sdk.NewAttribute("address", msg.Address),
		sdk.NewAttribute("reason", msg.Reason),
		sdk.NewAttribute("by", msg.Creator),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgExemptAddressResponse{}, nil
}
GOEOF

echo "✅ msg_server_exempt_address.go"

# ── msg_server_update_protocol_params.go ──────────────────
cat > ~/oan/x/oanprotocol/keeper/msg_server_update_protocol_params.go << 'GOEOF'
package keeper

import (
	"context"
	"fmt"
	"oan/x/oanprotocol/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateProtocolParams(goCtx context.Context, msg *types.MsgUpdateProtocolParams) (*types.MsgUpdateProtocolParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if msg.Creator != k.GetAuthority() {
		return nil, fmt.Errorf("unauthorized: only governance can update protocol params")
	}
	params := k.GetParams(ctx)
	k.SetParams(ctx, params)
	ctx.EventManager().EmitEvent(sdk.NewEvent("protocol_params_updated",
		sdk.NewAttribute("updated_by", msg.Creator),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgUpdateProtocolParamsResponse{}, nil
}
GOEOF

echo "✅ msg_server_update_protocol_params.go"

# ── query_address_status.go ───────────────────────────────
cat > ~/oan/x/oanprotocol/keeper/query_address_status.go << 'GOEOF'
package keeper

import (
	"context"
	"fmt"
	"oan/x/oanprotocol/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AddressStatus(goCtx context.Context, req *types.QueryAddressStatusRequest) (*types.QueryAddressStatusResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := k.GetParams(ctx)
	tier := k.GetAddressTier(ctx, req.Address)
	dailyReceived := k.GetDailyReceived(ctx, req.Address)

	dailyLimit := params.MaxDailyReceive
	maxBalance := params.MaxWalletBalance
	switch tier {
	case "unverified":
		dailyLimit = params.UnverifiedDailyReceive
		maxBalance = params.UnverifiedMaxBalance
	case "verified":
		maxBalance = params.VerifiedMaxBalance
	}

	remaining := uint64(0)
	if dailyLimit > dailyReceived {
		remaining = dailyLimit - dailyReceived
	}
	_ = remaining

	return &types.QueryAddressStatusResponse{
		Tier:          tier,
		DailyReceived: fmt.Sprintf("%d", dailyReceived),
		DailyLimit:    fmt.Sprintf("%d", dailyLimit),
		MaxBalance:    fmt.Sprintf("%d", maxBalance),
	}, nil
}
GOEOF

echo "✅ query_address_status.go"

echo ""
echo "All files written. Building..."
cd ~/oan
ignite chain build 2>&1 | tail -40
