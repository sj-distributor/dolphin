package templates

var ResolverCore = `package gen

import (
	"context"
)

type ResolutionHandlers struct {
	OnEvent func(ctx context.Context, r *GeneratedResolver, e *Event) error
	{{range $obj := .Model.ObjectEntities}}
		Create{{$obj.Name}} func (ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *{{$obj.Name}}, err error)
		Update{{$obj.Name}} func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *{{$obj.Name}}, err error)
		Delete{{$obj.PluralName}} func (ctx context.Context, r *GeneratedResolver, id []string, unscoped *bool) (bool, error)
		Recovery{{$obj.PluralName}} func (ctx context.Context, r *GeneratedResolver, id []string) (bool, error)
		Query{{$obj.Name}} func (ctx context.Context, r *GeneratedResolver, id string) (*{{$obj.Name}}, error)
		Query{{$obj.PluralName}} func (ctx context.Context, r *GeneratedResolver, opts Query{{$obj.PluralName}}HandlerOptions) (*{{$obj.Name}}ResultType, error)
		{{range $rel := $obj.Relationships}}
			{{if $rel.IsToMany}}
				{{$obj.Name}}{{$rel.MethodName}} func (ctx context.Context,r *GeneratedResolver, obj *{{$obj.Name}}, input map[string]interface{}) (res {{$rel.ReturnType}}, err error)
			{{else}}
				{{$obj.Name}}{{$rel.MethodName}} func (ctx context.Context,r *GeneratedResolver, obj *{{$obj.Name}}) (res {{$rel.ReturnType}}, err error)
			{{end}}
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
	}
	return handlers
}

type GeneratedResolver struct {
	Handlers ResolutionHandlers
	DB *DB
	EventController *EventController
}
`
