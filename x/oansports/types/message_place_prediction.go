package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgPlacePrediction{}

func NewMsgPlacePrediction(creator string, matchId string, prediction string, stake uint64) *MsgPlacePrediction {
	return &MsgPlacePrediction{
		Creator:    creator,
		MatchId:    matchId,
		Prediction: prediction,
		Stake:      stake,
	}
}

func (msg *MsgPlacePrediction) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
