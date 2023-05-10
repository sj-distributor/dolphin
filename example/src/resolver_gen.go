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

type UserResultTypeResolver struct {
	*gen.GeneratedUserResultTypeResolver
}

func (r *Resolver) UserResultType() gen.UserResultTypeResolver {
	return &UserResultTypeResolver{&gen.GeneratedUserResultTypeResolver{GeneratedResolver: r.GeneratedResolver}}
}

type UserResolver struct {
	*gen.GeneratedUserResolver
}

func (r *Resolver) User() gen.UserResolver {
	return &UserResolver{&gen.GeneratedUserResolver{GeneratedResolver: r.GeneratedResolver}}
}

type AccountResultTypeResolver struct {
	*gen.GeneratedAccountResultTypeResolver
}

func (r *Resolver) AccountResultType() gen.AccountResultTypeResolver {
	return &AccountResultTypeResolver{&gen.GeneratedAccountResultTypeResolver{GeneratedResolver: r.GeneratedResolver}}
}

type AccountResolver struct {
	*gen.GeneratedAccountResolver
}

func (r *Resolver) Account() gen.AccountResolver {
	return &AccountResolver{&gen.GeneratedAccountResolver{GeneratedResolver: r.GeneratedResolver}}
}

type TransactionResultTypeResolver struct {
	*gen.GeneratedTransactionResultTypeResolver
}

func (r *Resolver) TransactionResultType() gen.TransactionResultTypeResolver {
	return &TransactionResultTypeResolver{&gen.GeneratedTransactionResultTypeResolver{GeneratedResolver: r.GeneratedResolver}}
}

type TransactionResolver struct {
	*gen.GeneratedTransactionResolver
}

func (r *Resolver) Transaction() gen.TransactionResolver {
	return &TransactionResolver{&gen.GeneratedTransactionResolver{GeneratedResolver: r.GeneratedResolver}}
}

type TodoResultTypeResolver struct {
	*gen.GeneratedTodoResultTypeResolver
}

func (r *Resolver) TodoResultType() gen.TodoResultTypeResolver {
	return &TodoResultTypeResolver{&gen.GeneratedTodoResultTypeResolver{GeneratedResolver: r.GeneratedResolver}}
}

type TodoResolver struct {
	*gen.GeneratedTodoResolver
}

func (r *Resolver) Todo() gen.TodoResolver {
	return &TodoResolver{&gen.GeneratedTodoResolver{GeneratedResolver: r.GeneratedResolver}}
}
