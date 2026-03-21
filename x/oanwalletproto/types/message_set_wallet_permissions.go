package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetWalletPermissions{}

func NewMsgSetWalletPermissions(creator string, walletId string, allowMsgs bool, allowAiAgent bool, allowBridge bool) *MsgSetWalletPermissions {
	return &MsgSetWalletPermissions{
		Creator:      creator,
		WalletId:     walletId,
		AllowMsgs:    allowMsgs,
		AllowAiAgent: allowAiAgent,
		AllowBridge:  allowBridge,
	}
}

func (msg *MsgSetWalletPermissions) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
