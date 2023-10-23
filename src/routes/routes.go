package routes

import (
	"eth-api/src/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App, ethBlockController *controllers.EthBlockController, ethTransactionController *controllers.EthTransactionController) {

	api := app.Group("v1")

	ethBlocks := api.Group("eth-blocks")
	ethBlocks.Get("latest", func(c *fiber.Ctx) error {
		return ethBlockController.GetLatestBlocksController(c)
	})
	ethBlocks.Get(":number", func(c *fiber.Ctx) error {
		return ethBlockController.GetBlockByNumberController(c)
	})

	ethTransactions := api.Group("eth-transactions")
	ethTransactions.Get(":hash", func(c *fiber.Ctx) error {
		return ethTransactionController.GetTransactionByHashController(c)
	})
	ethTransactionsByAddress := ethTransactions.Group("address")
	ethTransactionsByAddress.Get(":address", func(c *fiber.Ctx) error {
		return ethTransactionController.GetTransactionByAddressController(c)
	})

	htmx := app.Group("htmx")

	htmx.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})
}
