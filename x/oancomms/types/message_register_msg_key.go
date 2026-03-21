package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRegisterMsgKey{}

func NewMsgRegisterMsgKey(creator string, keyId string, keyType string, publicKeyHash string, algorithm string) *MsgRegisterMsgKey {
	return &MsgRegisterMsgKey{
		Creator:       creator,
		KeyId:         keyId,
		KeyType:       keyType,
		PublicKeyHash: publicKeyHash,
		Algorithm:     algorithm,
	}
}

func (msg *MsgRegisterMsgKey) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
