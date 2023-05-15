package gen

import (
	"context"
	"math"

	"github.com/99designs/gqlgen/graphql"
	"github.com/graph-gophers/dataloader"
	"github.com/vektah/gqlparser/v2/ast"
)

type GeneratedQueryResolver struct{ *GeneratedResolver }

func (r *GeneratedQueryResolver) Order(ctx context.Context, id string) (*Order, error) {
	return r.Handlers.QueryOrder(ctx, r.GeneratedResolver, id)
}
func QueryOrderHandler(ctx context.Context, r *GeneratedResolver, id string) (*Order, error) {
	selection := []ast.Selection{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		selection = append(selection, f.Field)
	}
	selectionSet := ast.SelectionSet(selection)

	query := OrderQueryFilter{}
	rt := &OrderResultType{
		EntityResultType: EntityResultType{
			Query:        &query,
			SelectionSet: &selectionSet,
		},
	}
	qb := r.DB.Query()
	qb = qb.Where(TableName("orders")+".id = ?", id)

	var items []*Order
	giOpts := GetItemsOptions{
		Alias:      TableName("orders"),
		Preloaders: []string{},
		Item:       &Order{},
	}
	err := rt.GetData(ctx, qb, giOpts, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, &NotFoundError{Entity: "Order"}
	}
	return items[0], err
}

type QueryOrdersHandlerOptions struct {
	CurrentPage *int
	PerPage     *int
	Q           *string
	Sort        []*OrderSortType
	Filter      *OrderFilterType
	Rand        *bool
}

func (r *GeneratedQueryResolver) Orders(ctx context.Context, current_page *int, per_page *int, q *string, sort []*OrderSortType, filter *OrderFilterType, rand *bool) (*OrderResultType, error) {
	opts := QueryOrdersHandlerOptions{
		CurrentPage: current_page,
		PerPage:     per_page,
		Q:           q,
		Sort:        sort,
		Filter:      filter,
		Rand:        rand,
	}
	return r.Handlers.QueryOrders(ctx, r.GeneratedResolver, opts)
}
func QueryOrdersHandler(ctx context.Context, r *GeneratedResolver, opts QueryOrdersHandlerOptions) (*OrderResultType, error) {
	query := OrderQueryFilter{opts.Q}

	var selectionSet *ast.SelectionSet
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		if f.Field.Name == "data" {
			selectionSet = &f.Field.SelectionSet
		}
	}

	_sort := []EntitySort{}
	for _, sort := range opts.Sort {
		_sort = append(_sort, sort)
	}

	return &OrderResultType{
		EntityResultType: EntityResultType{
			CurrentPage:  opts.CurrentPage,
			PerPage:      opts.PerPage,
			Rand:         opts.Rand,
			Query:        &query,
			Sort:         _sort,
			Filter:       opts.Filter,
			SelectionSet: selectionSet,
		},
	}, nil
}

type GeneratedOrderResultTypeResolver struct{ *GeneratedResolver }

func (r *GeneratedOrderResultTypeResolver) Data(ctx context.Context, obj *OrderResultType) (items []*Order, err error) {
	giOpts := GetItemsOptions{
		Alias:      TableName("orders"),
		Preloaders: []string{},
		Item:       &Order{},
	}
	err = obj.GetData(ctx, r.DB.db, giOpts, &items)

	uniqueItems := []*Order{}
	idMap := map[string]bool{}
	for _, item := range items {
		if _, ok := idMap[item.ID]; !ok {
			idMap[item.ID] = true
			uniqueItems = append(uniqueItems, item)
		}
	}
	items = uniqueItems

	return
}

func (r *GeneratedOrderResultTypeResolver) Total(ctx context.Context, obj *OrderResultType) (count int, err error) {
	return obj.GetTotal(ctx, r.DB.db, TableName("orders"), &Order{})
}

func (r *GeneratedOrderResultTypeResolver) TotalPage(ctx context.Context, obj *OrderResultType) (count int, err error) {
	total, _ := r.Total(ctx, obj)
	perPage, _ := r.PerPage(ctx, obj)
	totalPage := int(math.Ceil(float64(total) / float64(perPage)))
	if totalPage < 0 {
		totalPage = 0
	}

	return totalPage, nil
}

func (r *GeneratedOrderResultTypeResolver) CurrentPage(ctx context.Context, obj *OrderResultType) (count int, err error) {
	return int(*obj.EntityResultType.CurrentPage), nil
}

