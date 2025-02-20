package types

const (
	// ModuleName defines the module name
	ModuleName = "faucet"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_faucet"

	// TotalRequestKey is used to track total amount requested per address
    TotalRequestKey = "TotalRequest/value/"
)

var (
	ParamsKey = []byte("p_faucet")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
