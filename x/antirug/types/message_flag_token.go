package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgFlagToken{}

func NewMsgFlagToken(creator string, tokenId string, reason string, evidence string, severity string) *MsgFlagToken {
	return &MsgFlagToken{
		Creator:  creator,
		TokenId:  tokenId,
		Reason:   reason,
		Evidence: evidence,
		Severity: severity,
	}
}

func (msg *MsgFlagToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
