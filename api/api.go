package api

import (
	"github.com/julienschmidt/httprouter"
	"github.com/mistikel/inventoy/inventory"
)

type API struct {
	Inventory *inventory.InventoryModule
}

func Api() *API {
	return &API{
		Inventory: inventory.NewInventoryModule(),
	}
}

func (api *API) Register(r *httprouter.Router) {
	//item
	r.GET("/items", api.Inventory.GetItems)
	r.GET("/items/:id", api.Inventory.GetItem)
	r.POST("/items", api.Inventory.StoreItem)
	r.PUT("/items/:id", api.Inventory.PutItem)
	r.DELETE("/items/:id", api.Inventory.RemoveItem)

	//inbound item
	r.GET("/inbound_items", api.Inventory.GetInboundItems)
	r.GET("/inbound_items/:id", api.Inventory.GetInboundItem)
	r.POST("/inbound_items", api.Inventory.StoreInboundItem)
	r.PUT("/inbound_items/:id", api.Inventory.PutInboundItem)
	r.DELETE("/inbound_item/:id", api.Inventory.RemoveInboundItem)

	//outbound item
	r.GET("/outbound_items", api.Inventory.GetOutboundItems)
	r.GET("/outbound_items/:id", api.Inventory.GetOutboundItem)
	r.POST("/outbound_items", api.Inventory.StoreOutboundItem)
	r.PUT("/outbound_items/:id", api.Inventory.PutOutboundItem)
	r.DELETE("/outbound_items/:id", api.Inventory.RemoveOutboundItem)
}
