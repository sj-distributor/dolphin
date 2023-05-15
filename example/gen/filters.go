package gen

import (
	"context"
	"fmt"
	"strings"
)

func (f *OrderFilterType) IsEmpty(ctx context.Context) bool {
	wheres := []string{}
	values := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, "companies", &wheres, &values, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0
}
func (f *OrderFilterType) Apply(ctx context.Context, wheres *[]string, values *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, TableName("orders"), wheres, values, joins)
}
func (f *OrderFilterType) ApplyWithAlias(ctx context.Context, alias string, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := alias + "."

	_where, _values := f.WhereContent(aliasPrefix)
	*wheres = append(*wheres, _where...)
	*values = append(*values, _values...)

	if f.Or != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			_cs := []string{}
			err := or.ApplyWithAlias(ctx, alias, &_cs, &vs, &js)
			if err != nil {
				return err
			}
			cs = append(cs, strings.Join(_cs, " AND "))
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, "("+strings.Join(cs, " OR ")+")")
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, alias, &cs, &vs, &js)
			if err != nil {
				return err
			}
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, strings.Join(cs, " AND "))
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}

	return nil
}

func (f *OrderFilterType) WhereContent(aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" = ?")
		values = append(values, f.ID)
	}

	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" != ?")
		values = append(values, f.IDNe)
	}

	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" > ?")
		values = append(values, f.IDGt)
	}

	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" < ?")
		values = append(values, f.IDLt)
	}

	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" >= ?")
		values = append(values, f.IDGte)
	}

	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" <= ?")
		values = append(values, f.IDLte)
	}

	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" IN (?)")
		values = append(values, f.IDIn)
	}

	if f.IDNull != nil {
		if *f.IDNull {
			conditions = append(conditions, aliasPrefix+SnakeString("id")+" IS NULL"+" OR "+aliasPrefix+SnakeString("id")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("id")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("id")+" <> ''")
		}
	}

	if f.OrderNo != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("orderNo")+" = ?")
		values = append(values, f.OrderNo)
	}

	if f.OrderNoNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("orderNo")+" != ?")
		values = append(values, f.OrderNoNe)
	}

	if f.OrderNoGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("orderNo")+" > ?")
		values = append(values, f.OrderNoGt)
	}

	if f.OrderNoLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("orderNo")+" < ?")
		values = append(values, f.OrderNoLt)
	}

	if f.OrderNoGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("orderNo")+" >= ?")
		values = append(values, f.OrderNoGte)
	}

	if f.OrderNoLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("orderNo")+" <= ?")
		values = append(values, f.OrderNoLte)
	}

	if f.OrderNoIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("orderNo")+" IN (?)")
		values = append(values, f.OrderNoIn)
	}

	if f.OrderNoLike != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("orderNo")+" LIKE ?")
		values = append(values, "%"+strings.Replace(strings.Replace(*f.OrderNoLike, "?", "_", -1), "*", "%", -1)+"%")
	}

	if f.OrderNoPrefix != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("orderNo")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.OrderNoPrefix))
	}

	if f.OrderNoSuffix != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("orderNo")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.OrderNoSuffix))
	}

	if f.OrderNoNull != nil {
		if *f.OrderNoNull {
			conditions = append(conditions, aliasPrefix+SnakeString("orderNo")+" IS NULL"+" OR "+aliasPrefix+SnakeString("orderNo")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("orderNo")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("orderNo")+" <> ''")
		}
	}

	if f.CustomerInfo != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("customerInfo")+" = ?")
		values = append(values, f.CustomerInfo)
	}

	if f.CustomerInfoNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("customerInfo")+" != ?")
		values = append(values, f.CustomerInfoNe)
	}

	if f.CustomerInfoGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("customerInfo")+" > ?")
		values = append(values, f.CustomerInfoGt)
	}

	if f.CustomerInfoLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("customerInfo")+" < ?")
		values = append(values, f.CustomerInfoLt)
	}

	if f.CustomerInfoGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("customerInfo")+" >= ?")
		values = append(values, f.CustomerInfoGte)
	}

	if f.CustomerInfoLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("customerInfo")+" <= ?")
		values = append(values, f.CustomerInfoLte)
	}

	if f.CustomerInfoIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("customerInfo")+" IN (?)")
		values = append(values, f.CustomerInfoIn)
	}

	if f.CustomerInfoLike != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("customerInfo")+" LIKE ?")
		values = append(values, "%"+strings.Replace(strings.Replace(*f.CustomerInfoLike, "?", "_", -1), "*", "%", -1)+"%")
	}

	if f.CustomerInfoPrefix != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("customerInfo")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.CustomerInfoPrefix))
	}

	if f.CustomerInfoSuffix != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("customerInfo")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.CustomerInfoSuffix))
	}

	if f.CustomerInfoNull != nil {
		if *f.CustomerInfoNull {
			conditions = append(conditions, aliasPrefix+SnakeString("customerInfo")+" IS NULL"+" OR "+aliasPrefix+SnakeString("customerInfo")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("customerInfo")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("customerInfo")+" <> ''")
		}
	}

	if f.GoodsInfo != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("goodsInfo")+" = ?")
		values = append(values, f.GoodsInfo)
	}

	if f.GoodsInfoNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("goodsInfo")+" != ?")
		values = append(values, f.GoodsInfoNe)
	}

	if f.GoodsInfoGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("goodsInfo")+" > ?")
		values = append(values, f.GoodsInfoGt)
	}

	if f.GoodsInfoLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("goodsInfo")+" < ?")
		values = append(values, f.GoodsInfoLt)
	}

	if f.GoodsInfoGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("goodsInfo")+" >= ?")
		values = append(values, f.GoodsInfoGte)
	}

	if f.GoodsInfoLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("goodsInfo")+" <= ?")
		values = append(values, f.GoodsInfoLte)
	}

	if f.GoodsInfoIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("goodsInfo")+" IN (?)")
		values = append(values, f.GoodsInfoIn)
	}

	if f.GoodsInfoLike != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("goodsInfo")+" LIKE ?")
		values = append(values, "%"+strings.Replace(strings.Replace(*f.GoodsInfoLike, "?", "_", -1), "*", "%", -1)+"%")
	}

	if f.GoodsInfoPrefix != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("goodsInfo")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.GoodsInfoPrefix))
	}

	if f.GoodsInfoSuffix != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("goodsInfo")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.GoodsInfoSuffix))
	}

	if f.GoodsInfoNull != nil {
		if *f.GoodsInfoNull {
			conditions = append(conditions, aliasPrefix+SnakeString("goodsInfo")+" IS NULL"+" OR "+aliasPrefix+SnakeString("goodsInfo")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("goodsInfo")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("goodsInfo")+" <> ''")
		}
	}

	if f.DeletedBy != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" = ?")
		values = append(values, f.DeletedBy)
	}

	if f.DeletedByNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" != ?")
		values = append(values, f.DeletedByNe)
	}

	if f.DeletedByGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" > ?")
		values = append(values, f.DeletedByGt)
	}

	if f.DeletedByLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" < ?")
		values = append(values, f.DeletedByLt)
	}

	if f.DeletedByGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" >= ?")
		values = append(values, f.DeletedByGte)
	}

	if f.DeletedByLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" <= ?")
		values = append(values, f.DeletedByLte)
	}

	if f.DeletedByIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" IN (?)")
		values = append(values, f.DeletedByIn)
	}

	if f.DeletedByNull != nil {
		if *f.DeletedByNull {
			conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" IS NULL"+" OR "+aliasPrefix+SnakeString("deletedBy")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("deletedBy")+" <> ''")
		}
	}

	if f.UpdatedBy != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" = ?")
		values = append(values, f.UpdatedBy)
	}

	if f.UpdatedByNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" != ?")
		values = append(values, f.UpdatedByNe)
	}

	if f.UpdatedByGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" > ?")
		values = append(values, f.UpdatedByGt)
	}

	if f.UpdatedByLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" < ?")
		values = append(values, f.UpdatedByLt)
	}

	if f.UpdatedByGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" >= ?")
		values = append(values, f.UpdatedByGte)
	}

	if f.UpdatedByLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" <= ?")
		values = append(values, f.UpdatedByLte)
	}

	if f.UpdatedByIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" IN (?)")
		values = append(values, f.UpdatedByIn)
	}

	if f.UpdatedByNull != nil {
		if *f.UpdatedByNull {
			conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" IS NULL"+" OR "+aliasPrefix+SnakeString("updatedBy")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("updatedBy")+" <> ''")
		}
	}

	if f.CreatedBy != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" = ?")
		values = append(values, f.CreatedBy)
	}

	if f.CreatedByNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" != ?")
		values = append(values, f.CreatedByNe)
	}

	if f.CreatedByGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" > ?")
		values = append(values, f.CreatedByGt)
	}

	if f.CreatedByLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" < ?")
		values = append(values, f.CreatedByLt)
	}

	if f.CreatedByGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" >= ?")
		values = append(values, f.CreatedByGte)
	}

	if f.CreatedByLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" <= ?")
		values = append(values, f.CreatedByLte)
	}

	if f.CreatedByIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" IN (?)")
		values = append(values, f.CreatedByIn)
	}

	if f.CreatedByNull != nil {
		if *f.CreatedByNull {
			conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" IS NULL"+" OR "+aliasPrefix+SnakeString("createdBy")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("createdBy")+" <> ''")
		}
	}

	if f.DeletedAt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" = ?")
		values = append(values, f.DeletedAt)
	}

	if f.DeletedAtNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" != ?")
		values = append(values, f.DeletedAtNe)
	}

	if f.DeletedAtGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" > ?")
		values = append(values, f.DeletedAtGt)
	}

	if f.DeletedAtLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" < ?")
		values = append(values, f.DeletedAtLt)
	}

	if f.DeletedAtGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" >= ?")
		values = append(values, f.DeletedAtGte)
	}

	if f.DeletedAtLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" <= ?")
		values = append(values, f.DeletedAtLte)
	}

	if f.DeletedAtIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" IN (?)")
		values = append(values, f.DeletedAtIn)
	}

	if f.DeletedAtNull != nil {
		if *f.DeletedAtNull {
			conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" IS NULL"+" OR "+aliasPrefix+SnakeString("deletedAt")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("deletedAt")+" <> ''")
		}
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" = ?")
		values = append(values, f.UpdatedAt)
	}

	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" != ?")
		values = append(values, f.UpdatedAtNe)
	}

	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" > ?")
		values = append(values, f.UpdatedAtGt)
	}

	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" < ?")
		values = append(values, f.UpdatedAtLt)
	}

	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" >= ?")
		values = append(values, f.UpdatedAtGte)
	}

	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" <= ?")
		values = append(values, f.UpdatedAtLte)
	}

	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.UpdatedAtNull != nil {
		if *f.UpdatedAtNull {
			conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" IS NULL"+" OR "+aliasPrefix+SnakeString("updatedAt")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("updatedAt")+" <> ''")
		}
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" = ?")
		values = append(values, f.CreatedAt)
	}

	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" != ?")
		values = append(values, f.CreatedAtNe)
	}

	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" > ?")
		values = append(values, f.CreatedAtGt)
	}

	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" < ?")
		values = append(values, f.CreatedAtLt)
	}

	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" >= ?")
		values = append(values, f.CreatedAtGte)
	}

	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" <= ?")
		values = append(values, f.CreatedAtLte)
	}

	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	if f.CreatedAtNull != nil {
		if *f.CreatedAtNull {
			conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" IS NULL"+" OR "+aliasPrefix+SnakeString("createdAt")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("createdAt")+" <> ''")
		}
	}

	return
}

