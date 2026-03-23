package keeper

import (
	"context"
	"fmt"
	"qcb/x/qcbmarket/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) HireAgent(goCtx context.Context, msg *types.MsgHireAgent) (*types.MsgHireAgentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := k.storeService.OpenKVStore(ctx)

	if msg.ListingId == "" {
		return nil, fmt.Errorf("listingId cannot be empty")
	}
	if msg.Budget == 0 {
		return nil, fmt.Errorf("budget must be greater than 0")
	}
	// SECURITY — minimum amount anti-dust
	if msg.Budget < 100 {
		return nil, fmt.Errorf("minimum hire budget is 100 charmbits")
	}
	// SECURITY — overflow check
	if msg.Budget > 9223372036854775807 {
		return nil, fmt.Errorf("budget amount overflow")
	}

	listingKey := fmt.Sprintf("listing-%s", msg.ListingId)
	listingData, _ := store.Get([]byte(listingKey))
	if listingData == nil {
		return nil, fmt.Errorf("listing %s not found", msg.ListingId)
	}

	// Get owner from listing
	ownerKey := fmt.Sprintf("listing-owner-%s", msg.ListingId)
	ownerBytes, _ := store.Get([]byte(ownerKey))
	if ownerBytes == nil {
		return nil, fmt.Errorf("listing owner not found for %s", msg.ListingId)
	}

	owner, err := sdk.AccAddressFromBech32(string(ownerBytes))
	if err != nil {
		return nil, fmt.Errorf("invalid owner address: %s", err)
	}
	hirer, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, fmt.Errorf("invalid hirer address: %s", err)
	}

	// Cannot hire your own agent
	if string(ownerBytes) == msg.Creator {
		return nil, fmt.Errorf("cannot hire your own agent")
	}

	params := k.GetParams(ctx)
	fee := params.MarketFee
	if fee == 0 {
		fee = 20
	}
	feeAmount := msg.Budget * fee / 1000
	ownerAmount := msg.Budget - feeAmount

	// REAL TOKEN TRANSFER — pay owner
	ownerCoins := sdk.NewCoins(sdk.NewInt64Coin("charmbits", int64(ownerAmount)))
	if err := k.bankKeeper.SendCoins(ctx, hirer, owner, ownerCoins); err != nil {
		return nil, fmt.Errorf("payment failed — insufficient balance: %s", err)
	}

	// Collect marketplace fee
	if feeAmount > 0 {
		feeCoins := sdk.NewCoins(sdk.NewInt64Coin("charmbits", int64(feeAmount)))
		if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, hirer, types.ModuleName, feeCoins); err != nil {
			fmt.Printf("Fee collection skipped: %s\n", err)
		}
	}

	contractId := fmt.Sprintf("contract-%d-%s", ctx.BlockHeight(), msg.ListingId)
	store.Set([]byte(fmt.Sprintf("contract-%s", contractId)),
		[]byte(fmt.Sprintf("%s|%s|%s|%d|active",
			contractId, msg.ListingId, msg.Creator, ctx.BlockHeight())))

	ctx.EventManager().EmitEvent(sdk.NewEvent("agent_hired",
		sdk.NewAttribute("contract_id", contractId),
		sdk.NewAttribute("listing_id", msg.ListingId),
		sdk.NewAttribute("hirer", msg.Creator),
		sdk.NewAttribute("budget", fmt.Sprintf("%d", msg.Budget)),
		sdk.NewAttribute("owner_received", fmt.Sprintf("%d", ownerAmount)),
		sdk.NewAttribute("fee", fmt.Sprintf("%d", feeAmount)),
		sdk.NewAttribute("denom", "charmbits"),
	))
	return &types.MsgHireAgentResponse{ContractId: contractId}, nil
}
