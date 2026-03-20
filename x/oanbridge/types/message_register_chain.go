package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRegisterChain{}

func NewMsgRegisterChain(creator string, chainId string, chainName string, bridgeType string, endpoint string) *MsgRegisterChain {
	return &MsgRegisterChain{
		Creator:    creator,
		ChainId:    chainId,
		ChainName:  chainName,
		BridgeType: bridgeType,
		Endpoint:   endpoint,
	}
}

func (msg *MsgRegisterChain) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
