package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgListAgentForHire{}

func NewMsgListAgentForHire(creator string, agentId string, pricePerTask uint64, skills string) *MsgListAgentForHire {
	return &MsgListAgentForHire{
		Creator:      creator,
		AgentId:      agentId,
		PricePerTask: pricePerTask,
		Skills:       skills,
	}
}

func (msg *MsgListAgentForHire) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
