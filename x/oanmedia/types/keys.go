package types

const (
	// ModuleName defines the module name
	ModuleName = "oanmedia"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_oanmedia"
)

var (
	ParamsKey = []byte("p_oanmedia")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
