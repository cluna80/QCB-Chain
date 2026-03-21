package keeper

import (
	"context"
	"fmt"
	"oan/x/oannode/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ReportNode(goCtx context.Context, msg *types.MsgReportNode) (*types.MsgReportNodeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := k.storeService.OpenKVStore(ctx)

	if msg.Evidence == "" {
		return nil, fmt.Errorf("evidence required to report a node")
	}
	validViolations := map[string]bool{
		"offline": true, "double-sign": true, "bad-inference": true,
		"malicious": true, "sybil": true, "spam": true,
	}
	if !validViolations[msg.ViolationType] {
		return nil, fmt.Errorf("violationType must be offline, double-sign, bad-inference, malicious, sybil, or spam")
	}

	// Rate limit reports — max 3 reports per reporter per day
	reportCountKey := fmt.Sprintf("report-count-%s-%d", msg.Creator, ctx.BlockHeight()/14400)
	countBytes, _ := store.Get([]byte(reportCountKey))
	count := 0
	if countBytes != nil {
		fmt.Sscanf(string(countBytes), "%d", &count)
	}
	if count >= 3 {
		return nil, fmt.Errorf("maximum 3 reports per day — prevents report spam attacks")
	}
	store.Set([]byte(reportCountKey), []byte(fmt.Sprintf("%d", count+1)))

	reportId := fmt.Sprintf("report-%d-%s", ctx.BlockHeight(), msg.NodeId)
	store.Set([]byte(fmt.Sprintf("nodereport-%s", reportId)),
		[]byte(fmt.Sprintf("%s|%s|%s|%s|pending",
			reportId, msg.NodeId, msg.ViolationType, msg.Creator)))

	ctx.EventManager().EmitEvent(sdk.NewEvent("node_reported",
		sdk.NewAttribute("report_id", reportId),
		sdk.NewAttribute("node_id", msg.NodeId),
		sdk.NewAttribute("violation", msg.ViolationType),
		sdk.NewAttribute("reporter", msg.Creator),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgReportNodeResponse{ReportId: reportId, Status: "pending"}, nil
}
