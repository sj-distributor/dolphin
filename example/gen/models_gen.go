// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gen

import (
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
)

type FileField struct {
	Hash string         `json:"hash"`
	File graphql.Upload `json:"file"`
}

type TodoCreateRelationship struct {
	Title  string  `json:"title"`
	Age    *int    `json:"age,omitempty"`
	Money  int     `json:"money"`
	Remark *string `json:"remark,omitempty"`
	UserID *string `json:"userId,omitempty"`
}

type TodoFilterType struct {
	And           []*TodoFilterType `json:"AND,omitempty"`
	Or            []*TodoFilterType `json:"OR,omitempty"`
	ID            *string           `json:"id,omitempty"`
	IDNe          *string           `json:"id_ne,omitempty"`
	IDGt          *string           `json:"id_gt,omitempty"`
	IDLt          *string           `json:"id_lt,omitempty"`
	IDGte         *string           `json:"id_gte,omitempty"`
	IDLte         *string           `json:"id_lte,omitempty"`
	IDIn          []string          `json:"id_in,omitempty"`
	IDNull        *bool             `json:"id_null,omitempty"`
	Title         *string           `json:"title,omitempty"`
	TitleNe       *string           `json:"title_ne,omitempty"`
	TitleGt       *string           `json:"title_gt,omitempty"`
	TitleLt       *string           `json:"title_lt,omitempty"`
	TitleGte      *string           `json:"title_gte,omitempty"`
	TitleLte      *string           `json:"title_lte,omitempty"`
	TitleIn       []string          `json:"title_in,omitempty"`
	TitleLike     *string           `json:"title_like,omitempty"`
	TitlePrefix   *string           `json:"title_prefix,omitempty"`
	TitleSuffix   *string           `json:"title_suffix,omitempty"`
	TitleNull     *bool             `json:"title_null,omitempty"`
	Age           *int              `json:"age,omitempty"`
	AgeNe         *int              `json:"age_ne,omitempty"`
	AgeGt         *int              `json:"age_gt,omitempty"`
	AgeLt         *int              `json:"age_lt,omitempty"`
	AgeGte        *int              `json:"age_gte,omitempty"`
	AgeLte        *int              `json:"age_lte,omitempty"`
	AgeIn         []int             `json:"age_in,omitempty"`
	AgeNull       *bool             `json:"age_null,omitempty"`
	Money         *int              `json:"money,omitempty"`
	MoneyNe       *int              `json:"money_ne,omitempty"`
	MoneyGt       *int              `json:"money_gt,omitempty"`
	MoneyLt       *int              `json:"money_lt,omitempty"`
	MoneyGte      *int              `json:"money_gte,omitempty"`
	MoneyLte      *int              `json:"money_lte,omitempty"`
	MoneyIn       []int             `json:"money_in,omitempty"`
	MoneyNull     *bool             `json:"money_null,omitempty"`
	Remark        *string           `json:"remark,omitempty"`
	RemarkNe      *string           `json:"remark_ne,omitempty"`
	RemarkGt      *string           `json:"remark_gt,omitempty"`
	RemarkLt      *string           `json:"remark_lt,omitempty"`
	RemarkGte     *string           `json:"remark_gte,omitempty"`
	RemarkLte     *string           `json:"remark_lte,omitempty"`
	RemarkIn      []string          `json:"remark_in,omitempty"`
	RemarkLike    *string           `json:"remark_like,omitempty"`
	RemarkPrefix  *string           `json:"remark_prefix,omitempty"`
	RemarkSuffix  *string           `json:"remark_suffix,omitempty"`
	RemarkNull    *bool             `json:"remark_null,omitempty"`
	UserID        *string           `json:"userId,omitempty"`
	UserIDNe      *string           `json:"userId_ne,omitempty"`
	UserIDGt      *string           `json:"userId_gt,omitempty"`
	UserIDLt      *string           `json:"userId_lt,omitempty"`
	UserIDGte     *string           `json:"userId_gte,omitempty"`
	UserIDLte     *string           `json:"userId_lte,omitempty"`
	UserIDIn      []string          `json:"userId_in,omitempty"`
	UserIDNull    *bool             `json:"userId_null,omitempty"`
	DeletedBy     *string           `json:"deletedBy,omitempty"`
	DeletedByNe   *string           `json:"deletedBy_ne,omitempty"`
	DeletedByGt   *string           `json:"deletedBy_gt,omitempty"`
	DeletedByLt   *string           `json:"deletedBy_lt,omitempty"`
	DeletedByGte  *string           `json:"deletedBy_gte,omitempty"`
	DeletedByLte  *string           `json:"deletedBy_lte,omitempty"`
	DeletedByIn   []string          `json:"deletedBy_in,omitempty"`
	DeletedByNull *bool             `json:"deletedBy_null,omitempty"`
	UpdatedBy     *string           `json:"updatedBy,omitempty"`
	UpdatedByNe   *string           `json:"updatedBy_ne,omitempty"`
	UpdatedByGt   *string           `json:"updatedBy_gt,omitempty"`
	UpdatedByLt   *string           `json:"updatedBy_lt,omitempty"`
	UpdatedByGte  *string           `json:"updatedBy_gte,omitempty"`
	UpdatedByLte  *string           `json:"updatedBy_lte,omitempty"`
	UpdatedByIn   []string          `json:"updatedBy_in,omitempty"`
	UpdatedByNull *bool             `json:"updatedBy_null,omitempty"`
	CreatedBy     *string           `json:"createdBy,omitempty"`
	CreatedByNe   *string           `json:"createdBy_ne,omitempty"`
	CreatedByGt   *string           `json:"createdBy_gt,omitempty"`
	CreatedByLt   *string           `json:"createdBy_lt,omitempty"`
	CreatedByGte  *string           `json:"createdBy_gte,omitempty"`
	CreatedByLte  *string           `json:"createdBy_lte,omitempty"`
	CreatedByIn   []string          `json:"createdBy_in,omitempty"`
	CreatedByNull *bool             `json:"createdBy_null,omitempty"`
	DeletedAt     *int              `json:"deletedAt,omitempty"`
	DeletedAtNe   *int              `json:"deletedAt_ne,omitempty"`
	DeletedAtGt   *int              `json:"deletedAt_gt,omitempty"`
	DeletedAtLt   *int              `json:"deletedAt_lt,omitempty"`
	DeletedAtGte  *int              `json:"deletedAt_gte,omitempty"`
	DeletedAtLte  *int              `json:"deletedAt_lte,omitempty"`
	DeletedAtIn   []int             `json:"deletedAt_in,omitempty"`
	DeletedAtNull *bool             `json:"deletedAt_null,omitempty"`
	UpdatedAt     *int              `json:"updatedAt,omitempty"`
	UpdatedAtNe   *int              `json:"updatedAt_ne,omitempty"`
	UpdatedAtGt   *int              `json:"updatedAt_gt,omitempty"`
	UpdatedAtLt   *int              `json:"updatedAt_lt,omitempty"`
	UpdatedAtGte  *int              `json:"updatedAt_gte,omitempty"`
	UpdatedAtLte  *int              `json:"updatedAt_lte,omitempty"`
	UpdatedAtIn   []int             `json:"updatedAt_in,omitempty"`
	UpdatedAtNull *bool             `json:"updatedAt_null,omitempty"`
	CreatedAt     *int              `json:"createdAt,omitempty"`
	CreatedAtNe   *int              `json:"createdAt_ne,omitempty"`
	CreatedAtGt   *int              `json:"createdAt_gt,omitempty"`
	CreatedAtLt   *int              `json:"createdAt_lt,omitempty"`
	CreatedAtGte  *int              `json:"createdAt_gte,omitempty"`
	CreatedAtLte  *int              `json:"createdAt_lte,omitempty"`
	CreatedAtIn   []int             `json:"createdAt_in,omitempty"`
	CreatedAtNull *bool             `json:"createdAt_null,omitempty"`
	User          *UserFilterType   `json:"user,omitempty"`
}

