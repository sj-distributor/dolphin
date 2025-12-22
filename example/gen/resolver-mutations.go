package gen

import (
	"context"
	"errors"
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

	// ToMany: userRoles - 不能同时传入 IDs 和嵌套对象
	if !utils.IsNil(input["userRoles"]) && !utils.IsNil(input["userRolesIds"]) {
		return nil, fmt.Errorf("userRolesIds and userRoles cannot coexist")
	}

	// ========== 处理 ManyToOne/OneToOne 关系（当前实体持有外键） ==========

	// ---------- OneToOne: profile (当前表持有外键 user.profile_id) ----------
	if _, ok := input["profile"]; ok && !utils.IsNil(input["profile"]) {
		v := changes.Profile

		if !utils.IsEmpty(v.ID) {
			// 更新现有关联对象
			v.UpdatedAt = &timestampMillis
			v.UpdatedBy = principalID

			if err := auth.CheckAuthorization(ctx, "UpdateProfile"); err != nil {
				return item, errors.New("UpdateProfile " + err.Error())
			}
			if err := auth.CheckAuthorization(ctx, "Profile"); err != nil {
				return item, errors.New("Profile Detail " + err.Error())
			}

			profileInput := utils.StructToMap(*v)
			if _, err := r.Handlers.UpdateProfile(ctx, r, v.ID, profileInput); err != nil {
				return item, errors.New("Profile ID " + v.ID + " " + err.Error())
			}

			// 设置外键

			item.ProfileID = &v.ID

		} else {
			// 创建新关联对象
			if err := auth.CheckAuthorization(ctx, "CreateProfile"); err != nil {
				return item, errors.New("CreateProfile " + err.Error())
			}

			v.ID = uuid.Must(uuid.NewV4()).String()
			v.CreatedAt = timestampMillis
			v.CreatedBy = principalID

			// 先保存关联对象
			if err := tx.Omit(clause.Associations).Table(TableName("profiles", ctx)).Create(v).Error; err != nil {
				return item, err
			}

			// 设置外键

			item.ProfileID = &v.ID

		}

		item.Profile = v
		event.AddNewValue("profile", item.Profile)
		event.AddNewValue("profileId", item.ProfileID)
	}

	// ========== 处理普通字段 ==========

	if _, ok := input["phone"]; ok && !utils.IsEmpty(input["phone"]) {

		if item.Phone != changes.Phone {

			item.Phone = changes.Phone

			event.AddNewValue("phone", changes.Phone)
		}
	}

	if _, ok := input["password"]; ok && !utils.IsEmpty(input["password"]) {

		if item.Password != changes.Password {

			item.Password = changes.Password

			event.AddNewValue("password", changes.Password)
		}
	}

	if _, ok := input["email"]; ok && changes.Email != nil {

		if (item.Email != changes.Email) || (*item.Email != *changes.Email) {

			item.Email = changes.Email

			event.AddNewValue("email", changes.Email)
		}
	}

	if _, ok := input["nickname"]; ok && changes.Nickname != nil {

		if (item.Nickname != changes.Nickname) || (*item.Nickname != *changes.Nickname) {

			item.Nickname = changes.Nickname

			event.AddNewValue("nickname", changes.Nickname)
		}
	}

	if _, ok := input["age"]; ok && changes.Age != nil {

		if (item.Age != changes.Age) || (*item.Age != *changes.Age) {

			item.Age = changes.Age

			event.AddNewValue("age", changes.Age)
		}
	}

	if _, ok := input["profileId"]; ok && changes.ProfileID != nil {

		if (item.ProfileID != changes.ProfileID) || (*item.ProfileID != *changes.ProfileID) {

			if !utils.IsNil(input["profileId"]) {
				if err := tx.Select("id").Where("id = ?", input["profileId"]).First(&Profile{}).Error; err != nil {
					return nil, fmt.Errorf("profileId " + err.Error())
				}
			}

			item.ProfileID = changes.ProfileID

			event.AddNewValue("profileId", changes.ProfileID)
		}
	}

	if _, ok := input["isDelete"]; ok && changes.IsDelete != nil {

		if (item.IsDelete != changes.IsDelete) || (*item.IsDelete != *changes.IsDelete) {

			item.IsDelete = changes.IsDelete

			event.AddNewValue("isDelete", changes.IsDelete)
		}
	}

	if _, ok := input["weight"]; ok && changes.Weight != nil {

		if (item.Weight != changes.Weight) || (*item.Weight != *changes.Weight) {

			item.Weight = changes.Weight

			event.AddNewValue("weight", changes.Weight)
		}
	}

	if _, ok := input["state"]; ok && changes.State != nil {

		if (item.State != changes.State) || (*item.State != *changes.State) {

			item.State = changes.State

			event.AddNewValue("state", changes.State)
		}
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
				return item, errors.New("Task Detail " + err.Error())
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
				return item, fmt.Errorf("tasksIds " + strings.Join(differenceIds, ",") + " not found")
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

		for index, v := range changes.Tasks {
			weight := int64(index + 1)
			v.Weight = &weight

			if !utils.IsEmpty(v.ID) {
				// 更新现有记录
				v.UpdatedAt = &timestampMillis
				v.UpdatedBy = principalID

				if err := auth.CheckAuthorization(ctx, "UpdateTask"); err != nil {
					return item, errors.New("UpdateTask " + err.Error())
				}
				if err := auth.CheckAuthorization(ctx, "Task"); err != nil {
					return item, errors.New("Task Detail " + err.Error())
				}

				tasksInput := utils.StructToMap(*v)
				if _, err := r.Handlers.UpdateTask(ctx, r, tasksInput["id"].(string), tasksInput); err != nil {
					return item, errors.New("Task ID " + v.ID + " " + err.Error())
				}

				// OneToMany: 设置外键指向当前实体
				if err := tx.Model(v).Update("user_id", item.ID).Error; err != nil {
					return item, err
				}

				updateTasks = append(updateTasks, v)
			} else {
				// 创建新记录
				if err := auth.CheckAuthorization(ctx, "CreateTask"); err != nil {
					return item, errors.New("CreateTask " + err.Error())
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

	// ---------- ToMany: userRoles (ManyToMany 中间表) ----------

	// 方式1：通过 IDs 关联现有记录
	if ids, ok := input["userRolesIds"]; ok && !utils.IsNil(input["userRolesIds"]) {
		items := []*UserRole{}
		itemIds := []string{}
		findIds := []string{}

		for _, v := range ids.([]interface{}) {
			itemIds = append(itemIds, v.(string))
		}

		if len(itemIds) > 0 {
			// 权限检查
			if err := auth.CheckAuthorization(ctx, "UserRole"); err != nil {
				return item, errors.New("UserRole Detail " + err.Error())
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
				return item, fmt.Errorf("userRolesIds " + strings.Join(differenceIds, ",") + " not found")
			}

			// ManyToMany: 使用 Replace 更新中间表
			if err := tx.Model(item).Association("UserRoles").Replace(items); err != nil {
				return item, err
			}

		}
		event.AddNewValue("userRoles", items)
	}

	// 方式2：通过嵌套对象创建/更新
	if _, ok := input["userRoles"]; ok && !utils.IsNil(input["userRoles"]) {
		newUserRoles := []*UserRole{}
		updateUserRoles := []*UserRole{}

		for index, v := range changes.UserRoles {
			weight := int64(index + 1)
			v.Weight = &weight

			if !utils.IsEmpty(v.ID) {
				// 更新现有记录
				v.UpdatedAt = &timestampMillis
				v.UpdatedBy = principalID

				if err := auth.CheckAuthorization(ctx, "UpdateUserRole"); err != nil {
					return item, errors.New("UpdateUserRole " + err.Error())
				}
				if err := auth.CheckAuthorization(ctx, "UserRole"); err != nil {
					return item, errors.New("UserRole Detail " + err.Error())
				}

				userRolesInput := utils.StructToMap(*v)
				if _, err := r.Handlers.UpdateUserRole(ctx, r, userRolesInput["id"].(string), userRolesInput); err != nil {
					return item, errors.New("UserRole ID " + v.ID + " " + err.Error())
				}

				updateUserRoles = append(updateUserRoles, v)
			} else {
				// 创建新记录
				if err := auth.CheckAuthorization(ctx, "CreateUserRole"); err != nil {
					return item, errors.New("CreateUserRole " + err.Error())
				}

				v.ID = uuid.Must(uuid.NewV4()).String()
				v.CreatedAt = timestampMillis
				v.CreatedBy = principalID

				// 保存新记录
				if err := tx.Omit(clause.Associations).Table(TableName("user_roles", ctx)).Create(v).Error; err != nil {
					return item, err
				}

				newUserRoles = append(newUserRoles, v)
			}
		}

		allItems := append(updateUserRoles, newUserRoles...)

		// ManyToMany: 使用 Replace 更新中间表
		if err := tx.Model(item).Association("UserRoles").Replace(allItems); err != nil {
			return item, err
		}

		event.AddNewValue("userRoles", allItems)
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

	if !utils.IsNil(input["profile"]) && !utils.IsNil(input["profileId"]) {
		return nil, fmt.Errorf("profileId and profile cannot coexist")
	}

	if !utils.IsNil(input["tasks"]) && !utils.IsNil(input["tasksIds"]) {
		return nil, fmt.Errorf("tasksIds and tasks cannot coexist")
	}

	if !utils.IsNil(input["userRoles"]) && !utils.IsNil(input["userRolesIds"]) {
		return nil, fmt.Errorf("userRolesIds and userRoles cannot coexist")
	}

	// 获取现有实体
	if err = GetItem(ctx, tx, TableName("users", ctx), item, &id); err != nil {
		return nil, err
	}

	// 更新 UpdatedBy
	if item.UpdatedBy != nil && principalID != nil && *item.UpdatedBy != *principalID {
		newItem.UpdatedBy = principalID
	}

	// 字段变更追踪
	changedFields := []string{}

	// ========== 处理 ManyToOne/OneToOne 关系 ==========

	// ---------- OneToOne: profile ----------
	if _, ok := input["profile"]; ok && !utils.IsNil(input["profile"]) {
		v := changes.Profile

		if !utils.IsEmpty(v.ID) {
			// 更新现有关联对象
			v.UpdatedAt = &timestampMillis
			v.UpdatedBy = principalID

			if err := auth.CheckAuthorization(ctx, "UpdateProfile"); err != nil {
				return item, errors.New("UpdateProfile " + err.Error())
			}
			if err := auth.CheckAuthorization(ctx, "Profile"); err != nil {
				return item, errors.New("Profile Detail " + err.Error())
			}

			profileInput := utils.StructToMap(*v)
			if _, err := r.Handlers.UpdateProfile(ctx, r, v.ID, profileInput); err != nil {
				return item, errors.New("Profile ID " + v.ID + " " + err.Error())
			}

			// 更新外键

			item.ProfileID = &v.ID
			newItem.ProfileID = &v.ID

			changedFields = append(changedFields, "user_id")
			isChange = true
		} else {
			// 创建新关联对象
			if err := auth.CheckAuthorization(ctx, "CreateProfile"); err != nil {
				return item, errors.New("CreateProfile " + err.Error())
			}

			v.ID = uuid.Must(uuid.NewV4()).String()
			v.CreatedAt = timestampMillis
			v.CreatedBy = principalID

			// 保存新关联对象
			if err := tx.Omit(clause.Associations).Table(TableName("profiles", ctx)).Create(v).Error; err != nil {
				return item, err
			}

			// 更新外键

			item.ProfileID = &v.ID
			newItem.ProfileID = &v.ID

			changedFields = append(changedFields, "user_id")
			isChange = true
		}
	}

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

	if _, ok := input["profileId"]; ok {
		// 只要 input 中包含该字段，且值发生了变化（包括变为 null），就进行更新
		if (item.ProfileID != changes.ProfileID) && (item.ProfileID == nil || changes.ProfileID == nil || *item.ProfileID != *changes.ProfileID) {

			if !utils.IsNil(input["profileId"]) {
				if err := tx.Select("id").Where("id = ?", input["profileId"]).First(&Profile{}).Error; err != nil {
					return nil, fmt.Errorf("profileId " + err.Error())
				}
			}

			event.AddOldValue("profileId", item.ProfileID)
			event.AddNewValue("profileId", changes.ProfileID)

			item.ProfileID = changes.ProfileID
			newItem.ProfileID = changes.ProfileID
			changedFields = append(changedFields, "profile_id")
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
		// 如果有更新 UpdatedBy，也需要添加到 Select 中
		if newItem.UpdatedBy != nil {
			changedFields = append(changedFields, "updated_by")
		}

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
				return item, errors.New("Task Detail " + err.Error())
			}
			if err := tx.Find(&items, "id IN (?)", itemIds).Error; err != nil {
				return item, err
			}
			for _, v := range items {
				findIds = append(findIds, v.ID)
			}

			differenceIds := utils.Difference(itemIds, findIds)
			if len(differenceIds) > 0 {
				return item, fmt.Errorf("tasksIds " + strings.Join(differenceIds, ",") + " not found")
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

		for index, v := range changes.Tasks {
			weight := int64(index + 1)
			v.Weight = &weight

			if !utils.IsEmpty(v.ID) {
				// 更新现有记录
				v.UpdatedAt = &timestampMillis
				v.UpdatedBy = principalID

				if err := auth.CheckAuthorization(ctx, "UpdateTask"); err != nil {
					return item, errors.New("UpdateTask " + err.Error())
				}
				if err := auth.CheckAuthorization(ctx, "Task"); err != nil {
					return item, errors.New("Task Detail " + err.Error())
				}

				tasksInput := utils.StructToMap(*v)
				if _, err := r.Handlers.UpdateTask(ctx, r, tasksInput["id"].(string), tasksInput); err != nil {
					return item, errors.New("Task ID " + v.ID + " " + err.Error())
				}

				if err := tx.Model(v).Update("user_id", item.ID).Error; err != nil {
					return item, err
				}

				updateTasks = append(updateTasks, v)
			} else {
				// 创建新记录
				if err := auth.CheckAuthorization(ctx, "CreateTask"); err != nil {
					return item, errors.New("CreateTask " + err.Error())
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

	// ---------- ToMany: userRoles ----------

	// 方式1：通过 IDs 关联
	if ids, ok := input["userRolesIds"]; ok && !utils.IsNil(input["userRolesIds"]) {
		items := []*UserRole{}
		itemIds := []string{}
		findIds := []string{}

		for _, v := range ids.([]interface{}) {
			itemIds = append(itemIds, v.(string))
		}

		if len(itemIds) > 0 {
			if err := auth.CheckAuthorization(ctx, "UserRole"); err != nil {
				return item, errors.New("UserRole Detail " + err.Error())
			}
			if err := tx.Find(&items, "id IN (?)", itemIds).Error; err != nil {
				return item, err
			}
			for _, v := range items {
				findIds = append(findIds, v.ID)
			}

			differenceIds := utils.Difference(itemIds, findIds)
			if len(differenceIds) > 0 {
				return item, fmt.Errorf("userRolesIds " + strings.Join(differenceIds, ",") + " not found")
			}

			if err := tx.Model(item).Association("UserRoles").Replace(items); err != nil {
				return item, err
			}

		} else {
			// 清空关联

			if err := tx.Model(item).Association("UserRoles").Clear(); err != nil {
				return item, err
			}

		}
		event.AddNewValue("userRoles", items)
	}

	// 方式2：通过嵌套对象创建/更新
	if _, ok := input["userRoles"]; ok && !utils.IsNil(input["userRoles"]) {
		newUserRoles := []*UserRole{}
		updateUserRoles := []*UserRole{}

		for index, v := range changes.UserRoles {
			weight := int64(index + 1)
			v.Weight = &weight

			if !utils.IsEmpty(v.ID) {
				// 更新现有记录
				v.UpdatedAt = &timestampMillis
				v.UpdatedBy = principalID

				if err := auth.CheckAuthorization(ctx, "UpdateUserRole"); err != nil {
					return item, errors.New("UpdateUserRole " + err.Error())
				}
				if err := auth.CheckAuthorization(ctx, "UserRole"); err != nil {
					return item, errors.New("UserRole Detail " + err.Error())
				}

				userRolesInput := utils.StructToMap(*v)
				if _, err := r.Handlers.UpdateUserRole(ctx, r, userRolesInput["id"].(string), userRolesInput); err != nil {
					return item, errors.New("UserRole ID " + v.ID + " " + err.Error())
				}

				updateUserRoles = append(updateUserRoles, v)
			} else {
				// 创建新记录
				if err := auth.CheckAuthorization(ctx, "CreateUserRole"); err != nil {
					return item, errors.New("CreateUserRole " + err.Error())
				}

				v.ID = uuid.Must(uuid.NewV4()).String()
				v.CreatedAt = timestampMillis
				v.CreatedBy = principalID

				if err := tx.Omit(clause.Associations).Table(TableName("user_roles", ctx)).Create(v).Error; err != nil {
					return item, err
				}

				newUserRoles = append(newUserRoles, v)
			}
		}

		allItems := append(updateUserRoles, newUserRoles...)

		if err := tx.Model(item).Association("UserRoles").Replace(allItems); err != nil {
			return item, err
		}

		event.AddNewValue("userRoles", allItems)
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
// Profile - Create
// ============================================================

// CreateProfile 创建 Profile 实体的解析器入口
func (r *GeneratedMutationResolver) CreateProfile(ctx context.Context, input map[string]interface{}) (item *Profile, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreateProfile(ctx, r.GeneratedResolver, input)
	if err != nil {
		RollbackMutationContext(ctx, r.GeneratedResolver)
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// CreateProfileHandler 处理 Profile 创建逻辑
func CreateProfileHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Profile, err error) {
	item = &Profile{}
	now := time.Now()
	timestampMillis := now.UnixNano() / 1e6
	principalID := GetPrincipalIDFromContext(ctx)
	tx := GetTransaction(ctx)

	// 创建事件记录
	event := NewEvent(EventMetadata{
		Type:        EventTypeCreated,
		Entity:      "Profile",
		EntityID:    item.ID,
		Date:        timestampMillis,
		PrincipalID: principalID,
	})

	// 解析输入变更
	var changes ProfileChanges
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

	// ========== 处理 ManyToOne/OneToOne 关系（当前实体持有外键） ==========

	// ---------- OneToOne: user (当前表持有外键 profile.user_id) ----------
	if _, ok := input["user"]; ok && !utils.IsNil(input["user"]) {
		v := changes.User

		if !utils.IsEmpty(v.ID) {
			// 更新现有关联对象
			v.UpdatedAt = &timestampMillis
			v.UpdatedBy = principalID

			if err := auth.CheckAuthorization(ctx, "UpdateUser"); err != nil {
				return item, errors.New("UpdateUser " + err.Error())
			}
			if err := auth.CheckAuthorization(ctx, "User"); err != nil {
				return item, errors.New("User Detail " + err.Error())
			}

			userInput := utils.StructToMap(*v)
			if _, err := r.Handlers.UpdateUser(ctx, r, v.ID, userInput); err != nil {
				return item, errors.New("User ID " + v.ID + " " + err.Error())
			}

			// 设置外键

			item.UserID = v.ID

		} else {
			// 创建新关联对象
			if err := auth.CheckAuthorization(ctx, "CreateUser"); err != nil {
				return item, errors.New("CreateUser " + err.Error())
			}

			v.ID = uuid.Must(uuid.NewV4()).String()
			v.CreatedAt = timestampMillis
			v.CreatedBy = principalID

			// 先保存关联对象
			if err := tx.Omit(clause.Associations).Table(TableName("users", ctx)).Create(v).Error; err != nil {
				return item, err
			}

			// 设置外键

			item.UserID = v.ID

		}

		item.User = v
		event.AddNewValue("user", item.User)
		event.AddNewValue("userId", item.UserID)
	}

	// ========== 处理普通字段 ==========

	if _, ok := input["avatar"]; ok && changes.Avatar != nil {

		if (item.Avatar != changes.Avatar) || (*item.Avatar != *changes.Avatar) {

			item.Avatar = changes.Avatar

			event.AddNewValue("avatar", changes.Avatar)
		}
	}

	if _, ok := input["bio"]; ok && changes.Bio != nil {

		if (item.Bio != changes.Bio) || (*item.Bio != *changes.Bio) {

			item.Bio = changes.Bio

			event.AddNewValue("bio", changes.Bio)
		}
	}

	if _, ok := input["birthday"]; ok && changes.Birthday != nil {

		if (item.Birthday != changes.Birthday) || (*item.Birthday != *changes.Birthday) {

			item.Birthday = changes.Birthday

			event.AddNewValue("birthday", changes.Birthday)
		}
	}

	if _, ok := input["address"]; ok && changes.Address != nil {

		if (item.Address != changes.Address) || (*item.Address != *changes.Address) {

			item.Address = changes.Address

			event.AddNewValue("address", changes.Address)
		}
	}

	if _, ok := input["userId"]; ok && !utils.IsEmpty(input["userId"]) {

		if item.UserID != changes.UserID {

			if !utils.IsNil(input["userId"]) {
				if err := tx.Select("id").Where("id = ?", input["userId"]).First(&User{}).Error; err != nil {
					return nil, fmt.Errorf("userId " + err.Error())
				}
			}

			item.UserID = changes.UserID

			event.AddNewValue("userId", changes.UserID)
		}
	}

	if _, ok := input["isDelete"]; ok && changes.IsDelete != nil {

		if (item.IsDelete != changes.IsDelete) || (*item.IsDelete != *changes.IsDelete) {

			item.IsDelete = changes.IsDelete

			event.AddNewValue("isDelete", changes.IsDelete)
		}
	}

	if _, ok := input["weight"]; ok && changes.Weight != nil {

		if (item.Weight != changes.Weight) || (*item.Weight != *changes.Weight) {

			item.Weight = changes.Weight

			event.AddNewValue("weight", changes.Weight)
		}
	}

	if _, ok := input["state"]; ok && changes.State != nil {

		if (item.State != changes.State) || (*item.State != *changes.State) {

			item.State = changes.State

			event.AddNewValue("state", changes.State)
		}
	}

	// ========== 保存主实体 ==========
	if err := tx.Omit(clause.Associations).Table(TableName("profiles", ctx)).Create(item).Error; err != nil {
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
// Profile - Update
// ============================================================

// UpdateProfile 更新 Profile 实体的解析器入口
func (r *GeneratedMutationResolver) UpdateProfile(ctx context.Context, id string, input map[string]interface{}) (item *Profile, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdateProfile(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		RollbackMutationContext(ctx, r.GeneratedResolver)
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// UpdateProfileHandler 处理 Profile 更新逻辑
func UpdateProfileHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Profile, err error) {
	item = &Profile{}
	newItem := &Profile{}
	isChange := false

	now := time.Now()
	timestampMillis := now.UnixNano() / 1e6
	principalID := GetPrincipalIDFromContext(ctx)
	tx := GetTransaction(ctx)

	// 创建事件记录
	event := NewEvent(EventMetadata{
		Type:        EventTypeUpdated,
		Entity:      "Profile",
		EntityID:    id,
		Date:        timestampMillis,
		PrincipalID: principalID,
	})

	// 解析输入变更
	var changes ProfileChanges
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
	if err = GetItem(ctx, tx, TableName("profiles", ctx), item, &id); err != nil {
		return nil, err
	}

	// 更新 UpdatedBy
	if item.UpdatedBy != nil && principalID != nil && *item.UpdatedBy != *principalID {
		newItem.UpdatedBy = principalID
	}

	// 字段变更追踪
	changedFields := []string{}

	// ========== 处理 ManyToOne/OneToOne 关系 ==========

	// ---------- OneToOne: user ----------
	if _, ok := input["user"]; ok && !utils.IsNil(input["user"]) {
		v := changes.User

		if !utils.IsEmpty(v.ID) {
			// 更新现有关联对象
			v.UpdatedAt = &timestampMillis
			v.UpdatedBy = principalID

			if err := auth.CheckAuthorization(ctx, "UpdateUser"); err != nil {
				return item, errors.New("UpdateUser " + err.Error())
			}
			if err := auth.CheckAuthorization(ctx, "User"); err != nil {
				return item, errors.New("User Detail " + err.Error())
			}

			userInput := utils.StructToMap(*v)
			if _, err := r.Handlers.UpdateUser(ctx, r, v.ID, userInput); err != nil {
				return item, errors.New("User ID " + v.ID + " " + err.Error())
			}

			// 更新外键

			item.UserID = v.ID
			newItem.UserID = v.ID

			changedFields = append(changedFields, "profile_id")
			isChange = true
		} else {
			// 创建新关联对象
			if err := auth.CheckAuthorization(ctx, "CreateUser"); err != nil {
				return item, errors.New("CreateUser " + err.Error())
			}

			v.ID = uuid.Must(uuid.NewV4()).String()
			v.CreatedAt = timestampMillis
			v.CreatedBy = principalID

			// 保存新关联对象
			if err := tx.Omit(clause.Associations).Table(TableName("users", ctx)).Create(v).Error; err != nil {
				return item, err
			}

			// 更新外键

			item.UserID = v.ID
			newItem.UserID = v.ID

			changedFields = append(changedFields, "profile_id")
			isChange = true
		}
	}

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

	if _, ok := input["avatar"]; ok {
		// 只要 input 中包含该字段，且值发生了变化（包括变为 null），就进行更新
		if (item.Avatar != changes.Avatar) && (item.Avatar == nil || changes.Avatar == nil || *item.Avatar != *changes.Avatar) {

			event.AddOldValue("avatar", item.Avatar)
			event.AddNewValue("avatar", changes.Avatar)

			item.Avatar = changes.Avatar
			newItem.Avatar = changes.Avatar
			changedFields = append(changedFields, "avatar")
			isChange = true
		}
	}

	if _, ok := input["bio"]; ok {
		// 只要 input 中包含该字段，且值发生了变化（包括变为 null），就进行更新
		if (item.Bio != changes.Bio) && (item.Bio == nil || changes.Bio == nil || *item.Bio != *changes.Bio) {

			event.AddOldValue("bio", item.Bio)
			event.AddNewValue("bio", changes.Bio)

			item.Bio = changes.Bio
			newItem.Bio = changes.Bio
			changedFields = append(changedFields, "bio")
			isChange = true
		}
	}

	if _, ok := input["birthday"]; ok {
		// 只要 input 中包含该字段，且值发生了变化（包括变为 null），就进行更新
		if (item.Birthday != changes.Birthday) && (item.Birthday == nil || changes.Birthday == nil || *item.Birthday != *changes.Birthday) {

			event.AddOldValue("birthday", item.Birthday)
			event.AddNewValue("birthday", changes.Birthday)

			item.Birthday = changes.Birthday
			newItem.Birthday = changes.Birthday
			changedFields = append(changedFields, "birthday")
			isChange = true
		}
	}

	if _, ok := input["address"]; ok {
		// 只要 input 中包含该字段，且值发生了变化（包括变为 null），就进行更新
		if (item.Address != changes.Address) && (item.Address == nil || changes.Address == nil || *item.Address != *changes.Address) {

			event.AddOldValue("address", item.Address)
			event.AddNewValue("address", changes.Address)

			item.Address = changes.Address
			newItem.Address = changes.Address
			changedFields = append(changedFields, "address")
			isChange = true
		}
	}

	if _, ok := input["userId"]; ok {
		// 只要 input 中包含该字段，且值发生了变化（包括变为 null），就进行更新
		if item.UserID != changes.UserID {

			if !utils.IsNil(input["userId"]) {
				if err := tx.Select("id").Where("id = ?", input["userId"]).First(&User{}).Error; err != nil {
					return nil, fmt.Errorf("userId " + err.Error())
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
		// 如果有更新 UpdatedBy，也需要添加到 Select 中
		if newItem.UpdatedBy != nil {
			changedFields = append(changedFields, "updated_by")
		}

		if err := tx.Table(TableName("profiles", ctx)).Where("id = ?", id).Select(changedFields).Updates(newItem).Error; err != nil {
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
// Profile - Delete
// ============================================================

// DeleteProfileFunc 执行删除或恢复操作
func DeleteProfileFunc(ctx context.Context, r *GeneratedResolver, id string, operationType string, unscoped *bool) (err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item := &Profile{}
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
	if err = tx.Unscoped().Table(TableName("profiles", ctx)).Where("is_delete = ? and id = ?", status, id).First(item).Error; err != nil {
		return err
	}

	deletedAt := now.UnixNano() / 1e6

	// 创建事件记录
	event := NewEvent(EventMetadata{
		Type:        EventTypeDeleted,
		Entity:      "Profile",
		EntityID:    id,
		Date:        deletedAt,
		PrincipalID: principalID,
	})

	// 执行删除或恢复
	if operationType == "recovery" {
		if err := tx.Unscoped().Table(TableName("profiles", ctx)).Model(&item).Updates(map[string]interface{}{
			"IsDelete":  1,
			"DeletedAt": nil,
			"DeletedBy": nil,
		}).Error; err != nil {
			return err
		}
	} else {
		if unscoped != nil && *unscoped {
			// 物理删除
			if err := tx.Unscoped().Table(TableName("profiles", ctx)).Model(&item).Delete(item).Error; err != nil {
				return err
			}
		} else {
			// 软删除
			if err := tx.Model(&item).Table(TableName("profiles", ctx)).Updates(Profile{
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

// DeleteProfiles 批量删除 Profile 实体
func (r *GeneratedMutationResolver) DeleteProfiles(ctx context.Context, id []string, unscoped *bool) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteProfiles(ctx, r.GeneratedResolver, id, unscoped)
	if err != nil {
		RollbackMutationContext(ctx, r.GeneratedResolver)
		return done, err
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

// DeleteProfilesHandler 处理批量删除逻辑
func DeleteProfilesHandler(ctx context.Context, r *GeneratedResolver, id []string, unscoped *bool) (bool, error) {
	for _, itemID := range id {
		if err := DeleteProfileFunc(ctx, r, itemID, "delete", unscoped); err != nil {
			return false, err
		}
	}
	return true, nil
}

// ============================================================
// Profile - Recovery
// ============================================================

// RecoveryProfiles 批量恢复 Profile 实体
func (r *GeneratedMutationResolver) RecoveryProfiles(ctx context.Context, id []string) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.RecoveryProfiles(ctx, r.GeneratedResolver, id)
	if err != nil {
		return done, err
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

// RecoveryProfilesHandler 处理批量恢复逻辑
func RecoveryProfilesHandler(ctx context.Context, r *GeneratedResolver, id []string) (bool, error) {
	unscoped := false
	for _, itemID := range id {
		if err := DeleteProfileFunc(ctx, r, itemID, "recovery", &unscoped); err != nil {
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

	// ToMany: tags - 不能同时传入 IDs 和嵌套对象
	if !utils.IsNil(input["tags"]) && !utils.IsNil(input["tagsIds"]) {
		return nil, fmt.Errorf("tagsIds and tags cannot coexist")
	}

	// ========== 处理 ManyToOne/OneToOne 关系（当前实体持有外键） ==========

	// ========== 处理普通字段 ==========

	if _, ok := input["title"]; ok && !utils.IsEmpty(input["title"]) {

		if item.Title != changes.Title {

			item.Title = changes.Title

			event.AddNewValue("title", changes.Title)
		}
	}

	if _, ok := input["description"]; ok && changes.Description != nil {

		if (item.Description != changes.Description) || (*item.Description != *changes.Description) {

			item.Description = changes.Description

			event.AddNewValue("description", changes.Description)
		}
	}

	if _, ok := input["completed"]; ok && changes.Completed != nil {

		if (item.Completed != changes.Completed) || (*item.Completed != *changes.Completed) {

			item.Completed = changes.Completed

			event.AddNewValue("completed", changes.Completed)
		}
	}

	if _, ok := input["dueDate"]; ok && changes.DueDate != nil {

		if (item.DueDate != changes.DueDate) || (*item.DueDate != *changes.DueDate) {

			item.DueDate = changes.DueDate

			event.AddNewValue("dueDate", changes.DueDate)
		}
	}

	if _, ok := input["priority"]; ok && changes.Priority != nil {

		if (item.Priority != changes.Priority) || (*item.Priority != *changes.Priority) {

			item.Priority = changes.Priority

			event.AddNewValue("priority", changes.Priority)
		}
	}

	if _, ok := input["userId"]; ok && changes.UserID != nil {

		if (item.UserID != changes.UserID) || (*item.UserID != *changes.UserID) {

			if !utils.IsNil(input["userId"]) {
				if err := tx.Select("id").Where("id = ?", input["userId"]).First(&User{}).Error; err != nil {
					return nil, fmt.Errorf("userId " + err.Error())
				}
			}

			item.UserID = changes.UserID

			event.AddNewValue("userId", changes.UserID)
		}
	}

	if _, ok := input["isDelete"]; ok && changes.IsDelete != nil {

		if (item.IsDelete != changes.IsDelete) || (*item.IsDelete != *changes.IsDelete) {

			item.IsDelete = changes.IsDelete

			event.AddNewValue("isDelete", changes.IsDelete)
		}
	}

	if _, ok := input["weight"]; ok && changes.Weight != nil {

		if (item.Weight != changes.Weight) || (*item.Weight != *changes.Weight) {

			item.Weight = changes.Weight

			event.AddNewValue("weight", changes.Weight)
		}
	}

	if _, ok := input["state"]; ok && changes.State != nil {

		if (item.State != changes.State) || (*item.State != *changes.State) {

			item.State = changes.State

			event.AddNewValue("state", changes.State)
		}
	}

	// ========== 保存主实体 ==========
	if err := tx.Omit(clause.Associations).Table(TableName("tasks", ctx)).Create(item).Error; err != nil {
		return item, err
	}

	// ========== 处理 OneToMany/ManyToMany 关系（关联表持有外键或中间表） ==========

	// ---------- ToMany: tags (ManyToMany 中间表) ----------

	// 方式1：通过 IDs 关联现有记录
	if ids, ok := input["tagsIds"]; ok && !utils.IsNil(input["tagsIds"]) {
		items := []*Tag{}
		itemIds := []string{}
		findIds := []string{}

		for _, v := range ids.([]interface{}) {
			itemIds = append(itemIds, v.(string))
		}

		if len(itemIds) > 0 {
			// 权限检查
			if err := auth.CheckAuthorization(ctx, "Tag"); err != nil {
				return item, errors.New("Tag Detail " + err.Error())
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
				return item, fmt.Errorf("tagsIds " + strings.Join(differenceIds, ",") + " not found")
			}

			// ManyToMany: 使用 Replace 更新中间表
			if err := tx.Model(item).Association("Tags").Replace(items); err != nil {
				return item, err
			}

		}
		event.AddNewValue("tags", items)
	}

	// 方式2：通过嵌套对象创建/更新
	if _, ok := input["tags"]; ok && !utils.IsNil(input["tags"]) {
		newTags := []*Tag{}
		updateTags := []*Tag{}

		for index, v := range changes.Tags {
			weight := int64(index + 1)
			v.Weight = &weight

			if !utils.IsEmpty(v.ID) {
				// 更新现有记录
				v.UpdatedAt = &timestampMillis
				v.UpdatedBy = principalID

				if err := auth.CheckAuthorization(ctx, "UpdateTag"); err != nil {
					return item, errors.New("UpdateTag " + err.Error())
				}
				if err := auth.CheckAuthorization(ctx, "Tag"); err != nil {
					return item, errors.New("Tag Detail " + err.Error())
				}

				tagsInput := utils.StructToMap(*v)
				if _, err := r.Handlers.UpdateTag(ctx, r, tagsInput["id"].(string), tagsInput); err != nil {
					return item, errors.New("Tag ID " + v.ID + " " + err.Error())
				}

				updateTags = append(updateTags, v)
			} else {
				// 创建新记录
				if err := auth.CheckAuthorization(ctx, "CreateTag"); err != nil {
					return item, errors.New("CreateTag " + err.Error())
				}

				v.ID = uuid.Must(uuid.NewV4()).String()
				v.CreatedAt = timestampMillis
				v.CreatedBy = principalID

				// 保存新记录
				if err := tx.Omit(clause.Associations).Table(TableName("tags", ctx)).Create(v).Error; err != nil {
					return item, err
				}

				newTags = append(newTags, v)
			}
		}

		allItems := append(updateTags, newTags...)

		// ManyToMany: 使用 Replace 更新中间表
		if err := tx.Model(item).Association("Tags").Replace(allItems); err != nil {
			return item, err
		}

		event.AddNewValue("tags", allItems)
	}

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

	if !utils.IsNil(input["tags"]) && !utils.IsNil(input["tagsIds"]) {
		return nil, fmt.Errorf("tagsIds and tags cannot coexist")
	}

	// 获取现有实体
	if err = GetItem(ctx, tx, TableName("tasks", ctx), item, &id); err != nil {
		return nil, err
	}

	// 更新 UpdatedBy
	if item.UpdatedBy != nil && principalID != nil && *item.UpdatedBy != *principalID {
		newItem.UpdatedBy = principalID
	}

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
		if item.Title != changes.Title {

			event.AddOldValue("title", item.Title)
			event.AddNewValue("title", changes.Title)

			item.Title = changes.Title
			newItem.Title = changes.Title
			changedFields = append(changedFields, "title")
			isChange = true
		}
	}

	if _, ok := input["description"]; ok {
		// 只要 input 中包含该字段，且值发生了变化（包括变为 null），就进行更新
		if (item.Description != changes.Description) && (item.Description == nil || changes.Description == nil || *item.Description != *changes.Description) {

			event.AddOldValue("description", item.Description)
			event.AddNewValue("description", changes.Description)

			item.Description = changes.Description
			newItem.Description = changes.Description
			changedFields = append(changedFields, "description")
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

	if _, ok := input["priority"]; ok {
		// 只要 input 中包含该字段，且值发生了变化（包括变为 null），就进行更新
		if (item.Priority != changes.Priority) && (item.Priority == nil || changes.Priority == nil || *item.Priority != *changes.Priority) {

			event.AddOldValue("priority", item.Priority)
			event.AddNewValue("priority", changes.Priority)

			item.Priority = changes.Priority
			newItem.Priority = changes.Priority
			changedFields = append(changedFields, "priority")
			isChange = true
		}
	}

	if _, ok := input["userId"]; ok {
		// 只要 input 中包含该字段，且值发生了变化（包括变为 null），就进行更新
		if (item.UserID != changes.UserID) && (item.UserID == nil || changes.UserID == nil || *item.UserID != *changes.UserID) {

			if !utils.IsNil(input["userId"]) {
				if err := tx.Select("id").Where("id = ?", input["userId"]).First(&User{}).Error; err != nil {
					return nil, fmt.Errorf("userId " + err.Error())
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
		// 如果有更新 UpdatedBy，也需要添加到 Select 中
		if newItem.UpdatedBy != nil {
			changedFields = append(changedFields, "updated_by")
		}

		if err := tx.Table(TableName("tasks", ctx)).Where("id = ?", id).Select(changedFields).Updates(newItem).Error; err != nil {
			return item, err
		}
	}

	// ========== 处理 OneToMany/ManyToMany 关系 ==========

	// ---------- ToMany: tags ----------

	// 方式1：通过 IDs 关联
	if ids, ok := input["tagsIds"]; ok && !utils.IsNil(input["tagsIds"]) {
		items := []*Tag{}
		itemIds := []string{}
		findIds := []string{}

		for _, v := range ids.([]interface{}) {
			itemIds = append(itemIds, v.(string))
		}

		if len(itemIds) > 0 {
			if err := auth.CheckAuthorization(ctx, "Tag"); err != nil {
				return item, errors.New("Tag Detail " + err.Error())
			}
			if err := tx.Find(&items, "id IN (?)", itemIds).Error; err != nil {
				return item, err
			}
			for _, v := range items {
				findIds = append(findIds, v.ID)
			}

			differenceIds := utils.Difference(itemIds, findIds)
			if len(differenceIds) > 0 {
				return item, fmt.Errorf("tagsIds " + strings.Join(differenceIds, ",") + " not found")
			}

			if err := tx.Model(item).Association("Tags").Replace(items); err != nil {
				return item, err
			}

		} else {
			// 清空关联

			if err := tx.Model(item).Association("Tags").Clear(); err != nil {
				return item, err
			}

		}
		event.AddNewValue("tags", items)
	}

	// 方式2：通过嵌套对象创建/更新
	if _, ok := input["tags"]; ok && !utils.IsNil(input["tags"]) {
		newTags := []*Tag{}
		updateTags := []*Tag{}

		for index, v := range changes.Tags {
			weight := int64(index + 1)
			v.Weight = &weight

			if !utils.IsEmpty(v.ID) {
				// 更新现有记录
				v.UpdatedAt = &timestampMillis
				v.UpdatedBy = principalID

				if err := auth.CheckAuthorization(ctx, "UpdateTag"); err != nil {
					return item, errors.New("UpdateTag " + err.Error())
				}
				if err := auth.CheckAuthorization(ctx, "Tag"); err != nil {
					return item, errors.New("Tag Detail " + err.Error())
				}

				tagsInput := utils.StructToMap(*v)
				if _, err := r.Handlers.UpdateTag(ctx, r, tagsInput["id"].(string), tagsInput); err != nil {
					return item, errors.New("Tag ID " + v.ID + " " + err.Error())
				}

				updateTags = append(updateTags, v)
			} else {
				// 创建新记录
				if err := auth.CheckAuthorization(ctx, "CreateTag"); err != nil {
					return item, errors.New("CreateTag " + err.Error())
				}

				v.ID = uuid.Must(uuid.NewV4()).String()
				v.CreatedAt = timestampMillis
				v.CreatedBy = principalID

				if err := tx.Omit(clause.Associations).Table(TableName("tags", ctx)).Create(v).Error; err != nil {
					return item, err
				}

				newTags = append(newTags, v)
			}
		}

		allItems := append(updateTags, newTags...)

		if err := tx.Model(item).Association("Tags").Replace(allItems); err != nil {
			return item, err
		}

		event.AddNewValue("tags", allItems)
	}

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

// ============================================================
// UserRole - Create
// ============================================================

// CreateUserRole 创建 UserRole 实体的解析器入口
func (r *GeneratedMutationResolver) CreateUserRole(ctx context.Context, input map[string]interface{}) (item *UserRole, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreateUserRole(ctx, r.GeneratedResolver, input)
	if err != nil {
		RollbackMutationContext(ctx, r.GeneratedResolver)
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// CreateUserRoleHandler 处理 UserRole 创建逻辑
func CreateUserRoleHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *UserRole, err error) {
	item = &UserRole{}
	now := time.Now()
	timestampMillis := now.UnixNano() / 1e6
	principalID := GetPrincipalIDFromContext(ctx)
	tx := GetTransaction(ctx)

	// 创建事件记录
	event := NewEvent(EventMetadata{
		Type:        EventTypeCreated,
		Entity:      "UserRole",
		EntityID:    item.ID,
		Date:        timestampMillis,
		PrincipalID: principalID,
	})

	// 解析输入变更
	var changes UserRoleChanges
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

	// ToMany: users - 不能同时传入 IDs 和嵌套对象
	if !utils.IsNil(input["users"]) && !utils.IsNil(input["usersIds"]) {
		return nil, fmt.Errorf("usersIds and users cannot coexist")
	}

	// ========== 处理 ManyToOne/OneToOne 关系（当前实体持有外键） ==========

	// ========== 处理普通字段 ==========

	if _, ok := input["name"]; ok && !utils.IsEmpty(input["name"]) {

		if item.Name != changes.Name {

			item.Name = changes.Name

			event.AddNewValue("name", changes.Name)
		}
	}

	if _, ok := input["description"]; ok && changes.Description != nil {

		if (item.Description != changes.Description) || (*item.Description != *changes.Description) {

			item.Description = changes.Description

			event.AddNewValue("description", changes.Description)
		}
	}

	if _, ok := input["permissions"]; ok && changes.Permissions != nil {

		if (item.Permissions != changes.Permissions) || (*item.Permissions != *changes.Permissions) {

			item.Permissions = changes.Permissions

			event.AddNewValue("permissions", changes.Permissions)
		}
	}

	if _, ok := input["isDelete"]; ok && changes.IsDelete != nil {

		if (item.IsDelete != changes.IsDelete) || (*item.IsDelete != *changes.IsDelete) {

			item.IsDelete = changes.IsDelete

			event.AddNewValue("isDelete", changes.IsDelete)
		}
	}

	if _, ok := input["weight"]; ok && changes.Weight != nil {

		if (item.Weight != changes.Weight) || (*item.Weight != *changes.Weight) {

			item.Weight = changes.Weight

			event.AddNewValue("weight", changes.Weight)
		}
	}

	if _, ok := input["state"]; ok && changes.State != nil {

		if (item.State != changes.State) || (*item.State != *changes.State) {

			item.State = changes.State

			event.AddNewValue("state", changes.State)
		}
	}

	// ========== 保存主实体 ==========
	if err := tx.Omit(clause.Associations).Table(TableName("user_roles", ctx)).Create(item).Error; err != nil {
		return item, err
	}

	// ========== 处理 OneToMany/ManyToMany 关系（关联表持有外键或中间表） ==========

	// ---------- ToMany: users (ManyToMany 中间表) ----------

	// 方式1：通过 IDs 关联现有记录
	if ids, ok := input["usersIds"]; ok && !utils.IsNil(input["usersIds"]) {
		items := []*User{}
		itemIds := []string{}
		findIds := []string{}

		for _, v := range ids.([]interface{}) {
			itemIds = append(itemIds, v.(string))
		}

		if len(itemIds) > 0 {
			// 权限检查
			if err := auth.CheckAuthorization(ctx, "User"); err != nil {
				return item, errors.New("User Detail " + err.Error())
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
				return item, fmt.Errorf("usersIds " + strings.Join(differenceIds, ",") + " not found")
			}

			// ManyToMany: 使用 Replace 更新中间表
			if err := tx.Model(item).Association("Users").Replace(items); err != nil {
				return item, err
			}

		}
		event.AddNewValue("users", items)
	}

	// 方式2：通过嵌套对象创建/更新
	if _, ok := input["users"]; ok && !utils.IsNil(input["users"]) {
		newUsers := []*User{}
		updateUsers := []*User{}

		for index, v := range changes.Users {
			weight := int64(index + 1)
			v.Weight = &weight

			if !utils.IsEmpty(v.ID) {
				// 更新现有记录
				v.UpdatedAt = &timestampMillis
				v.UpdatedBy = principalID

				if err := auth.CheckAuthorization(ctx, "UpdateUser"); err != nil {
					return item, errors.New("UpdateUser " + err.Error())
				}
				if err := auth.CheckAuthorization(ctx, "User"); err != nil {
					return item, errors.New("User Detail " + err.Error())
				}

				usersInput := utils.StructToMap(*v)
				if _, err := r.Handlers.UpdateUser(ctx, r, usersInput["id"].(string), usersInput); err != nil {
					return item, errors.New("User ID " + v.ID + " " + err.Error())
				}

				updateUsers = append(updateUsers, v)
			} else {
				// 创建新记录
				if err := auth.CheckAuthorization(ctx, "CreateUser"); err != nil {
					return item, errors.New("CreateUser " + err.Error())
				}

				v.ID = uuid.Must(uuid.NewV4()).String()
				v.CreatedAt = timestampMillis
				v.CreatedBy = principalID

				// 保存新记录
				if err := tx.Omit(clause.Associations).Table(TableName("users", ctx)).Create(v).Error; err != nil {
					return item, err
				}

				newUsers = append(newUsers, v)
			}
		}

		allItems := append(updateUsers, newUsers...)

		// ManyToMany: 使用 Replace 更新中间表
		if err := tx.Model(item).Association("Users").Replace(allItems); err != nil {
			return item, err
		}

		event.AddNewValue("users", allItems)
	}

	// 记录事件
	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}

// ============================================================
// UserRole - Update
// ============================================================

// UpdateUserRole 更新 UserRole 实体的解析器入口
func (r *GeneratedMutationResolver) UpdateUserRole(ctx context.Context, id string, input map[string]interface{}) (item *UserRole, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdateUserRole(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		RollbackMutationContext(ctx, r.GeneratedResolver)
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// UpdateUserRoleHandler 处理 UserRole 更新逻辑
func UpdateUserRoleHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *UserRole, err error) {
	item = &UserRole{}
	newItem := &UserRole{}
	isChange := false

	now := time.Now()
	timestampMillis := now.UnixNano() / 1e6
	principalID := GetPrincipalIDFromContext(ctx)
	tx := GetTransaction(ctx)

	// 创建事件记录
	event := NewEvent(EventMetadata{
		Type:        EventTypeUpdated,
		Entity:      "UserRole",
		EntityID:    id,
		Date:        timestampMillis,
		PrincipalID: principalID,
	})

	// 解析输入变更
	var changes UserRoleChanges
	if err = ApplyChanges(input, &changes); err != nil {
		return
	}

	// 验证必填字段
	if err = CheckStructFieldIsEmpty(item, input); err != nil {
		return nil, err
	}

	// ========== 验证关系字段冲突 ==========

	if !utils.IsNil(input["users"]) && !utils.IsNil(input["usersIds"]) {
		return nil, fmt.Errorf("usersIds and users cannot coexist")
	}

	// 获取现有实体
	if err = GetItem(ctx, tx, TableName("user_roles", ctx), item, &id); err != nil {
		return nil, err
	}

	// 更新 UpdatedBy
	if item.UpdatedBy != nil && principalID != nil && *item.UpdatedBy != *principalID {
		newItem.UpdatedBy = principalID
	}

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

	if _, ok := input["name"]; ok {
		// 只要 input 中包含该字段，且值发生了变化（包括变为 null），就进行更新
		if item.Name != changes.Name {

			event.AddOldValue("name", item.Name)
			event.AddNewValue("name", changes.Name)

			item.Name = changes.Name
			newItem.Name = changes.Name
			changedFields = append(changedFields, "name")
			isChange = true
		}
	}

	if _, ok := input["description"]; ok {
		// 只要 input 中包含该字段，且值发生了变化（包括变为 null），就进行更新
		if (item.Description != changes.Description) && (item.Description == nil || changes.Description == nil || *item.Description != *changes.Description) {

			event.AddOldValue("description", item.Description)
			event.AddNewValue("description", changes.Description)

			item.Description = changes.Description
			newItem.Description = changes.Description
			changedFields = append(changedFields, "description")
			isChange = true
		}
	}

	if _, ok := input["permissions"]; ok {
		// 只要 input 中包含该字段，且值发生了变化（包括变为 null），就进行更新
		if (item.Permissions != changes.Permissions) && (item.Permissions == nil || changes.Permissions == nil || *item.Permissions != *changes.Permissions) {

			event.AddOldValue("permissions", item.Permissions)
			event.AddNewValue("permissions", changes.Permissions)

			item.Permissions = changes.Permissions
			newItem.Permissions = changes.Permissions
			changedFields = append(changedFields, "permissions")
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
		// 如果有更新 UpdatedBy，也需要添加到 Select 中
		if newItem.UpdatedBy != nil {
			changedFields = append(changedFields, "updated_by")
		}

		if err := tx.Table(TableName("user_roles", ctx)).Where("id = ?", id).Select(changedFields).Updates(newItem).Error; err != nil {
			return item, err
		}
	}

	// ========== 处理 OneToMany/ManyToMany 关系 ==========

	// ---------- ToMany: users ----------

	// 方式1：通过 IDs 关联
	if ids, ok := input["usersIds"]; ok && !utils.IsNil(input["usersIds"]) {
		items := []*User{}
		itemIds := []string{}
		findIds := []string{}

		for _, v := range ids.([]interface{}) {
			itemIds = append(itemIds, v.(string))
		}

		if len(itemIds) > 0 {
			if err := auth.CheckAuthorization(ctx, "User"); err != nil {
				return item, errors.New("User Detail " + err.Error())
			}
			if err := tx.Find(&items, "id IN (?)", itemIds).Error; err != nil {
				return item, err
			}
			for _, v := range items {
				findIds = append(findIds, v.ID)
			}

			differenceIds := utils.Difference(itemIds, findIds)
			if len(differenceIds) > 0 {
				return item, fmt.Errorf("usersIds " + strings.Join(differenceIds, ",") + " not found")
			}

			if err := tx.Model(item).Association("Users").Replace(items); err != nil {
				return item, err
			}

		} else {
			// 清空关联

			if err := tx.Model(item).Association("Users").Clear(); err != nil {
				return item, err
			}

		}
		event.AddNewValue("users", items)
	}

	// 方式2：通过嵌套对象创建/更新
	if _, ok := input["users"]; ok && !utils.IsNil(input["users"]) {
		newUsers := []*User{}
		updateUsers := []*User{}

		for index, v := range changes.Users {
			weight := int64(index + 1)
			v.Weight = &weight

			if !utils.IsEmpty(v.ID) {
				// 更新现有记录
				v.UpdatedAt = &timestampMillis
				v.UpdatedBy = principalID

				if err := auth.CheckAuthorization(ctx, "UpdateUser"); err != nil {
					return item, errors.New("UpdateUser " + err.Error())
				}
				if err := auth.CheckAuthorization(ctx, "User"); err != nil {
					return item, errors.New("User Detail " + err.Error())
				}

				usersInput := utils.StructToMap(*v)
				if _, err := r.Handlers.UpdateUser(ctx, r, usersInput["id"].(string), usersInput); err != nil {
					return item, errors.New("User ID " + v.ID + " " + err.Error())
				}

				updateUsers = append(updateUsers, v)
			} else {
				// 创建新记录
				if err := auth.CheckAuthorization(ctx, "CreateUser"); err != nil {
					return item, errors.New("CreateUser " + err.Error())
				}

				v.ID = uuid.Must(uuid.NewV4()).String()
				v.CreatedAt = timestampMillis
				v.CreatedBy = principalID

				if err := tx.Omit(clause.Associations).Table(TableName("users", ctx)).Create(v).Error; err != nil {
					return item, err
				}

				newUsers = append(newUsers, v)
			}
		}

		allItems := append(updateUsers, newUsers...)

		if err := tx.Model(item).Association("Users").Replace(allItems); err != nil {
			return item, err
		}

		event.AddNewValue("users", allItems)
	}

	// 记录事件
	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}

// ============================================================
// UserRole - Delete
// ============================================================

// DeleteUserRoleFunc 执行删除或恢复操作
func DeleteUserRoleFunc(ctx context.Context, r *GeneratedResolver, id string, operationType string, unscoped *bool) (err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item := &UserRole{}
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
	if err = tx.Unscoped().Table(TableName("user_roles", ctx)).Where("is_delete = ? and id = ?", status, id).First(item).Error; err != nil {
		return err
	}

	deletedAt := now.UnixNano() / 1e6

	// 创建事件记录
	event := NewEvent(EventMetadata{
		Type:        EventTypeDeleted,
		Entity:      "UserRole",
		EntityID:    id,
		Date:        deletedAt,
		PrincipalID: principalID,
	})

	// 执行删除或恢复
	if operationType == "recovery" {
		if err := tx.Unscoped().Table(TableName("user_roles", ctx)).Model(&item).Updates(map[string]interface{}{
			"IsDelete":  1,
			"DeletedAt": nil,
			"DeletedBy": nil,
		}).Error; err != nil {
			return err
		}
	} else {
		if unscoped != nil && *unscoped {
			// 物理删除
			if err := tx.Unscoped().Table(TableName("user_roles", ctx)).Model(&item).Delete(item).Error; err != nil {
				return err
			}
		} else {
			// 软删除
			if err := tx.Model(&item).Table(TableName("user_roles", ctx)).Updates(UserRole{
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

// DeleteUserRoles 批量删除 UserRole 实体
func (r *GeneratedMutationResolver) DeleteUserRoles(ctx context.Context, id []string, unscoped *bool) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteUserRoles(ctx, r.GeneratedResolver, id, unscoped)
	if err != nil {
		RollbackMutationContext(ctx, r.GeneratedResolver)
		return done, err
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

// DeleteUserRolesHandler 处理批量删除逻辑
func DeleteUserRolesHandler(ctx context.Context, r *GeneratedResolver, id []string, unscoped *bool) (bool, error) {
	for _, itemID := range id {
		if err := DeleteUserRoleFunc(ctx, r, itemID, "delete", unscoped); err != nil {
			return false, err
		}
	}
	return true, nil
}

// ============================================================
// UserRole - Recovery
// ============================================================

// RecoveryUserRoles 批量恢复 UserRole 实体
func (r *GeneratedMutationResolver) RecoveryUserRoles(ctx context.Context, id []string) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.RecoveryUserRoles(ctx, r.GeneratedResolver, id)
	if err != nil {
		return done, err
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

// RecoveryUserRolesHandler 处理批量恢复逻辑
func RecoveryUserRolesHandler(ctx context.Context, r *GeneratedResolver, id []string) (bool, error) {
	unscoped := false
	for _, itemID := range id {
		if err := DeleteUserRoleFunc(ctx, r, itemID, "recovery", &unscoped); err != nil {
			return false, err
		}
	}
	return true, nil
}

// ============================================================
// Tag - Create
// ============================================================

// CreateTag 创建 Tag 实体的解析器入口
func (r *GeneratedMutationResolver) CreateTag(ctx context.Context, input map[string]interface{}) (item *Tag, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreateTag(ctx, r.GeneratedResolver, input)
	if err != nil {
		RollbackMutationContext(ctx, r.GeneratedResolver)
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// CreateTagHandler 处理 Tag 创建逻辑
func CreateTagHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Tag, err error) {
	item = &Tag{}
	now := time.Now()
	timestampMillis := now.UnixNano() / 1e6
	principalID := GetPrincipalIDFromContext(ctx)
	tx := GetTransaction(ctx)

	// 创建事件记录
	event := NewEvent(EventMetadata{
		Type:        EventTypeCreated,
		Entity:      "Tag",
		EntityID:    item.ID,
		Date:        timestampMillis,
		PrincipalID: principalID,
	})

	// 解析输入变更
	var changes TagChanges
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

	if _, ok := input["name"]; ok && !utils.IsEmpty(input["name"]) {

		if item.Name != changes.Name {

			item.Name = changes.Name

			event.AddNewValue("name", changes.Name)
		}
	}

	if _, ok := input["color"]; ok && changes.Color != nil {

		if (item.Color != changes.Color) || (*item.Color != *changes.Color) {

			item.Color = changes.Color

			event.AddNewValue("color", changes.Color)
		}
	}

	if _, ok := input["isDelete"]; ok && changes.IsDelete != nil {

		if (item.IsDelete != changes.IsDelete) || (*item.IsDelete != *changes.IsDelete) {

			item.IsDelete = changes.IsDelete

			event.AddNewValue("isDelete", changes.IsDelete)
		}
	}

	if _, ok := input["weight"]; ok && changes.Weight != nil {

		if (item.Weight != changes.Weight) || (*item.Weight != *changes.Weight) {

			item.Weight = changes.Weight

			event.AddNewValue("weight", changes.Weight)
		}
	}

	if _, ok := input["state"]; ok && changes.State != nil {

		if (item.State != changes.State) || (*item.State != *changes.State) {

			item.State = changes.State

			event.AddNewValue("state", changes.State)
		}
	}

	// ========== 保存主实体 ==========
	if err := tx.Omit(clause.Associations).Table(TableName("tags", ctx)).Create(item).Error; err != nil {
		return item, err
	}

	// ========== 处理 OneToMany/ManyToMany 关系（关联表持有外键或中间表） ==========

	// ---------- ToMany: tasks (ManyToMany 中间表) ----------

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
				return item, errors.New("Task Detail " + err.Error())
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
				return item, fmt.Errorf("tasksIds " + strings.Join(differenceIds, ",") + " not found")
			}

			// ManyToMany: 使用 Replace 更新中间表
			if err := tx.Model(item).Association("Tasks").Replace(items); err != nil {
				return item, err
			}

		}
		event.AddNewValue("tasks", items)
	}

	// 方式2：通过嵌套对象创建/更新
	if _, ok := input["tasks"]; ok && !utils.IsNil(input["tasks"]) {
		newTasks := []*Task{}
		updateTasks := []*Task{}

		for index, v := range changes.Tasks {
			weight := int64(index + 1)
			v.Weight = &weight

			if !utils.IsEmpty(v.ID) {
				// 更新现有记录
				v.UpdatedAt = &timestampMillis
				v.UpdatedBy = principalID

				if err := auth.CheckAuthorization(ctx, "UpdateTask"); err != nil {
					return item, errors.New("UpdateTask " + err.Error())
				}
				if err := auth.CheckAuthorization(ctx, "Task"); err != nil {
					return item, errors.New("Task Detail " + err.Error())
				}

				tasksInput := utils.StructToMap(*v)
				if _, err := r.Handlers.UpdateTask(ctx, r, tasksInput["id"].(string), tasksInput); err != nil {
					return item, errors.New("Task ID " + v.ID + " " + err.Error())
				}

				updateTasks = append(updateTasks, v)
			} else {
				// 创建新记录
				if err := auth.CheckAuthorization(ctx, "CreateTask"); err != nil {
					return item, errors.New("CreateTask " + err.Error())
				}

				v.ID = uuid.Must(uuid.NewV4()).String()
				v.CreatedAt = timestampMillis
				v.CreatedBy = principalID

				// 保存新记录
				if err := tx.Omit(clause.Associations).Table(TableName("tasks", ctx)).Create(v).Error; err != nil {
					return item, err
				}

				newTasks = append(newTasks, v)
			}
		}

		allItems := append(updateTasks, newTasks...)

		// ManyToMany: 使用 Replace 更新中间表
		if err := tx.Model(item).Association("Tasks").Replace(allItems); err != nil {
			return item, err
		}

		event.AddNewValue("tasks", allItems)
	}

	// 记录事件
	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}

// ============================================================
// Tag - Update
// ============================================================

// UpdateTag 更新 Tag 实体的解析器入口
func (r *GeneratedMutationResolver) UpdateTag(ctx context.Context, id string, input map[string]interface{}) (item *Tag, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdateTag(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		RollbackMutationContext(ctx, r.GeneratedResolver)
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// UpdateTagHandler 处理 Tag 更新逻辑
func UpdateTagHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Tag, err error) {
	item = &Tag{}
	newItem := &Tag{}
	isChange := false

	now := time.Now()
	timestampMillis := now.UnixNano() / 1e6
	principalID := GetPrincipalIDFromContext(ctx)
	tx := GetTransaction(ctx)

	// 创建事件记录
	event := NewEvent(EventMetadata{
		Type:        EventTypeUpdated,
		Entity:      "Tag",
		EntityID:    id,
		Date:        timestampMillis,
		PrincipalID: principalID,
	})

	// 解析输入变更
	var changes TagChanges
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
	if err = GetItem(ctx, tx, TableName("tags", ctx), item, &id); err != nil {
		return nil, err
	}

	// 更新 UpdatedBy
	if item.UpdatedBy != nil && principalID != nil && *item.UpdatedBy != *principalID {
		newItem.UpdatedBy = principalID
	}

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

	if _, ok := input["name"]; ok {
		// 只要 input 中包含该字段，且值发生了变化（包括变为 null），就进行更新
		if item.Name != changes.Name {

			event.AddOldValue("name", item.Name)
			event.AddNewValue("name", changes.Name)

			item.Name = changes.Name
			newItem.Name = changes.Name
			changedFields = append(changedFields, "name")
			isChange = true
		}
	}

	if _, ok := input["color"]; ok {
		// 只要 input 中包含该字段，且值发生了变化（包括变为 null），就进行更新
		if (item.Color != changes.Color) && (item.Color == nil || changes.Color == nil || *item.Color != *changes.Color) {

			event.AddOldValue("color", item.Color)
			event.AddNewValue("color", changes.Color)

			item.Color = changes.Color
			newItem.Color = changes.Color
			changedFields = append(changedFields, "color")
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
		// 如果有更新 UpdatedBy，也需要添加到 Select 中
		if newItem.UpdatedBy != nil {
			changedFields = append(changedFields, "updated_by")
		}

		if err := tx.Table(TableName("tags", ctx)).Where("id = ?", id).Select(changedFields).Updates(newItem).Error; err != nil {
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
				return item, errors.New("Task Detail " + err.Error())
			}
			if err := tx.Find(&items, "id IN (?)", itemIds).Error; err != nil {
				return item, err
			}
			for _, v := range items {
				findIds = append(findIds, v.ID)
			}

			differenceIds := utils.Difference(itemIds, findIds)
			if len(differenceIds) > 0 {
				return item, fmt.Errorf("tasksIds " + strings.Join(differenceIds, ",") + " not found")
			}

			if err := tx.Model(item).Association("Tasks").Replace(items); err != nil {
				return item, err
			}

		} else {
			// 清空关联

			if err := tx.Model(item).Association("Tasks").Clear(); err != nil {
				return item, err
			}

		}
		event.AddNewValue("tasks", items)
	}

	// 方式2：通过嵌套对象创建/更新
	if _, ok := input["tasks"]; ok && !utils.IsNil(input["tasks"]) {
		newTasks := []*Task{}
		updateTasks := []*Task{}

		for index, v := range changes.Tasks {
			weight := int64(index + 1)
			v.Weight = &weight

			if !utils.IsEmpty(v.ID) {
				// 更新现有记录
				v.UpdatedAt = &timestampMillis
				v.UpdatedBy = principalID

				if err := auth.CheckAuthorization(ctx, "UpdateTask"); err != nil {
					return item, errors.New("UpdateTask " + err.Error())
				}
				if err := auth.CheckAuthorization(ctx, "Task"); err != nil {
					return item, errors.New("Task Detail " + err.Error())
				}

				tasksInput := utils.StructToMap(*v)
				if _, err := r.Handlers.UpdateTask(ctx, r, tasksInput["id"].(string), tasksInput); err != nil {
					return item, errors.New("Task ID " + v.ID + " " + err.Error())
				}

				updateTasks = append(updateTasks, v)
			} else {
				// 创建新记录
				if err := auth.CheckAuthorization(ctx, "CreateTask"); err != nil {
					return item, errors.New("CreateTask " + err.Error())
				}

				v.ID = uuid.Must(uuid.NewV4()).String()
				v.CreatedAt = timestampMillis
				v.CreatedBy = principalID

				if err := tx.Omit(clause.Associations).Table(TableName("tasks", ctx)).Create(v).Error; err != nil {
					return item, err
				}

				newTasks = append(newTasks, v)
			}
		}

		allItems := append(updateTasks, newTasks...)

		if err := tx.Model(item).Association("Tasks").Replace(allItems); err != nil {
			return item, err
		}

		event.AddNewValue("tasks", allItems)
	}

	// 记录事件
	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}

// ============================================================
// Tag - Delete
// ============================================================

// DeleteTagFunc 执行删除或恢复操作
func DeleteTagFunc(ctx context.Context, r *GeneratedResolver, id string, operationType string, unscoped *bool) (err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item := &Tag{}
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
	if err = tx.Unscoped().Table(TableName("tags", ctx)).Where("is_delete = ? and id = ?", status, id).First(item).Error; err != nil {
		return err
	}

	deletedAt := now.UnixNano() / 1e6

	// 创建事件记录
	event := NewEvent(EventMetadata{
		Type:        EventTypeDeleted,
		Entity:      "Tag",
		EntityID:    id,
		Date:        deletedAt,
		PrincipalID: principalID,
	})

	// 执行删除或恢复
	if operationType == "recovery" {
		if err := tx.Unscoped().Table(TableName("tags", ctx)).Model(&item).Updates(map[string]interface{}{
			"IsDelete":  1,
			"DeletedAt": nil,
			"DeletedBy": nil,
		}).Error; err != nil {
			return err
		}
	} else {
		if unscoped != nil && *unscoped {
			// 物理删除
			if err := tx.Unscoped().Table(TableName("tags", ctx)).Model(&item).Delete(item).Error; err != nil {
				return err
			}
		} else {
			// 软删除
			if err := tx.Model(&item).Table(TableName("tags", ctx)).Updates(Tag{
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

// DeleteTags 批量删除 Tag 实体
func (r *GeneratedMutationResolver) DeleteTags(ctx context.Context, id []string, unscoped *bool) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteTags(ctx, r.GeneratedResolver, id, unscoped)
	if err != nil {
		RollbackMutationContext(ctx, r.GeneratedResolver)
		return done, err
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

// DeleteTagsHandler 处理批量删除逻辑
func DeleteTagsHandler(ctx context.Context, r *GeneratedResolver, id []string, unscoped *bool) (bool, error) {
	for _, itemID := range id {
		if err := DeleteTagFunc(ctx, r, itemID, "delete", unscoped); err != nil {
			return false, err
		}
	}
	return true, nil
}

// ============================================================
// Tag - Recovery
// ============================================================

// RecoveryTags 批量恢复 Tag 实体
func (r *GeneratedMutationResolver) RecoveryTags(ctx context.Context, id []string) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.RecoveryTags(ctx, r.GeneratedResolver, id)
	if err != nil {
		return done, err
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

// RecoveryTagsHandler 处理批量恢复逻辑
func RecoveryTagsHandler(ctx context.Context, r *GeneratedResolver, id []string) (bool, error) {
	unscoped := false
	for _, itemID := range id {
		if err := DeleteTagFunc(ctx, r, itemID, "recovery", &unscoped); err != nil {
			return false, err
		}
	}
	return true, nil
}
