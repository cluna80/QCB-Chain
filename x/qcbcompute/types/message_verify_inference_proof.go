package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgVerifyInferenceProof{}

func NewMsgVerifyInferenceProof(creator string, jobId string, proofHash string, proofType string) *MsgVerifyInferenceProof {
	return &MsgVerifyInferenceProof{
		Creator:   creator,
		JobId:     jobId,
		ProofHash: proofHash,
		ProofType: proofType,
	}
}

func (msg *MsgVerifyInferenceProof) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