type TodoSortType struct {
	ID        *ObjectSortType `json:"id,omitempty"`
	Title     *ObjectSortType `json:"title,omitempty"`
	Age       *ObjectSortType `json:"age,omitempty"`
	Money     *ObjectSortType `json:"money,omitempty"`
	Remark    *ObjectSortType `json:"remark,omitempty"`
	UserID    *ObjectSortType `json:"userId,omitempty"`
	DeletedBy *ObjectSortType `json:"deletedBy,omitempty"`
	UpdatedBy *ObjectSortType `json:"updatedBy,omitempty"`
	CreatedBy *ObjectSortType `json:"createdBy,omitempty"`
	DeletedAt *ObjectSortType `json:"deletedAt,omitempty"`
	UpdatedAt *ObjectSortType `json:"updatedAt,omitempty"`
	CreatedAt *ObjectSortType `json:"createdAt,omitempty"`
	User      *UserSortType   `json:"user,omitempty"`
}

type TodoUpdateRelationship struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Age    *int    `json:"age,omitempty"`
	Money  int     `json:"money"`
	Remark *string `json:"remark,omitempty"`
	UserID *string `json:"userId,omitempty"`
}

type UserCreateRelationship struct {
	Username string  `json:"username"`
	TodoID   *string `json:"todoId,omitempty"`
}

