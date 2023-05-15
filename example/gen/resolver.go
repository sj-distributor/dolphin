//go:generate go run github.com/99designs/gqlgen generate
package gen

import (
	"context"
)

type ResolutionHandlers struct {
	OnEvent func(ctx context.Context, r *GeneratedResolver, e *Event) error

	CreateOrder    func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Order, err error)
	UpdateOrder    func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Order, err error)
	DeleteOrders   func(ctx context.Context, r *GeneratedResolver, id []string, unscoped *bool) (bool, error)
	RecoveryOrders func(ctx context.Context, r *GeneratedResolver, id []string) (bool, error)
	QueryOrder     func(ctx context.Context, r *GeneratedResolver, id string) (*Order, error)
	QueryOrders    func(ctx context.Context, r *GeneratedResolver, opts QueryOrdersHandlerOptions) (*OrderResultType, error)

	CreateShipment    func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Shipment, err error)
	UpdateShipment    func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Shipment, err error)
	DeleteShipments   func(ctx context.Context, r *GeneratedResolver, id []string, unscoped *bool) (bool, error)
	RecoveryShipments func(ctx context.Context, r *GeneratedResolver, id []string) (bool, error)
	QueryShipment     func(ctx context.Context, r *GeneratedResolver, id string) (*Shipment, error)
	QueryShipments    func(ctx context.Context, r *GeneratedResolver, opts QueryShipmentsHandlerOptions) (*ShipmentResultType, error)

	ShipmentStartLocation func(ctx context.Context, r *GeneratedResolver, obj *Shipment) (res *Location, err error)

	ShipmentEndLocation func(ctx context.Context, r *GeneratedResolver, obj *Shipment) (res *Location, err error)

	CreateCarrier    func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Carrier, err error)
	UpdateCarrier    func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Carrier, err error)
	DeleteCarriers   func(ctx context.Context, r *GeneratedResolver, id []string, unscoped *bool) (bool, error)
	RecoveryCarriers func(ctx context.Context, r *GeneratedResolver, id []string) (bool, error)
	QueryCarrier     func(ctx context.Context, r *GeneratedResolver, id string) (*Carrier, error)
	QueryCarriers    func(ctx context.Context, r *GeneratedResolver, opts QueryCarriersHandlerOptions) (*CarrierResultType, error)

	CreateLocation    func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Location, err error)
	UpdateLocation    func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Location, err error)
	DeleteLocations   func(ctx context.Context, r *GeneratedResolver, id []string, unscoped *bool) (bool, error)
	RecoveryLocations func(ctx context.Context, r *GeneratedResolver, id []string) (bool, error)
	QueryLocation     func(ctx context.Context, r *GeneratedResolver, id string) (*Location, error)
	QueryLocations    func(ctx context.Context, r *GeneratedResolver, opts QueryLocationsHandlerOptions) (*LocationResultType, error)

	LocationStartShipments func(ctx context.Context, r *GeneratedResolver, obj *Location) (res []*Shipment, err error)

	LocationEndShipments func(ctx context.Context, r *GeneratedResolver, obj *Location) (res []*Shipment, err error)

	CreateEquipmentd    func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Equipmentd, err error)
	UpdateEquipmentd    func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Equipmentd, err error)
	DeleteEquipmentds   func(ctx context.Context, r *GeneratedResolver, id []string, unscoped *bool) (bool, error)
	RecoveryEquipmentds func(ctx context.Context, r *GeneratedResolver, id []string) (bool, error)
	QueryEquipmentd     func(ctx context.Context, r *GeneratedResolver, id string) (*Equipmentd, error)
	QueryEquipmentds    func(ctx context.Context, r *GeneratedResolver, opts QueryEquipmentdsHandlerOptions) (*EquipmentdResultType, error)
}

func DefaultResolutionHandlers() ResolutionHandlers {
	handlers := ResolutionHandlers{
		OnEvent: func(ctx context.Context, r *GeneratedResolver, e *Event) error { return nil },

		CreateOrder:    CreateOrderHandler,
		UpdateOrder:    UpdateOrderHandler,
		DeleteOrders:   DeleteOrdersHandler,
		RecoveryOrders: RecoveryOrdersHandler,
		QueryOrder:     QueryOrderHandler,
		QueryOrders:    QueryOrdersHandler,

		CreateShipment:    CreateShipmentHandler,
		UpdateShipment:    UpdateShipmentHandler,
		DeleteShipments:   DeleteShipmentsHandler,
		RecoveryShipments: RecoveryShipmentsHandler,
		QueryShipment:     QueryShipmentHandler,
		QueryShipments:    QueryShipmentsHandler,

		ShipmentStartLocation: ShipmentStartLocationHandler,

		ShipmentEndLocation: ShipmentEndLocationHandler,

		CreateCarrier:    CreateCarrierHandler,
		UpdateCarrier:    UpdateCarrierHandler,
		DeleteCarriers:   DeleteCarriersHandler,
		RecoveryCarriers: RecoveryCarriersHandler,
		QueryCarrier:     QueryCarrierHandler,
		QueryCarriers:    QueryCarriersHandler,

		CreateLocation:    CreateLocationHandler,
		UpdateLocation:    UpdateLocationHandler,
		DeleteLocations:   DeleteLocationsHandler,
		RecoveryLocations: RecoveryLocationsHandler,
		QueryLocation:     QueryLocationHandler,
		QueryLocations:    QueryLocationsHandler,

		LocationStartShipments: LocationStartShipmentsHandler,

		LocationEndShipments: LocationEndShipmentsHandler,

		CreateEquipmentd:    CreateEquipmentdHandler,
		UpdateEquipmentd:    UpdateEquipmentdHandler,
		DeleteEquipmentds:   DeleteEquipmentdsHandler,
		RecoveryEquipmentds: RecoveryEquipmentdsHandler,
		QueryEquipmentd:     QueryEquipmentdHandler,
		QueryEquipmentds:    QueryEquipmentdsHandler,
	}
	return handlers
}

type GeneratedResolver struct {
	Handlers        ResolutionHandlers
	DB              *DB
	EventController *EventController
}
