package repository

import (
	"context"
)

type UnitOfWorkBlock func(AtomicStore) error

type AtomicRepository interface {
	Do(context.Context, UnitOfWorkBlock) error
}

type AtomicStore interface {
	WeatherHistoryRepo() WeatherHistoryRepository
}