type UserFilterType struct {
	And            []*UserFilterType `json:"AND,omitempty"`
	Or             []*UserFilterType `json:"OR,omitempty"`
	ID             *string           `json:"id,omitempty"`
	IDNe           *string           `json:"id_ne,omitempty"`
	IDGt           *string           `json:"id_gt,omitempty"`
	IDLt           *string           `json:"id_lt,omitempty"`
	IDGte          *string           `json:"id_gte,omitempty"`
	IDLte          *string           `json:"id_lte,omitempty"`
	IDIn           []string          `json:"id_in,omitempty"`
	IDNull         *bool             `json:"id_null,omitempty"`
	Username       *string           `json:"username,omitempty"`
	UsernameNe     *string           `json:"username_ne,omitempty"`
	UsernameGt     *string           `json:"username_gt,omitempty"`
	UsernameLt     *string           `json:"username_lt,omitempty"`
	UsernameGte    *string           `json:"username_gte,omitempty"`
	UsernameLte    *string           `json:"username_lte,omitempty"`
	UsernameIn     []string          `json:"username_in,omitempty"`
	UsernameLike   *string           `json:"username_like,omitempty"`
	UsernamePrefix *string           `json:"username_prefix,omitempty"`
	UsernameSuffix *string           `json:"username_suffix,omitempty"`
	UsernameNull   *bool             `json:"username_null,omitempty"`
	TodoID         *string           `json:"todoId,omitempty"`
	TodoIDNe       *string           `json:"todoId_ne,omitempty"`
	TodoIDGt       *string           `json:"todoId_gt,omitempty"`
	TodoIDLt       *string           `json:"todoId_lt,omitempty"`
	TodoIDGte      *string           `json:"todoId_gte,omitempty"`
	TodoIDLte      *string           `json:"todoId_lte,omitempty"`
	TodoIDIn       []string          `json:"todoId_in,omitempty"`
	TodoIDNull     *bool             `json:"todoId_null,omitempty"`
	DeletedBy      *string           `json:"deletedBy,omitempty"`
	DeletedByNe    *string           `json:"deletedBy_ne,omitempty"`
	DeletedByGt    *string           `json:"deletedBy_gt,omitempty"`
	DeletedByLt    *string           `json:"deletedBy_lt,omitempty"`
	DeletedByGte   *string           `json:"deletedBy_gte,omitempty"`
	DeletedByLte   *string           `json:"deletedBy_lte,omitempty"`
	DeletedByIn    []string          `json:"deletedBy_in,omitempty"`
	DeletedByNull  *bool             `json:"deletedBy_null,omitempty"`
	UpdatedBy      *string           `json:"updatedBy,omitempty"`
	UpdatedByNe    *string           `json:"updatedBy_ne,omitempty"`
	UpdatedByGt    *string           `json:"updatedBy_gt,omitempty"`
	UpdatedByLt    *string           `json:"updatedBy_lt,omitempty"`
	UpdatedByGte   *string           `json:"updatedBy_gte,omitempty"`
	UpdatedByLte   *string           `json:"updatedBy_lte,omitempty"`
	UpdatedByIn    []string          `json:"updatedBy_in,omitempty"`
	UpdatedByNull  *bool             `json:"updatedBy_null,omitempty"`
	CreatedBy      *string           `json:"createdBy,omitempty"`
	CreatedByNe    *string           `json:"createdBy_ne,omitempty"`
	CreatedByGt    *string           `json:"createdBy_gt,omitempty"`
	CreatedByLt    *string           `json:"createdBy_lt,omitempty"`
	CreatedByGte   *string           `json:"createdBy_gte,omitempty"`
	CreatedByLte   *string           `json:"createdBy_lte,omitempty"`
	CreatedByIn    []string          `json:"createdBy_in,omitempty"`
	CreatedByNull  *bool             `json:"createdBy_null,omitempty"`
	DeletedAt      *int              `json:"deletedAt,omitempty"`
	DeletedAtNe    *int              `json:"deletedAt_ne,omitempty"`
	DeletedAtGt    *int              `json:"deletedAt_gt,omitempty"`
	DeletedAtLt    *int              `json:"deletedAt_lt,omitempty"`
	DeletedAtGte   *int              `json:"deletedAt_gte,omitempty"`
	DeletedAtLte   *int              `json:"deletedAt_lte,omitempty"`
	DeletedAtIn    []int             `json:"deletedAt_in,omitempty"`
	DeletedAtNull  *bool             `json:"deletedAt_null,omitempty"`
	UpdatedAt      *int              `json:"updatedAt,omitempty"`
	UpdatedAtNe    *int              `json:"updatedAt_ne,omitempty"`
	UpdatedAtGt    *int              `json:"updatedAt_gt,omitempty"`
	UpdatedAtLt    *int              `json:"updatedAt_lt,omitempty"`
	UpdatedAtGte   *int              `json:"updatedAt_gte,omitempty"`
	UpdatedAtLte   *int              `json:"updatedAt_lte,omitempty"`
	UpdatedAtIn    []int             `json:"updatedAt_in,omitempty"`
	UpdatedAtNull  *bool             `json:"updatedAt_null,omitempty"`
	CreatedAt      *int              `json:"createdAt,omitempty"`
	CreatedAtNe    *int              `json:"createdAt_ne,omitempty"`
	CreatedAtGt    *int              `json:"createdAt_gt,omitempty"`
	CreatedAtLt    *int              `json:"createdAt_lt,omitempty"`
	CreatedAtGte   *int              `json:"createdAt_gte,omitempty"`
	CreatedAtLte   *int              `json:"createdAt_lte,omitempty"`
	CreatedAtIn    []int             `json:"createdAt_in,omitempty"`
	CreatedAtNull  *bool             `json:"createdAt_null,omitempty"`
	Todo           *TodoFilterType   `json:"todo,omitempty"`
}