func (r *GeneratedOrderResultTypeResolver) PerPage(ctx context.Context, obj *OrderResultType) (count int, err error) {
	return int(*obj.EntityResultType.PerPage), nil
}

func (r *GeneratedQueryResolver) Shipment(ctx context.Context, id string) (*Shipment, error) {
	return r.Handlers.QueryShipment(ctx, r.GeneratedResolver, id)
}
func QueryShipmentHandler(ctx context.Context, r *GeneratedResolver, id string) (*Shipment, error) {
	selection := []ast.Selection{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		selection = append(selection, f.Field)
	}
	selectionSet := ast.SelectionSet(selection)

	query := ShipmentQueryFilter{}
	rt := &ShipmentResultType{
		EntityResultType: EntityResultType{
			Query:        &query,
			SelectionSet: &selectionSet,
		},
	}
	qb := r.DB.Query()
	qb = qb.Where(TableName("shipments")+".id = ?", id)

	var items []*Shipment
	giOpts := GetItemsOptions{
		Alias:      TableName("shipments"),
		Preloaders: []string{},
		Item:       &Shipment{},
	}
	err := rt.GetData(ctx, qb, giOpts, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, &NotFoundError{Entity: "Shipment"}
	}
	return items[0], err
}

type QueryShipmentsHandlerOptions struct {
	CurrentPage *int
	PerPage     *int
	Q           *string
	Sort        []*ShipmentSortType
	Filter      *ShipmentFilterType
	Rand        *bool
}

func (r *GeneratedQueryResolver) Shipments(ctx context.Context, current_page *int, per_page *int, q *string, sort []*ShipmentSortType, filter *ShipmentFilterType, rand *bool) (*ShipmentResultType, error) {
	opts := QueryShipmentsHandlerOptions{
		CurrentPage: current_page,
		PerPage:     per_page,
		Q:           q,
		Sort:        sort,
		Filter:      filter,
		Rand:        rand,
	}
	return r.Handlers.QueryShipments(ctx, r.GeneratedResolver, opts)
}
func QueryShipmentsHandler(ctx context.Context, r *GeneratedResolver, opts QueryShipmentsHandlerOptions) (*ShipmentResultType, error) {
	query := ShipmentQueryFilter{opts.Q}

	var selectionSet *ast.SelectionSet
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		if f.Field.Name == "data" {
			selectionSet = &f.Field.SelectionSet
		}
	}

	_sort := []EntitySort{}
	for _, sort := range opts.Sort {
		_sort = append(_sort, sort)
	}

	return &ShipmentResultType{
		EntityResultType: EntityResultType{
			CurrentPage:  opts.CurrentPage,
			PerPage:      opts.PerPage,
			Rand:         opts.Rand,
			Query:        &query,
			Sort:         _sort,
			Filter:       opts.Filter,
			SelectionSet: selectionSet,
		},
	}, nil
}

type GeneratedShipmentResultTypeResolver struct{ *GeneratedResolver }

func (r *GeneratedShipmentResultTypeResolver) Data(ctx context.Context, obj *ShipmentResultType) (items []*Shipment, err error) {
	giOpts := GetItemsOptions{
		Alias:      TableName("shipments"),
		Preloaders: []string{},
		Item:       &Shipment{},
	}
	err = obj.GetData(ctx, r.DB.db, giOpts, &items)

	uniqueItems := []*Shipment{}
	idMap := map[string]bool{}
	for _, item := range items {
		if _, ok := idMap[item.ID]; !ok {
			idMap[item.ID] = true
			uniqueItems = append(uniqueItems, item)
		}
	}
	items = uniqueItems

	return
}

func (r *GeneratedShipmentResultTypeResolver) Total(ctx context.Context, obj *ShipmentResultType) (count int, err error) {
	return obj.GetTotal(ctx, r.DB.db, TableName("shipments"), &Shipment{})
}

func (r *GeneratedShipmentResultTypeResolver) TotalPage(ctx context.Context, obj *ShipmentResultType) (count int, err error) {
	total, _ := r.Total(ctx, obj)
	perPage, _ := r.PerPage(ctx, obj)
	totalPage := int(math.Ceil(float64(total) / float64(perPage)))
	if totalPage < 0 {
		totalPage = 0
	}

	return totalPage, nil
}

