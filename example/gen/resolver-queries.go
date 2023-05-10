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

func (r *GeneratedUserResolver) Accounts(ctx context.Context, obj *User) (res []*Account, err error) {
	return r.Handlers.UserAccounts(ctx, r.GeneratedResolver, obj)
}
func UserAccountsHandler(ctx context.Context, r *GeneratedResolver, obj *User) (items []*Account, err error) {
	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)

	item, _ := loaders["AccountOwner"].Load(ctx, dataloader.StringKey(obj.ID))()
	items = []*Account{}
	if item != nil {
		items = item.([]*Account)
	}

	return
}

func (r *GeneratedUserResolver) AccountsIds(ctx context.Context, obj *User) (ids []string, err error) {
	ids = []string{}
	items := []*Account{}

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	item, _ := loaders["AccountOwner"].Load(ctx, dataloader.StringKey(obj.ID))()

	if item != nil {
		items = item.([]*Account)
	}

	for _, v := range items {
		ids = append(ids, v.ID)
	}

	return
}

func (r *GeneratedUserResolver) Todo(ctx context.Context, obj *User) (res []*Todo, err error) {
	return r.Handlers.UserTodo(ctx, r.GeneratedResolver, obj)
}
func UserTodoHandler(ctx context.Context, r *GeneratedResolver, obj *User) (items []*Todo, err error) {
	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)

	item, _ := loaders["TodoAccount"].Load(ctx, dataloader.StringKey(obj.ID))()
	items = []*Todo{}
	if item != nil {
		items = item.([]*Todo)
	}

	return
}

func (r *GeneratedUserResolver) TodoIds(ctx context.Context, obj *User) (ids []string, err error) {
	ids = []string{}
	items := []*Todo{}

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	item, _ := loaders["TodoAccount"].Load(ctx, dataloader.StringKey(obj.ID))()

	if item != nil {
		items = item.([]*Todo)
	}

	for _, v := range items {
		ids = append(ids, v.ID)
	}

	return
}

func (r *GeneratedQueryResolver) Account(ctx context.Context, id string) (*Account, error) {
	return r.Handlers.QueryAccount(ctx, r.GeneratedResolver, id)
}
func QueryAccountHandler(ctx context.Context, r *GeneratedResolver, id string) (*Account, error) {
	selection := []ast.Selection{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		selection = append(selection, f.Field)
	}
	selectionSet := ast.SelectionSet(selection)

	query := AccountQueryFilter{}
	rt := &AccountResultType{
		EntityResultType: EntityResultType{
			Query:        &query,
			SelectionSet: &selectionSet,
		},
	}
	qb := r.DB.Query()
	qb = qb.Where(TableName("accounts")+".id = ?", id)

	var items []*Account
	giOpts := GetItemsOptions{
		Alias:      TableName("accounts"),
		Preloaders: []string{},
		Item:       &Account{},
	}
	err := rt.GetData(ctx, qb, giOpts, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, &NotFoundError{Entity: "Account"}
	}
	return items[0], err
}

type QueryAccountsHandlerOptions struct {
	CurrentPage *int
	PerPage     *int
	Q           *string
	Sort        []*AccountSortType
	Filter      *AccountFilterType
	Rand        *bool
}

