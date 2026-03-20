package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgBreedAgent{}

func NewMsgBreedAgent(creator string, parentA string, parentB string, childId string, childName string) *MsgBreedAgent {
	return &MsgBreedAgent{
		Creator:   creator,
		ParentA:   parentA,
		ParentB:   parentB,
		ChildId:   childId,
		ChildName: childName,
	}
}

func (msg *MsgBreedAgent) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
