package gen

import (
	"context"
	"strings"

	"github.com/vektah/gqlparser/v2/ast"
	"gorm.io/gorm"
)

func GetItem(ctx context.Context, db *gorm.DB, table string, out interface{}, id *string) error {
	return db.First(out, table+".id = ?", id).Error
}

type EntityFilter interface {
	Apply(ctx context.Context, wheres *[]string, values *[]interface{}, joins *[]string) error
}

type EntityFilterQuery interface {
	Apply(ctx context.Context, db *gorm.DB, selectionSet *ast.SelectionSet, wheres *[]string, values *[]interface{}, joins *[]string) error
}

type EntitySort interface {
	Apply(ctx context.Context, sorts *[]string, joins *[]string) error
}

type EntityResultType struct {
	Offset       *int
	Limit        *int
	CurrentPage  *int
	PerPage      *int
	Rand         *bool
	Query        EntityFilterQuery
	Sort         []EntitySort
	Filter       EntityFilter
	Fields       []*ast.Field
	SelectionSet *ast.SelectionSet
}

type GetItemsOptions struct {
	Alias      string
	Preloaders []string
	Item       interface{}
}

type CountResult struct {
	Count int
}

func (r *EntityResultType) GetData(ctx context.Context, db *gorm.DB, opts GetItemsOptions, out interface{}) error {
	q := db

	selects := GetFieldsRequested(ctx, opts.Alias)
	if len(selects) > 0 && IndexOf(selects, opts.Alias+".id") == -1 {
		selects = append(selects, opts.Alias+".id")
	}

	if len(selects) > 0 {
		q = q.Select(selects)
	}

	if r.PerPage != nil {
		if int(*r.PerPage) != 0 {
			q = q.Limit(*r.PerPage)
		}
	}

	if r.CurrentPage != nil {
		q = q.Offset((int(*r.CurrentPage) - 1) * int(*r.PerPage))
	}

	wheres := []string{}
	values := []interface{}{}
	joins := []string{}
	sorts := []string{}

	if r.Rand != nil && *r.Rand {
		sorts = append(sorts, "Rand()")
	}

	err := r.Query.Apply(ctx, db, r.SelectionSet, &wheres, &values, &joins)
	if err != nil {
		return err
	}

	for _, sort := range r.Sort {
		sort.Apply(ctx, &sorts, &joins)
	}

	if r.Filter != nil {
		err = r.Filter.Apply(ctx, &wheres, &values, &joins)
		if err != nil {
			return err
		}
	}

	isAt := false

	for _, s := range sorts {
		if strings.Contains(s, "_at") {
			isAt = true
		}
	}

	if !isAt {
		sorts = append(sorts, opts.Alias+".created_at DESC")
	}

	if len(sorts) > 0 {
		q = q.Order(strings.Join(sorts, ", "))
	}

	if len(wheres) > 0 {
		q = q.Where(strings.Join(wheres, " AND "), values...)
	}

	uniqueJoinsMap := map[string]bool{}
	uniqueJoins := []string{}
	for _, join := range joins {
		if !uniqueJoinsMap[join] {
			uniqueJoinsMap[join] = true
			uniqueJoins = append(uniqueJoins, join)
		}
	}
	for _, join := range uniqueJoins {
		q = q.Joins(join)
	}
	if len(opts.Preloaders) > 0 {
		for _, p := range opts.Preloaders {
			q = q.Preload(p)
		}
	}

	return q.Find(out).Error
}

// GetTotal ...
func (r *EntityResultType) GetTotal(ctx context.Context, db *gorm.DB, table string, out interface{}) (count int, err error) {
	q := db.Model(out).Select(table + ".id")

	wheres := []string{}
	values := []interface{}{}
	joins := []string{}

	err = r.Query.Apply(ctx, db, r.SelectionSet, &wheres, &values, &joins)
	if err != nil {
		return 0, err
	}

	if r.Filter != nil {
		err = r.Filter.Apply(ctx, &wheres, &values, &joins)
		if err != nil {
			return 0, err
		}
	}

	if len(wheres) > 0 {
		q = q.Where(strings.Join(wheres, " AND "), values...)
	}

	uniqueJoinsMap := map[string]bool{}
	uniqueJoins := []string{}
	for _, join := range joins {
		if !uniqueJoinsMap[join] {
			uniqueJoinsMap[join] = true
			uniqueJoins = append(uniqueJoins, join)
		}
	}

	for _, join := range uniqueJoins {
		q = q.Joins(join)
	}

	var result CountResult

	err = q.Select("COUNT(DISTINCT " + table + ".id) as count").Scan(&result).Error

	count = result.Count

	return
}

func (r *EntityResultType) GetSortStrings() []string {
	return []string{}
}
