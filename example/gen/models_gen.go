// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gen

import (
	"fmt"
	"io"
	"strconv"
)

type TodoCreateRelationship struct {
	Title  string                         `json:"title"`
	Age    *int                           `json:"age"`
	Money  int                            `json:"money"`
	Remark *string                        `json:"remark"`
	User   *UserCreateReverseRelationship `json:"user"`
	UserID *string                        `json:"userId"`
}

type TodoCreateReverseRelationship struct {
	Title  string  `json:"title"`
	Age    *int    `json:"age"`
	Money  int     `json:"money"`
	Remark *string `json:"remark"`
	UserID *string `json:"userId"`
}

type TodoFilterType struct {
	And           []*TodoFilterType `json:"AND"`
	Or            []*TodoFilterType `json:"OR"`
	ID            *string           `json:"id"`
	IDNe          *string           `json:"id_ne"`
	IDGt          *string           `json:"id_gt"`
	IDLt          *string           `json:"id_lt"`
	IDGte         *string           `json:"id_gte"`
	IDLte         *string           `json:"id_lte"`
	IDIn          []string          `json:"id_in"`
	IDNull        *bool             `json:"id_null"`
	Title         *string           `json:"title"`
	TitleNe       *string           `json:"title_ne"`
	TitleGt       *string           `json:"title_gt"`
	TitleLt       *string           `json:"title_lt"`
	TitleGte      *string           `json:"title_gte"`
	TitleLte      *string           `json:"title_lte"`
	TitleIn       []string          `json:"title_in"`
	TitleLike     *string           `json:"title_like"`
	TitlePrefix   *string           `json:"title_prefix"`
	TitleSuffix   *string           `json:"title_suffix"`
	TitleNull     *bool             `json:"title_null"`
	Age           *int              `json:"age"`
	AgeNe         *int              `json:"age_ne"`
	AgeGt         *int              `json:"age_gt"`
	AgeLt         *int              `json:"age_lt"`
	AgeGte        *int              `json:"age_gte"`
	AgeLte        *int              `json:"age_lte"`
	AgeIn         []int             `json:"age_in"`
	AgeNull       *bool             `json:"age_null"`
	Money         *int              `json:"money"`
	MoneyNe       *int              `json:"money_ne"`
	MoneyGt       *int              `json:"money_gt"`
	MoneyLt       *int              `json:"money_lt"`
	MoneyGte      *int              `json:"money_gte"`
	MoneyLte      *int              `json:"money_lte"`
	MoneyIn       []int             `json:"money_in"`
	MoneyNull     *bool             `json:"money_null"`
	Remark        *string           `json:"remark"`
	RemarkNe      *string           `json:"remark_ne"`
	RemarkGt      *string           `json:"remark_gt"`
	RemarkLt      *string           `json:"remark_lt"`
	RemarkGte     *string           `json:"remark_gte"`
	RemarkLte     *string           `json:"remark_lte"`
	RemarkIn      []string          `json:"remark_in"`
	RemarkLike    *string           `json:"remark_like"`
	RemarkPrefix  *string           `json:"remark_prefix"`
	RemarkSuffix  *string           `json:"remark_suffix"`
	RemarkNull    *bool             `json:"remark_null"`
	UserID        *string           `json:"userId"`
	UserIDNe      *string           `json:"userId_ne"`
	UserIDGt      *string           `json:"userId_gt"`
	UserIDLt      *string           `json:"userId_lt"`
	UserIDGte     *string           `json:"userId_gte"`
	UserIDLte     *string           `json:"userId_lte"`
	UserIDIn      []string          `json:"userId_in"`
	UserIDNull    *bool             `json:"userId_null"`
	DeletedBy     *string           `json:"deletedBy"`
	DeletedByNe   *string           `json:"deletedBy_ne"`
	DeletedByGt   *string           `json:"deletedBy_gt"`
	DeletedByLt   *string           `json:"deletedBy_lt"`
	DeletedByGte  *string           `json:"deletedBy_gte"`
	DeletedByLte  *string           `json:"deletedBy_lte"`
	DeletedByIn   []string          `json:"deletedBy_in"`
	DeletedByNull *bool             `json:"deletedBy_null"`
	UpdatedBy     *string           `json:"updatedBy"`
	UpdatedByNe   *string           `json:"updatedBy_ne"`
	UpdatedByGt   *string           `json:"updatedBy_gt"`
	UpdatedByLt   *string           `json:"updatedBy_lt"`
	UpdatedByGte  *string           `json:"updatedBy_gte"`
	UpdatedByLte  *string           `json:"updatedBy_lte"`
	UpdatedByIn   []string          `json:"updatedBy_in"`
	UpdatedByNull *bool             `json:"updatedBy_null"`
	CreatedBy     *string           `json:"createdBy"`
	CreatedByNe   *string           `json:"createdBy_ne"`
	CreatedByGt   *string           `json:"createdBy_gt"`
	CreatedByLt   *string           `json:"createdBy_lt"`
	CreatedByGte  *string           `json:"createdBy_gte"`
	CreatedByLte  *string           `json:"createdBy_lte"`
	CreatedByIn   []string          `json:"createdBy_in"`
	CreatedByNull *bool             `json:"createdBy_null"`
	DeletedAt     *int              `json:"deletedAt"`
	DeletedAtNe   *int              `json:"deletedAt_ne"`
	DeletedAtGt   *int              `json:"deletedAt_gt"`
	DeletedAtLt   *int              `json:"deletedAt_lt"`
	DeletedAtGte  *int              `json:"deletedAt_gte"`
	DeletedAtLte  *int              `json:"deletedAt_lte"`
	DeletedAtIn   []int             `json:"deletedAt_in"`
	DeletedAtNull *bool             `json:"deletedAt_null"`
	UpdatedAt     *int              `json:"updatedAt"`
	UpdatedAtNe   *int              `json:"updatedAt_ne"`
	UpdatedAtGt   *int              `json:"updatedAt_gt"`
	UpdatedAtLt   *int              `json:"updatedAt_lt"`
	UpdatedAtGte  *int              `json:"updatedAt_gte"`
	UpdatedAtLte  *int              `json:"updatedAt_lte"`
	UpdatedAtIn   []int             `json:"updatedAt_in"`
	UpdatedAtNull *bool             `json:"updatedAt_null"`
	CreatedAt     *int              `json:"createdAt"`
	CreatedAtNe   *int              `json:"createdAt_ne"`
	CreatedAtGt   *int              `json:"createdAt_gt"`
	CreatedAtLt   *int              `json:"createdAt_lt"`
	CreatedAtGte  *int              `json:"createdAt_gte"`
	CreatedAtLte  *int              `json:"createdAt_lte"`
	CreatedAtIn   []int             `json:"createdAt_in"`
	CreatedAtNull *bool             `json:"createdAt_null"`
	User          *UserFilterType   `json:"user"`
}