func (r *GeneratedQueryResolver) Accounts(ctx context.Context, current_page *int, per_page *int, q *string, sort []*AccountSortType, filter *AccountFilterType, rand *bool) (*AccountResultType, error) {
	opts := QueryAccountsHandlerOptions{
		CurrentPage: current_page,
		PerPage:     per_page,
		Q:           q,
		Sort:        sort,
		Filter:      filter,
		Rand:        rand,
	}
	return r.Handlers.QueryAccounts(ctx, r.GeneratedResolver, opts)
}
func QueryAccountsHandler(ctx context.Context, r *GeneratedResolver, opts QueryAccountsHandlerOptions) (*AccountResultType, error) {
	query := AccountQueryFilter{opts.Q}

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

	return &AccountResultType{
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

type GeneratedAccountResultTypeResolver struct{ *GeneratedResolver }

func (r *GeneratedAccountResultTypeResolver) Data(ctx context.Context, obj *AccountResultType) (items []*Account, err error) {
	giOpts := GetItemsOptions{
		Alias:      TableName("accounts"),
		Preloaders: []string{},
		Item:       &Account{},
	}
	err = obj.GetData(ctx, r.DB.db, giOpts, &items)

	uniqueItems := []*Account{}
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

func (r *GeneratedAccountResultTypeResolver) Total(ctx context.Context, obj *AccountResultType) (count int, err error) {
	return obj.GetTotal(ctx, r.DB.db, TableName("accounts"), &Account{})
}

func (r *GeneratedAccountResultTypeResolver) TotalPage(ctx context.Context, obj *AccountResultType) (count int, err error) {
	total, _ := r.Total(ctx, obj)
	perPage, _ := r.PerPage(ctx, obj)
	totalPage := int(math.Ceil(float64(total) / float64(perPage)))
	if totalPage < 0 {
		totalPage = 0
	}

	return totalPage, nil
}

func (r *GeneratedAccountResultTypeResolver) CurrentPage(ctx context.Context, obj *AccountResultType) (count int, err error) {
	return int(*obj.EntityResultType.CurrentPage), nil
}

func (r *GeneratedAccountResultTypeResolver) PerPage(ctx context.Context, obj *AccountResultType) (count int, err error) {
	return int(*obj.EntityResultType.PerPage), nil
}

type GeneratedAccountResolver struct{ *GeneratedResolver }

func (r *GeneratedAccountResolver) Owner(ctx context.Context, obj *Account) (res *User, err error) {
	return r.Handlers.AccountOwner(ctx, r.GeneratedResolver, obj)
}
func AccountOwnerHandler(ctx context.Context, r *GeneratedResolver, obj *Account) (items *User, err error) {
	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)

	if obj.OwnerID != nil {
		item, _ := loaders["User"].Load(ctx, dataloader.StringKey(*obj.OwnerID))()
		items, _ = item.(*User)

		if items == nil {
			items = &User{}
		}

	}

	return
}

func (r *GeneratedAccountResolver) Transactions(ctx context.Context, obj *Account) (res []*Transaction, err error) {
	return r.Handlers.AccountTransactions(ctx, r.GeneratedResolver, obj)
}
func AccountTransactionsHandler(ctx context.Context, r *GeneratedResolver, obj *Account) (items []*Transaction, err error) {
	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)

	item, _ := loaders["TransactionAccount"].Load(ctx, dataloader.StringKey(obj.ID))()
	items = []*Transaction{}
	if item != nil {
		items = item.([]*Transaction)
	}

	return
}

func (r *GeneratedAccountResolver) TransactionsIds(ctx context.Context, obj *Account) (ids []string, err error) {
	ids = []string{}
	items := []*Transaction{}

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	item, _ := loaders["TransactionAccount"].Load(ctx, dataloader.StringKey(obj.ID))()

	if item != nil {
		items = item.([]*Transaction)
	}

	for _, v := range items {
		ids = append(ids, v.ID)
	}

	return
}

func (r *GeneratedQueryResolver) Transaction(ctx context.Context, id string) (*Transaction, error) {
	return r.Handlers.QueryTransaction(ctx, r.GeneratedResolver, id)
}
func QueryTransactionHandler(ctx context.Context, r *GeneratedResolver, id string) (*Transaction, error) {
	selection := []ast.Selection{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		selection = append(selection, f.Field)
	}
	selectionSet := ast.SelectionSet(selection)

	query := TransactionQueryFilter{}
	rt := &TransactionResultType{
		EntityResultType: EntityResultType{
			Query:        &query,
			SelectionSet: &selectionSet,
		},
	}
	qb := r.DB.Query()
	qb = qb.Where(TableName("transactions")+".id = ?", id)

	var items []*Transaction
	giOpts := GetItemsOptions{
		Alias:      TableName("transactions"),
		Preloaders: []string{},
		Item:       &Transaction{},
	}
	err := rt.GetData(ctx, qb, giOpts, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, &NotFoundError{Entity: "Transaction"}
	}
	return items[0], err
}

type QueryTransactionsHandlerOptions struct {
	CurrentPage *int
	PerPage     *int
	Q           *string
	Sort        []*TransactionSortType
	Filter      *TransactionFilterType
	Rand        *bool
}

func (r *GeneratedQueryResolver) Transactions(ctx context.Context, current_page *int, per_page *int, q *string, sort []*TransactionSortType, filter *TransactionFilterType, rand *bool) (*TransactionResultType, error) {
	opts := QueryTransactionsHandlerOptions{
		CurrentPage: current_page,
		PerPage:     per_page,
		Q:           q,
		Sort:        sort,
		Filter:      filter,
		Rand:        rand,
	}
	return r.Handlers.QueryTransactions(ctx, r.GeneratedResolver, opts)
}
func QueryTransactionsHandler(ctx context.Context, r *GeneratedResolver, opts QueryTransactionsHandlerOptions) (*TransactionResultType, error) {
	query := TransactionQueryFilter{opts.Q}

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

	return &TransactionResultType{
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

type GeneratedTransactionResultTypeResolver struct{ *GeneratedResolver }

func (r *GeneratedTransactionResultTypeResolver) Data(ctx context.Context, obj *TransactionResultType) (items []*Transaction, err error) {
	giOpts := GetItemsOptions{
		Alias:      TableName("transactions"),
		Preloaders: []string{},
		Item:       &Transaction{},
	}
	err = obj.GetData(ctx, r.DB.db, giOpts, &items)

	uniqueItems := []*Transaction{}
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

func (r *GeneratedTransactionResultTypeResolver) Total(ctx context.Context, obj *TransactionResultType) (count int, err error) {
	return obj.GetTotal(ctx, r.DB.db, TableName("transactions"), &Transaction{})
}

func (r *GeneratedTransactionResultTypeResolver) TotalPage(ctx context.Context, obj *TransactionResultType) (count int, err error) {
	total, _ := r.Total(ctx, obj)
	perPage, _ := r.PerPage(ctx, obj)
	totalPage := int(math.Ceil(float64(total) / float64(perPage)))
	if totalPage < 0 {
		totalPage = 0
	}

	return totalPage, nil
}

func (r *GeneratedTransactionResultTypeResolver) CurrentPage(ctx context.Context, obj *TransactionResultType) (count int, err error) {
	return int(*obj.EntityResultType.CurrentPage), nil
}

func (r *GeneratedTransactionResultTypeResolver) PerPage(ctx context.Context, obj *TransactionResultType) (count int, err error) {
	return int(*obj.EntityResultType.PerPage), nil
}

type GeneratedTransactionResolver struct{ *GeneratedResolver }

func (r *GeneratedTransactionResolver) Account(ctx context.Context, obj *Transaction) (res *Account, err error) {
	return r.Handlers.TransactionAccount(ctx, r.GeneratedResolver, obj)
}
func TransactionAccountHandler(ctx context.Context, r *GeneratedResolver, obj *Transaction) (items *Account, err error) {
	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)

	if obj.AccountID != nil {
		item, _ := loaders["Account"].Load(ctx, dataloader.StringKey(*obj.AccountID))()
		items, _ = item.(*Account)

		if items == nil {
			items = &Account{}
		}

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

func (r *GeneratedTodoResultTypeResolver) Total(ctx context.Context, obj *TodoResultType) (count int, err error) {
	return obj.GetTotal(ctx, r.DB.db, TableName("todos"), &Todo{})
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

func (r *GeneratedTodoResultTypeResolver) CurrentPage(ctx context.Context, obj *TodoResultType) (count int, err error) {
	return int(*obj.EntityResultType.CurrentPage), nil
}

func (r *GeneratedTodoResultTypeResolver) PerPage(ctx context.Context, obj *TodoResultType) (count int, err error) {
	return int(*obj.EntityResultType.PerPage), nil
}

type GeneratedTodoResolver struct{ *GeneratedResolver }

func (r *GeneratedTodoResolver) Account(ctx context.Context, obj *Todo) (res *User, err error) {
	return r.Handlers.TodoAccount(ctx, r.GeneratedResolver, obj)
}
func TodoAccountHandler(ctx context.Context, r *GeneratedResolver, obj *Todo) (items *User, err error) {
	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)

	if obj.AccountID != nil {
		item, _ := loaders["User"].Load(ctx, dataloader.StringKey(*obj.AccountID))()
		items, _ = item.(*User)

		if items == nil {
			items = &User{}
		}

	}

	return
}
