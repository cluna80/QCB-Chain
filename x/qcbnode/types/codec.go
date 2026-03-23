package types

import (
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	// this line is used by starport scaffolding # 1
)

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRegisterNode{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateNode{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgClaimNodeReward{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSlashNode{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeregisterNode{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetNodeConfig{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgReportNode{},
	)
	// this line is used by starport scaffolding # 3

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
