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
	Username  string  `json:"username" gorm:"type:varchar(32) comment '用戶名';NOT NULL;unique_index:username;" validator:"type:username;"`
	Password  string  `json:"password" gorm:"type:varchar(64) comment '登錄密碼';NOT NULL;" validator:"type:password;"`
	Email     *string `json:"email" gorm:"type:varchar(64) comment '用戶郵箱地址';default:null;" validator:"type:email;"`
	Nickname  *string `json:"nickname" gorm:"type:varchar(64) comment '暱稱';DEFAULT NULL;index:nickname;"`
	DeletedBy *string `json:"deletedBy" gorm:"type:varchar(36) comment 'deleted_by';default:null;index:deleted_by;"`
	UpdatedBy *string `json:"updatedBy" gorm:"type:varchar(36) comment 'updated_by';default:null;index:updated_by;"`
	CreatedBy *string `json:"createdBy" gorm:"type:varchar(36) comment 'created_by';default:null;index:created_by;"`
	DeletedAt *int64  `json:"deletedAt" gorm:"type:bigint(13) comment 'deleted_at';default:null;"`
	UpdatedAt *int64  `json:"updatedAt" gorm:"type:bigint(13) comment 'updated_at';default:null; autoUpdateTime:milli;"`
	CreatedAt int64   `json:"createdAt" gorm:"type:bigint(13) comment 'created_at';default:null; autoCreateTime:milli;"`

	Accounts []*Account `json:"accounts" gorm:"foreignkey:OwnerID"`

	Todo []*Todo `json:"todo" gorm:"foreignkey:AccountID"`
}

func (m *User) Is_Entity() {}

type UserChanges struct {
	ID        string
	Username  string
	Password  string
	Email     *string
	Nickname  *string
	DeletedBy *string
	UpdatedBy *string
	CreatedBy *string
	DeletedAt *int64
	UpdatedAt *int64
	CreatedAt int64

	Accounts []*Account
	Todo     []*Todo

	AccountsIDs []*string
	TodoIDs     []*string
}

type AccountResultType struct {
	EntityResultType
}

type Account struct {
	ID        string  `json:"id" gorm:"type:varchar(36) comment 'uuid';primary_key;unique_index;NOT NULL;"`
	Name      string  `json:"name" gorm:"type varchar(32) comment '賬戶名稱';NOT NULL;"`
	Balance   int64   `json:"balance" gorm:"type int(13) comment '賬戶餘額';default:0;" validator:"type:int;"`
	OwnerID   *string `json:"ownerId" gorm:"type:varchar(36) comment 'owner_id';default:null;"`
	DeletedBy *string `json:"deletedBy" gorm:"type:varchar(36) comment 'deleted_by';default:null;index:deleted_by;"`
	UpdatedBy *string `json:"updatedBy" gorm:"type:varchar(36) comment 'updated_by';default:null;index:updated_by;"`
	CreatedBy *string `json:"createdBy" gorm:"type:varchar(36) comment 'created_by';default:null;index:created_by;"`
	DeletedAt *int64  `json:"deletedAt" gorm:"type:bigint(13) comment 'deleted_at';default:null;"`
	UpdatedAt *int64  `json:"updatedAt" gorm:"type:bigint(13) comment 'updated_at';default:null; autoUpdateTime:milli;"`
	CreatedAt int64   `json:"createdAt" gorm:"type:bigint(13) comment 'created_at';default:null; autoCreateTime:milli;"`

	Owner *User `json:"owner"`

	Transactions []*Transaction `json:"transactions" gorm:"foreignkey:AccountID"`
}

func (m *Account) Is_Entity() {}

type AccountChanges struct {
	ID        string
	Name      string
	Balance   int64
	OwnerID   *string
	DeletedBy *string
	UpdatedBy *string
	CreatedBy *string
	DeletedAt *int64
	UpdatedAt *int64
	CreatedAt int64

	Owner        *User
	Transactions []*Transaction

	TransactionsIDs []*string
}

type TransactionResultType struct {
	EntityResultType
}

type Transaction struct {
	ID        string  `json:"id" gorm:"type:varchar(36) comment 'uuid';primary_key;unique_index;NOT NULL;"`
	Amount    int64   `json:"amount" gorm:"type int(13) comment '交易金額';NOT NULL;" validator:"type:int;"`
	Date      int64   `json:"date" gorm:"type int(13) comment '交易日期';NOT NULL;" validator:"type:int;"`
	Note      *string `json:"note" gorm:"type varchar(255) comment '備注信息';default:null;"`
	AccountID *string `json:"accountId" gorm:"type:varchar(36) comment 'account_id';default:null;"`
	DeletedBy *string `json:"deletedBy" gorm:"type:varchar(36) comment 'deleted_by';default:null;index:deleted_by;"`
	UpdatedBy *string `json:"updatedBy" gorm:"type:varchar(36) comment 'updated_by';default:null;index:updated_by;"`
	CreatedBy *string `json:"createdBy" gorm:"type:varchar(36) comment 'created_by';default:null;index:created_by;"`
	DeletedAt *int64  `json:"deletedAt" gorm:"type:bigint(13) comment 'deleted_at';default:null;"`
	UpdatedAt *int64  `json:"updatedAt" gorm:"type:bigint(13) comment 'updated_at';default:null; autoUpdateTime:milli;"`
	CreatedAt int64   `json:"createdAt" gorm:"type:bigint(13) comment 'created_at';default:null; autoCreateTime:milli;"`

	Account *Account `json:"account"`
}

func (m *Transaction) Is_Entity() {}

type TransactionChanges struct {
	ID        string
	Amount    int64
	Date      int64
	Note      *string
	AccountID *string
	DeletedBy *string
	UpdatedBy *string
	CreatedBy *string
	DeletedAt *int64
	UpdatedAt *int64
	CreatedAt int64

	Account *Account
}

type TodoResultType struct {
	EntityResultType
}

type Todo struct {
	ID        string  `json:"id" gorm:"type:varchar(36) comment 'uuid';primary_key;unique_index;NOT NULL;"`
	Name      string  `json:"name" gorm:"type varchar(32) comment 'todo名稱';NOT NULL;"`
	AccountID *string `json:"accountId" gorm:"type:varchar(36) comment 'account_id';default:null;"`
	DeletedBy *string `json:"deletedBy" gorm:"type:varchar(36) comment 'deleted_by';default:null;index:deleted_by;"`
	UpdatedBy *string `json:"updatedBy" gorm:"type:varchar(36) comment 'updated_by';default:null;index:updated_by;"`
	CreatedBy *string `json:"createdBy" gorm:"type:varchar(36) comment 'created_by';default:null;index:created_by;"`
	DeletedAt *int64  `json:"deletedAt" gorm:"type:bigint(13) comment 'deleted_at';default:null;"`
	UpdatedAt *int64  `json:"updatedAt" gorm:"type:bigint(13) comment 'updated_at';default:null; autoUpdateTime:milli;"`
	CreatedAt int64   `json:"createdAt" gorm:"type:bigint(13) comment 'created_at';default:null; autoCreateTime:milli;"`

	Account *User `json:"account"`
}

func (m *Todo) Is_Entity() {}

type TodoChanges struct {
	ID        string
	Name      string
	AccountID *string
	DeletedBy *string
	UpdatedBy *string
	CreatedBy *string
	DeletedAt *int64
	UpdatedAt *int64
	CreatedAt int64

	Account *User
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
