package model

import "context"

type Item struct {
	Id    int    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Sku   string `gorm:"type:varchar(100);unique" json:"sku"`
	Name  string `gorm:"not null" json:"name"`
	Stock int    `gorm:"not null" json:"stock"`
	Data
}

type ItemDatamodel struct{}

func NewItemModel(ctx context.Context) *ItemDatamodel {
	return &ItemDatamodel{}
}

func (itemDatamodel ItemDatamodel) GetMany(ctx context.Context) ([]Item, error) {
	db := initDB()
	defer db.Close()

	var items []Item
	err := db.Find(&items).Error

	return items, err
}

func (itemDatamodel ItemDatamodel) Get(ctx context.Context, id int) (Item, error) {
	db := initDB()
	defer db.Close()

	var item Item
	err := db.Find(&item, id).Error

	return item, err
}

func (itemDatamodel ItemDatamodel) Store(ctx context.Context, item Item) error {
	db := initDB()
	defer db.Close()

	err := db.Create(&item).Error

	return err
}

func (itemDatamodel ItemDatamodel) Update(ctx context.Context, id int, newItem Item) error {
	db := initDB()
	defer db.Close()

	var item Item
	err := db.First(&item, id).Error
	if err == nil {
		db.Save(newItem)
	} else {
		return err
	}

	return nil
}

func (itemDatamodel ItemDatamodel) Delete(ctx context.Context, id int) error {
	db := initDB()
	defer db.Close()

	var item Item
	err := db.First(&item, id).Error
	if err == nil {
		db.Delete(item)
	} else {
		return err
	}

	return nil
}

func (itemDatamodel ItemDatamodel) GetTotalSKU(ctx context.Context) int {
	db := initDB()
	defer db.Close()

	return 1
}
