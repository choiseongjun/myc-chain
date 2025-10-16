package keeper

import (
	"context"
	"errors"

	"myc-chain/x/payment/types"

	"cosmossdk.io/collections"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q queryServer) ListMerchant(ctx context.Context, req *types.QueryAllMerchantRequest) (*types.QueryAllMerchantResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	merchants, pageRes, err := query.CollectionPaginate(
		ctx,
		q.k.Merchant,
		req.Pagination,
		func(_ string, value types.Merchant) (types.Merchant, error) {
			return value, nil
		},
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllMerchantResponse{Merchant: merchants, Pagination: pageRes}, nil
}

func (q queryServer) GetMerchant(ctx context.Context, req *types.QueryGetMerchantRequest) (*types.QueryGetMerchantResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, err := q.k.Merchant.Get(ctx, req.Index)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "not found")
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &types.QueryGetMerchantResponse{Merchant: val}, nil
}
