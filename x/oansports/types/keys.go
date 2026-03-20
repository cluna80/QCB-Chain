package types

const (
	// ModuleName defines the module name
	ModuleName = "oansports"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_oansports"
)

var (
	ParamsKey = []byte("p_oansports")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
