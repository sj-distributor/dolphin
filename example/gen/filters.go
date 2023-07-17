package gen

import (
	"context"
	"fmt"
	"strings"
)

func (f *UserFilterType) IsEmpty(ctx context.Context) bool {
	wheres := []string{}
	values := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, "companies", &wheres, &values, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0
}
func (f *UserFilterType) Apply(ctx context.Context, wheres *[]string, values *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, TableName("users"), wheres, values, joins)
}
func (f *UserFilterType) ApplyWithAlias(ctx context.Context, alias string, wheres *[]string, values *[]interface{}, joins *[]string) error {
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

	if f.T != nil {
		_alias := alias + "_t"
		*joins = append(*joins, "LEFT JOIN "+"tasks"+" "+_alias+" ON "+_alias+".id = "+alias+"."+"t_id")
		err := f.T.ApplyWithAlias(ctx, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	if f.Tt != nil {
		_alias := alias + "_tt"
		*joins = append(*joins, "LEFT JOIN "+"tasks"+" "+_alias+" ON "+_alias+".id = "+alias+"."+"tt_id")
		err := f.Tt.ApplyWithAlias(ctx, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	if f.Ttt != nil {
		_alias := alias + "_ttt"
		*joins = append(*joins, "LEFT JOIN "+"tasks"+" "+_alias+" ON "+_alias+"."+"uuu_id"+" = "+alias+".id")
		err := f.Ttt.ApplyWithAlias(ctx, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	if f.Tttt != nil {
		_alias := alias + "_tttt"
		*joins = append(*joins, "LEFT JOIN "+"task_uuuu"+" "+_alias+"_jointable"+" ON "+alias+".id = "+_alias+"_jointable"+"."+"uuuu_id"+" LEFT JOIN "+TableName("tasks")+" "+_alias+" ON "+_alias+"_jointable"+"."+"tttt_id"+" = "+_alias+".id")
		err := f.Tttt.ApplyWithAlias(ctx, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *UserFilterType) WhereContent(aliasPrefix string) (conditions []string, values []interface{}) {
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

	if f.Phone != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("phone")+" = ?")
		values = append(values, f.Phone)
	}

	if f.PhoneNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("phone")+" != ?")
		values = append(values, f.PhoneNe)
	}

	if f.PhoneGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("phone")+" > ?")
		values = append(values, f.PhoneGt)
	}

	if f.PhoneLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("phone")+" < ?")
		values = append(values, f.PhoneLt)
	}

	if f.PhoneGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("phone")+" >= ?")
		values = append(values, f.PhoneGte)
	}

	if f.PhoneLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("phone")+" <= ?")
		values = append(values, f.PhoneLte)
	}

	if f.PhoneIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("phone")+" IN (?)")
		values = append(values, f.PhoneIn)
	}

	if f.PhoneLike != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("phone")+" LIKE ?")
		values = append(values, "%"+strings.Replace(strings.Replace(*f.PhoneLike, "?", "_", -1), "*", "%", -1)+"%")
	}

	if f.PhonePrefix != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("phone")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.PhonePrefix))
	}

	if f.PhoneSuffix != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("phone")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.PhoneSuffix))
	}

	if f.PhoneNull != nil {
		if *f.PhoneNull {
			conditions = append(conditions, aliasPrefix+SnakeString("phone")+" IS NULL"+" OR "+aliasPrefix+SnakeString("phone")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("phone")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("phone")+" <> ''")
		}
	}

	if f.TID != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("tId")+" = ?")
		values = append(values, f.TID)
	}

	if f.TIDNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("tId")+" != ?")
		values = append(values, f.TIDNe)
	}

	if f.TIDGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("tId")+" > ?")
		values = append(values, f.TIDGt)
	}

	if f.TIDLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("tId")+" < ?")
		values = append(values, f.TIDLt)
	}

	if f.TIDGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("tId")+" >= ?")
		values = append(values, f.TIDGte)
	}

	if f.TIDLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("tId")+" <= ?")
		values = append(values, f.TIDLte)
	}

	if f.TIDIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("tId")+" IN (?)")
		values = append(values, f.TIDIn)
	}

	if f.TIDNull != nil {
		if *f.TIDNull {
			conditions = append(conditions, aliasPrefix+SnakeString("tId")+" IS NULL"+" OR "+aliasPrefix+SnakeString("tId")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("tId")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("tId")+" <> ''")
		}
	}

	if f.TtID != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("ttId")+" = ?")
		values = append(values, f.TtID)
	}

	if f.TtIDNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("ttId")+" != ?")
		values = append(values, f.TtIDNe)
	}

	if f.TtIDGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("ttId")+" > ?")
		values = append(values, f.TtIDGt)
	}

	if f.TtIDLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("ttId")+" < ?")
		values = append(values, f.TtIDLt)
	}

	if f.TtIDGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("ttId")+" >= ?")
		values = append(values, f.TtIDGte)
	}

	if f.TtIDLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("ttId")+" <= ?")
		values = append(values, f.TtIDLte)
	}

	if f.TtIDIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("ttId")+" IN (?)")
		values = append(values, f.TtIDIn)
	}

	if f.TtIDNull != nil {
		if *f.TtIDNull {
			conditions = append(conditions, aliasPrefix+SnakeString("ttId")+" IS NULL"+" OR "+aliasPrefix+SnakeString("ttId")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("ttId")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("ttId")+" <> ''")
		}
	}

	if f.IsDelete != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("isDelete")+" = ?")
		values = append(values, f.IsDelete)
	}

	if f.IsDeleteNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("isDelete")+" != ?")
		values = append(values, f.IsDeleteNe)
	}

	if f.IsDeleteGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("isDelete")+" > ?")
		values = append(values, f.IsDeleteGt)
	}

	if f.IsDeleteLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("isDelete")+" < ?")
		values = append(values, f.IsDeleteLt)
	}

	if f.IsDeleteGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("isDelete")+" >= ?")
		values = append(values, f.IsDeleteGte)
	}

	if f.IsDeleteLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("isDelete")+" <= ?")
		values = append(values, f.IsDeleteLte)
	}

	if f.IsDeleteIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("isDelete")+" IN (?)")
		values = append(values, f.IsDeleteIn)
	}

	if f.IsDeleteNull != nil {
		if *f.IsDeleteNull {
			conditions = append(conditions, aliasPrefix+SnakeString("isDelete")+" IS NULL"+" OR "+aliasPrefix+SnakeString("isDelete")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("isDelete")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("isDelete")+" <> ''")
		}
	}

	if f.Weight != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("weight")+" = ?")
		values = append(values, f.Weight)
	}

	if f.WeightNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("weight")+" != ?")
		values = append(values, f.WeightNe)
	}

	if f.WeightGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("weight")+" > ?")
		values = append(values, f.WeightGt)
	}

	if f.WeightLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("weight")+" < ?")
		values = append(values, f.WeightLt)
	}

	if f.WeightGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("weight")+" >= ?")
		values = append(values, f.WeightGte)
	}

	if f.WeightLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("weight")+" <= ?")
		values = append(values, f.WeightLte)
	}

	if f.WeightIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("weight")+" IN (?)")
		values = append(values, f.WeightIn)
	}

	if f.WeightNull != nil {
		if *f.WeightNull {
			conditions = append(conditions, aliasPrefix+SnakeString("weight")+" IS NULL"+" OR "+aliasPrefix+SnakeString("weight")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("weight")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("weight")+" <> ''")
		}
	}

	if f.State != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("state")+" = ?")
		values = append(values, f.State)
	}

	if f.StateNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("state")+" != ?")
		values = append(values, f.StateNe)
	}

	if f.StateGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("state")+" > ?")
		values = append(values, f.StateGt)
	}

	if f.StateLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("state")+" < ?")
		values = append(values, f.StateLt)
	}

	if f.StateGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("state")+" >= ?")
		values = append(values, f.StateGte)
	}

	if f.StateLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("state")+" <= ?")
		values = append(values, f.StateLte)
	}

	if f.StateIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("state")+" IN (?)")
		values = append(values, f.StateIn)
	}

	if f.StateNull != nil {
		if *f.StateNull {
			conditions = append(conditions, aliasPrefix+SnakeString("state")+" IS NULL"+" OR "+aliasPrefix+SnakeString("state")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("state")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("state")+" <> ''")
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
func (f *UserFilterType) AndWith(f2 ...*UserFilterType) *UserFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &UserFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *UserFilterType) OrWith(f2 ...*UserFilterType) *UserFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &UserFilterType{
		Or: append(_f2, f),
	}
}

