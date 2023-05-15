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

type OrderResultType struct {
	EntityResultType
}

type Order struct {
	ID           string  `json:"id" gorm:"type:varchar(36) comment 'uuid';primary_key;unique_index;NOT NULL;"`
	OrderNo      string  `json:"orderNo" gorm:"type:varchar(64) comment '訂單號';NOT NULL;index:order_no;"`
	CustomerInfo *string `json:"customerInfo" gorm:"type:varchar(255) comment '客戶信息';default:null;"`
	GoodsInfo    *string `json:"goodsInfo" gorm:"type:varchar(255) comment '貨物信息';default:null;"`
	DeletedBy    *string `json:"deletedBy" gorm:"type:varchar(36) comment 'deleted_by';default:null;index:deleted_by;"`
	UpdatedBy    *string `json:"updatedBy" gorm:"type:varchar(36) comment 'updated_by';default:null;index:updated_by;"`
	CreatedBy    *string `json:"createdBy" gorm:"type:varchar(36) comment 'created_by';default:null;index:created_by;"`
	DeletedAt    *int64  `json:"deletedAt" gorm:"type:bigint(13) comment 'deleted_at';default:null;"`
	UpdatedAt    *int64  `json:"updatedAt" gorm:"type:bigint(13) comment 'updated_at';default:null; autoUpdateTime:milli;"`
	CreatedAt    int64   `json:"createdAt" gorm:"type:bigint(13) comment 'created_at';default:null; autoCreateTime:milli;"`
}

func (m *Order) Is_Entity() {}

type OrderChanges struct {
	ID           string
	OrderNo      string
	CustomerInfo *string
	GoodsInfo    *string
	DeletedBy    *string
	UpdatedBy    *string
	CreatedBy    *string
	DeletedAt    *int64
	UpdatedAt    *int64
	CreatedAt    int64
}

type ShipmentResultType struct {
	EntityResultType
}

type Shipment struct {
	ID                 string  `json:"id" gorm:"type:varchar(36) comment 'uuid';primary_key;unique_index;NOT NULL;"`
	ShipmentNo         string  `json:"shipmentNo" gorm:"type:varchar(64) comment '運輸單號';NOT NULL;index:shipment_no;"`
	TransportationMode *string `json:"transportationMode" gorm:"type:varchar(32) comment '運輸方式';default:null;"`
	StartLocationID    *string `json:"startLocationId" gorm:"type:varchar(36) comment 'start_location_id';default:null;"`
	EndLocationID      *string `json:"endLocationId" gorm:"type:varchar(36) comment 'end_location_id';default:null;"`
	DeletedBy          *string `json:"deletedBy" gorm:"type:varchar(36) comment 'deleted_by';default:null;index:deleted_by;"`
	UpdatedBy          *string `json:"updatedBy" gorm:"type:varchar(36) comment 'updated_by';default:null;index:updated_by;"`
	CreatedBy          *string `json:"createdBy" gorm:"type:varchar(36) comment 'created_by';default:null;index:created_by;"`
	DeletedAt          *int64  `json:"deletedAt" gorm:"type:bigint(13) comment 'deleted_at';default:null;"`
	UpdatedAt          *int64  `json:"updatedAt" gorm:"type:bigint(13) comment 'updated_at';default:null; autoUpdateTime:milli;"`
	CreatedAt          int64   `json:"createdAt" gorm:"type:bigint(13) comment 'created_at';default:null; autoCreateTime:milli;"`

	StartLocation *Location `json:"startLocation"`

	EndLocation *Location `json:"endLocation"`
}

func (m *Shipment) Is_Entity() {}

type ShipmentChanges struct {
	ID                 string
	ShipmentNo         string
	TransportationMode *string
	StartLocationID    *string
	EndLocationID      *string
	DeletedBy          *string
	UpdatedBy          *string
	CreatedBy          *string
	DeletedAt          *int64
	UpdatedAt          *int64
	CreatedAt          int64

	StartLocation *Location
	EndLocation   *Location
}

type CarrierResultType struct {
	EntityResultType
}

