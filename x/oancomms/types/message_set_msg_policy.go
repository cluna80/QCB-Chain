package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetMsgPolicy{}

func NewMsgSetMsgPolicy(creator string, allowList string, denyList string, maxInbound uint64, requirePqSig bool) *MsgSetMsgPolicy {
	return &MsgSetMsgPolicy{
		Creator:      creator,
		AllowList:    allowList,
		DenyList:     denyList,
		MaxInbound:   maxInbound,
		RequirePqSig: requirePqSig,
	}
}

func (msg *MsgSetMsgPolicy) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
