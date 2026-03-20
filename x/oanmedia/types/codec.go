package types

import (
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	// this line is used by starport scaffolding # 1
)

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateMediaNft{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgClaimRoyalty{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgLicenseContent{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRecordView{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgMintMusicNft{},
	)
	// this line is used by starport scaffolding # 3

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
