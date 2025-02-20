package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
    tests := []struct {
        name string
        gs   GenesisState
        err  error
    }{
        {
            name: "valid_genesis_state",
            gs: GenesisState{
                // Use the Params type from the same package
                Params: Params{
                    MaxPerRequest: 100_000_000,  // 100 tokens
                    MaxPerAddress: 500_000_000,  // 500 tokens
                },
            },
            err: nil,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := tt.gs.Validate()
            if tt.err != nil {
                require.ErrorIs(t, err, tt.err)
                return
            }
            require.NoError(t, err)
        })
    }
}