package gen

import (
	"context"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm/clause"
)

type GeneratedMutationResolver struct{ *GeneratedResolver }

type MutationEvents struct {
	Events []Event
}

func (r *GeneratedMutationResolver) CreateOrder(ctx context.Context, input map[string]interface{}) (item *Order, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreateOrder(ctx, r.GeneratedResolver, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func CreateOrderHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Order, err error) {
	item = &Order{}

	now := time.Now()
	milliTime := now.UnixNano() / 1e6
	principalID := GetPrincipalIDFromContext(ctx)

	tx := GetTransaction(ctx)
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	event := NewEvent(EventMetadata{
		Type:        EventTypeCreated,
		Entity:      "Order",
		EntityID:    item.ID,
		Date:        milliTime,
		PrincipalID: principalID,
	})

	var changes OrderChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	err = CheckStructFieldIsEmpty(item, input)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	item.ID = uuid.Must(uuid.NewV4()).String()
	item.CreatedAt = milliTime
	item.CreatedBy = principalID

	if _, ok := input["orderNo"]; ok && (item.OrderNo != changes.OrderNo) {
		item.OrderNo = changes.OrderNo
		event.AddNewValue("orderNo", changes.OrderNo)
	}

	if _, ok := input["customerInfo"]; ok && (item.CustomerInfo != changes.CustomerInfo) && (item.CustomerInfo == nil || changes.CustomerInfo == nil || *item.CustomerInfo != *changes.CustomerInfo) {
		item.CustomerInfo = changes.CustomerInfo
		event.AddNewValue("customerInfo", changes.CustomerInfo)
	}

	if _, ok := input["goodsInfo"]; ok && (item.GoodsInfo != changes.GoodsInfo) && (item.GoodsInfo == nil || changes.GoodsInfo == nil || *item.GoodsInfo != *changes.GoodsInfo) {
		item.GoodsInfo = changes.GoodsInfo
		event.AddNewValue("goodsInfo", changes.GoodsInfo)
	}

	if err := tx.Omit(clause.Associations).Create(item).Error; err != nil {
		tx.Rollback()
		return item, err
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}
func (r *GeneratedMutationResolver) UpdateOrder(ctx context.Context, id string, input map[string]interface{}) (item *Order, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdateOrder(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func UpdateOrderHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Order, err error) {
	item = &Order{}
	newItem := &Order{}

	isChange := false
	now := time.Now()
	milliTime := now.UnixNano() / 1e6
	principalID := GetPrincipalIDFromContext(ctx)

	tx := GetTransaction(ctx)
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	event := NewEvent(EventMetadata{
		Type:        EventTypeUpdated,
		Entity:      "Order",
		EntityID:    id,
		Date:        milliTime,
		PrincipalID: principalID,
	})

	var changes OrderChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	err = CheckStructFieldIsEmpty(item, input)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err = GetItem(ctx, tx, TableName("orders"), item, &id); err != nil {
		tx.Rollback()
		return nil, err
	}

	if item.UpdatedBy != nil && principalID != nil && *item.UpdatedBy != *principalID {
		newItem.UpdatedBy = principalID
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		event.AddOldValue("id", item.ID)
		event.AddNewValue("id", changes.ID)
		item.ID = changes.ID
		newItem.ID = changes.ID
		isChange = true
	}

	if _, ok := input["orderNo"]; ok && (item.OrderNo != changes.OrderNo) {
		event.AddOldValue("orderNo", item.OrderNo)
		event.AddNewValue("orderNo", changes.OrderNo)
		item.OrderNo = changes.OrderNo
		newItem.OrderNo = changes.OrderNo
		isChange = true
	}

	if _, ok := input["customerInfo"]; ok && (item.CustomerInfo != changes.CustomerInfo) && (item.CustomerInfo == nil || changes.CustomerInfo == nil || *item.CustomerInfo != *changes.CustomerInfo) {
		event.AddOldValue("customerInfo", item.CustomerInfo)
		event.AddNewValue("customerInfo", changes.CustomerInfo)
		item.CustomerInfo = changes.CustomerInfo
		newItem.CustomerInfo = changes.CustomerInfo
		isChange = true
	}

	if _, ok := input["goodsInfo"]; ok && (item.GoodsInfo != changes.GoodsInfo) && (item.GoodsInfo == nil || changes.GoodsInfo == nil || *item.GoodsInfo != *changes.GoodsInfo) {
		event.AddOldValue("goodsInfo", item.GoodsInfo)
		event.AddNewValue("goodsInfo", changes.GoodsInfo)
		item.GoodsInfo = changes.GoodsInfo
		newItem.GoodsInfo = changes.GoodsInfo
		isChange = true
	}

	// if err := validator.Struct(item); err != nil {
	// 	tx.Rollback()
	// 	return nil, err
	// }

	if !isChange {
		return item, nil
	}

	if err := tx.Model(&newItem).Where("id = ?", id).Omit(clause.Associations).Updates(newItem).Error; err != nil {
		tx.Rollback()
		return item, err
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}

func DeleteOrderFunc(ctx context.Context, r *GeneratedResolver, id string, tye string, unscoped *bool) (err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item := &Order{}
	now := time.Now()
	tx := GetTransaction(ctx)
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err = GetItem(ctx, tx, TableName("orders"), item, &id); err != nil {
		tx.Rollback()
		return err
	}

	deletedAt := now.UnixNano() / 1e6

	event := NewEvent(EventMetadata{
		Type:        EventTypeDeleted,
		Entity:      "Order",
		EntityID:    id,
		Date:        deletedAt,
		PrincipalID: principalID,
	})

	// 如果是恢复删除数据
	if tye == "recovery" {
		if err := tx.Unscoped().Model(&item).Updates(map[string]interface{}{"DeletedAt": nil, "DeletedBy": nil}).Error; err != nil {
			tx.Rollback()
			return err
		}
	} else {
		if unscoped != nil && *unscoped == true {
			if err := tx.Unscoped().Model(&item).Delete(item).Error; err != nil {
				tx.Rollback()
				return err
			}
		} else if err := tx.Model(&item).Updates(Order{DeletedAt: &deletedAt, DeletedBy: principalID, UpdatedBy: principalID}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}

func (r *GeneratedMutationResolver) DeleteOrders(ctx context.Context, id []string, unscoped *bool) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteOrders(ctx, r.GeneratedResolver, id, unscoped)
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

func DeleteOrdersHandler(ctx context.Context, r *GeneratedResolver, id []string, unscoped *bool) (bool, error) {
	var err error = nil

	if len(id) > 0 {
		for _, v := range id {
			err = DeleteOrderFunc(ctx, r, v, "delete", unscoped)
			if err != nil {
				break
			}
		}
	}

	if err != nil {
		return false, err
	}
	return true, err
}

func (r *GeneratedMutationResolver) RecoveryOrders(ctx context.Context, id []string) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.RecoveryOrders(ctx, r.GeneratedResolver, id)
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

func RecoveryOrdersHandler(ctx context.Context, r *GeneratedResolver, id []string) (bool, error) {
	var err error = nil

	var unscoped bool = false

	if len(id) > 0 {
		for _, v := range id {
			err = DeleteOrderFunc(ctx, r, v, "recovery", &unscoped)
			if err != nil {
				break
			}
		}
	}

	if err != nil {
		return false, err
	}
	return true, err
}

func (r *GeneratedMutationResolver) CreateShipment(ctx context.Context, input map[string]interface{}) (item *Shipment, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreateShipment(ctx, r.GeneratedResolver, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func CreateShipmentHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Shipment, err error) {
	item = &Shipment{}

	now := time.Now()
	milliTime := now.UnixNano() / 1e6
	principalID := GetPrincipalIDFromContext(ctx)

	tx := GetTransaction(ctx)
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	event := NewEvent(EventMetadata{
		Type:        EventTypeCreated,
		Entity:      "Shipment",
		EntityID:    item.ID,
		Date:        milliTime,
		PrincipalID: principalID,
	})

	var changes ShipmentChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	err = CheckStructFieldIsEmpty(item, input)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	item.ID = uuid.Must(uuid.NewV4()).String()
	item.CreatedAt = milliTime
	item.CreatedBy = principalID

	if input["startLocation"] != nil && input["startLocationId"] != nil {
		tx.Rollback()
		return nil, fmt.Errorf("startLocationId and startLocation cannot coexist")
	}

	if _, ok := input["startLocation"]; ok {
		var startLocation *Location
		startLocationInput := input["startLocation"].(map[string]interface{})

		if startLocationInput["id"] == nil {
			startLocation, err = r.Handlers.CreateLocation(ctx, r, startLocationInput)
		} else {
			startLocation, err = r.Handlers.UpdateLocation(ctx, r, startLocationInput["id"].(string), startLocationInput)
		}

		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf(fmt.Sprintf("startLocation %s", err.Error()))
		}

		event.AddNewValue("startLocation", changes.StartLocation)
		item.StartLocation = startLocation
		item.StartLocationID = &startLocation.ID
	}

	if input["endLocation"] != nil && input["endLocationId"] != nil {
		tx.Rollback()
		return nil, fmt.Errorf("endLocationId and endLocation cannot coexist")
	}

	if _, ok := input["endLocation"]; ok {
		var endLocation *Location
		endLocationInput := input["endLocation"].(map[string]interface{})

		if endLocationInput["id"] == nil {
			endLocation, err = r.Handlers.CreateLocation(ctx, r, endLocationInput)
		} else {
			endLocation, err = r.Handlers.UpdateLocation(ctx, r, endLocationInput["id"].(string), endLocationInput)
		}

		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf(fmt.Sprintf("endLocation %s", err.Error()))
		}

		event.AddNewValue("endLocation", changes.EndLocation)
		item.EndLocation = endLocation
		item.EndLocationID = &endLocation.ID
	}

	if _, ok := input["shipmentNo"]; ok && (item.ShipmentNo != changes.ShipmentNo) {
		item.ShipmentNo = changes.ShipmentNo
		event.AddNewValue("shipmentNo", changes.ShipmentNo)
	}

	if _, ok := input["transportationMode"]; ok && (item.TransportationMode != changes.TransportationMode) && (item.TransportationMode == nil || changes.TransportationMode == nil || *item.TransportationMode != *changes.TransportationMode) {
		item.TransportationMode = changes.TransportationMode
		event.AddNewValue("transportationMode", changes.TransportationMode)
	}

	if _, ok := input["startLocationId"]; ok && (item.StartLocationID != changes.StartLocationID) && (item.StartLocationID == nil || changes.StartLocationID == nil || *item.StartLocationID != *changes.StartLocationID) {

		// if err := tx.Select("id").Where("id", input["startLocationId"]).First(&StartLocation{}).Error; err != nil {
		// 	tx.Rollback()
		// 	return nil, fmt.Errorf("startLocationId " + err.Error())
		// }
		item.StartLocationID = changes.StartLocationID
		event.AddNewValue("startLocationId", changes.StartLocationID)
	}

	if _, ok := input["endLocationId"]; ok && (item.EndLocationID != changes.EndLocationID) && (item.EndLocationID == nil || changes.EndLocationID == nil || *item.EndLocationID != *changes.EndLocationID) {

		// if err := tx.Select("id").Where("id", input["endLocationId"]).First(&EndLocation{}).Error; err != nil {
		// 	tx.Rollback()
		// 	return nil, fmt.Errorf("endLocationId " + err.Error())
		// }
		item.EndLocationID = changes.EndLocationID
		event.AddNewValue("endLocationId", changes.EndLocationID)
	}

	if err := tx.Omit(clause.Associations).Create(item).Error; err != nil {
		tx.Rollback()
		return item, err
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}
func (r *GeneratedMutationResolver) UpdateShipment(ctx context.Context, id string, input map[string]interface{}) (item *Shipment, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdateShipment(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func UpdateShipmentHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Shipment, err error) {
	item = &Shipment{}
	newItem := &Shipment{}

	isChange := false
	now := time.Now()
	milliTime := now.UnixNano() / 1e6
	principalID := GetPrincipalIDFromContext(ctx)

	tx := GetTransaction(ctx)
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	event := NewEvent(EventMetadata{
		Type:        EventTypeUpdated,
		Entity:      "Shipment",
		EntityID:    id,
		Date:        milliTime,
		PrincipalID: principalID,
	})

	var changes ShipmentChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	err = CheckStructFieldIsEmpty(item, input)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if input["startLocation"] != nil && input["startLocationId"] != nil {
		tx.Rollback()
		return nil, fmt.Errorf("startLocationId and startLocation cannot coexist")
	}

	if input["endLocation"] != nil && input["endLocationId"] != nil {
		tx.Rollback()
		return nil, fmt.Errorf("endLocationId and endLocation cannot coexist")
	}

	if err = GetItem(ctx, tx, TableName("shipments"), item, &id); err != nil {
		tx.Rollback()
		return nil, err
	}

	if item.UpdatedBy != nil && principalID != nil && *item.UpdatedBy != *principalID {
		newItem.UpdatedBy = principalID
	}

	if _, ok := input["startLocation"]; ok {
		var startLocation *Location
		startLocationInput := input["startLocation"].(map[string]interface{})

		if startLocationInput["id"] == nil {
			startLocation, err = r.Handlers.CreateLocation(ctx, r, startLocationInput)
		} else {
			startLocation, err = r.Handlers.UpdateLocation(ctx, r, startLocationInput["id"].(string), startLocationInput)
		}

		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf(fmt.Sprintf("startLocation %s", err.Error()))
		}

		if err := tx.Model(&item).Association("StartLocation").Clear(); err != nil {
			tx.Rollback()
			return nil, err
		}

		if err := tx.Model(&Location{}).Where("id = ?", startLocation.ID).Update("start_shipments_id", item.ID).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		event.AddOldValue("startLocation", item.StartLocation)
		event.AddNewValue("startLocation", changes.StartLocation)
		item.StartLocation = startLocation
		newItem.StartLocationID = &startLocation.ID
		isChange = true
	}

	if _, ok := input["endLocation"]; ok {
		var endLocation *Location
		endLocationInput := input["endLocation"].(map[string]interface{})

		if endLocationInput["id"] == nil {
			endLocation, err = r.Handlers.CreateLocation(ctx, r, endLocationInput)
		} else {
			endLocation, err = r.Handlers.UpdateLocation(ctx, r, endLocationInput["id"].(string), endLocationInput)
		}

		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf(fmt.Sprintf("endLocation %s", err.Error()))
		}

		if err := tx.Model(&item).Association("EndLocation").Clear(); err != nil {
			tx.Rollback()
			return nil, err
		}

		if err := tx.Model(&Location{}).Where("id = ?", endLocation.ID).Update("end_shipments_id", item.ID).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		event.AddOldValue("endLocation", item.EndLocation)
		event.AddNewValue("endLocation", changes.EndLocation)
		item.EndLocation = endLocation
		newItem.EndLocationID = &endLocation.ID
		isChange = true
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		event.AddOldValue("id", item.ID)
		event.AddNewValue("id", changes.ID)
		item.ID = changes.ID
		newItem.ID = changes.ID
		isChange = true
	}

	if _, ok := input["shipmentNo"]; ok && (item.ShipmentNo != changes.ShipmentNo) {
		event.AddOldValue("shipmentNo", item.ShipmentNo)
		event.AddNewValue("shipmentNo", changes.ShipmentNo)
		item.ShipmentNo = changes.ShipmentNo
		newItem.ShipmentNo = changes.ShipmentNo
		isChange = true
	}

	if _, ok := input["transportationMode"]; ok && (item.TransportationMode != changes.TransportationMode) && (item.TransportationMode == nil || changes.TransportationMode == nil || *item.TransportationMode != *changes.TransportationMode) {
		event.AddOldValue("transportationMode", item.TransportationMode)
		event.AddNewValue("transportationMode", changes.TransportationMode)
		item.TransportationMode = changes.TransportationMode
		newItem.TransportationMode = changes.TransportationMode
		isChange = true
	}

	if _, ok := input["startLocationId"]; ok && (item.StartLocationID != changes.StartLocationID) && (item.StartLocationID == nil || changes.StartLocationID == nil || *item.StartLocationID != *changes.StartLocationID) {

		// if err := tx.Select("id").Where("id", input["startLocationId"]).First(&StartLocation{}).Error; err != nil {
		// 	tx.Rollback()
		// 	return nil, fmt.Errorf("startLocationId " + err.Error())
		// }
		event.AddOldValue("startLocationId", item.StartLocationID)
		event.AddNewValue("startLocationId", changes.StartLocationID)
		item.StartLocationID = changes.StartLocationID
		newItem.StartLocationID = changes.StartLocationID
		isChange = true
	}

	if _, ok := input["endLocationId"]; ok && (item.EndLocationID != changes.EndLocationID) && (item.EndLocationID == nil || changes.EndLocationID == nil || *item.EndLocationID != *changes.EndLocationID) {

		// if err := tx.Select("id").Where("id", input["endLocationId"]).First(&EndLocation{}).Error; err != nil {
		// 	tx.Rollback()
		// 	return nil, fmt.Errorf("endLocationId " + err.Error())
		// }
		event.AddOldValue("endLocationId", item.EndLocationID)
		event.AddNewValue("endLocationId", changes.EndLocationID)
		item.EndLocationID = changes.EndLocationID
		newItem.EndLocationID = changes.EndLocationID
		isChange = true
	}

	// if err := validator.Struct(item); err != nil {
	// 	tx.Rollback()
	// 	return nil, err
	// }

	if !isChange {
		return item, nil
	}

	if err := tx.Model(&newItem).Where("id = ?", id).Omit(clause.Associations).Updates(newItem).Error; err != nil {
		tx.Rollback()
		return item, err
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}

func DeleteShipmentFunc(ctx context.Context, r *GeneratedResolver, id string, tye string, unscoped *bool) (err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item := &Shipment{}
	now := time.Now()
	tx := GetTransaction(ctx)
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err = GetItem(ctx, tx, TableName("shipments"), item, &id); err != nil {
		tx.Rollback()
		return err
	}

	deletedAt := now.UnixNano() / 1e6

	event := NewEvent(EventMetadata{
		Type:        EventTypeDeleted,
		Entity:      "Shipment",
		EntityID:    id,
		Date:        deletedAt,
		PrincipalID: principalID,
	})

	// 如果是恢复删除数据
	if tye == "recovery" {
		if err := tx.Unscoped().Model(&item).Updates(map[string]interface{}{"DeletedAt": nil, "DeletedBy": nil}).Error; err != nil {
			tx.Rollback()
			return err
		}
	} else {
		if unscoped != nil && *unscoped == true {
			if err := tx.Unscoped().Model(&item).Delete(item).Error; err != nil {
				tx.Rollback()
				return err
			}
		} else if err := tx.Model(&item).Updates(Shipment{DeletedAt: &deletedAt, DeletedBy: principalID, UpdatedBy: principalID}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}

func (r *GeneratedMutationResolver) DeleteShipments(ctx context.Context, id []string, unscoped *bool) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteShipments(ctx, r.GeneratedResolver, id, unscoped)
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

func DeleteShipmentsHandler(ctx context.Context, r *GeneratedResolver, id []string, unscoped *bool) (bool, error) {
	var err error = nil

	if len(id) > 0 {
		for _, v := range id {
			err = DeleteShipmentFunc(ctx, r, v, "delete", unscoped)
			if err != nil {
				break
			}
		}
	}

	if err != nil {
		return false, err
	}
	return true, err
}

func (r *GeneratedMutationResolver) RecoveryShipments(ctx context.Context, id []string) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.RecoveryShipments(ctx, r.GeneratedResolver, id)
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

func RecoveryShipmentsHandler(ctx context.Context, r *GeneratedResolver, id []string) (bool, error) {
	var err error = nil

	var unscoped bool = false

	if len(id) > 0 {
		for _, v := range id {
			err = DeleteShipmentFunc(ctx, r, v, "recovery", &unscoped)
			if err != nil {
				break
			}
		}
	}

	if err != nil {
		return false, err
	}
	return true, err
}

func (r *GeneratedMutationResolver) CreateCarrier(ctx context.Context, input map[string]interface{}) (item *Carrier, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreateCarrier(ctx, r.GeneratedResolver, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func CreateCarrierHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Carrier, err error) {
	item = &Carrier{}

	now := time.Now()
	milliTime := now.UnixNano() / 1e6
	principalID := GetPrincipalIDFromContext(ctx)

	tx := GetTransaction(ctx)
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	event := NewEvent(EventMetadata{
		Type:        EventTypeCreated,
		Entity:      "Carrier",
		EntityID:    item.ID,
		Date:        milliTime,
		PrincipalID: principalID,
	})

	var changes CarrierChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	err = CheckStructFieldIsEmpty(item, input)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	item.ID = uuid.Must(uuid.NewV4()).String()
	item.CreatedAt = milliTime
	item.CreatedBy = principalID

	if _, ok := input["carrierName"]; ok && (item.CarrierName != changes.CarrierName) {
		item.CarrierName = changes.CarrierName
		event.AddNewValue("carrierName", changes.CarrierName)
	}

	if _, ok := input["contactPerson"]; ok && (item.ContactPerson != changes.ContactPerson) && (item.ContactPerson == nil || changes.ContactPerson == nil || *item.ContactPerson != *changes.ContactPerson) {
		item.ContactPerson = changes.ContactPerson
		event.AddNewValue("contactPerson", changes.ContactPerson)
	}

	if _, ok := input["contactInfo"]; ok && (item.ContactInfo != changes.ContactInfo) && (item.ContactInfo == nil || changes.ContactInfo == nil || *item.ContactInfo != *changes.ContactInfo) {
		item.ContactInfo = changes.ContactInfo
		event.AddNewValue("contactInfo", changes.ContactInfo)
	}

	if err := tx.Omit(clause.Associations).Create(item).Error; err != nil {
		tx.Rollback()
		return item, err
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}
func (r *GeneratedMutationResolver) UpdateCarrier(ctx context.Context, id string, input map[string]interface{}) (item *Carrier, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdateCarrier(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func UpdateCarrierHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Carrier, err error) {
	item = &Carrier{}
	newItem := &Carrier{}

	isChange := false
	now := time.Now()
	milliTime := now.UnixNano() / 1e6
	principalID := GetPrincipalIDFromContext(ctx)

	tx := GetTransaction(ctx)
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	event := NewEvent(EventMetadata{
		Type:        EventTypeUpdated,
		Entity:      "Carrier",
		EntityID:    id,
		Date:        milliTime,
		PrincipalID: principalID,
	})

	var changes CarrierChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	err = CheckStructFieldIsEmpty(item, input)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err = GetItem(ctx, tx, TableName("carriers"), item, &id); err != nil {
		tx.Rollback()
		return nil, err
	}

	if item.UpdatedBy != nil && principalID != nil && *item.UpdatedBy != *principalID {
		newItem.UpdatedBy = principalID
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		event.AddOldValue("id", item.ID)
		event.AddNewValue("id", changes.ID)
		item.ID = changes.ID
		newItem.ID = changes.ID
		isChange = true
	}

	if _, ok := input["carrierName"]; ok && (item.CarrierName != changes.CarrierName) {
		event.AddOldValue("carrierName", item.CarrierName)
		event.AddNewValue("carrierName", changes.CarrierName)
		item.CarrierName = changes.CarrierName
		newItem.CarrierName = changes.CarrierName
		isChange = true
	}

	if _, ok := input["contactPerson"]; ok && (item.ContactPerson != changes.ContactPerson) && (item.ContactPerson == nil || changes.ContactPerson == nil || *item.ContactPerson != *changes.ContactPerson) {
		event.AddOldValue("contactPerson", item.ContactPerson)
		event.AddNewValue("contactPerson", changes.ContactPerson)
		item.ContactPerson = changes.ContactPerson
		newItem.ContactPerson = changes.ContactPerson
		isChange = true
	}

	if _, ok := input["contactInfo"]; ok && (item.ContactInfo != changes.ContactInfo) && (item.ContactInfo == nil || changes.ContactInfo == nil || *item.ContactInfo != *changes.ContactInfo) {
		event.AddOldValue("contactInfo", item.ContactInfo)
		event.AddNewValue("contactInfo", changes.ContactInfo)
		item.ContactInfo = changes.ContactInfo
		newItem.ContactInfo = changes.ContactInfo
		isChange = true
	}

	// if err := validator.Struct(item); err != nil {
	// 	tx.Rollback()
	// 	return nil, err
	// }

	if !isChange {
		return item, nil
	}

	if err := tx.Model(&newItem).Where("id = ?", id).Omit(clause.Associations).Updates(newItem).Error; err != nil {
		tx.Rollback()
		return item, err
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}

func DeleteCarrierFunc(ctx context.Context, r *GeneratedResolver, id string, tye string, unscoped *bool) (err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item := &Carrier{}
	now := time.Now()
	tx := GetTransaction(ctx)
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err = GetItem(ctx, tx, TableName("carriers"), item, &id); err != nil {
		tx.Rollback()
		return err
	}

	deletedAt := now.UnixNano() / 1e6

	event := NewEvent(EventMetadata{
		Type:        EventTypeDeleted,
		Entity:      "Carrier",
		EntityID:    id,
		Date:        deletedAt,
		PrincipalID: principalID,
	})

	// 如果是恢复删除数据
	if tye == "recovery" {
		if err := tx.Unscoped().Model(&item).Updates(map[string]interface{}{"DeletedAt": nil, "DeletedBy": nil}).Error; err != nil {
			tx.Rollback()
			return err
		}
	} else {
		if unscoped != nil && *unscoped == true {
			if err := tx.Unscoped().Model(&item).Delete(item).Error; err != nil {
				tx.Rollback()
				return err
			}
		} else if err := tx.Model(&item).Updates(Carrier{DeletedAt: &deletedAt, DeletedBy: principalID, UpdatedBy: principalID}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}

func (r *GeneratedMutationResolver) DeleteCarriers(ctx context.Context, id []string, unscoped *bool) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteCarriers(ctx, r.GeneratedResolver, id, unscoped)
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

func DeleteCarriersHandler(ctx context.Context, r *GeneratedResolver, id []string, unscoped *bool) (bool, error) {
	var err error = nil

	if len(id) > 0 {
		for _, v := range id {
			err = DeleteCarrierFunc(ctx, r, v, "delete", unscoped)
			if err != nil {
				break
			}
		}
	}

	if err != nil {
		return false, err
	}
	return true, err
}

func (r *GeneratedMutationResolver) RecoveryCarriers(ctx context.Context, id []string) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.RecoveryCarriers(ctx, r.GeneratedResolver, id)
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

func RecoveryCarriersHandler(ctx context.Context, r *GeneratedResolver, id []string) (bool, error) {
	var err error = nil

	var unscoped bool = false

	if len(id) > 0 {
		for _, v := range id {
			err = DeleteCarrierFunc(ctx, r, v, "recovery", &unscoped)
			if err != nil {
				break
			}
		}
	}

	if err != nil {
		return false, err
	}
	return true, err
}

func (r *GeneratedMutationResolver) CreateLocation(ctx context.Context, input map[string]interface{}) (item *Location, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreateLocation(ctx, r.GeneratedResolver, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func CreateLocationHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Location, err error) {
	item = &Location{}

	now := time.Now()
	milliTime := now.UnixNano() / 1e6
	principalID := GetPrincipalIDFromContext(ctx)

	tx := GetTransaction(ctx)
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	event := NewEvent(EventMetadata{
		Type:        EventTypeCreated,
		Entity:      "Location",
		EntityID:    item.ID,
		Date:        milliTime,
		PrincipalID: principalID,
	})

	var changes LocationChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	err = CheckStructFieldIsEmpty(item, input)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	item.ID = uuid.Must(uuid.NewV4()).String()
	item.CreatedAt = milliTime
	item.CreatedBy = principalID

	if input["startShipments"] != nil && input["startShipmentsIds"] != nil {
		tx.Rollback()
		return nil, fmt.Errorf("startShipmentsIds and startShipments cannot coexist")
	}

	var startShipmentsIds []string

	if _, ok := input["startShipments"]; ok {
		var startShipmentsMaps []map[string]interface{}
		for _, v := range input["startShipments"].([]interface{}) {
			startShipmentsMaps = append(startShipmentsMaps, v.(map[string]interface{}))
		}

		for _, v := range startShipmentsMaps {
			var startShipments *Shipment
			if v["id"] == nil {
				startShipments, err = r.Handlers.CreateShipment(ctx, r, v)
			} else {
				startShipments, err = r.Handlers.UpdateShipment(ctx, r, v["id"].(string), v)
			}

			changes.StartShipments = append(changes.StartShipments, startShipments)
			startShipmentsIds = append(startShipmentsIds, startShipments.ID)
		}
		event.AddNewValue("startShipments", changes.StartShipments)
		item.StartShipments = changes.StartShipments
	}

	if ids, exists := input["startShipmentsIds"]; exists {
		for _, v := range ids.([]interface{}) {
			startShipmentsIds = append(startShipmentsIds, v.(string))
		}
	}

	if input["endShipments"] != nil && input["endShipmentsIds"] != nil {
		tx.Rollback()
		return nil, fmt.Errorf("endShipmentsIds and endShipments cannot coexist")
	}

	var endShipmentsIds []string

	if _, ok := input["endShipments"]; ok {
		var endShipmentsMaps []map[string]interface{}
		for _, v := range input["endShipments"].([]interface{}) {
			endShipmentsMaps = append(endShipmentsMaps, v.(map[string]interface{}))
		}

		for _, v := range endShipmentsMaps {
			var endShipments *Shipment
			if v["id"] == nil {
				endShipments, err = r.Handlers.CreateShipment(ctx, r, v)
			} else {
				endShipments, err = r.Handlers.UpdateShipment(ctx, r, v["id"].(string), v)
			}

			changes.EndShipments = append(changes.EndShipments, endShipments)
			endShipmentsIds = append(endShipmentsIds, endShipments.ID)
		}
		event.AddNewValue("endShipments", changes.EndShipments)
		item.EndShipments = changes.EndShipments
	}

	if ids, exists := input["endShipmentsIds"]; exists {
		for _, v := range ids.([]interface{}) {
			endShipmentsIds = append(endShipmentsIds, v.(string))
		}
	}

	if _, ok := input["warehouseAddress"]; ok && (item.WarehouseAddress != changes.WarehouseAddress) && (item.WarehouseAddress == nil || changes.WarehouseAddress == nil || *item.WarehouseAddress != *changes.WarehouseAddress) {
		item.WarehouseAddress = changes.WarehouseAddress
		event.AddNewValue("warehouseAddress", changes.WarehouseAddress)
	}

	if _, ok := input["loadingAddress"]; ok && (item.LoadingAddress != changes.LoadingAddress) && (item.LoadingAddress == nil || changes.LoadingAddress == nil || *item.LoadingAddress != *changes.LoadingAddress) {
		item.LoadingAddress = changes.LoadingAddress
		event.AddNewValue("loadingAddress", changes.LoadingAddress)
	}

	if err := tx.Omit(clause.Associations).Create(item).Error; err != nil {
		tx.Rollback()
		return item, err
	}

	if len(startShipmentsIds) > 0 {
		if err := tx.Model(&Shipment{}).Where("id IN(?)", startShipmentsIds).Update("start_location_id", item.ID).Error; err != nil {
			tx.Rollback()
			return item, err
		}
	}

	if len(endShipmentsIds) > 0 {
		if err := tx.Model(&Shipment{}).Where("id IN(?)", endShipmentsIds).Update("end_location_id", item.ID).Error; err != nil {
			tx.Rollback()
			return item, err
		}
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}
func (r *GeneratedMutationResolver) UpdateLocation(ctx context.Context, id string, input map[string]interface{}) (item *Location, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdateLocation(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func UpdateLocationHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Location, err error) {
	item = &Location{}
	newItem := &Location{}

	isChange := false
	now := time.Now()
	milliTime := now.UnixNano() / 1e6
	principalID := GetPrincipalIDFromContext(ctx)

	tx := GetTransaction(ctx)
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	event := NewEvent(EventMetadata{
		Type:        EventTypeUpdated,
		Entity:      "Location",
		EntityID:    id,
		Date:        milliTime,
		PrincipalID: principalID,
	})

	var changes LocationChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	err = CheckStructFieldIsEmpty(item, input)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if input["startShipments"] != nil && input["startShipmentsIds"] != nil {
		tx.Rollback()
		return nil, fmt.Errorf("startShipmentsIds and startShipments cannot coexist")
	}

	if input["endShipments"] != nil && input["endShipmentsIds"] != nil {
		tx.Rollback()
		return nil, fmt.Errorf("endShipmentsIds and endShipments cannot coexist")
	}

	if err = GetItem(ctx, tx, TableName("locations"), item, &id); err != nil {
		tx.Rollback()
		return nil, err
	}

	if item.UpdatedBy != nil && principalID != nil && *item.UpdatedBy != *principalID {
		newItem.UpdatedBy = principalID
	}

	var startShipmentsIds []string

	if _, ok := input["startShipments"]; ok {
		if err := tx.Unscoped().Model(&Shipment{}).Where("start_location_id = ?", id).Update("start_location_id", "").Error; err != nil {
			tx.Rollback()
			return item, err
		}

		var startShipmentsMaps []map[string]interface{}
		for _, v := range input["startShipments"].([]interface{}) {
			vMaps := v.(map[string]interface{})
			startShipmentsMaps = append(startShipmentsMaps, vMaps)
		}

		for _, v := range startShipmentsMaps {
			var startShipments *Shipment
			v["start_locationId"] = id
			if v["id"] == nil {
				startShipments, err = r.Handlers.CreateShipment(ctx, r, v)
			} else {
				startShipments, err = r.Handlers.UpdateShipment(ctx, r, v["id"].(string), v)
			}

			changes.StartShipments = append(changes.StartShipments, startShipments)
		}

		event.AddNewValue("startShipments", changes.StartShipments)
		item.StartShipments = changes.StartShipments
		newItem.StartShipments = changes.StartShipments
	}

	if ids, exists := input["startShipmentsIds"]; exists {
		if err := tx.Unscoped().Model(&Shipment{}).Where("start_location_id = ?", id).Update("start_location_id", "").Error; err != nil {
			tx.Rollback()
			return item, err
		}
		for _, v := range ids.([]interface{}) {
			startShipmentsIds = append(startShipmentsIds, v.(string))
		}

		if len(startShipmentsIds) > 0 {
			isChange = true
		}
	}

	var endShipmentsIds []string

	if _, ok := input["endShipments"]; ok {
		if err := tx.Unscoped().Model(&Shipment{}).Where("end_location_id = ?", id).Update("end_location_id", "").Error; err != nil {
			tx.Rollback()
			return item, err
		}

		var endShipmentsMaps []map[string]interface{}
		for _, v := range input["endShipments"].([]interface{}) {
			vMaps := v.(map[string]interface{})
			endShipmentsMaps = append(endShipmentsMaps, vMaps)
		}

		for _, v := range endShipmentsMaps {
			var endShipments *Shipment
			v["end_locationId"] = id
			if v["id"] == nil {
				endShipments, err = r.Handlers.CreateShipment(ctx, r, v)
			} else {
				endShipments, err = r.Handlers.UpdateShipment(ctx, r, v["id"].(string), v)
			}

			changes.EndShipments = append(changes.EndShipments, endShipments)
		}

		event.AddNewValue("endShipments", changes.EndShipments)
		item.EndShipments = changes.EndShipments
		newItem.EndShipments = changes.EndShipments
	}

	if ids, exists := input["endShipmentsIds"]; exists {
		if err := tx.Unscoped().Model(&Shipment{}).Where("end_location_id = ?", id).Update("end_location_id", "").Error; err != nil {
			tx.Rollback()
			return item, err
		}
		for _, v := range ids.([]interface{}) {
			endShipmentsIds = append(endShipmentsIds, v.(string))
		}

		if len(endShipmentsIds) > 0 {
			isChange = true
		}
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		event.AddOldValue("id", item.ID)
		event.AddNewValue("id", changes.ID)
		item.ID = changes.ID
		newItem.ID = changes.ID
		isChange = true
	}

	if _, ok := input["warehouseAddress"]; ok && (item.WarehouseAddress != changes.WarehouseAddress) && (item.WarehouseAddress == nil || changes.WarehouseAddress == nil || *item.WarehouseAddress != *changes.WarehouseAddress) {
		event.AddOldValue("warehouseAddress", item.WarehouseAddress)
		event.AddNewValue("warehouseAddress", changes.WarehouseAddress)
		item.WarehouseAddress = changes.WarehouseAddress
		newItem.WarehouseAddress = changes.WarehouseAddress
		isChange = true
	}

	if _, ok := input["loadingAddress"]; ok && (item.LoadingAddress != changes.LoadingAddress) && (item.LoadingAddress == nil || changes.LoadingAddress == nil || *item.LoadingAddress != *changes.LoadingAddress) {
		event.AddOldValue("loadingAddress", item.LoadingAddress)
		event.AddNewValue("loadingAddress", changes.LoadingAddress)
		item.LoadingAddress = changes.LoadingAddress
		newItem.LoadingAddress = changes.LoadingAddress
		isChange = true
	}

	// if err := validator.Struct(item); err != nil {
	// 	tx.Rollback()
	// 	return nil, err
	// }

	if !isChange {
		return item, nil
	}

	if err := tx.Model(&newItem).Where("id = ?", id).Omit(clause.Associations).Updates(newItem).Error; err != nil {
		tx.Rollback()
		return item, err
	}

	if len(startShipmentsIds) > 0 {
		if err := tx.Model(&Shipment{}).Where("id IN(?)", startShipmentsIds).Update("start_location_id", item.ID).Error; err != nil {
			tx.Rollback()
			return item, err
		}
	}

	if len(endShipmentsIds) > 0 {
		if err := tx.Model(&Shipment{}).Where("id IN(?)", endShipmentsIds).Update("end_location_id", item.ID).Error; err != nil {
			tx.Rollback()
			return item, err
		}
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}

func DeleteLocationFunc(ctx context.Context, r *GeneratedResolver, id string, tye string, unscoped *bool) (err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item := &Location{}
	now := time.Now()
	tx := GetTransaction(ctx)
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err = GetItem(ctx, tx, TableName("locations"), item, &id); err != nil {
		tx.Rollback()
		return err
	}

	deletedAt := now.UnixNano() / 1e6

	event := NewEvent(EventMetadata{
		Type:        EventTypeDeleted,
		Entity:      "Location",
		EntityID:    id,
		Date:        deletedAt,
		PrincipalID: principalID,
	})

	// 如果是恢复删除数据
	if tye == "recovery" {
		if err := tx.Unscoped().Model(&item).Updates(map[string]interface{}{"DeletedAt": nil, "DeletedBy": nil}).Error; err != nil {
			tx.Rollback()
			return err
		}
	} else {
		if unscoped != nil && *unscoped == true {
			if err := tx.Unscoped().Model(&item).Delete(item).Error; err != nil {
				tx.Rollback()
				return err
			}
		} else if err := tx.Model(&item).Updates(Location{DeletedAt: &deletedAt, DeletedBy: principalID, UpdatedBy: principalID}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}

func (r *GeneratedMutationResolver) DeleteLocations(ctx context.Context, id []string, unscoped *bool) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteLocations(ctx, r.GeneratedResolver, id, unscoped)
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

func DeleteLocationsHandler(ctx context.Context, r *GeneratedResolver, id []string, unscoped *bool) (bool, error) {
	var err error = nil

	if len(id) > 0 {
		for _, v := range id {
			err = DeleteLocationFunc(ctx, r, v, "delete", unscoped)
			if err != nil {
				break
			}
		}
	}

	if err != nil {
		return false, err
	}
	return true, err
}

func (r *GeneratedMutationResolver) RecoveryLocations(ctx context.Context, id []string) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.RecoveryLocations(ctx, r.GeneratedResolver, id)
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

func RecoveryLocationsHandler(ctx context.Context, r *GeneratedResolver, id []string) (bool, error) {
	var err error = nil

	var unscoped bool = false

	if len(id) > 0 {
		for _, v := range id {
			err = DeleteLocationFunc(ctx, r, v, "recovery", &unscoped)
			if err != nil {
				break
			}
		}
	}

	if err != nil {
		return false, err
	}
	return true, err
}

func (r *GeneratedMutationResolver) CreateEquipmentd(ctx context.Context, input map[string]interface{}) (item *Equipmentd, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreateEquipmentd(ctx, r.GeneratedResolver, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func CreateEquipmentdHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Equipmentd, err error) {
	item = &Equipmentd{}

	now := time.Now()
	milliTime := now.UnixNano() / 1e6
	principalID := GetPrincipalIDFromContext(ctx)

	tx := GetTransaction(ctx)
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	event := NewEvent(EventMetadata{
		Type:        EventTypeCreated,
		Entity:      "Equipmentd",
		EntityID:    item.ID,
		Date:        milliTime,
		PrincipalID: principalID,
	})

	var changes EquipmentdChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	err = CheckStructFieldIsEmpty(item, input)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	item.ID = uuid.Must(uuid.NewV4()).String()
	item.CreatedAt = milliTime
	item.CreatedBy = principalID

	if _, ok := input["vehicleType"]; ok && (item.VehicleType != changes.VehicleType) && (item.VehicleType == nil || changes.VehicleType == nil || *item.VehicleType != *changes.VehicleType) {
		item.VehicleType = changes.VehicleType
		event.AddNewValue("vehicleType", changes.VehicleType)
	}

	if _, ok := input["capacity"]; ok && (item.Capacity != changes.Capacity) && (item.Capacity == nil || changes.Capacity == nil || *item.Capacity != *changes.Capacity) {
		item.Capacity = changes.Capacity
		event.AddNewValue("capacity", changes.Capacity)
	}

	if err := tx.Omit(clause.Associations).Create(item).Error; err != nil {
		tx.Rollback()
		return item, err
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}
func (r *GeneratedMutationResolver) UpdateEquipmentd(ctx context.Context, id string, input map[string]interface{}) (item *Equipmentd, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdateEquipmentd(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func UpdateEquipmentdHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Equipmentd, err error) {
	item = &Equipmentd{}
	newItem := &Equipmentd{}

	isChange := false
	now := time.Now()
	milliTime := now.UnixNano() / 1e6
	principalID := GetPrincipalIDFromContext(ctx)

	tx := GetTransaction(ctx)
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	event := NewEvent(EventMetadata{
		Type:        EventTypeUpdated,
		Entity:      "Equipmentd",
		EntityID:    id,
		Date:        milliTime,
		PrincipalID: principalID,
	})

	var changes EquipmentdChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	err = CheckStructFieldIsEmpty(item, input)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err = GetItem(ctx, tx, TableName("equipmentds"), item, &id); err != nil {
		tx.Rollback()
		return nil, err
	}

	if item.UpdatedBy != nil && principalID != nil && *item.UpdatedBy != *principalID {
		newItem.UpdatedBy = principalID
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		event.AddOldValue("id", item.ID)
		event.AddNewValue("id", changes.ID)
		item.ID = changes.ID
		newItem.ID = changes.ID
		isChange = true
	}

	if _, ok := input["vehicleType"]; ok && (item.VehicleType != changes.VehicleType) && (item.VehicleType == nil || changes.VehicleType == nil || *item.VehicleType != *changes.VehicleType) {
		event.AddOldValue("vehicleType", item.VehicleType)
		event.AddNewValue("vehicleType", changes.VehicleType)
		item.VehicleType = changes.VehicleType
		newItem.VehicleType = changes.VehicleType
		isChange = true
	}

	if _, ok := input["capacity"]; ok && (item.Capacity != changes.Capacity) && (item.Capacity == nil || changes.Capacity == nil || *item.Capacity != *changes.Capacity) {
		event.AddOldValue("capacity", item.Capacity)
		event.AddNewValue("capacity", changes.Capacity)
		item.Capacity = changes.Capacity
		newItem.Capacity = changes.Capacity
		isChange = true
	}

	// if err := validator.Struct(item); err != nil {
	// 	tx.Rollback()
	// 	return nil, err
	// }

	if !isChange {
		return item, nil
	}

	if err := tx.Model(&newItem).Where("id = ?", id).Omit(clause.Associations).Updates(newItem).Error; err != nil {
		tx.Rollback()
		return item, err
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}

func DeleteEquipmentdFunc(ctx context.Context, r *GeneratedResolver, id string, tye string, unscoped *bool) (err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item := &Equipmentd{}
	now := time.Now()
	tx := GetTransaction(ctx)
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err = GetItem(ctx, tx, TableName("equipmentds"), item, &id); err != nil {
		tx.Rollback()
		return err
	}

	deletedAt := now.UnixNano() / 1e6

	event := NewEvent(EventMetadata{
		Type:        EventTypeDeleted,
		Entity:      "Equipmentd",
		EntityID:    id,
		Date:        deletedAt,
		PrincipalID: principalID,
	})

	// 如果是恢复删除数据
	if tye == "recovery" {
		if err := tx.Unscoped().Model(&item).Updates(map[string]interface{}{"DeletedAt": nil, "DeletedBy": nil}).Error; err != nil {
			tx.Rollback()
			return err
		}
	} else {
		if unscoped != nil && *unscoped == true {
			if err := tx.Unscoped().Model(&item).Delete(item).Error; err != nil {
				tx.Rollback()
				return err
			}
		} else if err := tx.Model(&item).Updates(Equipmentd{DeletedAt: &deletedAt, DeletedBy: principalID, UpdatedBy: principalID}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}

func (r *GeneratedMutationResolver) DeleteEquipmentds(ctx context.Context, id []string, unscoped *bool) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteEquipmentds(ctx, r.GeneratedResolver, id, unscoped)
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

func DeleteEquipmentdsHandler(ctx context.Context, r *GeneratedResolver, id []string, unscoped *bool) (bool, error) {
	var err error = nil

	if len(id) > 0 {
		for _, v := range id {
			err = DeleteEquipmentdFunc(ctx, r, v, "delete", unscoped)
			if err != nil {
				break
			}
		}
	}

	if err != nil {
		return false, err
	}
	return true, err
}

func (r *GeneratedMutationResolver) RecoveryEquipmentds(ctx context.Context, id []string) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.RecoveryEquipmentds(ctx, r.GeneratedResolver, id)
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

func RecoveryEquipmentdsHandler(ctx context.Context, r *GeneratedResolver, id []string) (bool, error) {
	var err error = nil

	var unscoped bool = false

	if len(id) > 0 {
		for _, v := range id {
			err = DeleteEquipmentdFunc(ctx, r, v, "recovery", &unscoped)
			if err != nil {
				break
			}
		}
	}

	if err != nil {
		return false, err
	}
	return true, err
}
