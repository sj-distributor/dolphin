package gen

import (
	"context"
	"math"

	"github.com/99designs/gqlgen/graphql"
	"github.com/graph-gophers/dataloader"
	"github.com/vektah/gqlparser/v2/ast"
)

type GeneratedQueryResolver struct{ *GeneratedResolver }

func (r *GeneratedQueryResolver) User(ctx context.Context, id string) (*User, error) {
	return r.Handlers.QueryUser(ctx, r.GeneratedResolver, id)
}
func QueryUserHandler(ctx context.Context, r *GeneratedResolver, id string) (*User, error) {
	selection := []ast.Selection{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		selection = append(selection, f.Field)
	}
	selectionSet := ast.SelectionSet(selection)

	query := UserQueryFilter{}
	rt := &UserResultType{
		EntityResultType: EntityResultType{
			Query:        &query,
			SelectionSet: &selectionSet,
		},
	}
	qb := r.DB.Query()
	qb = qb.Where(TableName("users")+".id = ?", id)

	var items []*User
	giOpts := GetItemsOptions{
		Alias:      TableName("users"),
		Preloaders: []string{},
		Item:       &User{},
	}
	err := rt.GetData(ctx, qb, giOpts, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, &NotFoundError{Entity: "User"}
	}
	return items[0], err
}

type QueryUsersHandlerOptions struct {
	CurrentPage *int
	PerPage     *int
	Q           *string
	Sort        []*UserSortType
	Filter      *UserFilterType
	Rand        *bool
}

