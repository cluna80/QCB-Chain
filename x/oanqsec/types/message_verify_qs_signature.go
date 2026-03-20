package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgVerifyQsSignature{}

func NewMsgVerifyQsSignature(creator string, txHash string, qsSig string, keyId string) *MsgVerifyQsSignature {
	return &MsgVerifyQsSignature{
		Creator: creator,
		TxHash:  txHash,
		QsSig:   qsSig,
		KeyId:   keyId,
	}
}

func (msg *MsgVerifyQsSignature) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
