package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgStakeTokens{}

func NewMsgStakeTokens(creator string, amount uint64, lockPeriod int32) *MsgStakeTokens {
	return &MsgStakeTokens{
		Creator:    creator,
		Amount:     amount,
		LockPeriod: lockPeriod,
	}
}

func (msg *MsgStakeTokens) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
