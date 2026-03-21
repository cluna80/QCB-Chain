package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgApproveToken{}

func NewMsgApproveToken(creator string, tokenId string, verdict string, justification string) *MsgApproveToken {
	return &MsgApproveToken{
		Creator:       creator,
		TokenId:       tokenId,
		Verdict:       verdict,
		Justification: justification,
	}
}

func (msg *MsgApproveToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
