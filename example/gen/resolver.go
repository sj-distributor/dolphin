//go:generate go run github.com/99designs/gqlgen generate
package gen

import (
	"context"
)

type ResolutionHandlers struct {
	OnEvent func(ctx context.Context, r *GeneratedResolver, e *Event) error

	CreateUser    func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *User, err error)
	UpdateUser    func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *User, err error)
	DeleteUsers   func(ctx context.Context, r *GeneratedResolver, id []string, unscoped *bool) (bool, error)
	RecoveryUsers func(ctx context.Context, r *GeneratedResolver, id []string) (bool, error)
	QueryUser     func(ctx context.Context, r *GeneratedResolver, id string) (*User, error)
	QueryUsers    func(ctx context.Context, r *GeneratedResolver, opts QueryUsersHandlerOptions) (*UserResultType, error)

	UserAccounts func(ctx context.Context, r *GeneratedResolver, obj *User) (res []*Account, err error)

	UserTodo func(ctx context.Context, r *GeneratedResolver, obj *User) (res []*Todo, err error)

	CreateAccount    func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Account, err error)
	UpdateAccount    func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Account, err error)
	DeleteAccounts   func(ctx context.Context, r *GeneratedResolver, id []string, unscoped *bool) (bool, error)
	RecoveryAccounts func(ctx context.Context, r *GeneratedResolver, id []string) (bool, error)
	QueryAccount     func(ctx context.Context, r *GeneratedResolver, id string) (*Account, error)
	QueryAccounts    func(ctx context.Context, r *GeneratedResolver, opts QueryAccountsHandlerOptions) (*AccountResultType, error)

	AccountOwner func(ctx context.Context, r *GeneratedResolver, obj *Account) (res *User, err error)

	AccountTransactions func(ctx context.Context, r *GeneratedResolver, obj *Account) (res []*Transaction, err error)

	CreateTransaction    func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Transaction, err error)
	UpdateTransaction    func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Transaction, err error)
	DeleteTransactions   func(ctx context.Context, r *GeneratedResolver, id []string, unscoped *bool) (bool, error)
	RecoveryTransactions func(ctx context.Context, r *GeneratedResolver, id []string) (bool, error)
	QueryTransaction     func(ctx context.Context, r *GeneratedResolver, id string) (*Transaction, error)
	QueryTransactions    func(ctx context.Context, r *GeneratedResolver, opts QueryTransactionsHandlerOptions) (*TransactionResultType, error)

	TransactionAccount func(ctx context.Context, r *GeneratedResolver, obj *Transaction) (res *Account, err error)

	CreateTodo    func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Todo, err error)
	UpdateTodo    func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Todo, err error)
	DeleteTodos   func(ctx context.Context, r *GeneratedResolver, id []string, unscoped *bool) (bool, error)
	RecoveryTodos func(ctx context.Context, r *GeneratedResolver, id []string) (bool, error)
	QueryTodo     func(ctx context.Context, r *GeneratedResolver, id string) (*Todo, error)
	QueryTodos    func(ctx context.Context, r *GeneratedResolver, opts QueryTodosHandlerOptions) (*TodoResultType, error)

	TodoAccount func(ctx context.Context, r *GeneratedResolver, obj *Todo) (res *User, err error)
}

func DefaultResolutionHandlers() ResolutionHandlers {
	handlers := ResolutionHandlers{
		OnEvent: func(ctx context.Context, r *GeneratedResolver, e *Event) error { return nil },

		CreateUser:    CreateUserHandler,
		UpdateUser:    UpdateUserHandler,
		DeleteUsers:   DeleteUsersHandler,
		RecoveryUsers: RecoveryUsersHandler,
		QueryUser:     QueryUserHandler,
		QueryUsers:    QueryUsersHandler,

		UserAccounts: UserAccountsHandler,

		UserTodo: UserTodoHandler,

		CreateAccount:    CreateAccountHandler,
		UpdateAccount:    UpdateAccountHandler,
		DeleteAccounts:   DeleteAccountsHandler,
		RecoveryAccounts: RecoveryAccountsHandler,
		QueryAccount:     QueryAccountHandler,
		QueryAccounts:    QueryAccountsHandler,

		AccountOwner: AccountOwnerHandler,

		AccountTransactions: AccountTransactionsHandler,

		CreateTransaction:    CreateTransactionHandler,
		UpdateTransaction:    UpdateTransactionHandler,
		DeleteTransactions:   DeleteTransactionsHandler,
		RecoveryTransactions: RecoveryTransactionsHandler,
		QueryTransaction:     QueryTransactionHandler,
		QueryTransactions:    QueryTransactionsHandler,

		TransactionAccount: TransactionAccountHandler,

		CreateTodo:    CreateTodoHandler,
		UpdateTodo:    UpdateTodoHandler,
		DeleteTodos:   DeleteTodosHandler,
		RecoveryTodos: RecoveryTodosHandler,
		QueryTodo:     QueryTodoHandler,
		QueryTodos:    QueryTodosHandler,

		TodoAccount: TodoAccountHandler,
	}
	return handlers
}

type GeneratedResolver struct {
	Handlers        ResolutionHandlers
	DB              *DB
	EventController *EventController
}
