package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUnlockWallet{}

func NewMsgUnlockWallet(creator string, walletId string, proof string) *MsgUnlockWallet {
	return &MsgUnlockWallet{
		Creator:  creator,
		WalletId: walletId,
		Proof:    proof,
	}
}

func (msg *MsgUnlockWallet) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
