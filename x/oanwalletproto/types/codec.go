package types

import (
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	// this line is used by starport scaffolding # 1
)

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRegisterWalletProfile{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetEncryptionKey{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetPqKey{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateWalletProfile{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetWalletPermissions{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgLockWallet{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUnlockWallet{},
	)
	// this line is used by starport scaffolding # 3

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
