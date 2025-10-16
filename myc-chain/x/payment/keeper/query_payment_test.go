package keeper_test

import (
	"context"
	"strconv"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"myc-chain/x/payment/keeper"
	"myc-chain/x/payment/types"
)

func createNPayment(keeper keeper.Keeper, ctx context.Context, n int) []types.Payment {
	items := make([]types.Payment, n)
	for i := range items {
		iu := uint64(i)
		items[i].Id = iu
		items[i].MerchantId = strconv.Itoa(i)
		items[i].CustomerId = strconv.Itoa(i)
		items[i].Amount = strconv.Itoa(i)
		items[i].Status = strconv.Itoa(i)
		items[i].CreatedAt = int64(i)
		_ = keeper.Payment.Set(ctx, iu, items[i])
		_ = keeper.PaymentSeq.Set(ctx, iu)
	}
	return items
}

func TestPaymentQuerySingle(t *testing.T) {
	f := initFixture(t)
	qs := keeper.NewQueryServerImpl(f.keeper)
	msgs := createNPayment(f.keeper, f.ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetPaymentRequest
		response *types.QueryGetPaymentResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetPaymentRequest{Id: msgs[0].Id},
			response: &types.QueryGetPaymentResponse{Payment: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetPaymentRequest{Id: msgs[1].Id},
			response: &types.QueryGetPaymentResponse{Payment: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetPaymentRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := qs.GetPayment(f.ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.EqualExportedValues(t, tc.response, response)
			}
		})
	}
}

func TestPaymentQueryPaginated(t *testing.T) {
	f := initFixture(t)
	qs := keeper.NewQueryServerImpl(f.keeper)
	msgs := createNPayment(f.keeper, f.ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllPaymentRequest {
		return &types.QueryAllPaymentRequest{
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
			resp, err := qs.ListPayment(f.ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Payment), step)
			require.Subset(t, msgs, resp.Payment)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListPayment(f.ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Payment), step)
			require.Subset(t, msgs, resp.Payment)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := qs.ListPayment(f.ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.EqualExportedValues(t, msgs, resp.Payment)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := qs.ListPayment(f.ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
