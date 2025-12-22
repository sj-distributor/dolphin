package gen

import (
	"fmt"
	"reflect"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/mitchellh/mapstructure"
)

type NotFoundError struct {
	Entity string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s not found", e.Entity)
}

type UserResultType struct {
	EntityResultType
}

type User struct {
	ID        string  `json:"id" gorm:"type:varchar(36);comment:'uuid';primaryKey;uniqueIndex;NOT NULL;"`
	Phone     string  `json:"phone" gorm:"type:varchar(32);comment:'phone';NOT NULL;index:phone;" validator:"required:true;type:phone"`
	Password  string  `json:"password" gorm:"type:varchar(64);comment:'password';NOT NULL;" validator:"required:true;type:password"`
	Email     *string `json:"email" gorm:"type:varchar(64);comment:'email';default:null;" validator:"type:email"`
	Nickname  *string `json:"nickname" gorm:"type:varchar(64);comment:'nickname';DEFAULT NULL;index:nickname;"`
	Age       *int64  `json:"age" gorm:"type:int(3);comment:'age';default:1;" validator:"type:int"`
	ProfileID *string `json:"profileId" gorm:"type:varchar(36);comment:'profile_id';default:null;"`
	IsDelete  *int64  `json:"isDelete" gorm:"type:int(2);comment:'是否删除：1/正常、2/删除';default:1;index:is_delete;"`
	Weight    *int64  `json:"weight" gorm:"type:int(11);comment:'权重：用来排序';default:1;index:weight;"`
	State     *int64  `json:"state" gorm:"type:int(2);comment:'状态：1/正常、2/禁用';default:1;index:state;"`
	DeletedBy *string `json:"deletedBy" gorm:"type:varchar(36);comment:'deleted_by';default:null;index:deleted_by;"`
	UpdatedBy *string `json:"updatedBy" gorm:"type:varchar(36);comment:'updated_by';default:null;index:updated_by;"`
	CreatedBy *string `json:"createdBy" gorm:"type:varchar(36);comment:'created_by';default:null;index:created_by;"`
	DeletedAt *int64  `json:"deletedAt" gorm:"type:bigint(13);comment:'deleted_at';default:null;"`
	UpdatedAt *int64  `json:"updatedAt" gorm:"type:bigint(13);comment:'updated_at';default:null; autoUpdateTime:milli;"`
	CreatedAt int64   `json:"createdAt" gorm:"type:bigint(13);comment:'created_at';default:null; autoCreateTime:milli;"`

	Profile *Profile `json:"profile"`

	Tasks []*Task `json:"tasks" gorm:"foreignkey:UserID"`

	UserRoles []*UserRole `json:"userRoles" gorm:"many2many:userRole_users;jointable_foreignkey:user_id;association_jointable_foreignkey:userRole_id"`
}

func (m *User) Is_Entity() {}

type UserChanges struct {
	ID        string
	Phone     string
	Password  string
	Email     *string
	Nickname  *string
	Age       *int64
	ProfileID *string
	IsDelete  *int64
	Weight    *int64
	State     *int64
	DeletedBy *string
	UpdatedBy *string
	CreatedBy *string
	DeletedAt *int64
	UpdatedAt *int64
	CreatedAt int64

	Profile   *Profile
	Tasks     []*Task
	UserRoles []*UserRole

	TasksIDs     []*string
	UserRolesIDs []*string
}

type ProfileResultType struct {
	EntityResultType
}

type Profile struct {
	ID        string     `json:"id" gorm:"type:varchar(36);comment:'uuid';primaryKey;uniqueIndex;NOT NULL;"`
	Avatar    *string    `json:"avatar" gorm:"type:varchar(255);comment:'avatar url';default:null;"`
	Bio       *string    `json:"bio" gorm:"type:text;comment:'bio';"`
	Birthday  *time.Time `json:"birthday" gorm:"comment:'birthday';"`
	Address   *string    `json:"address" gorm:"type:varchar(255);comment:'address';default:null;"`
	UserID    string     `json:"userId" gorm:"type:varchar(36);comment:'user_id';default:null;"`
	IsDelete  *int64     `json:"isDelete" gorm:"type:int(2);comment:'是否删除：1/正常、2/删除';default:1;index:is_delete;"`
	Weight    *int64     `json:"weight" gorm:"type:int(11);comment:'权重：用来排序';default:1;index:weight;"`
	State     *int64     `json:"state" gorm:"type:int(2);comment:'状态：1/正常、2/禁用';default:1;index:state;"`
	DeletedBy *string    `json:"deletedBy" gorm:"type:varchar(36);comment:'deleted_by';default:null;index:deleted_by;"`
	UpdatedBy *string    `json:"updatedBy" gorm:"type:varchar(36);comment:'updated_by';default:null;index:updated_by;"`
	CreatedBy *string    `json:"createdBy" gorm:"type:varchar(36);comment:'created_by';default:null;index:created_by;"`
	DeletedAt *int64     `json:"deletedAt" gorm:"type:bigint(13);comment:'deleted_at';default:null;"`
	UpdatedAt *int64     `json:"updatedAt" gorm:"type:bigint(13);comment:'updated_at';default:null; autoUpdateTime:milli;"`
	CreatedAt int64      `json:"createdAt" gorm:"type:bigint(13);comment:'created_at';default:null; autoCreateTime:milli;"`

	User *User `json:"user"`
}

