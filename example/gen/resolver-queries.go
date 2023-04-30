package gen

import (
	"context"
	"math"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/graph-gophers/dataloader"
	"github.com/vektah/gqlparser/v2/ast"
)

type GeneratedQueryResolver struct{ *GeneratedResolver }

func (r *GeneratedQueryResolver) BookCategory(ctx context.Context, id string) (*BookCategory, error) {
	return r.Handlers.QueryBookCategory(ctx, r.GeneratedResolver, id)
}
func QueryBookCategoryHandler(ctx context.Context, r *GeneratedResolver, id string) (*BookCategory, error) {
	selection := []ast.Selection{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		selection = append(selection, f.Field)
	}
	selectionSet := ast.SelectionSet(selection)

	query := BookCategoryQueryFilter{}
	rt := &BookCategoryResultType{
		EntityResultType: EntityResultType{
			Query:        &query,
			SelectionSet: &selectionSet,
		},
	}
	qb := r.DB.Query()
	qb = qb.Where(TableName("book_categories")+".id = ?", id)

	var items []*BookCategory
	giOpts := GetItemsOptions{
		Alias:      TableName("book_categories"),
		Preloaders: []string{},
		Item:       &BookCategory{},
	}
	err := rt.GetData(ctx, qb, giOpts, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, &NotFoundError{Entity: "BookCategory"}
	}
	return items[0], err
}

type QueryBookCategoriesHandlerOptions struct {
	CurrentPage *int
	PerPage     *int
	Q           *string
	Sort        []*BookCategorySortType
	Filter      *BookCategoryFilterType
	Rand        *bool
}