func (r *GeneratedShipmentResultTypeResolver) CurrentPage(ctx context.Context, obj *ShipmentResultType) (count int, err error) {
	return int(*obj.EntityResultType.CurrentPage), nil
}

func (r *GeneratedShipmentResultTypeResolver) PerPage(ctx context.Context, obj *ShipmentResultType) (count int, err error) {
	return int(*obj.EntityResultType.PerPage), nil
}

type GeneratedShipmentResolver struct{ *GeneratedResolver }

func (r *GeneratedShipmentResolver) StartLocation(ctx context.Context, obj *Shipment) (res *Location, err error) {
	return r.Handlers.ShipmentStartLocation(ctx, r.GeneratedResolver, obj)
}
func ShipmentStartLocationHandler(ctx context.Context, r *GeneratedResolver, obj *Shipment) (items *Location, err error) {
	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)

	if obj.StartLocationID != nil {
		item, _ := loaders["Location"].Load(ctx, dataloader.StringKey(*obj.StartLocationID))()
		items, _ = item.(*Location)

		if items == nil {
			items = &Location{}
		}

	}

	return
}

func (r *GeneratedShipmentResolver) EndLocation(ctx context.Context, obj *Shipment) (res *Location, err error) {
	return r.Handlers.ShipmentEndLocation(ctx, r.GeneratedResolver, obj)
}
func ShipmentEndLocationHandler(ctx context.Context, r *GeneratedResolver, obj *Shipment) (items *Location, err error) {
	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)

	if obj.EndLocationID != nil {
		item, _ := loaders["Location"].Load(ctx, dataloader.StringKey(*obj.EndLocationID))()
		items, _ = item.(*Location)

		if items == nil {
			items = &Location{}
		}

	}

	return
}

func (r *GeneratedQueryResolver) Carrier(ctx context.Context, id string) (*Carrier, error) {
	return r.Handlers.QueryCarrier(ctx, r.GeneratedResolver, id)
}
func QueryCarrierHandler(ctx context.Context, r *GeneratedResolver, id string) (*Carrier, error) {
	selection := []ast.Selection{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		selection = append(selection, f.Field)
	}
	selectionSet := ast.SelectionSet(selection)

	query := CarrierQueryFilter{}
	rt := &CarrierResultType{
		EntityResultType: EntityResultType{
			Query:        &query,
			SelectionSet: &selectionSet,
		},
	}
	qb := r.DB.Query()
	qb = qb.Where(TableName("carriers")+".id = ?", id)

	var items []*Carrier
	giOpts := GetItemsOptions{
		Alias:      TableName("carriers"),
		Preloaders: []string{},
		Item:       &Carrier{},
	}
	err := rt.GetData(ctx, qb, giOpts, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, &NotFoundError{Entity: "Carrier"}
	}
	return items[0], err
}

type QueryCarriersHandlerOptions struct {
	CurrentPage *int
	PerPage     *int
	Q           *string
	Sort        []*CarrierSortType
	Filter      *CarrierFilterType
	Rand        *bool
}

func (r *GeneratedQueryResolver) Carriers(ctx context.Context, current_page *int, per_page *int, q *string, sort []*CarrierSortType, filter *CarrierFilterType, rand *bool) (*CarrierResultType, error) {
	opts := QueryCarriersHandlerOptions{
		CurrentPage: current_page,
		PerPage:     per_page,
		Q:           q,
		Sort:        sort,
		Filter:      filter,
		Rand:        rand,
	}
	return r.Handlers.QueryCarriers(ctx, r.GeneratedResolver, opts)
}
func QueryCarriersHandler(ctx context.Context, r *GeneratedResolver, opts QueryCarriersHandlerOptions) (*CarrierResultType, error) {
	query := CarrierQueryFilter{opts.Q}

	var selectionSet *ast.SelectionSet
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		if f.Field.Name == "data" {
			selectionSet = &f.Field.SelectionSet
		}
	}

	_sort := []EntitySort{}
	for _, sort := range opts.Sort {
		_sort = append(_sort, sort)
	}

	return &CarrierResultType{
		EntityResultType: EntityResultType{
			CurrentPage:  opts.CurrentPage,
			PerPage:      opts.PerPage,
			Rand:         opts.Rand,
			Query:        &query,
			Sort:         _sort,
			Filter:       opts.Filter,
			SelectionSet: selectionSet,
		},
	}, nil
}

