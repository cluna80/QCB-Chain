package types

type ConfinementRecord struct {
	Address      string
	Tier         string
	DailyReceive int64
	LastResetDay int64
	TotalReceive int64
}

type CharmedPair struct {
	HumanAddress string
	AgentID      string
	BondedAt     int64
	YieldBonus   int64
	Active       bool
}

type CharmParams struct {
	EpochBlocks          int64
	FeesliceBps          int64
	UnverifiedMaxBalance int64
	UnverifiedDailyLimit int64
	VerifiedMaxBalance   int64
	VerifiedDailyLimit   int64
	DAOMaxBalance        int64
	CharmedPairBonus     int64
}

func DefaultCharmParams() CharmParams {
	return CharmParams{
		EpochBlocks:          14400,
		FeesliceBps:          50,
		UnverifiedMaxBalance: 4_000_000_000,
		UnverifiedDailyLimit: 40_000_000,
		VerifiedMaxBalance:   400_000_000_000,
		VerifiedDailyLimit:   2_000_000_000,
		DAOMaxBalance:        800_000_000_000,
		CharmedPairBonus:     1000,
	}
}
