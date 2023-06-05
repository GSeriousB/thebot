package main

import (
	"context"
	"tradebot/bybit/app/api/server"
	"tradebot/bybit/app/constants"
	"tradebot/bybit/app/service/logger"
	"tradebot/bybit/config"
)

func main() {
	var err error
	// Returns a struct with values from env variables
	constants.Config, err = config.LoadConfig()
	if err != nil {
		panic(err.Error())
	}
	logger.InitLogger()

	server.Init(context.Background())
}
