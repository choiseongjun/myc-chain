package keeper_test

import (
	"testing"

	"myc-chain/x/payment/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params:      types.DefaultParams(),
		MerchantMap: []types.Merchant{{Index: "0"}, {Index: "1"}}, PaymentList: []types.Payment{{Id: 0}, {Id: 1}},
		PaymentCount:    2,
		SettlementList:  []types.Settlement{{Id: 0}, {Id: 1}},
		SettlementCount: 2,
	}
	f := initFixture(t)
	err := f.keeper.InitGenesis(f.ctx, genesisState)
	require.NoError(t, err)
	got, err := f.keeper.ExportGenesis(f.ctx)
	require.NoError(t, err)
	require.NotNil(t, got)

	require.EqualExportedValues(t, genesisState.Params, got.Params)
	require.EqualExportedValues(t, genesisState.MerchantMap, got.MerchantMap)
	require.EqualExportedValues(t, genesisState.PaymentList, got.PaymentList)
	require.Equal(t, genesisState.PaymentCount, got.PaymentCount)
	require.EqualExportedValues(t, genesisState.SettlementList, got.SettlementList)
	require.Equal(t, genesisState.SettlementCount, got.SettlementCount)

}
