package keeper

import (
	"context"
	"zigtest/x/faucet/types"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math" // This is the correct import
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) Request(goCtx context.Context, msg *types.MsgRequest) (*types.MsgRequestResponse, error) {
    ctx := sdk.UnwrapSDKContext(goCtx)

    // Get module params
    params := k.GetParams(ctx)

    // Check if request amount is within maxPerRequest
    if msg.Amount > params.MaxPerRequest {
        return nil, errorsmod.Wrapf(types.ErrExceedsMaxRequest, "request amount %d exceeds limit %d", msg.Amount, params.MaxPerRequest)
    }

    // Get total requested for this address and check against maxPerAddress
    totalRequested := k.GetTotalRequested(ctx, msg.Creator)
    if totalRequested+msg.Amount > params.MaxPerAddress {
        return nil, errorsmod.Wrapf(types.ErrExceedsMaxAddress, "total requested %d would exceed limit %d", totalRequested+msg.Amount, params.MaxPerAddress)
    }

    // Parse creator address
    requester, err := sdk.AccAddressFromBech32(msg.Creator)
    if err != nil {
        return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, err.Error())
    }

    // Create coins using default bond denom
    amount := math.NewInt(int64(msg.Amount))
    coins := sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, amount))

    // Send coins from module to requester
    err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, requester, coins)
    if err != nil {
        return nil, errorsmod.Wrap(err, "failed to send coins")
    }

    // Update total requested amount
    k.AddToTotalRequested(ctx, msg.Creator, msg.Amount)

    return &types.MsgRequestResponse{}, nil
}