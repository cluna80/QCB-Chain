package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgClaimNodeReward{}

func NewMsgClaimNodeReward(creator string, nodeId string, epoch int32) *MsgClaimNodeReward {
	return &MsgClaimNodeReward{
		Creator: creator,
		NodeId:  nodeId,
		Epoch:   epoch,
	}
}

func (msg *MsgClaimNodeReward) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
