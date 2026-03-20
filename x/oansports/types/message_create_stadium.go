package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateStadium{}

func NewMsgCreateStadium(creator string, stadiumId string, name string, capacity uint64, location string) *MsgCreateStadium {
	return &MsgCreateStadium{
		Creator:   creator,
		StadiumId: stadiumId,
		Name:      name,
		Capacity:  capacity,
		Location:  location,
	}
}

func (msg *MsgCreateStadium) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
