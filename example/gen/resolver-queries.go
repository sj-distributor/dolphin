package gen

import (
	"context"
	"errors"
	"math"

	"github.com/99designs/gqlgen/graphql"
	"github.com/graph-gophers/dataloader"
	"github.com/sj-distributor/dolphin-example/auth"
	"github.com/vektah/gqlparser/v2/ast"
	"gorm.io/gorm"
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
	func() {
		defer func() { recover() }()
		for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
			selection = append(selection, f.Field)
		}
	}()
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
		qb = qb.Where(TableName("users", ctx)+".id = ?", *opts.ID)
	}

	var items []*User
	giOpts := GetItemsOptions{
		Alias:      TableName("users", ctx),
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
	func() {
		defer func() { recover() }()
		for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
			if f.Field.Name == "data" {
				selectionSet = &f.Field.SelectionSet
			}
		}
	}()

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
		Alias:      TableName("users", ctx),
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
	return obj.GetTotal(ctx, r.DB.db, TableName("users", ctx), &User{})
}

func (r *GeneratedUserResultTypeResolver) TotalPage(ctx context.Context, obj *UserResultType) (count int, err error) {
	total, _ := r.Total(ctx, obj)
	perPage, _ := r.PerPage(ctx, obj)
	totalPage := int(math.Ceil(float64(total) / float64(perPage)))
	if totalPage < 0 {
		totalPage = 0
	} else if perPage <= 0 {
		totalPage = total
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

func (r *GeneratedUserResolver) Profile(ctx context.Context, obj *User) (res *Profile, err error) {
	return r.Handlers.UserProfile(ctx, r.GeneratedResolver, obj)
}
func UserProfileHandler(ctx context.Context, r *GeneratedResolver, obj *User) (items *Profile, err error) {

	// 判断是否有详情权限
	if err := auth.CheckAuthorization(ctx, "Profile"); err != nil {
		return items, errors.New("Profile " + err.Error())
	}

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	objKey := obj.ProfileID

	if objKey != nil {
		item, _ := loaders["Profile"].Load(ctx, dataloader.StringKey(*objKey))()

		items, _ = item.(*Profile)

	}

	return
}

func (r *GeneratedUserResolver) Tasks(ctx context.Context, obj *User) (res []*Task, err error) {
	return r.Handlers.UserTasks(ctx, r.GeneratedResolver, obj)
}
func UserTasksHandler(ctx context.Context, r *GeneratedResolver, obj *User) (items []*Task, err error) {

	items = []*Task{}

	// 判断是否有详情权限
	if err := auth.CheckAuthorization(ctx, "Tasks"); err != nil {
		return items, errors.New("Tasks " + err.Error())
	}

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	item, _ := loaders["TaskUser"].Load(ctx, dataloader.StringKey(obj.ID))()
	items = []*Task{}
	if item != nil {
		items = item.([]*Task)
	}

	return
}

func (r *GeneratedUserResolver) TasksIds(ctx context.Context, obj *User) (ids []string, err error) {

	items := []*Task{}

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	item, _ := loaders["UserAndTaskIds"].Load(ctx, dataloader.StringKey(obj.ID))()

	if item != nil {
		items = item.([]*Task)
	}

	for _, v := range items {
		ids = append(ids, v.ID)
	}

	return
}

func (r *GeneratedUserResolver) UserRoles(ctx context.Context, obj *User) (res []*UserRole, err error) {
	return r.Handlers.UserUserRoles(ctx, r.GeneratedResolver, obj)
}
func UserUserRolesHandler(ctx context.Context, r *GeneratedResolver, obj *User) (items []*UserRole, err error) {

	items = []*UserRole{}

	// 判断是否有详情权限
	if err := auth.CheckAuthorization(ctx, "UserRoles"); err != nil {
		return items, errors.New("UserRoles " + err.Error())
	}

	// selects := GetFieldsRequested(ctx, strings.ToLower(TableName("user_roles", ctx)))
	// wheres  := []string{}
	// values  := []interface{}{}
	// err = tx.Select(selects).Where(strings.Join(wheres, " AND "), values...).Model(obj).Related(&items, "UserRoles").Error
	// err = r.DB.Query().Select(selects).Where(strings.Join(wheres, " AND "), values...).Model(&UserRole{}).Find(&items).Error

	err = r.DB.Query().Model(obj).Order("weight ASC, created_at ASC").Preload("UserRoles").First(&obj).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 记录不存在
			return items, nil
		}
		return items, err
	}

	items = obj.UserRoles

	return
}

