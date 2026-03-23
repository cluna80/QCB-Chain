package keeper

import (
	"context"
	"fmt"
	"qcb/x/qcbbridge/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SendIbcAgent(goCtx context.Context, msg *types.MsgSendIbcAgent) (*types.MsgSendIbcAgentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	validChains := map[string]bool{
		"xrpl-evm": true, "flare": true, "ethereum": true,
		"cosmos": true, "solana": true, "osmosis": true,
	}
	if !validChains[msg.DestChain] {
		return nil, fmt.Errorf("destChain %s not supported — use xrpl-evm, flare, ethereum, cosmos, solana, or osmosis", msg.DestChain)
	}
	if msg.DestAddr == "" {
		return nil, fmt.Errorf("destAddr cannot be empty")
	}
	packetId := fmt.Sprintf("ibc-pkt-%d-%s", ctx.BlockHeight(), msg.AgentId)
	ctx.EventManager().EmitEvent(sdk.NewEvent("ibc_agent_sent",
		sdk.NewAttribute("packet_id", packetId),
		sdk.NewAttribute("agent_id", msg.AgentId),
		sdk.NewAttribute("dest_chain", msg.DestChain),
		sdk.NewAttribute("dest_addr", msg.DestAddr),
		sdk.NewAttribute("sender", msg.Creator),
		sdk.NewAttribute("block", fmt.Sprintf("%d", ctx.BlockHeight())),
	))
	return &types.MsgSendIbcAgentResponse{PacketId: packetId, DestChain: msg.DestChain, Status: "sent"}, nil
}
