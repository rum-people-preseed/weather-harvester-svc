package interfaces

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"

	"github.com/rum-people-preseed/weather-harvester-svc/internal/interfaces/handler"
)

type RouterConfigDeps struct {
	dig.In

	WeatherHistory handler.WeatherHistory
}

func Router(ex *echo.Echo, deps RouterConfigDeps) {
	ex.POST("/weather-history", deps.WeatherHistory.GetWeatherHistory)
}