func (r *GeneratedUserResolver) UserRolesIds(ctx context.Context, obj *User) (ids []string, err error) {

	err = r.DB.Query().Order("weight ASC, created_at ASC").Preload("UserRoles").First(&obj).Error
	if err != nil {
		return
	}

	for _, item := range obj.UserRoles {
		ids = append(ids, item.ID)
	}

	return
}

type QueryProfileHandlerOptions struct {
	ID     *string
	Filter *ProfileFilterType
}

func (r *GeneratedQueryResolver) Profile(ctx context.Context, id *string, filter *ProfileFilterType) (*Profile, error) {
	opts := QueryProfileHandlerOptions{
		ID:     id,
		Filter: filter,
	}
	return r.Handlers.QueryProfile(ctx, r.GeneratedResolver, opts)
}
func QueryProfileHandler(ctx context.Context, r *GeneratedResolver, opts QueryProfileHandlerOptions) (*Profile, error) {
	selection := []ast.Selection{}
	func() {
		defer func() { recover() }()
		for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
			selection = append(selection, f.Field)
		}
	}()
	selectionSet := ast.SelectionSet(selection)

	query := ProfileQueryFilter{}
	rt := &ProfileResultType{
		EntityResultType: EntityResultType{
			Query:        &query,
			Filter:       opts.Filter,
			SelectionSet: &selectionSet,
		},
	}
	qb := r.DB.Query()
	if opts.ID != nil {
		qb = qb.Where(TableName("profiles", ctx)+".id = ?", *opts.ID)
	}

	var items []*Profile
	giOpts := GetItemsOptions{
		Alias:      TableName("profiles", ctx),
		Preloaders: []string{},
		Item:       &Profile{},
	}
	err := rt.GetData(ctx, qb, giOpts, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, &NotFoundError{Entity: "Profile"}
	}
	return items[0], err
}

type QueryProfilesHandlerOptions struct {
	CurrentPage *int
	PerPage     *int
	Q           *string
	Sort        []*ProfileSortType
	Filter      *ProfileFilterType
	Rand        *bool
}

