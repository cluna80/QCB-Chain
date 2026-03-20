package types

const (
	// ModuleName defines the module name
	ModuleName = "oandao"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_oandao"
)

var (
	ParamsKey = []byte("p_oandao")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
