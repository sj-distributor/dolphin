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

type SubscriptionResolver struct {
	*gen.GeneratedSubscriptionResolver
}

func (r *Resolver) Mutation() gen.MutationResolver {
	return &MutationResolver{&gen.GeneratedMutationResolver{GeneratedResolver: r.GeneratedResolver}}
}
func (r *Resolver) Query() gen.QueryResolver {
	return &QueryResolver{&gen.GeneratedQueryResolver{GeneratedResolver: r.GeneratedResolver}}
}

func (r *Resolver) Subscription() gen.SubscriptionResolver {
	return &SubscriptionResolver{&gen.GeneratedSubscriptionResolver{GeneratedResolver: r.GeneratedResolver}}
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

type ProfileResultTypeResolver struct {
	*gen.GeneratedProfileResultTypeResolver
}

func (r *Resolver) ProfileResultType() gen.ProfileResultTypeResolver {
	return &ProfileResultTypeResolver{&gen.GeneratedProfileResultTypeResolver{GeneratedResolver: r.GeneratedResolver}}
}

type ProfileResolver struct {
	*gen.GeneratedProfileResolver
}

func (r *Resolver) Profile() gen.ProfileResolver {
	return &ProfileResolver{&gen.GeneratedProfileResolver{GeneratedResolver: r.GeneratedResolver}}
}

type TaskResultTypeResolver struct {
	*gen.GeneratedTaskResultTypeResolver
}

func (r *Resolver) TaskResultType() gen.TaskResultTypeResolver {
	return &TaskResultTypeResolver{&gen.GeneratedTaskResultTypeResolver{GeneratedResolver: r.GeneratedResolver}}
}

type TaskResolver struct {
	*gen.GeneratedTaskResolver
}

func (r *Resolver) Task() gen.TaskResolver {
	return &TaskResolver{&gen.GeneratedTaskResolver{GeneratedResolver: r.GeneratedResolver}}
}

type UserRoleResultTypeResolver struct {
	*gen.GeneratedUserRoleResultTypeResolver
}

func (r *Resolver) UserRoleResultType() gen.UserRoleResultTypeResolver {
	return &UserRoleResultTypeResolver{&gen.GeneratedUserRoleResultTypeResolver{GeneratedResolver: r.GeneratedResolver}}
}

type UserRoleResolver struct {
	*gen.GeneratedUserRoleResolver
}

func (r *Resolver) UserRole() gen.UserRoleResolver {
	return &UserRoleResolver{&gen.GeneratedUserRoleResolver{GeneratedResolver: r.GeneratedResolver}}
}

type TagResultTypeResolver struct {
	*gen.GeneratedTagResultTypeResolver
}

func (r *Resolver) TagResultType() gen.TagResultTypeResolver {
	return &TagResultTypeResolver{&gen.GeneratedTagResultTypeResolver{GeneratedResolver: r.GeneratedResolver}}
}

type TagResolver struct {
	*gen.GeneratedTagResolver
}

func (r *Resolver) Tag() gen.TagResolver {
	return &TagResolver{&gen.GeneratedTagResolver{GeneratedResolver: r.GeneratedResolver}}
}
