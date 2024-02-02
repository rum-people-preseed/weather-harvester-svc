package handler

import "go.uber.org/dig"

type DepsOut struct {
	dig.Out

	WeatherHistoryHandler WeatherHistory
}

func ProvideHandler() interface{} {
	return func(weatherHistoryDepsIn WeatherHistoryConfig) DepsOut {
		return DepsOut{
			WeatherHistoryHandler: NewWeatherHistoryHandler(weatherHistoryDepsIn),
		}
	}
}
