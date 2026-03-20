package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetThreatLevel{}

func NewMsgSetThreatLevel(creator string, level uint64, evidence string, justification string) *MsgSetThreatLevel {
	return &MsgSetThreatLevel{
		Creator:       creator,
		Level:         level,
		Evidence:      evidence,
		Justification: justification,
	}
}

func (msg *MsgSetThreatLevel) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
