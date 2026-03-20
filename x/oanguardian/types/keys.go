package types

const (
	// ModuleName defines the module name
	ModuleName = "oanguardian"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_oanguardian"
)

var (
	ParamsKey = []byte("p_oanguardian")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
