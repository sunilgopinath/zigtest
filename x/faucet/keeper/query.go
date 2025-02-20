package keeper

import (
	"zigtest/x/faucet/types"
)

var _ types.QueryServer = Keeper{}