// AndWith convenience method for combining two or more filters with AND statement
func (f *OrderFilterType) AndWith(f2 ...*OrderFilterType) *OrderFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &OrderFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *OrderFilterType) OrWith(f2 ...*OrderFilterType) *OrderFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &OrderFilterType{
		Or: append(_f2, f),
	}
}

func (f *ShipmentFilterType) IsEmpty(ctx context.Context) bool {
	wheres := []string{}
	values := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, "companies", &wheres, &values, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0
}
func (f *ShipmentFilterType) Apply(ctx context.Context, wheres *[]string, values *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, TableName("shipments"), wheres, values, joins)
}
func (f *ShipmentFilterType) ApplyWithAlias(ctx context.Context, alias string, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := alias + "."

	_where, _values := f.WhereContent(aliasPrefix)
	*wheres = append(*wheres, _where...)
	*values = append(*values, _values...)

	if f.Or != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			_cs := []string{}
			err := or.ApplyWithAlias(ctx, alias, &_cs, &vs, &js)
			if err != nil {
				return err
			}
			cs = append(cs, strings.Join(_cs, " AND "))
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, "("+strings.Join(cs, " OR ")+")")
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, alias, &cs, &vs, &js)
			if err != nil {
				return err
			}
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, strings.Join(cs, " AND "))
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}

	if f.StartLocation != nil {
		_alias := alias + "_startLocation"
		*joins = append(*joins, "LEFT JOIN "+"locations"+" "+_alias+" ON "+_alias+".id = "+alias+"."+"start_location_id")
		err := f.StartLocation.ApplyWithAlias(ctx, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	if f.EndLocation != nil {
		_alias := alias + "_endLocation"
		*joins = append(*joins, "LEFT JOIN "+"locations"+" "+_alias+" ON "+_alias+".id = "+alias+"."+"end_location_id")
		err := f.EndLocation.ApplyWithAlias(ctx, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *ShipmentFilterType) WhereContent(aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" = ?")
		values = append(values, f.ID)
	}

	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" != ?")
		values = append(values, f.IDNe)
	}

	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" > ?")
		values = append(values, f.IDGt)
	}

	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" < ?")
		values = append(values, f.IDLt)
	}

	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" >= ?")
		values = append(values, f.IDGte)
	}

	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" <= ?")
		values = append(values, f.IDLte)
	}

	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" IN (?)")
		values = append(values, f.IDIn)
	}

	if f.IDNull != nil {
		if *f.IDNull {
			conditions = append(conditions, aliasPrefix+SnakeString("id")+" IS NULL"+" OR "+aliasPrefix+SnakeString("id")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("id")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("id")+" <> ''")
		}
	}

	if f.ShipmentNo != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("shipmentNo")+" = ?")
		values = append(values, f.ShipmentNo)
	}

	if f.ShipmentNoNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("shipmentNo")+" != ?")
		values = append(values, f.ShipmentNoNe)
	}

	if f.ShipmentNoGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("shipmentNo")+" > ?")
		values = append(values, f.ShipmentNoGt)
	}

	if f.ShipmentNoLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("shipmentNo")+" < ?")
		values = append(values, f.ShipmentNoLt)
	}

	if f.ShipmentNoGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("shipmentNo")+" >= ?")
		values = append(values, f.ShipmentNoGte)
	}

	if f.ShipmentNoLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("shipmentNo")+" <= ?")
		values = append(values, f.ShipmentNoLte)
	}

	if f.ShipmentNoIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("shipmentNo")+" IN (?)")
		values = append(values, f.ShipmentNoIn)
	}

	if f.ShipmentNoLike != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("shipmentNo")+" LIKE ?")
		values = append(values, "%"+strings.Replace(strings.Replace(*f.ShipmentNoLike, "?", "_", -1), "*", "%", -1)+"%")
	}

	if f.ShipmentNoPrefix != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("shipmentNo")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.ShipmentNoPrefix))
	}

	if f.ShipmentNoSuffix != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("shipmentNo")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.ShipmentNoSuffix))
	}

	if f.ShipmentNoNull != nil {
		if *f.ShipmentNoNull {
			conditions = append(conditions, aliasPrefix+SnakeString("shipmentNo")+" IS NULL"+" OR "+aliasPrefix+SnakeString("shipmentNo")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("shipmentNo")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("shipmentNo")+" <> ''")
		}
	}

	if f.TransportationMode != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("transportationMode")+" = ?")
		values = append(values, f.TransportationMode)
	}

	if f.TransportationModeNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("transportationMode")+" != ?")
		values = append(values, f.TransportationModeNe)
	}

	if f.TransportationModeGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("transportationMode")+" > ?")
		values = append(values, f.TransportationModeGt)
	}

	if f.TransportationModeLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("transportationMode")+" < ?")
		values = append(values, f.TransportationModeLt)
	}

	if f.TransportationModeGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("transportationMode")+" >= ?")
		values = append(values, f.TransportationModeGte)
	}

	if f.TransportationModeLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("transportationMode")+" <= ?")
		values = append(values, f.TransportationModeLte)
	}

	if f.TransportationModeIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("transportationMode")+" IN (?)")
		values = append(values, f.TransportationModeIn)
	}

	if f.TransportationModeLike != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("transportationMode")+" LIKE ?")
		values = append(values, "%"+strings.Replace(strings.Replace(*f.TransportationModeLike, "?", "_", -1), "*", "%", -1)+"%")
	}

	if f.TransportationModePrefix != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("transportationMode")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.TransportationModePrefix))
	}

	if f.TransportationModeSuffix != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("transportationMode")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.TransportationModeSuffix))
	}

	if f.TransportationModeNull != nil {
		if *f.TransportationModeNull {
			conditions = append(conditions, aliasPrefix+SnakeString("transportationMode")+" IS NULL"+" OR "+aliasPrefix+SnakeString("transportationMode")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("transportationMode")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("transportationMode")+" <> ''")
		}
	}

	if f.StartLocationID != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("startLocationId")+" = ?")
		values = append(values, f.StartLocationID)
	}

	if f.StartLocationIDNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("startLocationId")+" != ?")
		values = append(values, f.StartLocationIDNe)
	}

	if f.StartLocationIDGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("startLocationId")+" > ?")
		values = append(values, f.StartLocationIDGt)
	}

	if f.StartLocationIDLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("startLocationId")+" < ?")
		values = append(values, f.StartLocationIDLt)
	}

	if f.StartLocationIDGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("startLocationId")+" >= ?")
		values = append(values, f.StartLocationIDGte)
	}

	if f.StartLocationIDLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("startLocationId")+" <= ?")
		values = append(values, f.StartLocationIDLte)
	}

	if f.StartLocationIDIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("startLocationId")+" IN (?)")
		values = append(values, f.StartLocationIDIn)
	}

	if f.StartLocationIDNull != nil {
		if *f.StartLocationIDNull {
			conditions = append(conditions, aliasPrefix+SnakeString("startLocationId")+" IS NULL"+" OR "+aliasPrefix+SnakeString("startLocationId")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("startLocationId")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("startLocationId")+" <> ''")
		}
	}

	if f.EndLocationID != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("endLocationId")+" = ?")
		values = append(values, f.EndLocationID)
	}

	if f.EndLocationIDNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("endLocationId")+" != ?")
		values = append(values, f.EndLocationIDNe)
	}

	if f.EndLocationIDGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("endLocationId")+" > ?")
		values = append(values, f.EndLocationIDGt)
	}

	if f.EndLocationIDLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("endLocationId")+" < ?")
		values = append(values, f.EndLocationIDLt)
	}

	if f.EndLocationIDGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("endLocationId")+" >= ?")
		values = append(values, f.EndLocationIDGte)
	}

	if f.EndLocationIDLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("endLocationId")+" <= ?")
		values = append(values, f.EndLocationIDLte)
	}

	if f.EndLocationIDIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("endLocationId")+" IN (?)")
		values = append(values, f.EndLocationIDIn)
	}

	if f.EndLocationIDNull != nil {
		if *f.EndLocationIDNull {
			conditions = append(conditions, aliasPrefix+SnakeString("endLocationId")+" IS NULL"+" OR "+aliasPrefix+SnakeString("endLocationId")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("endLocationId")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("endLocationId")+" <> ''")
		}
	}

	if f.DeletedBy != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" = ?")
		values = append(values, f.DeletedBy)
	}

	if f.DeletedByNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" != ?")
		values = append(values, f.DeletedByNe)
	}

	if f.DeletedByGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" > ?")
		values = append(values, f.DeletedByGt)
	}

	if f.DeletedByLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" < ?")
		values = append(values, f.DeletedByLt)
	}

	if f.DeletedByGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" >= ?")
		values = append(values, f.DeletedByGte)
	}

	if f.DeletedByLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" <= ?")
		values = append(values, f.DeletedByLte)
	}

	if f.DeletedByIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" IN (?)")
		values = append(values, f.DeletedByIn)
	}

	if f.DeletedByNull != nil {
		if *f.DeletedByNull {
			conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" IS NULL"+" OR "+aliasPrefix+SnakeString("deletedBy")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("deletedBy")+" <> ''")
		}
	}

	if f.UpdatedBy != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" = ?")
		values = append(values, f.UpdatedBy)
	}

	if f.UpdatedByNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" != ?")
		values = append(values, f.UpdatedByNe)
	}

	if f.UpdatedByGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" > ?")
		values = append(values, f.UpdatedByGt)
	}

	if f.UpdatedByLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" < ?")
		values = append(values, f.UpdatedByLt)
	}

	if f.UpdatedByGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" >= ?")
		values = append(values, f.UpdatedByGte)
	}

	if f.UpdatedByLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" <= ?")
		values = append(values, f.UpdatedByLte)
	}

	if f.UpdatedByIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" IN (?)")
		values = append(values, f.UpdatedByIn)
	}

	if f.UpdatedByNull != nil {
		if *f.UpdatedByNull {
			conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" IS NULL"+" OR "+aliasPrefix+SnakeString("updatedBy")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("updatedBy")+" <> ''")
		}
	}

	if f.CreatedBy != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" = ?")
		values = append(values, f.CreatedBy)
	}

	if f.CreatedByNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" != ?")
		values = append(values, f.CreatedByNe)
	}

	if f.CreatedByGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" > ?")
		values = append(values, f.CreatedByGt)
	}

	if f.CreatedByLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" < ?")
		values = append(values, f.CreatedByLt)
	}

	if f.CreatedByGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" >= ?")
		values = append(values, f.CreatedByGte)
	}

	if f.CreatedByLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" <= ?")
		values = append(values, f.CreatedByLte)
	}

	if f.CreatedByIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" IN (?)")
		values = append(values, f.CreatedByIn)
	}

	if f.CreatedByNull != nil {
		if *f.CreatedByNull {
			conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" IS NULL"+" OR "+aliasPrefix+SnakeString("createdBy")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("createdBy")+" <> ''")
		}
	}

	if f.DeletedAt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" = ?")
		values = append(values, f.DeletedAt)
	}

	if f.DeletedAtNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" != ?")
		values = append(values, f.DeletedAtNe)
	}

	if f.DeletedAtGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" > ?")
		values = append(values, f.DeletedAtGt)
	}

	if f.DeletedAtLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" < ?")
		values = append(values, f.DeletedAtLt)
	}

	if f.DeletedAtGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" >= ?")
		values = append(values, f.DeletedAtGte)
	}

	if f.DeletedAtLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" <= ?")
		values = append(values, f.DeletedAtLte)
	}

	if f.DeletedAtIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" IN (?)")
		values = append(values, f.DeletedAtIn)
	}

	if f.DeletedAtNull != nil {
		if *f.DeletedAtNull {
			conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" IS NULL"+" OR "+aliasPrefix+SnakeString("deletedAt")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("deletedAt")+" <> ''")
		}
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" = ?")
		values = append(values, f.UpdatedAt)
	}

	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" != ?")
		values = append(values, f.UpdatedAtNe)
	}

	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" > ?")
		values = append(values, f.UpdatedAtGt)
	}

	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" < ?")
		values = append(values, f.UpdatedAtLt)
	}

	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" >= ?")
		values = append(values, f.UpdatedAtGte)
	}

	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" <= ?")
		values = append(values, f.UpdatedAtLte)
	}

	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.UpdatedAtNull != nil {
		if *f.UpdatedAtNull {
			conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" IS NULL"+" OR "+aliasPrefix+SnakeString("updatedAt")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("updatedAt")+" <> ''")
		}
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" = ?")
		values = append(values, f.CreatedAt)
	}

	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" != ?")
		values = append(values, f.CreatedAtNe)
	}

	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" > ?")
		values = append(values, f.CreatedAtGt)
	}

	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" < ?")
		values = append(values, f.CreatedAtLt)
	}

	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" >= ?")
		values = append(values, f.CreatedAtGte)
	}

	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" <= ?")
		values = append(values, f.CreatedAtLte)
	}

	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	if f.CreatedAtNull != nil {
		if *f.CreatedAtNull {
			conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" IS NULL"+" OR "+aliasPrefix+SnakeString("createdAt")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("createdAt")+" <> ''")
		}
	}

	return
}

