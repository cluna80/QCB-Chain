package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDistributeRewards{}

func NewMsgDistributeRewards(creator string, epoch int32) *MsgDistributeRewards {
	return &MsgDistributeRewards{
		Creator: creator,
		Epoch:   epoch,
	}
}

func (msg *MsgDistributeRewards) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
