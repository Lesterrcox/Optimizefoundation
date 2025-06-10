package asset

import (
	"context"

	"optimizeglobalcoin/x/asset/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func EndBlocker(goCtx context.Context, k keeper.Keeper) error {
	ctx := sdk.UnwrapSDKContext(goCtx)

	k.UpdateAssetsValueByVotes(ctx)

	return nil
}
