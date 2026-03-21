package types

import (
	"testing"

	"oan/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgUnfreezeToken_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUnfreezeToken
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUnfreezeToken{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUnfreezeToken{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
