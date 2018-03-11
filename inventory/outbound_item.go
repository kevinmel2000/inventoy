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

type AddOutboundItemForm struct {
	ItemId     int    `valid:"required"`
	SellAmount int    `valid:"required"`
	Price      int    `valid:"required"`
	Total      int    `valid:"required"`
	Notes      string `valid:"required"`
}

func (aif *AddOutboundItemForm) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&aif.ItemId:     "item_id",
		&aif.SellAmount: "sell_amount",
		&aif.Price:      "price",
		&aif.Total:      "total",
		&aif.Notes:      "notes",
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
		ItemId:     addOutboundItemForm.ItemId,
		SellAmount: addOutboundItemForm.SellAmount,
		Price:      addOutboundItemForm.Price,
		Total:      addOutboundItemForm.Total,
		Notes:      addOutboundItemForm.Notes,
	}

	err = itemDatamodel.Store(ctx, outboundItem)

	if err != nil {
		errors.InternalServer(ctx, w, err)
		return
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

	itemDatamodel := model.NewOutboundItemModel(ctx)

	outboundItem := model.OutboundItem{
		ItemId:     addOutboundItemForm.ItemId,
		SellAmount: addOutboundItemForm.SellAmount,
		Price:      addOutboundItemForm.Price,
		Total:      addOutboundItemForm.Total,
		Notes:      addOutboundItemForm.Notes,
	}

	id, _ := strconv.Atoi(p.ByName("id"))

	err = itemDatamodel.Update(ctx, id, outboundItem)
	if err != nil {
		errors.DataNotFound(ctx, w)
		return
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
