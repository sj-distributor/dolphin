package templates

var Loaders = `package gen
import (
	"context"
	"errors"

	"github.com/graph-gophers/dataloader"
	"gorm.io/gorm"
)

func GetLoaders(db *DB) map[string]*dataloader.Loader {
	loaders := map[string]*dataloader.Loader{}
	
	{{range $object := .Model.ObjectEntities}}
		{{range $rel := .Relationships}}
			{{if $rel.IsToOne}}
				{{$object.TableName}}{{$rel.MethodName}}BatchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
					var results []*dataloader.Result
					
					ids := make([]string, len(keys))
					for i, key := range keys {
						ids[i] = key.String()
					}
			
					items := &[]{{$object.Name}}{}
					selects := GetFieldsRequested(ctx, "{{$object.ToSnakePluraName}}")
					if IndexOf(selects, "{{$object.ToSnakePluraName}}" + ".id") == -1 {
						selects = append(selects, "{{$object.ToSnakePluraName}}" + ".id")
					}

					if IndexOf(selects, "{{$object.ToSnakePluraName}}"+".{{$rel.ToSnakeName}}_id") == -1 {
						selects = append(selects, "{{$object.ToSnakePluraName}}"+".{{$rel.ToSnakeName}}_id")
					}
			
					res := db.Query().Select(selects).Order("weight ASC, created_at ASC").Find(items, "{{$rel.ToSnakeName}}_id IN (?)", ids)
					if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
						return []*dataloader.Result{
							{Error: res.Error},
						}
					}

					itemMap := make(map[string][]*{{$object.Name}}, len(keys))
					for _, v := range *items {
						item := v
						if itemMap[*item.{{$rel.MethodName}}ID] == nil {
							itemMap[*item.{{$rel.MethodName}}ID] = []*{{$object.Name}}{}
						}
						itemMap[*item.{{$rel.MethodName}}ID] = append(itemMap[*item.{{$rel.MethodName}}ID], &item)
					}
			
					for _, key := range keys {
						id := key.String()
						item, ok := itemMap[id]
						if !ok {
							results = append(results, &dataloader.Result{
								Data: nil,
								Error: nil,
								// Error: fmt.Errorf("{{$object.Name}} with id '%s' not found", id),
							})
						} else {
							results = append(results, &dataloader.Result{
								Data:  item,
								Error: nil,
							})
						}
					}
					return results
				}
				loaders["{{$object.Name}}{{$rel.MethodName}}"] = dataloader.NewBatchedLoader({{$object.TableName}}{{$rel.MethodName}}BatchFn, dataloader.WithClearCacheOnBatch())
				loaders["{{$object.Name}}Ids"] = dataloader.NewBatchedLoader({{$object.TableName}}{{$rel.MethodName}}BatchFn, dataloader.WithClearCacheOnBatch())

			{{end}}
		{{end}}

		{{$object.TableName}}BatchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
			var results []*dataloader.Result
			
			ids := make([]string, len(keys))
			for i, key := range keys {
				ids[i] = key.String()
			}

			items := &[]{{$object.Name}}{}
			selects := GetFieldsRequested(ctx, "{{$object.ToSnakePluraName}}")
			if len(selects) > 0 && IndexOf(selects, "{{$object.ToSnakePluraName}}" + ".id") == -1 {
				selects = append(selects, "{{$object.ToSnakePluraName}}" + ".id")
			}

			res := db.Query().Select(selects).Order("weight ASC, created_at ASC").Find(items, "id IN (?)", ids)
			if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
				return []*dataloader.Result{
					{Error: res.Error},
				}
			}

			itemMap := make(map[string]{{$object.Name}}, len(keys))
			for _, item := range *items {
				itemMap[item.ID] = item
			}

			for _, key := range keys {
				id := key.String()
				item, ok := itemMap[id]
				if !ok {
					results = append(results, &dataloader.Result{
						Data: nil,
						Error: nil,
						// Error: fmt.Errorf("{{$object.Name}} with id '%s' not found", id),
					})
				} else {
					results = append(results, &dataloader.Result{
						Data:  &item,
						Error: nil,
					})
				}
			}
			return results
		}

		loaders["{{$object.Name}}"] = dataloader.NewBatchedLoader({{$object.TableName}}BatchFn, dataloader.WithClearCacheOnBatch())

	{{end}}

	return loaders
}
`
