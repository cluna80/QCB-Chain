package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"oan/x/oancomms/types"
)

func (k msgServer) SetMsgPolicy(goCtx context.Context, msg *types.MsgSetMsgPolicy) (*types.MsgSetMsgPolicyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := k.storeService.OpenKVStore(ctx)

	if msg.MaxInbound == 0 {
		return nil, fmt.Errorf("maxInbound must be greater than 0")
	}

	policyId := fmt.Sprintf("policy-%s", msg.Creator)
	store.Set([]byte(fmt.Sprintf("msgpolicy-%s", msg.Creator)),
		[]byte(fmt.Sprintf("%s|%s|%d|%v",
			msg.AllowList, msg.DenyList, msg.MaxInbound, msg.RequirePqSig)))

	// Write deny entries for quick lookup
	if msg.DenyList != "" {
		store.Set([]byte(fmt.Sprintf("msgpolicy-denylist-%s", msg.Creator)),
			[]byte(msg.DenyList))
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent("msg_policy_set",
		sdk.NewAttribute("policy_id", policyId),
		sdk.NewAttribute("owner", msg.Creator),
		sdk.NewAttribute("max_inbound", fmt.Sprintf("%d", msg.MaxInbound)),
		sdk.NewAttribute("pq_required", fmt.Sprintf("%v", msg.RequirePqSig)),
	))
	return &types.MsgSetMsgPolicyResponse{PolicyId: policyId, Active: true}, nil
}
