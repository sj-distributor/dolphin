package gen

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gofrs/uuid"
	"github.com/sj-distributor/dolphin-example/auth"
	"github.com/sj-distributor/dolphin-example/utils"
	"gorm.io/gorm/clause"
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

// ============================================================
// User - Create
// ============================================================

// CreateUser 创建 User 实体的解析器入口
func (r *GeneratedMutationResolver) CreateUser(ctx context.Context, input map[string]interface{}) (item *User, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreateUser(ctx, r.GeneratedResolver, input)
	if err != nil {
		RollbackMutationContext(ctx, r.GeneratedResolver)
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// CreateUserHandler 处理 User 创建逻辑
func CreateUserHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *User, err error) {
	item = &User{}
	now := time.Now()
	timestampMillis := now.UnixNano() / 1e6
	principalID := GetPrincipalIDFromContext(ctx)
	tx := GetTransaction(ctx)

	// 创建事件记录
	event := NewEvent(EventMetadata{
		Type:        EventTypeCreated,
		Entity:      "User",
		EntityID:    item.ID,
		Date:        timestampMillis,
		PrincipalID: principalID,
	})

	// 解析输入变更
	var changes UserChanges
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

	// ToMany: tasks - 不能同时传入 IDs 和嵌套对象
	if !utils.IsNil(input["tasks"]) && !utils.IsNil(input["tasksIds"]) {
		return nil, fmt.Errorf("tasksIds and tasks cannot coexist")
	}

	// ========== 处理 ManyToOne/OneToOne 关系（当前实体持有外键） ==========

	// ========== 处理普通字段 ==========

	if _, ok := input["phone"]; ok {

		item.Phone = changes.Phone

		event.AddNewValue("phone", changes.Phone)
	}

	if _, ok := input["password"]; ok {

		item.Password = changes.Password

		event.AddNewValue("password", changes.Password)
	}

	if _, ok := input["email"]; ok && changes.Email != nil {

		item.Email = changes.Email

		event.AddNewValue("email", changes.Email)
	}

	if _, ok := input["nickname"]; ok && changes.Nickname != nil {

		item.Nickname = changes.Nickname

		event.AddNewValue("nickname", changes.Nickname)
	}

	if _, ok := input["age"]; ok && changes.Age != nil {

		item.Age = changes.Age

		event.AddNewValue("age", changes.Age)
	}

	if _, ok := input["lastName"]; ok && changes.LastName != nil {

		item.LastName = changes.LastName

		event.AddNewValue("lastName", changes.LastName)
	}

	if _, ok := input["isDelete"]; ok && changes.IsDelete != nil {

		item.IsDelete = changes.IsDelete

		event.AddNewValue("isDelete", changes.IsDelete)
	}

	if _, ok := input["weight"]; ok && changes.Weight != nil {

		item.Weight = changes.Weight

		event.AddNewValue("weight", changes.Weight)
	}

	if _, ok := input["state"]; ok && changes.State != nil {

		item.State = changes.State

		event.AddNewValue("state", changes.State)
	}

	// ========== 保存主实体 ==========
	if err := tx.Omit(clause.Associations).Table(TableName("users", ctx)).Create(item).Error; err != nil {
		return item, err
	}

	// ========== 处理 OneToMany/ManyToMany 关系（关联表持有外键或中间表） ==========

	// ---------- ToMany: tasks (OneToMany 外键在 Task.user_id) ----------

	// 方式1：通过 IDs 关联现有记录
	if ids, ok := input["tasksIds"]; ok && !utils.IsNil(input["tasksIds"]) {
		items := []*Task{}
		itemIds := []string{}
		findIds := []string{}

		for _, v := range ids.([]interface{}) {
			itemIds = append(itemIds, v.(string))
		}

		if len(itemIds) > 0 {
			// 权限检查
			if err := auth.CheckAuthorization(ctx, "Task"); err != nil {
				return item, fmt.Errorf("Task Detail: %w", err)
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
				return item, fmt.Errorf("tasksIds %s not found", strings.Join(differenceIds, ","))
			}

			// OneToMany: 更新关联记录的外键
			for _, relItem := range items {
				if err := tx.Model(relItem).Update("user_id", item.ID).Error; err != nil {
					return item, err
				}
			}

		}
		event.AddNewValue("tasks", items)
	}

	// 方式2：通过嵌套对象创建/更新
	if _, ok := input["tasks"]; ok && !utils.IsNil(input["tasks"]) {
		newTasks := []*Task{}
		updateTasks := []*Task{}

		hasCreateTasks := false
		hasUpdateTasks := false

		for index, v := range changes.Tasks {
			weight := int64(index + 1)
			v.Weight = &weight

			if !utils.IsEmpty(v.ID) {
				// 更新现有记录
				v.UpdatedAt = &timestampMillis
				v.UpdatedBy = principalID

				if !hasUpdateTasks {
					if err := auth.CheckAuthorization(ctx, "UpdateTask"); err != nil {
						return item, fmt.Errorf("UpdateTask: %w", err)
					}
					if err := auth.CheckAuthorization(ctx, "Task"); err != nil {
						return item, fmt.Errorf("Task Detail: %w", err)
					}
					hasUpdateTasks = true
				}

				tasksInput := utils.StructToMap(*v)
				if _, err := r.Handlers.UpdateTask(ctx, r, tasksInput["id"].(string), tasksInput); err != nil {
					return item, fmt.Errorf("Task ID %s: %w", v.ID, err)
				}

				// OneToMany: 设置外键指向当前实体
				if err := tx.Model(v).Update("user_id", item.ID).Error; err != nil {
					return item, err
				}

				updateTasks = append(updateTasks, v)
			} else {
				// 创建新记录
				if !hasCreateTasks {
					if err := auth.CheckAuthorization(ctx, "CreateTask"); err != nil {
						return item, fmt.Errorf("CreateTask: %w", err)
					}
					hasCreateTasks = true
				}

				v.ID = uuid.Must(uuid.NewV4()).String()
				v.CreatedAt = timestampMillis
				v.CreatedBy = principalID

				// OneToMany: 设置外键指向当前实体

				v.UserID = &item.ID

				// 保存新记录
				if err := tx.Omit(clause.Associations).Table(TableName("tasks", ctx)).Create(v).Error; err != nil {
					return item, err
				}

				newTasks = append(newTasks, v)
			}
		}

		allItems := append(updateTasks, newTasks...)

		event.AddNewValue("tasks", allItems)
	}

	// 记录事件
	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}

// ============================================================
// User - Update
// ============================================================

// UpdateUser 更新 User 实体的解析器入口
func (r *GeneratedMutationResolver) UpdateUser(ctx context.Context, id string, input map[string]interface{}) (item *User, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdateUser(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		RollbackMutationContext(ctx, r.GeneratedResolver)
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// UpdateUserHandler 处理 User 更新逻辑
func UpdateUserHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *User, err error) {
	item = &User{}
	newItem := &User{}
	isChange := false

	now := time.Now()
	timestampMillis := now.UnixNano() / 1e6
	principalID := GetPrincipalIDFromContext(ctx)
	tx := GetTransaction(ctx)

	// 创建事件记录
	event := NewEvent(EventMetadata{
		Type:        EventTypeUpdated,
		Entity:      "User",
		EntityID:    id,
		Date:        timestampMillis,
		PrincipalID: principalID,
	})

	// 解析输入变更
	var changes UserChanges
	if err = ApplyChanges(input, &changes); err != nil {
		return
	}

	// 验证必填字段
	if err = CheckStructFieldIsEmpty(item, input); err != nil {
		return nil, err
	}

	// ========== 验证关系字段冲突 ==========

	if !utils.IsNil(input["tasks"]) && !utils.IsNil(input["tasksIds"]) {
		return nil, fmt.Errorf("tasksIds and tasks cannot coexist")
	}

	// 获取现有实体
	if err = GetItem(ctx, tx, TableName("users", ctx), item, &id); err != nil {
		return nil, err
	}

	// 设置审计字段
	newItem.UpdatedAt = &timestampMillis
	newItem.UpdatedBy = principalID

	// 字段变更追踪
	changedFields := []string{}

	// ========== 处理 ManyToOne/OneToOne 关系 ==========

	// ========== 处理普通字段 ==========
	// changedFields := []string{} (Moved to top)

	if _, ok := input["id"]; ok {
		// 只要 input 中包含该字段，且值发生了变化（包括变为 null），就进行更新
		if item.ID != changes.ID {

			event.AddOldValue("id", item.ID)
			event.AddNewValue("id", changes.ID)

			item.ID = changes.ID
			newItem.ID = changes.ID
			changedFields = append(changedFields, "id")
			isChange = true
		}
	}

	if _, ok := input["phone"]; ok {
		// 只要 input 中包含该字段，且值发生了变化（包括变为 null），就进行更新
		if item.Phone != changes.Phone {

			event.AddOldValue("phone", item.Phone)
			event.AddNewValue("phone", changes.Phone)

			item.Phone = changes.Phone
			newItem.Phone = changes.Phone
			changedFields = append(changedFields, "phone")
			isChange = true
		}
	}

	if _, ok := input["password"]; ok {
		// 只要 input 中包含该字段，且值发生了变化（包括变为 null），就进行更新
		if item.Password != changes.Password {

			event.AddOldValue("password", item.Password)
			event.AddNewValue("password", changes.Password)

			item.Password = changes.Password
			newItem.Password = changes.Password
			changedFields = append(changedFields, "password")
			isChange = true
		}
	}

	if _, ok := input["email"]; ok {
		// 只要 input 中包含该字段，且值发生了变化（包括变为 null），就进行更新
		if (item.Email != changes.Email) && (item.Email == nil || changes.Email == nil || *item.Email != *changes.Email) {

			event.AddOldValue("email", item.Email)
			event.AddNewValue("email", changes.Email)

			item.Email = changes.Email
			newItem.Email = changes.Email
			changedFields = append(changedFields, "email")
			isChange = true
		}
	}

	if _, ok := input["nickname"]; ok {
		// 只要 input 中包含该字段，且值发生了变化（包括变为 null），就进行更新
		if (item.Nickname != changes.Nickname) && (item.Nickname == nil || changes.Nickname == nil || *item.Nickname != *changes.Nickname) {

			event.AddOldValue("nickname", item.Nickname)
			event.AddNewValue("nickname", changes.Nickname)

			item.Nickname = changes.Nickname
			newItem.Nickname = changes.Nickname
			changedFields = append(changedFields, "nickname")
			isChange = true
		}
	}

	if _, ok := input["age"]; ok {
		// 只要 input 中包含该字段，且值发生了变化（包括变为 null），就进行更新
		if (item.Age != changes.Age) && (item.Age == nil || changes.Age == nil || *item.Age != *changes.Age) {

			event.AddOldValue("age", item.Age)
			event.AddNewValue("age", changes.Age)

			item.Age = changes.Age
			newItem.Age = changes.Age
			changedFields = append(changedFields, "age")
			isChange = true
		}
	}

	if _, ok := input["lastName"]; ok {
		// 只要 input 中包含该字段，且值发生了变化（包括变为 null），就进行更新
		if (item.LastName != changes.LastName) && (item.LastName == nil || changes.LastName == nil || *item.LastName != *changes.LastName) {

			event.AddOldValue("lastName", item.LastName)
			event.AddNewValue("lastName", changes.LastName)

			item.LastName = changes.LastName
			newItem.LastName = changes.LastName
			changedFields = append(changedFields, "last_name")
			isChange = true
		}
	}

	if _, ok := input["isDelete"]; ok {
		// 只要 input 中包含该字段，且值发生了变化（包括变为 null），就进行更新
		if (item.IsDelete != changes.IsDelete) && (item.IsDelete == nil || changes.IsDelete == nil || *item.IsDelete != *changes.IsDelete) {

			event.AddOldValue("isDelete", item.IsDelete)
			event.AddNewValue("isDelete", changes.IsDelete)

			item.IsDelete = changes.IsDelete
			newItem.IsDelete = changes.IsDelete
			changedFields = append(changedFields, "is_delete")
			isChange = true
		}
	}

	if _, ok := input["weight"]; ok {
		// 只要 input 中包含该字段，且值发生了变化（包括变为 null），就进行更新
		if (item.Weight != changes.Weight) && (item.Weight == nil || changes.Weight == nil || *item.Weight != *changes.Weight) {

			event.AddOldValue("weight", item.Weight)
			event.AddNewValue("weight", changes.Weight)

			item.Weight = changes.Weight
			newItem.Weight = changes.Weight
			changedFields = append(changedFields, "weight")
			isChange = true
		}
	}

	if _, ok := input["state"]; ok {
		// 只要 input 中包含该字段，且值发生了变化（包括变为 null），就进行更新
		if (item.State != changes.State) && (item.State == nil || changes.State == nil || *item.State != *changes.State) {

			event.AddOldValue("state", item.State)
			event.AddNewValue("state", changes.State)

			item.State = changes.State
			newItem.State = changes.State
			changedFields = append(changedFields, "state")
			isChange = true
		}
	}

	// ========== 保存主实体变更 ==========
	if isChange {
		changedFields = append(changedFields, "updated_at", "updated_by")

		if err := tx.Table(TableName("users", ctx)).Where("id = ?", id).Select(changedFields).Updates(newItem).Error; err != nil {
			return item, err
		}
	}

	// ========== 处理 OneToMany/ManyToMany 关系 ==========

	// ---------- ToMany: tasks ----------

	// 方式1：通过 IDs 关联
	if ids, ok := input["tasksIds"]; ok && !utils.IsNil(input["tasksIds"]) {
		items := []*Task{}
		itemIds := []string{}
		findIds := []string{}

		for _, v := range ids.([]interface{}) {
			itemIds = append(itemIds, v.(string))
		}

		if len(itemIds) > 0 {
			if err := auth.CheckAuthorization(ctx, "Task"); err != nil {
				return item, fmt.Errorf("Task Detail: %w", err)
			}
			if err := tx.Find(&items, "id IN (?)", itemIds).Error; err != nil {
				return item, err
			}
			for _, v := range items {
				findIds = append(findIds, v.ID)
			}

			differenceIds := utils.Difference(itemIds, findIds)
			if len(differenceIds) > 0 {
				return item, fmt.Errorf("tasksIds %s not found", strings.Join(differenceIds, ","))
			}

			// OneToMany: 先清除旧关联，再设置新关联
			if err := tx.Model(&Task{}).Where("user_id = ?", item.ID).Update("user_id", nil).Error; err != nil {
				return item, err
			}
			for _, relItem := range items {
				if err := tx.Model(relItem).Update("user_id", item.ID).Error; err != nil {
					return item, err
				}
			}

		} else {
			// 清空关联

			if err := tx.Model(&Task{}).Where("user_id = ?", item.ID).Update("user_id", nil).Error; err != nil {
				return item, err
			}

		}
		event.AddNewValue("tasks", items)
	}

	// 方式2：通过嵌套对象创建/更新
	if _, ok := input["tasks"]; ok && !utils.IsNil(input["tasks"]) {
		newTasks := []*Task{}
		updateTasks := []*Task{}

		// OneToMany: 先清除旧关联（与 IDs 方式行为一致）
		if err := tx.Model(&Task{}).Where("user_id = ?", item.ID).Update("user_id", nil).Error; err != nil {
			return item, err
		}

		hasCreateTasks := false
		hasUpdateTasks := false

		for index, v := range changes.Tasks {
			weight := int64(index + 1)
			v.Weight = &weight

			if !utils.IsEmpty(v.ID) {
				// 更新现有记录
				v.UpdatedAt = &timestampMillis
				v.UpdatedBy = principalID

				if !hasUpdateTasks {
					if err := auth.CheckAuthorization(ctx, "UpdateTask"); err != nil {
						return item, fmt.Errorf("UpdateTask: %w", err)
					}
					if err := auth.CheckAuthorization(ctx, "Task"); err != nil {
						return item, fmt.Errorf("Task Detail: %w", err)
					}
					hasUpdateTasks = true
				}

				tasksInput := utils.StructToMap(*v)
				if _, err := r.Handlers.UpdateTask(ctx, r, tasksInput["id"].(string), tasksInput); err != nil {
					return item, fmt.Errorf("Task ID %s: %w", v.ID, err)
				}

				if err := tx.Model(v).Update("user_id", item.ID).Error; err != nil {
					return item, err
				}

				updateTasks = append(updateTasks, v)
			} else {
				// 创建新记录
				if !hasCreateTasks {
					if err := auth.CheckAuthorization(ctx, "CreateTask"); err != nil {
						return item, fmt.Errorf("CreateTask: %w", err)
					}
					hasCreateTasks = true
				}

				v.ID = uuid.Must(uuid.NewV4()).String()
				v.CreatedAt = timestampMillis
				v.CreatedBy = principalID

				v.UserID = &item.ID

				if err := tx.Omit(clause.Associations).Table(TableName("tasks", ctx)).Create(v).Error; err != nil {
					return item, err
				}

				newTasks = append(newTasks, v)
			}
		}

		allItems := append(updateTasks, newTasks...)

		event.AddNewValue("tasks", allItems)
	}

	// 记录事件
	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}

// ============================================================
// User - Delete
// ============================================================

// DeleteUserFunc 执行删除或恢复操作
func DeleteUserFunc(ctx context.Context, r *GeneratedResolver, id string, operationType string, unscoped *bool) (err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item := &User{}
	now := time.Now()
	tx := GetTransaction(ctx)

	// 检查主从关系约束

	// 确定操作类型
	var status int64 = 1
	var isDelete int64 = 2
	if operationType == "recovery" {
		isDelete = 1
		status = 2
	}

	// 获取现有实体
	if err = tx.Unscoped().Table(TableName("users", ctx)).Where("is_delete = ? and id = ?", status, id).First(item).Error; err != nil {
		return err
	}

	deletedAt := now.UnixNano() / 1e6

	// 创建事件记录
	event := NewEvent(EventMetadata{
		Type:        EventTypeDeleted,
		Entity:      "User",
		EntityID:    id,
		Date:        deletedAt,
		PrincipalID: principalID,
	})

	// 执行删除或恢复
	if operationType == "recovery" {
		if err := tx.Unscoped().Table(TableName("users", ctx)).Model(&item).Updates(map[string]interface{}{
			"IsDelete":  1,
			"DeletedAt": nil,
			"DeletedBy": nil,
		}).Error; err != nil {
			return err
		}
	} else {
		if unscoped != nil && *unscoped {
			// 物理删除
			if err := tx.Unscoped().Table(TableName("users", ctx)).Model(&item).Delete(item).Error; err != nil {
				return err
			}
		} else {
			// 软删除
			if err := tx.Model(&item).Table(TableName("users", ctx)).Updates(User{
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

// DeleteUsers 批量删除 User 实体
func (r *GeneratedMutationResolver) DeleteUsers(ctx context.Context, id []string, unscoped *bool) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteUsers(ctx, r.GeneratedResolver, id, unscoped)
	if err != nil {
		RollbackMutationContext(ctx, r.GeneratedResolver)
		return done, err
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

// DeleteUsersHandler 处理批量删除逻辑
func DeleteUsersHandler(ctx context.Context, r *GeneratedResolver, id []string, unscoped *bool) (bool, error) {
	for _, itemID := range id {
		if err := DeleteUserFunc(ctx, r, itemID, "delete", unscoped); err != nil {
			return false, err
		}
	}
	return true, nil
}

// ============================================================
// User - Recovery
// ============================================================

// RecoveryUsers 批量恢复 User 实体
func (r *GeneratedMutationResolver) RecoveryUsers(ctx context.Context, id []string) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.RecoveryUsers(ctx, r.GeneratedResolver, id)
	if err != nil {
		return done, err
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

// RecoveryUsersHandler 处理批量恢复逻辑
func RecoveryUsersHandler(ctx context.Context, r *GeneratedResolver, id []string) (bool, error) {
	unscoped := false
	for _, itemID := range id {
		if err := DeleteUserFunc(ctx, r, itemID, "recovery", &unscoped); err != nil {
			return false, err
		}
	}
	return true, nil
}

// ============================================================
// Task - Create
// ============================================================

// CreateTask 创建 Task 实体的解析器入口
func (r *GeneratedMutationResolver) CreateTask(ctx context.Context, input map[string]interface{}) (item *Task, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreateTask(ctx, r.GeneratedResolver, input)
	if err != nil {
		RollbackMutationContext(ctx, r.GeneratedResolver)
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// CreateTaskHandler 处理 Task 创建逻辑
func CreateTaskHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Task, err error) {
	item = &Task{}
	now := time.Now()
	timestampMillis := now.UnixNano() / 1e6
	principalID := GetPrincipalIDFromContext(ctx)
	tx := GetTransaction(ctx)

	// 创建事件记录
	event := NewEvent(EventMetadata{
		Type:        EventTypeCreated,
		Entity:      "Task",
		EntityID:    item.ID,
		Date:        timestampMillis,
		PrincipalID: principalID,
	})

	// 解析输入变更
	var changes TaskChanges
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

	if !utils.IsNil(input["user"]) && !utils.IsNil(input["userId"]) {
		return nil, fmt.Errorf("userId and user cannot coexist")
	}

	// ========== 处理 ManyToOne/OneToOne 关系（当前实体持有外键） ==========

	// ========== 处理普通字段 ==========

	if _, ok := input["title"]; ok && changes.Title != nil {

		item.Title = changes.Title

		event.AddNewValue("title", changes.Title)
	}

	if _, ok := input["completed"]; ok && changes.Completed != nil {

		item.Completed = changes.Completed

		event.AddNewValue("completed", changes.Completed)
	}

	if _, ok := input["dueDate"]; ok && changes.DueDate != nil {

		item.DueDate = changes.DueDate

		event.AddNewValue("dueDate", changes.DueDate)
	}

	if _, ok := input["userId"]; ok && changes.UserID != nil {

		if !utils.IsNil(input["userId"]) {
			if err := tx.Select("id").Where("id = ?", input["userId"]).First(&User{}).Error; err != nil {
				return nil, fmt.Errorf("userId: %w", err)
			}
		}

		item.UserID = changes.UserID

		event.AddNewValue("userId", changes.UserID)
	}

	if _, ok := input["isDelete"]; ok && changes.IsDelete != nil {

		item.IsDelete = changes.IsDelete

		event.AddNewValue("isDelete", changes.IsDelete)
	}

	if _, ok := input["weight"]; ok && changes.Weight != nil {

		item.Weight = changes.Weight

		event.AddNewValue("weight", changes.Weight)
	}

	if _, ok := input["state"]; ok && changes.State != nil {

		item.State = changes.State

		event.AddNewValue("state", changes.State)
	}

	// ========== 保存主实体 ==========
	if err := tx.Omit(clause.Associations).Table(TableName("tasks", ctx)).Create(item).Error; err != nil {
		return item, err
	}

	// ========== 处理 OneToMany/ManyToMany 关系（关联表持有外键或中间表） ==========

	// 记录事件
	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}

// ============================================================
// Task - Update
// ============================================================

// UpdateTask 更新 Task 实体的解析器入口
func (r *GeneratedMutationResolver) UpdateTask(ctx context.Context, id string, input map[string]interface{}) (item *Task, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdateTask(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		RollbackMutationContext(ctx, r.GeneratedResolver)
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// UpdateTaskHandler 处理 Task 更新逻辑
func UpdateTaskHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Task, err error) {
	item = &Task{}
	newItem := &Task{}
	isChange := false

	now := time.Now()
	timestampMillis := now.UnixNano() / 1e6
	principalID := GetPrincipalIDFromContext(ctx)
	tx := GetTransaction(ctx)

	// 创建事件记录
	event := NewEvent(EventMetadata{
		Type:        EventTypeUpdated,
		Entity:      "Task",
		EntityID:    id,
		Date:        timestampMillis,
		PrincipalID: principalID,
	})

	// 解析输入变更
	var changes TaskChanges
	if err = ApplyChanges(input, &changes); err != nil {
		return
	}

	// 验证必填字段
	if err = CheckStructFieldIsEmpty(item, input); err != nil {
		return nil, err
	}

	// ========== 验证关系字段冲突 ==========

	if !utils.IsNil(input["user"]) && !utils.IsNil(input["userId"]) {
		return nil, fmt.Errorf("userId and user cannot coexist")
	}

	// 获取现有实体
	if err = GetItem(ctx, tx, TableName("tasks", ctx), item, &id); err != nil {
		return nil, err
	}

	// 设置审计字段
	newItem.UpdatedAt = &timestampMillis
	newItem.UpdatedBy = principalID

	// 字段变更追踪
	changedFields := []string{}

	// ========== 处理 ManyToOne/OneToOne 关系 ==========

	// ========== 处理普通字段 ==========
	// changedFields := []string{} (Moved to top)

	if _, ok := input["id"]; ok {
		// 只要 input 中包含该字段，且值发生了变化（包括变为 null），就进行更新
		if item.ID != changes.ID {

			event.AddOldValue("id", item.ID)
			event.AddNewValue("id", changes.ID)

			item.ID = changes.ID
			newItem.ID = changes.ID
			changedFields = append(changedFields, "id")
			isChange = true
		}
	}

	if _, ok := input["title"]; ok {
		// 只要 input 中包含该字段，且值发生了变化（包括变为 null），就进行更新
		if (item.Title != changes.Title) && (item.Title == nil || changes.Title == nil || *item.Title != *changes.Title) {

			event.AddOldValue("title", item.Title)
			event.AddNewValue("title", changes.Title)

			item.Title = changes.Title
			newItem.Title = changes.Title
			changedFields = append(changedFields, "title")
			isChange = true
		}
	}

	if _, ok := input["completed"]; ok {
		// 只要 input 中包含该字段，且值发生了变化（包括变为 null），就进行更新
		if (item.Completed != changes.Completed) && (item.Completed == nil || changes.Completed == nil || *item.Completed != *changes.Completed) {

			event.AddOldValue("completed", item.Completed)
			event.AddNewValue("completed", changes.Completed)

			item.Completed = changes.Completed
			newItem.Completed = changes.Completed
			changedFields = append(changedFields, "completed")
			isChange = true
		}
	}

	if _, ok := input["dueDate"]; ok {
		// 只要 input 中包含该字段，且值发生了变化（包括变为 null），就进行更新
		if (item.DueDate != changes.DueDate) && (item.DueDate == nil || changes.DueDate == nil || *item.DueDate != *changes.DueDate) {

			event.AddOldValue("dueDate", item.DueDate)
			event.AddNewValue("dueDate", changes.DueDate)

			item.DueDate = changes.DueDate
			newItem.DueDate = changes.DueDate
			changedFields = append(changedFields, "due_date")
			isChange = true
		}
	}

	if _, ok := input["userId"]; ok {
		// 只要 input 中包含该字段，且值发生了变化（包括变为 null），就进行更新
		if (item.UserID != changes.UserID) && (item.UserID == nil || changes.UserID == nil || *item.UserID != *changes.UserID) {

			if !utils.IsNil(input["userId"]) {
				if err := tx.Select("id").Where("id = ?", input["userId"]).First(&User{}).Error; err != nil {
					return nil, fmt.Errorf("userId: %w", err)
				}
			}

			event.AddOldValue("userId", item.UserID)
			event.AddNewValue("userId", changes.UserID)

			item.UserID = changes.UserID
			newItem.UserID = changes.UserID
			changedFields = append(changedFields, "user_id")
			isChange = true
		}
	}

	if _, ok := input["isDelete"]; ok {
		// 只要 input 中包含该字段，且值发生了变化（包括变为 null），就进行更新
		if (item.IsDelete != changes.IsDelete) && (item.IsDelete == nil || changes.IsDelete == nil || *item.IsDelete != *changes.IsDelete) {

			event.AddOldValue("isDelete", item.IsDelete)
			event.AddNewValue("isDelete", changes.IsDelete)

			item.IsDelete = changes.IsDelete
			newItem.IsDelete = changes.IsDelete
			changedFields = append(changedFields, "is_delete")
			isChange = true
		}
	}

	if _, ok := input["weight"]; ok {
		// 只要 input 中包含该字段，且值发生了变化（包括变为 null），就进行更新
		if (item.Weight != changes.Weight) && (item.Weight == nil || changes.Weight == nil || *item.Weight != *changes.Weight) {

			event.AddOldValue("weight", item.Weight)
			event.AddNewValue("weight", changes.Weight)

			item.Weight = changes.Weight
			newItem.Weight = changes.Weight
			changedFields = append(changedFields, "weight")
			isChange = true
		}
	}

	if _, ok := input["state"]; ok {
		// 只要 input 中包含该字段，且值发生了变化（包括变为 null），就进行更新
		if (item.State != changes.State) && (item.State == nil || changes.State == nil || *item.State != *changes.State) {

			event.AddOldValue("state", item.State)
			event.AddNewValue("state", changes.State)

			item.State = changes.State
			newItem.State = changes.State
			changedFields = append(changedFields, "state")
			isChange = true
		}
	}

	// ========== 保存主实体变更 ==========
	if isChange {
		changedFields = append(changedFields, "updated_at", "updated_by")

		if err := tx.Table(TableName("tasks", ctx)).Where("id = ?", id).Select(changedFields).Updates(newItem).Error; err != nil {
			return item, err
		}
	}

	// ========== 处理 OneToMany/ManyToMany 关系 ==========

	// 记录事件
	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}

// ============================================================
// Task - Delete
// ============================================================

// DeleteTaskFunc 执行删除或恢复操作
func DeleteTaskFunc(ctx context.Context, r *GeneratedResolver, id string, operationType string, unscoped *bool) (err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item := &Task{}
	now := time.Now()
	tx := GetTransaction(ctx)

	// 检查主从关系约束

	// 确定操作类型
	var status int64 = 1
	var isDelete int64 = 2
	if operationType == "recovery" {
		isDelete = 1
		status = 2
	}

	// 获取现有实体
	if err = tx.Unscoped().Table(TableName("tasks", ctx)).Where("is_delete = ? and id = ?", status, id).First(item).Error; err != nil {
		return err
	}

	deletedAt := now.UnixNano() / 1e6

	// 创建事件记录
	event := NewEvent(EventMetadata{
		Type:        EventTypeDeleted,
		Entity:      "Task",
		EntityID:    id,
		Date:        deletedAt,
		PrincipalID: principalID,
	})

	// 执行删除或恢复
	if operationType == "recovery" {
		if err := tx.Unscoped().Table(TableName("tasks", ctx)).Model(&item).Updates(map[string]interface{}{
			"IsDelete":  1,
			"DeletedAt": nil,
			"DeletedBy": nil,
		}).Error; err != nil {
			return err
		}
	} else {
		if unscoped != nil && *unscoped {
			// 物理删除
			if err := tx.Unscoped().Table(TableName("tasks", ctx)).Model(&item).Delete(item).Error; err != nil {
				return err
			}
		} else {
			// 软删除
			if err := tx.Model(&item).Table(TableName("tasks", ctx)).Updates(Task{
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

// DeleteTasks 批量删除 Task 实体
func (r *GeneratedMutationResolver) DeleteTasks(ctx context.Context, id []string, unscoped *bool) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteTasks(ctx, r.GeneratedResolver, id, unscoped)
	if err != nil {
		RollbackMutationContext(ctx, r.GeneratedResolver)
		return done, err
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

// DeleteTasksHandler 处理批量删除逻辑
func DeleteTasksHandler(ctx context.Context, r *GeneratedResolver, id []string, unscoped *bool) (bool, error) {
	for _, itemID := range id {
		if err := DeleteTaskFunc(ctx, r, itemID, "delete", unscoped); err != nil {
			return false, err
		}
	}
	return true, nil
}

// ============================================================
// Task - Recovery
// ============================================================

// RecoveryTasks 批量恢复 Task 实体
func (r *GeneratedMutationResolver) RecoveryTasks(ctx context.Context, id []string) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.RecoveryTasks(ctx, r.GeneratedResolver, id)
	if err != nil {
		return done, err
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

// RecoveryTasksHandler 处理批量恢复逻辑
func RecoveryTasksHandler(ctx context.Context, r *GeneratedResolver, id []string) (bool, error) {
	unscoped := false
	for _, itemID := range id {
		if err := DeleteTaskFunc(ctx, r, itemID, "recovery", &unscoped); err != nil {
			return false, err
		}
	}
	return true, nil
}
