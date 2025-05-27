package keeper

import (
	"context"

	"optimizeglobalcoin/x/asset/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ValueVotesAll(ctx context.Context, req *types.QueryAllValueVotesRequest) (*types.QueryAllValueVotesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var valueVotess []types.ValueVotes

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	valueVotesStore := prefix.NewStore(store, types.KeyPrefix(types.ValueVotesKey))

	pageRes, err := query.Paginate(valueVotesStore, req.Pagination, func(key []byte, value []byte) error {
		var valueVotes types.ValueVotes
		if err := k.cdc.Unmarshal(value, &valueVotes); err != nil {
			return err
		}

		valueVotess = append(valueVotess, valueVotes)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllValueVotesResponse{ValueVotes: valueVotess, Pagination: pageRes}, nil
}

func (k Keeper) ValueVotes(ctx context.Context, req *types.QueryGetValueVotesRequest) (*types.QueryGetValueVotesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	valueVotes, found := k.GetValueVotes(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetValueVotesResponse{ValueVotes: valueVotes}, nil
}
