package keeper

import (
	"context"

	"myc-chain/x/payment/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func (k Keeper) InitGenesis(ctx context.Context, genState types.GenesisState) error {
	for _, elem := range genState.MerchantMap {
		if err := k.Merchant.Set(ctx, elem.Index, elem); err != nil {
			return err
		}
	}
	for _, elem := range genState.PaymentList {
		if err := k.Payment.Set(ctx, elem.Id, elem); err != nil {
			return err
		}
	}

	if err := k.PaymentSeq.Set(ctx, genState.PaymentCount); err != nil {
		return err
	}
	for _, elem := range genState.SettlementList {
		if err := k.Settlement.Set(ctx, elem.Id, elem); err != nil {
			return err
		}
	}

	if err := k.SettlementSeq.Set(ctx, genState.SettlementCount); err != nil {
		return err
	}

	return k.Params.Set(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis.
func (k Keeper) ExportGenesis(ctx context.Context) (*types.GenesisState, error) {
	var err error

	genesis := types.DefaultGenesis()
	genesis.Params, err = k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}
	if err := k.Merchant.Walk(ctx, nil, func(_ string, val types.Merchant) (stop bool, err error) {
		genesis.MerchantMap = append(genesis.MerchantMap, val)
		return false, nil
	}); err != nil {
		return nil, err
	}
	err = k.Payment.Walk(ctx, nil, func(key uint64, elem types.Payment) (bool, error) {
		genesis.PaymentList = append(genesis.PaymentList, elem)
		return false, nil
	})
	if err != nil {
		return nil, err
	}

	genesis.PaymentCount, err = k.PaymentSeq.Peek(ctx)
	if err != nil {
		return nil, err
	}
	err = k.Settlement.Walk(ctx, nil, func(key uint64, elem types.Settlement) (bool, error) {
		genesis.SettlementList = append(genesis.SettlementList, elem)
		return false, nil
	})
	if err != nil {
		return nil, err
	}

	genesis.SettlementCount, err = k.SettlementSeq.Peek(ctx)
	if err != nil {
		return nil, err
	}

	return genesis, nil
}
