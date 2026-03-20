package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSubmitProposal{}

func NewMsgSubmitProposal(creator string, title string, description string) *MsgSubmitProposal {
	return &MsgSubmitProposal{
		Creator:     creator,
		Title:       title,
		Description: description,
	}
}

func (msg *MsgSubmitProposal) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
