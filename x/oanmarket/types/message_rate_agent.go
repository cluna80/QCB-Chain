package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRateAgent{}

func NewMsgRateAgent(creator string, agentId string, contractId string, rating uint64, review string) *MsgRateAgent {
	return &MsgRateAgent{
		Creator:    creator,
		AgentId:    agentId,
		ContractId: contractId,
		Rating:     rating,
		Review:     review,
	}
}

func (msg *MsgRateAgent) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
