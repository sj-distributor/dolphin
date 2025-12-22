package gen

import (
	"context"
	"errors"

	"github.com/graph-gophers/dataloader"
	"gorm.io/gorm"
)

func GetLoaders(db *DB) map[string]*dataloader.Loader {
	loaders := map[string]*dataloader.Loader{}

	usersProfileBatchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result

		ids := make([]string, len(keys))
		for i, key := range keys {
			ids[i] = key.String()
		}

		items := &[]User{}
		selects := GetFieldsRequested(ctx, TableName("users", ctx))

		if IndexOf(selects, TableName("users", ctx)+".*") == -1 {
			if IndexOf(selects, TableName("users", ctx)+".id") == -1 {
				selects = append(selects, "users"+".id")
			}

			if IndexOf(selects, TableName("users", ctx)+".profile_id") == -1 {
				selects = append(selects, TableName("users", ctx)+".profile_id")
			}
		}

		res := db.Query().Table(TableName("users", ctx)).Select(selects).Order("weight ASC, created_at ASC").Find(items, "profile_id IN (?)", ids)
		if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return []*dataloader.Result{
				{Error: res.Error},
			}
		}

		itemMap := make(map[string][]*User, len(keys))
		for _, v := range *items {
			item := v

			mapKey := *item.ProfileID

			if itemMap[mapKey] == nil {
				itemMap[mapKey] = []*User{}
			}
			itemMap[mapKey] = append(itemMap[mapKey], &item)
		}

		for _, key := range keys {
			id := key.String()
			item, ok := itemMap[id]
			if !ok {
				results = append(results, &dataloader.Result{
					Data:  nil,
					Error: nil,
					// Error: fmt.Errorf("User with id '%s' not found", id),
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
	loaders["UserProfile"] = dataloader.NewBatchedLoader(usersProfileBatchFn, dataloader.WithClearCacheOnBatch())
	loaders["ProfileAndUserIds"] = dataloader.NewBatchedLoader(usersProfileBatchFn, dataloader.WithClearCacheOnBatch())

	usersBatchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result

		ids := make([]string, len(keys))
		for i, key := range keys {
			ids[i] = key.String()
		}

		items := &[]User{}
		selects := GetFieldsRequested(ctx, TableName("users", ctx))
		if len(selects) > 0 && IndexOf(selects, TableName("users", ctx)+".*") == -1 && IndexOf(selects, TableName("users", ctx)+".id") == -1 {
			selects = append(selects, TableName("users", ctx)+".id")
		}

		res := db.Query().Table(TableName("users", ctx)).Select(selects).Order("weight ASC, created_at ASC").Find(items, "id IN (?)", ids)
		if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return []*dataloader.Result{
				{Error: res.Error},
			}
		}

		itemMap := make(map[string]User, len(keys))
		for _, item := range *items {
			itemMap[item.ID] = item
		}

		for _, key := range keys {
			id := key.String()
			item, ok := itemMap[id]
			if !ok {
				results = append(results, &dataloader.Result{
					Data:  nil,
					Error: nil,
					// Error: fmt.Errorf("User with id '%s' not found", id),
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

	loaders["User"] = dataloader.NewBatchedLoader(usersBatchFn, dataloader.WithClearCacheOnBatch())

	profilesUserBatchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result

		ids := make([]string, len(keys))
		for i, key := range keys {
			ids[i] = key.String()
		}

		items := &[]Profile{}
		selects := GetFieldsRequested(ctx, TableName("profiles", ctx))

		if IndexOf(selects, TableName("profiles", ctx)+".*") == -1 {
			if IndexOf(selects, TableName("profiles", ctx)+".id") == -1 {
				selects = append(selects, "profiles"+".id")
			}

			if IndexOf(selects, TableName("profiles", ctx)+".user_id") == -1 {
				selects = append(selects, TableName("profiles", ctx)+".user_id")
			}
		}

		res := db.Query().Table(TableName("profiles", ctx)).Select(selects).Order("weight ASC, created_at ASC").Find(items, "user_id IN (?)", ids)
		if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return []*dataloader.Result{
				{Error: res.Error},
			}
		}

		itemMap := make(map[string][]*Profile, len(keys))
		for _, v := range *items {
			item := v

			mapKey := item.UserID

			if itemMap[mapKey] == nil {
				itemMap[mapKey] = []*Profile{}
			}
			itemMap[mapKey] = append(itemMap[mapKey], &item)
		}

		for _, key := range keys {
			id := key.String()
			item, ok := itemMap[id]
			if !ok {
				results = append(results, &dataloader.Result{
					Data:  nil,
					Error: nil,
					// Error: fmt.Errorf("Profile with id '%s' not found", id),
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
	loaders["ProfileUser"] = dataloader.NewBatchedLoader(profilesUserBatchFn, dataloader.WithClearCacheOnBatch())
	loaders["UserAndProfileIds"] = dataloader.NewBatchedLoader(profilesUserBatchFn, dataloader.WithClearCacheOnBatch())

	profilesBatchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result

		ids := make([]string, len(keys))
		for i, key := range keys {
			ids[i] = key.String()
		}

		items := &[]Profile{}
		selects := GetFieldsRequested(ctx, TableName("profiles", ctx))
		if len(selects) > 0 && IndexOf(selects, TableName("profiles", ctx)+".*") == -1 && IndexOf(selects, TableName("profiles", ctx)+".id") == -1 {
			selects = append(selects, TableName("profiles", ctx)+".id")
		}

		res := db.Query().Table(TableName("profiles", ctx)).Select(selects).Order("weight ASC, created_at ASC").Find(items, "id IN (?)", ids)
		if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return []*dataloader.Result{
				{Error: res.Error},
			}
		}

		itemMap := make(map[string]Profile, len(keys))
		for _, item := range *items {
			itemMap[item.ID] = item
		}

		for _, key := range keys {
			id := key.String()
			item, ok := itemMap[id]
			if !ok {
				results = append(results, &dataloader.Result{
					Data:  nil,
					Error: nil,
					// Error: fmt.Errorf("Profile with id '%s' not found", id),
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

	loaders["Profile"] = dataloader.NewBatchedLoader(profilesBatchFn, dataloader.WithClearCacheOnBatch())

	tasksUserBatchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result

		ids := make([]string, len(keys))
		for i, key := range keys {
			ids[i] = key.String()
		}

		items := &[]Task{}
		selects := GetFieldsRequested(ctx, TableName("tasks", ctx))

		if IndexOf(selects, TableName("tasks", ctx)+".*") == -1 {
			if IndexOf(selects, TableName("tasks", ctx)+".id") == -1 {
				selects = append(selects, "tasks"+".id")
			}

			if IndexOf(selects, TableName("tasks", ctx)+".user_id") == -1 {
				selects = append(selects, TableName("tasks", ctx)+".user_id")
			}
		}

		res := db.Query().Table(TableName("tasks", ctx)).Select(selects).Order("weight ASC, created_at ASC").Find(items, "user_id IN (?)", ids)
		if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return []*dataloader.Result{
				{Error: res.Error},
			}
		}

		itemMap := make(map[string][]*Task, len(keys))
		for _, v := range *items {
			item := v

			mapKey := *item.UserID

			if itemMap[mapKey] == nil {
				itemMap[mapKey] = []*Task{}
			}
			itemMap[mapKey] = append(itemMap[mapKey], &item)
		}

		for _, key := range keys {
			id := key.String()
			item, ok := itemMap[id]
			if !ok {
				results = append(results, &dataloader.Result{
					Data:  nil,
					Error: nil,
					// Error: fmt.Errorf("Task with id '%s' not found", id),
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
	loaders["TaskUser"] = dataloader.NewBatchedLoader(tasksUserBatchFn, dataloader.WithClearCacheOnBatch())
	loaders["UserAndTaskIds"] = dataloader.NewBatchedLoader(tasksUserBatchFn, dataloader.WithClearCacheOnBatch())

	tasksBatchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result

		ids := make([]string, len(keys))
		for i, key := range keys {
			ids[i] = key.String()
		}

		items := &[]Task{}
		selects := GetFieldsRequested(ctx, TableName("tasks", ctx))
		if len(selects) > 0 && IndexOf(selects, TableName("tasks", ctx)+".*") == -1 && IndexOf(selects, TableName("tasks", ctx)+".id") == -1 {
			selects = append(selects, TableName("tasks", ctx)+".id")
		}

		res := db.Query().Table(TableName("tasks", ctx)).Select(selects).Order("weight ASC, created_at ASC").Find(items, "id IN (?)", ids)
		if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return []*dataloader.Result{
				{Error: res.Error},
			}
		}

		itemMap := make(map[string]Task, len(keys))
		for _, item := range *items {
			itemMap[item.ID] = item
		}

		for _, key := range keys {
			id := key.String()
			item, ok := itemMap[id]
			if !ok {
				results = append(results, &dataloader.Result{
					Data:  nil,
					Error: nil,
					// Error: fmt.Errorf("Task with id '%s' not found", id),
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

	loaders["Task"] = dataloader.NewBatchedLoader(tasksBatchFn, dataloader.WithClearCacheOnBatch())

	user_rolesBatchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result

		ids := make([]string, len(keys))
		for i, key := range keys {
			ids[i] = key.String()
		}

		items := &[]UserRole{}
		selects := GetFieldsRequested(ctx, TableName("user_roles", ctx))
		if len(selects) > 0 && IndexOf(selects, TableName("user_roles", ctx)+".*") == -1 && IndexOf(selects, TableName("user_roles", ctx)+".id") == -1 {
			selects = append(selects, TableName("user_roles", ctx)+".id")
		}

		res := db.Query().Table(TableName("user_roles", ctx)).Select(selects).Order("weight ASC, created_at ASC").Find(items, "id IN (?)", ids)
		if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return []*dataloader.Result{
				{Error: res.Error},
			}
		}

		itemMap := make(map[string]UserRole, len(keys))
		for _, item := range *items {
			itemMap[item.ID] = item
		}

		for _, key := range keys {
			id := key.String()
			item, ok := itemMap[id]
			if !ok {
				results = append(results, &dataloader.Result{
					Data:  nil,
					Error: nil,
					// Error: fmt.Errorf("UserRole with id '%s' not found", id),
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

	loaders["UserRole"] = dataloader.NewBatchedLoader(user_rolesBatchFn, dataloader.WithClearCacheOnBatch())

	tagsBatchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result

		ids := make([]string, len(keys))
		for i, key := range keys {
			ids[i] = key.String()
		}

		items := &[]Tag{}
		selects := GetFieldsRequested(ctx, TableName("tags", ctx))
		if len(selects) > 0 && IndexOf(selects, TableName("tags", ctx)+".*") == -1 && IndexOf(selects, TableName("tags", ctx)+".id") == -1 {
			selects = append(selects, TableName("tags", ctx)+".id")
		}

		res := db.Query().Table(TableName("tags", ctx)).Select(selects).Order("weight ASC, created_at ASC").Find(items, "id IN (?)", ids)
		if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return []*dataloader.Result{
				{Error: res.Error},
			}
		}

		itemMap := make(map[string]Tag, len(keys))
		for _, item := range *items {
			itemMap[item.ID] = item
		}

		for _, key := range keys {
			id := key.String()
			item, ok := itemMap[id]
			if !ok {
				results = append(results, &dataloader.Result{
					Data:  nil,
					Error: nil,
					// Error: fmt.Errorf("Tag with id '%s' not found", id),
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

	loaders["Tag"] = dataloader.NewBatchedLoader(tagsBatchFn, dataloader.WithClearCacheOnBatch())

	return loaders
}
