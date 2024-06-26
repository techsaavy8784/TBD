package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateDenom{}, "tokenfactory/CreateDenom", nil)
	cdc.RegisterConcrete(&MsgUpdateDenom{}, "tokenfactory/UpdateDenom", nil)
	cdc.RegisterConcrete(&MsgMintAndSendTokens{}, "tokenfactory/MintAndSendTokens", nil)
	cdc.RegisterConcrete(&MsgUpdateOwner{}, "tokenfactory/UpdateOwner", nil)
	cdc.RegisterConcrete(&MsgBurnTokens{}, "tokenfactory/BurnTokens", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateDenom{},
		&MsgUpdateDenom{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgMintAndSendTokens{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateOwner{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBurnTokens{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
