package gen

import (
	"context"
	"fmt"
	"strings"
)

func (f *BookCategoryFilterType) IsEmpty(ctx context.Context) bool {
	wheres := []string{}
	values := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, "companies", &wheres, &values, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0
}
func (f *BookCategoryFilterType) Apply(ctx context.Context, wheres *[]string, values *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, TableName("book_categories"), wheres, values, joins)
}
func (f *BookCategoryFilterType) ApplyWithAlias(ctx context.Context, alias string, wheres *[]string, values *[]interface{}, joins *[]string) error {
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

	if f.Books != nil {
		_alias := alias + "_books"
		*joins = append(*joins, "LEFT JOIN "+"books"+" "+_alias+" ON "+_alias+"."+"category_id"+" = "+alias+".id")
		err := f.Books.ApplyWithAlias(ctx, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *BookCategoryFilterType) WhereContent(aliasPrefix string) (conditions []string, values []interface{}) {
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

	if f.Name != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("name")+" = ?")
		values = append(values, f.Name)
	}

	if f.NameNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("name")+" != ?")
		values = append(values, f.NameNe)
	}

	if f.NameGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("name")+" > ?")
		values = append(values, f.NameGt)
	}

	if f.NameLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("name")+" < ?")
		values = append(values, f.NameLt)
	}

	if f.NameGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("name")+" >= ?")
		values = append(values, f.NameGte)
	}

	if f.NameLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("name")+" <= ?")
		values = append(values, f.NameLte)
	}

	if f.NameIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("name")+" IN (?)")
		values = append(values, f.NameIn)
	}

	if f.NameLike != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("name")+" LIKE ?")
		values = append(values, "%"+strings.Replace(strings.Replace(*f.NameLike, "?", "_", -1), "*", "%", -1)+"%")
	}

	if f.NamePrefix != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("name")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NamePrefix))
	}

	if f.NameSuffix != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("name")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameSuffix))
	}

	if f.NameNull != nil {
		if *f.NameNull {
			conditions = append(conditions, aliasPrefix+SnakeString("name")+" IS NULL"+" OR "+aliasPrefix+SnakeString("name")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("name")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("name")+" <> ''")
		}
	}

	if f.Description != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("description")+" = ?")
		values = append(values, f.Description)
	}

	if f.DescriptionNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("description")+" != ?")
		values = append(values, f.DescriptionNe)
	}

	if f.DescriptionGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("description")+" > ?")
		values = append(values, f.DescriptionGt)
	}

	if f.DescriptionLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("description")+" < ?")
		values = append(values, f.DescriptionLt)
	}

	if f.DescriptionGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("description")+" >= ?")
		values = append(values, f.DescriptionGte)
	}

	if f.DescriptionLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("description")+" <= ?")
		values = append(values, f.DescriptionLte)
	}

	if f.DescriptionIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("description")+" IN (?)")
		values = append(values, f.DescriptionIn)
	}

	if f.DescriptionLike != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("description")+" LIKE ?")
		values = append(values, "%"+strings.Replace(strings.Replace(*f.DescriptionLike, "?", "_", -1), "*", "%", -1)+"%")
	}

	if f.DescriptionPrefix != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("description")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DescriptionPrefix))
	}

	if f.DescriptionSuffix != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("description")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DescriptionSuffix))
	}

	if f.DescriptionNull != nil {
		if *f.DescriptionNull {
			conditions = append(conditions, aliasPrefix+SnakeString("description")+" IS NULL"+" OR "+aliasPrefix+SnakeString("description")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("description")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("description")+" <> ''")
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
func (f *BookCategoryFilterType) AndWith(f2 ...*BookCategoryFilterType) *BookCategoryFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &BookCategoryFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *BookCategoryFilterType) OrWith(f2 ...*BookCategoryFilterType) *BookCategoryFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &BookCategoryFilterType{
		Or: append(_f2, f),
	}
}

