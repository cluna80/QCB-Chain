package types

const (
	// ModuleName defines the module name
	ModuleName = "oancompute"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_oancompute"
)

var (
	ParamsKey = []byte("p_oancompute")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
