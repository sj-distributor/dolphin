package templates

var ResolverSrcGen = `package src

import (
	"{{.Config.Package}}/gen"
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


{{range $obj := .Model.ObjectEntities}}
	type {{$obj.Name}}ResultTypeResolver struct {
		*gen.Generated{{$obj.Name}}ResultTypeResolver
	}
	func (r *Resolver) {{$obj.Name}}ResultType() gen.{{$obj.Name}}ResultTypeResolver {
		return &{{$obj.Name}}ResultTypeResolver{&gen.Generated{{$obj.Name}}ResultTypeResolver{GeneratedResolver: r.GeneratedResolver}}
	}
	{{if $obj.NeedsQueryResolver}}
		type {{$obj.Name}}Resolver struct {
			*gen.Generated{{$obj.Name}}Resolver
		}
		func (r *Resolver) {{$obj.Name}}() gen.{{$obj.Name}}Resolver {
			return &{{$obj.Name}}Resolver{&gen.Generated{{$obj.Name}}Resolver{GeneratedResolver: r.GeneratedResolver}}
		}
	{{end}}
{{end}}
{{range $ext := .Model.ObjectExtensions}}
	{{$obj := $ext.Object}}
	{{if not $ext.ExtendsLocalObject}}
		type {{$obj.Name}}Resolver struct {
			*gen.Generated{{$obj.Name}}Resolver
		}
		{{if or $obj.HasAnyRelationships $obj.HasReadonlyColumns}}
			func (r *Resolver) {{$obj.Name}}() gen.{{$obj.Name}}Resolver {
				return &{{$obj.Name}}Resolver{&gen.Generated{{$obj.Name}}Resolver{GeneratedResolver: r.GeneratedResolver}}
			}
		{{end}}
	{{end}}
{{end}}
`