func (r *GeneratedQueryResolver) Profiles(ctx context.Context, current_page *int, per_page *int, q *string, sort []*ProfileSortType, filter *ProfileFilterType, rand *bool) (*ProfileResultType, error) {
	opts := QueryProfilesHandlerOptions{
		CurrentPage: current_page,
		PerPage:     per_page,
		Q:           q,
		Sort:        sort,
		Filter:      filter,
		Rand:        rand,
	}
	return r.Handlers.QueryProfiles(ctx, r.GeneratedResolver, opts)
}
func QueryProfilesHandler(ctx context.Context, r *GeneratedResolver, opts QueryProfilesHandlerOptions) (*ProfileResultType, error) {
	query := ProfileQueryFilter{opts.Q}

	var selectionSet *ast.SelectionSet
	func() {
		defer func() { recover() }()
		for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
			if f.Field.Name == "data" {
				selectionSet = &f.Field.SelectionSet
			}
		}
	}()

	_sort := []EntitySort{}
	for _, sort := range opts.Sort {
		_sort = append(_sort, sort)
	}

	return &ProfileResultType{
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

type GeneratedProfileResultTypeResolver struct{ *GeneratedResolver }

func (r *GeneratedProfileResultTypeResolver) Data(ctx context.Context, obj *ProfileResultType) (items []*Profile, err error) {
	giOpts := GetItemsOptions{
		Alias:      TableName("profiles", ctx),
		Preloaders: []string{},
		Item:       &Profile{},
	}
	err = obj.GetData(ctx, r.DB.db, giOpts, &items)

	uniqueItems := []*Profile{}
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

func (r *GeneratedProfileResultTypeResolver) Total(ctx context.Context, obj *ProfileResultType) (count int, err error) {
	return obj.GetTotal(ctx, r.DB.db, TableName("profiles", ctx), &Profile{})
}

func (r *GeneratedProfileResultTypeResolver) TotalPage(ctx context.Context, obj *ProfileResultType) (count int, err error) {
	total, _ := r.Total(ctx, obj)
	perPage, _ := r.PerPage(ctx, obj)
	totalPage := int(math.Ceil(float64(total) / float64(perPage)))
	if totalPage < 0 {
		totalPage = 0
	} else if perPage <= 0 {
		totalPage = total
	}

	return totalPage, nil
}

func (r *GeneratedProfileResultTypeResolver) CurrentPage(ctx context.Context, obj *ProfileResultType) (count int, err error) {
	return int(*obj.EntityResultType.CurrentPage), nil
}

func (r *GeneratedProfileResultTypeResolver) PerPage(ctx context.Context, obj *ProfileResultType) (count int, err error) {
	return int(*obj.EntityResultType.PerPage), nil
}

type GeneratedProfileResolver struct{ *GeneratedResolver }

func (r *GeneratedProfileResolver) User(ctx context.Context, obj *Profile) (res *User, err error) {
	return r.Handlers.ProfileUser(ctx, r.GeneratedResolver, obj)
}
func ProfileUserHandler(ctx context.Context, r *GeneratedResolver, obj *Profile) (items *User, err error) {

	// 判断是否有详情权限
	if err := auth.CheckAuthorization(ctx, "User"); err != nil {
		return items, errors.New("User " + err.Error())
	}

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	objKey := obj.UserID

	if objKey != "" {
		item, _ := loaders["User"].Load(ctx, dataloader.StringKey(objKey))()

		items, _ = item.(*User)

		if items == nil {
			items = &User{}
		}

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
	func() {
		defer func() { recover() }()
		for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
			selection = append(selection, f.Field)
		}
	}()
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
		qb = qb.Where(TableName("tasks", ctx)+".id = ?", *opts.ID)
	}

	var items []*Task
	giOpts := GetItemsOptions{
		Alias:      TableName("tasks", ctx),
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
	func() {
		defer func() { recover() }()
		for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
			if f.Field.Name == "data" {
				selectionSet = &f.Field.SelectionSet
			}
		}
	}()

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
		Alias:      TableName("tasks", ctx),
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
	return obj.GetTotal(ctx, r.DB.db, TableName("tasks", ctx), &Task{})
}

func (r *GeneratedTaskResultTypeResolver) TotalPage(ctx context.Context, obj *TaskResultType) (count int, err error) {
	total, _ := r.Total(ctx, obj)
	perPage, _ := r.PerPage(ctx, obj)
	totalPage := int(math.Ceil(float64(total) / float64(perPage)))
	if totalPage < 0 {
		totalPage = 0
	} else if perPage <= 0 {
		totalPage = total
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

func (r *GeneratedTaskResolver) User(ctx context.Context, obj *Task) (res *User, err error) {
	return r.Handlers.TaskUser(ctx, r.GeneratedResolver, obj)
}
func TaskUserHandler(ctx context.Context, r *GeneratedResolver, obj *Task) (items *User, err error) {

	// 判断是否有详情权限
	if err := auth.CheckAuthorization(ctx, "User"); err != nil {
		return items, errors.New("User " + err.Error())
	}

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	objKey := obj.UserID

	if objKey != nil {
		item, _ := loaders["User"].Load(ctx, dataloader.StringKey(*objKey))()

		items, _ = item.(*User)

	}

	return
}

func (r *GeneratedTaskResolver) Tags(ctx context.Context, obj *Task) (res []*Tag, err error) {
	return r.Handlers.TaskTags(ctx, r.GeneratedResolver, obj)
}
func TaskTagsHandler(ctx context.Context, r *GeneratedResolver, obj *Task) (items []*Tag, err error) {

	items = []*Tag{}

	// 判断是否有详情权限
	if err := auth.CheckAuthorization(ctx, "Tags"); err != nil {
		return items, errors.New("Tags " + err.Error())
	}

	// selects := GetFieldsRequested(ctx, strings.ToLower(TableName("tags", ctx)))
	// wheres  := []string{}
	// values  := []interface{}{}
	// err = tx.Select(selects).Where(strings.Join(wheres, " AND "), values...).Model(obj).Related(&items, "Tags").Error
	// err = r.DB.Query().Select(selects).Where(strings.Join(wheres, " AND "), values...).Model(&Tag{}).Find(&items).Error

	err = r.DB.Query().Model(obj).Order("weight ASC, created_at ASC").Preload("Tags").First(&obj).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 记录不存在
			return items, nil
		}
		return items, err
	}

	items = obj.Tags

	return
}

func (r *GeneratedTaskResolver) TagsIds(ctx context.Context, obj *Task) (ids []string, err error) {

	err = r.DB.Query().Order("weight ASC, created_at ASC").Preload("Tags").First(&obj).Error
	if err != nil {
		return
	}

	for _, item := range obj.Tags {
		ids = append(ids, item.ID)
	}

	return
}

type QueryUserRoleHandlerOptions struct {
	ID     *string
	Filter *UserRoleFilterType
}

func (r *GeneratedQueryResolver) UserRole(ctx context.Context, id *string, filter *UserRoleFilterType) (*UserRole, error) {
	opts := QueryUserRoleHandlerOptions{
		ID:     id,
		Filter: filter,
	}
	return r.Handlers.QueryUserRole(ctx, r.GeneratedResolver, opts)
}
func QueryUserRoleHandler(ctx context.Context, r *GeneratedResolver, opts QueryUserRoleHandlerOptions) (*UserRole, error) {
	selection := []ast.Selection{}
	func() {
		defer func() { recover() }()
		for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
			selection = append(selection, f.Field)
		}
	}()
	selectionSet := ast.SelectionSet(selection)

	query := UserRoleQueryFilter{}
	rt := &UserRoleResultType{
		EntityResultType: EntityResultType{
			Query:        &query,
			Filter:       opts.Filter,
			SelectionSet: &selectionSet,
		},
	}
	qb := r.DB.Query()
	if opts.ID != nil {
		qb = qb.Where(TableName("user_roles", ctx)+".id = ?", *opts.ID)
	}

	var items []*UserRole
	giOpts := GetItemsOptions{
		Alias:      TableName("user_roles", ctx),
		Preloaders: []string{},
		Item:       &UserRole{},
	}
	err := rt.GetData(ctx, qb, giOpts, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, &NotFoundError{Entity: "UserRole"}
	}
	return items[0], err
}

