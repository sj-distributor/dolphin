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

func (r *GeneratedMutationResolver) CreateBookCategory(ctx context.Context, input map[string]interface{}) (item *BookCategory, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreateBookCategory(ctx, r.GeneratedResolver, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func CreateBookCategoryHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *BookCategory, err error) {
	item = &BookCategory{}

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
		Entity:      "BookCategory",
		EntityID:    item.ID,
		Date:        milliTime,
		PrincipalID: principalID,
	})

	var changes BookCategoryChanges
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

	if input["books"] != nil && input["booksId"] != nil {
		tx.Rollback()
		return nil, fmt.Errorf("booksId and books cannot coexist")
	}

	var booksIds []string

	if _, ok := input["books"]; ok {
		var booksMaps []map[string]interface{}
		for _, v := range input["books"].([]interface{}) {
			booksMaps = append(booksMaps, v.(map[string]interface{}))
		}

		for _, v := range booksMaps {
			var books *Book
			if v["id"] == nil {
				books, err = r.Handlers.CreateBook(ctx, r, v)
			} else {
				books, err = r.Handlers.UpdateBook(ctx, r, v["id"].(string), v)
			}

			changes.Books = append(changes.Books, books)
			booksIds = append(booksIds, books.ID)
		}
		event.AddNewValue("books", changes.Books)
		item.Books = changes.Books
	}

	if _, ok := input["name"]; ok && (item.Name != changes.Name) {
		item.Name = changes.Name
		event.AddNewValue("name", changes.Name)
	}

	if _, ok := input["description"]; ok && (item.Description != changes.Description) && (item.Description == nil || changes.Description == nil || *item.Description != *changes.Description) {
		item.Description = changes.Description
		event.AddNewValue("description", changes.Description)
	}

	if err := tx.Omit(clause.Associations).Create(item).Error; err != nil {
		tx.Rollback()
		return item, err
	}

	if len(booksIds) > 0 {
		if err := tx.Model(&Book{}).Where("id IN(?)", booksIds).Update("category_id", item.ID).Error; err != nil {
			tx.Rollback()
			return item, err
		}
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}
func (r *GeneratedMutationResolver) UpdateBookCategory(ctx context.Context, id string, input map[string]interface{}) (item *BookCategory, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdateBookCategory(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func UpdateBookCategoryHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *BookCategory, err error) {
	item = &BookCategory{}
	newItem := &BookCategory{}

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
		Entity:      "BookCategory",
		EntityID:    id,
		Date:        milliTime,
		PrincipalID: principalID,
	})

	var changes BookCategoryChanges
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

	if input["books"] != nil && input["booksId"] != nil {
		tx.Rollback()
		return nil, fmt.Errorf("booksId and books cannot coexist")
	}

	if err = GetItem(ctx, tx, TableName("book_categories"), item, &id); err != nil {
		tx.Rollback()
		return nil, err
	}

	if item.UpdatedBy != nil && principalID != nil && *item.UpdatedBy != *principalID {
		newItem.UpdatedBy = principalID
	}

	if _, ok := input["books"]; ok {
		if err := tx.Unscoped().Model(&Book{}).Where("category_id = ?", id).Update("category_id", "").Error; err != nil {
			tx.Rollback()
			return item, err
		}

		var booksMaps []map[string]interface{}
		for _, v := range input["books"].([]interface{}) {
			vMaps := v.(map[string]interface{})
			booksMaps = append(booksMaps, vMaps)
		}

		for _, v := range booksMaps {
			books := &Book{}
			v["categoryId"] = id
			if v["id"] == nil {
				books, err = r.Handlers.CreateBook(ctx, r, v)
			} else {
				books, err = r.Handlers.UpdateBook(ctx, r, v["id"].(string), v)
			}

			changes.Books = append(changes.Books, books)
		}

		event.AddNewValue("books", changes.Books)
		item.Books = changes.Books
		newItem.Books = changes.Books
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

	if _, ok := input["description"]; ok && (item.Description != changes.Description) && (item.Description == nil || changes.Description == nil || *item.Description != *changes.Description) {
		event.AddOldValue("description", item.Description)
		event.AddNewValue("description", changes.Description)
		item.Description = changes.Description
		newItem.Description = changes.Description
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

func DeleteBookCategoryFunc(ctx context.Context, r *GeneratedResolver, id string, tye string, unscoped *bool) (err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item := &BookCategory{}
	now := time.Now()
	tx := GetTransaction(ctx)
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err = GetItem(ctx, tx, TableName("book_categories"), item, &id); err != nil {
		tx.Rollback()
		return err
	}

	deletedAt := now.UnixNano() / 1e6

	event := NewEvent(EventMetadata{
		Type:        EventTypeDeleted,
		Entity:      "BookCategory",
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
		} else if err := tx.Model(&item).Updates(BookCategory{DeletedAt: &deletedAt, DeletedBy: principalID, UpdatedBy: principalID}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}

func (r *GeneratedMutationResolver) DeleteBookCategories(ctx context.Context, id []string, unscoped *bool) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteBookCategories(ctx, r.GeneratedResolver, id, unscoped)
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

func DeleteBookCategoriesHandler(ctx context.Context, r *GeneratedResolver, id []string, unscoped *bool) (bool, error) {
	var err error = nil

	if len(id) > 0 {
		for _, v := range id {
			err = DeleteBookCategoryFunc(ctx, r, v, "delete", unscoped)
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

func (r *GeneratedMutationResolver) RecoveryBookCategories(ctx context.Context, id []string) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.RecoveryBookCategories(ctx, r.GeneratedResolver, id)
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

func RecoveryBookCategoriesHandler(ctx context.Context, r *GeneratedResolver, id []string) (bool, error) {
	var err error = nil

	var unscoped bool = false

	if len(id) > 0 {
		for _, v := range id {
			err = DeleteBookCategoryFunc(ctx, r, v, "recovery", &unscoped)
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

func (r *GeneratedMutationResolver) CreateBook(ctx context.Context, input map[string]interface{}) (item *Book, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreateBook(ctx, r.GeneratedResolver, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func CreateBookHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Book, err error) {
	item = &Book{}

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
		Entity:      "Book",
		EntityID:    item.ID,
		Date:        milliTime,
		PrincipalID: principalID,
	})

	var changes BookChanges
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

	if input["category"] != nil && input["categoryId"] != nil {
		tx.Rollback()
		return nil, fmt.Errorf("categoryId and category cannot coexist")
	}

	if _, ok := input["category"]; ok {
		var category *BookCategory
		categoryInput := input["category"].(map[string]interface{})

		if categoryInput["id"] == nil {
			category, err = r.Handlers.CreateBookCategory(ctx, r, categoryInput)
		} else {
			category, err = r.Handlers.UpdateBookCategory(ctx, r, categoryInput["id"].(string), categoryInput)
		}

		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf(fmt.Sprintf("category %s", err.Error()))
		}

		event.AddNewValue("category", changes.Category)
		item.Category = category
		item.CategoryID = &category.ID
	}

	if _, ok := input["title"]; ok && (item.Title != changes.Title) {
		item.Title = changes.Title
		event.AddNewValue("title", changes.Title)
	}

	if _, ok := input["author"]; ok && (item.Author != changes.Author) {
		item.Author = changes.Author
		event.AddNewValue("author", changes.Author)
	}

	if _, ok := input["price"]; ok && (item.Price != changes.Price) && (item.Price == nil || changes.Price == nil || *item.Price != *changes.Price) {
		item.Price = changes.Price
		event.AddNewValue("price", changes.Price)
	}

	if _, ok := input["publishDateAt"]; ok && (item.PublishDateAt != changes.PublishDateAt) && (item.PublishDateAt == nil || changes.PublishDateAt == nil || *item.PublishDateAt != *changes.PublishDateAt) {
		item.PublishDateAt = changes.PublishDateAt
		event.AddNewValue("publishDateAt", changes.PublishDateAt)
	}

	if _, ok := input["categoryId"]; ok && (item.CategoryID != changes.CategoryID) && (item.CategoryID == nil || changes.CategoryID == nil || *item.CategoryID != *changes.CategoryID) {

		// if err := tx.Select("id").Where("id", input["categoryId"]).First(&Category{}).Error; err != nil {
		// 	tx.Rollback()
		// 	return nil, fmt.Errorf("categoryId " + err.Error())
		// }
		item.CategoryID = changes.CategoryID
		event.AddNewValue("categoryId", changes.CategoryID)
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
func (r *GeneratedMutationResolver) UpdateBook(ctx context.Context, id string, input map[string]interface{}) (item *Book, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdateBook(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func UpdateBookHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Book, err error) {
	item = &Book{}
	newItem := &Book{}

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
		Entity:      "Book",
		EntityID:    id,
		Date:        milliTime,
		PrincipalID: principalID,
	})

	var changes BookChanges
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

	if input["category"] != nil && input["categoryId"] != nil {
		tx.Rollback()
		return nil, fmt.Errorf("categoryId and category cannot coexist")
	}

	if err = GetItem(ctx, tx, TableName("books"), item, &id); err != nil {
		tx.Rollback()
		return nil, err
	}

	if item.UpdatedBy != nil && principalID != nil && *item.UpdatedBy != *principalID {
		newItem.UpdatedBy = principalID
	}

	if _, ok := input["category"]; ok {
		var category *BookCategory
		categoryInput := input["category"].(map[string]interface{})

		if categoryInput["id"] == nil {
			category, err = r.Handlers.CreateBookCategory(ctx, r, categoryInput)
		} else {
			category, err = r.Handlers.UpdateBookCategory(ctx, r, categoryInput["id"].(string), categoryInput)
		}

		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf(fmt.Sprintf("category %s", err.Error()))
		}

		if err := tx.Model(&item).Association("Category").Clear(); err != nil {
			tx.Rollback()
			return nil, err
		}

		if err := tx.Model(&BookCategory{}).Where("id = ?", category.ID).Update("books_id", item.ID).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		event.AddOldValue("category", item.Category)
		event.AddNewValue("category", changes.Category)
		item.Category = category
		newItem.CategoryID = &category.ID
		isChange = true
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		event.AddOldValue("id", item.ID)
		event.AddNewValue("id", changes.ID)
		item.ID = changes.ID
		newItem.ID = changes.ID
		isChange = true
	}

	if _, ok := input["title"]; ok && (item.Title != changes.Title) {
		event.AddOldValue("title", item.Title)
		event.AddNewValue("title", changes.Title)
		item.Title = changes.Title
		newItem.Title = changes.Title
		isChange = true
	}

	if _, ok := input["author"]; ok && (item.Author != changes.Author) {
		event.AddOldValue("author", item.Author)
		event.AddNewValue("author", changes.Author)
		item.Author = changes.Author
		newItem.Author = changes.Author
		isChange = true
	}

	if _, ok := input["price"]; ok && (item.Price != changes.Price) && (item.Price == nil || changes.Price == nil || *item.Price != *changes.Price) {
		event.AddOldValue("price", item.Price)
		event.AddNewValue("price", changes.Price)
		item.Price = changes.Price
		newItem.Price = changes.Price
		isChange = true
	}

	if _, ok := input["publishDateAt"]; ok && (item.PublishDateAt != changes.PublishDateAt) && (item.PublishDateAt == nil || changes.PublishDateAt == nil || *item.PublishDateAt != *changes.PublishDateAt) {
		event.AddOldValue("publishDateAt", item.PublishDateAt)
		event.AddNewValue("publishDateAt", changes.PublishDateAt)
		item.PublishDateAt = changes.PublishDateAt
		newItem.PublishDateAt = changes.PublishDateAt
		isChange = true
	}

	if _, ok := input["categoryId"]; ok && (item.CategoryID != changes.CategoryID) && (item.CategoryID == nil || changes.CategoryID == nil || *item.CategoryID != *changes.CategoryID) {

		// if err := tx.Select("id").Where("id", input["categoryId"]).First(&Category{}).Error; err != nil {
		// 	tx.Rollback()
		// 	return nil, fmt.Errorf("categoryId " + err.Error())
		// }
		event.AddOldValue("categoryId", item.CategoryID)
		event.AddNewValue("categoryId", changes.CategoryID)
		item.CategoryID = changes.CategoryID
		newItem.CategoryID = changes.CategoryID
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

func DeleteBookFunc(ctx context.Context, r *GeneratedResolver, id string, tye string, unscoped *bool) (err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item := &Book{}
	now := time.Now()
	tx := GetTransaction(ctx)
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err = GetItem(ctx, tx, TableName("books"), item, &id); err != nil {
		tx.Rollback()
		return err
	}

	deletedAt := now.UnixNano() / 1e6

	event := NewEvent(EventMetadata{
		Type:        EventTypeDeleted,
		Entity:      "Book",
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
		} else if err := tx.Model(&item).Updates(Book{DeletedAt: &deletedAt, DeletedBy: principalID, UpdatedBy: principalID}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}

func (r *GeneratedMutationResolver) DeleteBooks(ctx context.Context, id []string, unscoped *bool) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteBooks(ctx, r.GeneratedResolver, id, unscoped)
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

func DeleteBooksHandler(ctx context.Context, r *GeneratedResolver, id []string, unscoped *bool) (bool, error) {
	var err error = nil

	if len(id) > 0 {
		for _, v := range id {
			err = DeleteBookFunc(ctx, r, v, "delete", unscoped)
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

func (r *GeneratedMutationResolver) RecoveryBooks(ctx context.Context, id []string) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.RecoveryBooks(ctx, r.GeneratedResolver, id)
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

func RecoveryBooksHandler(ctx context.Context, r *GeneratedResolver, id []string) (bool, error) {
	var err error = nil

	var unscoped bool = false

	if len(id) > 0 {
		for _, v := range id {
			err = DeleteBookFunc(ctx, r, v, "recovery", &unscoped)
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
