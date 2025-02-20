package types

import (
	"testing"

	"zigtest/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgRequest_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgRequest
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgRequest{
				Creator: "invalid_address",
				Amount: 1,
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgRequest{
				Creator: sample.AccAddress(),
				Amount: 1,
			},
		},
		{
			name: "invalid amount",
			msg: MsgRequest{
				Creator: sample.AccAddress(),
				Amount: 0,
			},
			err: sdkerrors.ErrInvalidRequest,
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
