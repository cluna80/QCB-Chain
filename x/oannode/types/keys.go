package types

const (
	// ModuleName defines the module name
	ModuleName = "oannode"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_oannode"
)

var (
	ParamsKey = []byte("p_oannode")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
