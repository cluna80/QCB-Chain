package types

const (
	// ModuleName defines the module name
	ModuleName = "qcbmarket"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_qcbmarket"
)

var (
	ParamsKey = []byte("p_qcbmarket")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
