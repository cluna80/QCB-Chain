package types

import (
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	// this line is used by starport scaffolding # 1
)

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRegisterQsKey{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRotateQsKey{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSubmitHybridSig{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetThreatLevel{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRegisterAlgorithm{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeprecateAlgorithm{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgVerifyQsSignature{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgEmergencyCryptoSwap{},
	)
	// this line is used by starport scaffolding # 3

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
