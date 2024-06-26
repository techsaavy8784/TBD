package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateOwner = "update_owner"

var _ sdk.Msg = &MsgUpdateOwner{}

func NewMsgUpdateOwner(manager string, denom string, newManager string) *MsgUpdateOwner {
	return &MsgUpdateOwner{
		Manager:    manager,
		Denom:      denom,
		NewManager: newManager,
	}
}

func (msg *MsgUpdateOwner) Route() string {
	return RouterKey
}

func (msg *MsgUpdateOwner) Type() string {
	return TypeMsgUpdateOwner
}

func (msg *MsgUpdateOwner) GetSigners() []sdk.AccAddress {
	manager, err := sdk.AccAddressFromBech32(msg.Manager)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{manager}
}

func (msg *MsgUpdateOwner) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateOwner) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Manager)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid manager address (%s)", err)
	}
	return nil
}
