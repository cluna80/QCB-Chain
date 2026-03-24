package types

type GenesisState struct {
	Params CharmParams
}

func DefaultGenesis() *GenesisState {
	return &GenesisState{Params: DefaultCharmParams()}
}

func (gs GenesisState) Validate() error { return nil }

func (gs *GenesisState) Reset()         {}
func (gs *GenesisState) String() string { return "" }
func (gs *GenesisState) ProtoMessage()  {}
