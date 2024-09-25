package gen

import (
	"context"
	"errors"

	"github.com/graph-gophers/dataloader"
	"gorm.io/gorm"
)

func GetLoaders(db *DB) map[string]*dataloader.Loader {
	loaders := map[string]*dataloader.Loader{}

	usersBatchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result

		ids := make([]string, len(keys))
		for i, key := range keys {
			ids[i] = key.String()
		}

		items := &[]User{}
		selects := GetFieldsRequested(ctx, "users")
		if len(selects) > 0 && IndexOf(selects, "users"+".id") == -1 {
			selects = append(selects, "users"+".id")
		}

		res := db.Query().Select(selects).Order("weight ASC, created_at ASC").Find(items, "id IN (?)", ids)
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

	tasksUserBatchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result

		ids := make([]string, len(keys))
		for i, key := range keys {
			ids[i] = key.String()
		}

		items := &[]Task{}
		selects := GetFieldsRequested(ctx, "tasks")
		if IndexOf(selects, "tasks"+".id") == -1 {
			selects = append(selects, "tasks"+".id")
		}

		if IndexOf(selects, "tasks"+".user_id") == -1 {
			selects = append(selects, "tasks"+".user_id")
		}

		res := db.Query().Select(selects).Order("weight ASC, created_at ASC").Find(items, "user_id IN (?)", ids)
		if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return []*dataloader.Result{
				{Error: res.Error},
			}
		}

		itemMap := make(map[string][]*Task, len(keys))
		for _, v := range *items {
			item := v
			if itemMap[*item.UserID] == nil {
				itemMap[*item.UserID] = []*Task{}
			}
			itemMap[*item.UserID] = append(itemMap[*item.UserID], &item)
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
		selects := GetFieldsRequested(ctx, "tasks")
		if len(selects) > 0 && IndexOf(selects, "tasks"+".id") == -1 {
			selects = append(selects, "tasks"+".id")
		}

		res := db.Query().Select(selects).Order("weight ASC, created_at ASC").Find(items, "id IN (?)", ids)
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

	return loaders
}
