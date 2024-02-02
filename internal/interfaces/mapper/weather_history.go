package mapper

import (
	"github.com/shopspring/decimal"

	"github.com/rum-people-preseed/weather-harvester-svc/dto"
	"github.com/rum-people-preseed/weather-harvester-svc/internal/domain/entity"
)

func NewWeatherHistoryMapper() BaseMapper[entity.WeatherHistory, dto.WeatherHistoryDTO] {
	return &weatherHistoryMapper{}
}

type weatherHistoryMapper struct{}

func (m *weatherHistoryMapper) ToDTO(e entity.WeatherHistory) dto.WeatherHistoryDTO {
	return dto.WeatherHistoryDTO{
		Timestamp:   int(e.Timestamp),
		Temperature: e.Temperature.InexactFloat64(),
	}
}

func (m *weatherHistoryMapper) FromDTO(d dto.WeatherHistoryDTO) entity.WeatherHistory {
	return entity.WeatherHistory{
		Timestamp:   int64(d.Timestamp),
		Temperature: decimal.NewFromFloat(d.Temperature),
	}
}

func (m *weatherHistoryMapper) ToManyDTO(wh []entity.WeatherHistory) []dto.WeatherHistoryDTO {
	items := make([]dto.WeatherHistoryDTO, len(wh))
	for k, v := range wh {
		items[k] = m.ToDTO(v)
	}
	return items
}

func (m *weatherHistoryMapper) FromManyDTO(wh []dto.WeatherHistoryDTO) []entity.WeatherHistory {
	items := make([]entity.WeatherHistory, len(wh))
	for k, v := range wh {
		items[k] = m.FromDTO(v)
	}
	return items
}
