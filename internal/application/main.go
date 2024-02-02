package application

import (
	"go.uber.org/dig"

	"github.com/rum-people-preseed/weather-harvester-svc/internal/application/services"
)

type serviceDependencies struct {
	dig.In

	services.DigWeatherHistoryService
}

type DepsServiceOut struct {
	dig.Out

	services.WeatherHistoryService
}

func ProvideService() interface{} {
	return func(deps serviceDependencies) DepsServiceOut {
		return DepsServiceOut{
			WeatherHistoryService: services.NewWeatherHistoryService(deps.DigWeatherHistoryService),
		}
	}
}
