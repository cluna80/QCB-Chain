package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgClaimPredictionReward{}

func NewMsgClaimPredictionReward(creator string, predictionId string) *MsgClaimPredictionReward {
	return &MsgClaimPredictionReward{
		Creator:      creator,
		PredictionId: predictionId,
	}
}

func (msg *MsgClaimPredictionReward) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
