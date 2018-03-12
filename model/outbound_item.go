package model

import (
	"context"
	"time"
)

type OutboundItem struct {
	Id    int    `json:"id"`
	Notes string `json:"notes"`
	Data
}

type SellingReport struct {
	Id           int
	Time         time.Time
	Sku          string
	Name         string
	TotalItem    int
	SellingPrice int
	TotalPrice   int
	BuyingPrice  int
	Laba         int
}

type OutboundItemDatamodel struct{}

func NewOutboundItemModel(ctx context.Context) *OutboundItemDatamodel {
	return &OutboundItemDatamodel{}
}

func (outboundItemDatamodel OutboundItemDatamodel) GetMany(ctx context.Context) ([]OutboundItem, error) {
	db := initDB()
	defer db.Close()

	var items []OutboundItem
	err := db.Find(&items).Error

	return items, err
}

func (outboundItemDatamodel OutboundItemDatamodel) Get(ctx context.Context, id int) (OutboundItem, error) {
	db := initDB()
	defer db.Close()

	var item OutboundItem
	err := db.Find(&item, id).Error

	return item, err
}

func (outboundItemDatamodel OutboundItemDatamodel) Store(ctx context.Context, item *OutboundItem) error {
	db := initDB()
	defer db.Close()

	err := db.Create(&item).Error

	return err
}

func (outboundItemDatamodel OutboundItemDatamodel) Update(ctx context.Context, id int, newItem OutboundItem) error {
	db := initDB()
	defer db.Close()

	var item OutboundItem
	err := db.First(&item, id).Error
	if err == nil {
		db.Save(newItem)
	} else {
		return err
	}

	return nil
}

func (outboundItemDatamodel OutboundItemDatamodel) Delete(ctx context.Context, id int) error {
	db := initDB()
	defer db.Close()

	var item OutboundItem
	err := db.First(&item, id).Error
	if err == nil {
		db.Delete(item)
	} else {
		return err
	}

	return nil
}

func (outboundItemDatamodel OutboundItemDatamodel) GetRecord(ctx context.Context, dt1 string, dt2 string) ([]SellingReport, int, int, int, int) {
	db := initDB()
	defer db.Close()
	var omset, laba, selling, totalItem int
	var outbounds []OutboundItem
	db.Where("created_at >= ? ", dt1).Where("created_at < ? ", dt2).Find(&outbounds)
	itemDatamodel := NewItemModel(ctx)
	orderItemDatamodel := NewOrderItemModel(ctx)
	batchDatamodel := NewStockBatchodel(ctx)
	var r SellingReport
	var record []SellingReport
	for _, outbound := range outbounds {
		orderItems, _ := orderItemDatamodel.GetByOutbound(ctx, outbound.Id)
		for _, orderItem := range orderItems {
			item, _ := itemDatamodel.Get(ctx, orderItem.ItemID)
			batch, _ := batchDatamodel.Get(ctx, orderItem.BatchID)
			r.Id = outbound.Id
			r.Time = outbound.CreatedAt
			r.Sku = item.Sku
			r.Name = item.Name
			r.SellingPrice = orderItem.Price
			r.BuyingPrice = batch.Price
			r.TotalItem = orderItem.SellAmount
			r.TotalPrice = orderItem.Total
			r.Laba = r.SellingPrice - r.BuyingPrice

			omset += orderItem.Total
			laba += r.Laba
			totalItem += orderItem.SellAmount

			record = append(record, r)
		}

		selling += 1

	}

	return record, omset, laba, selling, totalItem

}
