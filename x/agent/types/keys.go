package types

const (
	// ModuleName defines the module name
	ModuleName = "agent"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_agent"
)

var (
	ParamsKey = []byte("p_agent")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
