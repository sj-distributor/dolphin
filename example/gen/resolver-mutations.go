package gen

import (
	"context"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/sj-distributor/dolphin-example/enums"
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
		return nil, err
	}

	err = CheckStructFieldIsEmpty(item, input)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	item.ID = uuid.Must(uuid.NewV4()).String()
	item.CreatedAt = milliTime
	item.CreatedBy = principalID

	if _, ok := input["title"]; ok && (item.Title != changes.Title) {
		event.AddNewValue("title", changes.Title)
		item.Title = changes.Title
	}

	if _, ok := input["remark"]; ok && (item.Remark != changes.Remark) {
		event.AddNewValue("remark", changes.Remark)
		item.Remark = changes.Remark
	}

	if err := tx.Create(item).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	AddMutationEvent(ctx, event)

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
		return nil, err
	}

	err = CheckStructFieldIsEmpty(item, input)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err = GetItem(ctx, tx, "todos", item, &id); err != nil {
		tx.Rollback()
		return nil, err
	}

	if *item.UpdatedBy != *principalID {
		newItem.UpdatedBy = principalID
	}

	if _, ok := input["title"]; ok && (item.Title != changes.Title) {
		event.AddOldValue("title", item.Title)
		event.AddNewValue("title", changes.Title)
		item.Title = changes.Title
		newItem.Title = changes.Title
		isChange = true
	}

	if _, ok := input["remark"]; ok && (item.Remark != changes.Remark) {
		event.AddOldValue("remark", item.Remark)
		event.AddNewValue("remark", changes.Remark)
		item.Remark = changes.Remark
		newItem.Remark = changes.Remark
		isChange = true
	}

	if !isChange {
		return nil, fmt.Errorf(enums.DataNotChange)
	}

	if err := tx.Where("id = ?", item.ID).Updates(newItem).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	AddMutationEvent(ctx, event)

	return
}