type QueryUserRolesHandlerOptions struct {
	CurrentPage *int
	PerPage     *int
	Q           *string
	Sort        []*UserRoleSortType
	Filter      *UserRoleFilterType
	Rand        *bool
}

func (r *GeneratedQueryResolver) UserRoles(ctx context.Context, current_page *int, per_page *int, q *string, sort []*UserRoleSortType, filter *UserRoleFilterType, rand *bool) (*UserRoleResultType, error) {
	opts := QueryUserRolesHandlerOptions{
		CurrentPage: current_page,
		PerPage:     per_page,
		Q:           q,
		Sort:        sort,
		Filter:      filter,
		Rand:        rand,
	}
	return r.Handlers.QueryUserRoles(ctx, r.GeneratedResolver, opts)
}
func QueryUserRolesHandler(ctx context.Context, r *GeneratedResolver, opts QueryUserRolesHandlerOptions) (*UserRoleResultType, error) {
	query := UserRoleQueryFilter{opts.Q}

	var selectionSet *ast.SelectionSet
	func() {
		defer func() { recover() }()
		for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
			if f.Field.Name == "data" {
				selectionSet = &f.Field.SelectionSet
			}
		}
	}()

	_sort := []EntitySort{}
	for _, sort := range opts.Sort {
		_sort = append(_sort, sort)
	}

	return &UserRoleResultType{
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

type GeneratedUserRoleResultTypeResolver struct{ *GeneratedResolver }

func (r *GeneratedUserRoleResultTypeResolver) Data(ctx context.Context, obj *UserRoleResultType) (items []*UserRole, err error) {
	giOpts := GetItemsOptions{
		Alias:      TableName("user_roles", ctx),
		Preloaders: []string{},
		Item:       &UserRole{},
	}
	err = obj.GetData(ctx, r.DB.db, giOpts, &items)

	uniqueItems := []*UserRole{}
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

func (r *GeneratedUserRoleResultTypeResolver) Total(ctx context.Context, obj *UserRoleResultType) (count int, err error) {
	return obj.GetTotal(ctx, r.DB.db, TableName("user_roles", ctx), &UserRole{})
}

func (r *GeneratedUserRoleResultTypeResolver) TotalPage(ctx context.Context, obj *UserRoleResultType) (count int, err error) {
	total, _ := r.Total(ctx, obj)
	perPage, _ := r.PerPage(ctx, obj)
	totalPage := int(math.Ceil(float64(total) / float64(perPage)))
	if totalPage < 0 {
		totalPage = 0
	} else if perPage <= 0 {
		totalPage = total
	}

	return totalPage, nil
}

func (r *GeneratedUserRoleResultTypeResolver) CurrentPage(ctx context.Context, obj *UserRoleResultType) (count int, err error) {
	return int(*obj.EntityResultType.CurrentPage), nil
}

func (r *GeneratedUserRoleResultTypeResolver) PerPage(ctx context.Context, obj *UserRoleResultType) (count int, err error) {
	return int(*obj.EntityResultType.PerPage), nil
}

type GeneratedUserRoleResolver struct{ *GeneratedResolver }

func (r *GeneratedUserRoleResolver) Users(ctx context.Context, obj *UserRole) (res []*User, err error) {
	return r.Handlers.UserRoleUsers(ctx, r.GeneratedResolver, obj)
}
func UserRoleUsersHandler(ctx context.Context, r *GeneratedResolver, obj *UserRole) (items []*User, err error) {

	items = []*User{}

	// 判断是否有详情权限
	if err := auth.CheckAuthorization(ctx, "Users"); err != nil {
		return items, errors.New("Users " + err.Error())
	}

	// selects := GetFieldsRequested(ctx, strings.ToLower(TableName("users", ctx)))
	// wheres  := []string{}
	// values  := []interface{}{}
	// err = tx.Select(selects).Where(strings.Join(wheres, " AND "), values...).Model(obj).Related(&items, "Users").Error
	// err = r.DB.Query().Select(selects).Where(strings.Join(wheres, " AND "), values...).Model(&User{}).Find(&items).Error

	err = r.DB.Query().Model(obj).Order("weight ASC, created_at ASC").Preload("Users").First(&obj).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 记录不存在
			return items, nil
		}
		return items, err
	}

	items = obj.Users

	return
}

