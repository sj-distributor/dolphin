package gen

import (
	"context"
	"regexp"
	"strings"

	"github.com/99designs/gqlgen/graphql"
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

func GetFieldsRequested(ctx context.Context, alias string) []string {
	reqCtx := graphql.GetRequestContext(ctx)
	fieldSelections := graphql.GetResolverContext(ctx).Field.Selections
	return recurseSelectionSets(reqCtx, []string{}, fieldSelections, alias)
}

func recurseSelectionSets(reqCtx *graphql.RequestContext, fields []string, selection ast.SelectionSet, alias string) []string {
	goTypeMap := []string{"String", "Time", "ID", "Float", "Int", "Boolean"}

	for _, sel := range selection {
		switch sel := sel.(type) {
		case *ast.Field:
			if !strings.HasPrefix(sel.Name, "__") && strings.Index(sel.Name, "Ids") == -1 {
				if len(sel.SelectionSet) == 0 && IndexOf(goTypeMap, sel.Definition.Type.Name()) != -1 {
					if alias != "" {
						fields = append(fields, alias+"."+SnakeString(sel.Name))
					} else {
						fields = append(fields, SnakeString(sel.Name))
					}
				} else {
					IsToMany, _ := regexp.MatchString("^\\[(.+?)\\]", sel.Definition.Type.String())
					if IsToMany == false {
						if alias != "" {
							fields = append(fields, alias+"."+SnakeString(sel.Name)+"_id")
						} else {
							fields = append(fields, SnakeString(sel.Name)+"_id")
						}
					}
				}
			}
		}
	}

	return fields
}