// AndWith convenience method for combining two or more filters with AND statement
func (f *ShipmentFilterType) AndWith(f2 ...*ShipmentFilterType) *ShipmentFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &ShipmentFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *ShipmentFilterType) OrWith(f2 ...*ShipmentFilterType) *ShipmentFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &ShipmentFilterType{
		Or: append(_f2, f),
	}
}

func (f *CarrierFilterType) IsEmpty(ctx context.Context) bool {
	wheres := []string{}
	values := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, "companies", &wheres, &values, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0
}
func (f *CarrierFilterType) Apply(ctx context.Context, wheres *[]string, values *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, TableName("carriers"), wheres, values, joins)
}
func (f *CarrierFilterType) ApplyWithAlias(ctx context.Context, alias string, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := alias + "."

	_where, _values := f.WhereContent(aliasPrefix)
	*wheres = append(*wheres, _where...)
	*values = append(*values, _values...)

	if f.Or != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			_cs := []string{}
			err := or.ApplyWithAlias(ctx, alias, &_cs, &vs, &js)
			if err != nil {
				return err
			}
			cs = append(cs, strings.Join(_cs, " AND "))
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, "("+strings.Join(cs, " OR ")+")")
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, alias, &cs, &vs, &js)
			if err != nil {
				return err
			}
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, strings.Join(cs, " AND "))
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}

	return nil
}

