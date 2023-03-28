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

	UserTodo func(ctx context.Context, r *GeneratedResolver, obj *User) (res *Todo, err error)

	CreateTodo    func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Todo, err error)
	UpdateTodo    func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Todo, err error)
	DeleteTodos   func(ctx context.Context, r *GeneratedResolver, id []string, unscoped *bool) (bool, error)
	RecoveryTodos func(ctx context.Context, r *GeneratedResolver, id []string) (bool, error)
	QueryTodo     func(ctx context.Context, r *GeneratedResolver, id string) (*Todo, error)
	QueryTodos    func(ctx context.Context, r *GeneratedResolver, opts QueryTodosHandlerOptions) (*TodoResultType, error)

	TodoUser func(ctx context.Context, r *GeneratedResolver, obj *Todo) (res *User, err error)
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

		UserTodo: UserTodoHandler,

		CreateTodo:    CreateTodoHandler,
		UpdateTodo:    UpdateTodoHandler,
		DeleteTodos:   DeleteTodosHandler,
		RecoveryTodos: RecoveryTodosHandler,
		QueryTodo:     QueryTodoHandler,
		QueryTodos:    QueryTodosHandler,

		TodoUser: TodoUserHandler,
	}
	return handlers
}

type GeneratedResolver struct {
	Handlers        ResolutionHandlers
	DB              *DB
	EventController *EventController
}
