package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateMediaNft{}

func NewMsgCreateMediaNft(creator string, title string, mediaType string, contentHash string, creatorShare uint64) *MsgCreateMediaNft {
	return &MsgCreateMediaNft{
		Creator:      creator,
		Title:        title,
		MediaType:    mediaType,
		ContentHash:  contentHash,
		CreatorShare: creatorShare,
	}
}

func (msg *MsgCreateMediaNft) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