func (f *TaskFilterType) IsEmpty(ctx context.Context) bool {
	wheres := []string{}
	values := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, "companies", &wheres, &values, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0
}
func (f *TaskFilterType) Apply(ctx context.Context, wheres *[]string, values *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, TableName("tasks"), wheres, values, joins)
}
func (f *TaskFilterType) ApplyWithAlias(ctx context.Context, alias string, wheres *[]string, values *[]interface{}, joins *[]string) error {
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

	if f.U != nil {
		_alias := alias + "_u"
		*joins = append(*joins, "LEFT JOIN "+"users"+" "+_alias+" ON "+_alias+".id = "+alias+"."+"u_id")
		err := f.U.ApplyWithAlias(ctx, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	if f.Uu != nil {
		_alias := alias + "_uu"
		*joins = append(*joins, "LEFT JOIN "+"users"+" "+_alias+" ON "+_alias+"."+"tt_id"+" = "+alias+".id")
		err := f.Uu.ApplyWithAlias(ctx, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	if f.Uuu != nil {
		_alias := alias + "_uuu"
		*joins = append(*joins, "LEFT JOIN "+"users"+" "+_alias+" ON "+_alias+".id = "+alias+"."+"uuu_id")
		err := f.Uuu.ApplyWithAlias(ctx, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	if f.Uuuu != nil {
		_alias := alias + "_uuuu"
		*joins = append(*joins, "LEFT JOIN "+"task_uuuu"+" "+_alias+"_jointable"+" ON "+alias+".id = "+_alias+"_jointable"+"."+"tttt_id"+" LEFT JOIN "+TableName("users")+" "+_alias+" ON "+_alias+"_jointable"+"."+"uuuu_id"+" = "+_alias+".id")
		err := f.Uuuu.ApplyWithAlias(ctx, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *TaskFilterType) WhereContent(aliasPrefix string) (conditions []string, values []interface{}) {
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

	if f.Title != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("title")+" = ?")
		values = append(values, f.Title)
	}

	if f.TitleNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("title")+" != ?")
		values = append(values, f.TitleNe)
	}

	if f.TitleGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("title")+" > ?")
		values = append(values, f.TitleGt)
	}

	if f.TitleLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("title")+" < ?")
		values = append(values, f.TitleLt)
	}

	if f.TitleGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("title")+" >= ?")
		values = append(values, f.TitleGte)
	}

	if f.TitleLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("title")+" <= ?")
		values = append(values, f.TitleLte)
	}

	if f.TitleIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("title")+" IN (?)")
		values = append(values, f.TitleIn)
	}

	if f.TitleLike != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("title")+" LIKE ?")
		values = append(values, "%"+strings.Replace(strings.Replace(*f.TitleLike, "?", "_", -1), "*", "%", -1)+"%")
	}

	if f.TitlePrefix != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("title")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.TitlePrefix))
	}

	if f.TitleSuffix != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("title")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.TitleSuffix))
	}

	if f.TitleNull != nil {
		if *f.TitleNull {
			conditions = append(conditions, aliasPrefix+SnakeString("title")+" IS NULL"+" OR "+aliasPrefix+SnakeString("title")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("title")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("title")+" <> ''")
		}
	}

	if f.UID != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("uId")+" = ?")
		values = append(values, f.UID)
	}

	if f.UIDNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("uId")+" != ?")
		values = append(values, f.UIDNe)
	}

	if f.UIDGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("uId")+" > ?")
		values = append(values, f.UIDGt)
	}

	if f.UIDLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("uId")+" < ?")
		values = append(values, f.UIDLt)
	}

	if f.UIDGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("uId")+" >= ?")
		values = append(values, f.UIDGte)
	}

	if f.UIDLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("uId")+" <= ?")
		values = append(values, f.UIDLte)
	}

	if f.UIDIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("uId")+" IN (?)")
		values = append(values, f.UIDIn)
	}

	if f.UIDNull != nil {
		if *f.UIDNull {
			conditions = append(conditions, aliasPrefix+SnakeString("uId")+" IS NULL"+" OR "+aliasPrefix+SnakeString("uId")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("uId")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("uId")+" <> ''")
		}
	}

	if f.UuuID != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("uuuId")+" = ?")
		values = append(values, f.UuuID)
	}

	if f.UuuIDNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("uuuId")+" != ?")
		values = append(values, f.UuuIDNe)
	}

	if f.UuuIDGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("uuuId")+" > ?")
		values = append(values, f.UuuIDGt)
	}

	if f.UuuIDLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("uuuId")+" < ?")
		values = append(values, f.UuuIDLt)
	}

	if f.UuuIDGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("uuuId")+" >= ?")
		values = append(values, f.UuuIDGte)
	}

	if f.UuuIDLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("uuuId")+" <= ?")
		values = append(values, f.UuuIDLte)
	}

	if f.UuuIDIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("uuuId")+" IN (?)")
		values = append(values, f.UuuIDIn)
	}

	if f.UuuIDNull != nil {
		if *f.UuuIDNull {
			conditions = append(conditions, aliasPrefix+SnakeString("uuuId")+" IS NULL"+" OR "+aliasPrefix+SnakeString("uuuId")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("uuuId")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("uuuId")+" <> ''")
		}
	}

	if f.IsDelete != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("isDelete")+" = ?")
		values = append(values, f.IsDelete)
	}

	if f.IsDeleteNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("isDelete")+" != ?")
		values = append(values, f.IsDeleteNe)
	}

	if f.IsDeleteGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("isDelete")+" > ?")
		values = append(values, f.IsDeleteGt)
	}

	if f.IsDeleteLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("isDelete")+" < ?")
		values = append(values, f.IsDeleteLt)
	}

	if f.IsDeleteGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("isDelete")+" >= ?")
		values = append(values, f.IsDeleteGte)
	}

	if f.IsDeleteLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("isDelete")+" <= ?")
		values = append(values, f.IsDeleteLte)
	}

	if f.IsDeleteIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("isDelete")+" IN (?)")
		values = append(values, f.IsDeleteIn)
	}

	if f.IsDeleteNull != nil {
		if *f.IsDeleteNull {
			conditions = append(conditions, aliasPrefix+SnakeString("isDelete")+" IS NULL"+" OR "+aliasPrefix+SnakeString("isDelete")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("isDelete")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("isDelete")+" <> ''")
		}
	}

	if f.Weight != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("weight")+" = ?")
		values = append(values, f.Weight)
	}

	if f.WeightNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("weight")+" != ?")
		values = append(values, f.WeightNe)
	}

	if f.WeightGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("weight")+" > ?")
		values = append(values, f.WeightGt)
	}

	if f.WeightLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("weight")+" < ?")
		values = append(values, f.WeightLt)
	}

	if f.WeightGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("weight")+" >= ?")
		values = append(values, f.WeightGte)
	}

	if f.WeightLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("weight")+" <= ?")
		values = append(values, f.WeightLte)
	}

	if f.WeightIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("weight")+" IN (?)")
		values = append(values, f.WeightIn)
	}

	if f.WeightNull != nil {
		if *f.WeightNull {
			conditions = append(conditions, aliasPrefix+SnakeString("weight")+" IS NULL"+" OR "+aliasPrefix+SnakeString("weight")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("weight")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("weight")+" <> ''")
		}
	}

	if f.State != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("state")+" = ?")
		values = append(values, f.State)
	}

	if f.StateNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("state")+" != ?")
		values = append(values, f.StateNe)
	}

	if f.StateGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("state")+" > ?")
		values = append(values, f.StateGt)
	}

	if f.StateLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("state")+" < ?")
		values = append(values, f.StateLt)
	}

	if f.StateGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("state")+" >= ?")
		values = append(values, f.StateGte)
	}

	if f.StateLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("state")+" <= ?")
		values = append(values, f.StateLte)
	}

	if f.StateIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("state")+" IN (?)")
		values = append(values, f.StateIn)
	}

	if f.StateNull != nil {
		if *f.StateNull {
			conditions = append(conditions, aliasPrefix+SnakeString("state")+" IS NULL"+" OR "+aliasPrefix+SnakeString("state")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("state")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("state")+" <> ''")
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
func (f *TaskFilterType) AndWith(f2 ...*TaskFilterType) *TaskFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &TaskFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *TaskFilterType) OrWith(f2 ...*TaskFilterType) *TaskFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &TaskFilterType{
		Or: append(_f2, f),
	}
}
