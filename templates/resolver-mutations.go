package templates

var ResolverMutations = `package gen

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/graph-gophers/dataloader"
	"github.com/99designs/gqlgen/graphql"
	"{{.Config.Package}}/utils"
)

type GeneratedMutationResolver struct{ *GeneratedResolver }

type MutationEvents struct {
	Events []Event
}
{{range $obj := .Model.ObjectEntities}}
	func (r *GeneratedMutationResolver) Create{{$obj.Name}}(ctx context.Context, input map[string]interface{}) (item *{{$obj.Name}}, err error) {
		ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
		item, err = r.Handlers.Create{{$obj.Name}}(ctx, r.GeneratedResolver, input, true)
		if err!=nil{
			return
		}
		err = FinishMutationContext(ctx, r.GeneratedResolver)
		return
	}
	func Create{{$obj.Name}}Handler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}, authType bool) (item *{{$obj.Name}}, err error) {
		item = &{{$obj.Name}}{}
		if err := auth.CheckRouterAuth(ctx, authType); err != nil {
			return item, err
		}
	
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
			return nil, err
		}

		item.ID        = uuid.Must(uuid.NewV4()).String()
		item.CreatedAt = milliTime
		item.CreatedBy = principalID

		{{range $col := .Columns}}
			{{if and (not $col.IsHasUpperId) $col.IsCreatable}}
				if _, ok := input["{{$col.Name}}"]; ok && (item.{{$col.MethodName}} != changes.{{$col.MethodName}}){{if $col.IsOptional}} && (item.{{$col.MethodName}} == nil || changes.{{$col.MethodName}} == nil || *item.{{$col.MethodName}} != *changes.{{$col.MethodName}}){{end}} {
					{{if $col.IsRelationshipIdentifier}}
						if err := tx.Select("id").Where("id", input["{{$col.Name}}"]).First(&{{$col.RelationshipTypeName}}{}).Error; err != nil {
							return nil, fmt.Errorf("{{$col.Name}} " + err.Error())
						}
					{{end}}item.{{$col.MethodName}} = changes.{{$col.MethodName}}
					{{if $col.IsIdentifier}}event.EntityID = item.{{$col.MethodName}}
					{{end}}event.AddNewValue("{{$col.Name}}", changes.{{$col.MethodName}})
				}
			{{end}}
		{{end}}

		if err := utils.Validate(item); err != nil {
			return nil, err
		}
		
	  if err := tx.Omit(clause.Associations).Create(item).Error; err != nil {
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
		item,err = r.Handlers.Update{{$obj.Name}}(ctx, r.GeneratedResolver, id, input, true)
		if err!=nil{
			return
		}
		err = FinishMutationContext(ctx, r.GeneratedResolver)
		return
	}
	func Update{{$obj.Name}}Handler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}, authType bool) (item *{{$obj.Name}}, err error) {
		item = &{{$obj.Name}}{}
		newItem := &{{$obj.Name}}{}
		if err := auth.CheckRouterAuth(ctx, authType); err != nil {
			return item, err
		}

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
			return nil, err
		}

		if err = GetItem(ctx, tx, TableName("{{$obj.TableName}}"), item, &id); err != nil {
			return nil, err
		}
	
		if item.UpdatedBy != nil && principalID != nil && *item.UpdatedBy != *principalID {
			newItem.UpdatedBy = principalID
		}

		{{range $col := .Columns}}
			{{if and (not $col.IsHasUpperId) $col.IsUpdatable}}
				if _, ok := input["{{$col.Name}}"]; ok && (item.{{$col.MethodName}} != changes.{{$col.MethodName}}){{if $col.IsOptional}} && (item.{{$col.MethodName}} == nil || changes.{{$col.MethodName}} == nil || *item.{{$col.MethodName}} != *changes.{{$col.MethodName}}){{end}} {
					{{if $col.IsRelationshipIdentifier}}
						if err := tx.Select("id").Where("id", input["{{$col.Name}}"]).First(&{{$col.RelationshipName}}{}).Error; err != nil {
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

		if err := utils.Validate(item); err != nil {
			return nil, err
		}
	
		if !isChange {
			return item, nil
		}

		if err := tx.Model(&newItem).Where("id = ?", id).Updates(newItem).Error; err != nil {
	  	tx.Rollback()
	    return item, err
	  }
		
    {{range $rel := $obj.Relationships}}
    {{if $rel.IsToMany}}{{if not $rel.Target.IsExtended}}
      if ids,exists:=input["{{$rel.Name}}Ids"]; exists {
      	if len(ids.([]interface{})) > 0 {
	        items := []{{$rel.TargetType}}{}
	        tx.Find(&items, "id IN (?)", ids)
	        if err := tx.Model(&item).Association("{{$rel.MethodName}}").Replace(items); err != nil {
						tx.Rollback()
						return item, err
					}
      	}
      }
    {{end}}{{end}}
    {{end}}

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

		{{range $rel := .Relationships}}
			{{if $rel.IsMaster}}
				if err := tx.Where("is_delete = ? and {{$rel.ToSnakeRelationshipName}}_id = ?", 1, id).First(&{{$rel.TargetType}}{}).Error; err == nil {
					return fmt.Errorf("{{$rel.TargetType}} exists, cannot be deleted")
				}
			{{end}}
		{{end}}
		
		var status int64 = 1
		var isDelete int64 = 2
		if tye == "recovery" {
			isDelete = 1
			status = 2
		}

		if err = tx.Unscoped().Where("is_delete = ? and id = ?", status, id).First(item).Error; err != nil {
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
			if err := tx.Unscoped().Model(&item).Updates(map[string]interface{}{"IsDelete": 1, "DeletedAt": nil, "DeletedBy": nil}).Error; err != nil {
				return err
			}
		} else {
			if unscoped != nil && *unscoped {
				if err := tx.Unscoped().Model(&item).Delete(item).Error; err != nil {
					return err
				}
			} else if err := tx.Model(&item).Updates({{$obj.Name}}{IsDelete: &isDelete, DeletedAt: &deletedAt, DeletedBy: principalID, UpdatedBy: principalID}).Error; err != nil {
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
		done,err := r.Handlers.Delete{{$obj.PluralName}}(ctx, r.GeneratedResolver, id, unscoped, true)
		if err != nil {
			return done, err
		}
		err = FinishMutationContext(ctx, r.GeneratedResolver)
		return done,err
	}

	func Delete{{$obj.PluralName}}Handler(ctx context.Context, r *GeneratedResolver, id []string, unscoped *bool, authType bool) (bool,error) {
		tx := GetTransaction(ctx)
		var err error = nil
		if err := auth.CheckRouterAuth(ctx, authType); err != nil {
			return false, err
		}

		if len(id) > 0 {
			for _, v := range id {
				err = Delete{{$obj.Name}}Func(ctx, r, v, "delete", unscoped)
				if err != nil {
					break
				}
			}
		}

		if err != nil {
			tx.Rollback()
			return false, err
		}
		return true, err
	}

	func (r *GeneratedMutationResolver) Recovery{{$obj.PluralName}}(ctx context.Context, id []string) (bool, error) {
		ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
		done,err := r.Handlers.Recovery{{$obj.PluralName}}(ctx, r.GeneratedResolver, id, true)
		if err != nil {
			return done, err
		}
		err = FinishMutationContext(ctx, r.GeneratedResolver)
		return done,err
	}

	func Recovery{{$obj.PluralName}}Handler(ctx context.Context, r *GeneratedResolver, id []string, authType bool) (bool,error) {
		var err error = nil
		if err := auth.CheckRouterAuth(ctx, authType); err != nil {
			return false, err
		}
	
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
