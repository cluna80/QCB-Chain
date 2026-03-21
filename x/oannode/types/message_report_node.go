package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgReportNode{}

func NewMsgReportNode(creator string, nodeId string, evidence string, violationType string) *MsgReportNode {
	return &MsgReportNode{
		Creator:       creator,
		NodeId:        nodeId,
		Evidence:      evidence,
		ViolationType: violationType,
	}
}

func (msg *MsgReportNode) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
