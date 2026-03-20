package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgBroadcastAgentState{}

func NewMsgBroadcastAgentState(creator string, agentId string, stateHash string, targetChains string) *MsgBroadcastAgentState {
	return &MsgBroadcastAgentState{
		Creator:      creator,
		AgentId:      agentId,
		StateHash:    stateHash,
		TargetChains: targetChains,
	}
}

func (msg *MsgBroadcastAgentState) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
