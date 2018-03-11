package model

type Item struct {
	Id    int    `json:"id"`
	Sku   string `json:"sku"`
	Name  string `json:"name"`
	Stock int    `json:"stock"`
}
