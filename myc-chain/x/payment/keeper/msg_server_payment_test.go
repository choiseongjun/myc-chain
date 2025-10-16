package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"myc-chain/x/payment/keeper"
	"myc-chain/x/payment/types"
)

func TestPaymentMsgServerCreate(t *testing.T) {
	f := initFixture(t)
	srv := keeper.NewMsgServerImpl(f.keeper)

	creator, err := f.addressCodec.BytesToString([]byte("signerAddr__________________"))
	require.NoError(t, err)

	for i := 0; i < 5; i++ {
		resp, err := srv.CreatePayment(f.ctx, &types.MsgCreatePayment{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestPaymentMsgServerUpdate(t *testing.T) {
	f := initFixture(t)
	srv := keeper.NewMsgServerImpl(f.keeper)

	creator, err := f.addressCodec.BytesToString([]byte("signerAddr__________________"))
	require.NoError(t, err)

	unauthorizedAddr, err := f.addressCodec.BytesToString([]byte("unauthorizedAddr___________"))
	require.NoError(t, err)

	_, err = srv.CreatePayment(f.ctx, &types.MsgCreatePayment{Creator: creator})
	require.NoError(t, err)

	tests := []struct {
		desc    string
		request *types.MsgUpdatePayment
		err     error
	}{
		{
			desc:    "invalid address",
			request: &types.MsgUpdatePayment{Creator: "invalid"},
			err:     sdkerrors.ErrInvalidAddress,
		},
		{
			desc:    "unauthorized",
			request: &types.MsgUpdatePayment{Creator: unauthorizedAddr},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "key not found",
			request: &types.MsgUpdatePayment{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc:    "completed",
			request: &types.MsgUpdatePayment{Creator: creator},
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			_, err = srv.UpdatePayment(f.ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestPaymentMsgServerDelete(t *testing.T) {
	f := initFixture(t)
	srv := keeper.NewMsgServerImpl(f.keeper)

	creator, err := f.addressCodec.BytesToString([]byte("signerAddr__________________"))
	require.NoError(t, err)

	unauthorizedAddr, err := f.addressCodec.BytesToString([]byte("unauthorizedAddr___________"))
	require.NoError(t, err)

	_, err = srv.CreatePayment(f.ctx, &types.MsgCreatePayment{Creator: creator})
	require.NoError(t, err)

	tests := []struct {
		desc    string
		request *types.MsgDeletePayment
		err     error
	}{
		{
			desc:    "invalid address",
			request: &types.MsgDeletePayment{Creator: "invalid"},
			err:     sdkerrors.ErrInvalidAddress,
		},
		{
			desc:    "unauthorized",
			request: &types.MsgDeletePayment{Creator: unauthorizedAddr},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "key not found",
			request: &types.MsgDeletePayment{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc:    "completed",
			request: &types.MsgDeletePayment{Creator: creator},
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			_, err = srv.DeletePayment(f.ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