func (r *GeneratedQueryResolver) Users(ctx context.Context, current_page *int, per_page *int, q *string, sort []*UserSortType, filter *UserFilterType, rand *bool) (*UserResultType, error) {
	opts := QueryUsersHandlerOptions{
		CurrentPage: current_page,
		PerPage:     per_page,
		Q:           q,
		Sort:        sort,
		Filter:      filter,
		Rand:        rand,
	}
	return r.Handlers.QueryUsers(ctx, r.GeneratedResolver, opts)
}
func QueryUsersHandler(ctx context.Context, r *GeneratedResolver, opts QueryUsersHandlerOptions) (*UserResultType, error) {
	query := UserQueryFilter{opts.Q}

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

	return &UserResultType{
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

type GeneratedUserResultTypeResolver struct{ *GeneratedResolver }

func (r *GeneratedUserResultTypeResolver) Data(ctx context.Context, obj *UserResultType) (items []*User, err error) {
	giOpts := GetItemsOptions{
		Alias:      TableName("users"),
		Preloaders: []string{},
		Item:       &User{},
	}
	err = obj.GetData(ctx, r.DB.db, giOpts, &items)

	uniqueItems := []*User{}
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

func (r *GeneratedUserResultTypeResolver) Pages(ctx context.Context, obj *UserResultType) (interface{}, error) {
	total, _ := r.Total(ctx, obj)
	totalPage, _ := r.TotalPage(ctx, obj)
	perPage, _ := r.PerPage(ctx, obj)
	currentPage, _ := r.CurrentPage(ctx, obj)

	return map[string]int{
		"total":        total,
		"total_page":   totalPage,
		"per_page":     perPage,
		"current_page": currentPage,
	}, nil
}

func (r *GeneratedUserResultTypeResolver) Total(ctx context.Context, obj *UserResultType) (count int, err error) {
	return obj.GetTotal(ctx, r.DB.db, TableName("users"), &User{})
}

func (r *GeneratedUserResultTypeResolver) CurrentPage(ctx context.Context, obj *UserResultType) (count int, err error) {
	return int(*obj.EntityResultType.CurrentPage), nil
}

func (r *GeneratedUserResultTypeResolver) PerPage(ctx context.Context, obj *UserResultType) (count int, err error) {
	return int(*obj.EntityResultType.PerPage), nil
}

func (r *GeneratedUserResultTypeResolver) TotalPage(ctx context.Context, obj *UserResultType) (count int, err error) {
	total, _ := r.Total(ctx, obj)
	perPage, _ := r.PerPage(ctx, obj)
	totalPage := int(math.Ceil(float64(total) / float64(perPage)))
	if totalPage < 0 {
		totalPage = 0
	}

	return totalPage, nil
}

type GeneratedUserResolver struct{ *GeneratedResolver }

func (r *GeneratedUserResolver) Todo(ctx context.Context, obj *User) (res *Todo, err error) {

	return r.Handlers.UserTodo(ctx, r.GeneratedResolver, obj)

}

func UserTodoHandler(ctx context.Context, r *GeneratedResolver, obj *User) (res *Todo, err error) {

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	if obj.TodoID != nil {
		item, _ := loaders["Todo"].Load(ctx, dataloader.StringKey(*obj.TodoID))()
		res, _ = item.(*Todo)

		// err = _err
	}

	return
}

func (r *GeneratedQueryResolver) Todo(ctx context.Context, id string) (*Todo, error) {
	return r.Handlers.QueryTodo(ctx, r.GeneratedResolver, id)
}
func QueryTodoHandler(ctx context.Context, r *GeneratedResolver, id string) (*Todo, error) {
	selection := []ast.Selection{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		selection = append(selection, f.Field)
	}
	selectionSet := ast.SelectionSet(selection)

	query := TodoQueryFilter{}
	rt := &TodoResultType{
		EntityResultType: EntityResultType{
			Query:        &query,
			SelectionSet: &selectionSet,
		},
	}
	qb := r.DB.Query()
	qb = qb.Where(TableName("todos")+".id = ?", id)

	var items []*Todo
	giOpts := GetItemsOptions{
		Alias:      TableName("todos"),
		Preloaders: []string{},
		Item:       &Todo{},
	}
	err := rt.GetData(ctx, qb, giOpts, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, &NotFoundError{Entity: "Todo"}
	}
	return items[0], err
}

type QueryTodosHandlerOptions struct {
	CurrentPage *int
	PerPage     *int
	Q           *string
	Sort        []*TodoSortType
	Filter      *TodoFilterType
	Rand        *bool
}

func (r *GeneratedQueryResolver) Todos(ctx context.Context, current_page *int, per_page *int, q *string, sort []*TodoSortType, filter *TodoFilterType, rand *bool) (*TodoResultType, error) {
	opts := QueryTodosHandlerOptions{
		CurrentPage: current_page,
		PerPage:     per_page,
		Q:           q,
		Sort:        sort,
		Filter:      filter,
		Rand:        rand,
	}
	return r.Handlers.QueryTodos(ctx, r.GeneratedResolver, opts)
}
func QueryTodosHandler(ctx context.Context, r *GeneratedResolver, opts QueryTodosHandlerOptions) (*TodoResultType, error) {
	query := TodoQueryFilter{opts.Q}

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

	return &TodoResultType{
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

type GeneratedTodoResultTypeResolver struct{ *GeneratedResolver }

func (r *GeneratedTodoResultTypeResolver) Data(ctx context.Context, obj *TodoResultType) (items []*Todo, err error) {
	giOpts := GetItemsOptions{
		Alias:      TableName("todos"),
		Preloaders: []string{},
		Item:       &Todo{},
	}
	err = obj.GetData(ctx, r.DB.db, giOpts, &items)

	uniqueItems := []*Todo{}
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

func (r *GeneratedTodoResultTypeResolver) Pages(ctx context.Context, obj *TodoResultType) (interface{}, error) {
	total, _ := r.Total(ctx, obj)
	totalPage, _ := r.TotalPage(ctx, obj)
	perPage, _ := r.PerPage(ctx, obj)
	currentPage, _ := r.CurrentPage(ctx, obj)

	return map[string]int{
		"total":        total,
		"total_page":   totalPage,
		"per_page":     perPage,
		"current_page": currentPage,
	}, nil
}

func (r *GeneratedTodoResultTypeResolver) Total(ctx context.Context, obj *TodoResultType) (count int, err error) {
	return obj.GetTotal(ctx, r.DB.db, TableName("todos"), &Todo{})
}

func (r *GeneratedTodoResultTypeResolver) CurrentPage(ctx context.Context, obj *TodoResultType) (count int, err error) {
	return int(*obj.EntityResultType.CurrentPage), nil
}

func (r *GeneratedTodoResultTypeResolver) PerPage(ctx context.Context, obj *TodoResultType) (count int, err error) {
	return int(*obj.EntityResultType.PerPage), nil
}

func (r *GeneratedTodoResultTypeResolver) TotalPage(ctx context.Context, obj *TodoResultType) (count int, err error) {
	total, _ := r.Total(ctx, obj)
	perPage, _ := r.PerPage(ctx, obj)
	totalPage := int(math.Ceil(float64(total) / float64(perPage)))
	if totalPage < 0 {
		totalPage = 0
	}

	return totalPage, nil
}

type GeneratedTodoResolver struct{ *GeneratedResolver }

func (r *GeneratedTodoResolver) User(ctx context.Context, obj *Todo) (res *User, err error) {

	return r.Handlers.TodoUser(ctx, r.GeneratedResolver, obj)

}

func TodoUserHandler(ctx context.Context, r *GeneratedResolver, obj *Todo) (res *User, err error) {

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	if obj.UserID != nil {
		item, _ := loaders["User"].Load(ctx, dataloader.StringKey(*obj.UserID))()
		res, _ = item.(*User)

		// err = _err
	}

	return
}