func (m *Profile) Is_Entity() {}

type ProfileChanges struct {
	ID        string
	Avatar    *string
	Bio       *string
	Birthday  *time.Time
	Address   *string
	UserID    string
	IsDelete  *int64
	Weight    *int64
	State     *int64
	DeletedBy *string
	UpdatedBy *string
	CreatedBy *string
	DeletedAt *int64
	UpdatedAt *int64
	CreatedAt int64

	User *User
}

type TaskResultType struct {
	EntityResultType
}

type Task struct {
	ID          string     `json:"id" gorm:"type:varchar(36);comment:'uuid';primaryKey;uniqueIndex;NOT NULL;"`
	Title       string     `json:"title" gorm:"type:varchar(128);comment:'title';NOT NULL;"`
	Description *string    `json:"description" gorm:"type:text;comment:'description';"`
	Completed   *bool      `json:"completed" gorm:"type:tinyint(1);comment:'completed';default:0;"`
	DueDate     *time.Time `json:"dueDate" gorm:"comment:'due date';"`
	Priority    *int64     `json:"priority" gorm:"type:int(2);comment:'priority';default:2;"`
	UserID      *string    `json:"userId" gorm:"type:varchar(36);comment:'user_id';default:null;"`
	IsDelete    *int64     `json:"isDelete" gorm:"type:int(2);comment:'是否删除：1/正常、2/删除';default:1;index:is_delete;"`
	Weight      *int64     `json:"weight" gorm:"type:int(11);comment:'权重：用来排序';default:1;index:weight;"`
	State       *int64     `json:"state" gorm:"type:int(2);comment:'状态：1/正常、2/禁用';default:1;index:state;"`
	DeletedBy   *string    `json:"deletedBy" gorm:"type:varchar(36);comment:'deleted_by';default:null;index:deleted_by;"`
	UpdatedBy   *string    `json:"updatedBy" gorm:"type:varchar(36);comment:'updated_by';default:null;index:updated_by;"`
	CreatedBy   *string    `json:"createdBy" gorm:"type:varchar(36);comment:'created_by';default:null;index:created_by;"`
	DeletedAt   *int64     `json:"deletedAt" gorm:"type:bigint(13);comment:'deleted_at';default:null;"`
	UpdatedAt   *int64     `json:"updatedAt" gorm:"type:bigint(13);comment:'updated_at';default:null; autoUpdateTime:milli;"`
	CreatedAt   int64      `json:"createdAt" gorm:"type:bigint(13);comment:'created_at';default:null; autoCreateTime:milli;"`

	User *User `json:"user"`

	Tags []*Tag `json:"tags" gorm:"many2many:tag_tasks;jointable_foreignkey:task_id;association_jointable_foreignkey:tag_id"`
}

func (m *Task) Is_Entity() {}

type TaskChanges struct {
	ID          string
	Title       string
	Description *string
	Completed   *bool
	DueDate     *time.Time
	Priority    *int64
	UserID      *string
	IsDelete    *int64
	Weight      *int64
	State       *int64
	DeletedBy   *string
	UpdatedBy   *string
	CreatedBy   *string
	DeletedAt   *int64
	UpdatedAt   *int64
	CreatedAt   int64

	User *User
	Tags []*Tag

	TagsIDs []*string
}

type UserRoleResultType struct {
	EntityResultType
}

type UserRole struct {
	ID          string  `json:"id" gorm:"type:varchar(36);comment:'uuid';primaryKey;uniqueIndex;NOT NULL;"`
	Name        string  `json:"name" gorm:"type:varchar(64);comment:'name';NOT NULL;uniqueIndex;"`
	Description *string `json:"description" gorm:"type:varchar(255);comment:'description';default:null;"`
	Permissions *string `json:"permissions" gorm:"type:text;comment:'permissions';"`
	IsDelete    *int64  `json:"isDelete" gorm:"type:int(2);comment:'是否删除：1/正常、2/删除';default:1;index:is_delete;"`
	Weight      *int64  `json:"weight" gorm:"type:int(11);comment:'权重：用来排序';default:1;index:weight;"`
	State       *int64  `json:"state" gorm:"type:int(2);comment:'状态：1/正常、2/禁用';default:1;index:state;"`
	DeletedBy   *string `json:"deletedBy" gorm:"type:varchar(36);comment:'deleted_by';default:null;index:deleted_by;"`
	UpdatedBy   *string `json:"updatedBy" gorm:"type:varchar(36);comment:'updated_by';default:null;index:updated_by;"`
	CreatedBy   *string `json:"createdBy" gorm:"type:varchar(36);comment:'created_by';default:null;index:created_by;"`
	DeletedAt   *int64  `json:"deletedAt" gorm:"type:bigint(13);comment:'deleted_at';default:null;"`
	UpdatedAt   *int64  `json:"updatedAt" gorm:"type:bigint(13);comment:'updated_at';default:null; autoUpdateTime:milli;"`
	CreatedAt   int64   `json:"createdAt" gorm:"type:bigint(13);comment:'created_at';default:null; autoCreateTime:milli;"`

	Users []*User `json:"users" gorm:"many2many:userRole_users;jointable_foreignkey:userRole_id;association_jointable_foreignkey:user_id"`
}

