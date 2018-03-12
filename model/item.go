package model

import "context"

type Item struct {
	Id    int    `json:"id"`
	Sku   string `json:"sku"`
	Name  string `json:"name"`
	Stock int    `json:"stock"`
	Data
}

type ReportRecord struct {
	Sku        string
	Name       string
	TotalItem  int
	Avarage    int
	TotalValue int
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

func (itemDatamodel ItemDatamodel) GetTotalItem(ctx context.Context) int {
	db := initSQL()
	defer db.Close()
	var total int
	rows, _ := db.Query("SELECT SUM(received_amount) as total FROM inbound_items")
	for rows.Next() {
		rows.Scan(&total)
	}

	return total
}

func (itemDatamodel ItemDatamodel) GetTotalValue(ctx context.Context) int {
	db := initSQL()
	defer db.Close()
	var total int
	rows, _ := db.Query("SELECT SUM(total) as total FROM inbound_items")
	for rows.Next() {
		rows.Scan(&total)
	}

	return total
}

func (itemDatamodel ItemDatamodel) GetRecord(ctx context.Context) []ReportRecord {
	db := initSQL()
	defer db.Close()
	var record []ReportRecord
	var r ReportRecord
	var sku, name string
	var stock, total int
	rows, _ := db.Query("SELECT sku, name, SUM(received_amount) as stock, SUM(total) as total FROM items JOIN inbound_items on items.id = inbound_items.item_id group by items.id")
	for rows.Next() {
		rows.Scan(&sku, &name, &stock, &total)
		r.Sku = sku
		r.Name = name
		r.TotalItem = stock
		r.TotalValue = total
		r.Avarage = total / stock

		record = append(record, r)
	}

	return record

}
