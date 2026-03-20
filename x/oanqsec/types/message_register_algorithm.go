package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRegisterAlgorithm{}

func NewMsgRegisterAlgorithm(creator string, algorithmId string, algorithmType string, securityLevel uint64, specification string) *MsgRegisterAlgorithm {
	return &MsgRegisterAlgorithm{
		Creator:       creator,
		AlgorithmId:   algorithmId,
		AlgorithmType: algorithmType,
		SecurityLevel: securityLevel,
		Specification: specification,
	}
}

func (msg *MsgRegisterAlgorithm) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
