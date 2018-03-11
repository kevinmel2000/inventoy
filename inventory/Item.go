package inventory

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"

	"github.com/mholt/binding"

	"github.com/julienschmidt/httprouter"
	"github.com/mistikel/inventoy/errors"
	"github.com/mistikel/inventoy/model"
)

type AddItemForm struct {
	Sku   string `valid:"required"`
	Name  string `valid:"required"`
	Stock int    `valid:"required"`
}

func (aif *AddItemForm) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&aif.Name:  "name",
		&aif.Sku:   "sku",
		&aif.Stock: "stock",
	}
}

func (inventoryModule *InventoryModule) GetItems(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := context.Background()
	itemDatamodel := model.NewItemModel(ctx)

	items, err := itemDatamodel.GetMany(ctx)
	if err != nil {
		errors.InternalServer(ctx, w, err)
		return
	}
	res, _ := json.Marshal(items)

	w.Write(res)
}

func (inventoryModule *InventoryModule) GetItem(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx := context.Background()
	itemDatamodel := model.NewItemModel(ctx)

	id, _ := strconv.Atoi(p.ByName("id"))

	items, err := itemDatamodel.Get(ctx, id)
	if err != nil {
		errors.DataNotFound(ctx, w)
		return
	}

	res, _ := json.Marshal(items)
	w.Write(res)
}

func (inventoryModule *InventoryModule) StoreItem(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := context.Background()

	addItemForm := new(AddItemForm)
	if err := binding.Bind(r, addItemForm); err != nil {
		errors.InternalServer(ctx, w, err)
		return
	}

	valid, err := govalidator.ValidateStruct(addItemForm)
	if !valid && err != nil {
		errors.InternalServer(ctx, w, err)
		return
	}

	itemDatamodel := model.NewItemModel(ctx)

	item := model.Item{
		Name:  addItemForm.Name,
		Sku:   addItemForm.Sku,
		Stock: addItemForm.Stock,
	}
	itemDatamodel.Store(ctx, item)

	w.Write([]byte(`{ "Status" : "Ok" }`))
}

func (inventoryModule *InventoryModule) PutItem(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx := context.Background()

	addItemForm := new(AddItemForm)
	if err := binding.Bind(r, addItemForm); err != nil {
		errors.InternalServer(ctx, w, err)
		return
	}

	valid, err := govalidator.ValidateStruct(addItemForm)
	if !valid && err != nil {
		errors.InternalServer(ctx, w, err)
		return
	}

	itemDatamodel := model.NewItemModel(ctx)

	id, _ := strconv.Atoi(p.ByName("id"))
	item := model.Item{
		Id:    id,
		Name:  addItemForm.Name,
		Sku:   addItemForm.Sku,
		Stock: addItemForm.Stock,
	}

	err = itemDatamodel.Update(ctx, id, item)
	if err != nil {
		errors.DataNotFound(ctx, w)
		return
	}

	w.Write([]byte(`{ "Status" : "Ok" }`))
}

func (inventoryModule *InventoryModule) RemoveItem(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx := context.Background()
	itemDatamodel := model.NewItemModel(ctx)

	id, _ := strconv.Atoi(p.ByName("id"))

	err := itemDatamodel.Delete(ctx, id)
	if err != nil {
		errors.DataNotFound(ctx, w)
		return
	}

	w.Write([]byte(`{ "Status" : "Ok" }`))
}
