package types

import (
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	// this line is used by starport scaffolding # 1
)

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRegisterToken{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgLockLiquidity{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeclareMintLimit{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgFlagToken{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgFreezeToken{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUnfreezeToken{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRequestUpgrade{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgApproveToken{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetModuleActive{},
	)
	// this line is used by starport scaffolding # 3

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
