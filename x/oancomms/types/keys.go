package types

const (
	// ModuleName defines the module name
	ModuleName = "oancomms"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_oancomms"
)

var (
	ParamsKey = []byte("p_oancomms")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
