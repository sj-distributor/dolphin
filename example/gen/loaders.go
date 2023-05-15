package gen

import (
	"context"
	"errors"

	"github.com/graph-gophers/dataloader"
	"gorm.io/gorm"
)

func GetLoaders(db *DB) map[string]*dataloader.Loader {
	loaders := map[string]*dataloader.Loader{}

	ordersBatchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result

		ids := make([]string, len(keys))
		for i, key := range keys {
			ids[i] = key.String()
		}

		items := &[]Order{}
		selects := GetFieldsRequested(ctx, "orders")
		if len(selects) > 0 && IndexOf(selects, "orders"+".id") == -1 {
			selects = append(selects, "orders"+".id")
		}

		res := db.Query().Select(selects).Find(items, "id IN (?)", ids)
		if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return []*dataloader.Result{
				{Error: res.Error},
			}
		}

		itemMap := make(map[string]Order, len(keys))
		for _, item := range *items {
			itemMap[item.ID] = item
		}

		for _, key := range keys {
			id := key.String()
			item, ok := itemMap[id]
			if !ok {
				results = append(results, &dataloader.Result{
					Data:  nil,
					Error: nil,
					// Error: fmt.Errorf("Order with id '%s' not found", id),
				})
			} else {
				results = append(results, &dataloader.Result{
					Data:  &item,
					Error: nil,
				})
			}
		}
		return results
	}

	loaders["Order"] = dataloader.NewBatchedLoader(ordersBatchFn, dataloader.WithClearCacheOnBatch())

	shipmentsStartLocationBatchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result

		ids := make([]string, len(keys))
		for i, key := range keys {
			ids[i] = key.String()
		}

		items := &[]Shipment{}
		selects := GetFieldsRequested(ctx, "shipments")
		if IndexOf(selects, "shipments"+".id") == -1 {
			selects = append(selects, "shipments"+".id")
		}

		if IndexOf(selects, "shipments"+".startLocation_id") == -1 {
			selects = append(selects, "shipments"+".startLocation_id")
		}

		res := db.Query().Select(selects).Find(items, "startLocation_id IN (?)", ids)
		if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return []*dataloader.Result{
				{Error: res.Error},
			}
		}

		itemMap := make(map[string][]*Shipment, len(keys))
		for _, v := range *items {
			item := v
			if itemMap[*item.StartLocationID] == nil {
				itemMap[*item.StartLocationID] = []*Shipment{}
			}
			itemMap[*item.StartLocationID] = append(itemMap[*item.StartLocationID], &item)
		}

		for _, key := range keys {
			id := key.String()
			item, ok := itemMap[id]
			if !ok {
				results = append(results, &dataloader.Result{
					Data:  nil,
					Error: nil,
					// Error: fmt.Errorf("Shipment with id '%s' not found", id),
				})
			} else {
				results = append(results, &dataloader.Result{
					Data:  item,
					Error: nil,
				})
			}
		}
		return results
	}
	loaders["ShipmentStartLocation"] = dataloader.NewBatchedLoader(shipmentsStartLocationBatchFn, dataloader.WithClearCacheOnBatch())

	shipmentsEndLocationBatchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result

		ids := make([]string, len(keys))
		for i, key := range keys {
			ids[i] = key.String()
		}

		items := &[]Shipment{}
		selects := GetFieldsRequested(ctx, "shipments")
		if IndexOf(selects, "shipments"+".id") == -1 {
			selects = append(selects, "shipments"+".id")
		}

		if IndexOf(selects, "shipments"+".endLocation_id") == -1 {
			selects = append(selects, "shipments"+".endLocation_id")
		}

		res := db.Query().Select(selects).Find(items, "endLocation_id IN (?)", ids)
		if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return []*dataloader.Result{
				{Error: res.Error},
			}
		}

		itemMap := make(map[string][]*Shipment, len(keys))
		for _, v := range *items {
			item := v
			if itemMap[*item.EndLocationID] == nil {
				itemMap[*item.EndLocationID] = []*Shipment{}
			}
			itemMap[*item.EndLocationID] = append(itemMap[*item.EndLocationID], &item)
		}

		for _, key := range keys {
			id := key.String()
			item, ok := itemMap[id]
			if !ok {
				results = append(results, &dataloader.Result{
					Data:  nil,
					Error: nil,
					// Error: fmt.Errorf("Shipment with id '%s' not found", id),
				})
			} else {
				results = append(results, &dataloader.Result{
					Data:  item,
					Error: nil,
				})
			}
		}
		return results
	}
	loaders["ShipmentEndLocation"] = dataloader.NewBatchedLoader(shipmentsEndLocationBatchFn, dataloader.WithClearCacheOnBatch())

	shipmentsBatchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result

		ids := make([]string, len(keys))
		for i, key := range keys {
			ids[i] = key.String()
		}

		items := &[]Shipment{}
		selects := GetFieldsRequested(ctx, "shipments")
		if len(selects) > 0 && IndexOf(selects, "shipments"+".id") == -1 {
			selects = append(selects, "shipments"+".id")
		}

		res := db.Query().Select(selects).Find(items, "id IN (?)", ids)
		if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return []*dataloader.Result{
				{Error: res.Error},
			}
		}

		itemMap := make(map[string]Shipment, len(keys))
		for _, item := range *items {
			itemMap[item.ID] = item
		}

		for _, key := range keys {
			id := key.String()
			item, ok := itemMap[id]
			if !ok {
				results = append(results, &dataloader.Result{
					Data:  nil,
					Error: nil,
					// Error: fmt.Errorf("Shipment with id '%s' not found", id),
				})
			} else {
				results = append(results, &dataloader.Result{
					Data:  &item,
					Error: nil,
				})
			}
		}
		return results
	}

	loaders["Shipment"] = dataloader.NewBatchedLoader(shipmentsBatchFn, dataloader.WithClearCacheOnBatch())

	carriersBatchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result

		ids := make([]string, len(keys))
		for i, key := range keys {
			ids[i] = key.String()
		}

		items := &[]Carrier{}
		selects := GetFieldsRequested(ctx, "carriers")
		if len(selects) > 0 && IndexOf(selects, "carriers"+".id") == -1 {
			selects = append(selects, "carriers"+".id")
		}

		res := db.Query().Select(selects).Find(items, "id IN (?)", ids)
		if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return []*dataloader.Result{
				{Error: res.Error},
			}
		}

		itemMap := make(map[string]Carrier, len(keys))
		for _, item := range *items {
			itemMap[item.ID] = item
		}

		for _, key := range keys {
			id := key.String()
			item, ok := itemMap[id]
			if !ok {
				results = append(results, &dataloader.Result{
					Data:  nil,
					Error: nil,
					// Error: fmt.Errorf("Carrier with id '%s' not found", id),
				})
			} else {
				results = append(results, &dataloader.Result{
					Data:  &item,
					Error: nil,
				})
			}
		}
		return results
	}

	loaders["Carrier"] = dataloader.NewBatchedLoader(carriersBatchFn, dataloader.WithClearCacheOnBatch())

	locationsBatchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result

		ids := make([]string, len(keys))
		for i, key := range keys {
			ids[i] = key.String()
		}

		items := &[]Location{}
		selects := GetFieldsRequested(ctx, "locations")
		if len(selects) > 0 && IndexOf(selects, "locations"+".id") == -1 {
			selects = append(selects, "locations"+".id")
		}

		res := db.Query().Select(selects).Find(items, "id IN (?)", ids)
		if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return []*dataloader.Result{
				{Error: res.Error},
			}
		}

		itemMap := make(map[string]Location, len(keys))
		for _, item := range *items {
			itemMap[item.ID] = item
		}

		for _, key := range keys {
			id := key.String()
			item, ok := itemMap[id]
			if !ok {
				results = append(results, &dataloader.Result{
					Data:  nil,
					Error: nil,
					// Error: fmt.Errorf("Location with id '%s' not found", id),
				})
			} else {
				results = append(results, &dataloader.Result{
					Data:  &item,
					Error: nil,
				})
			}
		}
		return results
	}

	loaders["Location"] = dataloader.NewBatchedLoader(locationsBatchFn, dataloader.WithClearCacheOnBatch())

	equipmentdsBatchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result

		ids := make([]string, len(keys))
		for i, key := range keys {
			ids[i] = key.String()
		}

		items := &[]Equipmentd{}
		selects := GetFieldsRequested(ctx, "equipmentds")
		if len(selects) > 0 && IndexOf(selects, "equipmentds"+".id") == -1 {
			selects = append(selects, "equipmentds"+".id")
		}

		res := db.Query().Select(selects).Find(items, "id IN (?)", ids)
		if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return []*dataloader.Result{
				{Error: res.Error},
			}
		}

		itemMap := make(map[string]Equipmentd, len(keys))
		for _, item := range *items {
			itemMap[item.ID] = item
		}

		for _, key := range keys {
			id := key.String()
			item, ok := itemMap[id]
			if !ok {
				results = append(results, &dataloader.Result{
					Data:  nil,
					Error: nil,
					// Error: fmt.Errorf("Equipmentd with id '%s' not found", id),
				})
			} else {
				results = append(results, &dataloader.Result{
					Data:  &item,
					Error: nil,
				})
			}
		}
		return results
	}

	loaders["Equipmentd"] = dataloader.NewBatchedLoader(equipmentdsBatchFn, dataloader.WithClearCacheOnBatch())

	return loaders
}
