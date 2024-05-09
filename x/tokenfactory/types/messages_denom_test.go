package types

import (
	"testing"

	"tbd/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateDenom_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateDenom
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateDenom{
				Manager: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateDenom{
				Manager: sample.AccAddress(),
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

func TestMsgUpdateDenom_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateDenom
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateDenom{
				Manager: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateDenom{
				Manager: sample.AccAddress(),
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
