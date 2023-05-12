package gen

import (
	"context"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm/clause"
)

type GeneratedMutationResolver struct{ *GeneratedResolver }

type MutationEvents struct {
	Events []Event
}

func (r *GeneratedMutationResolver) CreateUser(ctx context.Context, input map[string]interface{}) (item *User, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreateUser(ctx, r.GeneratedResolver, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func CreateUserHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *User, err error) {
	item = &User{}

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
		tx.Rollback()
		return nil, err
	}

	item.ID = uuid.Must(uuid.NewV4()).String()
	item.CreatedAt = milliTime
	item.CreatedBy = principalID

	if input["tasks"] != nil && input["tasksIds"] != nil {
		tx.Rollback()
		return nil, fmt.Errorf("tasksIds and tasks cannot coexist")
	}

	var tasksIds []string

	if _, ok := input["tasks"]; ok {
		var tasksMaps []map[string]interface{}
		for _, v := range input["tasks"].([]interface{}) {
			tasksMaps = append(tasksMaps, v.(map[string]interface{}))
		}

		for _, v := range tasksMaps {
			var tasks *Task
			if v["id"] == nil {
				tasks, err = r.Handlers.CreateTask(ctx, r, v)
			} else {
				tasks, err = r.Handlers.UpdateTask(ctx, r, v["id"].(string), v)
			}

			changes.Tasks = append(changes.Tasks, tasks)
			tasksIds = append(tasksIds, tasks.ID)
		}
		event.AddNewValue("tasks", changes.Tasks)
		item.Tasks = changes.Tasks
	}

	if ids, exists := input["tasksIds"]; exists {
		for _, v := range ids.([]interface{}) {
			tasksIds = append(tasksIds, v.(string))
		}
	}

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

	if err := tx.Omit(clause.Associations).Create(item).Error; err != nil {
		tx.Rollback()
		return item, err
	}

	if len(tasksIds) > 0 {
		if err := tx.Model(&Task{}).Where("id IN(?)", tasksIds).Update("assignee_id", item.ID).Error; err != nil {
			tx.Rollback()
			return item, err
		}
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}
func (r *GeneratedMutationResolver) UpdateUser(ctx context.Context, id string, input map[string]interface{}) (item *User, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdateUser(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func UpdateUserHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *User, err error) {
	item = &User{}
	newItem := &User{}

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
		tx.Rollback()
		return nil, err
	}

	if input["tasks"] != nil && input["tasksIds"] != nil {
		tx.Rollback()
		return nil, fmt.Errorf("tasksIds and tasks cannot coexist")
	}

	if err = GetItem(ctx, tx, TableName("users"), item, &id); err != nil {
		tx.Rollback()
		return nil, err
	}

	if item.UpdatedBy != nil && principalID != nil && *item.UpdatedBy != *principalID {
		newItem.UpdatedBy = principalID
	}

	var tasksIds []string

	if _, ok := input["tasks"]; ok {
		if err := tx.Unscoped().Model(&Task{}).Where("assignee_id = ?", id).Update("assignee_id", "").Error; err != nil {
			tx.Rollback()
			return item, err
		}

		var tasksMaps []map[string]interface{}
		for _, v := range input["tasks"].([]interface{}) {
			vMaps := v.(map[string]interface{})
			tasksMaps = append(tasksMaps, vMaps)
		}

		for _, v := range tasksMaps {
			var tasks *Task
			v["assigneeId"] = id
			if v["id"] == nil {
				tasks, err = r.Handlers.CreateTask(ctx, r, v)
			} else {
				tasks, err = r.Handlers.UpdateTask(ctx, r, v["id"].(string), v)
			}

			changes.Tasks = append(changes.Tasks, tasks)
		}

		event.AddNewValue("tasks", changes.Tasks)
		item.Tasks = changes.Tasks
		newItem.Tasks = changes.Tasks
	}

	if ids, exists := input["tasksIds"]; exists {
		if err := tx.Unscoped().Model(&Task{}).Where("assignee_id = ?", id).Update("assignee_id", "").Error; err != nil {
			tx.Rollback()
			return item, err
		}
		for _, v := range ids.([]interface{}) {
			tasksIds = append(tasksIds, v.(string))
		}

		if len(tasksIds) > 0 {
			isChange = true
		}
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

	// if err := validator.Struct(item); err != nil {
	// 	tx.Rollback()
	// 	return nil, err
	// }

	if !isChange {
		return item, nil
	}

	if err := tx.Model(&newItem).Where("id = ?", id).Omit(clause.Associations).Updates(newItem).Error; err != nil {
		tx.Rollback()
		return item, err
	}

	if len(tasksIds) > 0 {
		if err := tx.Model(&Task{}).Where("id IN(?)", tasksIds).Update("assignee_id", item.ID).Error; err != nil {
			tx.Rollback()
			return item, err
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

	if err = GetItem(ctx, tx, TableName("users"), item, &id); err != nil {
		tx.Rollback()
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
		} else if err := tx.Model(&item).Updates(User{DeletedAt: &deletedAt, DeletedBy: principalID, UpdatedBy: principalID}).Error; err != nil {
			tx.Rollback()
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
	done, err := r.Handlers.DeleteUsers(ctx, r.GeneratedResolver, id, unscoped)
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

func DeleteUsersHandler(ctx context.Context, r *GeneratedResolver, id []string, unscoped *bool) (bool, error) {
	var err error = nil

	if len(id) > 0 {
		for _, v := range id {
			err = DeleteUserFunc(ctx, r, v, "delete", unscoped)
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

func (r *GeneratedMutationResolver) RecoveryUsers(ctx context.Context, id []string) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.RecoveryUsers(ctx, r.GeneratedResolver, id)
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

func RecoveryUsersHandler(ctx context.Context, r *GeneratedResolver, id []string) (bool, error) {
	var err error = nil

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
	item, err = r.Handlers.CreateTask(ctx, r.GeneratedResolver, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func CreateTaskHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Task, err error) {
	item = &Task{}

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
		tx.Rollback()
		return nil, err
	}

	item.ID = uuid.Must(uuid.NewV4()).String()
	item.CreatedAt = milliTime
	item.CreatedBy = principalID

	if input["assignee"] != nil && input["assigneeId"] != nil {
		tx.Rollback()
		return nil, fmt.Errorf("assigneeId and assignee cannot coexist")
	}

	if _, ok := input["assignee"]; ok {
		var assignee *User
		assigneeInput := input["assignee"].(map[string]interface{})

		if assigneeInput["id"] == nil {
			assignee, err = r.Handlers.CreateUser(ctx, r, assigneeInput)
		} else {
			assignee, err = r.Handlers.UpdateUser(ctx, r, assigneeInput["id"].(string), assigneeInput)
		}

		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf(fmt.Sprintf("assignee %s", err.Error()))
		}

		event.AddNewValue("assignee", changes.Assignee)
		item.Assignee = assignee
		item.AssigneeID = &assignee.ID
	}

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

	if _, ok := input["assigneeId"]; ok && (item.AssigneeID != changes.AssigneeID) && (item.AssigneeID == nil || changes.AssigneeID == nil || *item.AssigneeID != *changes.AssigneeID) {

		// if err := tx.Select("id").Where("id", input["assigneeId"]).First(&Assignee{}).Error; err != nil {
		// 	tx.Rollback()
		// 	return nil, fmt.Errorf("assigneeId " + err.Error())
		// }
		item.AssigneeID = changes.AssigneeID
		event.AddNewValue("assigneeId", changes.AssigneeID)
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
	item, err = r.Handlers.UpdateTask(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func UpdateTaskHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Task, err error) {
	item = &Task{}
	newItem := &Task{}

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
		tx.Rollback()
		return nil, err
	}

	if input["assignee"] != nil && input["assigneeId"] != nil {
		tx.Rollback()
		return nil, fmt.Errorf("assigneeId and assignee cannot coexist")
	}

	if err = GetItem(ctx, tx, TableName("tasks"), item, &id); err != nil {
		tx.Rollback()
		return nil, err
	}

	if item.UpdatedBy != nil && principalID != nil && *item.UpdatedBy != *principalID {
		newItem.UpdatedBy = principalID
	}

	if _, ok := input["assignee"]; ok {
		var assignee *User
		assigneeInput := input["assignee"].(map[string]interface{})

		if assigneeInput["id"] == nil {
			assignee, err = r.Handlers.CreateUser(ctx, r, assigneeInput)
		} else {
			assignee, err = r.Handlers.UpdateUser(ctx, r, assigneeInput["id"].(string), assigneeInput)
		}

		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf(fmt.Sprintf("assignee %s", err.Error()))
		}

		if err := tx.Model(&item).Association("Assignee").Clear(); err != nil {
			tx.Rollback()
			return nil, err
		}

		if err := tx.Model(&User{}).Where("id = ?", assignee.ID).Update("tasks_id", item.ID).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		event.AddOldValue("assignee", item.Assignee)
		event.AddNewValue("assignee", changes.Assignee)
		item.Assignee = assignee
		newItem.AssigneeID = &assignee.ID
		isChange = true
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

	if _, ok := input["assigneeId"]; ok && (item.AssigneeID != changes.AssigneeID) && (item.AssigneeID == nil || changes.AssigneeID == nil || *item.AssigneeID != *changes.AssigneeID) {

		// if err := tx.Select("id").Where("id", input["assigneeId"]).First(&Assignee{}).Error; err != nil {
		// 	tx.Rollback()
		// 	return nil, fmt.Errorf("assigneeId " + err.Error())
		// }
		event.AddOldValue("assigneeId", item.AssigneeID)
		event.AddNewValue("assigneeId", changes.AssigneeID)
		item.AssigneeID = changes.AssigneeID
		newItem.AssigneeID = changes.AssigneeID
		isChange = true
	}

	// if err := validator.Struct(item); err != nil {
	// 	tx.Rollback()
	// 	return nil, err
	// }

	if !isChange {
		return item, nil
	}

	if err := tx.Model(&newItem).Where("id = ?", id).Omit(clause.Associations).Updates(newItem).Error; err != nil {
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

	if err = GetItem(ctx, tx, TableName("tasks"), item, &id); err != nil {
		tx.Rollback()
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
		} else if err := tx.Model(&item).Updates(Task{DeletedAt: &deletedAt, DeletedBy: principalID, UpdatedBy: principalID}).Error; err != nil {
			tx.Rollback()
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
	done, err := r.Handlers.DeleteTasks(ctx, r.GeneratedResolver, id, unscoped)
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

func DeleteTasksHandler(ctx context.Context, r *GeneratedResolver, id []string, unscoped *bool) (bool, error) {
	var err error = nil

	if len(id) > 0 {
		for _, v := range id {
			err = DeleteTaskFunc(ctx, r, v, "delete", unscoped)
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

func (r *GeneratedMutationResolver) RecoveryTasks(ctx context.Context, id []string) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.RecoveryTasks(ctx, r.GeneratedResolver, id)
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

func RecoveryTasksHandler(ctx context.Context, r *GeneratedResolver, id []string) (bool, error) {
	var err error = nil

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

func (r *GeneratedMutationResolver) CreateUploadFile(ctx context.Context, input map[string]interface{}) (item *UploadFile, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreateUploadFile(ctx, r.GeneratedResolver, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func CreateUploadFileHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *UploadFile, err error) {
	item = &UploadFile{}

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
		Entity:      "UploadFile",
		EntityID:    item.ID,
		Date:        milliTime,
		PrincipalID: principalID,
	})

	var changes UploadFileChanges
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

	item.ID = uuid.Must(uuid.NewV4()).String()
	item.CreatedAt = milliTime
	item.CreatedBy = principalID

	if _, ok := input["name"]; ok && (item.Name != changes.Name) {
		item.Name = changes.Name
		event.AddNewValue("name", changes.Name)
	}

	if _, ok := input["hash"]; ok && (item.Hash != changes.Hash) {
		item.Hash = changes.Hash
		event.AddNewValue("hash", changes.Hash)
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
func (r *GeneratedMutationResolver) UpdateUploadFile(ctx context.Context, id string, input map[string]interface{}) (item *UploadFile, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdateUploadFile(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func UpdateUploadFileHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *UploadFile, err error) {
	item = &UploadFile{}
	newItem := &UploadFile{}

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
		Entity:      "UploadFile",
		EntityID:    id,
		Date:        milliTime,
		PrincipalID: principalID,
	})

	var changes UploadFileChanges
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

	if err = GetItem(ctx, tx, TableName("upload_files"), item, &id); err != nil {
		tx.Rollback()
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

	if _, ok := input["name"]; ok && (item.Name != changes.Name) {
		event.AddOldValue("name", item.Name)
		event.AddNewValue("name", changes.Name)
		item.Name = changes.Name
		newItem.Name = changes.Name
		isChange = true
	}

	if _, ok := input["hash"]; ok && (item.Hash != changes.Hash) {
		event.AddOldValue("hash", item.Hash)
		event.AddNewValue("hash", changes.Hash)
		item.Hash = changes.Hash
		newItem.Hash = changes.Hash
		isChange = true
	}

	// if err := validator.Struct(item); err != nil {
	// 	tx.Rollback()
	// 	return nil, err
	// }

	if !isChange {
		return item, nil
	}

	if err := tx.Model(&newItem).Where("id = ?", id).Omit(clause.Associations).Updates(newItem).Error; err != nil {
		tx.Rollback()
		return item, err
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}

func DeleteUploadFileFunc(ctx context.Context, r *GeneratedResolver, id string, tye string, unscoped *bool) (err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item := &UploadFile{}
	now := time.Now()
	tx := GetTransaction(ctx)
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err = GetItem(ctx, tx, TableName("upload_files"), item, &id); err != nil {
		tx.Rollback()
		return err
	}

	deletedAt := now.UnixNano() / 1e6

	event := NewEvent(EventMetadata{
		Type:        EventTypeDeleted,
		Entity:      "UploadFile",
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
		} else if err := tx.Model(&item).Updates(UploadFile{DeletedAt: &deletedAt, DeletedBy: principalID, UpdatedBy: principalID}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}

func (r *GeneratedMutationResolver) DeleteUploadFiles(ctx context.Context, id []string, unscoped *bool) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteUploadFiles(ctx, r.GeneratedResolver, id, unscoped)
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

func DeleteUploadFilesHandler(ctx context.Context, r *GeneratedResolver, id []string, unscoped *bool) (bool, error) {
	var err error = nil

	if len(id) > 0 {
		for _, v := range id {
			err = DeleteUploadFileFunc(ctx, r, v, "delete", unscoped)
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

func (r *GeneratedMutationResolver) RecoveryUploadFiles(ctx context.Context, id []string) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.RecoveryUploadFiles(ctx, r.GeneratedResolver, id)
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

func RecoveryUploadFilesHandler(ctx context.Context, r *GeneratedResolver, id []string) (bool, error) {
	var err error = nil

	var unscoped bool = false

	if len(id) > 0 {
		for _, v := range id {
			err = DeleteUploadFileFunc(ctx, r, v, "recovery", &unscoped)
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
