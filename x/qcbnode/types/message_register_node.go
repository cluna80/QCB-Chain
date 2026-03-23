package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRegisterNode{}

func NewMsgRegisterNode(creator string, nodeType string, endpoint string, nodeId string) *MsgRegisterNode {
	return &MsgRegisterNode{
		Creator:  creator,
		NodeType: nodeType,
		Endpoint: endpoint,
		NodeId:   nodeId,
	}
}

func (msg *MsgRegisterNode) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
