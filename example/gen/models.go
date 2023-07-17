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
	Phone     string  `json:"phone" gorm:"type:varchar(32) comment '账号：使用手机号码';NOT NULL;index:phone;" validator:"type:phone;"`
	TID       *string `json:"tId" gorm:"type:varchar(36) comment 't_id';default:null;"`
	TtID      *string `json:"ttId" gorm:"type:varchar(36) comment 'tt_id';default:null;"`
	IsDelete  *int64  `json:"isDelete" gorm:"type:int(2) comment '是否删除：1/正常、2/删除';default:1;index:is_delete;"`
	Weight    *int64  `json:"weight" gorm:"type:int(11) comment '权重：用来排序';default:1;index:weight;"`
	State     *int64  `json:"state" gorm:"type:int(2) comment '状态：1/正常、2/禁用';default:1;index:state;"`
	DeletedBy *string `json:"deletedBy" gorm:"type:varchar(36) comment 'deleted_by';default:null;index:deleted_by;"`
	UpdatedBy *string `json:"updatedBy" gorm:"type:varchar(36) comment 'updated_by';default:null;index:updated_by;"`
	CreatedBy *string `json:"createdBy" gorm:"type:varchar(36) comment 'created_by';default:null;index:created_by;"`
	DeletedAt *int64  `json:"deletedAt" gorm:"type:bigint(13) comment 'deleted_at';default:null;"`
	UpdatedAt *int64  `json:"updatedAt" gorm:"type:bigint(13) comment 'updated_at';default:null; autoUpdateTime:milli;"`
	CreatedAt int64   `json:"createdAt" gorm:"type:bigint(13) comment 'created_at';default:null; autoCreateTime:milli;"`

	T *Task `json:"t"`

	Tt *Task `json:"tt"`

	Ttt []*Task `json:"ttt" gorm:"foreignkey:UuuID"`

	Tttt []*Task `json:"tttt" gorm:"many2many:task_uuuu;jointable_foreignkey:uuuu_id;association_jointable_foreignkey:tttt_id"`
}

func (m *User) Is_Entity() {}

type UserChanges struct {
	ID        string
	Phone     string
	TID       *string
	TtID      *string
	IsDelete  *int64
	Weight    *int64
	State     *int64
	DeletedBy *string
	UpdatedBy *string
	CreatedBy *string
	DeletedAt *int64
	UpdatedAt *int64
	CreatedAt int64

	T    *Task
	Tt   *Task
	Ttt  []*Task
	Tttt []*Task

	TttIDs  []*string
	TtttIDs []*string
}

type TaskResultType struct {
	EntityResultType
}

type Task struct {
	ID        string  `json:"id" gorm:"type:varchar(36) comment 'uuid';primary_key;unique_index;NOT NULL;"`
	Title     *string `json:"title" gorm:"type:varchar(64) comment '标题';NOT NULL;"`
	UID       *string `json:"uId" gorm:"type:varchar(36) comment 'u_id';default:null;"`
	UuuID     *string `json:"uuuId" gorm:"type:varchar(36) comment 'uuu_id';default:null;"`
	IsDelete  *int64  `json:"isDelete" gorm:"type:int(2) comment '是否删除：1/正常、2/删除';default:1;index:is_delete;"`
	Weight    *int64  `json:"weight" gorm:"type:int(11) comment '权重：用来排序';default:1;index:weight;"`
	State     *int64  `json:"state" gorm:"type:int(2) comment '状态：1/正常、2/禁用';default:1;index:state;"`
	DeletedBy *string `json:"deletedBy" gorm:"type:varchar(36) comment 'deleted_by';default:null;index:deleted_by;"`
	UpdatedBy *string `json:"updatedBy" gorm:"type:varchar(36) comment 'updated_by';default:null;index:updated_by;"`
	CreatedBy *string `json:"createdBy" gorm:"type:varchar(36) comment 'created_by';default:null;index:created_by;"`
	DeletedAt *int64  `json:"deletedAt" gorm:"type:bigint(13) comment 'deleted_at';default:null;"`
	UpdatedAt *int64  `json:"updatedAt" gorm:"type:bigint(13) comment 'updated_at';default:null; autoUpdateTime:milli;"`
	CreatedAt int64   `json:"createdAt" gorm:"type:bigint(13) comment 'created_at';default:null; autoCreateTime:milli;"`

	U *User `json:"u"`

	Uu []*User `json:"uu" gorm:"foreignkey:TtID"`

	Uuu *User `json:"uuu"`

	Uuuu []*User `json:"uuuu" gorm:"many2many:task_uuuu;jointable_foreignkey:tttt_id;association_jointable_foreignkey:uuuu_id"`
}

func (m *Task) Is_Entity() {}

type TaskChanges struct {
	ID        string
	Title     *string
	UID       *string
	UuuID     *string
	IsDelete  *int64
	Weight    *int64
	State     *int64
	DeletedBy *string
	UpdatedBy *string
	CreatedBy *string
	DeletedAt *int64
	UpdatedAt *int64
	CreatedAt int64

	U    *User
	Uu   []*User
	Uuu  *User
	Uuuu []*User

	UuIDs   []*string
	UuuuIDs []*string
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