func (r *GeneratedUserRoleResolver) UsersIds(ctx context.Context, obj *UserRole) (ids []string, err error) {

	err = r.DB.Query().Order("weight ASC, created_at ASC").Preload("Users").First(&obj).Error
	if err != nil {
		return
	}

	for _, item := range obj.Users {
		ids = append(ids, item.ID)
	}

	return
}

type QueryTagHandlerOptions struct {
	ID     *string
	Filter *TagFilterType
}

func (r *GeneratedQueryResolver) Tag(ctx context.Context, id *string, filter *TagFilterType) (*Tag, error) {
	opts := QueryTagHandlerOptions{
		ID:     id,
		Filter: filter,
	}
	return r.Handlers.QueryTag(ctx, r.GeneratedResolver, opts)
}
func QueryTagHandler(ctx context.Context, r *GeneratedResolver, opts QueryTagHandlerOptions) (*Tag, error) {
	selection := []ast.Selection{}
	func() {
		defer func() { recover() }()
		for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
			selection = append(selection, f.Field)
		}
	}()
	selectionSet := ast.SelectionSet(selection)

	query := TagQueryFilter{}
	rt := &TagResultType{
		EntityResultType: EntityResultType{
			Query:        &query,
			Filter:       opts.Filter,
			SelectionSet: &selectionSet,
		},
	}
	qb := r.DB.Query()
	if opts.ID != nil {
		qb = qb.Where(TableName("tags", ctx)+".id = ?", *opts.ID)
	}

	var items []*Tag
	giOpts := GetItemsOptions{
		Alias:      TableName("tags", ctx),
		Preloaders: []string{},
		Item:       &Tag{},
	}
	err := rt.GetData(ctx, qb, giOpts, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, &NotFoundError{Entity: "Tag"}
	}
	return items[0], err
}

type QueryTagsHandlerOptions struct {
	CurrentPage *int
	PerPage     *int
	Q           *string
	Sort        []*TagSortType
	Filter      *TagFilterType
	Rand        *bool
}

