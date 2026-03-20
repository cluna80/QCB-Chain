package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		AgentList: []Agent{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in agent
	agentIndexMap := make(map[string]struct{})

	for _, elem := range gs.AgentList {
		index := string(AgentKey(elem.Index))
		if _, ok := agentIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for agent")
		}
		agentIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
