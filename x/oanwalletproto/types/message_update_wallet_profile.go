package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateWalletProfile{}

func NewMsgUpdateWalletProfile(creator string, walletId string, displayName string, avatarHash string) *MsgUpdateWalletProfile {
	return &MsgUpdateWalletProfile{
		Creator:     creator,
		WalletId:    walletId,
		DisplayName: displayName,
		AvatarHash:  avatarHash,
	}
}

func (msg *MsgUpdateWalletProfile) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
