package model

import "context"

type StockBatch struct {
	Id        int
	ItemID    int
	InboundId int
	Price     int
	Stock     int
	Item      Item
}

type StockBatchDatanmodel struct{}

func NewStockBatchodel(ctx context.Context) *StockBatchDatanmodel {
	return &StockBatchDatanmodel{}
}

func (itockBatchDatanmodel StockBatchDatanmodel) GetMany(ctx context.Context) ([]StockBatch, error) {
	db := initDB()
	defer db.Close()

	var stockBatchs []StockBatch
	err := db.Find(&stockBatchs).Error

	return stockBatchs, err
}

func (itockBatchDatanmodel StockBatchDatanmodel) Get(ctx context.Context, id int) (StockBatch, error) {
	db := initDB()
	defer db.Close()

	var stockBatch StockBatch
	err := db.Find(&stockBatch, id).Error

	return stockBatch, err
}

func (itockBatchDatanmodel StockBatchDatanmodel) Store(ctx context.Context, stockBatch StockBatch) error {
	db := initDB()
	defer db.Close()

	err := db.Create(&stockBatch).Error

	return err
}

func (itockBatchDatanmodel StockBatchDatanmodel) Update(ctx context.Context, id int, newStockBatch StockBatch) error {
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

func (itockBatchDatanmodel StockBatchDatanmodel) Delete(ctx context.Context, id int) error {
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
