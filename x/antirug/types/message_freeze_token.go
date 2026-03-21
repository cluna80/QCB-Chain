package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgFreezeToken{}

func NewMsgFreezeToken(creator string, tokenId string, reason string, evidence string) *MsgFreezeToken {
	return &MsgFreezeToken{
		Creator:  creator,
		TokenId:  tokenId,
		Reason:   reason,
		Evidence: evidence,
	}
}

func (msg *MsgFreezeToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