type UserSortType struct {
	ID        *ObjectSortType `json:"id,omitempty"`
	Username  *ObjectSortType `json:"username,omitempty"`
	TodoID    *ObjectSortType `json:"todoId,omitempty"`
	DeletedBy *ObjectSortType `json:"deletedBy,omitempty"`
	UpdatedBy *ObjectSortType `json:"updatedBy,omitempty"`
	CreatedBy *ObjectSortType `json:"createdBy,omitempty"`
	DeletedAt *ObjectSortType `json:"deletedAt,omitempty"`
	UpdatedAt *ObjectSortType `json:"updatedAt,omitempty"`
	CreatedAt *ObjectSortType `json:"createdAt,omitempty"`
	Todo      *TodoSortType   `json:"todo,omitempty"`
}

type UserUpdateRelationship struct {
	ID       string  `json:"id"`
	Username string  `json:"username"`
	TodoID   *string `json:"todoId,omitempty"`
}

type ObjectSortType string

const (
	ObjectSortTypeAsc  ObjectSortType = "ASC"
	ObjectSortTypeDesc ObjectSortType = "DESC"
)

var AllObjectSortType = []ObjectSortType{
	ObjectSortTypeAsc,
	ObjectSortTypeDesc,
}

func (e ObjectSortType) IsValid() bool {
	switch e {
	case ObjectSortTypeAsc, ObjectSortTypeDesc:
		return true
	}
	return false
}

func (e ObjectSortType) String() string {
	return string(e)
}

func (e *ObjectSortType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ObjectSortType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ObjectSortType", str)
	}
	return nil
}

func (e ObjectSortType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
