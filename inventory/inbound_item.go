package inventory

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/julienschmidt/httprouter"
	"github.com/mholt/binding"
	"github.com/mistikel/inventoy/errors"
	"github.com/mistikel/inventoy/model"
)

type AddInboundItemForm struct {
	ItemId         int    `valid:"required"`
	Status         int    `valid:"required"`
	OrderAmount    int    `valid:"required"`
	ReceivedAmount int    `valid:"required"`
	Price          int    `valid:"required"`
	Total          int    `valid:"required"`
	ReceiptNumber  string `valid:"required"`
	Notes          string `valid:"required"`
}

func (aif *AddInboundItemForm) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&aif.ItemId:        "item_id",
		&aif.Status:        "status",
		&aif.OrderAmount:   "order_amount",
		&aif.ReceiptNumber: "received_amount",
		&aif.Price:         "price",
		&aif.Total:         "total",
		&aif.ReceiptNumber: "receipt_number",
		&aif.Notes:         "notes",
	}
}

func (inventoryModule *InventoryModule) GetInboundItems(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := context.Background()
	itemDatamodel := model.NewInboundItemModel(ctx)

	items, err := itemDatamodel.GetMany(ctx)
	if err != nil {
		errors.InternalServer(ctx, w, err)
		return
	}
	res, _ := json.Marshal(items)

	w.Write(res)
}

func (inventoryModule *InventoryModule) GetInboundItem(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx := context.Background()
	itemDatamodel := model.NewInboundItemModel(ctx)

	id, _ := strconv.Atoi(p.ByName("id"))

	items, err := itemDatamodel.Get(ctx, id)
	if err != nil {
		errors.DataNotFound(ctx, w)
		return
	}

	res, _ := json.Marshal(items)
	w.Write(res)
}

func (inventoryModule *InventoryModule) StoreInboundItem(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := context.Background()

	addInboundItemForm := new(AddInboundItemForm)
	if err := binding.Bind(r, addInboundItemForm); err != nil {
		errors.InternalServer(ctx, w, err)
		return
	}

	valid, err := govalidator.ValidateStruct(addInboundItemForm)
	if !valid && err != nil {
		errors.InternalServer(ctx, w, err)
		return
	}

	itemDatamodel := model.NewInboundItemModel(ctx)

	inboundItem := model.InboundItem{
		ItemId:         addInboundItemForm.ItemId,
		Status:         addInboundItemForm.Status,
		OrderAmount:    addInboundItemForm.OrderAmount,
		ReceivedAmount: addInboundItemForm.ReceivedAmount,
		Price:          addInboundItemForm.Price,
		Total:          addInboundItemForm.Total,
		ReceiptNumber:  addInboundItemForm.ReceiptNumber,
		Notes:          addInboundItemForm.Notes,
	}

	err = itemDatamodel.Store(ctx, inboundItem)

	if err != nil {
		errors.InternalServer(ctx, w, err)
		return
	}

	w.Write([]byte(`{ "Status" : "Ok" }`))
}

func (inventoryModule *InventoryModule) PutInboundItem(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx := context.Background()

	addInboundItemForm := new(AddInboundItemForm)
	if err := binding.Bind(r, addInboundItemForm); err != nil {
		errors.InternalServer(ctx, w, err)
		return
	}

	valid, err := govalidator.ValidateStruct(addInboundItemForm)
	if !valid && err != nil {
		errors.InternalServer(ctx, w, err)
		return
	}

	itemDatamodel := model.NewInboundItemModel(ctx)

	id, _ := strconv.Atoi(p.ByName("id"))
	inboundItem := model.InboundItem{
		ItemId:         addInboundItemForm.ItemId,
		Status:         addInboundItemForm.Status,
		OrderAmount:    addInboundItemForm.OrderAmount,
		ReceivedAmount: addInboundItemForm.ReceivedAmount,
		Price:          addInboundItemForm.Price,
		Total:          addInboundItemForm.Total,
		ReceiptNumber:  addInboundItemForm.ReceiptNumber,
		Notes:          addInboundItemForm.Notes,
	}

	err = itemDatamodel.Update(ctx, id, inboundItem)
	if err != nil {
		errors.DataNotFound(ctx, w)
		return
	}

	w.Write([]byte(`{ "Status" : "Ok" }`))
}

func (inventoryModule *InventoryModule) RemoveInboundItem(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx := context.Background()
	itemDatamodel := model.NewInboundItemModel(ctx)

	id, _ := strconv.Atoi(p.ByName("id"))

	err := itemDatamodel.Delete(ctx, id)
	if err != nil {
		errors.DataNotFound(ctx, w)
		return
	}

	w.Write([]byte(`{ "Status" : "Ok" }`))
}
