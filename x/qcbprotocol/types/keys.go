package types

const (
	// ModuleName defines the module name
	ModuleName = "qcbprotocol"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_qcbprotocol"
)

var (
	ParamsKey = []byte("p_qcbprotocol")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
