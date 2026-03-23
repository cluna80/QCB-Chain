package keeper

import (
	"context"
	"fmt"
	"qcb/x/qcbmedia/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) MintMusicNft(goCtx context.Context, msg *types.MsgMintMusicNft) (*types.MsgMintMusicNftResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if msg.AudioHash == "" {
		return nil, fmt.Errorf("audioHash cannot be empty")
	}
	if msg.Bpm == 0 || msg.Bpm > 300 {
		return nil, fmt.Errorf("BPM must be between 1 and 300")
	}
	if msg.Genre == "" {
		return nil, fmt.Errorf("genre cannot be empty")
	}
	nftId := fmt.Sprintf("music-%d-%s", ctx.BlockHeight(), msg.AgentId)
	store := k.storeService.OpenKVStore(ctx)
	store.Set([]byte(fmt.Sprintf("nft-%s", nftId)),
		[]byte(fmt.Sprintf("%s|%s|music|%s|%d|%s|active", nftId, msg.Title, msg.AgentId, msg.Bpm, msg.Genre)))
	store.Set([]byte(fmt.Sprintf("nft-owner-%s", nftId)), []byte(msg.Creator))
	ctx.EventManager().EmitEvent(sdk.NewEvent("music_nft_minted",
		sdk.NewAttribute("nft_id", nftId),
		sdk.NewAttribute("title", msg.Title),
		sdk.NewAttribute("agent_id", msg.AgentId),
		sdk.NewAttribute("audio_hash", msg.AudioHash),
		sdk.NewAttribute("bpm", fmt.Sprintf("%d", msg.Bpm)),
		sdk.NewAttribute("genre", msg.Genre),
		sdk.NewAttribute("creator", msg.Creator),
	))
	return &types.MsgMintMusicNftResponse{NftId: nftId, AudioHash: msg.AudioHash}, nil
}