func (f *CarrierFilterType) WhereContent(aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" = ?")
		values = append(values, f.ID)
	}

	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" != ?")
		values = append(values, f.IDNe)
	}

	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" > ?")
		values = append(values, f.IDGt)
	}

	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" < ?")
		values = append(values, f.IDLt)
	}

	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" >= ?")
		values = append(values, f.IDGte)
	}

	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" <= ?")
		values = append(values, f.IDLte)
	}

	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" IN (?)")
		values = append(values, f.IDIn)
	}

	if f.IDNull != nil {
		if *f.IDNull {
			conditions = append(conditions, aliasPrefix+SnakeString("id")+" IS NULL"+" OR "+aliasPrefix+SnakeString("id")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("id")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("id")+" <> ''")
		}
	}

	if f.CarrierName != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("carrierName")+" = ?")
		values = append(values, f.CarrierName)
	}

	if f.CarrierNameNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("carrierName")+" != ?")
		values = append(values, f.CarrierNameNe)
	}

	if f.CarrierNameGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("carrierName")+" > ?")
		values = append(values, f.CarrierNameGt)
	}

	if f.CarrierNameLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("carrierName")+" < ?")
		values = append(values, f.CarrierNameLt)
	}

	if f.CarrierNameGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("carrierName")+" >= ?")
		values = append(values, f.CarrierNameGte)
	}

	if f.CarrierNameLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("carrierName")+" <= ?")
		values = append(values, f.CarrierNameLte)
	}

	if f.CarrierNameIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("carrierName")+" IN (?)")
		values = append(values, f.CarrierNameIn)
	}

	if f.CarrierNameLike != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("carrierName")+" LIKE ?")
		values = append(values, "%"+strings.Replace(strings.Replace(*f.CarrierNameLike, "?", "_", -1), "*", "%", -1)+"%")
	}

	if f.CarrierNamePrefix != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("carrierName")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.CarrierNamePrefix))
	}

	if f.CarrierNameSuffix != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("carrierName")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.CarrierNameSuffix))
	}

	if f.CarrierNameNull != nil {
		if *f.CarrierNameNull {
			conditions = append(conditions, aliasPrefix+SnakeString("carrierName")+" IS NULL"+" OR "+aliasPrefix+SnakeString("carrierName")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("carrierName")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("carrierName")+" <> ''")
		}
	}

	if f.ContactPerson != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("contactPerson")+" = ?")
		values = append(values, f.ContactPerson)
	}

	if f.ContactPersonNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("contactPerson")+" != ?")
		values = append(values, f.ContactPersonNe)
	}

	if f.ContactPersonGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("contactPerson")+" > ?")
		values = append(values, f.ContactPersonGt)
	}

	if f.ContactPersonLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("contactPerson")+" < ?")
		values = append(values, f.ContactPersonLt)
	}

	if f.ContactPersonGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("contactPerson")+" >= ?")
		values = append(values, f.ContactPersonGte)
	}

	if f.ContactPersonLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("contactPerson")+" <= ?")
		values = append(values, f.ContactPersonLte)
	}

	if f.ContactPersonIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("contactPerson")+" IN (?)")
		values = append(values, f.ContactPersonIn)
	}

	if f.ContactPersonLike != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("contactPerson")+" LIKE ?")
		values = append(values, "%"+strings.Replace(strings.Replace(*f.ContactPersonLike, "?", "_", -1), "*", "%", -1)+"%")
	}

	if f.ContactPersonPrefix != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("contactPerson")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.ContactPersonPrefix))
	}

	if f.ContactPersonSuffix != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("contactPerson")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.ContactPersonSuffix))
	}

	if f.ContactPersonNull != nil {
		if *f.ContactPersonNull {
			conditions = append(conditions, aliasPrefix+SnakeString("contactPerson")+" IS NULL"+" OR "+aliasPrefix+SnakeString("contactPerson")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("contactPerson")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("contactPerson")+" <> ''")
		}
	}

	if f.ContactInfo != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("contactInfo")+" = ?")
		values = append(values, f.ContactInfo)
	}

	if f.ContactInfoNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("contactInfo")+" != ?")
		values = append(values, f.ContactInfoNe)
	}

	if f.ContactInfoGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("contactInfo")+" > ?")
		values = append(values, f.ContactInfoGt)
	}

	if f.ContactInfoLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("contactInfo")+" < ?")
		values = append(values, f.ContactInfoLt)
	}

	if f.ContactInfoGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("contactInfo")+" >= ?")
		values = append(values, f.ContactInfoGte)
	}

	if f.ContactInfoLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("contactInfo")+" <= ?")
		values = append(values, f.ContactInfoLte)
	}

	if f.ContactInfoIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("contactInfo")+" IN (?)")
		values = append(values, f.ContactInfoIn)
	}

	if f.ContactInfoLike != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("contactInfo")+" LIKE ?")
		values = append(values, "%"+strings.Replace(strings.Replace(*f.ContactInfoLike, "?", "_", -1), "*", "%", -1)+"%")
	}

	if f.ContactInfoPrefix != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("contactInfo")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.ContactInfoPrefix))
	}

	if f.ContactInfoSuffix != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("contactInfo")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.ContactInfoSuffix))
	}

	if f.ContactInfoNull != nil {
		if *f.ContactInfoNull {
			conditions = append(conditions, aliasPrefix+SnakeString("contactInfo")+" IS NULL"+" OR "+aliasPrefix+SnakeString("contactInfo")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("contactInfo")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("contactInfo")+" <> ''")
		}
	}

	if f.DeletedBy != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" = ?")
		values = append(values, f.DeletedBy)
	}

	if f.DeletedByNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" != ?")
		values = append(values, f.DeletedByNe)
	}

	if f.DeletedByGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" > ?")
		values = append(values, f.DeletedByGt)
	}

	if f.DeletedByLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" < ?")
		values = append(values, f.DeletedByLt)
	}

	if f.DeletedByGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" >= ?")
		values = append(values, f.DeletedByGte)
	}

	if f.DeletedByLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" <= ?")
		values = append(values, f.DeletedByLte)
	}

	if f.DeletedByIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" IN (?)")
		values = append(values, f.DeletedByIn)
	}

	if f.DeletedByNull != nil {
		if *f.DeletedByNull {
			conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" IS NULL"+" OR "+aliasPrefix+SnakeString("deletedBy")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("deletedBy")+" <> ''")
		}
	}

	if f.UpdatedBy != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" = ?")
		values = append(values, f.UpdatedBy)
	}

	if f.UpdatedByNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" != ?")
		values = append(values, f.UpdatedByNe)
	}

	if f.UpdatedByGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" > ?")
		values = append(values, f.UpdatedByGt)
	}

	if f.UpdatedByLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" < ?")
		values = append(values, f.UpdatedByLt)
	}

	if f.UpdatedByGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" >= ?")
		values = append(values, f.UpdatedByGte)
	}

	if f.UpdatedByLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" <= ?")
		values = append(values, f.UpdatedByLte)
	}

	if f.UpdatedByIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" IN (?)")
		values = append(values, f.UpdatedByIn)
	}

	if f.UpdatedByNull != nil {
		if *f.UpdatedByNull {
			conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" IS NULL"+" OR "+aliasPrefix+SnakeString("updatedBy")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("updatedBy")+" <> ''")
		}
	}

	if f.CreatedBy != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" = ?")
		values = append(values, f.CreatedBy)
	}

	if f.CreatedByNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" != ?")
		values = append(values, f.CreatedByNe)
	}

	if f.CreatedByGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" > ?")
		values = append(values, f.CreatedByGt)
	}

	if f.CreatedByLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" < ?")
		values = append(values, f.CreatedByLt)
	}

	if f.CreatedByGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" >= ?")
		values = append(values, f.CreatedByGte)
	}

	if f.CreatedByLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" <= ?")
		values = append(values, f.CreatedByLte)
	}

	if f.CreatedByIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" IN (?)")
		values = append(values, f.CreatedByIn)
	}

	if f.CreatedByNull != nil {
		if *f.CreatedByNull {
			conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" IS NULL"+" OR "+aliasPrefix+SnakeString("createdBy")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("createdBy")+" <> ''")
		}
	}

	if f.DeletedAt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" = ?")
		values = append(values, f.DeletedAt)
	}

	if f.DeletedAtNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" != ?")
		values = append(values, f.DeletedAtNe)
	}

	if f.DeletedAtGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" > ?")
		values = append(values, f.DeletedAtGt)
	}

	if f.DeletedAtLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" < ?")
		values = append(values, f.DeletedAtLt)
	}

	if f.DeletedAtGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" >= ?")
		values = append(values, f.DeletedAtGte)
	}

	if f.DeletedAtLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" <= ?")
		values = append(values, f.DeletedAtLte)
	}

	if f.DeletedAtIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" IN (?)")
		values = append(values, f.DeletedAtIn)
	}

	if f.DeletedAtNull != nil {
		if *f.DeletedAtNull {
			conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" IS NULL"+" OR "+aliasPrefix+SnakeString("deletedAt")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("deletedAt")+" <> ''")
		}
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" = ?")
		values = append(values, f.UpdatedAt)
	}

	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" != ?")
		values = append(values, f.UpdatedAtNe)
	}

	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" > ?")
		values = append(values, f.UpdatedAtGt)
	}

	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" < ?")
		values = append(values, f.UpdatedAtLt)
	}

	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" >= ?")
		values = append(values, f.UpdatedAtGte)
	}

	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" <= ?")
		values = append(values, f.UpdatedAtLte)
	}

	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.UpdatedAtNull != nil {
		if *f.UpdatedAtNull {
			conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" IS NULL"+" OR "+aliasPrefix+SnakeString("updatedAt")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("updatedAt")+" <> ''")
		}
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" = ?")
		values = append(values, f.CreatedAt)
	}

	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" != ?")
		values = append(values, f.CreatedAtNe)
	}

	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" > ?")
		values = append(values, f.CreatedAtGt)
	}

	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" < ?")
		values = append(values, f.CreatedAtLt)
	}

	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" >= ?")
		values = append(values, f.CreatedAtGte)
	}

	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" <= ?")
		values = append(values, f.CreatedAtLte)
	}

	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	if f.CreatedAtNull != nil {
		if *f.CreatedAtNull {
			conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" IS NULL"+" OR "+aliasPrefix+SnakeString("createdAt")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("createdAt")+" <> ''")
		}
	}

	return
}

