package model

import (
	"context"
)

type InboundItem struct {
	Id             int    `json:"id"`
	ItemID         int    `json:"id_item"`
	Status         int    `json:"status"`
	OrderAmount    int    `json:"order_amount"`
	ReceivedAmount int    `json:"received_amount"`
	Price          int    `json:"price"`
	Total          int    `json:"total_price"`
	ReceiptNumber  string `json:"receipt_number"`
	Notes          string `json:"notes"`
	Data
	Item       Item
	StockBatch StockBatch
}

type InboundItemDatamodel struct{}

func NewInboundItemModel(ctx context.Context) *InboundItemDatamodel {
	return &InboundItemDatamodel{}
}

func (inboundItemDatamodel InboundItemDatamodel) GetMany(ctx context.Context) ([]InboundItem, error) {
	db := initDB()
	defer db.Close()

	var items []InboundItem
	err := db.Find(&items).Error

	return items, err
}

func (inboundItemDatamodel InboundItemDatamodel) Get(ctx context.Context, id int) (InboundItem, error) {
	db := initDB()
	defer db.Close()

	var item InboundItem
	err := db.Find(&item, id).Error

	return item, err
}

func (inboundItemDatamodel InboundItemDatamodel) Store(ctx context.Context, item *InboundItem) error {
	db := initDB()
	defer db.Close()

	err := db.Create(&item).Error
	return err
}

func (inboundItemDatamodel InboundItemDatamodel) Update(ctx context.Context, id int, newItem InboundItem) error {
	db := initDB()
	defer db.Close()

	var item InboundItem
	err := db.First(&item, id).Error
	if err == nil {
		db.Save(newItem)
	} else {
		return err
	}

	return nil
}

func (inboundItemDatamodel InboundItemDatamodel) Delete(ctx context.Context, id int) error {
	db := initDB()
	defer db.Close()

	var item InboundItem
	err := db.First(&item, id).Error
	if err == nil {
		db.Delete(item)
	} else {
		return err
	}

	return nil
}
