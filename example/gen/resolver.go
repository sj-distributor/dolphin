//go:generate go run github.com/99designs/gqlgen generate
package gen

import (
	"context"
)

type ResolutionHandlers struct {
	OnEvent func(ctx context.Context, r *GeneratedResolver, e *Event) error

	CreateUser    func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}, authType bool) (item *User, err error)
	UpdateUser    func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}, authType bool) (item *User, err error)
	DeleteUsers   func(ctx context.Context, r *GeneratedResolver, id []string, unscoped *bool, authType bool) (bool, error)
	RecoveryUsers func(ctx context.Context, r *GeneratedResolver, id []string, authType bool) (bool, error)
	QueryUser     func(ctx context.Context, r *GeneratedResolver, opts QueryUserHandlerOptions, authType bool) (*User, error)
	QueryUsers    func(ctx context.Context, r *GeneratedResolver, opts QueryUsersHandlerOptions, authType bool) (*UserResultType, error)

	UserT func(ctx context.Context, r *GeneratedResolver, obj *User, authType bool) (res *Task, err error)

	UserTt func(ctx context.Context, r *GeneratedResolver, obj *User, authType bool) (res *Task, err error)

	UserTtt func(ctx context.Context, r *GeneratedResolver, obj *User, authType bool) (res []*Task, err error)

	UserTttt func(ctx context.Context, r *GeneratedResolver, obj *User, authType bool) (res []*Task, err error)

	CreateTask    func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}, authType bool) (item *Task, err error)
	UpdateTask    func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}, authType bool) (item *Task, err error)
	DeleteTasks   func(ctx context.Context, r *GeneratedResolver, id []string, unscoped *bool, authType bool) (bool, error)
	RecoveryTasks func(ctx context.Context, r *GeneratedResolver, id []string, authType bool) (bool, error)
	QueryTask     func(ctx context.Context, r *GeneratedResolver, opts QueryTaskHandlerOptions, authType bool) (*Task, error)
	QueryTasks    func(ctx context.Context, r *GeneratedResolver, opts QueryTasksHandlerOptions, authType bool) (*TaskResultType, error)

	TaskU func(ctx context.Context, r *GeneratedResolver, obj *Task, authType bool) (res *User, err error)

	TaskUu func(ctx context.Context, r *GeneratedResolver, obj *Task, authType bool) (res []*User, err error)

	TaskUuu func(ctx context.Context, r *GeneratedResolver, obj *Task, authType bool) (res *User, err error)

	TaskUuuu func(ctx context.Context, r *GeneratedResolver, obj *Task, authType bool) (res []*User, err error)
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

		UserT: UserTHandler,

		UserTt: UserTtHandler,

		UserTtt: UserTttHandler,

		UserTttt: UserTtttHandler,

		CreateTask:    CreateTaskHandler,
		UpdateTask:    UpdateTaskHandler,
		DeleteTasks:   DeleteTasksHandler,
		RecoveryTasks: RecoveryTasksHandler,
		QueryTask:     QueryTaskHandler,
		QueryTasks:    QueryTasksHandler,

		TaskU: TaskUHandler,

		TaskUu: TaskUuHandler,

		TaskUuu: TaskUuuHandler,

		TaskUuuu: TaskUuuuHandler,
	}
	return handlers
}

type GeneratedResolver struct {
	Handlers        ResolutionHandlers
	DB              *DB
	EventController *EventController
}
