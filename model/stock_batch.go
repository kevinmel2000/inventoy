package model

import "context"

type StockBatch struct {
	Id        int
	ItemID    int
	InboundId int
	Price     int
	Stock     int
	Data
	Item Item
}

type StockBatchDatanmodel struct{}

func NewStockBatchodel(ctx context.Context) *StockBatchDatanmodel {
	return &StockBatchDatanmodel{}
}

func (stockBatchDatanmodel StockBatchDatanmodel) GetMany(ctx context.Context) ([]StockBatch, error) {
	db := initDB()
	defer db.Close()

	var stockBatchs []StockBatch
	err := db.Find(&stockBatchs).Error

	return stockBatchs, err
}

func (stockBatchDatanmodel StockBatchDatanmodel) Get(ctx context.Context, id int) (StockBatch, error) {
	db := initDB()
	defer db.Close()

	var stockBatch StockBatch
	err := db.Find(&stockBatch, id).Error

	return stockBatch, err
}

func (stockBatchDatanmodel StockBatchDatanmodel) Store(ctx context.Context, stockBatch StockBatch) error {
	db := initDB()
	defer db.Close()

	err := db.Create(&stockBatch).Error

	return err
}

func (stockBatchDatanmodel StockBatchDatanmodel) Update(ctx context.Context, id int, newStockBatch StockBatch) error {
	db := initDB()
	defer db.Close()

	var stockBatch StockBatch
	err := db.First(&stockBatch, id).Error
	if err == nil {
		db.Save(newStockBatch)
	} else {
		return err
	}

	return nil
}

func (stockBatchDatanmodel StockBatchDatanmodel) Delete(ctx context.Context, id int) error {
	db := initDB()
	defer db.Close()

	var stockBatch StockBatch
	err := db.First(&stockBatch, id).Error
	if err == nil {
		db.Delete(stockBatch)
	} else {
		return err
	}

	return nil
}

func (stockBatchDatanmodel StockBatchDatanmodel) GetItemStock(ctx context.Context, itemID int) (StockBatch, error) {
	db := initDB()
	defer db.Close()

	var stockBatch StockBatch
	err := db.Where("stock > 0").Where("item_id = ?", itemID).Order("created_at asc").Find(&stockBatch).Error

	return stockBatch, err
}
