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
	Title     string  `json:"title" gorm:"default:null" validator:"type:password;"`
	Age       *int64  `json:"age" gorm:"default:null"`
	Money     int64   `json:"money" gorm:"default:null" validator:"max:2;"`
	Remark    *string `json:"remark" gorm:"default:null" validator:"type:password;len:12;"`
	DeletedBy *string `json:"deletedBy" gorm:"type:varchar(36) comment 'deletedBy';default:null;"`
	UpdatedBy *string `json:"updatedBy" gorm:"type:varchar(36) comment 'updatedBy';default:null;"`
	CreatedBy *string `json:"createdBy" gorm:"type:varchar(36) comment 'createdBy';default:null;"`
	DeletedAt *int64  `json:"deletedAt" gorm:"type:bigint(13) comment 'deletedAt';default:null;"`
	UpdatedAt *int64  `json:"updatedAt" gorm:"type:bigint(13) comment 'updatedAt';default:null;"`
	CreatedAt int64   `json:"createdAt" gorm:"type:bigint(13) comment 'createdAt';default:null;index:created_at;"`
}

func (m *Todo) Is_Entity() {}

type TodoChanges struct {
	ID        string
	Title     string
	Age       *int64
	Money     int64
	Remark    *string
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
					return v, fmt.Errorf("unable to parse date from %v", v)
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
