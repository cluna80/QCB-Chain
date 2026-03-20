package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRetireAgent{}

func NewMsgRetireAgent(creator string, nodeId string, reason string) *MsgRetireAgent {
	return &MsgRetireAgent{
		Creator: creator,
		NodeId:  nodeId,
		Reason:  reason,
	}
}

func (msg *MsgRetireAgent) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
