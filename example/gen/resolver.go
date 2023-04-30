//go:generate go run github.com/99designs/gqlgen generate
package gen

import (
	"context"
)

type ResolutionHandlers struct {
	OnEvent func(ctx context.Context, r *GeneratedResolver, e *Event) error

	CreateBookCategory     func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *BookCategory, err error)
	UpdateBookCategory     func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *BookCategory, err error)
	DeleteBookCategories   func(ctx context.Context, r *GeneratedResolver, id []string, unscoped *bool) (bool, error)
	RecoveryBookCategories func(ctx context.Context, r *GeneratedResolver, id []string) (bool, error)
	QueryBookCategory      func(ctx context.Context, r *GeneratedResolver, id string) (*BookCategory, error)
	QueryBookCategories    func(ctx context.Context, r *GeneratedResolver, opts QueryBookCategoriesHandlerOptions) (*BookCategoryResultType, error)

	BookCategoryBooks func(ctx context.Context, r *GeneratedResolver, obj *BookCategory, input map[string]interface{}) (res []*Book, err error)

	CreateBook    func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Book, err error)
	UpdateBook    func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Book, err error)
	DeleteBooks   func(ctx context.Context, r *GeneratedResolver, id []string, unscoped *bool) (bool, error)
	RecoveryBooks func(ctx context.Context, r *GeneratedResolver, id []string) (bool, error)
	QueryBook     func(ctx context.Context, r *GeneratedResolver, id string) (*Book, error)
	QueryBooks    func(ctx context.Context, r *GeneratedResolver, opts QueryBooksHandlerOptions) (*BookResultType, error)

	BookCategory func(ctx context.Context, r *GeneratedResolver, obj *Book) (res *BookCategory, err error)
}

func DefaultResolutionHandlers() ResolutionHandlers {
	handlers := ResolutionHandlers{
		OnEvent: func(ctx context.Context, r *GeneratedResolver, e *Event) error { return nil },

		CreateBookCategory:     CreateBookCategoryHandler,
		UpdateBookCategory:     UpdateBookCategoryHandler,
		DeleteBookCategories:   DeleteBookCategoriesHandler,
		RecoveryBookCategories: RecoveryBookCategoriesHandler,
		QueryBookCategory:      QueryBookCategoryHandler,
		QueryBookCategories:    QueryBookCategoriesHandler,

		BookCategoryBooks: BookCategoryBooksHandler,

		CreateBook:    CreateBookHandler,
		UpdateBook:    UpdateBookHandler,
		DeleteBooks:   DeleteBooksHandler,
		RecoveryBooks: RecoveryBooksHandler,
		QueryBook:     QueryBookHandler,
		QueryBooks:    QueryBooksHandler,

		BookCategory: BookCategoryHandler,
	}
	return handlers
}

type GeneratedResolver struct {
	Handlers        ResolutionHandlers
	DB              *DB
	EventController *EventController
}
