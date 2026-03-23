package types

const (
	// ModuleName defines the module name
	ModuleName = "qcbguardian"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_qcbguardian"
)

var (
	ParamsKey = []byte("p_qcbguardian")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
