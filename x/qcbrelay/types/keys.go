package types

const (
	// ModuleName defines the module name
	ModuleName = "qcbrelay"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_qcbrelay"
)

var (
	ParamsKey = []byte("p_qcbrelay")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
