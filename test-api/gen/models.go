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
	ID        string  `json:"id" gorm:"type:varchar(36) comment 'uuid';primary_key;unique_index;NOT NULL;"`
	Phone     string  `json:"phone" gorm:"type:varchar(32) comment '账号：使用手机号码';NOT NULL;index:phone;" validator:"required:true;type:phone;repeat:no;relation:no;edit:no;"`
	Password  string  `json:"password" gorm:"type:varchar(64) comment '登录密码';NOT NULL;" validator:"required:true;type:password;"`
	Email     *string `json:"email" gorm:"type:varchar(64) comment '用户邮箱地址';default:null;" validator:"required:true;type:email;"`
	Nickname  *string `json:"nickname" gorm:"type:varchar(64) comment '昵称';DEFAULT NULL;index:nickname;"`
	Age       *int64  `json:"age" gorm:"type:int(3) comment '年龄';default:1;" validator:"type:int;"`
	LastName  *string `json:"lastName" gorm:"type:varchar(255) comment 'last_name';default:null;"`
	DeletedBy *string `json:"deletedBy" gorm:"type:varchar(36) comment 'deleted_by';default:null;index:deleted_by;"`
	UpdatedBy *string `json:"updatedBy" gorm:"type:varchar(36) comment 'updated_by';default:null;index:updated_by;"`
	CreatedBy *string `json:"createdBy" gorm:"type:varchar(36) comment 'created_by';default:null;index:created_by;"`
	DeletedAt *int64  `json:"deletedAt" gorm:"type:bigint(13) comment 'deleted_at';default:null;"`
	UpdatedAt *int64  `json:"updatedAt" gorm:"type:bigint(13) comment 'updated_at';default:null; autoUpdateTime:milli;"`
	CreatedAt int64   `json:"createdAt" gorm:"type:bigint(13) comment 'created_at';default:null; autoCreateTime:milli;"`

	Tasks []*Task `json:"tasks" gorm:"foreignkey:AssigneeID"`
}

func (m *User) Is_Entity() {}

type UserChanges struct {
	ID        string
	Phone     string
	Password  string
	Email     *string
	Nickname  *string
	Age       *int64
	LastName  *string
	DeletedBy *string
	UpdatedBy *string
	CreatedBy *string
	DeletedAt *int64
	UpdatedAt *int64
	CreatedAt int64

	Tasks []*Task

	TasksIDs []*string
}

type TaskResultType struct {
	EntityResultType
}

type Task struct {
	ID         string     `json:"id" gorm:"type:varchar(36) comment 'uuid';primary_key;unique_index;NOT NULL;"`
	Title      *string    `json:"title" gorm:"type:varchar(255) comment 'title';default:null;"`
	Completed  *bool      `json:"completed" gorm:"default:null"`
	DueDate    *time.Time `json:"dueDate" gorm:"default:null"`
	AssigneeID *string    `json:"assigneeId" gorm:"type:varchar(36) comment 'assignee_id';default:null;"`
	DeletedBy  *string    `json:"deletedBy" gorm:"type:varchar(36) comment 'deleted_by';default:null;index:deleted_by;"`
	UpdatedBy  *string    `json:"updatedBy" gorm:"type:varchar(36) comment 'updated_by';default:null;index:updated_by;"`
	CreatedBy  *string    `json:"createdBy" gorm:"type:varchar(36) comment 'created_by';default:null;index:created_by;"`
	DeletedAt  *int64     `json:"deletedAt" gorm:"type:bigint(13) comment 'deleted_at';default:null;"`
	UpdatedAt  *int64     `json:"updatedAt" gorm:"type:bigint(13) comment 'updated_at';default:null; autoUpdateTime:milli;"`
	CreatedAt  int64      `json:"createdAt" gorm:"type:bigint(13) comment 'created_at';default:null; autoCreateTime:milli;"`

	Assignee *User `json:"assignee"`
}

func (m *Task) Is_Entity() {}

type TaskChanges struct {
	ID         string
	Title      *string
	Completed  *bool
	DueDate    *time.Time
	AssigneeID *string
	DeletedBy  *string
	UpdatedBy  *string
	CreatedBy  *string
	DeletedAt  *int64
	UpdatedAt  *int64
	CreatedAt  int64

	Assignee *User
}

type UploadFileResultType struct {
	EntityResultType
}

type UploadFile struct {
	ID        string  `json:"id" gorm:"type:varchar(36) comment 'uuid';primary_key;unique_index;NOT NULL;"`
	Name      string  `json:"name" gorm:"type:varchar(255) comment '文件名称';NOT NULL;index:name;" validator:"required:true;repeat:no;"`
	Hash      string  `json:"hash" gorm:"type:text comment '文件hash值';NOT NULL;" validator:"required:true;repeat:no;"`
	DeletedBy *string `json:"deletedBy" gorm:"type:varchar(36) comment 'deleted_by';default:null;index:deleted_by;"`
	UpdatedBy *string `json:"updatedBy" gorm:"type:varchar(36) comment 'updated_by';default:null;index:updated_by;"`
	CreatedBy *string `json:"createdBy" gorm:"type:varchar(36) comment 'created_by';default:null;index:created_by;"`
	DeletedAt *int64  `json:"deletedAt" gorm:"type:bigint(13) comment 'deleted_at';default:null;"`
	UpdatedAt *int64  `json:"updatedAt" gorm:"type:bigint(13) comment 'updated_at';default:null; autoUpdateTime:milli;"`
	CreatedAt int64   `json:"createdAt" gorm:"type:bigint(13) comment 'created_at';default:null; autoCreateTime:milli;"`
}

func (m *UploadFile) Is_Entity() {}

type UploadFileChanges struct {
	ID        string
	Name      string
	Hash      string
	DeletedBy *string
	UpdatedBy *string
	CreatedBy *string
	DeletedAt *int64
	UpdatedAt *int64
	CreatedAt int64
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
