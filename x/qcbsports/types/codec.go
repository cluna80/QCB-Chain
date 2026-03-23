package types

import (
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	// this line is used by starport scaffolding # 1
)

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRegisterAthlete{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateStadium{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgScheduleMatch{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRecordMatchResult{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgPlacePrediction{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgClaimPredictionReward{},
	)
	// this line is used by starport scaffolding # 3

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
