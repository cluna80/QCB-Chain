package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgTokenizeOutput{}

func NewMsgTokenizeOutput(creator string, agentId string, outputType string, contentHash string, value uint64) *MsgTokenizeOutput {
	return &MsgTokenizeOutput{
		Creator:     creator,
		AgentId:     agentId,
		OutputType:  outputType,
		ContentHash: contentHash,
		Value:       value,
	}
}

func (msg *MsgTokenizeOutput) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