func (m *UserRole) Is_Entity() {}

type UserRoleChanges struct {
	ID          string
	Name        string
	Description *string
	Permissions *string
	IsDelete    *int64
	Weight      *int64
	State       *int64
	DeletedBy   *string
	UpdatedBy   *string
	CreatedBy   *string
	DeletedAt   *int64
	UpdatedAt   *int64
	CreatedAt   int64

	Users []*User

	UsersIDs []*string
}

type TagResultType struct {
	EntityResultType
}

type Tag struct {
	ID        string  `json:"id" gorm:"type:varchar(36);comment:'uuid';primaryKey;uniqueIndex;NOT NULL;"`
	Name      string  `json:"name" gorm:"type:varchar(32);comment:'name';NOT NULL;uniqueIndex;"`
	Color     *string `json:"color" gorm:"type:varchar(16);comment:'color';default:'blue';"`
	IsDelete  *int64  `json:"isDelete" gorm:"type:int(2);comment:'是否删除：1/正常、2/删除';default:1;index:is_delete;"`
	Weight    *int64  `json:"weight" gorm:"type:int(11);comment:'权重：用来排序';default:1;index:weight;"`
	State     *int64  `json:"state" gorm:"type:int(2);comment:'状态：1/正常、2/禁用';default:1;index:state;"`
	DeletedBy *string `json:"deletedBy" gorm:"type:varchar(36);comment:'deleted_by';default:null;index:deleted_by;"`
	UpdatedBy *string `json:"updatedBy" gorm:"type:varchar(36);comment:'updated_by';default:null;index:updated_by;"`
	CreatedBy *string `json:"createdBy" gorm:"type:varchar(36);comment:'created_by';default:null;index:created_by;"`
	DeletedAt *int64  `json:"deletedAt" gorm:"type:bigint(13);comment:'deleted_at';default:null;"`
	UpdatedAt *int64  `json:"updatedAt" gorm:"type:bigint(13);comment:'updated_at';default:null; autoUpdateTime:milli;"`
	CreatedAt int64   `json:"createdAt" gorm:"type:bigint(13);comment:'created_at';default:null; autoCreateTime:milli;"`

	Tasks []*Task `json:"tasks" gorm:"many2many:tag_tasks;jointable_foreignkey:tag_id;association_jointable_foreignkey:task_id"`
}

func (m *Tag) Is_Entity() {}

type TagChanges struct {
	ID        string
	Name      string
	Color     *string
	IsDelete  *int64
	Weight    *int64
	State     *int64
	DeletedBy *string
	UpdatedBy *string
	CreatedBy *string
	DeletedAt *int64
	UpdatedAt *int64
	CreatedAt int64

	Tasks []*Task

	TasksIDs []*string
}

// used to convert map[string]interface{} to EntityChanges struct
func ApplyChanges(changes map[string]interface{}, to interface{}) error {
	dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		ErrorUnused: true,
		TagName:     "json",
		Result:      to,
		ZeroFields:  true,
		// This is needed to get mapstructure to call the gqlgen unmarshaler func for custom scalars (eg Date)
		DecodeHook: func(a reflect.Type, b reflect.Type, v interface{}) (interface{}, error) {

			if b == reflect.TypeOf(time.Time{}) {
				switch a.Kind() {
				case reflect.String:
					return time.Parse(time.RFC3339, v.(string))
				case reflect.Float64:
					return time.Unix(0, int64(v.(float64))*int64(time.Millisecond)), nil
				case reflect.Int64:
					return time.Unix(0, v.(int64)*int64(time.Millisecond)), nil
				case reflect.Struct:
					if t, ok := v.(time.Time); ok {
						return t, nil
					}
					return v, nil
				default:
					return v, fmt.Errorf("Unable to parse date from %v", v)
				}
			}

			if reflect.PtrTo(b).Implements(reflect.TypeOf((*graphql.Unmarshaler)(nil)).Elem()) {
				resultType := reflect.New(b)
				result := resultType.MethodByName("UnmarshalGQL").Call([]reflect.Value{reflect.ValueOf(v)})
				err, _ := result[0].Interface().(error)
				return resultType.Elem().Interface(), err
			}

			return v, nil
		},
	})

	if err != nil {
		return err
	}

	return dec.Decode(changes)
}
