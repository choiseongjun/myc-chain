package payment

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	"myc-chain/x/payment/types"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: types.Query_serviceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "ListMerchant",
					Use:       "list-merchant",
					Short:     "List all merchant",
				},
				{
					RpcMethod:      "GetMerchant",
					Use:            "get-merchant [id]",
					Short:          "Gets a merchant",
					Alias:          []string{"show-merchant"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod: "ListPayment",
					Use:       "list-payment",
					Short:     "List all payment",
				},
				{
					RpcMethod:      "GetPayment",
					Use:            "get-payment [id]",
					Short:          "Gets a payment by id",
					Alias:          []string{"show-payment"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod: "ListSettlement",
					Use:       "list-settlement",
					Short:     "List all settlement",
				},
				{
					RpcMethod:      "GetSettlement",
					Use:            "get-settlement [id]",
					Short:          "Gets a settlement by id",
					Alias:          []string{"show-settlement"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              types.Msg_serviceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateMerchant",
					Use:            "create-merchant [index] [name] [status] [registered-at]",
					Short:          "Create a new merchant",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "name"}, {ProtoField: "status"}, {ProtoField: "registered_at"}},
				},
				{
					RpcMethod:      "UpdateMerchant",
					Use:            "update-merchant [index] [name] [status] [registered-at]",
					Short:          "Update merchant",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "name"}, {ProtoField: "status"}, {ProtoField: "registered_at"}},
				},
				{
					RpcMethod:      "DeleteMerchant",
					Use:            "delete-merchant [index]",
					Short:          "Delete merchant",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod:      "CreatePayment",
					Use:            "create-payment [merchant-id] [customer-id] [amount] [status] [created-at]",
					Short:          "Create payment",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "merchant_id"}, {ProtoField: "customer_id"}, {ProtoField: "amount"}, {ProtoField: "status"}, {ProtoField: "created_at"}},
				},
				{
					RpcMethod:      "UpdatePayment",
					Use:            "update-payment [id] [merchant-id] [customer-id] [amount] [status] [created-at]",
					Short:          "Update payment",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "merchant_id"}, {ProtoField: "customer_id"}, {ProtoField: "amount"}, {ProtoField: "status"}, {ProtoField: "created_at"}},
				},
				{
					RpcMethod:      "DeletePayment",
					Use:            "delete-payment [id]",
					Short:          "Delete payment",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod:      "CreateSettlement",
					Use:            "create-settlement [merchant-id] [total-amount] [settlement-date] [status]",
					Short:          "Create settlement",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "merchant_id"}, {ProtoField: "total_amount"}, {ProtoField: "settlement_date"}, {ProtoField: "status"}},
				},
				{
					RpcMethod:      "UpdateSettlement",
					Use:            "update-settlement [id] [merchant-id] [total-amount] [settlement-date] [status]",
					Short:          "Update settlement",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "merchant_id"}, {ProtoField: "total_amount"}, {ProtoField: "settlement_date"}, {ProtoField: "status"}},
				},
				{
					RpcMethod:      "DeleteSettlement",
					Use:            "delete-settlement [id]",
					Short:          "Delete settlement",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
