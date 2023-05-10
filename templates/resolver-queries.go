package templates

var ResolverQueries = `package gen

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
	"github.com/graph-gophers/dataloader"
	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
	"{{.Config.Package}}/auth"
	"{{.Config.Package}}/utils"
)

type GeneratedQueryResolver struct{ *GeneratedResolver }

{{range $obj := .Model.ObjectEntities}}
	func (r *GeneratedQueryResolver) {{$obj.Name}}(ctx context.Context, id string) (*{{$obj.Name}}, error) {
		return r.Handlers.Query{{$obj.Name}}(ctx, r.GeneratedResolver, id)
	}
	func Query{{$obj.Name}}Handler(ctx context.Context, r *GeneratedResolver, id string) (*{{$obj.Name}}, error) {
		selection := []ast.Selection{}
		for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
			selection = append(selection, f.Field)
		}
		selectionSet := ast.SelectionSet(selection)

		query := {{$obj.Name}}QueryFilter{}
		rt := &{{$obj.Name}}ResultType{
			EntityResultType: EntityResultType{
				Query:        &query,
				SelectionSet: &selectionSet,
			},
		}
		qb := r.DB.Query()
		qb = qb.Where(TableName("{{$obj.TableName}}") + ".id = ?", id)

		var items []*{{$obj.Name}}
		giOpts := GetItemsOptions{
			Alias:TableName("{{$obj.TableName}}"),
			Preloaders:[]string{ {{range $r := $obj.PreloadableRelationships}}
				"{{$r.MethodName}}",{{end}}
			},
			Item: &{{$obj.Name}}{},
		}
		err := rt.GetData(ctx, qb, giOpts, &items)
		if err != nil {
			return nil, err
		}
		if len(items) == 0 {
			return nil, &NotFoundError{Entity: "{{$obj.Name}}"}
		}
		return items[0], err
	}

	type Query{{$obj.PluralName}}HandlerOptions struct {
		CurrentPage *int
		PerPage  *int
		Q      *string
		Sort   []*{{$obj.Name}}SortType
		Filter *{{$obj.Name}}FilterType
		Rand   *bool
	}
	func (r *GeneratedQueryResolver) {{$obj.PluralName}}(ctx context.Context, current_page *int, per_page *int, q *string, sort []*{{$obj.Name}}SortType, filter *{{$obj.Name}}FilterType, rand *bool) (*{{$obj.Name}}ResultType, error) {
		opts := Query{{$obj.PluralName}}HandlerOptions{
      CurrentPage: current_page,
      PerPage:  per_page,
			Q: q,
			Sort: sort,
			Filter: filter,
			Rand: rand,
		}
		return r.Handlers.Query{{$obj.PluralName}}(ctx, r.GeneratedResolver, opts)
	}
	func Query{{$obj.PluralName}}Handler(ctx context.Context, r *GeneratedResolver, opts Query{{$obj.PluralName}}HandlerOptions) (*{{$obj.Name}}ResultType, error) {
		query := {{$obj.Name}}QueryFilter{opts.Q}

		var selectionSet *ast.SelectionSet
		for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
			if f.Field.Name == "data" {
				selectionSet = &f.Field.SelectionSet
			}
		}

		_sort := []EntitySort{}
		for _, sort := range opts.Sort {
			_sort = append(_sort, sort)
		}

		return &{{$obj.Name}}ResultType{
			EntityResultType: EntityResultType{
				CurrentPage: opts.CurrentPage,
				PerPage:  opts.PerPage,
				Rand:  opts.Rand,
				Query:  &query,
				Sort: _sort,
				Filter: opts.Filter,
				SelectionSet: selectionSet,
			},
		}, nil
	}

	type Generated{{$obj.Name}}ResultTypeResolver struct{ *GeneratedResolver }

	func (r *Generated{{$obj.Name}}ResultTypeResolver) Data(ctx context.Context, obj *{{$obj.Name}}ResultType) (items []*{{$obj.Name}}, err error) {
		giOpts := GetItemsOptions{
			Alias:TableName("{{$obj.TableName}}"),
			Preloaders:[]string{ {{range $r := $obj.PreloadableRelationships}}
				"{{$r.MethodName}}",{{end}}
			},
			Item: &{{$obj.Name}}{},
		}
		err = obj.GetData(ctx, r.DB.db, giOpts, &items)
		{{if $obj.HasPreloadableRelationships}}
			for _, item := range items {
				{{range $rel := $obj.PreloadableRelationships}}
				item.{{$rel.MethodName}}Preloaded = true{{end}}
			}
		{{end}}

		uniqueItems := []*{{$obj.Name}}{}
		idMap := map[string]bool{}
		for _, item := range items {
			if _,ok := idMap[item.ID]; !ok {
				idMap[item.ID] = true
				uniqueItems = append(uniqueItems, item)
			}
		}
		items = uniqueItems

		return
	}

	func (r *Generated{{$obj.Name}}ResultTypeResolver) Total(ctx context.Context, obj *{{$obj.Name}}ResultType) (count int, err error) {
		return obj.GetTotal(ctx, r.DB.db, TableName("{{$obj.TableName}}"), &{{$obj.Name}}{})
	}

	func (r *Generated{{$obj.Name}}ResultTypeResolver) TotalPage(ctx context.Context, obj *{{$obj.Name}}ResultType) (count int, err error) {
	  total, _   := r.Total(ctx, obj)
	  perPage, _ := r.PerPage(ctx, obj)
	  totalPage  := int(math.Ceil(float64(total) / float64(perPage)))
		if totalPage < 0 {
			totalPage = 0
		}

	  return totalPage, nil
	}

	func (r *Generated{{$obj.Name}}ResultTypeResolver) CurrentPage(ctx context.Context, obj *{{$obj.Name}}ResultType) (count int, err error) {
	  return int(*obj.EntityResultType.CurrentPage), nil
	}

	func (r *Generated{{$obj.Name}}ResultTypeResolver) PerPage(ctx context.Context, obj *{{$obj.Name}}ResultType) (count int, err error) {
	  return int(*obj.EntityResultType.PerPage), nil
	}

	{{if $obj.NeedsQueryResolver}}
		type Generated{{$obj.Name}}Resolver struct { *GeneratedResolver }

		{{range $col := $obj.Fields}}
			{{if $col.NeedsQueryResolver}}
				func (r *Generated{{$obj.Name}}Resolver) {{$col.MethodName}}(ctx context.Context, obj *{{$obj.Name}}) (res {{$col.GoType}}, err error) {
					return r.Handlers.{{$obj.Name}}{{$col.MethodName}}(ctx, r.GeneratedResolver, obj)
				}
				func {{$obj.Name}}{{$col.MethodName}}Handler(ctx context.Context,r *GeneratedResolver, obj *{{$obj.Name}}) (res {{$col.GoType}}, err error) {
					{{if and (not $col.IsList) $col.HasTargetTypeWithIDField ($obj.HasColumn (print $col.Name "Id"))}}
						if obj.{{$col.MethodName}}ID != nil {
							res = &{{$col.TargetType}}{ID: *obj.{{$col.MethodName}}ID}
						}
					{{else}}
						err = fmt.Errorf("Resolver handler for {{$obj.Name}}{{$col.MethodName}} not implemented")
					{{end}}
					return
				}
			{{end}}
		{{end}}

		{{range $index, $rel := $obj.Relationships}}
			func (r *Generated{{$obj.Name}}Resolver) {{$rel.MethodName}}(ctx context.Context, obj *{{$obj.Name}}) (res {{$rel.ReturnType}}, err error) {
				return r.Handlers.{{$obj.Name}}{{$rel.MethodName}}(ctx, r.GeneratedResolver, obj)
			}
			func {{$obj.Name}}{{$rel.MethodName}}Handler(ctx context.Context,r *GeneratedResolver, obj *{{$obj.Name}}) (items {{$rel.ReturnType}}, err error) {
				loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
				{{if $rel.IsToMany}}
						item, _ := loaders["{{$rel.TargetType}}{{$rel.UpperRelationshipName}}"].Load(ctx, dataloader.StringKey(obj.ID))()
						items = {{$rel.ReturnType}}{}
						if item != nil {
							items = item.({{$rel.ReturnType}})
						}
				{{else}}
					if obj.{{$rel.MethodName}}ID != nil {
						item, _ := loaders["{{$rel.Target.Name}}"].Load(ctx, dataloader.StringKey(*obj.{{$rel.MethodName}}ID))()
						items, _ = item.({{$rel.ReturnType}})
						{{if $rel.IsNonNull}}
							if items == nil {
								items = &{{$rel.Target.Name}}{}
							}
						{{end}}
					}
				{{end}}
				return
			}
			{{if $rel.IsToMany}}
				func (r *Generated{{$obj.Name}}Resolver) {{$rel.MethodName}}Ids(ctx context.Context, obj *{{$obj.Name}}) (ids []string, err error) {
					ids = []string{}
					items := {{$rel.ReturnType}}{}

					loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
					item, _ := loaders["{{$rel.TargetType}}{{$rel.UpperRelationshipName}}"].Load(ctx, dataloader.StringKey(obj.ID))()
				
					if item != nil {
						items = item.({{$rel.ReturnType}})
					}

					for _, v := range items {
						ids = append(ids, v.ID)
					}

					return
				}
			{{end}}

		{{end}}

	{{end}}

{{end}}
`
