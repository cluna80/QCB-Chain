package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSlashNode{}

func NewMsgSlashNode(creator string, nodeId string, evidence string, slashType string) *MsgSlashNode {
	return &MsgSlashNode{
		Creator:   creator,
		NodeId:    nodeId,
		Evidence:  evidence,
		SlashType: slashType,
	}
}

func (msg *MsgSlashNode) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
