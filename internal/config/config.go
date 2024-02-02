package config

import (
	"os"

	"go.uber.org/dig"
)

type OWMConfig struct {
	ApiKey string
}

type DbConfig struct {
	Host        string
	Port        string
	User        string
	Password    string
	DBName      string
	AutoMigrate string
}

type ServerConfig struct {
	Port string
}

type Config struct {
	dig.Out

	OWMConfig *OWMConfig
	Server    *ServerConfig
	DB        *DbConfig
}

func ProvideConfig() interface{} {
	return func() Config {
		return Config{
			OWMConfig: &OWMConfig{
				ApiKey: os.Getenv("OWM_API_KEY"),
			},
			Server: &ServerConfig{
				Port: os.Getenv("PORT"),
			},
			DB: &DbConfig{
				Host:        os.Getenv("DB_HOST"),
				Port:        os.Getenv("DB_PORT"),
				User:        os.Getenv("DB_USER"),
				Password:    os.Getenv("DB_PASSWORD"),
				DBName:      os.Getenv("DB_NAME"),
				AutoMigrate: os.Getenv("DB_MIGRATE"),
			},
		}
	}
}