type TodoSortType struct {
	ID        *ObjectSortType `json:"id"`
	Title     *ObjectSortType `json:"title"`
	Age       *ObjectSortType `json:"age"`
	Money     *ObjectSortType `json:"money"`
	Remark    *ObjectSortType `json:"remark"`
	UserID    *ObjectSortType `json:"userId"`
	DeletedBy *ObjectSortType `json:"deletedBy"`
	UpdatedBy *ObjectSortType `json:"updatedBy"`
	CreatedBy *ObjectSortType `json:"createdBy"`
	DeletedAt *ObjectSortType `json:"deletedAt"`
	UpdatedAt *ObjectSortType `json:"updatedAt"`
	CreatedAt *ObjectSortType `json:"createdAt"`
	User      *UserSortType   `json:"user"`
}

type TodoUpdateRelationship struct {
	ID     *string                        `json:"id"`
	Title  *string                        `json:"title"`
	Age    *int                           `json:"age"`
	Money  *int                           `json:"money"`
	Remark *string                        `json:"remark"`
	User   *UserUpdateReverseRelationship `json:"user"`
	UserID *string                        `json:"userId"`
}

type TodoUpdateReverseRelationship struct {
	ID     *string `json:"id"`
	Title  string  `json:"title"`
	Age    *int    `json:"age"`
	Money  int     `json:"money"`
	Remark *string `json:"remark"`
	UserID *string `json:"userId"`
}

type UserCreateRelationship struct {
	Username string                         `json:"username"`
	Todo     *TodoCreateReverseRelationship `json:"todo"`
	TodoID   *string                        `json:"todoId"`
}

type UserCreateReverseRelationship struct {
	Username string  `json:"username"`
	TodoID   *string `json:"todoId"`
}

