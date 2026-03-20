package oanmedia

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "oan/api/oan/oanmedia"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateMediaNft",
					Use:            "create-media-nft [title] [media-type] [content-hash] [creator-share]",
					Short:          "Send a create-media-nft tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "title"}, {ProtoField: "mediaType"}, {ProtoField: "contentHash"}, {ProtoField: "creatorShare"}},
				},
				{
					RpcMethod:      "ClaimRoyalty",
					Use:            "claim-royalty [nft-id] [period-start] [period-end]",
					Short:          "Send a claim-royalty tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "nftId"}, {ProtoField: "periodStart"}, {ProtoField: "periodEnd"}},
				},
				{
					RpcMethod:      "LicenseContent",
					Use:            "license-content [nft-id] [licensee-addr] [license-type] [duration] [fee]",
					Short:          "Send a license-content tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "nftId"}, {ProtoField: "licenseeAddr"}, {ProtoField: "licenseType"}, {ProtoField: "duration"}, {ProtoField: "fee"}},
				},
				{
					RpcMethod:      "RecordView",
					Use:            "record-view [nft-id] [viewer-addr] [payment-amount]",
					Short:          "Send a record-view tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "nftId"}, {ProtoField: "viewerAddr"}, {ProtoField: "paymentAmount"}},
				},
				{
					RpcMethod:      "MintMusicNft",
					Use:            "mint-music-nft [title] [audio-hash] [agent-id] [bpm] [genre]",
					Short:          "Send a mint-music-nft tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "title"}, {ProtoField: "audioHash"}, {ProtoField: "agentId"}, {ProtoField: "bpm"}, {ProtoField: "genre"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
