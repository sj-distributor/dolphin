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

	UserTasks func(ctx context.Context, r *GeneratedResolver, obj *User) (res []*Task, err error)

	CreateTask    func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Task, err error)
	UpdateTask    func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Task, err error)
	DeleteTasks   func(ctx context.Context, r *GeneratedResolver, id []string, unscoped *bool) (bool, error)
	RecoveryTasks func(ctx context.Context, r *GeneratedResolver, id []string) (bool, error)
	QueryTask     func(ctx context.Context, r *GeneratedResolver, id string) (*Task, error)
	QueryTasks    func(ctx context.Context, r *GeneratedResolver, opts QueryTasksHandlerOptions) (*TaskResultType, error)

	TaskAssignee func(ctx context.Context, r *GeneratedResolver, obj *Task) (res *User, err error)

	CreateUploadFile    func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *UploadFile, err error)
	UpdateUploadFile    func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *UploadFile, err error)
	DeleteUploadFiles   func(ctx context.Context, r *GeneratedResolver, id []string, unscoped *bool) (bool, error)
	RecoveryUploadFiles func(ctx context.Context, r *GeneratedResolver, id []string) (bool, error)
	QueryUploadFile     func(ctx context.Context, r *GeneratedResolver, id string) (*UploadFile, error)
	QueryUploadFiles    func(ctx context.Context, r *GeneratedResolver, opts QueryUploadFilesHandlerOptions) (*UploadFileResultType, error)
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

		TaskAssignee: TaskAssigneeHandler,

		CreateUploadFile:    CreateUploadFileHandler,
		UpdateUploadFile:    UpdateUploadFileHandler,
		DeleteUploadFiles:   DeleteUploadFilesHandler,
		RecoveryUploadFiles: RecoveryUploadFilesHandler,
		QueryUploadFile:     QueryUploadFileHandler,
		QueryUploadFiles:    QueryUploadFilesHandler,
	}
	return handlers
}

type GeneratedResolver struct {
	Handlers        ResolutionHandlers
	DB              *DB
	EventController *EventController
}
