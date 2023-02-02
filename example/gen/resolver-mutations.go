package gen

import (
	"context"
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

	tx := GetTransaction(ctx)
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	item.ID = uuid.Must(uuid.NewV4()).String()
	item.CreatedAt = milliTime
	item.CreatedBy = principalID

	if _, ok := input["title"]; ok && (item.Title != changes.Title) {
		item.Title = changes.Title

		event.AddNewValue("title", changes.Title)
	}

	if err := tx.Create(item).Error; err != nil {
		tx.Rollback()
		return item, err
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return item, nil
}
