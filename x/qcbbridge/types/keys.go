package types

const (
	// ModuleName defines the module name
	ModuleName = "qcbbridge"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_qcbbridge"
)

var (
	ParamsKey = []byte("p_qcbbridge")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
