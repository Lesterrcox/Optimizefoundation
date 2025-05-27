package asset

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "optimizeglobalcoin/api/optimizeglobalcoin/asset"
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
				{
					RpcMethod: "AssetAll",
					Use:       "list-asset",
					Short:     "List all asset",
				},
				{
					RpcMethod:      "Asset",
					Use:            "show-asset [id]",
					Short:          "Shows a asset by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod: "ValueVotesAll",
					Use:       "list-value-votes",
					Short:     "List all value-votes",
				},
				{
					RpcMethod:      "ValueVotes",
					Use:            "show-value-votes [id]",
					Short:          "Shows a value-votes by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
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
					RpcMethod:      "SubmitAsset",
					Use:            "submit-asset [name] [asset-type] [jurisdiction] [options]",
					Short:          "Send a submit-asset tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "name"}, {ProtoField: "asset_type"}, {ProtoField: "jurisdiction"}, {ProtoField: "options"}},
				},
				{
					RpcMethod:      "SubmitValue",
					Use:            "submit-value [asset-id] [value-usd]",
					Short:          "Send a submit-value tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "asset_id"}, {ProtoField: "value_usd"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
