// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gen

import (
	"fmt"
	"io"
	"strconv"
)

type Todo struct {
	ID        string  `json:"id"`
	Title     string  `json:"title"`
	CreatedBy *string `json:"createdBy"`
}

type TodoFilterType struct {
	And         []*TodoFilterType `json:"AND"`
	Or          []*TodoFilterType `json:"OR"`
	ID          *string           `json:"id"`
	IDNe        *string           `json:"id_ne"`
	IDGt        *string           `json:"id_gt"`
	IDLt        *string           `json:"id_lt"`
	IDGte       *string           `json:"id_gte"`
	IDLte       *string           `json:"id_lte"`
	IDIn        []string          `json:"id_in"`
	IDNull      *bool             `json:"id_null"`
	Title       *string           `json:"title"`
	TitleNe     *string           `json:"title_ne"`
	TitleGt     *string           `json:"title_gt"`
	TitleLt     *string           `json:"title_lt"`
	TitleGte    *string           `json:"title_gte"`
	TitleLte    *string           `json:"title_lte"`
	TitleIn     []string          `json:"title_in"`
	TitleLike   *string           `json:"title_like"`
	TitlePrefix *string           `json:"title_prefix"`
	TitleSuffix *string           `json:"title_suffix"`
	TitleNull   *bool             `json:"title_null"`
}

type TodoSortType struct {
	ID    *ObjectSortType `json:"id"`
	Title *ObjectSortType `json:"title"`
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
