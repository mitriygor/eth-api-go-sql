package controllers

import (
	"errors"
	"eth-api/src/helpers/logger"
	"eth-api/src/models/dto"
	"eth-api/src/services"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// EthTransactionController handles requests related to Ethereum transactions.
type EthTransactionController struct {
	EthTransactionService services.EthTransactionService
}

// NewEthTransactionController creates a new instance of EthTransactionController.
func NewEthTransactionController(service services.EthTransactionService) *EthTransactionController {
	return &EthTransactionController{
		EthTransactionService: service,
	}
}

// GetTransactionByHashController handles the request to get an Ethereum transaction by its hash.
func (e *EthTransactionController) GetTransactionByHashController(c *fiber.Ctx) error {
	hash := c.Params("hash")
	foundTransaction, err := e.EthTransactionService.GetTransactionByHashService(hash)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			dtoTransaction, err := e.EthTransactionService.FetchEthTransactionByHashFromAPI(hash)
			if err != nil {
				logger.Error("eth-api::ERROR::GetTransactionByHashController::FetchEthTransactionByHashFromAPI", "error", err)
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "Failed to fetch data from external API",
				})
			}

			foundTransaction = dto.ConvertDTOToEthTransaction(dtoTransaction)
			go func() {
				_, err = e.EthTransactionService.SaveTransactionService(foundTransaction)
				if err != nil {
					logger.Error("eth-api::ERROR::GetTransactionByHashController::SaveTransactionService", "error", err)
				}
			}()
		} else {
			logger.Error("eth-api::ERROR::GetTransactionByHashController::GetTransactionByHashService", "error", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Internal server error",
			})
		}
	}

	resTransaction := dto.ConvertEthTransactionToLightDTO(foundTransaction)

	return c.JSON(resTransaction)
}

// GetTransactionByAddressController handles the request to get Ethereum transactions associated with a given address.
func (e *EthTransactionController) GetTransactionByAddressController(c *fiber.Ctx) error {
	address := c.Params("address")

	foundTransactions, err := e.EthTransactionService.GetTransactionsByAddressService(address)
	if err != nil {
		logger.Error("eth-api::ERROR::GetTransactionByAddressController::GetTransactionsByAddressService", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	var resTransactions []*dto.EthTransactionLightDTO

	for _, transaction := range foundTransactions {
		resTransactions = append(resTransactions, dto.ConvertEthTransactionToLightDTO(&transaction))
	}

	return c.JSON(resTransactions)
}
