package config

import (

	// if using go modules

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type JwtConfig struct {
	JWT_ACCESS_SECRET  string `env:"JWT_ACCESS_SECRET"`
	JWT_REFRESH_SECRET string `env:"JWT_REFRESH_SECRET"`
	JWT_ACCESS_EXP     int    `env:"JWT_ACCESS_EXP"`
	JWT_REFRESH_EXP    int    `env:"JWT_REFRESH_EXP"`
}

type HTTPServerConfig struct {
	HTTPSERVER_LISTEN                      string `env:"HTTPSERVER_LISTEN"`
	HTTPSERVER_PORT                        string `env:"HTTPSERVER_PORT"`
	HTTPSERVER_READ_TIMEOUT                int    `env:"HTTPSERVER_READ_TIMEOUT"`
	HTTPSERVER_WRITE_TIMEOUT               int    `env:"HTTPSERVER_WRITE_TIMEOUT"`
	HTTPSERVER_MAX_CONNECTIONS_PER_IP      int    `env:"HTTPSERVER_MAX_CONNECTIONS_PER_IP"`
	HTTPSERVER_MAX_REQUESTS_PER_CONNECTION int    `env:"HTTPSERVER_MAX_REQUESTS_PER_CONNECTION"`
	HTTPSERVER_MAX_KEEP_ALIVE_DURATION     int    `env:"HTTPSERVER_MAX_KEEP_ALIVE_DURATION"`
}

type LogConfig struct {
	LOG_FILE_PATH      string `env:"LOG_FILE_PATH"`
	LOG_FILE_NAME      string `env:"LOG_FILE_NAME"`
	LOG_FILE_MAXSIZE   int    `env:"LOG_FILE_MAXSIZE"`
	LOG_FILE_MAXBACKUP int    `env:"LOG_FILE_MAXBACKUP"`
	LOG_FILE_MAXAGE    int    `env:"LOG_FILE_MAXAGE"`
}

type ByBitConfig struct {
	BYBIT_API_KEY    string `env:"BYBIT_API_KEY"`
	BYBIT_SECRET_KEY string `env:"BYBIT_SECRET_KEY"`
	BYBIT_ENDPOINT   string `env:"BYBIT_ENDPOINT"`
}

type CoreServiceConfig struct {
	Environment      string `env:"ENVIRONMENT"`
	ProjectVersion   string `env:"VERSION"`
	HTTPServerConfig HTTPServerConfig
	LogConfig        LogConfig
	ByBitConfig      ByBitConfig
}

var Config *CoreServiceConfig

func LoadConfig() (*CoreServiceConfig, error) {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file " + err.Error())
	}

	config := CoreServiceConfig{}
	if err := env.Parse(&config); err != nil {
		panic("unable to load env config " + err.Error())
	}

	return &config, nil
}
