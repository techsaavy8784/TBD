package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgBurnTokens = "burn_tokens"

var _ sdk.Msg = &MsgBurnTokens{}

func NewMsgBurnTokens(manager string, denom string, amount string) *MsgBurnTokens {
	return &MsgBurnTokens{
		Manager: manager,
		Denom:   denom,
		Amount:  amount,
	}
}

func (msg *MsgBurnTokens) Route() string {
	return RouterKey
}

func (msg *MsgBurnTokens) Type() string {
	return TypeMsgBurnTokens
}

func (msg *MsgBurnTokens) GetSigners() []sdk.AccAddress {
	manager, err := sdk.AccAddressFromBech32(msg.Manager)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{manager}
}

func (msg *MsgBurnTokens) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBurnTokens) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Manager)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid manager address (%s)", err)
	}
	return nil
}
