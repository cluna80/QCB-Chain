package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateRelayRegion{}

func NewMsgUpdateRelayRegion(creator string, relayId string, newRegion string, endpoint string) *MsgUpdateRelayRegion {
	return &MsgUpdateRelayRegion{
		Creator:   creator,
		RelayId:   relayId,
		NewRegion: newRegion,
		Endpoint:  endpoint,
	}
}

func (msg *MsgUpdateRelayRegion) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