type GeneratedCarrierResultTypeResolver struct{ *GeneratedResolver }

func (r *GeneratedCarrierResultTypeResolver) Data(ctx context.Context, obj *CarrierResultType) (items []*Carrier, err error) {
	giOpts := GetItemsOptions{
		Alias:      TableName("carriers"),
		Preloaders: []string{},
		Item:       &Carrier{},
	}
	err = obj.GetData(ctx, r.DB.db, giOpts, &items)

	uniqueItems := []*Carrier{}
	idMap := map[string]bool{}
	for _, item := range items {
		if _, ok := idMap[item.ID]; !ok {
			idMap[item.ID] = true
			uniqueItems = append(uniqueItems, item)
		}
	}
	items = uniqueItems

	return
}

func (r *GeneratedCarrierResultTypeResolver) Total(ctx context.Context, obj *CarrierResultType) (count int, err error) {
	return obj.GetTotal(ctx, r.DB.db, TableName("carriers"), &Carrier{})
}

func (r *GeneratedCarrierResultTypeResolver) TotalPage(ctx context.Context, obj *CarrierResultType) (count int, err error) {
	total, _ := r.Total(ctx, obj)
	perPage, _ := r.PerPage(ctx, obj)
	totalPage := int(math.Ceil(float64(total) / float64(perPage)))
	if totalPage < 0 {
		totalPage = 0
	}

	return totalPage, nil
}

func (r *GeneratedCarrierResultTypeResolver) CurrentPage(ctx context.Context, obj *CarrierResultType) (count int, err error) {
	return int(*obj.EntityResultType.CurrentPage), nil
}

func (r *GeneratedCarrierResultTypeResolver) PerPage(ctx context.Context, obj *CarrierResultType) (count int, err error) {
	return int(*obj.EntityResultType.PerPage), nil
}

func (r *GeneratedQueryResolver) Location(ctx context.Context, id string) (*Location, error) {
	return r.Handlers.QueryLocation(ctx, r.GeneratedResolver, id)
}
func QueryLocationHandler(ctx context.Context, r *GeneratedResolver, id string) (*Location, error) {
	selection := []ast.Selection{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		selection = append(selection, f.Field)
	}
	selectionSet := ast.SelectionSet(selection)

	query := LocationQueryFilter{}
	rt := &LocationResultType{
		EntityResultType: EntityResultType{
			Query:        &query,
			SelectionSet: &selectionSet,
		},
	}
	qb := r.DB.Query()
	qb = qb.Where(TableName("locations")+".id = ?", id)

	var items []*Location
	giOpts := GetItemsOptions{
		Alias:      TableName("locations"),
		Preloaders: []string{},
		Item:       &Location{},
	}
	err := rt.GetData(ctx, qb, giOpts, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, &NotFoundError{Entity: "Location"}
	}
	return items[0], err
}

type QueryLocationsHandlerOptions struct {
	CurrentPage *int
	PerPage     *int
	Q           *string
	Sort        []*LocationSortType
	Filter      *LocationFilterType
	Rand        *bool
}

func (r *GeneratedQueryResolver) Locations(ctx context.Context, current_page *int, per_page *int, q *string, sort []*LocationSortType, filter *LocationFilterType, rand *bool) (*LocationResultType, error) {
	opts := QueryLocationsHandlerOptions{
		CurrentPage: current_page,
		PerPage:     per_page,
		Q:           q,
		Sort:        sort,
		Filter:      filter,
		Rand:        rand,
	}
	return r.Handlers.QueryLocations(ctx, r.GeneratedResolver, opts)
}
func QueryLocationsHandler(ctx context.Context, r *GeneratedResolver, opts QueryLocationsHandlerOptions) (*LocationResultType, error) {
	query := LocationQueryFilter{opts.Q}

	var selectionSet *ast.SelectionSet
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		if f.Field.Name == "data" {
			selectionSet = &f.Field.SelectionSet
		}
	}

	_sort := []EntitySort{}
	for _, sort := range opts.Sort {
		_sort = append(_sort, sort)
	}

	return &LocationResultType{
		EntityResultType: EntityResultType{
			CurrentPage:  opts.CurrentPage,
			PerPage:      opts.PerPage,
			Rand:         opts.Rand,
			Query:        &query,
			Sort:         _sort,
			Filter:       opts.Filter,
			SelectionSet: selectionSet,
		},
	}, nil
}

