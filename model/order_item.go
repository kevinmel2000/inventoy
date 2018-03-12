package model

import "context"

type OrderItem struct {
	Id         int    `json:"id"`
	ItemID     int    `json:"id_item"`
	OutboundID int    `json:"id_outbound_item"`
	BatchID    int    `json:"-"`
	SellAmount int    `json:"sell_amount"`
	Price      int    `json:"price"`
	Total      int    `json:"total_price"`
	Notes      string `json:"notes"`
	Data
	Item         Item         `json:"-"`
	OutboundItem OutboundItem `json:"-"`
	StockBatch   StockBatch   `json:"-"`
}

type OrderItemDatamodel struct{}

func NewOrderItemModel(ctx context.Context) *OrderItemDatamodel {
	return &OrderItemDatamodel{}
}

func (orderItemDatamodel OrderItemDatamodel) GetMany(ctx context.Context) ([]OrderItem, error) {
	db := initDB()
	defer db.Close()

	var order []OrderItem
	err := db.Find(&order).Error

	return order, err
}

func (orderItemDatamodel OrderItemDatamodel) Get(ctx context.Context, id int) (OrderItem, error) {
	db := initDB()
	defer db.Close()

	var order OrderItem
	err := db.Find(&order, id).Error

	return order, err
}

func (orderItemDatamodel OrderItemDatamodel) Store(ctx context.Context, order OrderItem) error {
	db := initDB()
	defer db.Close()

	err := db.Create(&order).Error

	return err
}

func (orderItemDatamodel OrderItemDatamodel) Update(ctx context.Context, id int, newOrder OrderItem) error {
	db := initDB()
	defer db.Close()

	var order OrderItem
	err := db.First(&order, id).Error
	if err == nil {
		db.Save(newOrder)
	} else {
		return err
	}

	return nil
}

func (orderItemDatamodel OrderItemDatamodel) GetByOutbound(ctx context.Context, OutboundID int) (OrderItem, error) {
	db := initDB()
	defer db.Close()

	var order OrderItem
	err := db.Where("outbound_id = ?", OutboundID).Find(&order).Error

	return order, err
}
