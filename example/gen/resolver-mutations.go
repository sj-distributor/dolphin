package gen

import (
	"context"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/sj-distributor/dolphin-example/auth"
	"github.com/sj-distributor/dolphin-example/utils"
	"gorm.io/gorm/clause"
)

type GeneratedMutationResolver struct{ *GeneratedResolver }

type MutationEvents struct {
	Events []Event
}

func (r *GeneratedMutationResolver) CreateUser(ctx context.Context, input map[string]interface{}) (item *User, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreateUser(ctx, r.GeneratedResolver, input, true)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func CreateUserHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}, authType bool) (item *User, err error) {
	item = &User{}
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
		Entity:      "User",
		EntityID:    item.ID,
		Date:        milliTime,
		PrincipalID: principalID,
	})

	var changes UserChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	err = CheckStructFieldIsEmpty(item, input)
	if err != nil {
		return nil, err
	}

	item.ID = uuid.Must(uuid.NewV4()).String()
	item.CreatedAt = milliTime
	item.CreatedBy = principalID

	if _, ok := input["phone"]; ok && (item.Phone != changes.Phone) {
		item.Phone = changes.Phone
		event.AddNewValue("phone", changes.Phone)
	}

	if _, ok := input["password"]; ok && (item.Password != changes.Password) {
		item.Password = changes.Password
		event.AddNewValue("password", changes.Password)
	}

	if _, ok := input["email"]; ok && (item.Email != changes.Email) && (item.Email == nil || changes.Email == nil || *item.Email != *changes.Email) {
		item.Email = changes.Email
		event.AddNewValue("email", changes.Email)
	}

	if _, ok := input["nickname"]; ok && (item.Nickname != changes.Nickname) && (item.Nickname == nil || changes.Nickname == nil || *item.Nickname != *changes.Nickname) {
		item.Nickname = changes.Nickname
		event.AddNewValue("nickname", changes.Nickname)
	}

	if _, ok := input["age"]; ok && (item.Age != changes.Age) && (item.Age == nil || changes.Age == nil || *item.Age != *changes.Age) {
		item.Age = changes.Age
		event.AddNewValue("age", changes.Age)
	}

	if _, ok := input["lastName"]; ok && (item.LastName != changes.LastName) && (item.LastName == nil || changes.LastName == nil || *item.LastName != *changes.LastName) {
		item.LastName = changes.LastName
		event.AddNewValue("lastName", changes.LastName)
	}

	if _, ok := input["isDelete"]; ok && (item.IsDelete != changes.IsDelete) && (item.IsDelete == nil || changes.IsDelete == nil || *item.IsDelete != *changes.IsDelete) {
		item.IsDelete = changes.IsDelete
		event.AddNewValue("isDelete", changes.IsDelete)
	}

	if _, ok := input["weight"]; ok && (item.Weight != changes.Weight) && (item.Weight == nil || changes.Weight == nil || *item.Weight != *changes.Weight) {
		item.Weight = changes.Weight
		event.AddNewValue("weight", changes.Weight)
	}

	if _, ok := input["state"]; ok && (item.State != changes.State) && (item.State == nil || changes.State == nil || *item.State != *changes.State) {
		item.State = changes.State
		event.AddNewValue("state", changes.State)
	}

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
func (r *GeneratedMutationResolver) UpdateUser(ctx context.Context, id string, input map[string]interface{}) (item *User, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdateUser(ctx, r.GeneratedResolver, id, input, true)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func UpdateUserHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}, authType bool) (item *User, err error) {
	item = &User{}
	newItem := &User{}
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
		Entity:      "User",
		EntityID:    id,
		Date:        milliTime,
		PrincipalID: principalID,
	})

	var changes UserChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	err = CheckStructFieldIsEmpty(item, input)
	if err != nil {
		return nil, err
	}

	if err = GetItem(ctx, tx, TableName("users"), item, &id); err != nil {
		return nil, err
	}

	if item.UpdatedBy != nil && principalID != nil && *item.UpdatedBy != *principalID {
		newItem.UpdatedBy = principalID
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		event.AddOldValue("id", item.ID)
		event.AddNewValue("id", changes.ID)
		item.ID = changes.ID
		newItem.ID = changes.ID
		isChange = true
	}

	if _, ok := input["phone"]; ok && (item.Phone != changes.Phone) {
		event.AddOldValue("phone", item.Phone)
		event.AddNewValue("phone", changes.Phone)
		item.Phone = changes.Phone
		newItem.Phone = changes.Phone
		isChange = true
	}

	if _, ok := input["password"]; ok && (item.Password != changes.Password) {
		event.AddOldValue("password", item.Password)
		event.AddNewValue("password", changes.Password)
		item.Password = changes.Password
		newItem.Password = changes.Password
		isChange = true
	}

	if _, ok := input["email"]; ok && (item.Email != changes.Email) && (item.Email == nil || changes.Email == nil || *item.Email != *changes.Email) {
		event.AddOldValue("email", item.Email)
		event.AddNewValue("email", changes.Email)
		item.Email = changes.Email
		newItem.Email = changes.Email
		isChange = true
	}

	if _, ok := input["nickname"]; ok && (item.Nickname != changes.Nickname) && (item.Nickname == nil || changes.Nickname == nil || *item.Nickname != *changes.Nickname) {
		event.AddOldValue("nickname", item.Nickname)
		event.AddNewValue("nickname", changes.Nickname)
		item.Nickname = changes.Nickname
		newItem.Nickname = changes.Nickname
		isChange = true
	}

	if _, ok := input["age"]; ok && (item.Age != changes.Age) && (item.Age == nil || changes.Age == nil || *item.Age != *changes.Age) {
		event.AddOldValue("age", item.Age)
		event.AddNewValue("age", changes.Age)
		item.Age = changes.Age
		newItem.Age = changes.Age
		isChange = true
	}

	if _, ok := input["lastName"]; ok && (item.LastName != changes.LastName) && (item.LastName == nil || changes.LastName == nil || *item.LastName != *changes.LastName) {
		event.AddOldValue("lastName", item.LastName)
		event.AddNewValue("lastName", changes.LastName)
		item.LastName = changes.LastName
		newItem.LastName = changes.LastName
		isChange = true
	}

	if _, ok := input["isDelete"]; ok && (item.IsDelete != changes.IsDelete) && (item.IsDelete == nil || changes.IsDelete == nil || *item.IsDelete != *changes.IsDelete) {
		event.AddOldValue("isDelete", item.IsDelete)
		event.AddNewValue("isDelete", changes.IsDelete)
		item.IsDelete = changes.IsDelete
		newItem.IsDelete = changes.IsDelete
		isChange = true
	}

	if _, ok := input["weight"]; ok && (item.Weight != changes.Weight) && (item.Weight == nil || changes.Weight == nil || *item.Weight != *changes.Weight) {
		event.AddOldValue("weight", item.Weight)
		event.AddNewValue("weight", changes.Weight)
		item.Weight = changes.Weight
		newItem.Weight = changes.Weight
		isChange = true
	}

	if _, ok := input["state"]; ok && (item.State != changes.State) && (item.State == nil || changes.State == nil || *item.State != *changes.State) {
		event.AddOldValue("state", item.State)
		event.AddNewValue("state", changes.State)
		item.State = changes.State
		newItem.State = changes.State
		isChange = true
	}

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

	if ids, exists := input["tasksIds"]; exists {
		if len(ids.([]interface{})) > 0 {
			items := []Task{}
			tx.Find(&items, "id IN (?)", ids)
			if err := tx.Model(&item).Association("Tasks").Replace(items); err != nil {
				tx.Rollback()
				return item, err
			}
		}
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}

