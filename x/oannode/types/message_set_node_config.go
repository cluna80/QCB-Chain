package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetNodeConfig{}

func NewMsgSetNodeConfig(creator string, nodeId string, endpoint string, capabilities string) *MsgSetNodeConfig {
	return &MsgSetNodeConfig{
		Creator:      creator,
		NodeId:       nodeId,
		Endpoint:     endpoint,
		Capabilities: capabilities,
	}
}

func (msg *MsgSetNodeConfig) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
