package server

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/rum-people-preseed/weather-harvester-svc/internal/config"
	"github.com/rum-people-preseed/weather-harvester-svc/internal/infrastructure/http/server/response"
	"github.com/rum-people-preseed/weather-harvester-svc/internal/infrastructure/logger"
	"github.com/rum-people-preseed/weather-harvester-svc/internal/infrastructure/validators"
	"github.com/rum-people-preseed/weather-harvester-svc/internal/interfaces"
)

type HttpServer interface {
	Init() error
}

type EchoServer struct {
	cfg                *config.ServerConfig
	deps               *interfaces.RouterConfigDeps
	ignoreDumpBodyPath map[string]string
}

func (s *EchoServer) Init() error {
	port := s.cfg.Port
	if port == "" {
		port = "8000"
	}

	e := echo.New()

	// Id Generator Middleware for Logger
	e.Use(middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		Generator: func() string {
			return uuid.New().String()
		},
	}))

	// Logger Hook
	e.Use(logger.LoggerHook())

	// Dump handler req/resp
	e.Use(middleware.BodyDumpWithConfig(middleware.BodyDumpConfig{
		Skipper: func(c echo.Context) bool {
			method := c.Request().Method
			path := c.Path()
			// echo get http method from request
			httpMethod, ok := s.ignoreDumpBodyPath[path]
			if !ok {
				return false
			}

			if strings.EqualFold(method, httpMethod) {
				return true
			}
			return false
		},
		Handler: logger.BodyDumpHandler,
	}))

	e.Validator = validators.NewCustomValidator()
	e.HTTPErrorHandler = response.HttpErrorHandler

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "Ack")
	})

	interfaces.Router(e, *s.deps)

	return e.Start(fmt.Sprintf(":%s", port))
}

func ProvideEchoServer() interface{} {
	return func(config *config.ServerConfig, deps interfaces.RouterConfigDeps) HttpServer {
		return &EchoServer{
			cfg:  config,
			deps: &deps,
			ignoreDumpBodyPath: map[string]string{
				"/health": "GET",
			},
		}
	}
}
