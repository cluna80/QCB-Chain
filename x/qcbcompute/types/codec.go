package types

import (
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	// this line is used by starport scaffolding # 1
)

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSubmitInferenceJob{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCompleteInferenceJob{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSlashBadInference{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgStakeCompute{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRegisterModel{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgVerifyInferenceProof{},
	)
	// this line is used by starport scaffolding # 3

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
