package gen

import (
	"context"
)

type GeneratedQueryResolver struct{ *GeneratedResolver }

type GeneratedTodoResolver struct{ *GeneratedResolver }

// Todo is the resolver for the todo field.
func (r *GeneratedQueryResolver) Todo(ctx context.Context, id string) (*Todo, error) {
	return r.Handlers.Todo(ctx, r.GeneratedResolver, id)
}

func QueryTodoHandler(ctx context.Context, r *GeneratedResolver, id string) (*Todo, error) {
	db := r.DB.Query()

	var items []*Todo

	rt := &EntityResultType{}

	giOpts := GetItemsOptions{
		Alias:      TableName("todos"),
		Preloaders: []string{},
		Item:       &Todo{},
	}

	err := rt.GetData(ctx, db, giOpts, &items)
	if err != nil {
		return nil, err
	}

	return &Todo{ID: id, Title: "user " + id}, nil
}
