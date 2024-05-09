package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgMintAndSendTokens = "mint_and_send_tokens"

var _ sdk.Msg = &MsgMintAndSendTokens{}

func NewMsgMintAndSendTokens(manager string, denom string, supply string, recipient string) *MsgMintAndSendTokens {
	return &MsgMintAndSendTokens{
		Manager:   manager,
		Denom:     denom,
		Supply:    supply,
		Recipient: recipient,
	}
}

func (msg *MsgMintAndSendTokens) Route() string {
	return RouterKey
}

func (msg *MsgMintAndSendTokens) Type() string {
	return TypeMsgMintAndSendTokens
}

func (msg *MsgMintAndSendTokens) GetSigners() []sdk.AccAddress {
	manager, err := sdk.AccAddressFromBech32(msg.Manager)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{manager}
}

func (msg *MsgMintAndSendTokens) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMintAndSendTokens) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Manager)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid manager address (%s)", err)
	}
	return nil
}
