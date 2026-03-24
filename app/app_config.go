package app

import (
	"time"

	runtimev1alpha1 "cosmossdk.io/api/cosmos/app/runtime/v1alpha1"
	appv1alpha1 "cosmossdk.io/api/cosmos/app/v1alpha1"
	authmodulev1 "cosmossdk.io/api/cosmos/auth/module/v1"
	authzmodulev1 "cosmossdk.io/api/cosmos/authz/module/v1"
	bankmodulev1 "cosmossdk.io/api/cosmos/bank/module/v1"
	circuitmodulev1 "cosmossdk.io/api/cosmos/circuit/module/v1"
	consensusmodulev1 "cosmossdk.io/api/cosmos/consensus/module/v1"
	crisismodulev1 "cosmossdk.io/api/cosmos/crisis/module/v1"
	distrmodulev1 "cosmossdk.io/api/cosmos/distribution/module/v1"
	evidencemodulev1 "cosmossdk.io/api/cosmos/evidence/module/v1"
	feegrantmodulev1 "cosmossdk.io/api/cosmos/feegrant/module/v1"
	genutilmodulev1 "cosmossdk.io/api/cosmos/genutil/module/v1"
	govmodulev1 "cosmossdk.io/api/cosmos/gov/module/v1"
	groupmodulev1 "cosmossdk.io/api/cosmos/group/module/v1"
	mintmodulev1 "cosmossdk.io/api/cosmos/mint/module/v1"
	nftmodulev1 "cosmossdk.io/api/cosmos/nft/module/v1"
	paramsmodulev1 "cosmossdk.io/api/cosmos/params/module/v1"
	slashingmodulev1 "cosmossdk.io/api/cosmos/slashing/module/v1"
	stakingmodulev1 "cosmossdk.io/api/cosmos/staking/module/v1"
	txconfigv1 "cosmossdk.io/api/cosmos/tx/config/v1"
	upgrademodulev1 "cosmossdk.io/api/cosmos/upgrade/module/v1"
	vestingmodulev1 "cosmossdk.io/api/cosmos/vesting/module/v1"
	"cosmossdk.io/core/appconfig"
	circuittypes "cosmossdk.io/x/circuit/types"
	evidencetypes "cosmossdk.io/x/evidence/types"
	"cosmossdk.io/x/feegrant"
	"cosmossdk.io/x/nft"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	consensustypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/group"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	capabilitytypes "github.com/cosmos/ibc-go/modules/capability/types"
	icatypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/types"
	ibcfeetypes "github.com/cosmos/ibc-go/v8/modules/apps/29-fee/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	ibcexported "github.com/cosmos/ibc-go/v8/modules/core/exported"
	"google.golang.org/protobuf/types/known/durationpb"

	agentmodulev1 "qcb/api/qcb/agent/module"
	antirugmodulev1 "qcb/api/qcb/antirug/module"
	qcbagentmodulev1 "qcb/api/qcb/qcbagent/module"
	qcbbridgemodulev1 "qcb/api/qcb/qcbbridge/module"
	qcbcommsmodulev1 "qcb/api/qcb/qcbcomms/module"
	qcbcomputemodulev1 "qcb/api/qcb/qcbcompute/module"
	qcbdaomodulev1 "qcb/api/qcb/qcbdao/module"
	qcbeconomymodulev1 "qcb/api/qcb/qcbeconomy/module"
	qcbguardianmodulev1 "qcb/api/qcb/qcbguardian/module"
	qcbidentitymodulev1 "qcb/api/qcb/qcbidentity/module"
	qcbmarketmodulev1 "qcb/api/qcb/qcbmarket/module"
	qcbmediamodulev1 "qcb/api/qcb/qcbmedia/module"
	qcbnodemodulev1 "qcb/api/qcb/qcbnode/module"
	qcbprotocolmodulev1 "qcb/api/qcb/qcbprotocol/module"
	qcbqsecmodulev1 "qcb/api/qcb/qcbqsec/module"
	qcbrelaymodulev1 "qcb/api/qcb/qcbrelay/module"
	qcbsportsmodulev1 "qcb/api/qcb/qcbsports/module"
	qcbwalletprotomodulev1 "qcb/api/qcb/qcbwalletproto/module"
	_ "qcb/x/agent/module" // import for side-effects
	agentmoduletypes "qcb/x/agent/types"
	_ "qcb/x/antirug/module" // import for side-effects
	antirugmoduletypes "qcb/x/antirug/types"
	_ "qcb/x/qcbagent/module" // import for side-effects
	qcbagentmoduletypes "qcb/x/qcbagent/types"
	_ "qcb/x/qcbbridge/module" // import for side-effects
	qcbbridgemoduletypes "qcb/x/qcbbridge/types"
	_ "qcb/x/qcbcomms/module" // import for side-effects
	qcbcommsmoduletypes "qcb/x/qcbcomms/types"
	_ "qcb/x/qcbcompute/module" // import for side-effects
	qcbcomputemoduletypes "qcb/x/qcbcompute/types"
	_ "qcb/x/qcbdao/module" // import for side-effects
	qcbdaomoduletypes "qcb/x/qcbdao/types"
	_ "qcb/x/qcbeconomy/module" // import for side-effects
	qcbeconomymoduletypes "qcb/x/qcbeconomy/types"
	_ "qcb/x/charm/module" // import for side-effects
	charmtypes "qcb/x/charm/types"
	_ "qcb/x/qcbguardian/module" // import for side-effects
	qcbguardianmoduletypes "qcb/x/qcbguardian/types"
	_ "qcb/x/qcbidentity/module" // import for side-effects
	qcbidentitymoduletypes "qcb/x/qcbidentity/types"
	_ "qcb/x/qcbmarket/module" // import for side-effects
	qcbmarketmoduletypes "qcb/x/qcbmarket/types"
	_ "qcb/x/qcbmedia/module" // import for side-effects
	qcbmediamoduletypes "qcb/x/qcbmedia/types"
	_ "qcb/x/qcbnode/module" // import for side-effects
	qcbnodemoduletypes "qcb/x/qcbnode/types"
	_ "qcb/x/qcbprotocol/module" // import for side-effects
	qcbprotocolmoduletypes "qcb/x/qcbprotocol/types"
	_ "qcb/x/qcbqsec/module" // import for side-effects
	qcbqsecmoduletypes "qcb/x/qcbqsec/types"
	_ "qcb/x/qcbrelay/module" // import for side-effects
	qcbrelaymoduletypes "qcb/x/qcbrelay/types"
	_ "qcb/x/qcbsports/module" // import for side-effects
	qcbsportsmoduletypes "qcb/x/qcbsports/types"
	_ "qcb/x/qcbwalletproto/module" // import for side-effects
	qcbwalletprotomoduletypes "qcb/x/qcbwalletproto/types"
	// this line is used by starport scaffolding # stargate/app/moduleImport
)

