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
	QueryUser     func(ctx context.Context, r *GeneratedResolver, opts QueryUserHandlerOptions) (*User, error)
	QueryUsers    func(ctx context.Context, r *GeneratedResolver, opts QueryUsersHandlerOptions) (*UserResultType, error)

	UserTasks func(ctx context.Context, r *GeneratedResolver, obj *User) (res []*Task, err error)

	CreateTask    func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Task, err error)
	UpdateTask    func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Task, err error)
	DeleteTasks   func(ctx context.Context, r *GeneratedResolver, id []string, unscoped *bool) (bool, error)
	RecoveryTasks func(ctx context.Context, r *GeneratedResolver, id []string) (bool, error)
	QueryTask     func(ctx context.Context, r *GeneratedResolver, opts QueryTaskHandlerOptions) (*Task, error)
	QueryTasks    func(ctx context.Context, r *GeneratedResolver, opts QueryTasksHandlerOptions) (*TaskResultType, error)

	TaskUser func(ctx context.Context, r *GeneratedResolver, obj *Task) (res *User, err error)
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

		UserTasks: UserTasksHandler,

		CreateTask:    CreateTaskHandler,
		UpdateTask:    UpdateTaskHandler,
		DeleteTasks:   DeleteTasksHandler,
		RecoveryTasks: RecoveryTasksHandler,
		QueryTask:     QueryTaskHandler,
		QueryTasks:    QueryTasksHandler,

		TaskUser: TaskUserHandler,
	}
	return handlers
}

type GeneratedResolver struct {
	Handlers        ResolutionHandlers
	DB              *DB
	EventController *EventController
}
