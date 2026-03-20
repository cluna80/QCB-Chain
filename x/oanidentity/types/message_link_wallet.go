package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgLinkWallet{}

func NewMsgLinkWallet(creator string, did string, walletAddress string) *MsgLinkWallet {
	return &MsgLinkWallet{
		Creator:       creator,
		Did:           did,
		WalletAddress: walletAddress,
	}
}

func (msg *MsgLinkWallet) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