func (r *GeneratedQueryResolver) Tags(ctx context.Context, current_page *int, per_page *int, q *string, sort []*TagSortType, filter *TagFilterType, rand *bool) (*TagResultType, error) {
	opts := QueryTagsHandlerOptions{
		CurrentPage: current_page,
		PerPage:     per_page,
		Q:           q,
		Sort:        sort,
		Filter:      filter,
		Rand:        rand,
	}
	return r.Handlers.QueryTags(ctx, r.GeneratedResolver, opts)
}
func QueryTagsHandler(ctx context.Context, r *GeneratedResolver, opts QueryTagsHandlerOptions) (*TagResultType, error) {
	query := TagQueryFilter{opts.Q}

	var selectionSet *ast.SelectionSet
	func() {
		defer func() { recover() }()
		for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
			if f.Field.Name == "data" {
				selectionSet = &f.Field.SelectionSet
			}
		}
	}()

	_sort := []EntitySort{}
	for _, sort := range opts.Sort {
		_sort = append(_sort, sort)
	}

	return &TagResultType{
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

type GeneratedTagResultTypeResolver struct{ *GeneratedResolver }

func (r *GeneratedTagResultTypeResolver) Data(ctx context.Context, obj *TagResultType) (items []*Tag, err error) {
	giOpts := GetItemsOptions{
		Alias:      TableName("tags", ctx),
		Preloaders: []string{},
		Item:       &Tag{},
	}
	err = obj.GetData(ctx, r.DB.db, giOpts, &items)

	uniqueItems := []*Tag{}
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

func (r *GeneratedTagResultTypeResolver) Total(ctx context.Context, obj *TagResultType) (count int, err error) {
	return obj.GetTotal(ctx, r.DB.db, TableName("tags", ctx), &Tag{})
}

func (r *GeneratedTagResultTypeResolver) TotalPage(ctx context.Context, obj *TagResultType) (count int, err error) {
	total, _ := r.Total(ctx, obj)
	perPage, _ := r.PerPage(ctx, obj)
	totalPage := int(math.Ceil(float64(total) / float64(perPage)))
	if totalPage < 0 {
		totalPage = 0
	} else if perPage <= 0 {
		totalPage = total
	}

	return totalPage, nil
}

func (r *GeneratedTagResultTypeResolver) CurrentPage(ctx context.Context, obj *TagResultType) (count int, err error) {
	return int(*obj.EntityResultType.CurrentPage), nil
}

func (r *GeneratedTagResultTypeResolver) PerPage(ctx context.Context, obj *TagResultType) (count int, err error) {
	return int(*obj.EntityResultType.PerPage), nil
}

type GeneratedTagResolver struct{ *GeneratedResolver }

func (r *GeneratedTagResolver) Tasks(ctx context.Context, obj *Tag) (res []*Task, err error) {
	return r.Handlers.TagTasks(ctx, r.GeneratedResolver, obj)
}
func TagTasksHandler(ctx context.Context, r *GeneratedResolver, obj *Tag) (items []*Task, err error) {

	items = []*Task{}

	// 判断是否有详情权限
	if err := auth.CheckAuthorization(ctx, "Tasks"); err != nil {
		return items, errors.New("Tasks " + err.Error())
	}

	// selects := GetFieldsRequested(ctx, strings.ToLower(TableName("tasks", ctx)))
	// wheres  := []string{}
	// values  := []interface{}{}
	// err = tx.Select(selects).Where(strings.Join(wheres, " AND "), values...).Model(obj).Related(&items, "Tasks").Error
	// err = r.DB.Query().Select(selects).Where(strings.Join(wheres, " AND "), values...).Model(&Task{}).Find(&items).Error

	err = r.DB.Query().Model(obj).Order("weight ASC, created_at ASC").Preload("Tasks").First(&obj).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 记录不存在
			return items, nil
		}
		return items, err
	}

	items = obj.Tasks

	return
}

func (r *GeneratedTagResolver) TasksIds(ctx context.Context, obj *Tag) (ids []string, err error) {

	err = r.DB.Query().Order("weight ASC, created_at ASC").Preload("Tasks").First(&obj).Error
	if err != nil {
		return
	}

	for _, item := range obj.Tasks {
		ids = append(ids, item.ID)
	}

	return
}
