package types

const (
	// ModuleName defines the module name
	ModuleName = "oaneconomy"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_oaneconomy"
)

var (
	ParamsKey = []byte("p_oaneconomy")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
