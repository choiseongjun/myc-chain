package types

func NewMsgCreatePayment(creator string, merchantId string, customerId string, amount string, status string, createdAt int64) *MsgCreatePayment {
	return &MsgCreatePayment{
		Creator:    creator,
		MerchantId: merchantId,
		CustomerId: customerId,
		Amount:     amount,
		Status:     status,
		CreatedAt:  createdAt,
	}
}

func NewMsgUpdatePayment(creator string, id uint64, merchantId string, customerId string, amount string, status string, createdAt int64) *MsgUpdatePayment {
	return &MsgUpdatePayment{
		Id:         id,
		Creator:    creator,
		MerchantId: merchantId,
		CustomerId: customerId,
		Amount:     amount,
		Status:     status,
		CreatedAt:  createdAt,
	}
}

func NewMsgDeletePayment(creator string, id uint64) *MsgDeletePayment {
	return &MsgDeletePayment{
		Id:      id,
		Creator: creator,
	}
}
