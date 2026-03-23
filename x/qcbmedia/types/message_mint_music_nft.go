package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgMintMusicNft{}

func NewMsgMintMusicNft(creator string, title string, audioHash string, agentId string, bpm uint64, genre string) *MsgMintMusicNft {
	return &MsgMintMusicNft{
		Creator:   creator,
		Title:     title,
		AudioHash: audioHash,
		AgentId:   agentId,
		Bpm:       bpm,
		Genre:     genre,
	}
}

func (msg *MsgMintMusicNft) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
