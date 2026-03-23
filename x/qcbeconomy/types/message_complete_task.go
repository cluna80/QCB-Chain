package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCompleteTask{}

func NewMsgCompleteTask(creator string, taskId string, resultHash string) *MsgCompleteTask {
	return &MsgCompleteTask{
		Creator:    creator,
		TaskId:     taskId,
		ResultHash: resultHash,
	}
}

func (msg *MsgCompleteTask) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
