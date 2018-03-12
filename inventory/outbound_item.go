package inventory

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/julienschmidt/httprouter"
	"github.com/mholt/binding"
	"github.com/mistikel/inventoy/errors"
	"github.com/mistikel/inventoy/model"
)

type AddOutboundItemForm struct {
	Notes string `valid:"required"`
	Item  []AddOrderItemForm
}

type AddOrderItemForm struct {
	Id         int
	ItemId     int `valid:"required"`
	BatchId    int
	SellAmount int    `valid:"required"`
	Price      int    `valid:"required"`
	Total      int    `valid:"required"`
	Notes      string `valid:"required"`
}

func (aif *AddOutboundItemForm) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&aif.Notes: "notes",
		&aif.Item:  "item",
	}
}

func (inventoryModule *InventoryModule) GetOutboundItems(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := context.Background()
	itemDatamodel := model.NewOutboundItemModel(ctx)

	items, err := itemDatamodel.GetMany(ctx)
	if err != nil {
		errors.InternalServer(ctx, w, err)
		return
	}
	res, _ := json.Marshal(items)

	w.Write(res)
}

func (inventoryModule *InventoryModule) GetOutboundItem(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx := context.Background()
	itemDatamodel := model.NewOutboundItemModel(ctx)

	id, _ := strconv.Atoi(p.ByName("id"))

	items, err := itemDatamodel.Get(ctx, id)
	if err != nil {
		errors.DataNotFound(ctx, w)
		return
	}

	res, _ := json.Marshal(items)
	w.Write(res)
}

func (inventoryModule *InventoryModule) StoreOutboundItem(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := context.Background()

	addOutboundItemForm := new(AddOutboundItemForm)
	if err := binding.Bind(r, addOutboundItemForm); err != nil {
		log.Println("Binding Error : ", err)
		errors.InternalServer(ctx, w, err)
		return
	}

	valid, err := govalidator.ValidateStruct(addOutboundItemForm)
	if !valid && err != nil {
		errors.InternalServer(ctx, w, err)
		return
	}

	itemDatamodel := model.NewOutboundItemModel(ctx)

	outboundItem := model.OutboundItem{
		Notes: addOutboundItemForm.Notes,
	}

	err = itemDatamodel.Store(ctx, &outboundItem)
	if err != nil {
		errors.InternalServer(ctx, w, err)
		return
	}

	stockBatchModel := model.NewStockBatchodel(ctx)
	orderItemModel := model.NewOrderItemModel(ctx)
	itemModel := model.NewItemModel(ctx)
	for _, order := range addOutboundItemForm.Item {
		stock, err := stockBatchModel.GetItemStock(ctx, order.ItemId)
		if err != nil && stock.Id == 0 {
			log.Println("Stock not available", err)
		}
		// update stock
		stock.Stock -= order.SellAmount
		stockBatchModel.Update(ctx, stock.Id, stock)

		orderItem := model.OrderItem{
			ItemID:     order.ItemId,
			OutboundID: outboundItem.Id,
			BatchID:    stock.Id,
			SellAmount: order.SellAmount,
			Price:      order.Price,
			Total:      order.Total,
			Notes: 		order.Notes,
		}
		// update stock item
		item, _ := itemModel.Get(ctx, order.ItemId)
		item.Stock -= order.SellAmount
		itemModel.Update(ctx, item.Id, item)
		// order item
		orderItemModel.Store(ctx, orderItem)
	}

	w.Write([]byte(`{ "Status" : "Ok" }`))
}

func (inventoryModule *InventoryModule) PutOutboundItem(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx := context.Background()

	addOutboundItemForm := new(AddOutboundItemForm)
	if err := binding.Bind(r, addOutboundItemForm); err != nil {
		errors.InternalServer(ctx, w, err)
		return
	}

	valid, err := govalidator.ValidateStruct(addOutboundItemForm)
	if !valid && err != nil {
		errors.InternalServer(ctx, w, err)
		return
	}

	id, _ := strconv.Atoi(p.ByName("id"))

	itemDatamodel := model.NewOutboundItemModel(ctx)
	outboundItem := model.OutboundItem{
		Notes: addOutboundItemForm.Notes,
	}

	err = itemDatamodel.Update(ctx, id, outboundItem)
	if err != nil {
		errors.InternalServer(ctx, w, err)
		return
	}

	stockBatchModel := model.NewStockBatchodel(ctx)
	orderItemModel := model.NewOrderItemModel(ctx)
	for _, order := range addOutboundItemForm.Item {
		stock, err := stockBatchModel.Get(ctx, order.BatchId)
		if err != nil && stock.Id == 0 {
			log.Println("Stock not available", err)
		}
		orderItem, _ := orderItemModel.Get(ctx, order.Id)
		// update stock
		stock.Stock += orderItem.SellAmount
		stock.Stock += order.SellAmount
		stockBatchModel.Update(ctx, stock.Id, stock)

		//update order item
		orderItem.SellAmount = order.SellAmount
		orderItem.Price = order.Price
		orderItem.Total = order.Total
		orderItem.Notes = order.Notes
		orderItemModel.Update(ctx, orderItem.Id, orderItem)
	}

	w.Write([]byte(`{ "Status" : "Ok" }`))
}

func (inventoryModule *InventoryModule) RemoveOutboundItem(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx := context.Background()
	itemDatamodel := model.NewOutboundItemModel(ctx)

	id, _ := strconv.Atoi(p.ByName("id"))

	err := itemDatamodel.Delete(ctx, id)
	if err != nil {
		errors.DataNotFound(ctx, w)
		return
	}

	w.Write([]byte(`{ "Status" : "Ok" }`))
}
