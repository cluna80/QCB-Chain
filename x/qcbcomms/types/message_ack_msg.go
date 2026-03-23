package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgAckMsg{}

func NewMsgAckMsg(creator string, msgId string, ackType string) *MsgAckMsg {
	return &MsgAckMsg{
		Creator: creator,
		MsgId:   msgId,
		AckType: ackType,
	}
}

func (msg *MsgAckMsg) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
