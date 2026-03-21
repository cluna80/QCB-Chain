package types

const (
	// ModuleName defines the module name
	ModuleName = "oanwalletproto"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_oanwalletproto"
)

var (
	ParamsKey = []byte("p_oanwalletproto")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
