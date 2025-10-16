package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"myc-chain/x/payment/keeper"
	"myc-chain/x/payment/types"
)

func TestSettlementMsgServerCreate(t *testing.T) {
	f := initFixture(t)
	srv := keeper.NewMsgServerImpl(f.keeper)

	creator, err := f.addressCodec.BytesToString([]byte("signerAddr__________________"))
	require.NoError(t, err)

	for i := 0; i < 5; i++ {
		resp, err := srv.CreateSettlement(f.ctx, &types.MsgCreateSettlement{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestSettlementMsgServerUpdate(t *testing.T) {
	f := initFixture(t)
	srv := keeper.NewMsgServerImpl(f.keeper)

	creator, err := f.addressCodec.BytesToString([]byte("signerAddr__________________"))
	require.NoError(t, err)

	unauthorizedAddr, err := f.addressCodec.BytesToString([]byte("unauthorizedAddr___________"))
	require.NoError(t, err)

	_, err = srv.CreateSettlement(f.ctx, &types.MsgCreateSettlement{Creator: creator})
	require.NoError(t, err)

	tests := []struct {
		desc    string
		request *types.MsgUpdateSettlement
		err     error
	}{
		{
			desc:    "invalid address",
			request: &types.MsgUpdateSettlement{Creator: "invalid"},
			err:     sdkerrors.ErrInvalidAddress,
		},
		{
			desc:    "unauthorized",
			request: &types.MsgUpdateSettlement{Creator: unauthorizedAddr},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "key not found",
			request: &types.MsgUpdateSettlement{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc:    "completed",
			request: &types.MsgUpdateSettlement{Creator: creator},
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			_, err = srv.UpdateSettlement(f.ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestSettlementMsgServerDelete(t *testing.T) {
	f := initFixture(t)
	srv := keeper.NewMsgServerImpl(f.keeper)

	creator, err := f.addressCodec.BytesToString([]byte("signerAddr__________________"))
	require.NoError(t, err)

	unauthorizedAddr, err := f.addressCodec.BytesToString([]byte("unauthorizedAddr___________"))
	require.NoError(t, err)

	_, err = srv.CreateSettlement(f.ctx, &types.MsgCreateSettlement{Creator: creator})
	require.NoError(t, err)

	tests := []struct {
		desc    string
		request *types.MsgDeleteSettlement
		err     error
	}{
		{
			desc:    "invalid address",
			request: &types.MsgDeleteSettlement{Creator: "invalid"},
			err:     sdkerrors.ErrInvalidAddress,
		},
		{
			desc:    "unauthorized",
			request: &types.MsgDeleteSettlement{Creator: unauthorizedAddr},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "key not found",
			request: &types.MsgDeleteSettlement{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc:    "completed",
			request: &types.MsgDeleteSettlement{Creator: creator},
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			_, err = srv.DeleteSettlement(f.ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
