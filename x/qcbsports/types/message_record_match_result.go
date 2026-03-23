package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRecordMatchResult{}

func NewMsgRecordMatchResult(creator string, matchId string, winner string, loser string, scoreA uint64, scoreB uint64, statsHash string) *MsgRecordMatchResult {
	return &MsgRecordMatchResult{
		Creator:   creator,
		MatchId:   matchId,
		Winner:    winner,
		Loser:     loser,
		ScoreA:    scoreA,
		ScoreB:    scoreB,
		StatsHash: statsHash,
	}
}

func (msg *MsgRecordMatchResult) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
