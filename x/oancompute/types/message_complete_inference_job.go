package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCompleteInferenceJob{}

func NewMsgCompleteInferenceJob(creator string, jobId string, outputHash string, proof string) *MsgCompleteInferenceJob {
	return &MsgCompleteInferenceJob{
		Creator:    creator,
		JobId:      jobId,
		OutputHash: outputHash,
		Proof:      proof,
	}
}

func (msg *MsgCompleteInferenceJob) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
