package types

const (
	// ModuleName defines the module name
	ModuleName = "oanmarket"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_oanmarket"
)

var (
	ParamsKey = []byte("p_oanmarket")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
