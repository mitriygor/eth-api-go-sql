package app

import (
	"eth-api/src/config"
	"eth-api/src/controllers"
	"eth-api/src/database"
	"eth-api/src/middleware"
	"eth-api/src/models/dao"
	"eth-api/src/routes"
	"eth-api/src/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html/v2"
	"gorm.io/gorm"
)

// New creates and configures a new Fiber application with necessary middlewares, routes, and services.
func New() *fiber.App {
	engine := html.New("./src/views", ".html")
	app := fiber.New(fiber.Config{Views: engine})

	db := database.GetDB()
	conf := config.New()

	ethBlockController, ethBlockService, ethBlockDAO := initEthBlockController(db, conf)
	ethTransactionController := initEthTransactionController(db, conf, ethBlockService)
	cronService := initCronService(ethBlockDAO, ethBlockService, conf)

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET, OPTIONS",
		AllowHeaders: "*",
	}))

	app.Use(middleware.LoggingMiddleware)
	routes.Setup(app, ethBlockController, ethTransactionController)
	cronService.RunCron()

	return app
}

// initEthBlockController initializes and returns the EthBlock controller, service, and DAO.
func initEthBlockController(db *gorm.DB, config *config.Config) (*controllers.EthBlockController, services.EthBlockService, dao.EthBlockDAO) {
	ethBlockDAO := dao.NewEthBlockDAO(db)
	ethBlockService := services.NewEthBlockService(ethBlockDAO, config)
	ethBlockController := controllers.NewEthBlockController(ethBlockService)
	return ethBlockController, ethBlockService, ethBlockDAO
}

// initEthTransactionController initializes and returns the EthTransaction controller.
func initEthTransactionController(db *gorm.DB, config *config.Config, ethBlockService services.EthBlockService) *controllers.EthTransactionController {
	ethTransactionDAO := dao.NewEthTransactionDAO(db)
	ethTransactionService := services.NewEthTransactionService(ethTransactionDAO, config, ethBlockService)
	ethTransactionController := controllers.NewEthTransactionController(ethTransactionService)
	return ethTransactionController
}

// initCronService initializes and returns the cron service.
func initCronService(ethBlockDAO dao.EthBlockDAO, ethBlockService services.EthBlockService, config *config.Config) services.CronService {
	return services.NewCronService(ethBlockDAO, ethBlockService, config)
}
