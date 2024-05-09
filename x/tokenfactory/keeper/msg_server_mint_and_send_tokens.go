package keeper

import (
	"context"

	"tbd/x/tokenfactory/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) MintAndSendTokens(goCtx context.Context, msg *types.MsgMintAndSendTokens) (*types.MsgMintAndSendTokensResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valfound, isFound := k.GetDenom(ctx, msg.Denom)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "denom does not exist")
	}

	// Checks if the the msg manager is the same as the current manager
	if msg.Manager != valfound.Manager {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect manager")
	}

	moduleAcct := k.accountKeeper.GetModuleAddress(types.ModuleName)
	recipientAddress, err := sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid recipient address")
	}

	mintCoins, err := sdk.ParseCoinsNormalized(msg.Supply)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "failed to parse coins")
	}
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, mintCoins); err != nil {
		return nil, sdkerrors.Wrap(err, "failed to mint coins")
	}
	if err := k.bankKeeper.SendCoins(ctx, moduleAcct, recipientAddress, mintCoins); err != nil {
		return nil, sdkerrors.Wrap(err, "failed to send coins")
	}

	balance := k.bankKeeper.GetBalance(ctx, recipientAddress, msg.Denom)

	var denom = types.Denom{
		Denom:       valfound.Denom,
		Description: valfound.Description,
		Ticker:      valfound.Ticker,
		Precision:   valfound.Precision,
		Supply:      balance.Amount.String(),
		Manager:     valfound.Manager,
	}

	k.SetDenom(ctx, denom)
	return &types.MsgMintAndSendTokensResponse{}, nil
}