type Carrier struct {
	ID            string  `json:"id" gorm:"type:varchar(36) comment 'uuid';primary_key;unique_index;NOT NULL;"`
	CarrierName   string  `json:"carrierName" gorm:"type:varchar(64) comment '承運商名稱';NOT NULL;index:name;"`
	ContactPerson *string `json:"contactPerson" gorm:"type:varchar(32) comment '聯系人';default:null;"`
	ContactInfo   *string `json:"contactInfo" gorm:"type:varchar(32) comment '聯系方式';default:null;"`
	DeletedBy     *string `json:"deletedBy" gorm:"type:varchar(36) comment 'deleted_by';default:null;index:deleted_by;"`
	UpdatedBy     *string `json:"updatedBy" gorm:"type:varchar(36) comment 'updated_by';default:null;index:updated_by;"`
	CreatedBy     *string `json:"createdBy" gorm:"type:varchar(36) comment 'created_by';default:null;index:created_by;"`
	DeletedAt     *int64  `json:"deletedAt" gorm:"type:bigint(13) comment 'deleted_at';default:null;"`
	UpdatedAt     *int64  `json:"updatedAt" gorm:"type:bigint(13) comment 'updated_at';default:null; autoUpdateTime:milli;"`
	CreatedAt     int64   `json:"createdAt" gorm:"type:bigint(13) comment 'created_at';default:null; autoCreateTime:milli;"`
}

func (m *Carrier) Is_Entity() {}

type CarrierChanges struct {
	ID            string
	CarrierName   string
	ContactPerson *string
	ContactInfo   *string
	DeletedBy     *string
	UpdatedBy     *string
	CreatedBy     *string
	DeletedAt     *int64
	UpdatedAt     *int64
	CreatedAt     int64
}

type LocationResultType struct {
	EntityResultType
}

type Location struct {
	ID               string  `json:"id" gorm:"type:varchar(36) comment 'uuid';primary_key;unique_index;NOT NULL;"`
	WarehouseAddress *string `json:"warehouseAddress" gorm:"type:varchar(255) comment '倉庫地址';default:null;"`
	LoadingAddress   *string `json:"loadingAddress" gorm:"type:varchar(255) comment '裝卸點地址';default:null;"`
	DeletedBy        *string `json:"deletedBy" gorm:"type:varchar(36) comment 'deleted_by';default:null;index:deleted_by;"`
	UpdatedBy        *string `json:"updatedBy" gorm:"type:varchar(36) comment 'updated_by';default:null;index:updated_by;"`
	CreatedBy        *string `json:"createdBy" gorm:"type:varchar(36) comment 'created_by';default:null;index:created_by;"`
	DeletedAt        *int64  `json:"deletedAt" gorm:"type:bigint(13) comment 'deleted_at';default:null;"`
	UpdatedAt        *int64  `json:"updatedAt" gorm:"type:bigint(13) comment 'updated_at';default:null; autoUpdateTime:milli;"`
	CreatedAt        int64   `json:"createdAt" gorm:"type:bigint(13) comment 'created_at';default:null; autoCreateTime:milli;"`

	StartShipments []*Shipment `json:"startShipments" gorm:"foreignkey:StartLocationID"`

	EndShipments []*Shipment `json:"endShipments" gorm:"foreignkey:EndLocationID"`
}

func (m *Location) Is_Entity() {}

type LocationChanges struct {
	ID               string
	WarehouseAddress *string
	LoadingAddress   *string
	DeletedBy        *string
	UpdatedBy        *string
	CreatedBy        *string
	DeletedAt        *int64
	UpdatedAt        *int64
	CreatedAt        int64

	StartShipments []*Shipment
	EndShipments   []*Shipment

	StartShipmentsIDs []*string
	EndShipmentsIDs   []*string
}

type EquipmentdResultType struct {
	EntityResultType
}

type Equipmentd struct {
	ID          string  `json:"id" gorm:"type:varchar(36) comment 'uuid';primary_key;unique_index;NOT NULL;"`
	VehicleType *string `json:"vehicleType" gorm:"type:varchar(32) comment '車型';default:null;"`
	Capacity    *int64  `json:"capacity" gorm:"type:int(13) comment '容量';default:null;" validator:"type:int;"`
	DeletedBy   *string `json:"deletedBy" gorm:"type:varchar(36) comment 'deleted_by';default:null;index:deleted_by;"`
	UpdatedBy   *string `json:"updatedBy" gorm:"type:varchar(36) comment 'updated_by';default:null;index:updated_by;"`
	CreatedBy   *string `json:"createdBy" gorm:"type:varchar(36) comment 'created_by';default:null;index:created_by;"`
	DeletedAt   *int64  `json:"deletedAt" gorm:"type:bigint(13) comment 'deleted_at';default:null;"`
	UpdatedAt   *int64  `json:"updatedAt" gorm:"type:bigint(13) comment 'updated_at';default:null; autoUpdateTime:milli;"`
	CreatedAt   int64   `json:"createdAt" gorm:"type:bigint(13) comment 'created_at';default:null; autoCreateTime:milli;"`
}

func (m *Equipmentd) Is_Entity() {}

type EquipmentdChanges struct {
	ID          string
	VehicleType *string
	Capacity    *int64
	DeletedBy   *string
	UpdatedBy   *string
	CreatedBy   *string
	DeletedAt   *int64
	UpdatedAt   *int64
	CreatedAt   int64
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
