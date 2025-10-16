package payment

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"myc-chain/testutil/sample"
	paymentsimulation "myc-chain/x/payment/simulation"
	"myc-chain/x/payment/types"
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	paymentGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		MerchantMap: []types.Merchant{{Creator: sample.AccAddress(),
			Index: "0",
		}, {Creator: sample.AccAddress(),
			Index: "1",
		}}, PaymentList: []types.Payment{{Id: 0, Creator: sample.AccAddress()}, {Id: 1, Creator: sample.AccAddress()}}, PaymentCount: 2,
		SettlementList: []types.Settlement{{Id: 0, Creator: sample.AccAddress()}, {Id: 1, Creator: sample.AccAddress()}}, SettlementCount: 2,
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&paymentGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)
	const (
		opWeightMsgCreateMerchant          = "op_weight_msg_payment"
		defaultWeightMsgCreateMerchant int = 100
	)

	var weightMsgCreateMerchant int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateMerchant, &weightMsgCreateMerchant, nil,
		func(_ *rand.Rand) {
			weightMsgCreateMerchant = defaultWeightMsgCreateMerchant
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateMerchant,
		paymentsimulation.SimulateMsgCreateMerchant(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgUpdateMerchant          = "op_weight_msg_payment"
		defaultWeightMsgUpdateMerchant int = 100
	)

	var weightMsgUpdateMerchant int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateMerchant, &weightMsgUpdateMerchant, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateMerchant = defaultWeightMsgUpdateMerchant
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateMerchant,
		paymentsimulation.SimulateMsgUpdateMerchant(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgDeleteMerchant          = "op_weight_msg_payment"
		defaultWeightMsgDeleteMerchant int = 100
	)

	var weightMsgDeleteMerchant int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteMerchant, &weightMsgDeleteMerchant, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteMerchant = defaultWeightMsgDeleteMerchant
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteMerchant,
		paymentsimulation.SimulateMsgDeleteMerchant(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgCreatePayment          = "op_weight_msg_payment"
		defaultWeightMsgCreatePayment int = 100
	)

	var weightMsgCreatePayment int
	simState.AppParams.GetOrGenerate(opWeightMsgCreatePayment, &weightMsgCreatePayment, nil,
		func(_ *rand.Rand) {
			weightMsgCreatePayment = defaultWeightMsgCreatePayment
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreatePayment,
		paymentsimulation.SimulateMsgCreatePayment(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgUpdatePayment          = "op_weight_msg_payment"
		defaultWeightMsgUpdatePayment int = 100
	)

	var weightMsgUpdatePayment int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdatePayment, &weightMsgUpdatePayment, nil,
		func(_ *rand.Rand) {
			weightMsgUpdatePayment = defaultWeightMsgUpdatePayment
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdatePayment,
		paymentsimulation.SimulateMsgUpdatePayment(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgDeletePayment          = "op_weight_msg_payment"
		defaultWeightMsgDeletePayment int = 100
	)

	var weightMsgDeletePayment int
	simState.AppParams.GetOrGenerate(opWeightMsgDeletePayment, &weightMsgDeletePayment, nil,
		func(_ *rand.Rand) {
			weightMsgDeletePayment = defaultWeightMsgDeletePayment
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeletePayment,
		paymentsimulation.SimulateMsgDeletePayment(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgCreateSettlement          = "op_weight_msg_payment"
		defaultWeightMsgCreateSettlement int = 100
	)

	var weightMsgCreateSettlement int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateSettlement, &weightMsgCreateSettlement, nil,
		func(_ *rand.Rand) {
			weightMsgCreateSettlement = defaultWeightMsgCreateSettlement
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateSettlement,
		paymentsimulation.SimulateMsgCreateSettlement(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgUpdateSettlement          = "op_weight_msg_payment"
		defaultWeightMsgUpdateSettlement int = 100
	)

	var weightMsgUpdateSettlement int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateSettlement, &weightMsgUpdateSettlement, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateSettlement = defaultWeightMsgUpdateSettlement
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateSettlement,
		paymentsimulation.SimulateMsgUpdateSettlement(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgDeleteSettlement          = "op_weight_msg_payment"
		defaultWeightMsgDeleteSettlement int = 100
	)

	var weightMsgDeleteSettlement int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteSettlement, &weightMsgDeleteSettlement, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteSettlement = defaultWeightMsgDeleteSettlement
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteSettlement,
		paymentsimulation.SimulateMsgDeleteSettlement(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{}
}
