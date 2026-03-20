package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgLiftPause{}

func NewMsgLiftPause(creator string, pauseId string, justification string) *MsgLiftPause {
	return &MsgLiftPause{
		Creator:       creator,
		PauseId:       pauseId,
		Justification: justification,
	}
}

func (msg *MsgLiftPause) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
