package repository

import (
	"context"
)

type ICRUDStore[T any] interface {
	InitializeTable(context.Context, *T) error
	Insert(context.Context, *T) error
	Save(context.Context, *T) error
	Update(context.Context, *T) error
	FindByID(context.Context, uint) (T, error)
	FindAll(context.Context, Specification) ([]T, error)
	FindOne(context.Context, Specification) (T, error)
	Delete(context.Context, *T, Specification) error
	DeleteByID(context.Context, uint) error
	Paginated(context.Context, Specification, int, int) (PaginatedResults[T], error)
}

type Specification interface {
	Joins() []string
	Preload() []string
	GetQuery() string
	GetValues() []any
	Sort() []string
}

type PaginatedResults[T any] struct {
	Results    []T   `json:"results"`
	Count      int64 `json:"count"`
	TotalPages int64 `json:"total_pages"`
}
