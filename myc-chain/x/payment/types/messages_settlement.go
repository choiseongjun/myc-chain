package types

func NewMsgCreateSettlement(creator string, merchantId string, totalAmount string, settlementDate int64, status string) *MsgCreateSettlement {
	return &MsgCreateSettlement{
		Creator:        creator,
		MerchantId:     merchantId,
		TotalAmount:    totalAmount,
		SettlementDate: settlementDate,
		Status:         status,
	}
}

func NewMsgUpdateSettlement(creator string, id uint64, merchantId string, totalAmount string, settlementDate int64, status string) *MsgUpdateSettlement {
	return &MsgUpdateSettlement{
		Id:             id,
		Creator:        creator,
		MerchantId:     merchantId,
		TotalAmount:    totalAmount,
		SettlementDate: settlementDate,
		Status:         status,
	}
}

func NewMsgDeleteSettlement(creator string, id uint64) *MsgDeleteSettlement {
	return &MsgDeleteSettlement{
		Id:      id,
		Creator: creator,
	}
}
