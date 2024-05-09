package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"tbd/x/tokenfactory/types"
)

func (k msgServer) UpdateOwner(goCtx context.Context, msg *types.MsgUpdateOwner) (*types.MsgUpdateOwnerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	valFound, isFound := k.GetDenom(ctx, msg.Denom)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "denom does not exist")
	}

	// Checks if the the msg owner is the same as the current owner
	if msg.Manager != valFound.Manager {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect manager")
	}

	var denom = types.Denom{
		Manager:     msg.NewManager,
		Denom:       msg.Denom,
		Description: valFound.Description,
		Supply:      valFound.Supply,
		Precision:   valFound.Precision,
		Ticker:      valFound.Ticker,
	}

	k.SetDenom(ctx, denom)
	return &types.MsgUpdateOwnerResponse{}, nil
}
