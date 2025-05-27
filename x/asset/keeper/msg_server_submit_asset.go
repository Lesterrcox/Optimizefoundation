package keeper

import (
	"context"

	"optimizeglobalcoin/common"
	"optimizeglobalcoin/x/asset/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SubmitAsset(goCtx context.Context, msg *types.MsgSubmitAsset) (*types.MsgSubmitAssetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.validatorKeeper.CheckIfValidAccess(ctx, msg.Creator) {
		return &types.MsgSubmitAssetResponse{Code: common.RET_FAIL_INTERNAL, Message: "invalid access"}, nil
	}
	// Check if the name and asset type are valid.
	if msg.Name == "" || msg.AssetType == types.AssetType_OTHER_ASSET {
		return &types.MsgSubmitAssetResponse{Code: common.RET_FAIL_INTERNAL, Message: "Invalid amount parameter"}, nil
	}

	// Check if the same assets already exists.
	listAssets := k.GetAssetsByNameAndType(ctx, msg.Name, msg.AssetType)
	if len(listAssets) > 0 {
		return &types.MsgSubmitAssetResponse{Code: common.RET_FAIL_INTERNAL, Message: "same assets already exists"}, nil
	}

	asset := types.Asset{
		Name:         msg.Name,
		AssetType:    msg.AssetType,
		Jurisdiction: msg.Jurisdiction,
		Owner:        msg.Creator,
		Status:       types.AssetStatus_PENDING,
		Value: &types.ValueValidation{
			Validators: make([]string, 0),
			ValueUsd:   "0",
			Timestamp:  0,
		},
		Options:   msg.Options,
		Timestamp: uint64(ctx.BlockTime().Unix()),
	}

	k.AppendAsset(ctx, asset)

	return &types.MsgSubmitAssetResponse{}, nil
}
