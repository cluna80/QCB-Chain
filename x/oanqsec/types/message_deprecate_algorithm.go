package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeprecateAlgorithm{}

func NewMsgDeprecateAlgorithm(creator string, algorithmId string, reason string, replacementId string) *MsgDeprecateAlgorithm {
	return &MsgDeprecateAlgorithm{
		Creator:       creator,
		AlgorithmId:   algorithmId,
		Reason:        reason,
		ReplacementId: replacementId,
	}
}

func (msg *MsgDeprecateAlgorithm) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