type GeneratedLocationResultTypeResolver struct{ *GeneratedResolver }

func (r *GeneratedLocationResultTypeResolver) Data(ctx context.Context, obj *LocationResultType) (items []*Location, err error) {
	giOpts := GetItemsOptions{
		Alias:      TableName("locations"),
		Preloaders: []string{},
		Item:       &Location{},
	}
	err = obj.GetData(ctx, r.DB.db, giOpts, &items)

	uniqueItems := []*Location{}
	idMap := map[string]bool{}
	for _, item := range items {
		if _, ok := idMap[item.ID]; !ok {
			idMap[item.ID] = true
			uniqueItems = append(uniqueItems, item)
		}
	}
	items = uniqueItems

	return
}

func (r *GeneratedLocationResultTypeResolver) Total(ctx context.Context, obj *LocationResultType) (count int, err error) {
	return obj.GetTotal(ctx, r.DB.db, TableName("locations"), &Location{})
}

func (r *GeneratedLocationResultTypeResolver) TotalPage(ctx context.Context, obj *LocationResultType) (count int, err error) {
	total, _ := r.Total(ctx, obj)
	perPage, _ := r.PerPage(ctx, obj)
	totalPage := int(math.Ceil(float64(total) / float64(perPage)))
	if totalPage < 0 {
		totalPage = 0
	}

	return totalPage, nil
}

func (r *GeneratedLocationResultTypeResolver) CurrentPage(ctx context.Context, obj *LocationResultType) (count int, err error) {
	return int(*obj.EntityResultType.CurrentPage), nil
}

func (r *GeneratedLocationResultTypeResolver) PerPage(ctx context.Context, obj *LocationResultType) (count int, err error) {
	return int(*obj.EntityResultType.PerPage), nil
}

type GeneratedLocationResolver struct{ *GeneratedResolver }

func (r *GeneratedLocationResolver) StartShipments(ctx context.Context, obj *Location) (res []*Shipment, err error) {
	return r.Handlers.LocationStartShipments(ctx, r.GeneratedResolver, obj)
}
func LocationStartShipmentsHandler(ctx context.Context, r *GeneratedResolver, obj *Location) (items []*Shipment, err error) {
	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)

	item, _ := loaders["ShipmentStartLocation"].Load(ctx, dataloader.StringKey(obj.ID))()
	items = []*Shipment{}
	if item != nil {
		items = item.([]*Shipment)
	}

	return
}

func (r *GeneratedLocationResolver) StartShipmentsIds(ctx context.Context, obj *Location) (ids []string, err error) {
	ids = []string{}
	items := []*Shipment{}

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	item, _ := loaders["ShipmentStartLocation"].Load(ctx, dataloader.StringKey(obj.ID))()

	if item != nil {
		items = item.([]*Shipment)
	}

	for _, v := range items {
		ids = append(ids, v.ID)
	}

	return
}

func (r *GeneratedLocationResolver) EndShipments(ctx context.Context, obj *Location) (res []*Shipment, err error) {
	return r.Handlers.LocationEndShipments(ctx, r.GeneratedResolver, obj)
}
func LocationEndShipmentsHandler(ctx context.Context, r *GeneratedResolver, obj *Location) (items []*Shipment, err error) {
	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)

	item, _ := loaders["ShipmentEndLocation"].Load(ctx, dataloader.StringKey(obj.ID))()
	items = []*Shipment{}
	if item != nil {
		items = item.([]*Shipment)
	}

	return
}

func (r *GeneratedLocationResolver) EndShipmentsIds(ctx context.Context, obj *Location) (ids []string, err error) {
	ids = []string{}
	items := []*Shipment{}

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	item, _ := loaders["ShipmentEndLocation"].Load(ctx, dataloader.StringKey(obj.ID))()

	if item != nil {
		items = item.([]*Shipment)
	}

	for _, v := range items {
		ids = append(ids, v.ID)
	}

	return
}

