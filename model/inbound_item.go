package model

type InboundItem struct {
	Id             int    `json:"id"`
	ItemId         int    `json:"id_item"`
	Status         int    `json:"status"`
	OrderAmount    int    `json:"order_amount"`
	ReceivedAmount int    `json:"received_amount"`
	Price          int    `json:"price"`
	Total          int    `json:"total_price"`
	ReceiptNumber  string `json:"receipt_number"`
	Notes          string `json:"notes"`
	Item           Item
}
