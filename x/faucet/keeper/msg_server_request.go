package keeper

import (
	"context"

	"zigtest/x/faucet/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Request(goCtx context.Context, msg *types.MsgRequest) (*types.MsgRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgRequestResponse{}, nil
}
