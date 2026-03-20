package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgVerifyIdentity{}

func NewMsgVerifyIdentity(creator string, did string, proofType string, proofData string) *MsgVerifyIdentity {
	return &MsgVerifyIdentity{
		Creator:   creator,
		Did:       did,
		ProofType: proofType,
		ProofData: proofData,
	}
}

func (msg *MsgVerifyIdentity) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
