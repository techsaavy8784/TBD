package keeper

import (
	"context"

	"tbd/x/tokenfactory/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateDenom(goCtx context.Context, msg *types.MsgCreateDenom) (*types.MsgCreateDenomResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetDenom(
		ctx,
		msg.Denom,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "denom already set")
	}

	var denom = types.Denom{
		Manager:     msg.Manager,
		Denom:       msg.Denom,
		Description: msg.Description,
		Ticker:      msg.Ticker,
		Precision:   msg.Precision,
	}

	k.SetDenom(
		ctx,
		denom,
	)
	return &types.MsgCreateDenomResponse{}, nil
}

func (k msgServer) UpdateDenom(goCtx context.Context, msg *types.MsgUpdateDenom) (*types.MsgUpdateDenomResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetDenom(
		ctx,
		msg.Denom,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg manager is the same as the current owner
	if msg.Manager != valFound.Manager {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect manager")
	}

	var denom = types.Denom{
		Manager:     msg.Manager,
		Denom:       msg.Denom,
		Description: msg.Description,
		Ticker:      msg.Ticker,
		Precision:   valFound.Precision,
		Supply:      valFound.Supply,
	}

	k.SetDenom(ctx, denom)

	return &types.MsgUpdateDenomResponse{}, nil
}
