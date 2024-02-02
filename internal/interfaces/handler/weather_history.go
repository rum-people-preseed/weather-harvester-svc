package handler

import (
	"net/http"
	"strconv"

	owm "github.com/briandowns/openweathermap"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"

	"github.com/rum-people-preseed/weather-harvester-svc/internal/application/services"
	"github.com/rum-people-preseed/weather-harvester-svc/internal/infrastructure/http/server/response"
	"github.com/rum-people-preseed/weather-harvester-svc/models"
)

type WeatherHistoryConfig struct {
	dig.In

	WeatherHistoryService services.WeatherHistoryService
}

type WeatherHistory interface {
	GetWeatherHistory(c echo.Context) error
}

func NewWeatherHistoryHandler(deps WeatherHistoryConfig) WeatherHistory {
	return &weatherHistoryHandler{
		weatherHistoryService: deps.WeatherHistoryService,
	}
}

type weatherHistoryHandler struct {
	weatherHistoryService services.WeatherHistoryService
}

func (l *weatherHistoryHandler) GetWeatherHistory(c echo.Context) error {
	var location models.Coordinates
	if err := c.Bind(&location); err != nil {
		return response.HttpErrorBodyValidation(err)
	}

	lon, err := strconv.ParseFloat(location.Lon, 64)
	if err != nil {
		return response.HttpError(http.StatusInternalServerError)
	}
	lat, err := strconv.ParseFloat(location.Lat, 64)
	if err != nil {
		return response.HttpError(http.StatusInternalServerError)
	}
	coordinate := &owm.Coordinates{
		Longitude: lon,
		Latitude:  lat,
	}

	history, err := l.weatherHistoryService.GetWeatherHistory(coordinate)
	if err != nil {
		return response.HttpError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, history)
}
