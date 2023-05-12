package templates

var ResolverSrc = `package src
import (
	"context"

	"{{.Config.Package}}/gen"
)

func New(db *gen.DB, ec *gen.EventController) *Resolver {
	resolver := NewResolver(db, ec)

	// resolver.Handlers.CreateUser = func(ctx context.Context, r *gen.GeneratedMutationResolver, input map[string]interface{}) (item *gen.Company, err error) {
	// 	return gen.CreateUserHandler(ctx, r, input)
	// }
	// resolver.Handlers.Todo = func(ctx context.Context, id string) (*gen.Todo, error) {
	// 	return &gen.Todo{ID: id, Title: "user_ " + id}, nil
	// }

	// events
	resolver.Handlers.OnEvent = func(ctx context.Context, r *gen.GeneratedResolver, e *gen.Event) error {
		return nil
	}
	return resolver
}

// You can extend QueryResolver for adding custom fields in schema
// func (r *QueryResolver) Hello(ctx context.Context) (string, error) {
// 	return "world", nil
// }
`
