package gen

import (
	"context"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/iancoleman/strcase"
	"github.com/sj-distributor/dolphin-example/enums"
	"github.com/vektah/gqlparser/v2/ast"
	"gorm.io/gorm"
)

func IndexOf(str []string, data interface{}) int {
	for k, v := range str {
		if v == data {
			return k
		}
	}

	return -1
}

func SnakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}

func GetPrincipalIDFromContext(ctx context.Context) *string {
	v, _ := ctx.Value(KeyPrincipalID).(*string)
	return v
}

func EnrichContextWithMutations(ctx context.Context, r *GeneratedResolver) context.Context {
	_ctx := context.WithValue(ctx, KeyMutationTransaction, r.DB.db.Begin())
	_ctx = context.WithValue(_ctx, KeyMutationEvents, &MutationEvents{})
	return _ctx
}

func FinishMutationContext(ctx context.Context, r *GeneratedResolver) (err error) {
	s := GetMutationEventStore(ctx)

	for _, event := range s.Events {
		err = r.Handlers.OnEvent(ctx, r, &event)
		if err != nil {
			return
		}
	}

	tx := GetTransaction(ctx)
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return
	}

	// for _, event := range s.Events {
	// 	err = r.EventController.SendEvent(ctx, &event)
	// }

	return
}

func GetTransaction(ctx context.Context) *gorm.DB {
	return ctx.Value(KeyMutationTransaction).(*gorm.DB)
}

func GetMutationEventStore(ctx context.Context) *MutationEvents {
	return ctx.Value(KeyMutationEvents).(*MutationEvents)
}

func AddMutationEvent(ctx context.Context, e Event) {
	s := GetMutationEventStore(ctx)
	s.Events = append(s.Events, e)
}

// GetFieldsRequested ...
func GetFieldsRequested(ctx context.Context, alias string) []string {
	// result := graphql.CollectAllFields(ctx)
	reqCtx := graphql.GetOperationContext(ctx)

	// if IndexOf(result, alias) != -1 || alias != reqCtx.OperationName {
	// 	return []string{alias + ".*"}
	// }
	fieldSelections := graphql.GetFieldContext(ctx).Field.Selections
	return recurseSelectionSets(reqCtx, []string{}, fieldSelections, alias)
}

// recurseSelectionSets ...
func recurseSelectionSets(reqCtx *graphql.OperationContext, fields []string, selection ast.SelectionSet, alias string) []string {
	goTypeMap := []string{"String", "Time", "ID", "Float", "Int", "Boolean"}

	for _, sel := range selection {
		switch sel := sel.(type) {
		case *ast.Field:
			fieldName := ""
			// && !strings.Contains(sel.Name, "Ids")
			if strings.Contains(sel.Name, "Ids") && sel.Definition.Type.Name() != "ID" {
				fieldName = SnakeString(sel.Name)
				if alias != "" {
					fieldName = alias + "." + SnakeString(sel.Name)
				}
			} else if !strings.HasPrefix(sel.Name, "__") && !strings.Contains(sel.Name, "Ids") {
				if len(sel.SelectionSet) == 0 && IndexOf(goTypeMap, sel.Definition.Type.Name()) != -1 {
					fieldName = SnakeString(sel.Name)
					if alias != "" {
						fieldName = alias + "." + SnakeString(sel.Name)
					}
				} else {
					reg, _ := regexp.Compile("^\\[(.+?)\\]")
					IsToMany := reg.MatchString(sel.Definition.Type.String())
					if !IsToMany && len(sel.SelectionSet) != 0 && IndexOf(goTypeMap, sel.Definition.Type.Name()) == -1 {
						fieldName = SnakeString(sel.Name) + "_id"
						if alias != "" {
							fieldName = alias + "." + SnakeString(sel.Name) + "_id"
						}
					}
				}
			}
			if fieldName != "" && IndexOf(fields, fieldName) == -1 {
				fields = append(fields, fieldName)
			}
		}
	}
	return fields
}

// CheckStructFieldIsEmpty ...
func CheckStructFieldIsEmpty(item interface{}, input map[string]interface{}) (err error) {
	res := []string{}
	value := reflect.ValueOf(item)
	elem := value.Elem()
	elemKey := elem.Type()

	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i)
		if !strings.Contains(reflect.TypeOf(field.Interface()).String(), "*") {
			name := strcase.ToSnake(strcase.ToLowerCamel(elemKey.Field(i).Name))
			if v, ok := input[name]; ok {
				if v == "" || (strings.Contains(reflect.TypeOf(v).String(), "[]") && len(v.([]interface{})) <= 0) {
					res = append(res, name)
				}
			}
		}
	}

	if len(res) > 0 {
		err = fmt.Errorf(fmt.Sprintf(enums.CannotBeEmpty, strings.Join(res, "，")))
	}
	return err
}
