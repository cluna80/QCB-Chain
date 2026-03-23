package keeper

import (
	"context"
	"fmt"
	"qcb/x/antirug/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ApproveToken(goCtx context.Context, msg *types.MsgApproveToken) (*types.MsgApproveTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := k.storeService.OpenKVStore(ctx)
	params := k.GetParams(ctx)

	// DORMANCY CHECK
	if !params.Enabled {
		return nil, fmt.Errorf("antirug module not yet active — token approval optional until governance activates it")
	}

	// Guardian or DAO only
	guardianKey := fmt.Sprintf("guardian-auth-%s", msg.Creator)
	isGuardian, _ := store.Get([]byte(guardianKey))
	isDaoModule := msg.Creator == "oan10d07y265gmmuvt4z0w9aw880jnsr700j6zn9kn"
	if isGuardian == nil && !isDaoModule {
		return nil, fmt.Errorf("only guardians or DAO governance can approve tokens")
	}

	validVerdicts := map[string]bool{
		"approved": true, "rejected": true, "conditional": true,
	}
	if !validVerdicts[msg.Verdict] {
		return nil, fmt.Errorf("verdict must be approved, rejected, or conditional")
	}

	approved := msg.Verdict == "approved"
	safetyScore := uint64(0)
	if approved {
		safetyScore = 100
	}

	tokenKey := fmt.Sprintf("antirug-token-%s", msg.TokenId)
	store.Set([]byte(tokenKey+"-verdict"),
		[]byte(fmt.Sprintf("%s|%s|%d|%d",
			msg.Verdict, msg.Justification, ctx.BlockHeight(), safetyScore)))

	ctx.EventManager().EmitEvent(sdk.NewEvent("token_approval_decision",
		sdk.NewAttribute("token_id", msg.TokenId),
		sdk.NewAttribute("verdict", msg.Verdict),
		sdk.NewAttribute("safety_score", fmt.Sprintf("%d", safetyScore)),
		sdk.NewAttribute("justification", msg.Justification),
		sdk.NewAttribute("decided_by", msg.Creator),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgApproveTokenResponse{
		TokenId: msg.TokenId, Approved: approved, SafetyScore: safetyScore,
	}, nil
}
