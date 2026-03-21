package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRegisterToken{}

func NewMsgRegisterToken(creator string, tokenId string, tokenName string, symbol string, maxSupply uint64, liquidityLockBlocks uint64) *MsgRegisterToken {
	return &MsgRegisterToken{
		Creator:             creator,
		TokenId:             tokenId,
		TokenName:           tokenName,
		Symbol:              symbol,
		MaxSupply:           maxSupply,
		LiquidityLockBlocks: liquidityLockBlocks,
	}
}

func (msg *MsgRegisterToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
