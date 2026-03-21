package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRegisterRelay{}

func NewMsgRegisterRelay(creator string, relayId string, endpoint string, region string, pubKeyHash string) *MsgRegisterRelay {
	return &MsgRegisterRelay{
		Creator:    creator,
		RelayId:    relayId,
		Endpoint:   endpoint,
		Region:     region,
		PubKeyHash: pubKeyHash,
	}
}

func (msg *MsgRegisterRelay) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
