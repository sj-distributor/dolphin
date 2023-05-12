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

func (r *GeneratedUserResolver) Tasks(ctx context.Context, obj *User) (res []*Task, err error) {
	return r.Handlers.UserTasks(ctx, r.GeneratedResolver, obj)
}
func UserTasksHandler(ctx context.Context, r *GeneratedResolver, obj *User) (items []*Task, err error) {
	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)

	item, _ := loaders["TaskAssignee"].Load(ctx, dataloader.StringKey(obj.ID))()
	items = []*Task{}
	if item != nil {
		items = item.([]*Task)
	}

	return
}

func (r *GeneratedUserResolver) TasksIds(ctx context.Context, obj *User) (ids []string, err error) {
	ids = []string{}
	items := []*Task{}

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	item, _ := loaders["TaskAssignee"].Load(ctx, dataloader.StringKey(obj.ID))()

	if item != nil {
		items = item.([]*Task)
	}

	for _, v := range items {
		ids = append(ids, v.ID)
	}

	return
}

func (r *GeneratedQueryResolver) Task(ctx context.Context, id string) (*Task, error) {
	return r.Handlers.QueryTask(ctx, r.GeneratedResolver, id)
}
func QueryTaskHandler(ctx context.Context, r *GeneratedResolver, id string) (*Task, error) {
	selection := []ast.Selection{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		selection = append(selection, f.Field)
	}
	selectionSet := ast.SelectionSet(selection)

	query := TaskQueryFilter{}
	rt := &TaskResultType{
		EntityResultType: EntityResultType{
			Query:        &query,
			SelectionSet: &selectionSet,
		},
	}
	qb := r.DB.Query()
	qb = qb.Where(TableName("tasks")+".id = ?", id)

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

func (r *GeneratedTaskResolver) Assignee(ctx context.Context, obj *Task) (res *User, err error) {
	return r.Handlers.TaskAssignee(ctx, r.GeneratedResolver, obj)
}
func TaskAssigneeHandler(ctx context.Context, r *GeneratedResolver, obj *Task) (items *User, err error) {
	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)

	if obj.AssigneeID != nil {
		item, _ := loaders["User"].Load(ctx, dataloader.StringKey(*obj.AssigneeID))()
		items, _ = item.(*User)

	}

	return
}

func (r *GeneratedQueryResolver) UploadFile(ctx context.Context, id string) (*UploadFile, error) {
	return r.Handlers.QueryUploadFile(ctx, r.GeneratedResolver, id)
}
func QueryUploadFileHandler(ctx context.Context, r *GeneratedResolver, id string) (*UploadFile, error) {
	selection := []ast.Selection{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		selection = append(selection, f.Field)
	}
	selectionSet := ast.SelectionSet(selection)

	query := UploadFileQueryFilter{}
	rt := &UploadFileResultType{
		EntityResultType: EntityResultType{
			Query:        &query,
			SelectionSet: &selectionSet,
		},
	}
	qb := r.DB.Query()
	qb = qb.Where(TableName("upload_files")+".id = ?", id)

	var items []*UploadFile
	giOpts := GetItemsOptions{
		Alias:      TableName("upload_files"),
		Preloaders: []string{},
		Item:       &UploadFile{},
	}
	err := rt.GetData(ctx, qb, giOpts, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, &NotFoundError{Entity: "UploadFile"}
	}
	return items[0], err
}

type QueryUploadFilesHandlerOptions struct {
	CurrentPage *int
	PerPage     *int
	Q           *string
	Sort        []*UploadFileSortType
	Filter      *UploadFileFilterType
	Rand        *bool
}

func (r *GeneratedQueryResolver) UploadFiles(ctx context.Context, current_page *int, per_page *int, q *string, sort []*UploadFileSortType, filter *UploadFileFilterType, rand *bool) (*UploadFileResultType, error) {
	opts := QueryUploadFilesHandlerOptions{
		CurrentPage: current_page,
		PerPage:     per_page,
		Q:           q,
		Sort:        sort,
		Filter:      filter,
		Rand:        rand,
	}
	return r.Handlers.QueryUploadFiles(ctx, r.GeneratedResolver, opts)
}
func QueryUploadFilesHandler(ctx context.Context, r *GeneratedResolver, opts QueryUploadFilesHandlerOptions) (*UploadFileResultType, error) {
	query := UploadFileQueryFilter{opts.Q}

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

	return &UploadFileResultType{
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

type GeneratedUploadFileResultTypeResolver struct{ *GeneratedResolver }

func (r *GeneratedUploadFileResultTypeResolver) Data(ctx context.Context, obj *UploadFileResultType) (items []*UploadFile, err error) {
	giOpts := GetItemsOptions{
		Alias:      TableName("upload_files"),
		Preloaders: []string{},
		Item:       &UploadFile{},
	}
	err = obj.GetData(ctx, r.DB.db, giOpts, &items)

	uniqueItems := []*UploadFile{}
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

func (r *GeneratedUploadFileResultTypeResolver) Total(ctx context.Context, obj *UploadFileResultType) (count int, err error) {
	return obj.GetTotal(ctx, r.DB.db, TableName("upload_files"), &UploadFile{})
}

func (r *GeneratedUploadFileResultTypeResolver) TotalPage(ctx context.Context, obj *UploadFileResultType) (count int, err error) {
	total, _ := r.Total(ctx, obj)
	perPage, _ := r.PerPage(ctx, obj)
	totalPage := int(math.Ceil(float64(total) / float64(perPage)))
	if totalPage < 0 {
		totalPage = 0
	}

	return totalPage, nil
}

func (r *GeneratedUploadFileResultTypeResolver) CurrentPage(ctx context.Context, obj *UploadFileResultType) (count int, err error) {
	return int(*obj.EntityResultType.CurrentPage), nil
}

func (r *GeneratedUploadFileResultTypeResolver) PerPage(ctx context.Context, obj *UploadFileResultType) (count int, err error) {
	return int(*obj.EntityResultType.PerPage), nil
}
