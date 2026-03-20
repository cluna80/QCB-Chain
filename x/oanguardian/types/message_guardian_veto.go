package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgGuardianVeto{}

func NewMsgGuardianVeto(creator string, jobId string, reason string, severity string) *MsgGuardianVeto {
	return &MsgGuardianVeto{
		Creator:  creator,
		JobId:    jobId,
		Reason:   reason,
		Severity: severity,
	}
}

func (msg *MsgGuardianVeto) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
