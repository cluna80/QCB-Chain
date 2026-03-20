package types

const (
	// ModuleName defines the module name
	ModuleName = "oanqsec"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_oanqsec"
)

var (
	ParamsKey = []byte("p_oanqsec")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