// AndWith convenience method for combining two or more filters with AND statement
func (f *CarrierFilterType) AndWith(f2 ...*CarrierFilterType) *CarrierFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &CarrierFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *CarrierFilterType) OrWith(f2 ...*CarrierFilterType) *CarrierFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &CarrierFilterType{
		Or: append(_f2, f),
	}
}

func (f *LocationFilterType) IsEmpty(ctx context.Context) bool {
	wheres := []string{}
	values := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, "companies", &wheres, &values, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0
}
func (f *LocationFilterType) Apply(ctx context.Context, wheres *[]string, values *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, TableName("locations"), wheres, values, joins)
}
func (f *LocationFilterType) ApplyWithAlias(ctx context.Context, alias string, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := alias + "."

	_where, _values := f.WhereContent(aliasPrefix)
	*wheres = append(*wheres, _where...)
	*values = append(*values, _values...)

	if f.Or != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			_cs := []string{}
			err := or.ApplyWithAlias(ctx, alias, &_cs, &vs, &js)
			if err != nil {
				return err
			}
			cs = append(cs, strings.Join(_cs, " AND "))
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, "("+strings.Join(cs, " OR ")+")")
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, alias, &cs, &vs, &js)
			if err != nil {
				return err
			}
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, strings.Join(cs, " AND "))
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}

	if f.StartShipments != nil {
		_alias := alias + "_startShipments"
		*joins = append(*joins, "LEFT JOIN "+"shipments"+" "+_alias+" ON "+_alias+"."+"start_location_id"+" = "+alias+".id")
		err := f.StartShipments.ApplyWithAlias(ctx, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	if f.EndShipments != nil {
		_alias := alias + "_endShipments"
		*joins = append(*joins, "LEFT JOIN "+"shipments"+" "+_alias+" ON "+_alias+"."+"end_location_id"+" = "+alias+".id")
		err := f.EndShipments.ApplyWithAlias(ctx, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *LocationFilterType) WhereContent(aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" = ?")
		values = append(values, f.ID)
	}

	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" != ?")
		values = append(values, f.IDNe)
	}

	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" > ?")
		values = append(values, f.IDGt)
	}

	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" < ?")
		values = append(values, f.IDLt)
	}

	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" >= ?")
		values = append(values, f.IDGte)
	}

	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" <= ?")
		values = append(values, f.IDLte)
	}

	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" IN (?)")
		values = append(values, f.IDIn)
	}

	if f.IDNull != nil {
		if *f.IDNull {
			conditions = append(conditions, aliasPrefix+SnakeString("id")+" IS NULL"+" OR "+aliasPrefix+SnakeString("id")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("id")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("id")+" <> ''")
		}
	}

	if f.WarehouseAddress != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("warehouseAddress")+" = ?")
		values = append(values, f.WarehouseAddress)
	}

	if f.WarehouseAddressNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("warehouseAddress")+" != ?")
		values = append(values, f.WarehouseAddressNe)
	}

	if f.WarehouseAddressGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("warehouseAddress")+" > ?")
		values = append(values, f.WarehouseAddressGt)
	}

	if f.WarehouseAddressLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("warehouseAddress")+" < ?")
		values = append(values, f.WarehouseAddressLt)
	}

	if f.WarehouseAddressGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("warehouseAddress")+" >= ?")
		values = append(values, f.WarehouseAddressGte)
	}

	if f.WarehouseAddressLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("warehouseAddress")+" <= ?")
		values = append(values, f.WarehouseAddressLte)
	}

	if f.WarehouseAddressIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("warehouseAddress")+" IN (?)")
		values = append(values, f.WarehouseAddressIn)
	}

	if f.WarehouseAddressLike != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("warehouseAddress")+" LIKE ?")
		values = append(values, "%"+strings.Replace(strings.Replace(*f.WarehouseAddressLike, "?", "_", -1), "*", "%", -1)+"%")
	}

	if f.WarehouseAddressPrefix != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("warehouseAddress")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.WarehouseAddressPrefix))
	}

	if f.WarehouseAddressSuffix != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("warehouseAddress")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.WarehouseAddressSuffix))
	}

	if f.WarehouseAddressNull != nil {
		if *f.WarehouseAddressNull {
			conditions = append(conditions, aliasPrefix+SnakeString("warehouseAddress")+" IS NULL"+" OR "+aliasPrefix+SnakeString("warehouseAddress")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("warehouseAddress")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("warehouseAddress")+" <> ''")
		}
	}

	if f.LoadingAddress != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("loadingAddress")+" = ?")
		values = append(values, f.LoadingAddress)
	}

	if f.LoadingAddressNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("loadingAddress")+" != ?")
		values = append(values, f.LoadingAddressNe)
	}

	if f.LoadingAddressGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("loadingAddress")+" > ?")
		values = append(values, f.LoadingAddressGt)
	}

	if f.LoadingAddressLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("loadingAddress")+" < ?")
		values = append(values, f.LoadingAddressLt)
	}

	if f.LoadingAddressGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("loadingAddress")+" >= ?")
		values = append(values, f.LoadingAddressGte)
	}

	if f.LoadingAddressLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("loadingAddress")+" <= ?")
		values = append(values, f.LoadingAddressLte)
	}

	if f.LoadingAddressIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("loadingAddress")+" IN (?)")
		values = append(values, f.LoadingAddressIn)
	}

	if f.LoadingAddressLike != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("loadingAddress")+" LIKE ?")
		values = append(values, "%"+strings.Replace(strings.Replace(*f.LoadingAddressLike, "?", "_", -1), "*", "%", -1)+"%")
	}

	if f.LoadingAddressPrefix != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("loadingAddress")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.LoadingAddressPrefix))
	}

	if f.LoadingAddressSuffix != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("loadingAddress")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.LoadingAddressSuffix))
	}

	if f.LoadingAddressNull != nil {
		if *f.LoadingAddressNull {
			conditions = append(conditions, aliasPrefix+SnakeString("loadingAddress")+" IS NULL"+" OR "+aliasPrefix+SnakeString("loadingAddress")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("loadingAddress")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("loadingAddress")+" <> ''")
		}
	}

	if f.DeletedBy != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" = ?")
		values = append(values, f.DeletedBy)
	}

	if f.DeletedByNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" != ?")
		values = append(values, f.DeletedByNe)
	}

	if f.DeletedByGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" > ?")
		values = append(values, f.DeletedByGt)
	}

	if f.DeletedByLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" < ?")
		values = append(values, f.DeletedByLt)
	}

	if f.DeletedByGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" >= ?")
		values = append(values, f.DeletedByGte)
	}

	if f.DeletedByLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" <= ?")
		values = append(values, f.DeletedByLte)
	}

	if f.DeletedByIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" IN (?)")
		values = append(values, f.DeletedByIn)
	}

	if f.DeletedByNull != nil {
		if *f.DeletedByNull {
			conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" IS NULL"+" OR "+aliasPrefix+SnakeString("deletedBy")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("deletedBy")+" <> ''")
		}
	}

	if f.UpdatedBy != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" = ?")
		values = append(values, f.UpdatedBy)
	}

	if f.UpdatedByNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" != ?")
		values = append(values, f.UpdatedByNe)
	}

	if f.UpdatedByGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" > ?")
		values = append(values, f.UpdatedByGt)
	}

	if f.UpdatedByLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" < ?")
		values = append(values, f.UpdatedByLt)
	}

	if f.UpdatedByGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" >= ?")
		values = append(values, f.UpdatedByGte)
	}

	if f.UpdatedByLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" <= ?")
		values = append(values, f.UpdatedByLte)
	}

	if f.UpdatedByIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" IN (?)")
		values = append(values, f.UpdatedByIn)
	}

	if f.UpdatedByNull != nil {
		if *f.UpdatedByNull {
			conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" IS NULL"+" OR "+aliasPrefix+SnakeString("updatedBy")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("updatedBy")+" <> ''")
		}
	}

	if f.CreatedBy != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" = ?")
		values = append(values, f.CreatedBy)
	}

	if f.CreatedByNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" != ?")
		values = append(values, f.CreatedByNe)
	}

	if f.CreatedByGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" > ?")
		values = append(values, f.CreatedByGt)
	}

	if f.CreatedByLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" < ?")
		values = append(values, f.CreatedByLt)
	}

	if f.CreatedByGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" >= ?")
		values = append(values, f.CreatedByGte)
	}

	if f.CreatedByLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" <= ?")
		values = append(values, f.CreatedByLte)
	}

	if f.CreatedByIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" IN (?)")
		values = append(values, f.CreatedByIn)
	}

	if f.CreatedByNull != nil {
		if *f.CreatedByNull {
			conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" IS NULL"+" OR "+aliasPrefix+SnakeString("createdBy")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("createdBy")+" <> ''")
		}
	}

	if f.DeletedAt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" = ?")
		values = append(values, f.DeletedAt)
	}

	if f.DeletedAtNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" != ?")
		values = append(values, f.DeletedAtNe)
	}

	if f.DeletedAtGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" > ?")
		values = append(values, f.DeletedAtGt)
	}

	if f.DeletedAtLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" < ?")
		values = append(values, f.DeletedAtLt)
	}

	if f.DeletedAtGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" >= ?")
		values = append(values, f.DeletedAtGte)
	}

	if f.DeletedAtLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" <= ?")
		values = append(values, f.DeletedAtLte)
	}

	if f.DeletedAtIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" IN (?)")
		values = append(values, f.DeletedAtIn)
	}

	if f.DeletedAtNull != nil {
		if *f.DeletedAtNull {
			conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" IS NULL"+" OR "+aliasPrefix+SnakeString("deletedAt")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("deletedAt")+" <> ''")
		}
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" = ?")
		values = append(values, f.UpdatedAt)
	}

	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" != ?")
		values = append(values, f.UpdatedAtNe)
	}

	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" > ?")
		values = append(values, f.UpdatedAtGt)
	}

	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" < ?")
		values = append(values, f.UpdatedAtLt)
	}

	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" >= ?")
		values = append(values, f.UpdatedAtGte)
	}

	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" <= ?")
		values = append(values, f.UpdatedAtLte)
	}

	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.UpdatedAtNull != nil {
		if *f.UpdatedAtNull {
			conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" IS NULL"+" OR "+aliasPrefix+SnakeString("updatedAt")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("updatedAt")+" <> ''")
		}
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" = ?")
		values = append(values, f.CreatedAt)
	}

	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" != ?")
		values = append(values, f.CreatedAtNe)
	}

	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" > ?")
		values = append(values, f.CreatedAtGt)
	}

	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" < ?")
		values = append(values, f.CreatedAtLt)
	}

	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" >= ?")
		values = append(values, f.CreatedAtGte)
	}

	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" <= ?")
		values = append(values, f.CreatedAtLte)
	}

	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	if f.CreatedAtNull != nil {
		if *f.CreatedAtNull {
			conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" IS NULL"+" OR "+aliasPrefix+SnakeString("createdAt")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("createdAt")+" <> ''")
		}
	}

	return
}