func DeleteUserFunc(ctx context.Context, r *GeneratedResolver, id string, tye string, unscoped *bool) (err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item := &User{}
	now := time.Now()
	tx := GetTransaction(ctx)
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

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
		Entity:      "User",
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
		} else if err := tx.Model(&item).Updates(User{IsDelete: &isDelete, DeletedAt: &deletedAt, DeletedBy: principalID, UpdatedBy: principalID}).Error; err != nil {
			return err
		}
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}

func (r *GeneratedMutationResolver) DeleteUsers(ctx context.Context, id []string, unscoped *bool) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteUsers(ctx, r.GeneratedResolver, id, unscoped, true)
	if err != nil {
		return done, err
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

func DeleteUsersHandler(ctx context.Context, r *GeneratedResolver, id []string, unscoped *bool, authType bool) (bool, error) {
	tx := GetTransaction(ctx)
	var err error = nil
	if err := auth.CheckRouterAuth(ctx, authType); err != nil {
		return false, err
	}

	if len(id) > 0 {
		for _, v := range id {
			err = DeleteUserFunc(ctx, r, v, "delete", unscoped)
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

func (r *GeneratedMutationResolver) RecoveryUsers(ctx context.Context, id []string) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.RecoveryUsers(ctx, r.GeneratedResolver, id, true)
	if err != nil {
		return done, err
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

func RecoveryUsersHandler(ctx context.Context, r *GeneratedResolver, id []string, authType bool) (bool, error) {
	var err error = nil
	if err := auth.CheckRouterAuth(ctx, authType); err != nil {
		return false, err
	}

	var unscoped bool = false

	if len(id) > 0 {
		for _, v := range id {
			err = DeleteUserFunc(ctx, r, v, "recovery", &unscoped)
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

func (r *GeneratedMutationResolver) CreateTask(ctx context.Context, input map[string]interface{}) (item *Task, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreateTask(ctx, r.GeneratedResolver, input, true)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func CreateTaskHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}, authType bool) (item *Task, err error) {
	item = &Task{}
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
		Entity:      "Task",
		EntityID:    item.ID,
		Date:        milliTime,
		PrincipalID: principalID,
	})

	var changes TaskChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	err = CheckStructFieldIsEmpty(item, input)
	if err != nil {
		return nil, err
	}

	item.ID = uuid.Must(uuid.NewV4()).String()
	item.CreatedAt = milliTime
	item.CreatedBy = principalID

	if _, ok := input["title"]; ok && (item.Title != changes.Title) && (item.Title == nil || changes.Title == nil || *item.Title != *changes.Title) {
		item.Title = changes.Title
		event.AddNewValue("title", changes.Title)
	}

	if _, ok := input["completed"]; ok && (item.Completed != changes.Completed) && (item.Completed == nil || changes.Completed == nil || *item.Completed != *changes.Completed) {
		item.Completed = changes.Completed
		event.AddNewValue("completed", changes.Completed)
	}

	if _, ok := input["dueDate"]; ok && (item.DueDate != changes.DueDate) && (item.DueDate == nil || changes.DueDate == nil || *item.DueDate != *changes.DueDate) {
		item.DueDate = changes.DueDate
		event.AddNewValue("dueDate", changes.DueDate)
	}

	if _, ok := input["userId"]; ok && (item.UserID != changes.UserID) && (item.UserID == nil || changes.UserID == nil || *item.UserID != *changes.UserID) {

		if err := tx.Select("id").Where("id", input["userId"]).First(&User{}).Error; err != nil {
			return nil, fmt.Errorf("userId " + err.Error())
		}
		item.UserID = changes.UserID
		event.AddNewValue("userId", changes.UserID)
	}

	if _, ok := input["isDelete"]; ok && (item.IsDelete != changes.IsDelete) && (item.IsDelete == nil || changes.IsDelete == nil || *item.IsDelete != *changes.IsDelete) {
		item.IsDelete = changes.IsDelete
		event.AddNewValue("isDelete", changes.IsDelete)
	}

	if _, ok := input["weight"]; ok && (item.Weight != changes.Weight) && (item.Weight == nil || changes.Weight == nil || *item.Weight != *changes.Weight) {
		item.Weight = changes.Weight
		event.AddNewValue("weight", changes.Weight)
	}

	if _, ok := input["state"]; ok && (item.State != changes.State) && (item.State == nil || changes.State == nil || *item.State != *changes.State) {
		item.State = changes.State
		event.AddNewValue("state", changes.State)
	}

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
func (r *GeneratedMutationResolver) UpdateTask(ctx context.Context, id string, input map[string]interface{}) (item *Task, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdateTask(ctx, r.GeneratedResolver, id, input, true)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func UpdateTaskHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}, authType bool) (item *Task, err error) {
	item = &Task{}
	newItem := &Task{}
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
		Entity:      "Task",
		EntityID:    id,
		Date:        milliTime,
		PrincipalID: principalID,
	})

	var changes TaskChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	err = CheckStructFieldIsEmpty(item, input)
	if err != nil {
		return nil, err
	}

	if err = GetItem(ctx, tx, TableName("tasks"), item, &id); err != nil {
		return nil, err
	}

	if item.UpdatedBy != nil && principalID != nil && *item.UpdatedBy != *principalID {
		newItem.UpdatedBy = principalID
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		event.AddOldValue("id", item.ID)
		event.AddNewValue("id", changes.ID)
		item.ID = changes.ID
		newItem.ID = changes.ID
		isChange = true
	}

	if _, ok := input["title"]; ok && (item.Title != changes.Title) && (item.Title == nil || changes.Title == nil || *item.Title != *changes.Title) {
		event.AddOldValue("title", item.Title)
		event.AddNewValue("title", changes.Title)
		item.Title = changes.Title
		newItem.Title = changes.Title
		isChange = true
	}

	if _, ok := input["completed"]; ok && (item.Completed != changes.Completed) && (item.Completed == nil || changes.Completed == nil || *item.Completed != *changes.Completed) {
		event.AddOldValue("completed", item.Completed)
		event.AddNewValue("completed", changes.Completed)
		item.Completed = changes.Completed
		newItem.Completed = changes.Completed
		isChange = true
	}

	if _, ok := input["dueDate"]; ok && (item.DueDate != changes.DueDate) && (item.DueDate == nil || changes.DueDate == nil || *item.DueDate != *changes.DueDate) {
		event.AddOldValue("dueDate", item.DueDate)
		event.AddNewValue("dueDate", changes.DueDate)
		item.DueDate = changes.DueDate
		newItem.DueDate = changes.DueDate
		isChange = true
	}

	if _, ok := input["userId"]; ok && (item.UserID != changes.UserID) && (item.UserID == nil || changes.UserID == nil || *item.UserID != *changes.UserID) {

		if err := tx.Select("id").Where("id", input["userId"]).First(&User{}).Error; err != nil {
			return nil, fmt.Errorf("userId " + err.Error())
		}
		event.AddOldValue("userId", item.UserID)
		event.AddNewValue("userId", changes.UserID)
		item.UserID = changes.UserID
		newItem.UserID = changes.UserID
		isChange = true
	}

	if _, ok := input["isDelete"]; ok && (item.IsDelete != changes.IsDelete) && (item.IsDelete == nil || changes.IsDelete == nil || *item.IsDelete != *changes.IsDelete) {
		event.AddOldValue("isDelete", item.IsDelete)
		event.AddNewValue("isDelete", changes.IsDelete)
		item.IsDelete = changes.IsDelete
		newItem.IsDelete = changes.IsDelete
		isChange = true
	}

	if _, ok := input["weight"]; ok && (item.Weight != changes.Weight) && (item.Weight == nil || changes.Weight == nil || *item.Weight != *changes.Weight) {
		event.AddOldValue("weight", item.Weight)
		event.AddNewValue("weight", changes.Weight)
		item.Weight = changes.Weight
		newItem.Weight = changes.Weight
		isChange = true
	}

	if _, ok := input["state"]; ok && (item.State != changes.State) && (item.State == nil || changes.State == nil || *item.State != *changes.State) {
		event.AddOldValue("state", item.State)
		event.AddNewValue("state", changes.State)
		item.State = changes.State
		newItem.State = changes.State
		isChange = true
	}

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

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}

func DeleteTaskFunc(ctx context.Context, r *GeneratedResolver, id string, tye string, unscoped *bool) (err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item := &Task{}
	now := time.Now()
	tx := GetTransaction(ctx)
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

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
		Entity:      "Task",
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
		} else if err := tx.Model(&item).Updates(Task{IsDelete: &isDelete, DeletedAt: &deletedAt, DeletedBy: principalID, UpdatedBy: principalID}).Error; err != nil {
			return err
		}
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}

func (r *GeneratedMutationResolver) DeleteTasks(ctx context.Context, id []string, unscoped *bool) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteTasks(ctx, r.GeneratedResolver, id, unscoped, true)
	if err != nil {
		return done, err
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

func DeleteTasksHandler(ctx context.Context, r *GeneratedResolver, id []string, unscoped *bool, authType bool) (bool, error) {
	tx := GetTransaction(ctx)
	var err error = nil
	if err := auth.CheckRouterAuth(ctx, authType); err != nil {
		return false, err
	}

	if len(id) > 0 {
		for _, v := range id {
			err = DeleteTaskFunc(ctx, r, v, "delete", unscoped)
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

func (r *GeneratedMutationResolver) RecoveryTasks(ctx context.Context, id []string) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.RecoveryTasks(ctx, r.GeneratedResolver, id, true)
	if err != nil {
		return done, err
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

func RecoveryTasksHandler(ctx context.Context, r *GeneratedResolver, id []string, authType bool) (bool, error) {
	var err error = nil
	if err := auth.CheckRouterAuth(ctx, authType); err != nil {
		return false, err
	}

	var unscoped bool = false

	if len(id) > 0 {
		for _, v := range id {
			err = DeleteTaskFunc(ctx, r, v, "recovery", &unscoped)
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
