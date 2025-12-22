//go:generate go run github.com/99designs/gqlgen generate
package gen

import (
	"context"
)

type ResolutionHandlers struct {
	OnEvent   func(ctx context.Context, r *GeneratedResolver, e *Event) error
	WebSocket func(ctx context.Context, r *GeneratedResolver) (<-chan any, error)

	CreateUser    func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *User, err error)
	UpdateUser    func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *User, err error)
	DeleteUsers   func(ctx context.Context, r *GeneratedResolver, id []string, unscoped *bool) (bool, error)
	RecoveryUsers func(ctx context.Context, r *GeneratedResolver, id []string) (bool, error)
	QueryUser     func(ctx context.Context, r *GeneratedResolver, opts QueryUserHandlerOptions) (*User, error)
	QueryUsers    func(ctx context.Context, r *GeneratedResolver, opts QueryUsersHandlerOptions) (*UserResultType, error)

	UserProfile func(ctx context.Context, r *GeneratedResolver, obj *User) (res *Profile, err error)

	UserTasks func(ctx context.Context, r *GeneratedResolver, obj *User) (res []*Task, err error)

	UserUserRoles func(ctx context.Context, r *GeneratedResolver, obj *User) (res []*UserRole, err error)

	CreateProfile    func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Profile, err error)
	UpdateProfile    func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Profile, err error)
	DeleteProfiles   func(ctx context.Context, r *GeneratedResolver, id []string, unscoped *bool) (bool, error)
	RecoveryProfiles func(ctx context.Context, r *GeneratedResolver, id []string) (bool, error)
	QueryProfile     func(ctx context.Context, r *GeneratedResolver, opts QueryProfileHandlerOptions) (*Profile, error)
	QueryProfiles    func(ctx context.Context, r *GeneratedResolver, opts QueryProfilesHandlerOptions) (*ProfileResultType, error)

	ProfileUser func(ctx context.Context, r *GeneratedResolver, obj *Profile) (res *User, err error)

	CreateTask    func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Task, err error)
	UpdateTask    func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Task, err error)
	DeleteTasks   func(ctx context.Context, r *GeneratedResolver, id []string, unscoped *bool) (bool, error)
	RecoveryTasks func(ctx context.Context, r *GeneratedResolver, id []string) (bool, error)
	QueryTask     func(ctx context.Context, r *GeneratedResolver, opts QueryTaskHandlerOptions) (*Task, error)
	QueryTasks    func(ctx context.Context, r *GeneratedResolver, opts QueryTasksHandlerOptions) (*TaskResultType, error)

	TaskUser func(ctx context.Context, r *GeneratedResolver, obj *Task) (res *User, err error)

	TaskTags func(ctx context.Context, r *GeneratedResolver, obj *Task) (res []*Tag, err error)

	CreateUserRole    func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *UserRole, err error)
	UpdateUserRole    func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *UserRole, err error)
	DeleteUserRoles   func(ctx context.Context, r *GeneratedResolver, id []string, unscoped *bool) (bool, error)
	RecoveryUserRoles func(ctx context.Context, r *GeneratedResolver, id []string) (bool, error)
	QueryUserRole     func(ctx context.Context, r *GeneratedResolver, opts QueryUserRoleHandlerOptions) (*UserRole, error)
	QueryUserRoles    func(ctx context.Context, r *GeneratedResolver, opts QueryUserRolesHandlerOptions) (*UserRoleResultType, error)

	UserRoleUsers func(ctx context.Context, r *GeneratedResolver, obj *UserRole) (res []*User, err error)

	CreateTag    func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Tag, err error)
	UpdateTag    func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Tag, err error)
	DeleteTags   func(ctx context.Context, r *GeneratedResolver, id []string, unscoped *bool) (bool, error)
	RecoveryTags func(ctx context.Context, r *GeneratedResolver, id []string) (bool, error)
	QueryTag     func(ctx context.Context, r *GeneratedResolver, opts QueryTagHandlerOptions) (*Tag, error)
	QueryTags    func(ctx context.Context, r *GeneratedResolver, opts QueryTagsHandlerOptions) (*TagResultType, error)

	TagTasks func(ctx context.Context, r *GeneratedResolver, obj *Tag) (res []*Task, err error)
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

		UserProfile: UserProfileHandler,

		UserTasks: UserTasksHandler,

		UserUserRoles: UserUserRolesHandler,

		CreateProfile:    CreateProfileHandler,
		UpdateProfile:    UpdateProfileHandler,
		DeleteProfiles:   DeleteProfilesHandler,
		RecoveryProfiles: RecoveryProfilesHandler,
		QueryProfile:     QueryProfileHandler,
		QueryProfiles:    QueryProfilesHandler,

		ProfileUser: ProfileUserHandler,

		CreateTask:    CreateTaskHandler,
		UpdateTask:    UpdateTaskHandler,
		DeleteTasks:   DeleteTasksHandler,
		RecoveryTasks: RecoveryTasksHandler,
		QueryTask:     QueryTaskHandler,
		QueryTasks:    QueryTasksHandler,

		TaskUser: TaskUserHandler,

		TaskTags: TaskTagsHandler,

		CreateUserRole:    CreateUserRoleHandler,
		UpdateUserRole:    UpdateUserRoleHandler,
		DeleteUserRoles:   DeleteUserRolesHandler,
		RecoveryUserRoles: RecoveryUserRolesHandler,
		QueryUserRole:     QueryUserRoleHandler,
		QueryUserRoles:    QueryUserRolesHandler,

		UserRoleUsers: UserRoleUsersHandler,

		CreateTag:    CreateTagHandler,
		UpdateTag:    UpdateTagHandler,
		DeleteTags:   DeleteTagsHandler,
		RecoveryTags: RecoveryTagsHandler,
		QueryTag:     QueryTagHandler,
		QueryTags:    QueryTagsHandler,

		TagTasks: TagTasksHandler,

		WebSocket: WebSocketHandler,
	}
	return handlers
}

type GeneratedResolver struct {
	Handlers        ResolutionHandlers
	DB              *DB
	EventController *EventController
}