func (f *BookFilterType) IsEmpty(ctx context.Context) bool {
	wheres := []string{}
	values := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, "companies", &wheres, &values, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0
}
func (f *BookFilterType) Apply(ctx context.Context, wheres *[]string, values *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, TableName("books"), wheres, values, joins)
}
func (f *BookFilterType) ApplyWithAlias(ctx context.Context, alias string, wheres *[]string, values *[]interface{}, joins *[]string) error {
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

	if f.Category != nil {
		_alias := alias + "_category"
		*joins = append(*joins, "LEFT JOIN "+"book_categories"+" "+_alias+" ON "+_alias+".id = "+alias+"."+"category_id")
		err := f.Category.ApplyWithAlias(ctx, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *BookFilterType) WhereContent(aliasPrefix string) (conditions []string, values []interface{}) {
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

	if f.Author != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("author")+" = ?")
		values = append(values, f.Author)
	}

	if f.AuthorNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("author")+" != ?")
		values = append(values, f.AuthorNe)
	}

	if f.AuthorGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("author")+" > ?")
		values = append(values, f.AuthorGt)
	}

	if f.AuthorLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("author")+" < ?")
		values = append(values, f.AuthorLt)
	}

	if f.AuthorGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("author")+" >= ?")
		values = append(values, f.AuthorGte)
	}

	if f.AuthorLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("author")+" <= ?")
		values = append(values, f.AuthorLte)
	}

	if f.AuthorIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("author")+" IN (?)")
		values = append(values, f.AuthorIn)
	}

	if f.AuthorLike != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("author")+" LIKE ?")
		values = append(values, "%"+strings.Replace(strings.Replace(*f.AuthorLike, "?", "_", -1), "*", "%", -1)+"%")
	}

	if f.AuthorPrefix != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("author")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.AuthorPrefix))
	}

	if f.AuthorSuffix != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("author")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.AuthorSuffix))
	}

	if f.AuthorNull != nil {
		if *f.AuthorNull {
			conditions = append(conditions, aliasPrefix+SnakeString("author")+" IS NULL"+" OR "+aliasPrefix+SnakeString("author")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("author")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("author")+" <> ''")
		}
	}

	if f.Price != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("price")+" = ?")
		values = append(values, f.Price)
	}

	if f.PriceNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("price")+" != ?")
		values = append(values, f.PriceNe)
	}

	if f.PriceGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("price")+" > ?")
		values = append(values, f.PriceGt)
	}

	if f.PriceLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("price")+" < ?")
		values = append(values, f.PriceLt)
	}

	if f.PriceGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("price")+" >= ?")
		values = append(values, f.PriceGte)
	}

	if f.PriceLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("price")+" <= ?")
		values = append(values, f.PriceLte)
	}

	if f.PriceIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("price")+" IN (?)")
		values = append(values, f.PriceIn)
	}

	if f.PriceNull != nil {
		if *f.PriceNull {
			conditions = append(conditions, aliasPrefix+SnakeString("price")+" IS NULL"+" OR "+aliasPrefix+SnakeString("price")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("price")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("price")+" <> ''")
		}
	}

	if f.PublishDateAt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("publishDateAt")+" = ?")
		values = append(values, f.PublishDateAt)
	}

	if f.PublishDateAtNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("publishDateAt")+" != ?")
		values = append(values, f.PublishDateAtNe)
	}

	if f.PublishDateAtGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("publishDateAt")+" > ?")
		values = append(values, f.PublishDateAtGt)
	}

	if f.PublishDateAtLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("publishDateAt")+" < ?")
		values = append(values, f.PublishDateAtLt)
	}

	if f.PublishDateAtGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("publishDateAt")+" >= ?")
		values = append(values, f.PublishDateAtGte)
	}

	if f.PublishDateAtLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("publishDateAt")+" <= ?")
		values = append(values, f.PublishDateAtLte)
	}

	if f.PublishDateAtIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("publishDateAt")+" IN (?)")
		values = append(values, f.PublishDateAtIn)
	}

	if f.PublishDateAtNull != nil {
		if *f.PublishDateAtNull {
			conditions = append(conditions, aliasPrefix+SnakeString("publishDateAt")+" IS NULL"+" OR "+aliasPrefix+SnakeString("publishDateAt")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("publishDateAt")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("publishDateAt")+" <> ''")
		}
	}

	if f.CategoryID != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("categoryId")+" = ?")
		values = append(values, f.CategoryID)
	}

	if f.CategoryIDNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("categoryId")+" != ?")
		values = append(values, f.CategoryIDNe)
	}

	if f.CategoryIDGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("categoryId")+" > ?")
		values = append(values, f.CategoryIDGt)
	}

	if f.CategoryIDLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("categoryId")+" < ?")
		values = append(values, f.CategoryIDLt)
	}

	if f.CategoryIDGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("categoryId")+" >= ?")
		values = append(values, f.CategoryIDGte)
	}

	if f.CategoryIDLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("categoryId")+" <= ?")
		values = append(values, f.CategoryIDLte)
	}

	if f.CategoryIDIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("categoryId")+" IN (?)")
		values = append(values, f.CategoryIDIn)
	}

	if f.CategoryIDNull != nil {
		if *f.CategoryIDNull {
			conditions = append(conditions, aliasPrefix+SnakeString("categoryId")+" IS NULL"+" OR "+aliasPrefix+SnakeString("categoryId")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("categoryId")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("categoryId")+" <> ''")
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
func (f *BookFilterType) AndWith(f2 ...*BookFilterType) *BookFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &BookFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *BookFilterType) OrWith(f2 ...*BookFilterType) *BookFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &BookFilterType{
		Or: append(_f2, f),
	}
}
