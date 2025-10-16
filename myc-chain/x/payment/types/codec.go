package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterInterfaces(registrar codectypes.InterfaceRegistry) {
	registrar.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateSettlement{},
		&MsgUpdateSettlement{},
		&MsgDeleteSettlement{},
	)

	registrar.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreatePayment{},
		&MsgUpdatePayment{},
		&MsgDeletePayment{},
	)

	registrar.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateMerchant{},
		&MsgUpdateMerchant{},
		&MsgDeleteMerchant{},
	)

	registrar.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
	)
	msgservice.RegisterMsgServiceDesc(registrar, &_Msg_serviceDesc)
}
