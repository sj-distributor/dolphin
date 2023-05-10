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

	if input["accounts"] != nil && input["accountsId"] != nil {
		tx.Rollback()
		return nil, fmt.Errorf("accountsId and accounts cannot coexist")
	}

	var accountsIds []string

	if _, ok := input["accounts"]; ok {
		var accountsMaps []map[string]interface{}
		for _, v := range input["accounts"].([]interface{}) {
			accountsMaps = append(accountsMaps, v.(map[string]interface{}))
		}

		for _, v := range accountsMaps {
			var accounts *Account
			if v["id"] == nil {
				accounts, err = r.Handlers.CreateAccount(ctx, r, v)
			} else {
				accounts, err = r.Handlers.UpdateAccount(ctx, r, v["id"].(string), v)
			}

			changes.Accounts = append(changes.Accounts, accounts)
			accountsIds = append(accountsIds, accounts.ID)
		}
		event.AddNewValue("accounts", changes.Accounts)
		item.Accounts = changes.Accounts
	}

	if input["todo"] != nil && input["todoId"] != nil {
		tx.Rollback()
		return nil, fmt.Errorf("todoId and todo cannot coexist")
	}

	var todoIds []string

	if _, ok := input["todo"]; ok {
		var todoMaps []map[string]interface{}
		for _, v := range input["todo"].([]interface{}) {
			todoMaps = append(todoMaps, v.(map[string]interface{}))
		}

		for _, v := range todoMaps {
			var todo *Todo
			if v["id"] == nil {
				todo, err = r.Handlers.CreateTodo(ctx, r, v)
			} else {
				todo, err = r.Handlers.UpdateTodo(ctx, r, v["id"].(string), v)
			}

			changes.Todo = append(changes.Todo, todo)
			todoIds = append(todoIds, todo.ID)
		}
		event.AddNewValue("todo", changes.Todo)
		item.Todo = changes.Todo
	}

	if _, ok := input["username"]; ok && (item.Username != changes.Username) {
		item.Username = changes.Username
		event.AddNewValue("username", changes.Username)
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

	if err := tx.Omit(clause.Associations).Create(item).Error; err != nil {
		tx.Rollback()
		return item, err
	}

	if len(accountsIds) > 0 {
		if err := tx.Model(&Account{}).Where("id IN(?)", accountsIds).Update("owner_id", item.ID).Error; err != nil {
			tx.Rollback()
			return item, err
		}
	}

	if len(todoIds) > 0 {
		if err := tx.Model(&Todo{}).Where("id IN(?)", todoIds).Update("account_id", item.ID).Error; err != nil {
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

	if input["accounts"] != nil && input["accountsId"] != nil {
		tx.Rollback()
		return nil, fmt.Errorf("accountsId and accounts cannot coexist")
	}

	if input["todo"] != nil && input["todoId"] != nil {
		tx.Rollback()
		return nil, fmt.Errorf("todoId and todo cannot coexist")
	}

	if err = GetItem(ctx, tx, TableName("users"), item, &id); err != nil {
		tx.Rollback()
		return nil, err
	}

	if item.UpdatedBy != nil && principalID != nil && *item.UpdatedBy != *principalID {
		newItem.UpdatedBy = principalID
	}

	if _, ok := input["accounts"]; ok {
		if err := tx.Unscoped().Model(&Account{}).Where("owner_id = ?", id).Update("owner_id", "").Error; err != nil {
			tx.Rollback()
			return item, err
		}

		var accountsMaps []map[string]interface{}
		for _, v := range input["accounts"].([]interface{}) {
			vMaps := v.(map[string]interface{})
			accountsMaps = append(accountsMaps, vMaps)
		}

		for _, v := range accountsMaps {
			accounts := &Account{}
			v["ownerId"] = id
			if v["id"] == nil {
				accounts, err = r.Handlers.CreateAccount(ctx, r, v)
			} else {
				accounts, err = r.Handlers.UpdateAccount(ctx, r, v["id"].(string), v)
			}

			changes.Accounts = append(changes.Accounts, accounts)
		}

		event.AddNewValue("accounts", changes.Accounts)
		item.Accounts = changes.Accounts
		newItem.Accounts = changes.Accounts
	}

	if _, ok := input["todo"]; ok {
		if err := tx.Unscoped().Model(&Todo{}).Where("account_id = ?", id).Update("account_id", "").Error; err != nil {
			tx.Rollback()
			return item, err
		}

		var todoMaps []map[string]interface{}
		for _, v := range input["todo"].([]interface{}) {
			vMaps := v.(map[string]interface{})
			todoMaps = append(todoMaps, vMaps)
		}

		for _, v := range todoMaps {
			todo := &Todo{}
			v["accountId"] = id
			if v["id"] == nil {
				todo, err = r.Handlers.CreateTodo(ctx, r, v)
			} else {
				todo, err = r.Handlers.UpdateTodo(ctx, r, v["id"].(string), v)
			}

			changes.Todo = append(changes.Todo, todo)
		}

		event.AddNewValue("todo", changes.Todo)
		item.Todo = changes.Todo
		newItem.Todo = changes.Todo
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		event.AddOldValue("id", item.ID)
		event.AddNewValue("id", changes.ID)
		item.ID = changes.ID
		newItem.ID = changes.ID
		isChange = true
	}

	if _, ok := input["username"]; ok && (item.Username != changes.Username) {
		event.AddOldValue("username", item.Username)
		event.AddNewValue("username", changes.Username)
		item.Username = changes.Username
		newItem.Username = changes.Username
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

func (r *GeneratedMutationResolver) CreateAccount(ctx context.Context, input map[string]interface{}) (item *Account, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreateAccount(ctx, r.GeneratedResolver, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func CreateAccountHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Account, err error) {
	item = &Account{}

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
		Entity:      "Account",
		EntityID:    item.ID,
		Date:        milliTime,
		PrincipalID: principalID,
	})

	var changes AccountChanges
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

	if input["owner"] != nil && input["ownerId"] != nil {
		tx.Rollback()
		return nil, fmt.Errorf("ownerId and owner cannot coexist")
	}

	if _, ok := input["owner"]; ok {
		var owner *User
		ownerInput := input["owner"].(map[string]interface{})

		if ownerInput["id"] == nil {
			owner, err = r.Handlers.CreateUser(ctx, r, ownerInput)
		} else {
			owner, err = r.Handlers.UpdateUser(ctx, r, ownerInput["id"].(string), ownerInput)
		}

		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf(fmt.Sprintf("owner %s", err.Error()))
		}

		event.AddNewValue("owner", changes.Owner)
		item.Owner = owner
		item.OwnerID = &owner.ID
	}

	if input["transactions"] != nil && input["transactionsId"] != nil {
		tx.Rollback()
		return nil, fmt.Errorf("transactionsId and transactions cannot coexist")
	}

	var transactionsIds []string

	if _, ok := input["transactions"]; ok {
		var transactionsMaps []map[string]interface{}
		for _, v := range input["transactions"].([]interface{}) {
			transactionsMaps = append(transactionsMaps, v.(map[string]interface{}))
		}

		for _, v := range transactionsMaps {
			var transactions *Transaction
			if v["id"] == nil {
				transactions, err = r.Handlers.CreateTransaction(ctx, r, v)
			} else {
				transactions, err = r.Handlers.UpdateTransaction(ctx, r, v["id"].(string), v)
			}

			changes.Transactions = append(changes.Transactions, transactions)
			transactionsIds = append(transactionsIds, transactions.ID)
		}
		event.AddNewValue("transactions", changes.Transactions)
		item.Transactions = changes.Transactions
	}

	if _, ok := input["name"]; ok && (item.Name != changes.Name) {
		item.Name = changes.Name
		event.AddNewValue("name", changes.Name)
	}

	if _, ok := input["balance"]; ok && (item.Balance != changes.Balance) {
		item.Balance = changes.Balance
		event.AddNewValue("balance", changes.Balance)
	}

	if _, ok := input["ownerId"]; ok && (item.OwnerID != changes.OwnerID) && (item.OwnerID == nil || changes.OwnerID == nil || *item.OwnerID != *changes.OwnerID) {

		// if err := tx.Select("id").Where("id", input["ownerId"]).First(&Owner{}).Error; err != nil {
		// 	tx.Rollback()
		// 	return nil, fmt.Errorf("ownerId " + err.Error())
		// }
		item.OwnerID = changes.OwnerID
		event.AddNewValue("ownerId", changes.OwnerID)
	}

	if err := tx.Omit(clause.Associations).Create(item).Error; err != nil {
		tx.Rollback()
		return item, err
	}

	if len(transactionsIds) > 0 {
		if err := tx.Model(&Transaction{}).Where("id IN(?)", transactionsIds).Update("account_id", item.ID).Error; err != nil {
			tx.Rollback()
			return item, err
		}
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}
func (r *GeneratedMutationResolver) UpdateAccount(ctx context.Context, id string, input map[string]interface{}) (item *Account, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdateAccount(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func UpdateAccountHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Account, err error) {
	item = &Account{}
	newItem := &Account{}

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
		Entity:      "Account",
		EntityID:    id,
		Date:        milliTime,
		PrincipalID: principalID,
	})

	var changes AccountChanges
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

	if input["owner"] != nil && input["ownerId"] != nil {
		tx.Rollback()
		return nil, fmt.Errorf("ownerId and owner cannot coexist")
	}

	if input["transactions"] != nil && input["transactionsId"] != nil {
		tx.Rollback()
		return nil, fmt.Errorf("transactionsId and transactions cannot coexist")
	}

	if err = GetItem(ctx, tx, TableName("accounts"), item, &id); err != nil {
		tx.Rollback()
		return nil, err
	}

	if item.UpdatedBy != nil && principalID != nil && *item.UpdatedBy != *principalID {
		newItem.UpdatedBy = principalID
	}

	if _, ok := input["owner"]; ok {
		var owner *User
		ownerInput := input["owner"].(map[string]interface{})

		if ownerInput["id"] == nil {
			owner, err = r.Handlers.CreateUser(ctx, r, ownerInput)
		} else {
			owner, err = r.Handlers.UpdateUser(ctx, r, ownerInput["id"].(string), ownerInput)
		}

		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf(fmt.Sprintf("owner %s", err.Error()))
		}

		if err := tx.Model(&item).Association("Owner").Clear(); err != nil {
			tx.Rollback()
			return nil, err
		}

		if err := tx.Model(&User{}).Where("id = ?", owner.ID).Update("accounts_id", item.ID).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		event.AddOldValue("owner", item.Owner)
		event.AddNewValue("owner", changes.Owner)
		item.Owner = owner
		newItem.OwnerID = &owner.ID
		isChange = true
	}

	if _, ok := input["transactions"]; ok {
		if err := tx.Unscoped().Model(&Transaction{}).Where("account_id = ?", id).Update("account_id", "").Error; err != nil {
			tx.Rollback()
			return item, err
		}

		var transactionsMaps []map[string]interface{}
		for _, v := range input["transactions"].([]interface{}) {
			vMaps := v.(map[string]interface{})
			transactionsMaps = append(transactionsMaps, vMaps)
		}

		for _, v := range transactionsMaps {
			transactions := &Transaction{}
			v["accountId"] = id
			if v["id"] == nil {
				transactions, err = r.Handlers.CreateTransaction(ctx, r, v)
			} else {
				transactions, err = r.Handlers.UpdateTransaction(ctx, r, v["id"].(string), v)
			}

			changes.Transactions = append(changes.Transactions, transactions)
		}

		event.AddNewValue("transactions", changes.Transactions)
		item.Transactions = changes.Transactions
		newItem.Transactions = changes.Transactions
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

	if _, ok := input["balance"]; ok && (item.Balance != changes.Balance) {
		event.AddOldValue("balance", item.Balance)
		event.AddNewValue("balance", changes.Balance)
		item.Balance = changes.Balance
		newItem.Balance = changes.Balance
		isChange = true
	}

	if _, ok := input["ownerId"]; ok && (item.OwnerID != changes.OwnerID) && (item.OwnerID == nil || changes.OwnerID == nil || *item.OwnerID != *changes.OwnerID) {

		// if err := tx.Select("id").Where("id", input["ownerId"]).First(&Owner{}).Error; err != nil {
		// 	tx.Rollback()
		// 	return nil, fmt.Errorf("ownerId " + err.Error())
		// }
		event.AddOldValue("ownerId", item.OwnerID)
		event.AddNewValue("ownerId", changes.OwnerID)
		item.OwnerID = changes.OwnerID
		newItem.OwnerID = changes.OwnerID
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

func DeleteAccountFunc(ctx context.Context, r *GeneratedResolver, id string, tye string, unscoped *bool) (err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item := &Account{}
	now := time.Now()
	tx := GetTransaction(ctx)
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err = GetItem(ctx, tx, TableName("accounts"), item, &id); err != nil {
		tx.Rollback()
		return err
	}

	deletedAt := now.UnixNano() / 1e6

	event := NewEvent(EventMetadata{
		Type:        EventTypeDeleted,
		Entity:      "Account",
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
		} else if err := tx.Model(&item).Updates(Account{DeletedAt: &deletedAt, DeletedBy: principalID, UpdatedBy: principalID}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}

func (r *GeneratedMutationResolver) DeleteAccounts(ctx context.Context, id []string, unscoped *bool) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteAccounts(ctx, r.GeneratedResolver, id, unscoped)
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

func DeleteAccountsHandler(ctx context.Context, r *GeneratedResolver, id []string, unscoped *bool) (bool, error) {
	var err error = nil

	if len(id) > 0 {
		for _, v := range id {
			err = DeleteAccountFunc(ctx, r, v, "delete", unscoped)
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

func (r *GeneratedMutationResolver) RecoveryAccounts(ctx context.Context, id []string) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.RecoveryAccounts(ctx, r.GeneratedResolver, id)
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

func RecoveryAccountsHandler(ctx context.Context, r *GeneratedResolver, id []string) (bool, error) {
	var err error = nil

	var unscoped bool = false

	if len(id) > 0 {
		for _, v := range id {
			err = DeleteAccountFunc(ctx, r, v, "recovery", &unscoped)
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

func (r *GeneratedMutationResolver) CreateTransaction(ctx context.Context, input map[string]interface{}) (item *Transaction, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreateTransaction(ctx, r.GeneratedResolver, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func CreateTransactionHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Transaction, err error) {
	item = &Transaction{}

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
		Entity:      "Transaction",
		EntityID:    item.ID,
		Date:        milliTime,
		PrincipalID: principalID,
	})

	var changes TransactionChanges
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

	if input["account"] != nil && input["accountId"] != nil {
		tx.Rollback()
		return nil, fmt.Errorf("accountId and account cannot coexist")
	}

	if _, ok := input["account"]; ok {
		var account *Account
		accountInput := input["account"].(map[string]interface{})

		if accountInput["id"] == nil {
			account, err = r.Handlers.CreateAccount(ctx, r, accountInput)
		} else {
			account, err = r.Handlers.UpdateAccount(ctx, r, accountInput["id"].(string), accountInput)
		}

		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf(fmt.Sprintf("account %s", err.Error()))
		}

		event.AddNewValue("account", changes.Account)
		item.Account = account
		item.AccountID = &account.ID
	}

	if _, ok := input["amount"]; ok && (item.Amount != changes.Amount) {
		item.Amount = changes.Amount
		event.AddNewValue("amount", changes.Amount)
	}

	if _, ok := input["date"]; ok && (item.Date != changes.Date) {
		item.Date = changes.Date
		event.AddNewValue("date", changes.Date)
	}

	if _, ok := input["note"]; ok && (item.Note != changes.Note) && (item.Note == nil || changes.Note == nil || *item.Note != *changes.Note) {
		item.Note = changes.Note
		event.AddNewValue("note", changes.Note)
	}

	if _, ok := input["accountId"]; ok && (item.AccountID != changes.AccountID) && (item.AccountID == nil || changes.AccountID == nil || *item.AccountID != *changes.AccountID) {

		// if err := tx.Select("id").Where("id", input["accountId"]).First(&Account{}).Error; err != nil {
		// 	tx.Rollback()
		// 	return nil, fmt.Errorf("accountId " + err.Error())
		// }
		item.AccountID = changes.AccountID
		event.AddNewValue("accountId", changes.AccountID)
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
func (r *GeneratedMutationResolver) UpdateTransaction(ctx context.Context, id string, input map[string]interface{}) (item *Transaction, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdateTransaction(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func UpdateTransactionHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Transaction, err error) {
	item = &Transaction{}
	newItem := &Transaction{}

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
		Entity:      "Transaction",
		EntityID:    id,
		Date:        milliTime,
		PrincipalID: principalID,
	})

	var changes TransactionChanges
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

	if input["account"] != nil && input["accountId"] != nil {
		tx.Rollback()
		return nil, fmt.Errorf("accountId and account cannot coexist")
	}

	if err = GetItem(ctx, tx, TableName("transactions"), item, &id); err != nil {
		tx.Rollback()
		return nil, err
	}

	if item.UpdatedBy != nil && principalID != nil && *item.UpdatedBy != *principalID {
		newItem.UpdatedBy = principalID
	}

	if _, ok := input["account"]; ok {
		var account *Account
		accountInput := input["account"].(map[string]interface{})

		if accountInput["id"] == nil {
			account, err = r.Handlers.CreateAccount(ctx, r, accountInput)
		} else {
			account, err = r.Handlers.UpdateAccount(ctx, r, accountInput["id"].(string), accountInput)
		}

		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf(fmt.Sprintf("account %s", err.Error()))
		}

		if err := tx.Model(&item).Association("Account").Clear(); err != nil {
			tx.Rollback()
			return nil, err
		}

		if err := tx.Model(&Account{}).Where("id = ?", account.ID).Update("transactions_id", item.ID).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		event.AddOldValue("account", item.Account)
		event.AddNewValue("account", changes.Account)
		item.Account = account
		newItem.AccountID = &account.ID
		isChange = true
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		event.AddOldValue("id", item.ID)
		event.AddNewValue("id", changes.ID)
		item.ID = changes.ID
		newItem.ID = changes.ID
		isChange = true
	}

	if _, ok := input["amount"]; ok && (item.Amount != changes.Amount) {
		event.AddOldValue("amount", item.Amount)
		event.AddNewValue("amount", changes.Amount)
		item.Amount = changes.Amount
		newItem.Amount = changes.Amount
		isChange = true
	}

	if _, ok := input["date"]; ok && (item.Date != changes.Date) {
		event.AddOldValue("date", item.Date)
		event.AddNewValue("date", changes.Date)
		item.Date = changes.Date
		newItem.Date = changes.Date
		isChange = true
	}

	if _, ok := input["note"]; ok && (item.Note != changes.Note) && (item.Note == nil || changes.Note == nil || *item.Note != *changes.Note) {
		event.AddOldValue("note", item.Note)
		event.AddNewValue("note", changes.Note)
		item.Note = changes.Note
		newItem.Note = changes.Note
		isChange = true
	}

	if _, ok := input["accountId"]; ok && (item.AccountID != changes.AccountID) && (item.AccountID == nil || changes.AccountID == nil || *item.AccountID != *changes.AccountID) {

		// if err := tx.Select("id").Where("id", input["accountId"]).First(&Account{}).Error; err != nil {
		// 	tx.Rollback()
		// 	return nil, fmt.Errorf("accountId " + err.Error())
		// }
		event.AddOldValue("accountId", item.AccountID)
		event.AddNewValue("accountId", changes.AccountID)
		item.AccountID = changes.AccountID
		newItem.AccountID = changes.AccountID
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

func DeleteTransactionFunc(ctx context.Context, r *GeneratedResolver, id string, tye string, unscoped *bool) (err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item := &Transaction{}
	now := time.Now()
	tx := GetTransaction(ctx)
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err = GetItem(ctx, tx, TableName("transactions"), item, &id); err != nil {
		tx.Rollback()
		return err
	}

	deletedAt := now.UnixNano() / 1e6

	event := NewEvent(EventMetadata{
		Type:        EventTypeDeleted,
		Entity:      "Transaction",
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
		} else if err := tx.Model(&item).Updates(Transaction{DeletedAt: &deletedAt, DeletedBy: principalID, UpdatedBy: principalID}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}

func (r *GeneratedMutationResolver) DeleteTransactions(ctx context.Context, id []string, unscoped *bool) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteTransactions(ctx, r.GeneratedResolver, id, unscoped)
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

func DeleteTransactionsHandler(ctx context.Context, r *GeneratedResolver, id []string, unscoped *bool) (bool, error) {
	var err error = nil

	if len(id) > 0 {
		for _, v := range id {
			err = DeleteTransactionFunc(ctx, r, v, "delete", unscoped)
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

func (r *GeneratedMutationResolver) RecoveryTransactions(ctx context.Context, id []string) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.RecoveryTransactions(ctx, r.GeneratedResolver, id)
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

func RecoveryTransactionsHandler(ctx context.Context, r *GeneratedResolver, id []string) (bool, error) {
	var err error = nil

	var unscoped bool = false

	if len(id) > 0 {
		for _, v := range id {
			err = DeleteTransactionFunc(ctx, r, v, "recovery", &unscoped)
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

func (r *GeneratedMutationResolver) CreateTodo(ctx context.Context, input map[string]interface{}) (item *Todo, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreateTodo(ctx, r.GeneratedResolver, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func CreateTodoHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Todo, err error) {
	item = &Todo{}

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
		Entity:      "Todo",
		EntityID:    item.ID,
		Date:        milliTime,
		PrincipalID: principalID,
	})

	var changes TodoChanges
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

	if input["account"] != nil && input["accountId"] != nil {
		tx.Rollback()
		return nil, fmt.Errorf("accountId and account cannot coexist")
	}

	if _, ok := input["account"]; ok {
		var account *User
		accountInput := input["account"].(map[string]interface{})

		if accountInput["id"] == nil {
			account, err = r.Handlers.CreateUser(ctx, r, accountInput)
		} else {
			account, err = r.Handlers.UpdateUser(ctx, r, accountInput["id"].(string), accountInput)
		}

		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf(fmt.Sprintf("account %s", err.Error()))
		}

		event.AddNewValue("account", changes.Account)
		item.Account = account
		item.AccountID = &account.ID
	}

	if _, ok := input["name"]; ok && (item.Name != changes.Name) {
		item.Name = changes.Name
		event.AddNewValue("name", changes.Name)
	}

	if _, ok := input["accountId"]; ok && (item.AccountID != changes.AccountID) && (item.AccountID == nil || changes.AccountID == nil || *item.AccountID != *changes.AccountID) {

		// if err := tx.Select("id").Where("id", input["accountId"]).First(&Account{}).Error; err != nil {
		// 	tx.Rollback()
		// 	return nil, fmt.Errorf("accountId " + err.Error())
		// }
		item.AccountID = changes.AccountID
		event.AddNewValue("accountId", changes.AccountID)
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
func (r *GeneratedMutationResolver) UpdateTodo(ctx context.Context, id string, input map[string]interface{}) (item *Todo, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdateTodo(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func UpdateTodoHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Todo, err error) {
	item = &Todo{}
	newItem := &Todo{}

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
		Entity:      "Todo",
		EntityID:    id,
		Date:        milliTime,
		PrincipalID: principalID,
	})

	var changes TodoChanges
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

	if input["account"] != nil && input["accountId"] != nil {
		tx.Rollback()
		return nil, fmt.Errorf("accountId and account cannot coexist")
	}

	if err = GetItem(ctx, tx, TableName("todos"), item, &id); err != nil {
		tx.Rollback()
		return nil, err
	}

	if item.UpdatedBy != nil && principalID != nil && *item.UpdatedBy != *principalID {
		newItem.UpdatedBy = principalID
	}

	if _, ok := input["account"]; ok {
		var account *User
		accountInput := input["account"].(map[string]interface{})

		if accountInput["id"] == nil {
			account, err = r.Handlers.CreateUser(ctx, r, accountInput)
		} else {
			account, err = r.Handlers.UpdateUser(ctx, r, accountInput["id"].(string), accountInput)
		}

		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf(fmt.Sprintf("account %s", err.Error()))
		}

		if err := tx.Model(&item).Association("Account").Clear(); err != nil {
			tx.Rollback()
			return nil, err
		}

		if err := tx.Model(&User{}).Where("id = ?", account.ID).Update("todo_id", item.ID).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		event.AddOldValue("account", item.Account)
		event.AddNewValue("account", changes.Account)
		item.Account = account
		newItem.AccountID = &account.ID
		isChange = true
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

	if _, ok := input["accountId"]; ok && (item.AccountID != changes.AccountID) && (item.AccountID == nil || changes.AccountID == nil || *item.AccountID != *changes.AccountID) {

		// if err := tx.Select("id").Where("id", input["accountId"]).First(&Account{}).Error; err != nil {
		// 	tx.Rollback()
		// 	return nil, fmt.Errorf("accountId " + err.Error())
		// }
		event.AddOldValue("accountId", item.AccountID)
		event.AddNewValue("accountId", changes.AccountID)
		item.AccountID = changes.AccountID
		newItem.AccountID = changes.AccountID
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

func DeleteTodoFunc(ctx context.Context, r *GeneratedResolver, id string, tye string, unscoped *bool) (err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item := &Todo{}
	now := time.Now()
	tx := GetTransaction(ctx)
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err = GetItem(ctx, tx, TableName("todos"), item, &id); err != nil {
		tx.Rollback()
		return err
	}

	deletedAt := now.UnixNano() / 1e6

	event := NewEvent(EventMetadata{
		Type:        EventTypeDeleted,
		Entity:      "Todo",
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
		} else if err := tx.Model(&item).Updates(Todo{DeletedAt: &deletedAt, DeletedBy: principalID, UpdatedBy: principalID}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}

func (r *GeneratedMutationResolver) DeleteTodos(ctx context.Context, id []string, unscoped *bool) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteTodos(ctx, r.GeneratedResolver, id, unscoped)
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

func DeleteTodosHandler(ctx context.Context, r *GeneratedResolver, id []string, unscoped *bool) (bool, error) {
	var err error = nil

	if len(id) > 0 {
		for _, v := range id {
			err = DeleteTodoFunc(ctx, r, v, "delete", unscoped)
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

func (r *GeneratedMutationResolver) RecoveryTodos(ctx context.Context, id []string) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.RecoveryTodos(ctx, r.GeneratedResolver, id)
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

func RecoveryTodosHandler(ctx context.Context, r *GeneratedResolver, id []string) (bool, error) {
	var err error = nil

	var unscoped bool = false

	if len(id) > 0 {
		for _, v := range id {
			err = DeleteTodoFunc(ctx, r, v, "recovery", &unscoped)
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