func (r *GeneratedQueryResolver) Equipmentd(ctx context.Context, id string) (*Equipmentd, error) {
	return r.Handlers.QueryEquipmentd(ctx, r.GeneratedResolver, id)
}
func QueryEquipmentdHandler(ctx context.Context, r *GeneratedResolver, id string) (*Equipmentd, error) {
	selection := []ast.Selection{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		selection = append(selection, f.Field)
	}
	selectionSet := ast.SelectionSet(selection)

	query := EquipmentdQueryFilter{}
	rt := &EquipmentdResultType{
		EntityResultType: EntityResultType{
			Query:        &query,
			SelectionSet: &selectionSet,
		},
	}
	qb := r.DB.Query()
	qb = qb.Where(TableName("equipmentds")+".id = ?", id)

	var items []*Equipmentd
	giOpts := GetItemsOptions{
		Alias:      TableName("equipmentds"),
		Preloaders: []string{},
		Item:       &Equipmentd{},
	}
	err := rt.GetData(ctx, qb, giOpts, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, &NotFoundError{Entity: "Equipmentd"}
	}
	return items[0], err
}

type QueryEquipmentdsHandlerOptions struct {
	CurrentPage *int
	PerPage     *int
	Q           *string
	Sort        []*EquipmentdSortType
	Filter      *EquipmentdFilterType
	Rand        *bool
}

func (r *GeneratedQueryResolver) Equipmentds(ctx context.Context, current_page *int, per_page *int, q *string, sort []*EquipmentdSortType, filter *EquipmentdFilterType, rand *bool) (*EquipmentdResultType, error) {
	opts := QueryEquipmentdsHandlerOptions{
		CurrentPage: current_page,
		PerPage:     per_page,
		Q:           q,
		Sort:        sort,
		Filter:      filter,
		Rand:        rand,
	}
	return r.Handlers.QueryEquipmentds(ctx, r.GeneratedResolver, opts)
}
func QueryEquipmentdsHandler(ctx context.Context, r *GeneratedResolver, opts QueryEquipmentdsHandlerOptions) (*EquipmentdResultType, error) {
	query := EquipmentdQueryFilter{opts.Q}

	var selectionSet *ast.SelectionSet
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		if f.Field.Name == "data" {
			selectionSet = &f.Field.SelectionSet
		}
	}

	_sort := []EntitySort{}
	for _, sort := range opts.Sort {
		_sort = append(_sort, sort)
	}

	return &EquipmentdResultType{
		EntityResultType: EntityResultType{
			CurrentPage:  opts.CurrentPage,
			PerPage:      opts.PerPage,
			Rand:         opts.Rand,
			Query:        &query,
			Sort:         _sort,
			Filter:       opts.Filter,
			SelectionSet: selectionSet,
		},
	}, nil
}

type GeneratedEquipmentdResultTypeResolver struct{ *GeneratedResolver }

func (r *GeneratedEquipmentdResultTypeResolver) Data(ctx context.Context, obj *EquipmentdResultType) (items []*Equipmentd, err error) {
	giOpts := GetItemsOptions{
		Alias:      TableName("equipmentds"),
		Preloaders: []string{},
		Item:       &Equipmentd{},
	}
	err = obj.GetData(ctx, r.DB.db, giOpts, &items)

	uniqueItems := []*Equipmentd{}
	idMap := map[string]bool{}
	for _, item := range items {
		if _, ok := idMap[item.ID]; !ok {
			idMap[item.ID] = true
			uniqueItems = append(uniqueItems, item)
		}
	}
	items = uniqueItems

	return
}

func (r *GeneratedEquipmentdResultTypeResolver) Total(ctx context.Context, obj *EquipmentdResultType) (count int, err error) {
	return obj.GetTotal(ctx, r.DB.db, TableName("equipmentds"), &Equipmentd{})
}

func (r *GeneratedEquipmentdResultTypeResolver) TotalPage(ctx context.Context, obj *EquipmentdResultType) (count int, err error) {
	total, _ := r.Total(ctx, obj)
	perPage, _ := r.PerPage(ctx, obj)
	totalPage := int(math.Ceil(float64(total) / float64(perPage)))
	if totalPage < 0 {
		totalPage = 0
	}

	return totalPage, nil
}

func (r *GeneratedEquipmentdResultTypeResolver) CurrentPage(ctx context.Context, obj *EquipmentdResultType) (count int, err error) {
	return int(*obj.EntityResultType.CurrentPage), nil
}

func (r *GeneratedEquipmentdResultTypeResolver) PerPage(ctx context.Context, obj *EquipmentdResultType) (count int, err error) {
	return int(*obj.EntityResultType.PerPage), nil
}
