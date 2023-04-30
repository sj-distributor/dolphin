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

type BookCategoryResultType struct {
	EntityResultType
}

type BookCategory struct {
	ID          string  `json:"id" gorm:"type:varchar(36) comment 'uuid';primary_key;unique_index;NOT NULL;"`
	Name        string  `json:"name" gorm:"type:varchar(64) comment '分類名稱';NOT NULL;"`
	Description *string `json:"description" gorm:"type:text comment '分類描述';DEFAULT NULL;"`
	DeletedBy   *string `json:"deletedBy" gorm:"type:varchar(36) comment 'deleted_by';default:null;index:deleted_by;"`
	UpdatedBy   *string `json:"updatedBy" gorm:"type:varchar(36) comment 'updated_by';default:null;index:updated_by;"`
	CreatedBy   *string `json:"createdBy" gorm:"type:varchar(36) comment 'created_by';default:null;index:created_by;"`
	DeletedAt   *int64  `json:"deletedAt" gorm:"type:bigint(13) comment 'deleted_at';default:null;"`
	UpdatedAt   *int64  `json:"updatedAt" gorm:"type:bigint(13) comment 'updated_at';default:null; autoUpdateTime:milli;"`
	CreatedAt   int64   `json:"createdAt" gorm:"type:bigint(13) comment 'created_at';default:null; autoCreateTime:milli;"`

	Books []*Book `json:"books" gorm:"foreignkey:CategoryID"`
}

func (m *BookCategory) Is_Entity() {}

type BookCategoryChanges struct {
	ID          string
	Name        string
	Description *string
	DeletedBy   *string
	UpdatedBy   *string
	CreatedBy   *string
	DeletedAt   *int64
	UpdatedAt   *int64
	CreatedAt   int64

	Books []*Book

	BooksIDs []*string
}

type BookResultType struct {
	EntityResultType
}

type Book struct {
	ID            string  `json:"id" gorm:"type:varchar(36) comment 'uuid';primary_key;unique_index;NOT NULL;"`
	Title         string  `json:"title" gorm:"type:varchar(128) comment '圖書名稱';NOT NULL;"`
	Author        string  `json:"author" gorm:"type:varchar(64) comment '作者';NOT NULL;"`
	Price         *int64  `json:"price" gorm:"type:int(13) comment '價格';DEFAULT NULL;" validator:"type:int;"`
	PublishDateAt *int64  `json:"publishDateAt" gorm:"type:int(13) comment '出版日期';DEFAULT NULL;" validator:"type:int;"`
	CategoryID    *string `json:"categoryId" gorm:"type:varchar(36) comment 'category_id';default:null;"`
	DeletedBy     *string `json:"deletedBy" gorm:"type:varchar(36) comment 'deleted_by';default:null;index:deleted_by;"`
	UpdatedBy     *string `json:"updatedBy" gorm:"type:varchar(36) comment 'updated_by';default:null;index:updated_by;"`
	CreatedBy     *string `json:"createdBy" gorm:"type:varchar(36) comment 'created_by';default:null;index:created_by;"`
	DeletedAt     *int64  `json:"deletedAt" gorm:"type:bigint(13) comment 'deleted_at';default:null;"`
	UpdatedAt     *int64  `json:"updatedAt" gorm:"type:bigint(13) comment 'updated_at';default:null; autoUpdateTime:milli;"`
	CreatedAt     int64   `json:"createdAt" gorm:"type:bigint(13) comment 'created_at';default:null; autoCreateTime:milli;"`

	Category *BookCategory `json:"category"`
}

func (m *Book) Is_Entity() {}

type BookChanges struct {
	ID            string
	Title         string
	Author        string
	Price         *int64
	PublishDateAt *int64
	CategoryID    *string
	DeletedBy     *string
	UpdatedBy     *string
	CreatedBy     *string
	DeletedAt     *int64
	UpdatedAt     *int64
	CreatedAt     int64

	Category *BookCategory
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
