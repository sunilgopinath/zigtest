package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "zigtest/testutil/keeper"
	"zigtest/x/faucet/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.FaucetKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
