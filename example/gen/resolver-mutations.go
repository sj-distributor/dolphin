package gen

import (
	"context"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
)

type GeneratedMutationResolver struct{ *GeneratedResolver }

type MutationEvents struct {
	Events []Event
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
	// 获取操作人Id
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

	var changes Todo
	err = ApplyChanges(input, &changes)

	if err != nil {
		return item, err
	}

	item.ID = uuid.Must(uuid.NewV4()).String()
	item.CreatedAt = milliTime
	item.CreatedBy = principalID

	if _, ok := input["title"]; ok {
		if input["title"] == "" {
			return nil, fmt.Errorf("the title cannot be empty")
		}
		if item.Title != changes.Title {
			event.AddNewValue("title", changes.Title)
			item.Title = changes.Title
		}
	}

	if _, ok := input["remark"]; ok && (item.Remark != changes.Remark) {
		event.AddNewValue("remark", changes.Remark)
		item.Remark = changes.Remark
	}

	if err := tx.Create(item).Error; err != nil {
		tx.Rollback()
		return item, err
	}

	AddMutationEvent(ctx, event)

	return item, nil
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

	now := time.Now()
	milliTime := now.UnixNano() / 1e6
	// 获取操作人Id
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

	item.UpdatedAt = &milliTime
	item.UpdatedBy = principalID

	if err = GetItem(ctx, tx, item, &id); err != nil {
		tx.Rollback()
		return
	}

	fmt.Println("item", item)

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		event.AddOldValue("id", item.ID)
		event.AddNewValue("id", changes.ID)
		item.ID = changes.ID
	}

	if _, ok := input["title"]; ok {
		if input["title"] == "" {
			return nil, fmt.Errorf("the title cannot be empty")
		}
		if item.Title != changes.Title {
			event.AddOldValue("title", item.Title)
			event.AddNewValue("title", changes.Title)
			item.Title = changes.Title
		}
	}

	if _, ok := input["remark"]; ok && (item.Remark != changes.Remark) {
		event.AddOldValue("remark", item.Remark)
		event.AddNewValue("remark", changes.Remark)
		item.Remark = changes.Remark
	}

	if err := tx.Model(&item).Save(item).Error; err != nil {
		tx.Rollback()
		return item, err
	}

	AddMutationEvent(ctx, event)

	return
}
