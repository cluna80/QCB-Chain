package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSendIbcAgent{}

func NewMsgSendIbcAgent(creator string, agentId string, destChain string, destAddr string) *MsgSendIbcAgent {
	return &MsgSendIbcAgent{
		Creator:   creator,
		AgentId:   agentId,
		DestChain: destChain,
		DestAddr:  destAddr,
	}
}

func (msg *MsgSendIbcAgent) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