// AndWith convenience method for combining two or more filters with AND statement
func (f *LocationFilterType) AndWith(f2 ...*LocationFilterType) *LocationFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &LocationFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *LocationFilterType) OrWith(f2 ...*LocationFilterType) *LocationFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &LocationFilterType{
		Or: append(_f2, f),
	}
}

func (f *EquipmentdFilterType) IsEmpty(ctx context.Context) bool {
	wheres := []string{}
	values := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, "companies", &wheres, &values, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0
}
func (f *EquipmentdFilterType) Apply(ctx context.Context, wheres *[]string, values *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, TableName("equipmentds"), wheres, values, joins)
}
func (f *EquipmentdFilterType) ApplyWithAlias(ctx context.Context, alias string, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := alias + "."

	_where, _values := f.WhereContent(aliasPrefix)
	*wheres = append(*wheres, _where...)
	*values = append(*values, _values...)

	if f.Or != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			_cs := []string{}
			err := or.ApplyWithAlias(ctx, alias, &_cs, &vs, &js)
			if err != nil {
				return err
			}
			cs = append(cs, strings.Join(_cs, " AND "))
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, "("+strings.Join(cs, " OR ")+")")
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, alias, &cs, &vs, &js)
			if err != nil {
				return err
			}
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, strings.Join(cs, " AND "))
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}

	return nil
}

