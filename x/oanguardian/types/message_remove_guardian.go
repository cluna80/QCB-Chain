package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRemoveGuardian{}

func NewMsgRemoveGuardian(creator string, guardianAddr string, reason string) *MsgRemoveGuardian {
	return &MsgRemoveGuardian{
		Creator:      creator,
		GuardianAddr: guardianAddr,
		Reason:       reason,
	}
}

func (msg *MsgRemoveGuardian) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
