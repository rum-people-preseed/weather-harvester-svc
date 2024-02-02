package services

import (
	owm "github.com/briandowns/openweathermap"
	"go.uber.org/dig"

	"github.com/rum-people-preseed/weather-harvester-svc/dto"
	"github.com/rum-people-preseed/weather-harvester-svc/internal/domain/repository"
	"github.com/rum-people-preseed/weather-harvester-svc/internal/infrastructure/api"
)

type DigWeatherHistoryService struct {
	dig.In

	OWMApi             api.OWMApiClient
	WeatherHistoryRepo repository.WeatherHistoryRepository
}

type WeatherHistoryService interface {
	GetWeatherHistory(*owm.Coordinates) ([]dto.WeatherHistoryDTO, error)
}

type weatherHistoryService struct {
	weatherHistoryRepo repository.WeatherHistoryRepository
	ownApi             api.OWMApiClient
}

func NewWeatherHistoryService(deps DigWeatherHistoryService) WeatherHistoryService {
	return &weatherHistoryService{
		weatherHistoryRepo: deps.WeatherHistoryRepo,
		ownApi:             deps.OWMApi,
	}
}

func (s *weatherHistoryService) GetWeatherHistory(location *owm.Coordinates) ([]dto.WeatherHistoryDTO, error) {
	history, err := s.ownApi.GetWeatherHistory(location)
	if err != nil {
		return nil, err
	}
	return convertToDTO(history), nil
}

func convertToDTO(history *owm.HistoricalWeatherData) []dto.WeatherHistoryDTO {
	result := make([]dto.WeatherHistoryDTO, 0, len(history.List))
	for _, item := range history.List {
		entry := dto.WeatherHistoryDTO{
			Timestamp:   item.Dt,
			Temperature: item.Main.Temp,
		}
		result = append(result, entry)
	}
	return result
}
