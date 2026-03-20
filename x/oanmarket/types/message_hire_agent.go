package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgHireAgent{}

func NewMsgHireAgent(creator string, listingId string, taskDescription string, budget uint64) *MsgHireAgent {
	return &MsgHireAgent{
		Creator:         creator,
		ListingId:       listingId,
		TaskDescription: taskDescription,
		Budget:          budget,
	}
}

func (msg *MsgHireAgent) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
