package gen

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
)

type GeneratedQueryResolver struct{ *GeneratedResolver }

type GeneratedTodoResolver struct{ *GeneratedResolver }

// Todo is the resolver for the todo field.
func (r *GeneratedQueryResolver) Todo(ctx context.Context, id string) (*Todo, error) {
	return r.Handlers.Todo(ctx, r.GeneratedResolver, &id)
}

func QueryTodoHandler(ctx context.Context, r *GeneratedResolver, id *string) (*Todo, error) {
	selection := []ast.Selection{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		selection = append(selection, f.Field)
	}
	selectionSet := ast.SelectionSet(selection)

	query := TodoQueryFilter{}
	rt := &TodoResultType{
		EntityResultType: EntityResultType{
			Query:        &query,
			SelectionSet: &selectionSet,
		},
	}
	qb := r.DB.Query()
	if id != nil {
		qb = qb.Where(TableName("todos")+".id = ?", id)
	}

	var items []*Todo
	giOpts := GetItemsOptions{
		Alias:      TableName("todos"),
		Preloaders: []string{},
		Item:       &Todo{},
	}
	err := rt.GetData(ctx, qb, giOpts, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, &NotFoundError{Entity: "todo"}
	}
	return items[0], err
}
