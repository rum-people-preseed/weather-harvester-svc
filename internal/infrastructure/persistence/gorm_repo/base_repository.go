package gorm_repo

import (
	"context"
	"math"

	"gorm.io/gorm"

	"github.com/rum-people-preseed/weather-harvester-svc/internal/domain/repository"
)

const (
	minPageSize  = 1
	minLimitSize = 10
	maxLimitSize = 100
)

type BaseRepository[E any] struct {
	db *gorm.DB
}

func New[T any](db *gorm.DB) *BaseRepository[T] {
	return &BaseRepository[T]{db: db}
}

// InitializeTable create table in db if doesn't exist
func (r *BaseRepository[E]) InitializeTable(ctx context.Context, entity *E) error {
	if !r.db.WithContext(ctx).Migrator().HasTable(&entity) {
		return r.db.WithContext(ctx).Migrator().CreateTable(&entity)
	}
	return nil
}

// Insert create new record in the database
func (r *BaseRepository[E]) Insert(ctx context.Context, entity *E) error {
	err := r.db.WithContext(ctx).Create(&entity).Error
	if err != nil {
		return err
	}
	return nil
}

// Save updates value in database. If value doesn't contain a matching primary key, value is inserted.
func (r *BaseRepository[E]) Save(ctx context.Context, entity *E) error {
	err := r.db.WithContext(ctx).Save(&entity).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *BaseRepository[E]) Update(ctx context.Context, entity *E) error {
	err := r.db.WithContext(ctx).Updates(&entity).Error
	if err != nil {
		return err
	}
	return nil
}

// FindByID retrieve a record by id from a database
func (r *BaseRepository[E]) FindByID(ctx context.Context, id uint) (E, error) {
	var model E
	err := r.db.WithContext(ctx).First(&model, id).Error
	if err != nil {
		return model, err
	}
	return model, nil
}

// FindAll retrieve a record by specification
func (r *BaseRepository[E]) FindAll(ctx context.Context, specification repository.Specification) ([]E, error) {
	var models []E

	qb := r.db.WithContext(ctx)

	for _, query := range specification.Joins() {
		qb = qb.Joins(query)
	}

	for _, query := range specification.Preload() {
		qb = qb.Preload(query)
	}

	if len(specification.GetQuery()) > 0 {
		query := specification.GetQuery()
		value := specification.GetValues()
		qb = qb.Where(query, value...)
	}

	find := qb.Find(&models)
	err := find.Error
	if err != nil {
		return []E{}, err
	}
	return models, nil
}

func (r *BaseRepository[E]) FindOne(ctx context.Context, specification repository.Specification) (E, error) {
	var model E
	qb := r.db.WithContext(ctx)

	for _, query := range specification.Joins() {
		qb = qb.Joins(query)
	}
	for _, query := range specification.Preload() {
		qb = qb.Preload(query)
	}

	err := qb.Where(specification.GetQuery(), specification.GetValues()...).Take(&model).Error
	if err != nil {
		return model, err
	}
	return model, nil
}

// Delete record by specification and entity
func (r *BaseRepository[E]) Delete(ctx context.Context, entity *E, specification repository.Specification) error {
	err := r.db.WithContext(ctx).Where(specification.GetQuery(), specification.GetValues()).Delete(&entity).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteByID record by ID
func (r *BaseRepository[E]) DeleteByID(ctx context.Context, id uint) error {
	entity, err := r.FindByID(ctx, id)
	if err != nil {
		return err
	}
	err = r.db.WithContext(ctx).Delete(&entity).Error
	if err != nil {
		return err
	}
	return nil
}

// Paginated select paginated data results
func (r *BaseRepository[E]) Paginated(ctx context.Context, specification repository.Specification, page, limit int) (repository.PaginatedResults[E], error) {
	var models []E
	var count int64

	if page < minPageSize {
		page = minPageSize
	}
	if limit < minLimitSize {
		limit = minLimitSize
	}
	if limit > maxLimitSize {
		limit = maxLimitSize
	}

	qb := r.db.WithContext(ctx)

	for _, query := range specification.Joins() {
		qb = qb.Joins(query)
	}
	for _, query := range specification.Preload() {
		qb = qb.Preload(query)
	}
	if len(specification.GetQuery()) > 0 {
		qb = qb.Where(specification.GetQuery(), specification.GetValues()...)
	}

	// if len(specification.Sort()) > 0 {
	qb = qb.Order("id desc")
	// }
	err := qb.Offset((page - 1) * limit).Limit(limit).Find(&models).Error
	if err != nil {
		return repository.PaginatedResults[E]{}, err
	}

	// Get total count of records without pagination
	qbCount := r.db.WithContext(ctx)
	qbCount = qbCount.Model(&models)
	if len(specification.GetQuery()) > 0 {
		qbCount = qbCount.Where(specification.GetQuery(), specification.GetValues()...)
	}
	if err := qbCount.Count(&count).Error; err != nil {
		return repository.PaginatedResults[E]{}, err
	}
	totalPages := int64(math.Ceil(float64(count) / float64(limit)))

	return repository.PaginatedResults[E]{
		Results:    models,
		Count:      count,
		TotalPages: totalPages,
	}, nil
}
