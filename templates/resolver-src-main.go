package templates

var ResolverSrc = `package src
import (
	"context"

	"{{.Config.Package}}/gen"
)

func New(db *gen.DB, ec *gen.EventController) gen.Config {
	resolver := NewResolver(db, ec)

	// resolver.Handlers.CreateUser = func(ctx context.Context, r *gen.GeneratedResolver, input map[string]interface{}) (item *gen.User, err error) {
	// 	return gen.CreateUserHandler(ctx, r, input)
	// }

	// resolver.Handlers.QueryUsers = func(ctx context.Context, r *gen.GeneratedResolver, opts gen.QueryUsersHandlerOptions) (*gen.UserResultType, error) {
	// 	user, err := gen.QueryUserHandler(ctx, r, "userId")
	// 	fmt.Println(user, err)

	// 	return gen.QueryUsersHandler(ctx, r, opts)
	// }

	// events
	resolver.Handlers.OnEvent = func(ctx context.Context, r *gen.GeneratedResolver, e *gen.Event) error {
		return nil
	}

	c := gen.Config{Resolvers: resolver}

	c.Directives.HasRole = func(ctx context.Context, obj any, next graphql.Resolver, role gen.Role) (res any, err error) {
		return next(ctx)
	}

	return resolver
}
`
