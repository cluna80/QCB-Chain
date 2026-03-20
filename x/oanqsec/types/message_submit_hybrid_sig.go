package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSubmitHybridSig{}

func NewMsgSubmitHybridSig(creator string, txHash string, classicalSig string, qsSig string, keyId string) *MsgSubmitHybridSig {
	return &MsgSubmitHybridSig{
		Creator:      creator,
		TxHash:       txHash,
		ClassicalSig: classicalSig,
		QsSig:        qsSig,
		KeyId:        keyId,
	}
}

func (msg *MsgSubmitHybridSig) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
