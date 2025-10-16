package types

func NewMsgCreateMerchant(
	creator string,
	index string,
	name string,
	status string,
	registeredAt int64,

) *MsgCreateMerchant {
	return &MsgCreateMerchant{
		Creator:      creator,
		Index:        index,
		Name:         name,
		Status:       status,
		RegisteredAt: registeredAt,
	}
}

func NewMsgUpdateMerchant(
	creator string,
	index string,
	name string,
	status string,
	registeredAt int64,

) *MsgUpdateMerchant {
	return &MsgUpdateMerchant{
		Creator:      creator,
		Index:        index,
		Name:         name,
		Status:       status,
		RegisteredAt: registeredAt,
	}
}

func NewMsgDeleteMerchant(
	creator string,
	index string,

) *MsgDeleteMerchant {
	return &MsgDeleteMerchant{
		Creator: creator,
		Index:   index,
	}
}
