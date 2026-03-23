package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgScheduleMatch{}

func NewMsgScheduleMatch(creator string, matchId string, athleteA string, athleteB string, stadiumId string, scheduledAt int32) *MsgScheduleMatch {
	return &MsgScheduleMatch{
		Creator:     creator,
		MatchId:     matchId,
		AthleteA:    athleteA,
		AthleteB:    athleteB,
		StadiumId:   stadiumId,
		ScheduledAt: scheduledAt,
	}
}

func (msg *MsgScheduleMatch) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
