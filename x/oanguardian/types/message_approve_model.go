package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgApproveModel{}

func NewMsgApproveModel(creator string, modelId string, verdict string, justification string) *MsgApproveModel {
	return &MsgApproveModel{
		Creator:       creator,
		ModelId:       modelId,
		Verdict:       verdict,
		Justification: justification,
	}
}

func (msg *MsgApproveModel) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
