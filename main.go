package main

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/rum-people-preseed/weather-harvester-svc/internal"
	"github.com/rum-people-preseed/weather-harvester-svc/internal/infrastructure/http/server"
)

func main() {
	vp := viper.GetViper()
	vp.AutomaticEnv()

	digContainer := internal.SetupDependencies()
	err := digContainer.Invoke(func(server server.HttpServer) {
		zap.S().Info("Invoking Container")

		if err := server.Init(); err != nil {
			zap.S().Fatalf("failed server setup: %v", err)
		}
	})
	if err != nil {
		zap.S().Fatalf("can't setup dependencies: %v", err)
	}
}
