package keeper

import (
	"context"

	"tbd/x/tokenfactory/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) BurnTokens(goCtx context.Context, msg *types.MsgBurnTokens) (*types.MsgBurnTokensResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	valFound, isFound := k.GetDenom(ctx, msg.Denom)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "denom not found")
	}

	if valFound.Manager != msg.Manager {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	Manager, err := sdk.AccAddressFromBech32(msg.Manager)
	if err != nil {
		return nil, err
	}

	burnCoins, err := sdk.ParseCoinsNormalized(msg.Amount)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "failed to parse coins")
	}

	// Send the coins from the creator to the module and later it burns
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, Manager, types.ModuleName, burnCoins); err != nil {
		return nil, sdkerrors.Wrap(err, "failed to send coins from account to module")
	}

	if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, burnCoins); err != nil {
		return nil, sdkerrors.Wrap(err, "failed to burn coins from module")
	}

	balance := k.bankKeeper.GetBalance(ctx, Manager, msg.Denom)
	var denom = types.Denom{
		Denom:       valFound.Denom,
		Description: valFound.Description,
		Ticker:      valFound.Ticker,
		Precision:   valFound.Precision,
		Supply:      balance.Amount.String(),
		Manager:     valFound.Manager,
	}

	k.SetDenom(ctx, denom)

	return &types.MsgBurnTokensResponse{}, nil
}
