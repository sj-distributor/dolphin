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
		tx.Rollback()
		return nil, err
	}

	item.ID = uuid.Must(uuid.NewV4()).String()
	item.CreatedAt = milliTime
	item.CreatedBy = principalID

	if input["t"] != nil && input["tId"] != nil {
		tx.Rollback()
		return nil, fmt.Errorf("tId and t cannot coexist")
	}

	if _, ok := input["t"]; ok {
		var t *Task
		tInput := input["t"].(map[string]interface{})

		if tInput["id"] == nil {

			// one to one
			tInput["uId"] = item.ID

			t, err = r.Handlers.CreateTask(ctx, r, tInput, authType)
		} else {
			t, err = r.Handlers.UpdateTask(ctx, r, tInput["id"].(string), tInput, authType)
		}

		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf(fmt.Sprintf("t %s", err.Error()))
		}

		// if err := tx.Model(&Task{}).Where("id = ?", t.ID).Updates(Task{ UID: &item.ID}).Error; err != nil {
		// 	tx.Rollback()
		// 	return nil, err
		// }

		event.AddNewValue("t", changes.T)
		item.T = t
		item.TID = &t.ID
	}

	if _, ok := input["tt"]; ok {
		var tt *Task
		ttInput := input["tt"].(map[string]interface{})

		if ttInput["id"] == nil {

			tt, err = r.Handlers.CreateTask(ctx, r, ttInput, authType)
		} else {
			tt, err = r.Handlers.UpdateTask(ctx, r, ttInput["id"].(string), ttInput, authType)
		}

		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf(fmt.Sprintf("tt %s", err.Error()))
		}

		event.AddNewValue("tt", changes.Tt)
		item.Tt = tt
		item.TtID = &tt.ID
	}

	if input["ttt"] != nil && input["tttIds"] != nil {
		tx.Rollback()
		return nil, fmt.Errorf("tttIds and ttt cannot coexist")
	}

	var tttIds []string
	var createTtt []*Task
	var updateTtt []*Task
	if _, ok := input["ttt"]; ok {
		for _, v := range changes.Ttt {
			if v.ID == "" {
				v.ID = uuid.Must(uuid.NewV4()).String()
				v.CreatedAt = milliTime
				v.CreatedBy = principalID
				createTtt = append(createTtt, v)
			} else {
				opts := QueryTaskHandlerOptions{
					ID: &v.ID,
				}
				if _, err = r.Handlers.QueryTask(ctx, r, opts, authType); err != nil {
					tx.Rollback()
					return nil, err
				}
				v.UpdatedAt = &milliTime
				v.UpdatedBy = principalID
				updateTtt = append(updateTtt, v)
			}
		}
		event.AddNewValue("ttt", changes.Ttt)
		item.Ttt = changes.Ttt
	}

	if ids, exists := input["tttIds"]; exists {
		for _, v := range ids.([]interface{}) {
			tttIds = append(tttIds, v.(string))
		}
	}

	if input["tttt"] != nil && input["ttttIds"] != nil {
		tx.Rollback()
		return nil, fmt.Errorf("ttttIds and tttt cannot coexist")
	}

	var ttttIds []string
	var createTttt []*Task
	var updateTttt []*Task
	if _, ok := input["tttt"]; ok {
		for _, v := range changes.Tttt {
			if v.ID == "" {
				v.ID = uuid.Must(uuid.NewV4()).String()
				v.CreatedAt = milliTime
				v.CreatedBy = principalID
				createTttt = append(createTttt, v)
			} else {
				opts := QueryTaskHandlerOptions{
					ID: &v.ID,
				}
				if _, err = r.Handlers.QueryTask(ctx, r, opts, authType); err != nil {
					tx.Rollback()
					return nil, err
				}
				v.UpdatedAt = &milliTime
				v.UpdatedBy = principalID
				updateTttt = append(updateTttt, v)
			}
		}
		event.AddNewValue("tttt", changes.Tttt)
		item.Tttt = changes.Tttt
	}

	if ids, exists := input["ttttIds"]; exists {
		for _, v := range ids.([]interface{}) {
			ttttIds = append(ttttIds, v.(string))
		}
	}

	if _, ok := input["phone"]; ok && (item.Phone != changes.Phone) {
		item.Phone = changes.Phone
		event.AddNewValue("phone", changes.Phone)
	}

	if _, ok := input["tId"]; ok && (item.TID != changes.TID) && (item.TID == nil || changes.TID == nil || *item.TID != *changes.TID) {

		// if err := tx.Select("id").Where("id", input["tId"]).First(&T{}).Error; err != nil {
		// 	tx.Rollback()
		// 	return nil, fmt.Errorf("tId " + err.Error())
		// }
		item.TID = changes.TID
		event.AddNewValue("tId", changes.TID)
	}

	if _, ok := input["ttId"]; ok && (item.TtID != changes.TtID) && (item.TtID == nil || changes.TtID == nil || *item.TtID != *changes.TtID) {

		// if err := tx.Select("id").Where("id", input["ttId"]).First(&Tt{}).Error; err != nil {
		// 	tx.Rollback()
		// 	return nil, fmt.Errorf("ttId " + err.Error())
		// }
		item.TtID = changes.TtID
		event.AddNewValue("ttId", changes.TtID)
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
		tx.Rollback()
		return nil, err
	}

	if err := tx.Omit(clause.Associations).Create(item).Error; err != nil {
		tx.Rollback()
		return item, err
	}

	// todo 添加权限验证
	if len(createTtt) > 0 {
		if err := tx.Model(&item).Association("Ttt").Append(createTtt); err != nil {
			tx.Rollback()
			return item, err
		}
	}
	if len(updateTtt) > 0 {
		if err := tx.Model(&item).Association("Ttt").Replace(updateTtt); err != nil {
			tx.Rollback()
			return item, err
		}
	}

	// todo 添加权限验证
	if len(createTttt) > 0 {
		if err := tx.Model(&item).Association("Tttt").Append(createTttt); err != nil {
			tx.Rollback()
			return item, err
		}
	}
	if len(updateTttt) > 0 {
		if err := tx.Model(&item).Association("Tttt").Replace(updateTttt); err != nil {
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
		tx.Rollback()
		return nil, err
	}

	if input["t"] != nil && input["tId"] != nil {
		tx.Rollback()
		return nil, fmt.Errorf("tId and t cannot coexist")
	}

	if input["tt"] != nil && input["ttId"] != nil {
		tx.Rollback()
		return nil, fmt.Errorf("ttId and tt cannot coexist")
	}

	if input["ttt"] != nil && input["tttIds"] != nil {
		tx.Rollback()
		return nil, fmt.Errorf("tttIds and ttt cannot coexist")
	}

	if input["tttt"] != nil && input["ttttIds"] != nil {
		tx.Rollback()
		return nil, fmt.Errorf("ttttIds and tttt cannot coexist")
	}

	if err = GetItem(ctx, tx, TableName("users"), item, &id); err != nil {
		tx.Rollback()
		return nil, err
	}

	if item.UpdatedBy != nil && principalID != nil && *item.UpdatedBy != *principalID {
		newItem.UpdatedBy = principalID
	}

	if _, ok := input["t"]; ok {
		var t *Task
		tInput := input["t"].(map[string]interface{})

		if tInput["id"] == nil {
			t, err = r.Handlers.CreateTask(ctx, r, tInput, authType)
		} else {
			t, err = r.Handlers.UpdateTask(ctx, r, tInput["id"].(string), tInput, authType)
		}

		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf(fmt.Sprintf("t %s", err.Error()))
		}

		event.AddOldValue("t", item.T)
		event.AddNewValue("t", changes.T)
		item.T = t
		newItem.TID = &t.ID
		isChange = true
	}

	if _, ok := input["tt"]; ok {
		var tt *Task
		ttInput := input["tt"].(map[string]interface{})

		if ttInput["id"] == nil {
			tt, err = r.Handlers.CreateTask(ctx, r, ttInput, authType)
		} else {
			tt, err = r.Handlers.UpdateTask(ctx, r, ttInput["id"].(string), ttInput, authType)
		}

		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf(fmt.Sprintf("tt %s", err.Error()))
		}

		event.AddOldValue("tt", item.Tt)
		event.AddNewValue("tt", changes.Tt)
		item.Tt = tt
		newItem.TtID = &tt.ID
		isChange = true
	}

	var tttIds []string

	if _, ok := input["ttt"]; ok {
		if err := tx.Unscoped().Model(&Task{}).Where("uuu_id = ?", id).Update("uuu_id", "").Error; err != nil {
			tx.Rollback()
			return item, err
		}

		var tttMaps []map[string]interface{}
		for _, v := range input["ttt"].([]interface{}) {
			vMaps := v.(map[string]interface{})
			tttMaps = append(tttMaps, vMaps)
		}

		for _, v := range tttMaps {
			var ttt *Task
			v["uuuId"] = id
			if v["id"] == nil {
				ttt, err = r.Handlers.CreateTask(ctx, r, v, authType)
			} else {
				ttt, err = r.Handlers.UpdateTask(ctx, r, v["id"].(string), v, authType)
			}

			changes.Ttt = append(changes.Ttt, ttt)
		}

		event.AddNewValue("ttt", changes.Ttt)
		item.Ttt = changes.Ttt
		newItem.Ttt = changes.Ttt
	}

	if ids, exists := input["tttIds"]; exists {
		if err := tx.Unscoped().Model(&Task{}).Where("uuu_id = ?", id).Update("uuu_id", "").Error; err != nil {
			tx.Rollback()
			return item, err
		}
		for _, v := range ids.([]interface{}) {
			tttIds = append(tttIds, v.(string))
		}

		if len(tttIds) > 0 {
			isChange = true
		}
	}

	var ttttIds []string

	if _, ok := input["tttt"]; ok {
		if err := tx.Unscoped().Model(&Task{}).Where("uuuu_id = ?", id).Update("uuuu_id", "").Error; err != nil {
			tx.Rollback()
			return item, err
		}

		var ttttMaps []map[string]interface{}
		for _, v := range input["tttt"].([]interface{}) {
			vMaps := v.(map[string]interface{})
			ttttMaps = append(ttttMaps, vMaps)
		}

		for _, v := range ttttMaps {
			var tttt *Task
			v["uuuuId"] = id
			if v["id"] == nil {
				tttt, err = r.Handlers.CreateTask(ctx, r, v, authType)
			} else {
				tttt, err = r.Handlers.UpdateTask(ctx, r, v["id"].(string), v, authType)
			}

			changes.Tttt = append(changes.Tttt, tttt)
		}

		event.AddNewValue("tttt", changes.Tttt)
		item.Tttt = changes.Tttt
		newItem.Tttt = changes.Tttt
	}

	if ids, exists := input["ttttIds"]; exists {
		if err := tx.Unscoped().Model(&Task{}).Where("uuuu_id = ?", id).Update("uuuu_id", "").Error; err != nil {
			tx.Rollback()
			return item, err
		}
		for _, v := range ids.([]interface{}) {
			ttttIds = append(ttttIds, v.(string))
		}

		if len(ttttIds) > 0 {
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

	if _, ok := input["tId"]; ok && (item.TID != changes.TID) && (item.TID == nil || changes.TID == nil || *item.TID != *changes.TID) {

		// if err := tx.Select("id").Where("id", input["tId"]).First(&T{}).Error; err != nil {
		// 	tx.Rollback()
		// 	return nil, fmt.Errorf("tId " + err.Error())
		// }
		event.AddOldValue("tId", item.TID)
		event.AddNewValue("tId", changes.TID)
		item.TID = changes.TID
		newItem.TID = changes.TID
		isChange = true
	}

	if _, ok := input["ttId"]; ok && (item.TtID != changes.TtID) && (item.TtID == nil || changes.TtID == nil || *item.TtID != *changes.TtID) {

		// if err := tx.Select("id").Where("id", input["ttId"]).First(&Tt{}).Error; err != nil {
		// 	tx.Rollback()
		// 	return nil, fmt.Errorf("ttId " + err.Error())
		// }
		event.AddOldValue("ttId", item.TtID)
		event.AddNewValue("ttId", changes.TtID)
		item.TtID = changes.TtID
		newItem.TtID = changes.TtID
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
		tx.Rollback()
		return nil, err
	}

	if !isChange {
		return item, nil
	}

	if err := tx.Model(&newItem).Where("id = ?", id).Omit(clause.Associations).Updates(newItem).Error; err != nil {
		tx.Rollback()
		return item, err
	}

	if len(tttIds) > 0 {
		if err := tx.Model(&Task{}).Where("id IN(?)", tttIds).Update("uuu_id", item.ID).Error; err != nil {
			tx.Rollback()
			return item, err
		}
	}

	if len(ttttIds) > 0 {
		if err := tx.Model(&Task{}).Where("id IN(?)", ttttIds).Update("uuuu_id", item.ID).Error; err != nil {
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
		if unscoped != nil && *unscoped == true {
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
		tx.Rollback()
		return nil, err
	}

	item.ID = uuid.Must(uuid.NewV4()).String()
	item.CreatedAt = milliTime
	item.CreatedBy = principalID

	if input["u"] != nil && input["uId"] != nil {
		tx.Rollback()
		return nil, fmt.Errorf("uId and u cannot coexist")
	}

	if _, ok := input["u"]; ok {
		var u *User
		uInput := input["u"].(map[string]interface{})

		if uInput["id"] == nil {

			// one to one
			uInput["tId"] = item.ID

			u, err = r.Handlers.CreateUser(ctx, r, uInput, authType)
		} else {
			u, err = r.Handlers.UpdateUser(ctx, r, uInput["id"].(string), uInput, authType)
		}

		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf(fmt.Sprintf("u %s", err.Error()))
		}

		// if err := tx.Model(&User{}).Where("id = ?", u.ID).Updates(User{ TID: &item.ID}).Error; err != nil {
		// 	tx.Rollback()
		// 	return nil, err
		// }

		event.AddNewValue("u", changes.U)
		item.U = u
		item.UID = &u.ID
	}

	if input["uu"] != nil && input["uuIds"] != nil {
		tx.Rollback()
		return nil, fmt.Errorf("uuIds and uu cannot coexist")
	}

	var uuIds []string
	var createUu []*User
	var updateUu []*User
	if _, ok := input["uu"]; ok {
		for _, v := range changes.Uu {
			if v.ID == "" {
				v.ID = uuid.Must(uuid.NewV4()).String()
				v.CreatedAt = milliTime
				v.CreatedBy = principalID
				createUu = append(createUu, v)
			} else {
				opts := QueryUserHandlerOptions{
					ID: &v.ID,
				}
				if _, err = r.Handlers.QueryUser(ctx, r, opts, authType); err != nil {
					tx.Rollback()
					return nil, err
				}
				v.UpdatedAt = &milliTime
				v.UpdatedBy = principalID
				updateUu = append(updateUu, v)
			}
		}
		event.AddNewValue("uu", changes.Uu)
		item.Uu = changes.Uu
	}

	if ids, exists := input["uuIds"]; exists {
		for _, v := range ids.([]interface{}) {
			uuIds = append(uuIds, v.(string))
		}
	}

	if _, ok := input["uuu"]; ok {
		var uuu *User
		uuuInput := input["uuu"].(map[string]interface{})

		if uuuInput["id"] == nil {

			uuu, err = r.Handlers.CreateUser(ctx, r, uuuInput, authType)
		} else {
			uuu, err = r.Handlers.UpdateUser(ctx, r, uuuInput["id"].(string), uuuInput, authType)
		}

		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf(fmt.Sprintf("uuu %s", err.Error()))
		}

		event.AddNewValue("uuu", changes.Uuu)
		item.Uuu = uuu
		item.UuuID = &uuu.ID
	}

	if input["uuuu"] != nil && input["uuuuIds"] != nil {
		tx.Rollback()
		return nil, fmt.Errorf("uuuuIds and uuuu cannot coexist")
	}

	var uuuuIds []string
	var createUuuu []*User
	var updateUuuu []*User
	if _, ok := input["uuuu"]; ok {
		for _, v := range changes.Uuuu {
			if v.ID == "" {
				v.ID = uuid.Must(uuid.NewV4()).String()
				v.CreatedAt = milliTime
				v.CreatedBy = principalID
				createUuuu = append(createUuuu, v)
			} else {
				opts := QueryUserHandlerOptions{
					ID: &v.ID,
				}
				if _, err = r.Handlers.QueryUser(ctx, r, opts, authType); err != nil {
					tx.Rollback()
					return nil, err
				}
				v.UpdatedAt = &milliTime
				v.UpdatedBy = principalID
				updateUuuu = append(updateUuuu, v)
			}
		}
		event.AddNewValue("uuuu", changes.Uuuu)
		item.Uuuu = changes.Uuuu
	}

	if ids, exists := input["uuuuIds"]; exists {
		for _, v := range ids.([]interface{}) {
			uuuuIds = append(uuuuIds, v.(string))
		}
	}

	if _, ok := input["title"]; ok && (item.Title != changes.Title) && (item.Title == nil || changes.Title == nil || *item.Title != *changes.Title) {
		item.Title = changes.Title
		event.AddNewValue("title", changes.Title)
	}

	if _, ok := input["uId"]; ok && (item.UID != changes.UID) && (item.UID == nil || changes.UID == nil || *item.UID != *changes.UID) {

		// if err := tx.Select("id").Where("id", input["uId"]).First(&U{}).Error; err != nil {
		// 	tx.Rollback()
		// 	return nil, fmt.Errorf("uId " + err.Error())
		// }
		item.UID = changes.UID
		event.AddNewValue("uId", changes.UID)
	}

	if _, ok := input["uuuId"]; ok && (item.UuuID != changes.UuuID) && (item.UuuID == nil || changes.UuuID == nil || *item.UuuID != *changes.UuuID) {

		// if err := tx.Select("id").Where("id", input["uuuId"]).First(&Uuu{}).Error; err != nil {
		// 	tx.Rollback()
		// 	return nil, fmt.Errorf("uuuId " + err.Error())
		// }
		item.UuuID = changes.UuuID
		event.AddNewValue("uuuId", changes.UuuID)
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
		tx.Rollback()
		return nil, err
	}

	if err := tx.Omit(clause.Associations).Create(item).Error; err != nil {
		tx.Rollback()
		return item, err
	}

	// todo 添加权限验证
	if len(createUu) > 0 {
		if err := tx.Model(&item).Association("Uu").Append(createUu); err != nil {
			tx.Rollback()
			return item, err
		}
	}
	if len(updateUu) > 0 {
		if err := tx.Model(&item).Association("Uu").Replace(updateUu); err != nil {
			tx.Rollback()
			return item, err
		}
	}

	// todo 添加权限验证
	if len(createUuuu) > 0 {
		if err := tx.Model(&item).Association("Uuuu").Append(createUuuu); err != nil {
			tx.Rollback()
			return item, err
		}
	}
	if len(updateUuuu) > 0 {
		if err := tx.Model(&item).Association("Uuuu").Replace(updateUuuu); err != nil {
			tx.Rollback()
			return item, err
		}
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
		tx.Rollback()
		return nil, err
	}

	if input["u"] != nil && input["uId"] != nil {
		tx.Rollback()
		return nil, fmt.Errorf("uId and u cannot coexist")
	}

	if input["uu"] != nil && input["uuIds"] != nil {
		tx.Rollback()
		return nil, fmt.Errorf("uuIds and uu cannot coexist")
	}

	if input["uuu"] != nil && input["uuuId"] != nil {
		tx.Rollback()
		return nil, fmt.Errorf("uuuId and uuu cannot coexist")
	}

	if input["uuuu"] != nil && input["uuuuIds"] != nil {
		tx.Rollback()
		return nil, fmt.Errorf("uuuuIds and uuuu cannot coexist")
	}

	if err = GetItem(ctx, tx, TableName("tasks"), item, &id); err != nil {
		tx.Rollback()
		return nil, err
	}

	if item.UpdatedBy != nil && principalID != nil && *item.UpdatedBy != *principalID {
		newItem.UpdatedBy = principalID
	}

	if _, ok := input["u"]; ok {
		var u *User
		uInput := input["u"].(map[string]interface{})

		if uInput["id"] == nil {
			u, err = r.Handlers.CreateUser(ctx, r, uInput, authType)
		} else {
			u, err = r.Handlers.UpdateUser(ctx, r, uInput["id"].(string), uInput, authType)
		}

		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf(fmt.Sprintf("u %s", err.Error()))
		}

		event.AddOldValue("u", item.U)
		event.AddNewValue("u", changes.U)
		item.U = u
		newItem.UID = &u.ID
		isChange = true
	}

	var uuIds []string

	if _, ok := input["uu"]; ok {
		if err := tx.Unscoped().Model(&User{}).Where("tt_id = ?", id).Update("tt_id", "").Error; err != nil {
			tx.Rollback()
			return item, err
		}

		var uuMaps []map[string]interface{}
		for _, v := range input["uu"].([]interface{}) {
			vMaps := v.(map[string]interface{})
			uuMaps = append(uuMaps, vMaps)
		}

		for _, v := range uuMaps {
			var uu *User
			v["ttId"] = id
			if v["id"] == nil {
				uu, err = r.Handlers.CreateUser(ctx, r, v, authType)
			} else {
				uu, err = r.Handlers.UpdateUser(ctx, r, v["id"].(string), v, authType)
			}

			changes.Uu = append(changes.Uu, uu)
		}

		event.AddNewValue("uu", changes.Uu)
		item.Uu = changes.Uu
		newItem.Uu = changes.Uu
	}

	if ids, exists := input["uuIds"]; exists {
		if err := tx.Unscoped().Model(&User{}).Where("tt_id = ?", id).Update("tt_id", "").Error; err != nil {
			tx.Rollback()
			return item, err
		}
		for _, v := range ids.([]interface{}) {
			uuIds = append(uuIds, v.(string))
		}

		if len(uuIds) > 0 {
			isChange = true
		}
	}

	if _, ok := input["uuu"]; ok {
		var uuu *User
		uuuInput := input["uuu"].(map[string]interface{})

		if uuuInput["id"] == nil {
			uuu, err = r.Handlers.CreateUser(ctx, r, uuuInput, authType)
		} else {
			uuu, err = r.Handlers.UpdateUser(ctx, r, uuuInput["id"].(string), uuuInput, authType)
		}

		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf(fmt.Sprintf("uuu %s", err.Error()))
		}

		event.AddOldValue("uuu", item.Uuu)
		event.AddNewValue("uuu", changes.Uuu)
		item.Uuu = uuu
		newItem.UuuID = &uuu.ID
		isChange = true
	}

	var uuuuIds []string

	if _, ok := input["uuuu"]; ok {
		if err := tx.Unscoped().Model(&User{}).Where("tttt_id = ?", id).Update("tttt_id", "").Error; err != nil {
			tx.Rollback()
			return item, err
		}

		var uuuuMaps []map[string]interface{}
		for _, v := range input["uuuu"].([]interface{}) {
			vMaps := v.(map[string]interface{})
			uuuuMaps = append(uuuuMaps, vMaps)
		}

		for _, v := range uuuuMaps {
			var uuuu *User
			v["ttttId"] = id
			if v["id"] == nil {
				uuuu, err = r.Handlers.CreateUser(ctx, r, v, authType)
			} else {
				uuuu, err = r.Handlers.UpdateUser(ctx, r, v["id"].(string), v, authType)
			}

			changes.Uuuu = append(changes.Uuuu, uuuu)
		}

		event.AddNewValue("uuuu", changes.Uuuu)
		item.Uuuu = changes.Uuuu
		newItem.Uuuu = changes.Uuuu
	}

	if ids, exists := input["uuuuIds"]; exists {
		if err := tx.Unscoped().Model(&User{}).Where("tttt_id = ?", id).Update("tttt_id", "").Error; err != nil {
			tx.Rollback()
			return item, err
		}
		for _, v := range ids.([]interface{}) {
			uuuuIds = append(uuuuIds, v.(string))
		}

		if len(uuuuIds) > 0 {
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

	if _, ok := input["title"]; ok && (item.Title != changes.Title) && (item.Title == nil || changes.Title == nil || *item.Title != *changes.Title) {
		event.AddOldValue("title", item.Title)
		event.AddNewValue("title", changes.Title)
		item.Title = changes.Title
		newItem.Title = changes.Title
		isChange = true
	}

	if _, ok := input["uId"]; ok && (item.UID != changes.UID) && (item.UID == nil || changes.UID == nil || *item.UID != *changes.UID) {

		// if err := tx.Select("id").Where("id", input["uId"]).First(&U{}).Error; err != nil {
		// 	tx.Rollback()
		// 	return nil, fmt.Errorf("uId " + err.Error())
		// }
		event.AddOldValue("uId", item.UID)
		event.AddNewValue("uId", changes.UID)
		item.UID = changes.UID
		newItem.UID = changes.UID
		isChange = true
	}

	if _, ok := input["uuuId"]; ok && (item.UuuID != changes.UuuID) && (item.UuuID == nil || changes.UuuID == nil || *item.UuuID != *changes.UuuID) {

		// if err := tx.Select("id").Where("id", input["uuuId"]).First(&Uuu{}).Error; err != nil {
		// 	tx.Rollback()
		// 	return nil, fmt.Errorf("uuuId " + err.Error())
		// }
		event.AddOldValue("uuuId", item.UuuID)
		event.AddNewValue("uuuId", changes.UuuID)
		item.UuuID = changes.UuuID
		newItem.UuuID = changes.UuuID
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
		tx.Rollback()
		return nil, err
	}

	if !isChange {
		return item, nil
	}

	if err := tx.Model(&newItem).Where("id = ?", id).Omit(clause.Associations).Updates(newItem).Error; err != nil {
		tx.Rollback()
		return item, err
	}

	if len(uuIds) > 0 {
		if err := tx.Model(&User{}).Where("id IN(?)", uuIds).Update("tt_id", item.ID).Error; err != nil {
			tx.Rollback()
			return item, err
		}
	}

	if len(uuuuIds) > 0 {
		if err := tx.Model(&User{}).Where("id IN(?)", uuuuIds).Update("tttt_id", item.ID).Error; err != nil {
			tx.Rollback()
			return item, err
		}
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
		if unscoped != nil && *unscoped == true {
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
