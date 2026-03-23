package types

const (
	// ModuleName defines the module name
	ModuleName = "oanprotocol"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_oanprotocol"
)

var (
	ParamsKey = []byte("p_oanprotocol")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
