package types

const (
	// ModuleName defines the module name
	ModuleName = "qcbmedia"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_qcbmedia"
)

var (
	ParamsKey = []byte("p_qcbmedia")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
