package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRegisterWalletProfile{}

func NewMsgRegisterWalletProfile(creator string, walletId string, did string, displayName string, avatarHash string) *MsgRegisterWalletProfile {
	return &MsgRegisterWalletProfile{
		Creator:     creator,
		WalletId:    walletId,
		Did:         did,
		DisplayName: displayName,
		AvatarHash:  avatarHash,
	}
}

func (msg *MsgRegisterWalletProfile) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
