package types

import (
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	// this line is used by starport scaffolding # 1
)

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendIbcAgent{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBroadcastAgentState{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgPostStateRoot{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRegisterChain{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgTokenizeOutput{},
	)
	// this line is used by starport scaffolding # 3

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
