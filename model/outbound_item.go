package model

import "context"

type OutboundItem struct {
	Id    int    `json:"id"`
	Notes string `json:"notes"`
	Data
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
