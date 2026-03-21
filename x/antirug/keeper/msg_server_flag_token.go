package keeper

import (
	"context"
	"fmt"
	"oan/x/antirug/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) FlagToken(goCtx context.Context, msg *types.MsgFlagToken) (*types.MsgFlagTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := k.storeService.OpenKVStore(ctx)

	// NO DORMANCY CHECK — flagging always works to protect people
	validSeverity := map[string]bool{
		"low": true, "medium": true, "high": true, "critical": true,
	}
	if !validSeverity[msg.Severity] {
		return nil, fmt.Errorf("severity must be low, medium, high, or critical")
	}
	if msg.Evidence == "" {
		return nil, fmt.Errorf("evidence required to flag a token")
	}
	if msg.Reason == "" {
		return nil, fmt.Errorf("reason required to flag a token")
	}

	// Rate limit — max 5 flags per address per day
	rateLimitKey := fmt.Sprintf("antirug-flag-count-%s-%d", msg.Creator, ctx.BlockHeight()/14400)
	countBytes, _ := store.Get([]byte(rateLimitKey))
	count := 0
	if countBytes != nil {
		fmt.Sscanf(string(countBytes), "%d", &count)
	}
	if count >= 5 {
		return nil, fmt.Errorf("maximum 5 token flags per day — prevents flag spam attacks")
	}
	store.Set([]byte(rateLimitKey), []byte(fmt.Sprintf("%d", count+1)))

	flagId := fmt.Sprintf("flag-%d-%s", ctx.BlockHeight(), msg.TokenId)
	store.Set([]byte(fmt.Sprintf("antirug-flagged-%s", msg.TokenId)),
		[]byte(fmt.Sprintf("%s|%s|%s|%s|%d|open",
			flagId, msg.Reason, msg.Evidence, msg.Severity, ctx.BlockHeight())))

	ctx.EventManager().EmitEvent(sdk.NewEvent("token_flagged",
		sdk.NewAttribute("flag_id", flagId),
		sdk.NewAttribute("token_id", msg.TokenId),
		sdk.NewAttribute("severity", msg.Severity),
		sdk.NewAttribute("reason", msg.Reason),
		sdk.NewAttribute("flagged_by", msg.Creator),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgFlagTokenResponse{FlagId: flagId, TokenId: msg.TokenId, Status: "flagged"}, nil
}
