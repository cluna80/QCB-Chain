package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRequestUpgrade{}

func NewMsgRequestUpgrade(creator string, tokenId string, upgradeType string, description string) *MsgRequestUpgrade {
	return &MsgRequestUpgrade{
		Creator:     creator,
		TokenId:     tokenId,
		UpgradeType: upgradeType,
		Description: description,
	}
}

func (msg *MsgRequestUpgrade) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
