package model

type OutboundItem struct {
	Id         int    `json:"id"`
	ItemId     int    `json:"id_item"`
	SellAmount int    `json:"sell_amount"`
	Price      int    `json:"price"`
	Total      int    `json:"total_price"`
	Notes      string `json:"notes"`
	Item       Item
}
