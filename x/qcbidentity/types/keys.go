package types

const (
	// ModuleName defines the module name
	ModuleName = "qcbidentity"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_qcbidentity"
)

var (
	ParamsKey = []byte("p_qcbidentity")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
