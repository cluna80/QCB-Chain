package types

import (
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	// this line is used by starport scaffolding # 1
)

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddGuardian{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRemoveGuardian{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgGuardianVeto{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetAiLimits{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgEmergencyPause{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgLiftPause{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgApproveModel{},
	)
	// this line is used by starport scaffolding # 3

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
