package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeregisterNode{}

func NewMsgDeregisterNode(creator string, nodeId string, reason string) *MsgDeregisterNode {
	return &MsgDeregisterNode{
		Creator: creator,
		NodeId:  nodeId,
		Reason:  reason,
	}
}

func (msg *MsgDeregisterNode) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
