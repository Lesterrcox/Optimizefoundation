package keeper

import (
	"context"

	"optimizeglobalcoin/x/asset/types"

	common "optimizeglobalcoin/common"

	sdk "github.com/cosmos/cosmos-sdk/types"

	sdkmath "cosmossdk.io/math"
)

func (k msgServer) SubmitValue(goCtx context.Context, msg *types.MsgSubmitValue) (*types.MsgSubmitValueResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.validatorKeeper.CheckIfValidAccess(ctx, msg.Creator) {
		return &types.MsgSubmitValueResponse{Code: common.RET_FAIL_INTERNAL, Message: "invalid access"}, nil
	}

	_, found := k.GetAsset(ctx, msg.AssetId)
	if !found {
		return &types.MsgSubmitValueResponse{Code: common.RET_FAIL_INTERNAL, Message: "asset not found"}, nil
	}

	_, err := sdkmath.LegacyNewDecFromStr(msg.ValueUsd)
	if err != nil {
		return &types.MsgSubmitValueResponse{Code: common.RET_FAIL_INTERNAL, Message: "invalid value"}, nil
	}

	valueVotes := k.GetValueVotesByAssetIdAndVoter(ctx, msg.AssetId, msg.Creator)
	if len(valueVotes) > 0 {
		return &types.MsgSubmitValueResponse{Code: common.RET_FAIL_INTERNAL, Message: "value already submitted"}, nil
	}

	k.AppendValueVotes(ctx, types.ValueVotes{
		Voter:    msg.Creator,
		AssetId:  msg.AssetId,
		ValueUsd: msg.ValueUsd,
	})

	return &types.MsgSubmitValueResponse{}, nil
}