func (f *EquipmentdFilterType) WhereContent(aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" = ?")
		values = append(values, f.ID)
	}

	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" != ?")
		values = append(values, f.IDNe)
	}

	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" > ?")
		values = append(values, f.IDGt)
	}

	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" < ?")
		values = append(values, f.IDLt)
	}

	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" >= ?")
		values = append(values, f.IDGte)
	}

	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" <= ?")
		values = append(values, f.IDLte)
	}

	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" IN (?)")
		values = append(values, f.IDIn)
	}

	if f.IDNull != nil {
		if *f.IDNull {
			conditions = append(conditions, aliasPrefix+SnakeString("id")+" IS NULL"+" OR "+aliasPrefix+SnakeString("id")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("id")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("id")+" <> ''")
		}
	}

	if f.VehicleType != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("vehicleType")+" = ?")
		values = append(values, f.VehicleType)
	}

	if f.VehicleTypeNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("vehicleType")+" != ?")
		values = append(values, f.VehicleTypeNe)
	}

	if f.VehicleTypeGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("vehicleType")+" > ?")
		values = append(values, f.VehicleTypeGt)
	}

	if f.VehicleTypeLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("vehicleType")+" < ?")
		values = append(values, f.VehicleTypeLt)
	}

	if f.VehicleTypeGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("vehicleType")+" >= ?")
		values = append(values, f.VehicleTypeGte)
	}

	if f.VehicleTypeLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("vehicleType")+" <= ?")
		values = append(values, f.VehicleTypeLte)
	}

	if f.VehicleTypeIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("vehicleType")+" IN (?)")
		values = append(values, f.VehicleTypeIn)
	}

	if f.VehicleTypeLike != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("vehicleType")+" LIKE ?")
		values = append(values, "%"+strings.Replace(strings.Replace(*f.VehicleTypeLike, "?", "_", -1), "*", "%", -1)+"%")
	}

	if f.VehicleTypePrefix != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("vehicleType")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.VehicleTypePrefix))
	}

	if f.VehicleTypeSuffix != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("vehicleType")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.VehicleTypeSuffix))
	}

	if f.VehicleTypeNull != nil {
		if *f.VehicleTypeNull {
			conditions = append(conditions, aliasPrefix+SnakeString("vehicleType")+" IS NULL"+" OR "+aliasPrefix+SnakeString("vehicleType")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("vehicleType")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("vehicleType")+" <> ''")
		}
	}

	if f.Capacity != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("capacity")+" = ?")
		values = append(values, f.Capacity)
	}

	if f.CapacityNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("capacity")+" != ?")
		values = append(values, f.CapacityNe)
	}

	if f.CapacityGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("capacity")+" > ?")
		values = append(values, f.CapacityGt)
	}

	if f.CapacityLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("capacity")+" < ?")
		values = append(values, f.CapacityLt)
	}

	if f.CapacityGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("capacity")+" >= ?")
		values = append(values, f.CapacityGte)
	}

	if f.CapacityLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("capacity")+" <= ?")
		values = append(values, f.CapacityLte)
	}

	if f.CapacityIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("capacity")+" IN (?)")
		values = append(values, f.CapacityIn)
	}

	if f.CapacityNull != nil {
		if *f.CapacityNull {
			conditions = append(conditions, aliasPrefix+SnakeString("capacity")+" IS NULL"+" OR "+aliasPrefix+SnakeString("capacity")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("capacity")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("capacity")+" <> ''")
		}
	}

	if f.DeletedBy != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" = ?")
		values = append(values, f.DeletedBy)
	}

	if f.DeletedByNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" != ?")
		values = append(values, f.DeletedByNe)
	}

	if f.DeletedByGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" > ?")
		values = append(values, f.DeletedByGt)
	}

	if f.DeletedByLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" < ?")
		values = append(values, f.DeletedByLt)
	}

	if f.DeletedByGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" >= ?")
		values = append(values, f.DeletedByGte)
	}

	if f.DeletedByLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" <= ?")
		values = append(values, f.DeletedByLte)
	}

	if f.DeletedByIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" IN (?)")
		values = append(values, f.DeletedByIn)
	}

	if f.DeletedByNull != nil {
		if *f.DeletedByNull {
			conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" IS NULL"+" OR "+aliasPrefix+SnakeString("deletedBy")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("deletedBy")+" <> ''")
		}
	}

	if f.UpdatedBy != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" = ?")
		values = append(values, f.UpdatedBy)
	}

	if f.UpdatedByNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" != ?")
		values = append(values, f.UpdatedByNe)
	}

	if f.UpdatedByGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" > ?")
		values = append(values, f.UpdatedByGt)
	}

	if f.UpdatedByLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" < ?")
		values = append(values, f.UpdatedByLt)
	}

	if f.UpdatedByGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" >= ?")
		values = append(values, f.UpdatedByGte)
	}

	if f.UpdatedByLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" <= ?")
		values = append(values, f.UpdatedByLte)
	}

	if f.UpdatedByIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" IN (?)")
		values = append(values, f.UpdatedByIn)
	}

	if f.UpdatedByNull != nil {
		if *f.UpdatedByNull {
			conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" IS NULL"+" OR "+aliasPrefix+SnakeString("updatedBy")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("updatedBy")+" <> ''")
		}
	}

	if f.CreatedBy != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" = ?")
		values = append(values, f.CreatedBy)
	}

	if f.CreatedByNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" != ?")
		values = append(values, f.CreatedByNe)
	}

	if f.CreatedByGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" > ?")
		values = append(values, f.CreatedByGt)
	}

	if f.CreatedByLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" < ?")
		values = append(values, f.CreatedByLt)
	}

	if f.CreatedByGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" >= ?")
		values = append(values, f.CreatedByGte)
	}

	if f.CreatedByLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" <= ?")
		values = append(values, f.CreatedByLte)
	}

	if f.CreatedByIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" IN (?)")
		values = append(values, f.CreatedByIn)
	}

	if f.CreatedByNull != nil {
		if *f.CreatedByNull {
			conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" IS NULL"+" OR "+aliasPrefix+SnakeString("createdBy")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("createdBy")+" <> ''")
		}
	}

	if f.DeletedAt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" = ?")
		values = append(values, f.DeletedAt)
	}

	if f.DeletedAtNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" != ?")
		values = append(values, f.DeletedAtNe)
	}

	if f.DeletedAtGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" > ?")
		values = append(values, f.DeletedAtGt)
	}

	if f.DeletedAtLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" < ?")
		values = append(values, f.DeletedAtLt)
	}

	if f.DeletedAtGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" >= ?")
		values = append(values, f.DeletedAtGte)
	}

	if f.DeletedAtLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" <= ?")
		values = append(values, f.DeletedAtLte)
	}

	if f.DeletedAtIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" IN (?)")
		values = append(values, f.DeletedAtIn)
	}

	if f.DeletedAtNull != nil {
		if *f.DeletedAtNull {
			conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" IS NULL"+" OR "+aliasPrefix+SnakeString("deletedAt")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("deletedAt")+" <> ''")
		}
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" = ?")
		values = append(values, f.UpdatedAt)
	}

	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" != ?")
		values = append(values, f.UpdatedAtNe)
	}

	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" > ?")
		values = append(values, f.UpdatedAtGt)
	}

	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" < ?")
		values = append(values, f.UpdatedAtLt)
	}

	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" >= ?")
		values = append(values, f.UpdatedAtGte)
	}

	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" <= ?")
		values = append(values, f.UpdatedAtLte)
	}

	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.UpdatedAtNull != nil {
		if *f.UpdatedAtNull {
			conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" IS NULL"+" OR "+aliasPrefix+SnakeString("updatedAt")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("updatedAt")+" <> ''")
		}
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" = ?")
		values = append(values, f.CreatedAt)
	}

	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" != ?")
		values = append(values, f.CreatedAtNe)
	}

	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" > ?")
		values = append(values, f.CreatedAtGt)
	}

	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" < ?")
		values = append(values, f.CreatedAtLt)
	}

	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" >= ?")
		values = append(values, f.CreatedAtGte)
	}

	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" <= ?")
		values = append(values, f.CreatedAtLte)
	}

	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	if f.CreatedAtNull != nil {
		if *f.CreatedAtNull {
			conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" IS NULL"+" OR "+aliasPrefix+SnakeString("createdAt")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("createdAt")+" <> ''")
		}
	}

	return
}

// AndWith convenience method for combining two or more filters with AND statement
func (f *EquipmentdFilterType) AndWith(f2 ...*EquipmentdFilterType) *EquipmentdFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &EquipmentdFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *EquipmentdFilterType) OrWith(f2 ...*EquipmentdFilterType) *EquipmentdFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &EquipmentdFilterType{
		Or: append(_f2, f),
	}
}
