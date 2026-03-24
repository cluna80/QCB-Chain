package types

const (
	ModuleName = "charm"
	StoreKey   = ModuleName
	RouterKey  = ModuleName

	KeyConfinementPrefix  = "confinement/"
	KeyUBIPool            = "ubi-pool"
	KeyFeePool            = "fee-pool"
	KeyLastEpoch          = "last-epoch"
	KeyCharmedPairPrefix  = "charmed-pair/"
	KeyDailyReceivePrefix = "daily-receive/"
	KeyBalanceTierPrefix  = "balance-tier/"

	TierUnverified  = "unverified"
	TierVerified    = "verified"
	TierDAOApproved = "dao_approved"
	TierExempt      = "exempt"
)
