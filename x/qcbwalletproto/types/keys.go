package types

const (
	// ModuleName defines the module name
	ModuleName = "qcbwalletproto"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_qcbwalletproto"
)

var (
	ParamsKey = []byte("p_qcbwalletproto")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
