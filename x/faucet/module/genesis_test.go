package faucet_test

import (
	"testing"

	keepertest "zigtest/testutil/keeper"
	"zigtest/testutil/nullify"
	faucet "zigtest/x/faucet/module"
	"zigtest/x/faucet/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.FaucetKeeper(t)
	faucet.InitGenesis(ctx, k, genesisState)
	got := faucet.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
