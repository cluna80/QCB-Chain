package types

const (
	// ModuleName defines the module name
	ModuleName = "antirug"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_antirug"
)

var (
	ParamsKey = []byte("p_antirug")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
