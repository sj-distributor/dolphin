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

	accountsOwnerBatchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result

		ids := make([]string, len(keys))
		for i, key := range keys {
			ids[i] = key.String()
		}

		items := &[]Account{}
		selects := GetFieldsRequested(ctx, "accounts")
		if IndexOf(selects, "accounts"+".id") == -1 {
			selects = append(selects, "accounts"+".id")
		}

		if IndexOf(selects, "accounts"+".owner_id") == -1 {
			selects = append(selects, "accounts"+".owner_id")
		}

		res := db.Query().Select(selects).Find(items, "owner_id IN (?)", ids)
		if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return []*dataloader.Result{
				{Error: res.Error},
			}
		}

		itemMap := make(map[string][]*Account, len(keys))
		for _, v := range *items {
			item := v
			if itemMap[*item.OwnerID] == nil {
				itemMap[*item.OwnerID] = []*Account{}
			}
			itemMap[*item.OwnerID] = append(itemMap[*item.OwnerID], &item)
		}

		for _, key := range keys {
			id := key.String()
			item, ok := itemMap[id]
			if !ok {
				results = append(results, &dataloader.Result{
					Data:  nil,
					Error: nil,
					// Error: fmt.Errorf("Account with id '%s' not found", id),
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
	loaders["AccountOwner"] = dataloader.NewBatchedLoader(accountsOwnerBatchFn, dataloader.WithClearCacheOnBatch())

	accountsBatchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result

		ids := make([]string, len(keys))
		for i, key := range keys {
			ids[i] = key.String()
		}

		items := &[]Account{}
		selects := GetFieldsRequested(ctx, "accounts")
		if len(selects) > 0 && IndexOf(selects, "accounts"+".id") == -1 {
			selects = append(selects, "accounts"+".id")
		}

		res := db.Query().Select(selects).Find(items, "id IN (?)", ids)
		if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return []*dataloader.Result{
				{Error: res.Error},
			}
		}

		itemMap := make(map[string]Account, len(keys))
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
					// Error: fmt.Errorf("Account with id '%s' not found", id),
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

	loaders["Account"] = dataloader.NewBatchedLoader(accountsBatchFn, dataloader.WithClearCacheOnBatch())

	transactionsAccountBatchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result

		ids := make([]string, len(keys))
		for i, key := range keys {
			ids[i] = key.String()
		}

		items := &[]Transaction{}
		selects := GetFieldsRequested(ctx, "transactions")
		if IndexOf(selects, "transactions"+".id") == -1 {
			selects = append(selects, "transactions"+".id")
		}

		if IndexOf(selects, "transactions"+".account_id") == -1 {
			selects = append(selects, "transactions"+".account_id")
		}

		res := db.Query().Select(selects).Find(items, "account_id IN (?)", ids)
		if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return []*dataloader.Result{
				{Error: res.Error},
			}
		}

		itemMap := make(map[string][]*Transaction, len(keys))
		for _, v := range *items {
			item := v
			if itemMap[*item.AccountID] == nil {
				itemMap[*item.AccountID] = []*Transaction{}
			}
			itemMap[*item.AccountID] = append(itemMap[*item.AccountID], &item)
		}

		for _, key := range keys {
			id := key.String()
			item, ok := itemMap[id]
			if !ok {
				results = append(results, &dataloader.Result{
					Data:  nil,
					Error: nil,
					// Error: fmt.Errorf("Transaction with id '%s' not found", id),
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
	loaders["TransactionAccount"] = dataloader.NewBatchedLoader(transactionsAccountBatchFn, dataloader.WithClearCacheOnBatch())

	transactionsBatchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result

		ids := make([]string, len(keys))
		for i, key := range keys {
			ids[i] = key.String()
		}

		items := &[]Transaction{}
		selects := GetFieldsRequested(ctx, "transactions")
		if len(selects) > 0 && IndexOf(selects, "transactions"+".id") == -1 {
			selects = append(selects, "transactions"+".id")
		}

		res := db.Query().Select(selects).Find(items, "id IN (?)", ids)
		if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return []*dataloader.Result{
				{Error: res.Error},
			}
		}

		itemMap := make(map[string]Transaction, len(keys))
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
					// Error: fmt.Errorf("Transaction with id '%s' not found", id),
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

	loaders["Transaction"] = dataloader.NewBatchedLoader(transactionsBatchFn, dataloader.WithClearCacheOnBatch())

	todosAccountBatchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result

		ids := make([]string, len(keys))
		for i, key := range keys {
			ids[i] = key.String()
		}

		items := &[]Todo{}
		selects := GetFieldsRequested(ctx, "todos")
		if IndexOf(selects, "todos"+".id") == -1 {
			selects = append(selects, "todos"+".id")
		}

		if IndexOf(selects, "todos"+".account_id") == -1 {
			selects = append(selects, "todos"+".account_id")
		}

		res := db.Query().Select(selects).Find(items, "account_id IN (?)", ids)
		if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return []*dataloader.Result{
				{Error: res.Error},
			}
		}

		itemMap := make(map[string][]*Todo, len(keys))
		for _, v := range *items {
			item := v
			if itemMap[*item.AccountID] == nil {
				itemMap[*item.AccountID] = []*Todo{}
			}
			itemMap[*item.AccountID] = append(itemMap[*item.AccountID], &item)
		}

		for _, key := range keys {
			id := key.String()
			item, ok := itemMap[id]
			if !ok {
				results = append(results, &dataloader.Result{
					Data:  nil,
					Error: nil,
					// Error: fmt.Errorf("Todo with id '%s' not found", id),
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
	loaders["TodoAccount"] = dataloader.NewBatchedLoader(todosAccountBatchFn, dataloader.WithClearCacheOnBatch())

	todosBatchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result

		ids := make([]string, len(keys))
		for i, key := range keys {
			ids[i] = key.String()
		}

		items := &[]Todo{}
		selects := GetFieldsRequested(ctx, "todos")
		if len(selects) > 0 && IndexOf(selects, "todos"+".id") == -1 {
			selects = append(selects, "todos"+".id")
		}

		res := db.Query().Select(selects).Find(items, "id IN (?)", ids)
		if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return []*dataloader.Result{
				{Error: res.Error},
			}
		}

		itemMap := make(map[string]Todo, len(keys))
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
					// Error: fmt.Errorf("Todo with id '%s' not found", id),
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

	loaders["Todo"] = dataloader.NewBatchedLoader(todosBatchFn, dataloader.WithClearCacheOnBatch())

	return loaders
}
