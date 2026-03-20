package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateReputation{}

func NewMsgUpdateReputation(creator string, did string, delta int32, reason string) *MsgUpdateReputation {
	return &MsgUpdateReputation{
		Creator: creator,
		Did:     did,
		Delta:   delta,
		Reason:  reason,
	}
}

func (msg *MsgUpdateReputation) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
