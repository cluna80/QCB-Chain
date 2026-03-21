package types

import (
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	// this line is used by starport scaffolding # 1
)

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRegisterRelay{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRelayHeartbeat{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRouteMsg{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSlashRelay{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRemoveRelay{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateRelayRegion{},
	)
	// this line is used by starport scaffolding # 3

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
