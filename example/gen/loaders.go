package gen

import (
	"context"
	"errors"

	"github.com/graph-gophers/dataloader"
	"gorm.io/gorm"
)

func GetLoaders(db *DB) map[string]*dataloader.Loader {
	loaders := map[string]*dataloader.Loader{}

	usersTBatchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result

		ids := make([]string, len(keys))
		for i, key := range keys {
			ids[i] = key.String()
		}

		items := &[]User{}
		selects := GetFieldsRequested(ctx, "users")
		if IndexOf(selects, "users"+".id") == -1 {
			selects = append(selects, "users"+".id")
		}

		if IndexOf(selects, "users"+".t_id") == -1 {
			selects = append(selects, "users"+".t_id")
		}

		res := db.Query().Select(selects).Find(items, "t_id IN (?)", ids)
		if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return []*dataloader.Result{
				{Error: res.Error},
			}
		}

		itemMap := make(map[string][]*User, len(keys))
		for _, v := range *items {
			item := v
			if itemMap[*item.TID] == nil {
				itemMap[*item.TID] = []*User{}
			}
			itemMap[*item.TID] = append(itemMap[*item.TID], &item)
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
	loaders["UserT"] = dataloader.NewBatchedLoader(usersTBatchFn, dataloader.WithClearCacheOnBatch())

	usersTtBatchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result

		ids := make([]string, len(keys))
		for i, key := range keys {
			ids[i] = key.String()
		}

		items := &[]User{}
		selects := GetFieldsRequested(ctx, "users")
		if IndexOf(selects, "users"+".id") == -1 {
			selects = append(selects, "users"+".id")
		}

		if IndexOf(selects, "users"+".tt_id") == -1 {
			selects = append(selects, "users"+".tt_id")
		}

		res := db.Query().Select(selects).Find(items, "tt_id IN (?)", ids)
		if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return []*dataloader.Result{
				{Error: res.Error},
			}
		}

		itemMap := make(map[string][]*User, len(keys))
		for _, v := range *items {
			item := v
			if itemMap[*item.TtID] == nil {
				itemMap[*item.TtID] = []*User{}
			}
			itemMap[*item.TtID] = append(itemMap[*item.TtID], &item)
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
	loaders["UserTt"] = dataloader.NewBatchedLoader(usersTtBatchFn, dataloader.WithClearCacheOnBatch())

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

		res := db.Query().Select(selects).Find(items, "id IN (?)", ids)
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

	tasksUBatchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
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

		if IndexOf(selects, "tasks"+".u_id") == -1 {
			selects = append(selects, "tasks"+".u_id")
		}

		res := db.Query().Select(selects).Find(items, "u_id IN (?)", ids)
		if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return []*dataloader.Result{
				{Error: res.Error},
			}
		}

		itemMap := make(map[string][]*Task, len(keys))
		for _, v := range *items {
			item := v
			if itemMap[*item.UID] == nil {
				itemMap[*item.UID] = []*Task{}
			}
			itemMap[*item.UID] = append(itemMap[*item.UID], &item)
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
	loaders["TaskU"] = dataloader.NewBatchedLoader(tasksUBatchFn, dataloader.WithClearCacheOnBatch())

	tasksUuuBatchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
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

		if IndexOf(selects, "tasks"+".uuu_id") == -1 {
			selects = append(selects, "tasks"+".uuu_id")
		}

		res := db.Query().Select(selects).Find(items, "uuu_id IN (?)", ids)
		if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return []*dataloader.Result{
				{Error: res.Error},
			}
		}

		itemMap := make(map[string][]*Task, len(keys))
		for _, v := range *items {
			item := v
			if itemMap[*item.UuuID] == nil {
				itemMap[*item.UuuID] = []*Task{}
			}
			itemMap[*item.UuuID] = append(itemMap[*item.UuuID], &item)
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
	loaders["TaskUuu"] = dataloader.NewBatchedLoader(tasksUuuBatchFn, dataloader.WithClearCacheOnBatch())

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

		res := db.Query().Select(selects).Find(items, "id IN (?)", ids)
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
