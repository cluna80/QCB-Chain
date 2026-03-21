package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSendMsgHeader{}

func NewMsgSendMsgHeader(creator string, toAddr string, msgId string, encKeyId string, payloadHash string, msgType string) *MsgSendMsgHeader {
	return &MsgSendMsgHeader{
		Creator:     creator,
		ToAddr:      toAddr,
		MsgId:       msgId,
		EncKeyId:    encKeyId,
		PayloadHash: payloadHash,
		MsgType:     msgType,
	}
}

func (msg *MsgSendMsgHeader) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
