package templates

var ResolverCore = `//go:generate go run github.com/99designs/gqlgen generate
package gen

import (
	"context"
)

type ResolutionHandlers struct {
	OnEvent func(ctx context.Context, r *GeneratedResolver, e *Event) error
	WebSocket func(ctx context.Context, r *GeneratedResolver) (<-chan any, error)
	{{range $obj := .Model.ObjectEntities}}
		Create{{$obj.Name}} func (ctx context.Context, r *GeneratedResolver, input map[string]interface{}, authType bool) (item *{{$obj.Name}}, err error)
		Update{{$obj.Name}} func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}, authType bool) (item *{{$obj.Name}}, err error)
		Delete{{$obj.PluralName}} func (ctx context.Context, r *GeneratedResolver, id []string, unscoped *bool, authType bool) (bool, error)
		Recovery{{$obj.PluralName}} func (ctx context.Context, r *GeneratedResolver, id []string, authType bool) (bool, error)
		Query{{$obj.Name}} func (ctx context.Context, r *GeneratedResolver, opts Query{{$obj.Name}}HandlerOptions, authType bool) (*{{$obj.Name}}, error)
		Query{{$obj.PluralName}} func (ctx context.Context, r *GeneratedResolver, opts Query{{$obj.PluralName}}HandlerOptions, authType bool) (*{{$obj.Name}}ResultType, error)
		{{range $rel := $obj.Relationships}}
			{{$obj.Name}}{{$rel.MethodName}} func (ctx context.Context,r *GeneratedResolver, obj *{{$obj.Name}}, authType bool) (res {{$rel.ReturnType}}, err error)
		{{end}}
	{{end}}
}

func DefaultResolutionHandlers() ResolutionHandlers {
	handlers := ResolutionHandlers{
		OnEvent: func(ctx context.Context, r *GeneratedResolver, e *Event) error { return nil },
		{{range $obj := .Model.ObjectEntities}}
			Create{{$obj.Name}}: Create{{$obj.Name}}Handler,
			Update{{$obj.Name}}: Update{{$obj.Name}}Handler,
			Delete{{$obj.PluralName}}: Delete{{$obj.PluralName}}Handler,
			Recovery{{$obj.PluralName}}: Recovery{{$obj.PluralName}}Handler,
			Query{{$obj.Name}}: Query{{$obj.Name}}Handler,
			Query{{$obj.PluralName}}: Query{{$obj.PluralName}}Handler,
			{{range $rel := $obj.Relationships}}
				{{$obj.Name}}{{$rel.MethodName}}: {{$obj.Name}}{{$rel.MethodName}}Handler,
			{{end}}
		{{end}}
		WebSocket: WebSocketHandler,
	}
	return handlers
}

type GeneratedResolver struct {
	Handlers ResolutionHandlers
	DB *DB
	EventController *EventController
}
`
