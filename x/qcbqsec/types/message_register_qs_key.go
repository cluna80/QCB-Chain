package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRegisterQsKey{}

func NewMsgRegisterQsKey(creator string, walletAddr string, keyType string, publicKeyHash string, algorithm string) *MsgRegisterQsKey {
	return &MsgRegisterQsKey{
		Creator:       creator,
		WalletAddr:    walletAddr,
		KeyType:       keyType,
		PublicKeyHash: publicKeyHash,
		Algorithm:     algorithm,
	}
}

func (msg *MsgRegisterQsKey) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
