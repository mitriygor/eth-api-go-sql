package services

import (
	"encoding/json"
	"eth-api/src/helpers/logger"
	"eth-api/src/models"
	"eth-api/src/models/dao"

	"eth-api/src/config"
	"eth-api/src/helpers"
	"eth-api/src/models/dto"
	"fmt"
	"io"
	"net/http"
)

// EthBlockService defines the interface for services that handle Ethereum blocks.
type EthBlockService interface {
	FetchEthBlockByNumberFromAPI(number string) (*dto.EthBlockDTO, error)
	GetBlockByNumberService(number string) (*models.EthBlock, error)
	SaveBlockService(block *models.EthBlock) (*models.EthBlock, error)
	GetLatestBlocksService(page int) ([]models.EthBlock, error)
}

// ethBlockService is the concrete implementation of EthBlockService.
type ethBlockService struct {
	ethBlockDAO dao.EthBlockDAO
	config      *config.Config
}

// NewEthBlockService creates a new instance of EthBlockService.
func NewEthBlockService(ethBlockDAO dao.EthBlockDAO, config *config.Config) EthBlockService {
	return &ethBlockService{
		ethBlockDAO: ethBlockDAO,
		config:      config,
	}
}

// FetchEthBlockByNumberFromAPI fetches an Ethereum block by its number from an external API.
func (e *ethBlockService) FetchEthBlockByNumberFromAPI(number string) (*dto.EthBlockDTO, error) {
	params := map[string]string{
		"module":  "proxy",
		"action":  "eth_getBlockByNumber",
		"tag":     number,
		"boolean": "true",
	}
	url := helpers.GetUrl(e.config.ApiUrl, e.config.ApiKey, params)

	resp, err := http.Get(fmt.Sprintf("%s", url))
	if err != nil {
		logger.Error("eth-api::ERROR::FetchEthBlockByNumberFromAPI::http.Get", "error", err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("eth-api::ERROR::FetchEthBlockByNumberFromAPI::io.ReadAll", "error", err)
		return nil, err
	}

	var response struct {
		Jsonrpc string          `json:"jsonrpc"`
		Id      int             `json:"id"`
		Result  dto.EthBlockDTO `json:"result"`
	}

	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		logger.Error("eth-api::ERROR::FetchEthBlockByNumberFromAPI::json.Unmarshal", "error", err)
		return nil, err
	}

	return &response.Result, nil
}

// GetBlockByNumberService retrieves an Ethereum block by its number from the database.
func (e *ethBlockService) GetBlockByNumberService(number string) (*models.EthBlock, error) {
	block, err := e.ethBlockDAO.GetBlockByNumber(number)
	if err != nil {
		logger.Error("eth-api::ERROR::GetBlockByNumberService::GetBlockByNumber", "error", err)
		return nil, err
	}

	return block, nil
}

// GetLatestBlocksService retrieves the latest Ethereum blocks from the database.
func (e *ethBlockService) GetLatestBlocksService(page int) ([]models.EthBlock, error) {
	blocks, err := e.ethBlockDAO.GetLatestBlocks(page, e.config.Limit)
	if err != nil {
		logger.Error("eth-api::ERROR::GetLatestBlocksService::GetLatestBlocks", "error", err)
		return nil, err
	}

	return blocks, nil
}

// SaveBlockService saves an Ethereum block to the database.
func (e *ethBlockService) SaveBlockService(block *models.EthBlock) (*models.EthBlock, error) {
	savedBlock, err := e.ethBlockDAO.SaveBlock(block)
	if err != nil {
		logger.Error("eth-api::ERROR::SaveBlockService::SaveBlock", "error", err)
		return nil, err
	}

	return savedBlock, nil
}
