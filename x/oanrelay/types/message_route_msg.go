package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRouteMsg{}

func NewMsgRouteMsg(creator string, msgId string, fromAddr string, toAddr string, relayId string, payloadRef string) *MsgRouteMsg {
	return &MsgRouteMsg{
		Creator:    creator,
		MsgId:      msgId,
		FromAddr:   fromAddr,
		ToAddr:     toAddr,
		RelayId:    relayId,
		PayloadRef: payloadRef,
	}
}

func (msg *MsgRouteMsg) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
