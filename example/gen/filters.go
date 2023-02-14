package gen

import (
	"context"
	"fmt"
	"strings"
)

func (f *TodoFilterType) IsEmpty(ctx context.Context) bool {
	wheres := []string{}
	values := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, "companies", &wheres, &values, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0
}
func (f *TodoFilterType) Apply(ctx context.Context, wheres *[]string, values *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, TableName("todos"), wheres, values, joins)
}
func (f *TodoFilterType) ApplyWithAlias(ctx context.Context, alias string, wheres *[]string, values *[]interface{}, joins *[]string) error {
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

func (f *TodoFilterType) WhereContent(aliasPrefix string) (conditions []string, values []interface{}) {
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

	if f.Age != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("age")+" = ?")
		values = append(values, f.Age)
	}

	if f.AgeNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("age")+" != ?")
		values = append(values, f.AgeNe)
	}

	if f.AgeGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("age")+" > ?")
		values = append(values, f.AgeGt)
	}

	if f.AgeLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("age")+" < ?")
		values = append(values, f.AgeLt)
	}

	if f.AgeGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("age")+" >= ?")
		values = append(values, f.AgeGte)
	}

	if f.AgeLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("age")+" <= ?")
		values = append(values, f.AgeLte)
	}

	if f.AgeIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("age")+" IN (?)")
		values = append(values, f.AgeIn)
	}

	if f.AgeNull != nil {
		if *f.AgeNull {
			conditions = append(conditions, aliasPrefix+SnakeString("age")+" IS NULL"+" OR "+aliasPrefix+SnakeString("age")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("age")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("age")+" <> ''")
		}
	}

	if f.Money != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("money")+" = ?")
		values = append(values, f.Money)
	}

	if f.MoneyNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("money")+" != ?")
		values = append(values, f.MoneyNe)
	}

	if f.MoneyGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("money")+" > ?")
		values = append(values, f.MoneyGt)
	}

	if f.MoneyLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("money")+" < ?")
		values = append(values, f.MoneyLt)
	}

	if f.MoneyGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("money")+" >= ?")
		values = append(values, f.MoneyGte)
	}

	if f.MoneyLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("money")+" <= ?")
		values = append(values, f.MoneyLte)
	}

	if f.MoneyIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("money")+" IN (?)")
		values = append(values, f.MoneyIn)
	}

	if f.MoneyNull != nil {
		if *f.MoneyNull {
			conditions = append(conditions, aliasPrefix+SnakeString("money")+" IS NULL"+" OR "+aliasPrefix+SnakeString("money")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("money")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("money")+" <> ''")
		}
	}

	if f.Remark != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("remark")+" = ?")
		values = append(values, f.Remark)
	}

	if f.RemarkNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("remark")+" != ?")
		values = append(values, f.RemarkNe)
	}

	if f.RemarkGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("remark")+" > ?")
		values = append(values, f.RemarkGt)
	}

	if f.RemarkLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("remark")+" < ?")
		values = append(values, f.RemarkLt)
	}

	if f.RemarkGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("remark")+" >= ?")
		values = append(values, f.RemarkGte)
	}

	if f.RemarkLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("remark")+" <= ?")
		values = append(values, f.RemarkLte)
	}

	if f.RemarkIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("remark")+" IN (?)")
		values = append(values, f.RemarkIn)
	}

	if f.RemarkLike != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("remark")+" LIKE ?")
		values = append(values, "%"+strings.Replace(strings.Replace(*f.RemarkLike, "?", "_", -1), "*", "%", -1)+"%")
	}

	if f.RemarkPrefix != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("remark")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.RemarkPrefix))
	}

	if f.RemarkSuffix != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("remark")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.RemarkSuffix))
	}

	if f.RemarkNull != nil {
		if *f.RemarkNull {
			conditions = append(conditions, aliasPrefix+SnakeString("remark")+" IS NULL"+" OR "+aliasPrefix+SnakeString("remark")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("remark")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("remark")+" <> ''")
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
func (f *TodoFilterType) AndWith(f2 ...*TodoFilterType) *TodoFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &TodoFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *TodoFilterType) OrWith(f2 ...*TodoFilterType) *TodoFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &TodoFilterType{
		Or: append(_f2, f),
	}
}
