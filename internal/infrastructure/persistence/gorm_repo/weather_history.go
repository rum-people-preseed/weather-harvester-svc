package gorm_repo

import (
	"gorm.io/gorm"

	"github.com/rum-people-preseed/weather-harvester-svc/internal/domain/entity"
	"github.com/rum-people-preseed/weather-harvester-svc/internal/domain/repository"
)

var _ repository.WeatherHistoryRepository = &WeatherHistoryRepository{}

type WeatherHistoryRepository struct {
	*BaseRepository[entity.WeatherHistory]
}

func NewWeatherHistoryRepository(db *gorm.DB) *WeatherHistoryRepository {
	return &WeatherHistoryRepository{
		New[entity.WeatherHistory](db),
	}
}
