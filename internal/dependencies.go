package internal

import (
	"go.uber.org/dig"
	"go.uber.org/zap"

	"github.com/rum-people-preseed/weather-harvester-svc/internal/application"
	"github.com/rum-people-preseed/weather-harvester-svc/internal/config"
	"github.com/rum-people-preseed/weather-harvester-svc/internal/infrastructure/api"
	"github.com/rum-people-preseed/weather-harvester-svc/internal/infrastructure/http/server"
	"github.com/rum-people-preseed/weather-harvester-svc/internal/infrastructure/persistence"
	"github.com/rum-people-preseed/weather-harvester-svc/internal/interfaces/handler"
	"github.com/rum-people-preseed/weather-harvester-svc/internal/interfaces/mapper"
)

type dInjector struct {
	*dig.Container
	dependencies []interface{}
}

func (r *dInjector) Add(provider interface{}) {
	r.dependencies = append(r.dependencies, provider)
}

func SetupDependencies() *dig.Container {
	di := dInjector{dig.New(), []interface{}{}}

	di.Add(config.ProvideConfig())
	di.Add(persistence.ProvideDatabase())
	di.Add(mapper.ProvideMapper())
	di.Add(api.ProvideApi())
	di.Add(application.ProvideService())
	di.Add(handler.ProvideHandler())
	di.Add(server.ProvideEchoServer())

	for _, dep := range di.dependencies {
		err := di.Provide(dep)
		if err != nil {
			zap.S().Fatalf("DI Setup Failed: %v", err)
		}
	}
	di.dependencies = []interface{}{}
	return di.Container
}
