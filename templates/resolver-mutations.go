package templates

var ResolverMutations = `package gen

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/graph-gophers/dataloader"
	"github.com/99designs/gqlgen/graphql"
	"{{.Config.Package}}/validator"
)

type GeneratedMutationResolver struct{ *GeneratedResolver }

type MutationEvents struct {
	Events []Event
}
{{range $obj := .Model.ObjectEntities}}
	func (r *GeneratedMutationResolver) Create{{$obj.Name}}(ctx context.Context, input map[string]interface{}) (item *{{$obj.Name}}, err error) {
		ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
		item, err = r.Handlers.Create{{$obj.Name}}(ctx, r.GeneratedResolver, input)
		if err!=nil{
			return
		}
		err = FinishMutationContext(ctx, r.GeneratedResolver)
		return
	}
	func Create{{$obj.Name}}Handler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *{{$obj.Name}}, err error) {
		item = &{{$obj.Name}}{}

		now := time.Now()
		milliTime := now.UnixNano() / 1e6
		principalID := GetPrincipalIDFromContext(ctx)

		tx := GetTransaction(ctx)
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
			}
		}()

		event := NewEvent(EventMetadata{
			Type:        EventTypeCreated,
			Entity:      "{{$obj.Name}}",
			EntityID:    item.ID,
			Date:        milliTime,
			PrincipalID: principalID,
		})

		var changes {{$obj.Name}}Changes
		err = ApplyChanges(input, &changes)
		if err != nil {
			tx.Rollback()
			return
		}

		err = CheckStructFieldIsEmpty(item, input)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		item.ID        = uuid.Must(uuid.NewV4()).String()
		item.CreatedAt = milliTime
		item.CreatedBy = principalID

		{{range $col := .Columns}}
			{{if and (not $col.IsHasUpperId) $col.IsCreatable}}
				if _, ok := input["{{$col.Name}}"]; ok && (item.{{$col.MethodName}} != changes.{{$col.MethodName}}){{if $col.IsOptional}} && (item.{{$col.MethodName}} == nil || changes.{{$col.MethodName}} == nil || *item.{{$col.MethodName}} != *changes.{{$col.MethodName}}){{end}} {
					item.{{$col.MethodName}} = changes.{{$col.MethodName}}
					{{if $col.IsIdentifier}}event.EntityID = item.{{$col.MethodName}}{{end}}
					event.AddNewValue("{{$col.Name}}", changes.{{$col.MethodName}})
				}
			{{end}}
		{{end}}

	  if err := tx.Create(item).Error; err != nil {
	  	tx.Rollback()
	    return item, err
	  }

		if len(event.Changes) > 0 {
			AddMutationEvent(ctx, event)
		}

		return
	}
	func (r *GeneratedMutationResolver) Update{{$obj.Name}}(ctx context.Context, id string, input map[string]interface{}) (item *{{$obj.Name}}, err error) {
		ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
		item,err = r.Handlers.Update{{$obj.Name}}(ctx, r.GeneratedResolver, id, input)
		if err!=nil{
			return
		}
		err = FinishMutationContext(ctx, r.GeneratedResolver)
		return
	}
	func Update{{$obj.Name}}Handler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *{{$obj.Name}}, err error) {
		item = &{{$obj.Name}}{}
		newItem := &{{$obj.Name}}{}


		isChange := false
		now := time.Now()
		milliTime := now.UnixNano() / 1e6
		principalID := GetPrincipalIDFromContext(ctx)

		tx := GetTransaction(ctx)
    defer func() {
      if r := recover(); r != nil {
        tx.Rollback()
      }
    }()

		event := NewEvent(EventMetadata{
			Type:        EventTypeUpdated,
			Entity:      "{{$obj.Name}}",
			EntityID:    id,
			Date:        milliTime,
			PrincipalID: principalID,
		})

		var changes {{$obj.Name}}Changes
		err = ApplyChanges(input, &changes)
		if err != nil {
			tx.Rollback()
			return
		}

		err = CheckStructFieldIsEmpty(item, input)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		if err = GetItem(ctx, tx, TableName("{{$obj.TableName}}"), item, &id); err != nil {
			tx.Rollback()
			return nil, err
		}
	
		if item.UpdatedBy != nil && principalID != nil && *item.UpdatedBy != *principalID {
			newItem.UpdatedBy = principalID
		}

		{{range $col := .Columns}}
			{{if and (not $col.IsHasUpperId) $col.IsUpdatable}}
				if _, ok := input["{{$col.Name}}"]; ok && (item.{{$col.MethodName}} != changes.{{$col.MethodName}}){{if $col.IsOptional}} && (item.{{$col.MethodName}} == nil || changes.{{$col.MethodName}} == nil || *item.{{$col.MethodName}} != *changes.{{$col.MethodName}}){{end}} {
					{{if $col.IsRelationshipIdentifier}}if err := tx.Select("id").Where("id", input["{{$col.Name}}"]).First(&{{$col.RelationshipName}}{}).Error; err != nil {
						tx.Rollback()
						return nil, fmt.Errorf("{{$col.Name}} " + err.Error())
					}
					{{end}}event.AddOldValue("{{$col.Name}}", item.{{$col.MethodName}})
					event.AddNewValue("{{$col.Name}}", changes.{{$col.MethodName}})
					item.{{$col.MethodName}} = changes.{{$col.MethodName}}
					newItem.{{$col.MethodName}} = changes.{{$col.MethodName}}
					isChange = true
				}
			{{end}}
		{{end}}

		if err := validator.Struct(item); err != nil {
			tx.Rollback()
			return nil, err
		}
	
		if !isChange {
			return item, nil
		}

	  if err := tx.Model(&newItem).Where("id = ?", id).Updates(newItem).Error; err != nil {
	  	tx.Rollback()
	    return item, err
	  }

		if len(event.Changes) > 0 {
			AddMutationEvent(ctx, event)
		}

		return
	}

	func Delete{{$obj.Name}}Func(ctx context.Context, r *GeneratedResolver, id string, tye string, unscoped *bool) (err error) {
		principalID := GetPrincipalIDFromContext(ctx)
		item := &{{$obj.Name}}{}
		now := time.Now()
		tx := GetTransaction(ctx)
    defer func() {
      if r := recover(); r != nil {
        tx.Rollback()
      }
    }()

		if err = GetItem(ctx, tx, TableName("{{$obj.TableName}}"), item, &id); err != nil {
			tx.Rollback()
			return err
		}

		deletedAt := now.UnixNano() / 1e6

		event := NewEvent(EventMetadata{
			Type:        EventTypeDeleted,
			Entity:      "{{$obj.Name}}",
			EntityID:    id,
			Date:        deletedAt,
			PrincipalID: principalID,
		})

		// 如果是恢复删除数据
		if tye == "recovery" {
			if err := tx.Unscoped().Model(&item).Updates(map[string]interface{}{"DeletedAt": nil, "DeletedBy": nil}).Error; err != nil {
				tx.Rollback()
				return err
			}
		} else {
			if unscoped != nil && *unscoped == true {
				if err := tx.Unscoped().Model(&item).Delete(item).Error; err != nil {
					tx.Rollback()
					return err
				}
			} else if err := tx.Model(&item).Updates({{$obj.Name}}{DeletedAt: &deletedAt, DeletedBy: principalID, UpdatedBy: principalID}).Error; err != nil {
				tx.Rollback()
				return err
			}
		}

		if len(event.Changes) > 0 {
			AddMutationEvent(ctx, event)
		}

		return
	}

	func (r *GeneratedMutationResolver) Delete{{$obj.PluralName}}(ctx context.Context, id []string, unscoped *bool) (bool, error) {
		ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
		done,err:=r.Handlers.Delete{{$obj.PluralName}}(ctx, r.GeneratedResolver, id, unscoped)
		err = FinishMutationContext(ctx, r.GeneratedResolver)
		return done,err
	}

	func Delete{{$obj.PluralName}}Handler(ctx context.Context, r *GeneratedResolver, id []string, unscoped *bool) (bool,error) {
		var err error = nil

		if len(id) > 0 {
			for _, v := range id {
				err = Delete{{$obj.Name}}Func(ctx, r, v, "delete", unscoped)
				if err != nil {
					break
				}
			}
		}

		if err != nil {
			return false, err
		}
		return true, err
	}

	func (r *GeneratedMutationResolver) Recovery{{$obj.PluralName}}(ctx context.Context, id []string) (bool, error) {
		ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
		done,err:=r.Handlers.Recovery{{$obj.PluralName}}(ctx, r.GeneratedResolver, id)
		err = FinishMutationContext(ctx, r.GeneratedResolver)
		return done,err
	}

	func Recovery{{$obj.PluralName}}Handler(ctx context.Context, r *GeneratedResolver, id []string) (bool,error) {
		var err error = nil

		var unscoped bool = false

		if len(id) > 0 {
			for _, v := range id {
				err = Delete{{$obj.Name}}Func(ctx, r, v, "recovery", &unscoped)
				if err != nil {
					break
				}
			}
		}

		if err != nil {
			return false, err
		}
		return true, err
	}
{{end}}
`
