package repository

import (
	"github.com/rum-people-preseed/weather-harvester-svc/internal/domain/entity"
)

type WeatherHistoryRepository interface {
	ICRUDStore[entity.WeatherHistory]
}