type UserFilterType struct {
	And            []*UserFilterType `json:"AND"`
	Or             []*UserFilterType `json:"OR"`
	ID             *string           `json:"id"`
	IDNe           *string           `json:"id_ne"`
	IDGt           *string           `json:"id_gt"`
	IDLt           *string           `json:"id_lt"`
	IDGte          *string           `json:"id_gte"`
	IDLte          *string           `json:"id_lte"`
	IDIn           []string          `json:"id_in"`
	IDNull         *bool             `json:"id_null"`
	Username       *string           `json:"username"`
	UsernameNe     *string           `json:"username_ne"`
	UsernameGt     *string           `json:"username_gt"`
	UsernameLt     *string           `json:"username_lt"`
	UsernameGte    *string           `json:"username_gte"`
	UsernameLte    *string           `json:"username_lte"`
	UsernameIn     []string          `json:"username_in"`
	UsernameLike   *string           `json:"username_like"`
	UsernamePrefix *string           `json:"username_prefix"`
	UsernameSuffix *string           `json:"username_suffix"`
	UsernameNull   *bool             `json:"username_null"`
	TodoID         *string           `json:"todoId"`
	TodoIDNe       *string           `json:"todoId_ne"`
	TodoIDGt       *string           `json:"todoId_gt"`
	TodoIDLt       *string           `json:"todoId_lt"`
	TodoIDGte      *string           `json:"todoId_gte"`
	TodoIDLte      *string           `json:"todoId_lte"`
	TodoIDIn       []string          `json:"todoId_in"`
	TodoIDNull     *bool             `json:"todoId_null"`
	DeletedBy      *string           `json:"deletedBy"`
	DeletedByNe    *string           `json:"deletedBy_ne"`
	DeletedByGt    *string           `json:"deletedBy_gt"`
	DeletedByLt    *string           `json:"deletedBy_lt"`
	DeletedByGte   *string           `json:"deletedBy_gte"`
	DeletedByLte   *string           `json:"deletedBy_lte"`
	DeletedByIn    []string          `json:"deletedBy_in"`
	DeletedByNull  *bool             `json:"deletedBy_null"`
	UpdatedBy      *string           `json:"updatedBy"`
	UpdatedByNe    *string           `json:"updatedBy_ne"`
	UpdatedByGt    *string           `json:"updatedBy_gt"`
	UpdatedByLt    *string           `json:"updatedBy_lt"`
	UpdatedByGte   *string           `json:"updatedBy_gte"`
	UpdatedByLte   *string           `json:"updatedBy_lte"`
	UpdatedByIn    []string          `json:"updatedBy_in"`
	UpdatedByNull  *bool             `json:"updatedBy_null"`
	CreatedBy      *string           `json:"createdBy"`
	CreatedByNe    *string           `json:"createdBy_ne"`
	CreatedByGt    *string           `json:"createdBy_gt"`
	CreatedByLt    *string           `json:"createdBy_lt"`
	CreatedByGte   *string           `json:"createdBy_gte"`
	CreatedByLte   *string           `json:"createdBy_lte"`
	CreatedByIn    []string          `json:"createdBy_in"`
	CreatedByNull  *bool             `json:"createdBy_null"`
	DeletedAt      *int              `json:"deletedAt"`
	DeletedAtNe    *int              `json:"deletedAt_ne"`
	DeletedAtGt    *int              `json:"deletedAt_gt"`
	DeletedAtLt    *int              `json:"deletedAt_lt"`
	DeletedAtGte   *int              `json:"deletedAt_gte"`
	DeletedAtLte   *int              `json:"deletedAt_lte"`
	DeletedAtIn    []int             `json:"deletedAt_in"`
	DeletedAtNull  *bool             `json:"deletedAt_null"`
	UpdatedAt      *int              `json:"updatedAt"`
	UpdatedAtNe    *int              `json:"updatedAt_ne"`
	UpdatedAtGt    *int              `json:"updatedAt_gt"`
	UpdatedAtLt    *int              `json:"updatedAt_lt"`
	UpdatedAtGte   *int              `json:"updatedAt_gte"`
	UpdatedAtLte   *int              `json:"updatedAt_lte"`
	UpdatedAtIn    []int             `json:"updatedAt_in"`
	UpdatedAtNull  *bool             `json:"updatedAt_null"`
	CreatedAt      *int              `json:"createdAt"`
	CreatedAtNe    *int              `json:"createdAt_ne"`
	CreatedAtGt    *int              `json:"createdAt_gt"`
	CreatedAtLt    *int              `json:"createdAt_lt"`
	CreatedAtGte   *int              `json:"createdAt_gte"`
	CreatedAtLte   *int              `json:"createdAt_lte"`
	CreatedAtIn    []int             `json:"createdAt_in"`
	CreatedAtNull  *bool             `json:"createdAt_null"`
	Todo           *TodoFilterType   `json:"todo"`
}

type UserSortType struct {
	ID        *ObjectSortType `json:"id"`
	Username  *ObjectSortType `json:"username"`
	TodoID    *ObjectSortType `json:"todoId"`
	DeletedBy *ObjectSortType `json:"deletedBy"`
	UpdatedBy *ObjectSortType `json:"updatedBy"`
	CreatedBy *ObjectSortType `json:"createdBy"`
	DeletedAt *ObjectSortType `json:"deletedAt"`
	UpdatedAt *ObjectSortType `json:"updatedAt"`
	CreatedAt *ObjectSortType `json:"createdAt"`
	Todo      *TodoSortType   `json:"todo"`
}

type UserUpdateRelationship struct {
	ID       *string                        `json:"id"`
	Username *string                        `json:"username"`
	Todo     *TodoUpdateReverseRelationship `json:"todo"`
	TodoID   *string                        `json:"todoId"`
}

type UserUpdateReverseRelationship struct {
	ID       *string `json:"id"`
	Username string  `json:"username"`
	TodoID   *string `json:"todoId"`
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
