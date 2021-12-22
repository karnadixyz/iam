package infra

import (
	"os"
	"strconv"
)

type AppConfig struct {
	AppEnv string
	Debug  bool
}

type ServerConfig struct {
	Addr    string
	BaseURL string
}

type Config struct {
	Server *ServerConfig
}

var AppSrv = &Config{}

func (app *Config) Init() {
	app.Server = app.GetServerConfig()
}

func (app *Config) GetAppConfig() *AppConfig {
	debug, _ := strconv.ParseBool(os.Getenv("APP_DEBUG"))
	return &AppConfig{
		AppEnv: os.Getenv("APP_ENV"),
		Debug:  debug,
	}
}

func (app *Config) IsProduction() bool {
	config := app.GetAppConfig()
	return config.AppEnv == "production"
}

func (app *Config) GetServerConfig() *ServerConfig {
	return &ServerConfig{
		Addr:    ":8080",
		BaseURL: "localhost",
	}
}
