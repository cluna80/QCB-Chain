package types

const (
	// ModuleName defines the module name
	ModuleName = "qcbeconomy"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_qcbeconomy"
)

var (
	ParamsKey = []byte("p_qcbeconomy")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
