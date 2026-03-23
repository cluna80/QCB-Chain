package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRecordTrade{}

func NewMsgRecordTrade(creator string, agentId string, action string, amount int32, result string) *MsgRecordTrade {
	return &MsgRecordTrade{
		Creator: creator,
		AgentId: agentId,
		Action:  action,
		Amount:  amount,
		Result:  result,
	}
}

func (msg *MsgRecordTrade) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
