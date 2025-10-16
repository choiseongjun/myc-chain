package keeper

import (
	"fmt"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	corestore "cosmossdk.io/core/store"
	"github.com/cosmos/cosmos-sdk/codec"

	"myc-chain/x/payment/types"
)

type Keeper struct {
	storeService corestore.KVStoreService
	cdc          codec.Codec
	addressCodec address.Codec
	// Address capable of executing a MsgUpdateParams message.
	// Typically, this should be the x/gov module account.
	authority []byte

	Schema collections.Schema
	Params collections.Item[types.Params]

	bankKeeper    types.BankKeeper
	Merchant      collections.Map[string, types.Merchant]
	PaymentSeq    collections.Sequence
	Payment       collections.Map[uint64, types.Payment]
	SettlementSeq collections.Sequence
	Settlement    collections.Map[uint64, types.Settlement]
}

func NewKeeper(
	storeService corestore.KVStoreService,
	cdc codec.Codec,
	addressCodec address.Codec,
	authority []byte,

	bankKeeper types.BankKeeper,
) Keeper {
	if _, err := addressCodec.BytesToString(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address %s: %s", authority, err))
	}

	sb := collections.NewSchemaBuilder(storeService)

	k := Keeper{
		storeService: storeService,
		cdc:          cdc,
		addressCodec: addressCodec,
		authority:    authority,

		bankKeeper: bankKeeper,
		Params:     collections.NewItem(sb, types.ParamsKey, "params", codec.CollValue[types.Params](cdc)),
		Merchant:   collections.NewMap(sb, types.MerchantKey, "merchant", collections.StringKey, codec.CollValue[types.Merchant](cdc)), Payment: collections.NewMap(sb, types.PaymentKey, "payment", collections.Uint64Key, codec.CollValue[types.Payment](cdc)),
		PaymentSeq:    collections.NewSequence(sb, types.PaymentCountKey, "paymentSequence"),
		Settlement:    collections.NewMap(sb, types.SettlementKey, "settlement", collections.Uint64Key, codec.CollValue[types.Settlement](cdc)),
		SettlementSeq: collections.NewSequence(sb, types.SettlementCountKey, "settlementSequence"),
	}
	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}
	k.Schema = schema

	return k
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() []byte {
	return k.authority
}
