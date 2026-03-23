package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgAddGuardian{}

func NewMsgAddGuardian(creator string, guardianAddr string, displayName string, justification string) *MsgAddGuardian {
	return &MsgAddGuardian{
		Creator:       creator,
		GuardianAddr:  guardianAddr,
		DisplayName:   displayName,
		Justification: justification,
	}
}

func (msg *MsgAddGuardian) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
