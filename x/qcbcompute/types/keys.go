package types

const (
	// ModuleName defines the module name
	ModuleName = "qcbcompute"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_qcbcompute"
)

var (
	ParamsKey = []byte("p_qcbcompute")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
