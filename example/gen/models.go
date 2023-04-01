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

type TodoResultType struct {
	EntityResultType
}

type Todo struct {
	ID        string  `json:"id" gorm:"type:varchar(36) comment 'uuid';primary_key;unique_index;NOT NULL;"`
	Title     string  `json:"title" gorm:"type:varchar(255) comment 'title';default:null;" validator:"type:password;"`
	Age       *int64  `json:"age" gorm:"type:bigint(13) comment 'age';default:null;"`
	Money     int64   `json:"money" gorm:"type:bigint(13) comment 'money';default:null;"`
	Remark    *string `json:"remark" gorm:"type:varchar(255) comment 'remark';default:null;"`
	UserID    *string `json:"userId" gorm:"type:varchar(36) comment 'user_id';default:null;"`
	DeletedBy *string `json:"deletedBy" gorm:"type:varchar(36) comment 'deleted_by';default:null;index:deleted_by;"`
	UpdatedBy *string `json:"updatedBy" gorm:"type:varchar(36) comment 'updated_by';default:null;index:updated_by;"`
	CreatedBy *string `json:"createdBy" gorm:"type:varchar(36) comment 'created_by';default:null;index:created_by;"`
	DeletedAt *int64  `json:"deletedAt" gorm:"type:bigint(13) comment 'deleted_at';default:null;"`
	UpdatedAt *int64  `json:"updatedAt" gorm:"type:bigint(13) comment 'updated_at';default:null;"`
	CreatedAt int64   `json:"createdAt" gorm:"type:bigint(13) comment 'created_at';default:null;"`

	User *User `json:"user"`
}

func (m *Todo) Is_Entity() {}

type TodoChanges struct {
	ID        string
	Title     string
	Age       *int64
	Money     int64
	Remark    *string
	UserID    *string
	DeletedBy *string
	UpdatedBy *string
	CreatedBy *string
	DeletedAt *int64
	UpdatedAt *int64
	CreatedAt int64

	User *User
}

type UserResultType struct {
	EntityResultType
}

type User struct {
	ID        string  `json:"id" gorm:"type:varchar(36) comment 'uuid';primary_key;unique_index;NOT NULL;"`
	Username  string  `json:"username" gorm:"type:varchar(64) comment '用户名称';DEFAULT NULL;index:username;"`
	TodoID    *string `json:"todoId" gorm:"type:varchar(36) comment 'todo_id';default:null;"`
	DeletedBy *string `json:"deletedBy" gorm:"type:varchar(36) comment 'deleted_by';default:null;index:deleted_by;"`
	UpdatedBy *string `json:"updatedBy" gorm:"type:varchar(36) comment 'updated_by';default:null;index:updated_by;"`
	CreatedBy *string `json:"createdBy" gorm:"type:varchar(36) comment 'created_by';default:null;index:created_by;"`
	DeletedAt *int64  `json:"deletedAt" gorm:"type:bigint(13) comment 'deleted_at';default:null;"`
	UpdatedAt *int64  `json:"updatedAt" gorm:"type:bigint(13) comment 'updated_at';default:null;"`
	CreatedAt int64   `json:"createdAt" gorm:"type:bigint(13) comment 'created_at';default:null;"`

	Todo *Todo `json:"todo"`
}

func (m *User) Is_Entity() {}

type UserChanges struct {
	ID        string
	Username  string
	TodoID    *string
	DeletedBy *string
	UpdatedBy *string
	CreatedBy *string
	DeletedAt *int64
	UpdatedAt *int64
	CreatedAt int64

	Todo *Todo
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
