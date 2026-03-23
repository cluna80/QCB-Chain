package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgEmergencyCryptoSwap{}

func NewMsgEmergencyCryptoSwap(creator string, fromAlgorithm string, toAlgorithm string, evidence string, urgency string) *MsgEmergencyCryptoSwap {
	return &MsgEmergencyCryptoSwap{
		Creator:       creator,
		FromAlgorithm: fromAlgorithm,
		ToAlgorithm:   toAlgorithm,
		Evidence:      evidence,
		Urgency:       urgency,
	}
}

func (msg *MsgEmergencyCryptoSwap) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
