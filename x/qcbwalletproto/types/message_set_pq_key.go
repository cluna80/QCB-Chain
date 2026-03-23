package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetPqKey{}

func NewMsgSetPqKey(creator string, walletId string, pqKeyHash string, algorithm string) *MsgSetPqKey {
	return &MsgSetPqKey{
		Creator:   creator,
		WalletId:  walletId,
		PqKeyHash: pqKeyHash,
		Algorithm: algorithm,
	}
}

func (msg *MsgSetPqKey) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
