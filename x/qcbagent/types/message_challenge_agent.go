package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgChallengeAgent{}

func NewMsgChallengeAgent(creator string, targetId string, stake uint64) *MsgChallengeAgent {
	return &MsgChallengeAgent{
		Creator:  creator,
		TargetId: targetId,
		Stake:    stake,
	}
}

func (msg *MsgChallengeAgent) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