var (
	// NOTE: The genutils module must occur after staking so that pools are
	// properly initialized with tokens from genesis accounts.
	// NOTE: The genutils module must also occur after auth so that it can access the params from auth.
	// NOTE: Capability module must occur first so that it can initialize any capabilities
	// so that other modules that want to create or claim capabilities afterwards in InitChain
	// can do so safely.
	genesisModuleOrder = []string{
		// cosmos-sdk/ibc modules
		capabilitytypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		distrtypes.ModuleName,
		stakingtypes.ModuleName,
		slashingtypes.ModuleName,
		govtypes.ModuleName,
		minttypes.ModuleName,
		crisistypes.ModuleName,
		ibcexported.ModuleName,
		genutiltypes.ModuleName,
		evidencetypes.ModuleName,
		authz.ModuleName,
		ibctransfertypes.ModuleName,
		icatypes.ModuleName,
		ibcfeetypes.ModuleName,
		feegrant.ModuleName,
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
		vestingtypes.ModuleName,
		nft.ModuleName,
		group.ModuleName,
		consensustypes.ModuleName,
		circuittypes.ModuleName,
		// chain modules
		agentmoduletypes.ModuleName,
		qcbagentmoduletypes.ModuleName,
		qcbidentitymoduletypes.ModuleName,
		qcbdaomoduletypes.ModuleName,
		qcbeconomymoduletypes.ModuleName,
		qcbcomputemoduletypes.ModuleName,
		qcbmarketmoduletypes.ModuleName,
		qcbguardianmoduletypes.ModuleName,
		qcbqsecmoduletypes.ModuleName,
		qcbmediamoduletypes.ModuleName,
		qcbsportsmoduletypes.ModuleName,
		qcbbridgemoduletypes.ModuleName,
		qcbnodemoduletypes.ModuleName,
		antirugmoduletypes.ModuleName,
		qcbcommsmoduletypes.ModuleName,
		qcbrelaymoduletypes.ModuleName,
		qcbwalletprotomoduletypes.ModuleName,
		qcbprotocolmoduletypes.ModuleName,
		charmtypes.ModuleName,
		// this line is used by starport scaffolding # stargate/app/initGenesis
	}

	// During begin block slashing happens after distr.BeginBlocker so that
	// there is nothing left over in the validator fee pool, so as to keep the
	// CanWithdrawInvariant invariant.
	// NOTE: staking module is required if HistoricalEntries param > 0
	// NOTE: capability module's beginblocker must come before any modules using capabilities (e.g. IBC)
	beginBlockers = []string{
		// cosmos sdk modules
		minttypes.ModuleName,
		distrtypes.ModuleName,
		slashingtypes.ModuleName,
		evidencetypes.ModuleName,
		stakingtypes.ModuleName,
		authz.ModuleName,
		genutiltypes.ModuleName,
		// ibc modules
		capabilitytypes.ModuleName,
		ibcexported.ModuleName,
		ibctransfertypes.ModuleName,
		icatypes.ModuleName,
		ibcfeetypes.ModuleName,
		// chain modules
		agentmoduletypes.ModuleName,
		qcbagentmoduletypes.ModuleName,
		qcbidentitymoduletypes.ModuleName,
		qcbdaomoduletypes.ModuleName,
		qcbeconomymoduletypes.ModuleName,
		qcbcomputemoduletypes.ModuleName,
		qcbmarketmoduletypes.ModuleName,
		qcbguardianmoduletypes.ModuleName,
		qcbqsecmoduletypes.ModuleName,
		qcbmediamoduletypes.ModuleName,
		qcbsportsmoduletypes.ModuleName,
		qcbbridgemoduletypes.ModuleName,
		qcbnodemoduletypes.ModuleName,
		antirugmoduletypes.ModuleName,
		qcbcommsmoduletypes.ModuleName,
		qcbrelaymoduletypes.ModuleName,
		qcbwalletprotomoduletypes.ModuleName,
		qcbprotocolmoduletypes.ModuleName,
		charmtypes.ModuleName,
		// this line is used by starport scaffolding # stargate/app/beginBlockers
	}

	endBlockers = []string{
		// cosmos sdk modules
		crisistypes.ModuleName,
		govtypes.ModuleName,
		stakingtypes.ModuleName,
		feegrant.ModuleName,
		group.ModuleName,
		genutiltypes.ModuleName,
		// ibc modules
		ibcexported.ModuleName,
		ibctransfertypes.ModuleName,
		capabilitytypes.ModuleName,
		icatypes.ModuleName,
		ibcfeetypes.ModuleName,
		// chain modules
		agentmoduletypes.ModuleName,
		qcbagentmoduletypes.ModuleName,
		qcbidentitymoduletypes.ModuleName,
		qcbdaomoduletypes.ModuleName,
		qcbeconomymoduletypes.ModuleName,
		qcbcomputemoduletypes.ModuleName,
		qcbmarketmoduletypes.ModuleName,
		qcbguardianmoduletypes.ModuleName,
		qcbqsecmoduletypes.ModuleName,
		qcbmediamoduletypes.ModuleName,
		qcbsportsmoduletypes.ModuleName,
		qcbbridgemoduletypes.ModuleName,
		qcbnodemoduletypes.ModuleName,
		antirugmoduletypes.ModuleName,
		qcbcommsmoduletypes.ModuleName,
		qcbrelaymoduletypes.ModuleName,
		qcbwalletprotomoduletypes.ModuleName,
		qcbprotocolmoduletypes.ModuleName,
		charmtypes.ModuleName,
		// this line is used by starport scaffolding # stargate/app/endBlockers
	}

	preBlockers = []string{
		upgradetypes.ModuleName,
		// this line is used by starport scaffolding # stargate/app/preBlockers
	}

	// module account permissions
	moduleAccPerms = []*authmodulev1.ModuleAccountPermission{
		{Account: authtypes.FeeCollectorName},
		{Account: distrtypes.ModuleName},
		{Account: minttypes.ModuleName, Permissions: []string{authtypes.Minter}},
		{Account: stakingtypes.BondedPoolName, Permissions: []string{authtypes.Burner, stakingtypes.ModuleName}},
		{Account: stakingtypes.NotBondedPoolName, Permissions: []string{authtypes.Burner, stakingtypes.ModuleName}},
		{Account: govtypes.ModuleName, Permissions: []string{authtypes.Burner}},
		{Account: nft.ModuleName},
		{Account: ibctransfertypes.ModuleName, Permissions: []string{authtypes.Minter, authtypes.Burner}},
		{Account: ibcfeetypes.ModuleName},
		{Account: icatypes.ModuleName},
		{Account: "qcbeconomy", Permissions: []string{authtypes.Burner, authtypes.Minter}},
				{Account: "charm", Permissions: []string{authtypes.Burner, authtypes.Minter}},
		{Account: "qcbagent", Permissions: []string{authtypes.Burner}},
		{Account: "qcbnode", Permissions: []string{authtypes.Burner, authtypes.Minter}},
		{Account: "qcbmarket", Permissions: []string{authtypes.Burner}},
		{Account: "qcbcompute", Permissions: []string{authtypes.Burner, authtypes.Minter}},
		{Account: "qcbbridge", Permissions: []string{authtypes.Burner}},
		{Account: qcbprotocolmoduletypes.ModuleName, Permissions: []string{authtypes.Minter, authtypes.Burner, authtypes.Staking}},
		// this line is used by starport scaffolding # stargate/app/maccPerms
	}

	// blocked account addresses
	blockAccAddrs = []string{
		authtypes.FeeCollectorName,
		distrtypes.ModuleName,
		minttypes.ModuleName,
		stakingtypes.BondedPoolName,
		stakingtypes.NotBondedPoolName,
		nft.ModuleName,
		// We allow the following module accounts to receive funds:
		// govtypes.ModuleName
	}

	// appConfig application configuration (used by depinject)
	appConfig = appconfig.Compose(&appv1alpha1.Config{
		Modules: []*appv1alpha1.ModuleConfig{
			{
				Name: runtime.ModuleName,
				Config: appconfig.WrapAny(&runtimev1alpha1.Module{
					AppName:       Name,
					PreBlockers:   preBlockers,
					BeginBlockers: beginBlockers,
					EndBlockers:   endBlockers,
					InitGenesis:   genesisModuleOrder,
					OverrideStoreKeys: []*runtimev1alpha1.StoreKeyConfig{
						{
							ModuleName: authtypes.ModuleName,
							KvStoreKey: "acc",
						},
					},
					// When ExportGenesis is not specified, the export genesis module order
					// is equal to the init genesis order
					// ExportGenesis: genesisModuleOrder,
					// Uncomment if you want to set a custom migration order here.
					// OrderMigrations: nil,
				}),
			},
			{
				Name: authtypes.ModuleName,
				Config: appconfig.WrapAny(&authmodulev1.Module{
					Bech32Prefix:             AccountAddressPrefix,
					ModuleAccountPermissions: moduleAccPerms,
					// By default modules authority is the governance module. This is configurable with the following:
					// Authority: "group", // A custom module authority can be set using a module name
					// Authority: "cosmos1cwwv22j5ca08ggdv9c2uky355k908694z577tv", // or a specific address
				}),
			},
			{
				Name:   nft.ModuleName,
				Config: appconfig.WrapAny(&nftmodulev1.Module{}),
			},
			{
				Name:   vestingtypes.ModuleName,
				Config: appconfig.WrapAny(&vestingmodulev1.Module{}),
			},
			{
				Name: banktypes.ModuleName,
				Config: appconfig.WrapAny(&bankmodulev1.Module{
					BlockedModuleAccountsOverride: blockAccAddrs,
				}),
			},
			{
				Name: stakingtypes.ModuleName,
				Config: appconfig.WrapAny(&stakingmodulev1.Module{
					// NOTE: specifying a prefix is only necessary when using bech32 addresses
					// If not specfied, the auth Bech32Prefix appended with "valoper" and "valcons" is used by default
					Bech32PrefixValidator: AccountAddressPrefix + "valoper",
					Bech32PrefixConsensus: AccountAddressPrefix + "valcons",
				}),
			},
			{
				Name:   slashingtypes.ModuleName,
				Config: appconfig.WrapAny(&slashingmodulev1.Module{}),
			},
			{
				Name:   paramstypes.ModuleName,
				Config: appconfig.WrapAny(&paramsmodulev1.Module{}),
			},
			{
				Name:   "tx",
				Config: appconfig.WrapAny(&txconfigv1.Config{}),
			},
			{
				Name:   genutiltypes.ModuleName,
				Config: appconfig.WrapAny(&genutilmodulev1.Module{}),
			},
			{
				Name:   authz.ModuleName,
				Config: appconfig.WrapAny(&authzmodulev1.Module{}),
			},
			{
				Name:   upgradetypes.ModuleName,
				Config: appconfig.WrapAny(&upgrademodulev1.Module{}),
			},
			{
				Name:   distrtypes.ModuleName,
				Config: appconfig.WrapAny(&distrmodulev1.Module{}),
			},
			{
				Name:   evidencetypes.ModuleName,
				Config: appconfig.WrapAny(&evidencemodulev1.Module{}),
			},
			{
				Name:   minttypes.ModuleName,
				Config: appconfig.WrapAny(&mintmodulev1.Module{}),
			},
			{
				Name: group.ModuleName,
				Config: appconfig.WrapAny(&groupmodulev1.Module{
					MaxExecutionPeriod: durationpb.New(time.Second * 1209600),
					MaxMetadataLen:     255,
				}),
			},
			{
				Name:   feegrant.ModuleName,
				Config: appconfig.WrapAny(&feegrantmodulev1.Module{}),
			},
			{
				Name:   govtypes.ModuleName,
				Config: appconfig.WrapAny(&govmodulev1.Module{}),
			},
			{
				Name:   crisistypes.ModuleName,
				Config: appconfig.WrapAny(&crisismodulev1.Module{}),
			},
			{
				Name:   consensustypes.ModuleName,
				Config: appconfig.WrapAny(&consensusmodulev1.Module{}),
			},
			{
				Name:   circuittypes.ModuleName,
				Config: appconfig.WrapAny(&circuitmodulev1.Module{}),
			},
			{
				Name:   agentmoduletypes.ModuleName,
				Config: appconfig.WrapAny(&agentmodulev1.Module{}),
			},
			{
				Name:   qcbagentmoduletypes.ModuleName,
				Config: appconfig.WrapAny(&qcbagentmodulev1.Module{}),
			},
			{
				Name:   qcbidentitymoduletypes.ModuleName,
				Config: appconfig.WrapAny(&qcbidentitymodulev1.Module{}),
			},
			{
				Name:   qcbdaomoduletypes.ModuleName,
				Config: appconfig.WrapAny(&qcbdaomodulev1.Module{}),
			},
			{
				Name:   qcbeconomymoduletypes.ModuleName,
				Config: appconfig.WrapAny(&qcbeconomymodulev1.Module{}),
			},
			{
				Name:   qcbcomputemoduletypes.ModuleName,
				Config: appconfig.WrapAny(&qcbcomputemodulev1.Module{}),
			},
			{
				Name:   qcbmarketmoduletypes.ModuleName,
				Config: appconfig.WrapAny(&qcbmarketmodulev1.Module{}),
			},
			{
				Name:   qcbguardianmoduletypes.ModuleName,
				Config: appconfig.WrapAny(&qcbguardianmodulev1.Module{}),
			},
			{
				Name:   qcbqsecmoduletypes.ModuleName,
				Config: appconfig.WrapAny(&qcbqsecmodulev1.Module{}),
			},
			{
				Name:   qcbmediamoduletypes.ModuleName,
				Config: appconfig.WrapAny(&qcbmediamodulev1.Module{}),
			},
			{
				Name:   qcbsportsmoduletypes.ModuleName,
				Config: appconfig.WrapAny(&qcbsportsmodulev1.Module{}),
			},
			{
				Name:   qcbbridgemoduletypes.ModuleName,
				Config: appconfig.WrapAny(&qcbbridgemodulev1.Module{}),
			},
			{
				Name:   qcbnodemoduletypes.ModuleName,
				Config: appconfig.WrapAny(&qcbnodemodulev1.Module{}),
			},
			{
				Name:   antirugmoduletypes.ModuleName,
				Config: appconfig.WrapAny(&antirugmodulev1.Module{}),
			},
			{
				Name:   qcbcommsmoduletypes.ModuleName,
				Config: appconfig.WrapAny(&qcbcommsmodulev1.Module{}),
			},
			{
				Name:   qcbrelaymoduletypes.ModuleName,
				Config: appconfig.WrapAny(&qcbrelaymodulev1.Module{}),
			},
			{
				Name:   qcbwalletprotomoduletypes.ModuleName,
				Config: appconfig.WrapAny(&qcbwalletprotomodulev1.Module{}),
			},
			{
				Name:   qcbprotocolmoduletypes.ModuleName,
				Config: appconfig.WrapAny(&qcbprotocolmodulev1.Module{}),
			},
			// this line is used by starport scaffolding # stargate/app/moduleConfig
		},
	})
)
