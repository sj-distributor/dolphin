package templates

// ResolverMutations 生成 GraphQL Mutation 解析器代码的模板
//
// 关系类型说明（基于 GORM）：
// - OneToOne: 一对一关系，外键在当前表或关联表
// - OneToMany: 一对多关系，外键在关联表（如 User.tasks → Task.user_id）
// - ManyToOne: 多对一关系，外键在当前表（如 Task.user → Task.user_id）
// - ManyToMany: 多对多关系，需要中间表
//
// 主要功能：
// - Create: 创建新实体及其关系
// - Update: 更新现有实体及其关系
// - Delete: 软删除/硬删除实体
// - Recovery: 恢复已删除的实体
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

// ============================================================
// 类型定义
// ============================================================

// GeneratedMutationResolver 生成的 Mutation 解析器
type GeneratedMutationResolver struct{ *GeneratedResolver }

// MutationEvents 变更事件集合
type MutationEvents struct {
	Events []Event
}

// ============================================================
// 实体 Mutation 解析器
// ============================================================
{{range $obj := .Model.ObjectEntities}}

// ============================================================
// {{$obj.Name}} - Create
// ============================================================

// Create{{$obj.Name}} 创建 {{$obj.Name}} 实体的解析器入口
func (r *GeneratedMutationResolver) Create{{$obj.Name}}(ctx context.Context, input map[string]interface{}) (item *{{$obj.Name}}, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.Create{{$obj.Name}}(ctx, r.GeneratedResolver, input)
	if err != nil {
		RollbackMutationContext(ctx, r.GeneratedResolver)
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// Create{{$obj.Name}}Handler 处理 {{$obj.Name}} 创建逻辑
func Create{{$obj.Name}}Handler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *{{$obj.Name}}, err error) {
	item = &{{$obj.Name}}{}
	now := time.Now()
	timestampMillis := now.UnixNano() / 1e6
	principalID := GetPrincipalIDFromContext(ctx)
	tx := GetTransaction(ctx)

	// 创建事件记录
	event := NewEvent(EventMetadata{
		Type:        EventTypeCreated,
		Entity:      "{{$obj.Name}}",
		EntityID:    item.ID,
		Date:        timestampMillis,
		PrincipalID: principalID,
	})

	// 解析输入变更
	var changes {{$obj.Name}}Changes
	if err = ApplyChanges(input, &changes); err != nil {
		return
	}

	// 验证必填字段
	if err = CheckStructFieldIsEmpty(item, input); err != nil {
		return nil, err
	}

	// 设置基础字段
	item.ID = uuid.Must(uuid.NewV4()).String()
	item.CreatedAt = timestampMillis
	item.CreatedBy = principalID

	// ========== 验证关系字段冲突 ==========
	{{range $rel := .Relationships}}
	{{if $rel.IsToMany}}
	// ToMany: {{$rel.Name}} - 不能同时传入 IDs 和嵌套对象
	if !utils.IsNil(input["{{$rel.Name}}"]) && !utils.IsNil(input["{{$rel.Name}}Ids"]) {
		return nil, fmt.Errorf("{{$rel.Name}}Ids and {{$rel.Name}} cannot coexist")
	}
	{{end}}
	{{end}}

	// ========== 处理 ManyToOne/OneToOne 关系（当前实体持有外键） ==========
	{{range $rel := .Relationships}}
	{{if $rel.IsToOne}}
	{{if or $rel.IsManyToOne $rel.IsOneToOne}}
	// ---------- {{if $rel.IsManyToOne}}ManyToOne{{else}}OneToOne{{end}}: {{$rel.Name}} (当前表持有外键 {{$obj.ToSnakeName}}.{{$rel.ToSnakeName}}_id) ----------
	if _, ok := input["{{$rel.Name}}"]; ok && !utils.IsNil(input["{{$rel.Name}}"]) {
		v := changes.{{$rel.MethodName}}

		if !utils.IsEmpty(v.ID) {
			// 更新现有关联对象
			v.UpdatedAt = &timestampMillis
			v.UpdatedBy = principalID

			if err := auth.CheckAuthorization(ctx, "Update{{$rel.TargetType}}"); err != nil {
				return item, errors.New("Update{{$rel.TargetType}} " + err.Error())
			}
			if err := auth.CheckAuthorization(ctx, "{{$rel.TargetType}}"); err != nil {
				return item, errors.New("{{$rel.TargetType}} Detail " + err.Error())
			}

			{{$rel.Name}}Input := utils.StructToMap(*v)
			if _, err := r.Handlers.Update{{$rel.TargetType}}(ctx, r, v.ID, {{$rel.Name}}Input); err != nil {
				return item, errors.New("{{$rel.TargetType}} ID " + v.ID + " " + err.Error())
			}
			
			// 设置外键
			{{if $rel.IsNonNull}}
			item.{{$rel.MethodName}}ID = v.ID
			{{else}}
			item.{{$rel.MethodName}}ID = &v.ID
			{{end}}
		} else {
			// 创建新关联对象
			if err := auth.CheckAuthorization(ctx, "Create{{$rel.TargetType}}"); err != nil {
				return item, errors.New("Create{{$rel.TargetType}} " + err.Error())
			}
			
			v.ID = uuid.Must(uuid.NewV4()).String()
			v.CreatedAt = timestampMillis
			v.CreatedBy = principalID

			// 先保存关联对象
			if err := tx.Omit(clause.Associations).Table(TableName("{{$rel.TargetTypeToSnakeName}}s", ctx)).Create(v).Error; err != nil {
				return item, err
			}

			// 设置外键
			{{if $rel.IsNonNull}}
			item.{{$rel.MethodName}}ID = v.ID
			{{else}}
			item.{{$rel.MethodName}}ID = &v.ID
			{{end}}
		}

		item.{{$rel.MethodName}} = v
		event.AddNewValue("{{$rel.Name}}", item.{{$rel.MethodName}})
		event.AddNewValue("{{$rel.Name}}Id", item.{{$rel.MethodName}}ID)
	}
	{{end}}
	{{end}}
	{{end}}

	// ========== 处理普通字段 ==========
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
			{{end}}
			item.{{$col.MethodName}} = changes.{{$col.MethodName}}
			{{if $col.IsIdentifier}}
			event.EntityID = item.{{$col.MethodName}}
			{{end}}
			event.AddNewValue("{{$col.Name}}", changes.{{$col.MethodName}})
		}
	}
	{{end}}
	{{end}}

	// ========== 保存主实体 ==========
	if err := tx.Omit(clause.Associations).Table(TableName("{{$obj.TableName}}", ctx)).Create(item).Error; err != nil {
		return item, err
	}

	// ========== 处理 OneToMany/ManyToMany 关系（关联表持有外键或中间表） ==========
	{{range $rel := .Relationships}}
	{{if $rel.IsToMany}}
	// ---------- ToMany: {{$rel.Name}} ({{if $rel.IsManyToMany}}ManyToMany 中间表{{else}}OneToMany 外键在 {{$rel.TargetType}}.{{$rel.ToSnakeRelationshipName}}_id{{end}}) ----------
	
	// 方式1：通过 IDs 关联现有记录
	if ids, ok := input["{{$rel.Name}}Ids"]; ok && !utils.IsNil(input["{{$rel.Name}}Ids"]) {
		items := []*{{$rel.TargetType}}{}
		itemIds := []string{}
		findIds := []string{}

		for _, v := range ids.([]interface{}) {
			itemIds = append(itemIds, v.(string))
		}

		if len(itemIds) > 0 {
			// 权限检查
			if err := auth.CheckAuthorization(ctx, "{{$rel.TargetType}}"); err != nil {
				return item, errors.New("{{$rel.TargetType}} Detail " + err.Error())
			}

			if err := tx.Find(&items, "id IN (?)", itemIds).Error; err != nil {
				return item, err
			}

			for _, v := range items {
				findIds = append(findIds, v.ID)
			}

			// 验证所有 ID 都存在
			differenceIds := utils.Difference(itemIds, findIds)
			if len(differenceIds) > 0 {
				return item, fmt.Errorf("{{$rel.Name}}Ids " + strings.Join(differenceIds, ",") + " not found")
			}

			{{if $rel.IsManyToMany}}
			// ManyToMany: 使用 Replace 更新中间表
			if err := tx.Model(item).Association("{{$rel.MethodName}}").Replace(items); err != nil {
				return item, err
			}
			{{else}}
			// OneToMany: 更新关联记录的外键
			for _, relItem := range items {
				if err := tx.Model(relItem).Update("{{$rel.ToSnakeRelationshipName}}_id", item.ID).Error; err != nil {
					return item, err
				}
			}
			{{end}}
		}
		event.AddNewValue("{{$rel.Name}}", items)
	}

	// 方式2：通过嵌套对象创建/更新
	if _, ok := input["{{$rel.Name}}"]; ok && !utils.IsNil(input["{{$rel.Name}}"]) {
		new{{$rel.MethodName}} := []*{{$rel.TargetType}}{}
		update{{$rel.MethodName}} := []*{{$rel.TargetType}}{}

		for index, v := range changes.{{$rel.MethodName}} {
			weight := int64(index + 1)
			v.Weight = &weight

			if !utils.IsEmpty(v.ID) {
				// 更新现有记录
				v.UpdatedAt = &timestampMillis
				v.UpdatedBy = principalID

				if err := auth.CheckAuthorization(ctx, "Update{{$rel.TargetType}}"); err != nil {
					return item, errors.New("Update{{$rel.TargetType}} " + err.Error())
				}
				if err := auth.CheckAuthorization(ctx, "{{$rel.TargetType}}"); err != nil {
					return item, errors.New("{{$rel.TargetType}} Detail " + err.Error())
				}

				{{$rel.Name}}Input := utils.StructToMap(*v)
				if _, err := r.Handlers.Update{{$rel.TargetType}}(ctx, r, {{$rel.Name}}Input["id"].(string), {{$rel.Name}}Input); err != nil {
					return item, errors.New("{{$rel.TargetType}} ID " + v.ID + " " + err.Error())
				}
				
				{{if not $rel.IsManyToMany}}
				// OneToMany: 设置外键指向当前实体
				if err := tx.Model(v).Update("{{$rel.ToSnakeRelationshipName}}_id", item.ID).Error; err != nil {
					return item, err
				}
				{{end}}
				
				update{{$rel.MethodName}} = append(update{{$rel.MethodName}}, v)
			} else {
				// 创建新记录
				if err := auth.CheckAuthorization(ctx, "Create{{$rel.TargetType}}"); err != nil {
					return item, errors.New("Create{{$rel.TargetType}} " + err.Error())
				}
				
				v.ID = uuid.Must(uuid.NewV4()).String()
				v.CreatedAt = timestampMillis
				v.CreatedBy = principalID
				
				{{if not $rel.IsManyToMany}}
				// OneToMany: 设置外键指向当前实体
				{{if $rel.InverseRelationship.IsNonNull}}
				v.{{$rel.UpperRelationshipName}}ID = item.ID
				{{else}}
				v.{{$rel.UpperRelationshipName}}ID = &item.ID
				{{end}}
				{{end}}

				// 保存新记录
				if err := tx.Omit(clause.Associations).Table(TableName("{{$rel.TargetTypeToSnakeName}}s", ctx)).Create(v).Error; err != nil {
					return item, err
				}

				new{{$rel.MethodName}} = append(new{{$rel.MethodName}}, v)
			}
		}

		allItems := append(update{{$rel.MethodName}}, new{{$rel.MethodName}}...)
		
		{{if $rel.IsManyToMany}}
		// ManyToMany: 使用 Replace 更新中间表
		if err := tx.Model(item).Association("{{$rel.MethodName}}").Replace(allItems); err != nil {
			return item, err
		}
		{{end}}
		
		event.AddNewValue("{{$rel.Name}}", allItems)
	}
	{{end}}
	{{end}}

	// 记录事件
	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}

