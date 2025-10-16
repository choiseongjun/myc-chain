package keeper

import (
	"context"
	"errors"

	"myc-chain/x/payment/types"

	"cosmossdk.io/collections"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q queryServer) ListPayment(ctx context.Context, req *types.QueryAllPaymentRequest) (*types.QueryAllPaymentResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	payments, pageRes, err := query.CollectionPaginate(
		ctx,
		q.k.Payment,
		req.Pagination,
		func(_ uint64, value types.Payment) (types.Payment, error) {
			return value, nil
		},
	)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPaymentResponse{Payment: payments, Pagination: pageRes}, nil
}

func (q queryServer) GetPayment(ctx context.Context, req *types.QueryGetPaymentRequest) (*types.QueryGetPaymentResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	payment, err := q.k.Payment.Get(ctx, req.Id)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, sdkerrors.ErrKeyNotFound
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &types.QueryGetPaymentResponse{Payment: payment}, nil
}
