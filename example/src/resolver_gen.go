package src

import (
	"github.com/sj-distributor/dolphin-example/gen"
)

func NewResolver(db *gen.DB, ec *gen.EventController) *Resolver {
	handlers := gen.DefaultResolutionHandlers()
	return &Resolver{&gen.GeneratedResolver{Handlers: handlers, DB: db, EventController: ec}}
}

type Resolver struct {
	*gen.GeneratedResolver
}

type MutationResolver struct {
	*gen.GeneratedMutationResolver
}

type QueryResolver struct {
	*gen.GeneratedQueryResolver
}

func (r *Resolver) Mutation() gen.MutationResolver {
	return &MutationResolver{&gen.GeneratedMutationResolver{GeneratedResolver: r.GeneratedResolver}}
}
func (r *Resolver) Query() gen.QueryResolver {
	return &QueryResolver{&gen.GeneratedQueryResolver{GeneratedResolver: r.GeneratedResolver}}
}

type OrderResultTypeResolver struct {
	*gen.GeneratedOrderResultTypeResolver
}

func (r *Resolver) OrderResultType() gen.OrderResultTypeResolver {
	return &OrderResultTypeResolver{&gen.GeneratedOrderResultTypeResolver{GeneratedResolver: r.GeneratedResolver}}
}

type ShipmentResultTypeResolver struct {
	*gen.GeneratedShipmentResultTypeResolver
}

func (r *Resolver) ShipmentResultType() gen.ShipmentResultTypeResolver {
	return &ShipmentResultTypeResolver{&gen.GeneratedShipmentResultTypeResolver{GeneratedResolver: r.GeneratedResolver}}
}

type ShipmentResolver struct {
	*gen.GeneratedShipmentResolver
}

func (r *Resolver) Shipment() gen.ShipmentResolver {
	return &ShipmentResolver{&gen.GeneratedShipmentResolver{GeneratedResolver: r.GeneratedResolver}}
}

type CarrierResultTypeResolver struct {
	*gen.GeneratedCarrierResultTypeResolver
}

func (r *Resolver) CarrierResultType() gen.CarrierResultTypeResolver {
	return &CarrierResultTypeResolver{&gen.GeneratedCarrierResultTypeResolver{GeneratedResolver: r.GeneratedResolver}}
}

type LocationResultTypeResolver struct {
	*gen.GeneratedLocationResultTypeResolver
}

func (r *Resolver) LocationResultType() gen.LocationResultTypeResolver {
	return &LocationResultTypeResolver{&gen.GeneratedLocationResultTypeResolver{GeneratedResolver: r.GeneratedResolver}}
}

type LocationResolver struct {
	*gen.GeneratedLocationResolver
}

func (r *Resolver) Location() gen.LocationResolver {
	return &LocationResolver{&gen.GeneratedLocationResolver{GeneratedResolver: r.GeneratedResolver}}
}

type EquipmentdResultTypeResolver struct {
	*gen.GeneratedEquipmentdResultTypeResolver
}

func (r *Resolver) EquipmentdResultType() gen.EquipmentdResultTypeResolver {
	return &EquipmentdResultTypeResolver{&gen.GeneratedEquipmentdResultTypeResolver{GeneratedResolver: r.GeneratedResolver}}
}
