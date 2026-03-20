package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgVetoProposal{}

func NewMsgVetoProposal(creator string, proposalId string, reason string) *MsgVetoProposal {
	return &MsgVetoProposal{
		Creator:    creator,
		ProposalId: proposalId,
		Reason:     reason,
	}
}

func (msg *MsgVetoProposal) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
