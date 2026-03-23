package types

const (
	// ModuleName defines the module name
	ModuleName = "qcbqsec"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_qcbqsec"
)

var (
	ParamsKey = []byte("p_qcbqsec")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
