package gorm_repo

import (
	"context"

	"gorm.io/gorm"

	"github.com/rum-people-preseed/weather-harvester-svc/internal/domain/repository"
)

type atomicRepository struct {
	conn *gorm.DB
}

func NewAtomicRepository(db *gorm.DB) *atomicRepository {
	return &atomicRepository{conn: db}
}

// Do
// s.conn.WithContext(ctx).Transaction is just a wrapper function provided by Bun around a DB transaction,
// it takes care of proper handling the commit and rollback steps depending on the
// error returned by the given callback function
func (s *atomicRepository) Do(ctx context.Context, fn repository.UnitOfWorkBlock) error {
	return s.conn.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		store := &uowStore{
			weatherHistoryRepo: *NewWeatherHistoryRepository(tx),
		}
		return fn(store)
	})
}

type uowStore struct {
	weatherHistoryRepo WeatherHistoryRepository
}

func (u uowStore) WeatherHistoryRepo() repository.WeatherHistoryRepository {
	return u.weatherHistoryRepo
}
