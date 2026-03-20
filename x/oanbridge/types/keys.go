package types

const (
	// ModuleName defines the module name
	ModuleName = "oanbridge"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_oanbridge"
)

var (
	ParamsKey = []byte("p_oanbridge")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
