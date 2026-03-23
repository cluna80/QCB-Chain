package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSlashRelay{}

func NewMsgSlashRelay(creator string, relayId string, evidence string, slashType string) *MsgSlashRelay {
	return &MsgSlashRelay{
		Creator:   creator,
		RelayId:   relayId,
		Evidence:  evidence,
		SlashType: slashType,
	}
}

func (msg *MsgSlashRelay) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
