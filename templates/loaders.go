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
					selects := GetFieldsRequested(ctx, TableName("{{$object.ToSnakePluraName}}", ctx))

					if IndexOf(selects, TableName("{{$object.ToSnakePluraName}}", ctx) + ".*") == -1 {
						if IndexOf(selects, TableName("{{$object.ToSnakePluraName}}", ctx) + ".id") == -1 {
							selects = append(selects, "{{$object.ToSnakePluraName}}" + ".id")
						}
	
						if IndexOf(selects, TableName("{{$object.ToSnakePluraName}}", ctx)+".{{$rel.ToSnakeName}}_id") == -1 {
							selects = append(selects, TableName("{{$object.ToSnakePluraName}}", ctx)+".{{$rel.ToSnakeName}}_id")
						}
					}

					res := db.Query().Table(TableName("{{$object.ToSnakePluraName}}", ctx)).Select(selects).Order("weight ASC, created_at ASC").Find(items, "{{$rel.ToSnakeName}}_id IN (?)", ids)
					if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
						return []*dataloader.Result{
							{Error: res.Error},
						}
					}

					itemMap := make(map[string][]*{{$object.Name}}, len(keys))
					for _, v := range *items {
						item := v
						{{if $rel.IsNonNull}}
							mapKey := item.{{$rel.MethodName}}ID
						{{else}}
							mapKey := *item.{{$rel.MethodName}}ID
						{{end}}
						if itemMap[mapKey] == nil {
							itemMap[mapKey] = []*{{$object.Name}}{}
						}
						itemMap[mapKey] = append(itemMap[mapKey], &item)
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
				loaders["{{$rel.MethodName}}And{{$object.Name}}Ids"] = dataloader.NewBatchedLoader({{$object.TableName}}{{$rel.MethodName}}BatchFn, dataloader.WithClearCacheOnBatch())

			{{end}}
		{{end}}

		{{$object.TableName}}BatchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
			var results []*dataloader.Result
			
			ids := make([]string, len(keys))
			for i, key := range keys {
				ids[i] = key.String()
			}

			items := &[]{{$object.Name}}{}
			selects := GetFieldsRequested(ctx, TableName("{{$object.ToSnakePluraName}}", ctx))
			if len(selects) > 0 && IndexOf(selects, TableName("{{$object.ToSnakePluraName}}", ctx) + ".*") == -1 && IndexOf(selects, TableName("{{$object.ToSnakePluraName}}", ctx) + ".id") == -1 {
				selects = append(selects, TableName("{{$object.ToSnakePluraName}}", ctx) + ".id")
			}

			res := db.Query().Table(TableName("{{$object.ToSnakePluraName}}", ctx)).Select(selects).Order("weight ASC, created_at ASC").Find(items, "id IN (?)", ids)
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
