package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSlashBadInference{}

func NewMsgSlashBadInference(creator string, jobId string, validatorAddr string, evidence string) *MsgSlashBadInference {
	return &MsgSlashBadInference{
		Creator:       creator,
		JobId:         jobId,
		ValidatorAddr: validatorAddr,
		Evidence:      evidence,
	}
}

func (msg *MsgSlashBadInference) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
