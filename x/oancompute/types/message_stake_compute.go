package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgStakeCompute{}

func NewMsgStakeCompute(creator string, gpuType string, capacity uint64, pricePerJob uint64) *MsgStakeCompute {
	return &MsgStakeCompute{
		Creator:     creator,
		GpuType:     gpuType,
		Capacity:    capacity,
		PricePerJob: pricePerJob,
	}
}

func (msg *MsgStakeCompute) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
