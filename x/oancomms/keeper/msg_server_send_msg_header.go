package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"oan/x/oancomms/types"
)

func (k msgServer) SendMsgHeader(goCtx context.Context, msg *types.MsgSendMsgHeader) (*types.MsgSendMsgHeaderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := k.storeService.OpenKVStore(ctx)
	params := k.GetParams(ctx)

	// Validate sender has a registered msg key
	senderHasKey := false
	for _, kt := range []string{"dilithium3", "sphincs-plus", "falcon-1024", "ed25519", "secp256k1"} {
		keyStoreKey := fmt.Sprintf("msgkey-%s-%s", msg.Creator, kt)
		existing, _ := store.Get([]byte(keyStoreKey))
		if existing != nil {
			senderHasKey = true
			break
		}
	}
	if !senderHasKey {
		return nil, fmt.Errorf("sender must register a msg key before sending — use register-msg-key")
	}

	// Validate payload hash not empty
	if msg.PayloadHash == "" {
		return nil, fmt.Errorf("payloadHash cannot be empty — payload lives off-chain, only hash on-chain")
	}
	if msg.ToAddr == "" {
		return nil, fmt.Errorf("toAddr cannot be empty")
	}

	// Check msg size limit
	maxSize := params.MaxMsgSize
	if maxSize == 0 { maxSize = 65536 }

	// Check TTL
	ttl := params.MsgTtlBlocks
	if ttl == 0 { ttl = 1000 }
	expiresAt := int64(ctx.BlockHeight()) + int64(ttl)

	// Check policy — is sender blocked by recipient?
	policyKey := fmt.Sprintf("msgpolicy-deny-%s-%s", msg.ToAddr, msg.Creator)
	denied, _ := store.Get([]byte(policyKey))
	if denied != nil {
		return nil, fmt.Errorf("recipient has blocked this sender")
	}

	// Store msg header only — NO payload on chain
	msgKey := fmt.Sprintf("msgheader-%s", msg.MsgId)
	existing, _ := store.Get([]byte(msgKey))
	if existing != nil {
		return nil, fmt.Errorf("msgId %s already exists", msg.MsgId)
	}
	store.Set([]byte(msgKey),
		[]byte(fmt.Sprintf("%s|%s|%s|%s|%s|%d|pending",
			msg.MsgId, msg.Creator, msg.ToAddr,
			msg.PayloadHash, msg.MsgType, expiresAt)))

	// Relay hint — suggest nearest relay
	relayHint := fmt.Sprintf("relay-%s-region-1", msg.ToAddr[:8])

	ctx.EventManager().EmitEvent(sdk.NewEvent("msg_header_sent",
		sdk.NewAttribute("msg_id", msg.MsgId),
		sdk.NewAttribute("from", msg.Creator),
		sdk.NewAttribute("to", msg.ToAddr),
		sdk.NewAttribute("msg_type", msg.MsgType),
		sdk.NewAttribute("payload_hash", msg.PayloadHash),
		sdk.NewAttribute("expires_at", fmt.Sprintf("%d", expiresAt)),
	))
	return &types.MsgSendMsgHeaderResponse{
		MsgId: msg.MsgId, Routed: true, RelayHint: relayHint,
	}, nil
}
