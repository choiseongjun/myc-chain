package keeper_test

import (
	"context"
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"myc-chain/x/payment/keeper"
	"myc-chain/x/payment/types"
)

func createNMerchant(keeper keeper.Keeper, ctx context.Context, n int) []types.Merchant {
	items := make([]types.Merchant, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)
		items[i].Name = strconv.Itoa(i)
		items[i].Status = strconv.Itoa(i)
		items[i].RegisteredAt = int64(i)
		_ = keeper.Merchant.Set(ctx, items[i].Index, items[i])
	}
	return items
}

func TestMerchantQuerySingle(t *testing.T) {
	f := initFixture(t)
	qs := keeper.NewQueryServerImpl(f.keeper)
	msgs := createNMerchant(f.keeper, f.ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetMerchantRequest
		response *types.QueryGetMerchantResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetMerchantRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetMerchantResponse{Merchant: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetMerchantRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetMerchantResponse{Merchant: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetMerchantRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := qs.GetMerchant(f.ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.EqualExportedValues(t, tc.response, response)
			}
		})
	}
}

func TestMerchantQueryPaginated(t *testing.T) {
	f := initFixture(t)
	qs := keeper.NewQueryServerImpl(f.keeper)
	msgs := createNMerchant(f.keeper, f.ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllMerchantRequest {
		return &types.QueryAllMerchantRequest{
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
			resp, err := qs.ListMerchant(f.ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Merchant), step)
			require.Subset(t, msgs, resp.Merchant)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListMerchant(f.ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Merchant), step)
			require.Subset(t, msgs, resp.Merchant)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := qs.ListMerchant(f.ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.EqualExportedValues(t, msgs, resp.Merchant)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := qs.ListMerchant(f.ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
