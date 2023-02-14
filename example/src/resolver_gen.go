package src

import (
	"github.com/sj-distributor/dolphin-example/gen"
)

func NewResolver(db *gen.DB, ec *gen.EventController) *Resolver {
	handlers := gen.DefaultResolutionHandlers()
	return &Resolver{&gen.GeneratedResolver{Handlers: handlers, DB: db, EventController: ec}}
}

type Resolver struct {
	*gen.GeneratedResolver
}

type MutationResolver struct {
	*gen.GeneratedMutationResolver
}

type QueryResolver struct {
	*gen.GeneratedQueryResolver
}

func (r *Resolver) Mutation() gen.MutationResolver {
	return &MutationResolver{&gen.GeneratedMutationResolver{r.GeneratedResolver}}
}

func (r *Resolver) Query() gen.QueryResolver {
	return &QueryResolver{&gen.GeneratedQueryResolver{r.GeneratedResolver}}
}

type TodoResultTypeResolver struct {
	*gen.GeneratedTodoResultTypeResolver
}

func (r *Resolver) TodoResultType() gen.TodoResultTypeResolver {
	return &TodoResultTypeResolver{&gen.GeneratedTodoResultTypeResolver{r.GeneratedResolver}}
}
