package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDelegateVote{}

func NewMsgDelegateVote(creator string, delegateTo string, power uint64) *MsgDelegateVote {
	return &MsgDelegateVote{
		Creator:    creator,
		DelegateTo: delegateTo,
		Power:      power,
	}
}

func (msg *MsgDelegateVote) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
