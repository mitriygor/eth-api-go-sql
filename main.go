package main

import (
	"eth-api/src/app"
	"eth-api/src/database"
	"eth-api/src/helpers/logger"
	"github.com/gofiber/fiber/v2/log"
	"go.uber.org/zap"
)

func main() {

	err := logger.Initialize("info")
	if err != nil {
		panic(err)
	}
	defer func(Log *zap.SugaredLogger) {
		err := Log.Sync()
		if err != nil {
			log.Errorf("ERROR: Error the logger syn: %v", err)
		}
	}(logger.Log)

	database.Connect()
	database.AutoMigrate()

	ethApp := app.New()
	err = ethApp.Listen(":8000")
	if err != nil {
		log.Errorf("ERROR: Error starting eth-api server: %v", err)
		panic(err)
	}
}
