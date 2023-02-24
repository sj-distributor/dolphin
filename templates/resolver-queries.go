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

	func (r *Generated{{$obj.Name}}ResultTypeResolver) Pages(ctx context.Context, obj *{{$obj.Name}}ResultType) (interface{}, error) {
    total, _       := r.Total(ctx, obj)
    totalPage, _   := r.TotalPage(ctx, obj)
    perPage, _     := r.PerPage(ctx, obj)
    currentPage, _ := r.CurrentPage(ctx, obj)

    return map[string]int{
      "total"        : total,
      "total_page"   : totalPage,
      "per_page"     : perPage,
      "current_page" : currentPage,
    }, nil
	}

	func (r *Generated{{$obj.Name}}ResultTypeResolver) Total(ctx context.Context, obj *{{$obj.Name}}ResultType) (count int, err error) {
		return obj.GetTotal(ctx, r.DB.db, TableName("{{$obj.TableName}}"), &{{$obj.Name}}{})
	}

	func (r *Generated{{$obj.Name}}ResultTypeResolver) CurrentPage(ctx context.Context, obj *{{$obj.Name}}ResultType) (count int, err error) {
	  return int(*obj.EntityResultType.CurrentPage), nil
	}

	func (r *Generated{{$obj.Name}}ResultTypeResolver) PerPage(ctx context.Context, obj *{{$obj.Name}}ResultType) (count int, err error) {
	  return int(*obj.EntityResultType.PerPage), nil
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
				{{if $rel.IsToMany}}
					var input map[string]interface{}
					return r.Handlers.{{$obj.Name}}{{$rel.MethodName}}(ctx, r.GeneratedResolver, obj, input)
				{{else}}
					return r.Handlers.{{$obj.Name}}{{$rel.MethodName}}(ctx, r.GeneratedResolver, obj)
				{{end}}
			}
			{{if $rel.IsToMany}}
				func {{$obj.Name}}{{$rel.MethodName}}Handler(ctx context.Context,r *GeneratedResolver, obj *{{$obj.Name}}, input map[string]interface{}) (res {{$rel.ReturnType}}, err error) {
			{{else}}
				func {{$obj.Name}}{{$rel.MethodName}}Handler(ctx context.Context,r *GeneratedResolver, obj *{{$obj.Name}}) (res {{$rel.ReturnType}}, err error) {
			{{end}}
				{{if $rel.Preload}}
				if obj.{{$rel.MethodName}}Preloaded {
					res = obj.{{$rel.MethodName}}
				}else {
				{{end}}
					{{if $rel.IsToMany}}
							selects := GetFieldsRequested(ctx, strings.ToLower(TableName("{{$rel.Target.TableName}}")))
							items   := []*{{$rel.TargetType}}{}
							wheres  := []string{}
							values  := []interface{}{}

							// if input["label"] != nil && input["value"] != nil {
							// 	err = r.DB.Query().Where(input["label"], input["value"].([]interface{})...).Select(selects).Model(obj).Related(&items, "{{$rel.MethodName}}").Error
							// } else {
							// 	err = r.DB.Query().Select(selects).Model(obj).Related(&items, "{{$rel.MethodName}}").Error
							// }

							err = r.DB.Query().Select(selects).Where(strings.Join(wheres, " AND "), values...).Model(obj).Related(&items, "{{$rel.MethodName}}").Error
							res = items
					{{else}}
						loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
						if obj.{{$rel.MethodName}}ID != nil {
							item, _ := loaders["{{$rel.Target.Name}}"].Load(ctx, dataloader.StringKey(*obj.{{$rel.MethodName}}ID))()
							res, _ = item.({{$rel.ReturnType}})
							{{if $rel.IsNonNull}}
							if res == nil {
								res = &{{$rel.Target.Name}}{}
							// 	_err = fmt.Errorf("{{$rel.Target.Name}} with id '%s' not found",*obj.{{$rel.MethodName}}ID)
							}{{end}}
							// err = _err
						}
					{{end}}
				{{if $rel.Preload}}
				}
				{{end}}
				return
			}
			{{if $rel.IsToMany}}
				func (r *Generated{{$obj.Name}}Resolver) {{$rel.MethodName}}Ids(ctx context.Context, obj *{{$obj.Name}}) (ids []string, err error) {
					wheres  := []string{}
					values  := []interface{}{}

					ids = []string{}
					items := []*{{$rel.TargetType}}{}
					err = r.DB.Query().Model(obj).Select(TableName("{{$rel.Target.TableName}}") + ".id").Where(strings.Join(wheres, " AND "), values...).Related(&items, "{{$rel.MethodName}}").Error

					for _, item := range items {
						ids = append(ids, item.ID)
					}

					return
				}
			{{end}}

		{{end}}

	{{end}}

{{end}}
`
