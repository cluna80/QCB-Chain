package types

const (
	// ModuleName defines the module name
	ModuleName = "qcbdao"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_qcbdao"
)

var (
	ParamsKey = []byte("p_qcbdao")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
