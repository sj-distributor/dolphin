//go:generate go run github.com/99designs/gqlgen generate
package gen

import (
	"context"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// type Resolver struct{
// 	todos []*Todo
// 	todo *Todo
// }

type ResolutionHandlers struct {
	OnEvent func(ctx context.Context, r *GeneratedResolver, e *Event) error

	CreateTodo func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Todo, err error)
	UpdateTodo func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Todo, err error)
	QueryTodo  func(ctx context.Context, r *GeneratedResolver, id string) (*Todo, error)
	QueryTodos func(ctx context.Context, r *GeneratedResolver, opts QueryTodosHandlerOptions) (*TodoResultType, error)
}

func DefaultResolutionHandlers() ResolutionHandlers {
	handlers := ResolutionHandlers{
		OnEvent: func(ctx context.Context, r *GeneratedResolver, e *Event) error { return nil },

		CreateTodo: CreateTodoHandler,
		UpdateTodo: UpdateTodoHandler,
		QueryTodo:  QueryTodoHandler,
		QueryTodos: QueryTodosHandler,
	}
	return handlers
}

type GeneratedResolver struct {
	Handlers        ResolutionHandlers
	DB              *DB
	EventController *EventController
}
