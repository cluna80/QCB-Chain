package types

const (
	// ModuleName defines the module name
	ModuleName = "qcbcomms"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_qcbcomms"
)

var (
	ParamsKey = []byte("p_qcbcomms")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
