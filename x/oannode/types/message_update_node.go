package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateNode{}

func NewMsgUpdateNode(creator string, nodeId string, uptimeProof string, blockHeight int32) *MsgUpdateNode {
	return &MsgUpdateNode{
		Creator:     creator,
		NodeId:      nodeId,
		UptimeProof: uptimeProof,
		BlockHeight: blockHeight,
	}
}

func (msg *MsgUpdateNode) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
