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
		item, err = r.Handlers.Create{{$obj.Name}}(ctx, r.GeneratedResolver, input)
		if err!=nil{
			RollbackMutationContext(ctx, r.GeneratedResolver)
			return
		}
		err = FinishMutationContext(ctx, r.GeneratedResolver)
		return
	}
	func Create{{$obj.Name}}Handler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *{{$obj.Name}}, err error) {
		item = &{{$obj.Name}}{}
		now := time.Now()
		timestampMillis := now.UnixNano() / 1e6
		principalID := GetPrincipalIDFromContext(ctx)

		tx := GetTransaction(ctx)

		event := NewEvent(EventMetadata{
			Type:        EventTypeCreated,
			Entity:      "{{$obj.Name}}",
			EntityID:    item.ID,
			Date:        timestampMillis,
			PrincipalID: principalID,
		})

		var changes {{$obj.Name}}Changes
		err = ApplyChanges(input, &changes)
		if err != nil {
			return
		}

		err = CheckStructFieldIsEmpty(item, input)
		if err != nil {
			return nil, err
		}

		item.ID        = uuid.Must(uuid.NewV4()).String()
		item.CreatedAt = timestampMillis
		item.CreatedBy = principalID

		{{range $rel := .Relationships}}
			{{if $rel.IsToMany}}
				if !utils.IsNil(input["{{$rel.Name}}"]) && !utils.IsNil(input["{{$rel.Name}}Ids"]) {
					return nil, fmt.Errorf("{{$rel.Name}}Ids and {{$rel.Name}} cannot coexist")
				}

				if ids, ok := input["{{$rel.Name}}Ids"]; ok && !utils.IsNil(input["{{$rel.Name}}Ids"]) {
					items := []*{{$rel.TargetType}}{}
					itemIds := []string{}
					findIds := []string{}
			
					for _, v := range ids.([]interface{}) {
						itemIds = append(itemIds, v.(string))
					}
			
					if len(itemIds) > 0 {
						// 判断是否有详情权限
						if err := auth.CheckAuthorization(ctx, "{{$rel.TargetType}}"); err != nil {
							return item, errors.New("{{$rel.TargetType}} Detail " + err.Error())
						}
			
						if err := tx.Find(&items, "id IN (?)", itemIds).Error; err != nil {
							return item, err
						}
			
						for _, v := range items {
							findIds = append(findIds, v.ID)
						}
					}
			
					if len(findIds) > 0 {
						differenceIds := utils.Difference(itemIds, findIds)
						if len(differenceIds) > 0 {
							return item, fmt.Errorf("{{$rel.Name}}Ids " + strings.Join(differenceIds, ",") + " not found")
						}
					}
			
					if err := tx.Model(&item).Table("{{$rel.TargetTypeToSnakeName}}s").Association("{{$rel.MethodName}}").Replace(items); err != nil {
						return item, err
					}

					// item.{{$rel.TargetType}}s = items
					event.AddNewValue("{{$rel.Name}}", items)
				}

				if _, ok := input["{{$rel.Name}}"]; ok && !utils.IsNil(input["{{$rel.Name}}"]) {
					new{{$rel.MethodName}} := []*{{$rel.TargetType}}{}
					update{{$rel.MethodName}} := []*{{$rel.TargetType}}{}
					for index, v := range changes.{{$rel.MethodName}} {
						weight := int64(index + 1)
						v.Weight = &weight
						// 判断ID是否为空
						if !utils.IsEmpty(v.ID) {
							// 判断是否有Update权限
							v.UpdatedAt = &timestampMillis
							v.UpdatedBy = principalID
							if err := auth.CheckAuthorization(ctx, "Update{{$rel.TargetType}}"); err != nil {
								return item, errors.New("Update{{$rel.TargetType}} " + err.Error())
							}
			
							// 判断是否有详情权限
							if err := auth.CheckAuthorization(ctx, "{{$rel.TargetType}}"); err != nil {
								return item, errors.New("{{$rel.TargetType}} Detail " + err.Error())
							}
			
							{{$rel.Name}}Input := utils.StructToMap(*v)
							_, err := r.Handlers.Update{{$rel.TargetType}}(ctx, r, {{$rel.Name}}Input["id"].(string), {{$rel.Name}}Input)
							if err != nil {
								return item, errors.New("{{$rel.TargetType}} ID " + v.ID + " " + err.Error())
							}
			
							update{{$rel.MethodName}} = append(update{{$rel.MethodName}}, v)
						} else {
							// 判断是否有Create权限
							if err := auth.CheckAuthorization(ctx, "Create{{$rel.TargetType}}"); err != nil {
								return item, errors.New("Create{{$rel.TargetType}} " + err.Error())
							}
							v.ID = uuid.Must(uuid.NewV4()).String()
							v.CreatedAt = timestampMillis
							v.CreatedBy = principalID
							new{{$rel.MethodName}} = append(new{{$rel.MethodName}}, v)
						}
					}
					
					if err := tx.Model(&item).Table("{{$rel.TargetTypeToSnakeName}}s").Association("{{$rel.MethodName}}").Replace(append(update{{$rel.MethodName}}, new{{$rel.MethodName}}...)); err != nil {
						return item, err
					}

					event.AddNewValue("{{$rel.Name}}", append(update{{$rel.MethodName}}, new{{$rel.MethodName}}...))
				}
			{{else}}
				if _, ok := input["{{$rel.Name}}"]; ok && !utils.IsNil(input["{{$rel.Name}}"]) {
					v := changes.{{$rel.MethodName}}

					// 判断ID是否为空
					if !utils.IsEmpty(v.ID) {
						// 判断是否有Update权限
						v.UpdatedAt = &timestampMillis
						v.UpdatedBy = principalID
						if err := auth.CheckAuthorization(ctx, "Update{{$rel.TargetType}}"); err != nil {
							return item, errors.New("Update{{$rel.TargetType}} " + err.Error())
						}
		
						// 判断是否有详情权限
						if err := auth.CheckAuthorization(ctx, "{{$rel.TargetType}}"); err != nil {
							return item, errors.New("{{$rel.TargetType}} Detail " + err.Error())
						}
		
						{{$rel.Name}}Input := utils.StructToMap(*v)
						{{$rel.Name}}, err := r.Handlers.Update{{$rel.TargetType}}(ctx, r, v.ID, {{$rel.Name}}Input)
						if err != nil {
							return item, errors.New("{{$rel.TargetType}} ID " + v.ID + " " + err.Error())
						}

						if err := tx.Unscoped().Model(&{{$obj.Name}}{}).Where("{{$rel.Name}}_id = ?", {{$rel.Name}}.ID).Updates(map[string]interface{}{"{{$rel.Name}}_id": nil}).Error; err != nil {
							return nil, err
						}
		
					} else {
						// 判断是否有Create权限
						if err := auth.CheckAuthorization(ctx, "Create{{$rel.TargetType}}"); err != nil {
							return item, errors.New("Create{{$rel.TargetType}} " + err.Error())
						}
						v.ID = uuid.Must(uuid.NewV4()).String()
						v.CreatedAt = timestampMillis
						v.CreatedBy = principalID
						{{if $rel.IsOneToOne}}
						{{if $rel.IsNonNull}}
						v.{{$rel.UpperRelationshipName}}ID = &item.ID
						{{else}}
						v.{{$rel.UpperRelationshipName}}ID = item.ID
						{{end}}
						{{end}}
					}
					if err := tx.Model(&item).Table("{{$rel.TargetTypeToSnakeName}}s").Association("{{$rel.MethodName}}").Append(v); err != nil {
						return item, err
					}

					{{if $rel.IsNonNull}}
						item.{{$rel.MethodName}}ID = v.ID
					{{else}}
						item.{{$rel.MethodName}}ID = &v.ID
					{{end}}
					item.{{$rel.MethodName}} = v
					event.AddNewValue("{{$rel.Name}}", item.{{$rel.MethodName}})
					event.AddNewValue("{{$rel.Name}}Id", item.{{$rel.MethodName}}ID)
				}
			{{end}}
		{{end}}

		{{range $col := .Columns}}
			{{if and (not $col.IsHasUpperId) $col.IsCreatable}}
			{{if $col.IsOptional}}
			if _, ok := input["{{$col.Name}}"]; ok && changes.{{$col.MethodName}} != nil {
			{{else}}
			if _, ok := input["{{$col.Name}}"]; ok && !utils.IsEmpty(input["{{$col.Name}}"]) {
			{{end}}
			if (item.{{$col.MethodName}} != changes.{{$col.MethodName}}){{if $col.IsOptional}} || (*item.{{$col.MethodName}} != *changes.{{$col.MethodName}}){{end}} {

			{{if $col.IsRelationshipIdentifier}}
				if !utils.IsNil(input["{{$col.Name}}"]) {
					if err := tx.Select("id").Where("id = ?", input["{{$col.Name}}"]).First(&{{$col.RelationshipTypeName}}{}).Error; err != nil {
						return nil, fmt.Errorf("{{$col.Name}} " + err.Error())
					}
				}
				{{end}}item.{{$col.MethodName}} = changes.{{$col.MethodName}}
				{{if $col.IsIdentifier}}event.EntityID = item.{{$col.MethodName}}
				{{end}}event.AddNewValue("{{$col.Name}}", changes.{{$col.MethodName}})
				}
			}
			{{end}}
		{{end}}
		
	  if err := tx.Omit(clause.Associations).Table(TableName("{{$obj.TableName}}", ctx)).Create(item).Error; err != nil {
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
			RollbackMutationContext(ctx, r.GeneratedResolver)
			return
		}
		err = FinishMutationContext(ctx, r.GeneratedResolver)
		return
	}
	func Update{{$obj.Name}}Handler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *{{$obj.Name}}, err error) {
		item = &{{$obj.Name}}{}

		now := time.Now()
		timestampMillis := now.UnixNano() / 1e6
		principalID := GetPrincipalIDFromContext(ctx)

		tx := GetTransaction(ctx)

		event := NewEvent(EventMetadata{
			Type:        EventTypeUpdated,
			Entity:      "{{$obj.Name}}",
			EntityID:    id,
			Date:        timestampMillis,
			PrincipalID: principalID,
		})

		var changes {{$obj.Name}}Changes
		err = ApplyChanges(input, &changes)
		if err != nil {
			return
		}

		err = CheckStructFieldIsEmpty(item, input)
		if err != nil {
			return nil, err
		}

		{{range $rel := .Relationships}}
			{{if $rel.IsToMany}}
				if !utils.IsNil(input["{{$rel.Name}}"]) && !utils.IsNil(input["{{$rel.Name}}Ids"]) {
					return nil, fmt.Errorf("{{$rel.Name}}Ids and {{$rel.Name}} cannot coexist")
				}
			{{else}}
				if !utils.IsNil(input["{{$rel.Name}}"]) && !utils.IsNil(input["{{$rel.Name}}Id"]) {
					return nil, fmt.Errorf("{{$rel.Name}}Id and {{$rel.Name}} cannot coexist")
				}
			{{end}}
		{{end}}

		if err = GetItem(ctx, tx, TableName("{{$obj.TableName}}", ctx), item, &id); err != nil {
			return nil, err
		}
	
		if item.UpdatedBy != nil && principalID != nil && *item.UpdatedBy != *principalID {
			item.UpdatedBy = principalID
		}

		{{range $rel := .Relationships}}
			{{if $rel.IsToMany}}
				if ids, ok := input["{{$rel.Name}}Ids"]; ok && !utils.IsNil(input["{{$rel.Name}}Ids"]) {
					items := []*{{$rel.TargetType}}{}
					itemIds := []string{}
					findIds := []string{}
			
					for _, v := range ids.([]interface{}) {
						itemIds = append(itemIds, v.(string))
					}
			
					if len(itemIds) > 0 {
						// 判断是否有详情权限
						if err := auth.CheckAuthorization(ctx, "{{$rel.TargetType}}"); err != nil {
							return item, errors.New("{{$rel.TargetType}} Detail " + err.Error())
						}
			
						if err := tx.Find(&items, "id IN (?)", itemIds).Error; err != nil {
							return item, err
						}
			
						for _, v := range items {
							findIds = append(findIds, v.ID)
						}
					}
			
					if len(findIds) > 0 {						
						differenceIds := utils.Difference(itemIds, findIds)
						if len(differenceIds) > 0 {
							return item, fmt.Errorf("{{$rel.Name}}Ids " + strings.Join(differenceIds, ",") + " not found")
						}
					}
			
					if err := tx.Model(&item).Table("{{$rel.TargetTypeToSnakeName}}s").Association("{{$rel.MethodName}}").Replace(items); err != nil {
						return item, err
					}

					// item.{{$rel.TargetType}}s = items
					event.AddNewValue("{{$rel.Name}}", items)
				}

				if _, ok := input["{{$rel.Name}}"]; ok && !utils.IsNil(input["{{$rel.Name}}"]) {
					new{{$rel.MethodName}} := []*{{$rel.TargetType}}{}
					update{{$rel.MethodName}} := []*{{$rel.TargetType}}{}
					for index, v := range changes.{{$rel.MethodName}} {
						weight := int64(index + 1)
						v.Weight = &weight
						// 判断ID是否为空
						if !utils.IsEmpty(v.ID) {
							// 判断是否有Update权限
							v.UpdatedAt = &timestampMillis
							v.UpdatedBy = principalID
							if err := auth.CheckAuthorization(ctx, "Update{{$rel.TargetType}}"); err != nil {
								return item, errors.New("Update{{$rel.TargetType}} " + err.Error())
							}
			
							// 判断是否有详情权限
							if err := auth.CheckAuthorization(ctx, "{{$rel.TargetType}}"); err != nil {
								return item, errors.New("{{$rel.TargetType}} Detail " + err.Error())
							}
			
							{{$rel.Name}}Input := utils.StructToMap(*v)
							_, err := r.Handlers.Update{{$rel.TargetType}}(ctx, r, {{$rel.Name}}Input["id"].(string), {{$rel.Name}}Input)
							if err != nil {
								return item, errors.New("{{$rel.TargetType}} ID " + v.ID + " " + err.Error())
							}
			
							update{{$rel.MethodName}} = append(update{{$rel.MethodName}}, v)
						} else {
							// 判断是否有Create权限
							if err := auth.CheckAuthorization(ctx, "Create{{$rel.TargetType}}"); err != nil {
								return item, errors.New("Create{{$rel.TargetType}} " + err.Error())
							}
							v.ID = uuid.Must(uuid.NewV4()).String()
							v.CreatedAt = timestampMillis
							v.CreatedBy = principalID
							new{{$rel.MethodName}} = append(new{{$rel.MethodName}}, v)
						}
					}

					if err := tx.Model(&item).Table("{{$rel.TargetTypeToSnakeName}}s").Association("{{$rel.MethodName}}").Replace(append(update{{$rel.MethodName}}, new{{$rel.MethodName}}...)); err != nil {
						return item, err
					}

					event.AddNewValue("{{$rel.Name}}", append(update{{$rel.MethodName}}, new{{$rel.MethodName}}...))
				}
			{{else}}
				if _, ok := input["{{$rel.Name}}"]; ok && !utils.IsNil(input["{{$rel.Name}}"]) {
					v := changes.{{$rel.MethodName}}

					// 判断ID是否为空
					if !utils.IsEmpty(v.ID) {
						// 判断是否有Update权限
						v.UpdatedAt = &timestampMillis
						v.UpdatedBy = principalID
						if err := auth.CheckAuthorization(ctx, "Update{{$rel.TargetType}}"); err != nil {
							return item, errors.New("Update{{$rel.TargetType}} " + err.Error())
						}
		
						// 判断是否有详情权限
						if err := auth.CheckAuthorization(ctx, "{{$rel.TargetType}}"); err != nil {
							return item, errors.New("{{$rel.TargetType}} Detail " + err.Error())
						}
		
						{{$rel.Name}}Input := utils.StructToMap(*v)
						{{$rel.Name}}, err := r.Handlers.Update{{$rel.TargetType}}(ctx, r, v.ID, {{$rel.Name}}Input)
						if err != nil {
							return item, errors.New("{{$rel.TargetType}} ID " + v.ID + " " + err.Error())
						}

						if err := tx.Unscoped().Model(&{{$obj.Name}}{}).Where("{{$rel.Name}}_id = ?", {{$rel.Name}}.ID).Updates(map[string]interface{}{"{{$rel.Name}}_id": nil}).Error; err != nil {
							return nil, err
						}

					} else {
						// 判断是否有Create权限
						if err := auth.CheckAuthorization(ctx, "Create{{$rel.TargetType}}"); err != nil {
							return item, errors.New("Create{{$rel.TargetType}} " + err.Error())
						}
						v.ID = uuid.Must(uuid.NewV4()).String()
						v.CreatedAt = timestampMillis
						v.CreatedBy = principalID
					}

					// if err := tx.Model(&item).Association("{{$rel.MethodName}}").Append(v); err != nil {
					// 	return item, err
					// }

					{{if $rel.IsNonNull}}
						item.{{$rel.MethodName}}ID = v.ID
					{{else}}
						item.{{$rel.MethodName}}ID = &v.ID
					{{end}}

					// item.{{$rel.MethodName}} = v
					// event.AddNewValue("{{$rel.Name}}", item.{{$rel.MethodName}})
					// event.AddNewValue("{{$rel.Name}}Id", item.{{$rel.MethodName}}ID)
				}
			{{end}}
		{{end}}

		{{range $col := .Columns}}
			{{if and (not $col.IsHasUpperId) $col.IsUpdatable}}
			{{if $col.IsOptional}}
			if _, ok := input["{{$col.Name}}"]; ok {
			{{else}}
			if _, ok := input["{{$col.Name}}"]; ok {
			{{end}}
			// if (item.{{$col.MethodName}} != changes.{{$col.MethodName}}){{if $col.IsOptional}} || (*item.{{$col.MethodName}} != *changes.{{$col.MethodName}}){{end}} {

					{{if $col.IsRelationshipIdentifier}}
						if !utils.IsNil(input["{{$col.Name}}"]) {
							if err := tx.Select("id").Where("id = ?", input["{{$col.Name}}"]).First(&{{$col.RelationshipTypeName}}{}).Error; err != nil {
								return nil, fmt.Errorf("{{$col.Name}} " + err.Error())
							}
						}
					{{end}}event.AddOldValue("{{$col.Name}}", item.{{$col.MethodName}})
					event.AddNewValue("{{$col.Name}}", changes.{{$col.MethodName}})
					item.{{$col.MethodName}} = changes.{{$col.MethodName}}
					}
				// }
			{{end}}
		{{end}}

		if err := tx.Table(TableName("{{$obj.TableName}}", ctx)).Where("id = ?", id).Save(item).Error; err != nil {
	    return item, err
	  }
		
    {{range $rel := $obj.Relationships}}
    {{if $rel.IsToMany}}{{if not $rel.Target.IsExtended}}
      if ids,exists:=input["{{$rel.Name}}Ids"]; exists {
      	if len(ids.([]interface{})) > 0 {
	        items := []{{$rel.TargetType}}{}
	        tx.Find(&items, "id IN (?)", ids)
	        if err := tx.Model(&item).Association("{{$rel.MethodName}}").Replace(items); err != nil {
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
		done,err := r.Handlers.Delete{{$obj.PluralName}}(ctx, r.GeneratedResolver, id, unscoped)
		if err != nil {
			RollbackMutationContext(ctx, r.GeneratedResolver)
			return done, err
		}
		err = FinishMutationContext(ctx, r.GeneratedResolver)
		return done,err
	}

	func Delete{{$obj.PluralName}}Handler(ctx context.Context, r *GeneratedResolver, id []string, unscoped *bool) (bool, error) {
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
		done,err := r.Handlers.Recovery{{$obj.PluralName}}(ctx, r.GeneratedResolver, id)
		if err != nil {
			return done, err
		}
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
