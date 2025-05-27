package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "optimizeglobalcoin/testutil/keeper"
	"optimizeglobalcoin/testutil/nullify"
	"optimizeglobalcoin/x/validator/types"
)

func TestValidatorQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.OptimizeglobalcoinKeeper(t)
	msgs := createNValidator(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetValidatorRequest
		response *types.QueryGetValidatorResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetValidatorRequest{Id: msgs[0].Id},
			response: &types.QueryGetValidatorResponse{Validator: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetValidatorRequest{Id: msgs[1].Id},
			response: &types.QueryGetValidatorResponse{Validator: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetValidatorRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Validator(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestValidatorQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.OptimizeglobalcoinKeeper(t)
	msgs := createNValidator(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllValidatorRequest {
		return &types.QueryAllValidatorRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ValidatorAll(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Validators), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Validators),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ValidatorAll(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Validators), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Validators),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.ValidatorAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Validators),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ValidatorAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