// ============================================================
// {{$obj.Name}} - Update
// ============================================================

// Update{{$obj.Name}} 更新 {{$obj.Name}} 实体的解析器入口
func (r *GeneratedMutationResolver) Update{{$obj.Name}}(ctx context.Context, id string, input map[string]interface{}) (item *{{$obj.Name}}, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.Update{{$obj.Name}}(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		RollbackMutationContext(ctx, r.GeneratedResolver)
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// Update{{$obj.Name}}Handler 处理 {{$obj.Name}} 更新逻辑
func Update{{$obj.Name}}Handler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *{{$obj.Name}}, err error) {
	item = &{{$obj.Name}}{}
	newItem := &{{$obj.Name}}{}
	isChange := false

	now := time.Now()
	timestampMillis := now.UnixNano() / 1e6
	principalID := GetPrincipalIDFromContext(ctx)
	tx := GetTransaction(ctx)

	// 创建事件记录
	event := NewEvent(EventMetadata{
		Type:        EventTypeUpdated,
		Entity:      "{{$obj.Name}}",
		EntityID:    id,
		Date:        timestampMillis,
		PrincipalID: principalID,
	})

	// 解析输入变更
	var changes {{$obj.Name}}Changes
	if err = ApplyChanges(input, &changes); err != nil {
		return
	}

	// 验证必填字段
	if err = CheckStructFieldIsEmpty(item, input); err != nil {
		return nil, err
	}

	// ========== 验证关系字段冲突 ==========
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

	// 获取现有实体
	if err = GetItem(ctx, tx, TableName("{{$obj.TableName}}", ctx), item, &id); err != nil {
		return nil, err
	}

	// 更新 UpdatedBy
	if item.UpdatedBy != nil && principalID != nil && *item.UpdatedBy != *principalID {
		newItem.UpdatedBy = principalID
	}

	// ========== 处理 ManyToOne/OneToOne 关系 ==========
	{{range $rel := .Relationships}}
	{{if $rel.IsToOne}}
	{{if or $rel.IsManyToOne $rel.IsOneToOne}}
	// ---------- {{if $rel.IsManyToOne}}ManyToOne{{else}}OneToOne{{end}}: {{$rel.Name}} ----------
	if _, ok := input["{{$rel.Name}}"]; ok && !utils.IsNil(input["{{$rel.Name}}"]) {
		v := changes.{{$rel.MethodName}}

		if !utils.IsEmpty(v.ID) {
			// 更新现有关联对象
			v.UpdatedAt = &timestampMillis
			v.UpdatedBy = principalID

			if err := auth.CheckAuthorization(ctx, "Update{{$rel.TargetType}}"); err != nil {
				return item, errors.New("Update{{$rel.TargetType}} " + err.Error())
			}
			if err := auth.CheckAuthorization(ctx, "{{$rel.TargetType}}"); err != nil {
				return item, errors.New("{{$rel.TargetType}} Detail " + err.Error())
			}

			{{$rel.Name}}Input := utils.StructToMap(*v)
			if _, err := r.Handlers.Update{{$rel.TargetType}}(ctx, r, v.ID, {{$rel.Name}}Input); err != nil {
				return item, errors.New("{{$rel.TargetType}} ID " + v.ID + " " + err.Error())
			}

			// 更新外键
			{{if $rel.IsNonNull}}
			item.{{$rel.MethodName}}ID = v.ID
			newItem.{{$rel.MethodName}}ID = v.ID
			{{else}}
			item.{{$rel.MethodName}}ID = &v.ID
			newItem.{{$rel.MethodName}}ID = &v.ID
			{{end}}
			isChange = true
		} else {
			// 创建新关联对象
			if err := auth.CheckAuthorization(ctx, "Create{{$rel.TargetType}}"); err != nil {
				return item, errors.New("Create{{$rel.TargetType}} " + err.Error())
			}
			
			v.ID = uuid.Must(uuid.NewV4()).String()
			v.CreatedAt = timestampMillis
			v.CreatedBy = principalID

			// 保存新关联对象
			if err := tx.Omit(clause.Associations).Table(TableName("{{$rel.TargetTypeToSnakeName}}s", ctx)).Create(v).Error; err != nil {
				return item, err
			}

			// 更新外键
			{{if $rel.IsNonNull}}
			item.{{$rel.MethodName}}ID = v.ID
			newItem.{{$rel.MethodName}}ID = v.ID
			{{else}}
			item.{{$rel.MethodName}}ID = &v.ID
			newItem.{{$rel.MethodName}}ID = &v.ID
			{{end}}
			isChange = true
		}
	}
	{{end}}
	{{end}}
	{{end}}

	// ========== 处理普通字段 ==========
	{{range $col := .Columns}}
	{{if and (not $col.IsHasUpperId) $col.IsUpdatable}}
	{{if $col.IsOptional}}
	if _, ok := input["{{$col.Name}}"]; ok && (item.{{$col.MethodName}} != changes.{{$col.MethodName}}) && (item.{{$col.MethodName}} == nil || changes.{{$col.MethodName}} == nil || *item.{{$col.MethodName}} != *changes.{{$col.MethodName}}) && !utils.IsEmpty(input["{{$col.Name}}"]) {
	{{else}}
	if _, ok := input["{{$col.Name}}"]; ok && (item.{{$col.MethodName}} != changes.{{$col.MethodName}}){{if $col.IsOptional}} && (item.{{$col.MethodName}} == nil || changes.{{$col.MethodName}} == nil || *item.{{$col.MethodName}} != *changes.{{$col.MethodName}}){{end}} {
	{{end}}
		{{if $col.IsRelationshipIdentifier}}
		if err := tx.Select("id").Where("id = ?", input["{{$col.Name}}"]).First(&{{$col.RelationshipTypeName}}{}).Error; err != nil {
			return nil, fmt.Errorf("{{$col.Name}} " + err.Error())
		}
		{{end}}
		event.AddOldValue("{{$col.Name}}", item.{{$col.MethodName}})
		event.AddNewValue("{{$col.Name}}", changes.{{$col.MethodName}})
		item.{{$col.MethodName}} = changes.{{$col.MethodName}}
		newItem.{{$col.MethodName}} = changes.{{$col.MethodName}}
		isChange = true
	}
	{{end}}
	{{end}}

	// ========== 保存主实体变更 ==========
	if isChange {
		if err := tx.Table(TableName("{{$obj.TableName}}", ctx)).Where("id = ?", id).Updates(newItem).Error; err != nil {
			return item, err
		}
	}

	// ========== 处理 OneToMany/ManyToMany 关系 ==========
	{{range $rel := .Relationships}}
	{{if $rel.IsToMany}}
	// ---------- ToMany: {{$rel.Name}} ----------
	
	// 方式1：通过 IDs 关联
	if ids, ok := input["{{$rel.Name}}Ids"]; ok && !utils.IsNil(input["{{$rel.Name}}Ids"]) {
		items := []*{{$rel.TargetType}}{}
		itemIds := []string{}
		findIds := []string{}

		for _, v := range ids.([]interface{}) {
			itemIds = append(itemIds, v.(string))
		}

		if len(itemIds) > 0 {
			if err := auth.CheckAuthorization(ctx, "{{$rel.TargetType}}"); err != nil {
				return item, errors.New("{{$rel.TargetType}} Detail " + err.Error())
			}
			if err := tx.Find(&items, "id IN (?)", itemIds).Error; err != nil {
				return item, err
			}
			for _, v := range items {
				findIds = append(findIds, v.ID)
			}

			differenceIds := utils.Difference(itemIds, findIds)
			if len(differenceIds) > 0 {
				return item, fmt.Errorf("{{$rel.Name}}Ids " + strings.Join(differenceIds, ",") + " not found")
			}

			{{if $rel.IsManyToMany}}
			if err := tx.Model(item).Association("{{$rel.MethodName}}").Replace(items); err != nil {
				return item, err
			}
			{{else}}
			// OneToMany: 先清除旧关联，再设置新关联
			if err := tx.Model(&{{$rel.TargetType}}{}).Where("{{$rel.ToSnakeRelationshipName}}_id = ?", item.ID).Update("{{$rel.ToSnakeRelationshipName}}_id", nil).Error; err != nil {
				return item, err
			}
			for _, relItem := range items {
				if err := tx.Model(relItem).Update("{{$rel.ToSnakeRelationshipName}}_id", item.ID).Error; err != nil {
					return item, err
				}
			}
			{{end}}
		} else {
			// 清空关联
			{{if $rel.IsManyToMany}}
			if err := tx.Model(item).Association("{{$rel.MethodName}}").Clear(); err != nil {
				return item, err
			}
			{{else}}
			if err := tx.Model(&{{$rel.TargetType}}{}).Where("{{$rel.ToSnakeRelationshipName}}_id = ?", item.ID).Update("{{$rel.ToSnakeRelationshipName}}_id", nil).Error; err != nil {
				return item, err
			}
			{{end}}
		}
		event.AddNewValue("{{$rel.Name}}", items)
	}

	// 方式2：通过嵌套对象创建/更新
	if _, ok := input["{{$rel.Name}}"]; ok && !utils.IsNil(input["{{$rel.Name}}"]) {
		new{{$rel.MethodName}} := []*{{$rel.TargetType}}{}
		update{{$rel.MethodName}} := []*{{$rel.TargetType}}{}

		for index, v := range changes.{{$rel.MethodName}} {
			weight := int64(index + 1)
			v.Weight = &weight

			if !utils.IsEmpty(v.ID) {
				// 更新现有记录
				v.UpdatedAt = &timestampMillis
				v.UpdatedBy = principalID

				if err := auth.CheckAuthorization(ctx, "Update{{$rel.TargetType}}"); err != nil {
					return item, errors.New("Update{{$rel.TargetType}} " + err.Error())
				}
				if err := auth.CheckAuthorization(ctx, "{{$rel.TargetType}}"); err != nil {
					return item, errors.New("{{$rel.TargetType}} Detail " + err.Error())
				}

				{{$rel.Name}}Input := utils.StructToMap(*v)
				if _, err := r.Handlers.Update{{$rel.TargetType}}(ctx, r, {{$rel.Name}}Input["id"].(string), {{$rel.Name}}Input); err != nil {
					return item, errors.New("{{$rel.TargetType}} ID " + v.ID + " " + err.Error())
				}

				{{if not $rel.IsManyToMany}}
				if err := tx.Model(v).Update("{{$rel.ToSnakeRelationshipName}}_id", item.ID).Error; err != nil {
					return item, err
				}
				{{end}}

				update{{$rel.MethodName}} = append(update{{$rel.MethodName}}, v)
			} else {
				// 创建新记录
				if err := auth.CheckAuthorization(ctx, "Create{{$rel.TargetType}}"); err != nil {
					return item, errors.New("Create{{$rel.TargetType}} " + err.Error())
				}
				
				v.ID = uuid.Must(uuid.NewV4()).String()
				v.CreatedAt = timestampMillis
				v.CreatedBy = principalID

				{{if not $rel.IsManyToMany}}
				{{if $rel.InverseRelationship.IsNonNull}}
				v.{{$rel.UpperRelationshipName}}ID = item.ID
				{{else}}
				v.{{$rel.UpperRelationshipName}}ID = &item.ID
				{{end}}
				{{end}}

				if err := tx.Omit(clause.Associations).Table(TableName("{{$rel.TargetTypeToSnakeName}}s", ctx)).Create(v).Error; err != nil {
					return item, err
				}

				new{{$rel.MethodName}} = append(new{{$rel.MethodName}}, v)
			}
		}

		allItems := append(update{{$rel.MethodName}}, new{{$rel.MethodName}}...)
		
		{{if $rel.IsManyToMany}}
		if err := tx.Model(item).Association("{{$rel.MethodName}}").Replace(allItems); err != nil {
			return item, err
		}
		{{end}}

		event.AddNewValue("{{$rel.Name}}", allItems)
	}
	{{end}}
	{{end}}

	// 记录事件
	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}

// ============================================================
// {{$obj.Name}} - Delete
// ============================================================

// Delete{{$obj.Name}}Func 执行删除或恢复操作
func Delete{{$obj.Name}}Func(ctx context.Context, r *GeneratedResolver, id string, operationType string, unscoped *bool) (err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item := &{{$obj.Name}}{}
	now := time.Now()
	tx := GetTransaction(ctx)

	// 检查主从关系约束
	{{range $rel := .Relationships}}
	{{if $rel.IsMaster}}
	if err := tx.Where("is_delete = ? and {{$rel.ToSnakeRelationshipName}}_id = ?", 1, id).First(&{{$rel.TargetType}}{}).Error; err == nil {
		return fmt.Errorf("{{$rel.TargetType}} exists, cannot be deleted")
	}
	{{end}}
	{{end}}

	// 确定操作类型
	var status int64 = 1
	var isDelete int64 = 2
	if operationType == "recovery" {
		isDelete = 1
		status = 2
	}

	// 获取现有实体
	if err = tx.Unscoped().Table(TableName("{{$obj.TableName}}", ctx)).Where("is_delete = ? and id = ?", status, id).First(item).Error; err != nil {
		return err
	}

	deletedAt := now.UnixNano() / 1e6

	// 创建事件记录
	event := NewEvent(EventMetadata{
		Type:        EventTypeDeleted,
		Entity:      "{{$obj.Name}}",
		EntityID:    id,
		Date:        deletedAt,
		PrincipalID: principalID,
	})

	// 执行删除或恢复
	if operationType == "recovery" {
		if err := tx.Unscoped().Table(TableName("{{$obj.TableName}}", ctx)).Model(&item).Updates(map[string]interface{}{
			"IsDelete":  1,
			"DeletedAt": nil,
			"DeletedBy": nil,
		}).Error; err != nil {
			return err
		}
	} else {
		if unscoped != nil && *unscoped {
			// 物理删除
			if err := tx.Unscoped().Table(TableName("{{$obj.TableName}}", ctx)).Model(&item).Delete(item).Error; err != nil {
				return err
			}
		} else {
			// 软删除
			if err := tx.Model(&item).Table(TableName("{{$obj.TableName}}", ctx)).Updates({{$obj.Name}}{
				IsDelete:  &isDelete,
				DeletedAt: &deletedAt,
				DeletedBy: principalID,
				UpdatedBy: principalID,
			}).Error; err != nil {
				return err
			}
		}
	}

	// 记录事件
	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}

// Delete{{$obj.PluralName}} 批量删除 {{$obj.Name}} 实体
func (r *GeneratedMutationResolver) Delete{{$obj.PluralName}}(ctx context.Context, id []string, unscoped *bool) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.Delete{{$obj.PluralName}}(ctx, r.GeneratedResolver, id, unscoped)
	if err != nil {
		RollbackMutationContext(ctx, r.GeneratedResolver)
		return done, err
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

// Delete{{$obj.PluralName}}Handler 处理批量删除逻辑
func Delete{{$obj.PluralName}}Handler(ctx context.Context, r *GeneratedResolver, id []string, unscoped *bool) (bool, error) {
	for _, itemID := range id {
		if err := Delete{{$obj.Name}}Func(ctx, r, itemID, "delete", unscoped); err != nil {
			return false, err
		}
	}
	return true, nil
}

// ============================================================
// {{$obj.Name}} - Recovery
// ============================================================

// Recovery{{$obj.PluralName}} 批量恢复 {{$obj.Name}} 实体
func (r *GeneratedMutationResolver) Recovery{{$obj.PluralName}}(ctx context.Context, id []string) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.Recovery{{$obj.PluralName}}(ctx, r.GeneratedResolver, id)
	if err != nil {
		return done, err
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

// Recovery{{$obj.PluralName}}Handler 处理批量恢复逻辑
func Recovery{{$obj.PluralName}}Handler(ctx context.Context, r *GeneratedResolver, id []string) (bool, error) {
	unscoped := false
	for _, itemID := range id {
		if err := Delete{{$obj.Name}}Func(ctx, r, itemID, "recovery", &unscoped); err != nil {
			return false, err
		}
	}
	return true, nil
}

{{end}}
`
