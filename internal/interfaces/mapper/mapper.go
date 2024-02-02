package mapper

import (
	"go.uber.org/dig"

	"github.com/rum-people-preseed/weather-harvester-svc/dto"
	"github.com/rum-people-preseed/weather-harvester-svc/internal/domain/entity"
)

type MapperOut struct {
	dig.Out

	WeatherHistoryMapper BaseMapper[entity.WeatherHistory, dto.WeatherHistoryDTO]
}

func ProvideMapper() interface{} {
	return func() MapperOut {
		return MapperOut{
			WeatherHistoryMapper: NewWeatherHistoryMapper(),
		}
	}
}
