package gen

import (
	"context"
)

func (s OrderSortType) Apply(ctx context.Context, sorts *[]string, joins *[]string) error {
	return s.ApplyWithAlias(ctx, TableName("orders"), sorts, joins)
}
func (s OrderSortType) ApplyWithAlias(ctx context.Context, alias string, sorts *[]string, joins *[]string) error {
	aliasPrefix := alias + "."

	if s.ID != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("id")+" "+s.ID.String())
	}

	if s.OrderNo != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("orderNo")+" "+s.OrderNo.String())
	}

	if s.CustomerInfo != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("customerInfo")+" "+s.CustomerInfo.String())
	}

	if s.GoodsInfo != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("goodsInfo")+" "+s.GoodsInfo.String())
	}

	if s.DeletedBy != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("deletedBy")+" "+s.DeletedBy.String())
	}

	if s.UpdatedBy != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("updatedBy")+" "+s.UpdatedBy.String())
	}

	if s.CreatedBy != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("createdBy")+" "+s.CreatedBy.String())
	}

	if s.DeletedAt != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("deletedAt")+" "+s.DeletedAt.String())
	}

	if s.UpdatedAt != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("updatedAt")+" "+s.UpdatedAt.String())
	}

	if s.CreatedAt != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("createdAt")+" "+s.CreatedAt.String())
	}

	return nil
}

func (s ShipmentSortType) Apply(ctx context.Context, sorts *[]string, joins *[]string) error {
	return s.ApplyWithAlias(ctx, TableName("shipments"), sorts, joins)
}
func (s ShipmentSortType) ApplyWithAlias(ctx context.Context, alias string, sorts *[]string, joins *[]string) error {
	aliasPrefix := alias + "."

	if s.ID != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("id")+" "+s.ID.String())
	}

	if s.ShipmentNo != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("shipmentNo")+" "+s.ShipmentNo.String())
	}

	if s.TransportationMode != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("transportationMode")+" "+s.TransportationMode.String())
	}

	if s.StartLocationID != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("startLocationId")+" "+s.StartLocationID.String())
	}

	if s.EndLocationID != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("endLocationId")+" "+s.EndLocationID.String())
	}

	if s.DeletedBy != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("deletedBy")+" "+s.DeletedBy.String())
	}

	if s.UpdatedBy != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("updatedBy")+" "+s.UpdatedBy.String())
	}

	if s.CreatedBy != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("createdBy")+" "+s.CreatedBy.String())
	}

	if s.DeletedAt != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("deletedAt")+" "+s.DeletedAt.String())
	}

	if s.UpdatedAt != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("updatedAt")+" "+s.UpdatedAt.String())
	}

	if s.CreatedAt != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("createdAt")+" "+s.CreatedAt.String())
	}

	if s.StartLocation != nil {
		_alias := alias + "_startLocation"
		*joins = append(*joins, "LEFT JOIN "+"locations"+" "+_alias+" ON "+_alias+".id = "+alias+"."+"start_location_id")
		err := s.StartLocation.ApplyWithAlias(ctx, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	if s.EndLocation != nil {
		_alias := alias + "_endLocation"
		*joins = append(*joins, "LEFT JOIN "+"locations"+" "+_alias+" ON "+_alias+".id = "+alias+"."+"end_location_id")
		err := s.EndLocation.ApplyWithAlias(ctx, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s CarrierSortType) Apply(ctx context.Context, sorts *[]string, joins *[]string) error {
	return s.ApplyWithAlias(ctx, TableName("carriers"), sorts, joins)
}
func (s CarrierSortType) ApplyWithAlias(ctx context.Context, alias string, sorts *[]string, joins *[]string) error {
	aliasPrefix := alias + "."

	if s.ID != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("id")+" "+s.ID.String())
	}

	if s.CarrierName != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("carrierName")+" "+s.CarrierName.String())
	}

	if s.ContactPerson != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("contactPerson")+" "+s.ContactPerson.String())
	}

	if s.ContactInfo != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("contactInfo")+" "+s.ContactInfo.String())
	}

	if s.DeletedBy != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("deletedBy")+" "+s.DeletedBy.String())
	}

	if s.UpdatedBy != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("updatedBy")+" "+s.UpdatedBy.String())
	}

	if s.CreatedBy != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("createdBy")+" "+s.CreatedBy.String())
	}

	if s.DeletedAt != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("deletedAt")+" "+s.DeletedAt.String())
	}

	if s.UpdatedAt != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("updatedAt")+" "+s.UpdatedAt.String())
	}

	if s.CreatedAt != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("createdAt")+" "+s.CreatedAt.String())
	}

	return nil
}

func (s LocationSortType) Apply(ctx context.Context, sorts *[]string, joins *[]string) error {
	return s.ApplyWithAlias(ctx, TableName("locations"), sorts, joins)
}
func (s LocationSortType) ApplyWithAlias(ctx context.Context, alias string, sorts *[]string, joins *[]string) error {
	aliasPrefix := alias + "."

	if s.ID != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("id")+" "+s.ID.String())
	}

	if s.WarehouseAddress != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("warehouseAddress")+" "+s.WarehouseAddress.String())
	}

	if s.LoadingAddress != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("loadingAddress")+" "+s.LoadingAddress.String())
	}

	if s.DeletedBy != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("deletedBy")+" "+s.DeletedBy.String())
	}

	if s.UpdatedBy != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("updatedBy")+" "+s.UpdatedBy.String())
	}

	if s.CreatedBy != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("createdBy")+" "+s.CreatedBy.String())
	}

	if s.DeletedAt != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("deletedAt")+" "+s.DeletedAt.String())
	}

	if s.UpdatedAt != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("updatedAt")+" "+s.UpdatedAt.String())
	}

	if s.CreatedAt != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("createdAt")+" "+s.CreatedAt.String())
	}

	if s.StartShipments != nil {
		_alias := alias + "_startShipments"
		*joins = append(*joins, "LEFT JOIN "+"shipments"+" "+_alias+" ON "+_alias+"."+"start_location_id"+" = "+alias+".id")
		err := s.StartShipments.ApplyWithAlias(ctx, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	if s.EndShipments != nil {
		_alias := alias + "_endShipments"
		*joins = append(*joins, "LEFT JOIN "+"shipments"+" "+_alias+" ON "+_alias+"."+"end_location_id"+" = "+alias+".id")
		err := s.EndShipments.ApplyWithAlias(ctx, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s EquipmentdSortType) Apply(ctx context.Context, sorts *[]string, joins *[]string) error {
	return s.ApplyWithAlias(ctx, TableName("equipmentds"), sorts, joins)
}
func (s EquipmentdSortType) ApplyWithAlias(ctx context.Context, alias string, sorts *[]string, joins *[]string) error {
	aliasPrefix := alias + "."

	if s.ID != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("id")+" "+s.ID.String())
	}

	if s.VehicleType != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("vehicleType")+" "+s.VehicleType.String())
	}

	if s.Capacity != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("capacity")+" "+s.Capacity.String())
	}

	if s.DeletedBy != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("deletedBy")+" "+s.DeletedBy.String())
	}

	if s.UpdatedBy != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("updatedBy")+" "+s.UpdatedBy.String())
	}

	if s.CreatedBy != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("createdBy")+" "+s.CreatedBy.String())
	}

	if s.DeletedAt != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("deletedAt")+" "+s.DeletedAt.String())
	}

	if s.UpdatedAt != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("updatedAt")+" "+s.UpdatedAt.String())
	}

	if s.CreatedAt != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("createdAt")+" "+s.CreatedAt.String())
	}

	return nil
}
