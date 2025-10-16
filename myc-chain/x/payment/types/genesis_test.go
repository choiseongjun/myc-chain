package types_test

import (
	"testing"

	"myc-chain/x/payment/types"

	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc:     "valid genesis state",
			genState: &types.GenesisState{MerchantMap: []types.Merchant{{Index: "0"}, {Index: "1"}}, PaymentList: []types.Payment{{Id: 0}, {Id: 1}}, PaymentCount: 2, SettlementList: []types.Settlement{{Id: 0}, {Id: 1}}, SettlementCount: 2}, valid: true,
		}, {
			desc: "duplicated merchant",
			genState: &types.GenesisState{
				MerchantMap: []types.Merchant{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
				PaymentList: []types.Payment{{Id: 0}, {Id: 1}}, PaymentCount: 2,
				SettlementList: []types.Settlement{{Id: 0}, {Id: 1}}, SettlementCount: 2}, valid: false,
		}, {
			desc: "duplicated payment",
			genState: &types.GenesisState{
				PaymentList: []types.Payment{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
				SettlementList: []types.Settlement{{Id: 0}, {Id: 1}}, SettlementCount: 2,
			}, valid: false,
		}, {
			desc: "invalid payment count",
			genState: &types.GenesisState{
				PaymentList: []types.Payment{
					{
						Id: 1,
					},
				},
				PaymentCount:   0,
				SettlementList: []types.Settlement{{Id: 0}, {Id: 1}}, SettlementCount: 2,
			}, valid: false,
		}, {
			desc: "duplicated settlement",
			genState: &types.GenesisState{
				SettlementList: []types.Settlement{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		}, {
			desc: "invalid settlement count",
			genState: &types.GenesisState{
				SettlementList: []types.Settlement{
					{
						Id: 1,
					},
				},
				SettlementCount: 0,
			},
			valid: false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
