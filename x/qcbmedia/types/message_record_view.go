package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRecordView{}

func NewMsgRecordView(creator string, nftId string, viewerAddr string, paymentAmount uint64) *MsgRecordView {
	return &MsgRecordView{
		Creator:       creator,
		NftId:         nftId,
		ViewerAddr:    viewerAddr,
		PaymentAmount: paymentAmount,
	}
}

func (msg *MsgRecordView) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