func (r *GeneratedQueryResolver) BookCategories(ctx context.Context, current_page *int, per_page *int, q *string, sort []*BookCategorySortType, filter *BookCategoryFilterType, rand *bool) (*BookCategoryResultType, error) {
	opts := QueryBookCategoriesHandlerOptions{
		CurrentPage: current_page,
		PerPage:     per_page,
		Q:           q,
		Sort:        sort,
		Filter:      filter,
		Rand:        rand,
	}
	return r.Handlers.QueryBookCategories(ctx, r.GeneratedResolver, opts)
}
func QueryBookCategoriesHandler(ctx context.Context, r *GeneratedResolver, opts QueryBookCategoriesHandlerOptions) (*BookCategoryResultType, error) {
	query := BookCategoryQueryFilter{opts.Q}

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

	return &BookCategoryResultType{
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

type GeneratedBookCategoryResultTypeResolver struct{ *GeneratedResolver }

func (r *GeneratedBookCategoryResultTypeResolver) Data(ctx context.Context, obj *BookCategoryResultType) (items []*BookCategory, err error) {
	giOpts := GetItemsOptions{
		Alias:      TableName("book_categories"),
		Preloaders: []string{},
		Item:       &BookCategory{},
	}
	err = obj.GetData(ctx, r.DB.db, giOpts, &items)

	uniqueItems := []*BookCategory{}
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

func (r *GeneratedBookCategoryResultTypeResolver) Total(ctx context.Context, obj *BookCategoryResultType) (count int, err error) {
	return obj.GetTotal(ctx, r.DB.db, TableName("book_categories"), &BookCategory{})
}

func (r *GeneratedBookCategoryResultTypeResolver) TotalPage(ctx context.Context, obj *BookCategoryResultType) (count int, err error) {
	total, _ := r.Total(ctx, obj)
	perPage, _ := r.PerPage(ctx, obj)
	totalPage := int(math.Ceil(float64(total) / float64(perPage)))
	if totalPage < 0 {
		totalPage = 0
	}

	return totalPage, nil
}

func (r *GeneratedBookCategoryResultTypeResolver) CurrentPage(ctx context.Context, obj *BookCategoryResultType) (count int, err error) {
	return int(*obj.EntityResultType.CurrentPage), nil
}

func (r *GeneratedBookCategoryResultTypeResolver) PerPage(ctx context.Context, obj *BookCategoryResultType) (count int, err error) {
	return int(*obj.EntityResultType.PerPage), nil
}

type GeneratedBookCategoryResolver struct{ *GeneratedResolver }

func (r *GeneratedBookCategoryResolver) Books(ctx context.Context, obj *BookCategory) (res []*Book, err error) {

	var input map[string]interface{}
	return r.Handlers.BookCategoryBooks(ctx, r.GeneratedResolver, obj, input)

}

func BookCategoryBooksHandler(ctx context.Context, r *GeneratedResolver, obj *BookCategory, input map[string]interface{}) (items []*Book, err error) {

	selects := GetFieldsRequested(ctx, strings.ToLower(TableName("books")))
	items = []*Book{}
	wheres := []string{}
	values := []interface{}{}

	if err := r.DB.Query().Select(selects).Where(strings.Join(wheres, " AND "), values...).Model(obj).Association("Books").Find(&items); err != nil {
		return items, err
	}

	return
}

func (r *GeneratedBookCategoryResolver) BooksIds(ctx context.Context, obj *BookCategory) (ids []string, err error) {
	wheres := []string{}
	values := []interface{}{}

	ids = []string{}
	items := []*Book{}
	if err := r.DB.Query().Model(obj).Select(TableName("books")+".id").Where(strings.Join(wheres, " AND "), values...).Association("Books").Find(&items); err != nil {
		return ids, err
	}

	for _, item := range items {
		ids = append(ids, item.ID)
	}

	return
}

func (r *GeneratedQueryResolver) Book(ctx context.Context, id string) (*Book, error) {
	return r.Handlers.QueryBook(ctx, r.GeneratedResolver, id)
}
func QueryBookHandler(ctx context.Context, r *GeneratedResolver, id string) (*Book, error) {
	selection := []ast.Selection{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		selection = append(selection, f.Field)
	}
	selectionSet := ast.SelectionSet(selection)

	query := BookQueryFilter{}
	rt := &BookResultType{
		EntityResultType: EntityResultType{
			Query:        &query,
			SelectionSet: &selectionSet,
		},
	}
	qb := r.DB.Query()
	qb = qb.Where(TableName("books")+".id = ?", id)

	var items []*Book
	giOpts := GetItemsOptions{
		Alias:      TableName("books"),
		Preloaders: []string{},
		Item:       &Book{},
	}
	err := rt.GetData(ctx, qb, giOpts, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, &NotFoundError{Entity: "Book"}
	}
	return items[0], err
}

type QueryBooksHandlerOptions struct {
	CurrentPage *int
	PerPage     *int
	Q           *string
	Sort        []*BookSortType
	Filter      *BookFilterType
	Rand        *bool
}

func (r *GeneratedQueryResolver) Books(ctx context.Context, current_page *int, per_page *int, q *string, sort []*BookSortType, filter *BookFilterType, rand *bool) (*BookResultType, error) {
	opts := QueryBooksHandlerOptions{
		CurrentPage: current_page,
		PerPage:     per_page,
		Q:           q,
		Sort:        sort,
		Filter:      filter,
		Rand:        rand,
	}
	return r.Handlers.QueryBooks(ctx, r.GeneratedResolver, opts)
}
func QueryBooksHandler(ctx context.Context, r *GeneratedResolver, opts QueryBooksHandlerOptions) (*BookResultType, error) {
	query := BookQueryFilter{opts.Q}

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

	return &BookResultType{
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

type GeneratedBookResultTypeResolver struct{ *GeneratedResolver }

func (r *GeneratedBookResultTypeResolver) Data(ctx context.Context, obj *BookResultType) (items []*Book, err error) {
	giOpts := GetItemsOptions{
		Alias:      TableName("books"),
		Preloaders: []string{},
		Item:       &Book{},
	}
	err = obj.GetData(ctx, r.DB.db, giOpts, &items)

	uniqueItems := []*Book{}
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

func (r *GeneratedBookResultTypeResolver) Total(ctx context.Context, obj *BookResultType) (count int, err error) {
	return obj.GetTotal(ctx, r.DB.db, TableName("books"), &Book{})
}

func (r *GeneratedBookResultTypeResolver) TotalPage(ctx context.Context, obj *BookResultType) (count int, err error) {
	total, _ := r.Total(ctx, obj)
	perPage, _ := r.PerPage(ctx, obj)
	totalPage := int(math.Ceil(float64(total) / float64(perPage)))
	if totalPage < 0 {
		totalPage = 0
	}

	return totalPage, nil
}

func (r *GeneratedBookResultTypeResolver) CurrentPage(ctx context.Context, obj *BookResultType) (count int, err error) {
	return int(*obj.EntityResultType.CurrentPage), nil
}

func (r *GeneratedBookResultTypeResolver) PerPage(ctx context.Context, obj *BookResultType) (count int, err error) {
	return int(*obj.EntityResultType.PerPage), nil
}

type GeneratedBookResolver struct{ *GeneratedResolver }

func (r *GeneratedBookResolver) Category(ctx context.Context, obj *Book) (res *BookCategory, err error) {

	return r.Handlers.BookCategory(ctx, r.GeneratedResolver, obj)

}

func BookCategoryHandler(ctx context.Context, r *GeneratedResolver, obj *Book) (res *BookCategory, err error) {

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	if obj.CategoryID != nil {
		item, _ := loaders["BookCategory"].Load(ctx, dataloader.StringKey(*obj.CategoryID))()
		res, _ = item.(*BookCategory)

		if res == nil {
			res = &BookCategory{}
			// 	_err = fmt.Errorf("BookCategory with id '%s' not found",*obj.CategoryID)
		}
		// err = _err
	}

	return
}
