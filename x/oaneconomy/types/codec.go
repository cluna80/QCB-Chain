package types

import (
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	// this line is used by starport scaffolding # 1
)

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateTask{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCompleteTask{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAcceptTask{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDisputeTask{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgClaimUbi{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgStakeTokens{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDistributeRewards{},
	)
	// this line is used by starport scaffolding # 3

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
