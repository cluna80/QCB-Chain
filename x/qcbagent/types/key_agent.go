package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// AgentKeyPrefix is the prefix to retrieve all Agent
	AgentKeyPrefix = "Agent/value/"
)

// AgentKey returns the store key to retrieve a Agent from the index fields
func AgentKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
