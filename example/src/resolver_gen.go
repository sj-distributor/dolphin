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
	return &MutationResolver{&gen.GeneratedMutationResolver{GeneratedResolver: r.GeneratedResolver}}
}
func (r *Resolver) Query() gen.QueryResolver {
	return &QueryResolver{&gen.GeneratedQueryResolver{GeneratedResolver: r.GeneratedResolver}}
}

type BookCategoryResultTypeResolver struct {
	*gen.GeneratedBookCategoryResultTypeResolver
}

func (r *Resolver) BookCategoryResultType() gen.BookCategoryResultTypeResolver {
	return &BookCategoryResultTypeResolver{&gen.GeneratedBookCategoryResultTypeResolver{GeneratedResolver: r.GeneratedResolver}}
}

type BookCategoryResolver struct {
	*gen.GeneratedBookCategoryResolver
}

func (r *Resolver) BookCategory() gen.BookCategoryResolver {
	return &BookCategoryResolver{&gen.GeneratedBookCategoryResolver{GeneratedResolver: r.GeneratedResolver}}
}

type BookResultTypeResolver struct {
	*gen.GeneratedBookResultTypeResolver
}

func (r *Resolver) BookResultType() gen.BookResultTypeResolver {
	return &BookResultTypeResolver{&gen.GeneratedBookResultTypeResolver{GeneratedResolver: r.GeneratedResolver}}
}

type BookResolver struct {
	*gen.GeneratedBookResolver
}

func (r *Resolver) Book() gen.BookResolver {
	return &BookResolver{&gen.GeneratedBookResolver{GeneratedResolver: r.GeneratedResolver}}
}
