package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRegisterAthlete{}

func NewMsgRegisterAthlete(creator string, athleteId string, agentId string, sport string, position string) *MsgRegisterAthlete {
	return &MsgRegisterAthlete{
		Creator:   creator,
		AthleteId: athleteId,
		AgentId:   agentId,
		Sport:     sport,
		Position:  position,
	}
}

func (msg *MsgRegisterAthlete) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
