package types

import "fmt"

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params:      DefaultParams(),
		MerchantMap: []Merchant{}, PaymentList: []Payment{}, SettlementList: []Settlement{}}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	merchantIndexMap := make(map[string]struct{})

	for _, elem := range gs.MerchantMap {
		index := fmt.Sprint(elem.Index)
		if _, ok := merchantIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for merchant")
		}
		merchantIndexMap[index] = struct{}{}
	}
	paymentIdMap := make(map[uint64]bool)
	paymentCount := gs.GetPaymentCount()
	for _, elem := range gs.PaymentList {
		if _, ok := paymentIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for payment")
		}
		if elem.Id >= paymentCount {
			return fmt.Errorf("payment id should be lower or equal than the last id")
		}
		paymentIdMap[elem.Id] = true
	}
	settlementIdMap := make(map[uint64]bool)
	settlementCount := gs.GetSettlementCount()
	for _, elem := range gs.SettlementList {
		if _, ok := settlementIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for settlement")
		}
		if elem.Id >= settlementCount {
			return fmt.Errorf("settlement id should be lower or equal than the last id")
		}
		settlementIdMap[elem.Id] = true
	}

	return gs.Params.Validate()
}
