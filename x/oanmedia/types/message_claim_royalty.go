package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgClaimRoyalty{}

func NewMsgClaimRoyalty(creator string, nftId string, periodStart int32, periodEnd int32) *MsgClaimRoyalty {
	return &MsgClaimRoyalty{
		Creator:     creator,
		NftId:       nftId,
		PeriodStart: periodStart,
		PeriodEnd:   periodEnd,
	}
}

func (msg *MsgClaimRoyalty) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
