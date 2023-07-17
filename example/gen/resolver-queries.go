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

type QueryUserHandlerOptions struct {
	ID     *string
	Filter *UserFilterType
}

func (r *GeneratedQueryResolver) User(ctx context.Context, id *string, filter *UserFilterType) (*User, error) {
	opts := QueryUserHandlerOptions{
		ID:     id,
		Filter: filter,
	}
	return r.Handlers.QueryUser(ctx, r.GeneratedResolver, opts)
}
func QueryUserHandler(ctx context.Context, r *GeneratedResolver, opts QueryUserHandlerOptions) (*User, error) {
	selection := []ast.Selection{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		selection = append(selection, f.Field)
	}
	selectionSet := ast.SelectionSet(selection)

	query := UserQueryFilter{}
	rt := &UserResultType{
		EntityResultType: EntityResultType{
			Query:        &query,
			Filter:       opts.Filter,
			SelectionSet: &selectionSet,
		},
	}
	qb := r.DB.Query()
	if opts.ID != nil {
		qb = qb.Where(TableName("users")+".id = ?", *opts.ID)
	}

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

func (r *GeneratedUserResultTypeResolver) Total(ctx context.Context, obj *UserResultType) (count int, err error) {
	return obj.GetTotal(ctx, r.DB.db, TableName("users"), &User{})
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

func (r *GeneratedUserResultTypeResolver) CurrentPage(ctx context.Context, obj *UserResultType) (count int, err error) {
	return int(*obj.EntityResultType.CurrentPage), nil
}

func (r *GeneratedUserResultTypeResolver) PerPage(ctx context.Context, obj *UserResultType) (count int, err error) {
	return int(*obj.EntityResultType.PerPage), nil
}

type GeneratedUserResolver struct{ *GeneratedResolver }

func (r *GeneratedUserResolver) T(ctx context.Context, obj *User) (res *Task, err error) {
	return r.Handlers.UserT(ctx, r.GeneratedResolver, obj)
}
func UserTHandler(ctx context.Context, r *GeneratedResolver, obj *User) (items *Task, err error) {

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	if obj.TID != nil {
		item, _ := loaders["Task"].Load(ctx, dataloader.StringKey(*obj.TID))()
		items, _ = item.(*Task)

		if items == nil {
			items = &Task{}
		}

	}

	return
}

func (r *GeneratedUserResolver) Tt(ctx context.Context, obj *User) (res *Task, err error) {
	return r.Handlers.UserTt(ctx, r.GeneratedResolver, obj)
}
func UserTtHandler(ctx context.Context, r *GeneratedResolver, obj *User) (items *Task, err error) {

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	if obj.TtID != nil {
		item, _ := loaders["Task"].Load(ctx, dataloader.StringKey(*obj.TtID))()
		items, _ = item.(*Task)

		if items == nil {
			items = &Task{}
		}

	}

	return
}

func (r *GeneratedUserResolver) Ttt(ctx context.Context, obj *User) (res []*Task, err error) {
	return r.Handlers.UserTtt(ctx, r.GeneratedResolver, obj)
}
func UserTttHandler(ctx context.Context, r *GeneratedResolver, obj *User) (items []*Task, err error) {

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	item, _ := loaders["TaskUuu"].Load(ctx, dataloader.StringKey(obj.ID))()
	items = []*Task{}
	if item != nil {
		items = item.([]*Task)
	}

	return
}

func (r *GeneratedUserResolver) TttIds(ctx context.Context, obj *User) (ids []string, err error) {
	ids = []string{}
	items := []*Task{}

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	item, _ := loaders["TaskUuu"].Load(ctx, dataloader.StringKey(obj.ID))()

	if item != nil {
		items = item.([]*Task)
	}

	for _, v := range items {
		ids = append(ids, v.ID)
	}

	return
}

func (r *GeneratedUserResolver) Tttt(ctx context.Context, obj *User) (res []*Task, err error) {
	return r.Handlers.UserTttt(ctx, r.GeneratedResolver, obj)
}
func UserTtttHandler(ctx context.Context, r *GeneratedResolver, obj *User) (items []*Task, err error) {

	items = []*Task{}
	selects := GetFieldsRequested(ctx, strings.ToLower(TableName("tasks")))
	wheres := []string{}
	values := []interface{}{}
	// err = tx.Select(selects).Where(strings.Join(wheres, " AND "), values...).Model(obj).Related(&items, "Tttt").Error
	err = r.DB.Query().Select(selects).Where(strings.Join(wheres, " AND "), values...).Model(&Task{}).Preload("User").Find(&items).Error

	return
}

func (r *GeneratedUserResolver) TtttIds(ctx context.Context, obj *User) (ids []string, err error) {
	ids = []string{}
	items := []*Task{}

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	item, _ := loaders["TaskUuuu"].Load(ctx, dataloader.StringKey(obj.ID))()

	if item != nil {
		items = item.([]*Task)
	}

	for _, v := range items {
		ids = append(ids, v.ID)
	}

	return
}

type QueryTaskHandlerOptions struct {
	ID     *string
	Filter *TaskFilterType
}

func (r *GeneratedQueryResolver) Task(ctx context.Context, id *string, filter *TaskFilterType) (*Task, error) {
	opts := QueryTaskHandlerOptions{
		ID:     id,
		Filter: filter,
	}
	return r.Handlers.QueryTask(ctx, r.GeneratedResolver, opts)
}
func QueryTaskHandler(ctx context.Context, r *GeneratedResolver, opts QueryTaskHandlerOptions) (*Task, error) {
	selection := []ast.Selection{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		selection = append(selection, f.Field)
	}
	selectionSet := ast.SelectionSet(selection)

	query := TaskQueryFilter{}
	rt := &TaskResultType{
		EntityResultType: EntityResultType{
			Query:        &query,
			Filter:       opts.Filter,
			SelectionSet: &selectionSet,
		},
	}
	qb := r.DB.Query()
	if opts.ID != nil {
		qb = qb.Where(TableName("tasks")+".id = ?", *opts.ID)
	}

	var items []*Task
	giOpts := GetItemsOptions{
		Alias:      TableName("tasks"),
		Preloaders: []string{},
		Item:       &Task{},
	}
	err := rt.GetData(ctx, qb, giOpts, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, &NotFoundError{Entity: "Task"}
	}
	return items[0], err
}

type QueryTasksHandlerOptions struct {
	CurrentPage *int
	PerPage     *int
	Q           *string
	Sort        []*TaskSortType
	Filter      *TaskFilterType
	Rand        *bool
}

func (r *GeneratedQueryResolver) Tasks(ctx context.Context, current_page *int, per_page *int, q *string, sort []*TaskSortType, filter *TaskFilterType, rand *bool) (*TaskResultType, error) {
	opts := QueryTasksHandlerOptions{
		CurrentPage: current_page,
		PerPage:     per_page,
		Q:           q,
		Sort:        sort,
		Filter:      filter,
		Rand:        rand,
	}
	return r.Handlers.QueryTasks(ctx, r.GeneratedResolver, opts)
}
func QueryTasksHandler(ctx context.Context, r *GeneratedResolver, opts QueryTasksHandlerOptions) (*TaskResultType, error) {
	query := TaskQueryFilter{opts.Q}

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

	return &TaskResultType{
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

type GeneratedTaskResultTypeResolver struct{ *GeneratedResolver }

func (r *GeneratedTaskResultTypeResolver) Data(ctx context.Context, obj *TaskResultType) (items []*Task, err error) {
	giOpts := GetItemsOptions{
		Alias:      TableName("tasks"),
		Preloaders: []string{},
		Item:       &Task{},
	}
	err = obj.GetData(ctx, r.DB.db, giOpts, &items)

	uniqueItems := []*Task{}
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

func (r *GeneratedTaskResultTypeResolver) Total(ctx context.Context, obj *TaskResultType) (count int, err error) {
	return obj.GetTotal(ctx, r.DB.db, TableName("tasks"), &Task{})
}

func (r *GeneratedTaskResultTypeResolver) TotalPage(ctx context.Context, obj *TaskResultType) (count int, err error) {
	total, _ := r.Total(ctx, obj)
	perPage, _ := r.PerPage(ctx, obj)
	totalPage := int(math.Ceil(float64(total) / float64(perPage)))
	if totalPage < 0 {
		totalPage = 0
	}

	return totalPage, nil
}

func (r *GeneratedTaskResultTypeResolver) CurrentPage(ctx context.Context, obj *TaskResultType) (count int, err error) {
	return int(*obj.EntityResultType.CurrentPage), nil
}

func (r *GeneratedTaskResultTypeResolver) PerPage(ctx context.Context, obj *TaskResultType) (count int, err error) {
	return int(*obj.EntityResultType.PerPage), nil
}

type GeneratedTaskResolver struct{ *GeneratedResolver }

func (r *GeneratedTaskResolver) U(ctx context.Context, obj *Task) (res *User, err error) {
	return r.Handlers.TaskU(ctx, r.GeneratedResolver, obj)
}
func TaskUHandler(ctx context.Context, r *GeneratedResolver, obj *Task) (items *User, err error) {

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	if obj.UID != nil {
		item, _ := loaders["User"].Load(ctx, dataloader.StringKey(*obj.UID))()
		items, _ = item.(*User)

		if items == nil {
			items = &User{}
		}

	}

	return
}

func (r *GeneratedTaskResolver) Uu(ctx context.Context, obj *Task) (res []*User, err error) {
	return r.Handlers.TaskUu(ctx, r.GeneratedResolver, obj)
}
func TaskUuHandler(ctx context.Context, r *GeneratedResolver, obj *Task) (items []*User, err error) {

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	item, _ := loaders["UserTt"].Load(ctx, dataloader.StringKey(obj.ID))()
	items = []*User{}
	if item != nil {
		items = item.([]*User)
	}

	return
}

func (r *GeneratedTaskResolver) UuIds(ctx context.Context, obj *Task) (ids []string, err error) {
	ids = []string{}
	items := []*User{}

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	item, _ := loaders["UserTt"].Load(ctx, dataloader.StringKey(obj.ID))()

	if item != nil {
		items = item.([]*User)
	}

	for _, v := range items {
		ids = append(ids, v.ID)
	}

	return
}

func (r *GeneratedTaskResolver) Uuu(ctx context.Context, obj *Task) (res *User, err error) {
	return r.Handlers.TaskUuu(ctx, r.GeneratedResolver, obj)
}
func TaskUuuHandler(ctx context.Context, r *GeneratedResolver, obj *Task) (items *User, err error) {

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	if obj.UuuID != nil {
		item, _ := loaders["User"].Load(ctx, dataloader.StringKey(*obj.UuuID))()
		items, _ = item.(*User)

		if items == nil {
			items = &User{}
		}

	}

	return
}

func (r *GeneratedTaskResolver) Uuuu(ctx context.Context, obj *Task) (res []*User, err error) {
	return r.Handlers.TaskUuuu(ctx, r.GeneratedResolver, obj)
}
func TaskUuuuHandler(ctx context.Context, r *GeneratedResolver, obj *Task) (items []*User, err error) {

	items = []*User{}
	selects := GetFieldsRequested(ctx, strings.ToLower(TableName("users")))
	wheres := []string{}
	values := []interface{}{}
	// err = tx.Select(selects).Where(strings.Join(wheres, " AND "), values...).Model(obj).Related(&items, "Uuuu").Error
	err = r.DB.Query().Select(selects).Where(strings.Join(wheres, " AND "), values...).Model(&User{}).Preload("Task").Find(&items).Error

	return
}

func (r *GeneratedTaskResolver) UuuuIds(ctx context.Context, obj *Task) (ids []string, err error) {
	ids = []string{}
	items := []*User{}

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	item, _ := loaders["UserTttt"].Load(ctx, dataloader.StringKey(obj.ID))()

	if item != nil {
		items = item.([]*User)
	}

	for _, v := range items {
		ids = append(ids, v.ID)
	}

	return
}
