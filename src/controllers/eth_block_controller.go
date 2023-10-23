package controllers

import (
	"errors"
	"eth-api/src/helpers/logger"
	"eth-api/src/models/dto"
	"eth-api/src/services"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"strconv"
)

// EthBlockController handles requests related to Ethereum blocks.
type EthBlockController struct {
	EthBlockService services.EthBlockService
}

// NewEthBlockController creates a new instance of EthBlockController.
func NewEthBlockController(service services.EthBlockService) *EthBlockController {
	return &EthBlockController{
		EthBlockService: service,
	}
}

// GetBlockByNumberController handles the request to get an Ethereum block by its number.
func (e *EthBlockController) GetBlockByNumberController(c *fiber.Ctx) error {
	number := c.Params("number")
	foundBlock, err := e.EthBlockService.GetBlockByNumberService(number)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			dtoBlock, err := e.EthBlockService.FetchEthBlockByNumberFromAPI(number)
			if err != nil {
				logger.Error("eth-api::ERROR::GetBlockByNumberController::FetchEthBlockByNumberFromAPI", "error", err)
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "Failed to fetch data from external API",
				})
			}

			foundBlock = dto.ConvertDTOToEthBlock(dtoBlock)
			go func() {
				_, err = e.EthBlockService.SaveBlockService(foundBlock)
				if err != nil {
					logger.Error("eth-api::ERROR::GetBlockByNumberController::SaveBlockService", "error", err)
				}
			}()
		} else {
			logger.Error("eth-api::ERROR::GetBlockByNumberController::GetBlockByNumberService", "error", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Internal server error",
			})
		}
	}

	resBlock := dto.ConvertEthBlockToLightDTO(foundBlock)

	return c.JSON(resBlock)
}

// GetLatestBlocksController handles the request to get the latest Ethereum blocks.
func (e *EthBlockController) GetLatestBlocksController(c *fiber.Ctx) error {
	pageStr := c.Query("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	latestBlocks, err := e.EthBlockService.GetLatestBlocksService(page)
	if err != nil {
		logger.Error("eth-api::ERROR::GetLatestBlocksController::GetLatestBlocksService", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	var resBlocks []*dto.EthBlockLightDTO

	for _, blocks := range latestBlocks {
		resBlocks = append(resBlocks, dto.ConvertEthBlockToLightDTO(&blocks))
	}

	return c.JSON(resBlocks)
}
