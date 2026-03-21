package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRelayHeartbeat{}

func NewMsgRelayHeartbeat(creator string, relayId string, proofHash string, blocksOnline int32) *MsgRelayHeartbeat {
	return &MsgRelayHeartbeat{
		Creator:      creator,
		RelayId:      relayId,
		ProofHash:    proofHash,
		BlocksOnline: blocksOnline,
	}
}

func (msg *MsgRelayHeartbeat) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
