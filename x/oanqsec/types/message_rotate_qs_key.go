package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRotateQsKey{}

func NewMsgRotateQsKey(creator string, oldKeyId string, newPublicKeyHash string, reason string) *MsgRotateQsKey {
	return &MsgRotateQsKey{
		Creator:          creator,
		OldKeyId:         oldKeyId,
		NewPublicKeyHash: newPublicKeyHash,
		Reason:           reason,
	}
}

func (msg *MsgRotateQsKey) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
