package services

import (
	"encoding/json"
	"eth-api/src/config"
	"eth-api/src/helpers"
	"eth-api/src/helpers/logger"
	"eth-api/src/models"
	"eth-api/src/models/dao"
	"eth-api/src/models/dto"
	"fmt"
	"io"
	"net/http"
)

// EthTransactionService defines the interface for services that handle Ethereum transactions.
type EthTransactionService interface {
	GetTransactionByHashService(hash string) (*models.EthTransaction, error)
	GetTransactionsByAddressService(address string) ([]models.EthTransaction, error)
	FetchEthTransactionByHashFromAPI(hash string) (*dto.EthTransactionDTO, error)
	SaveTransactionService(transaction *models.EthTransaction) (*models.EthTransaction, error)
}

// ethTransactionService is the concrete implementation of EthTransactionService.
type ethTransactionService struct {
	ethTransactionDAO dao.EthTransactionDAO
	ethBlockService   EthBlockService
	config            *config.Config
}

// NewEthTransactionService creates a new instance of EthTransactionService.
func NewEthTransactionService(ethTransactionDAO dao.EthTransactionDAO, config *config.Config, ethBlockService EthBlockService) EthTransactionService {
	return &ethTransactionService{
		ethTransactionDAO: ethTransactionDAO,
		ethBlockService:   ethBlockService,
		config:            config,
	}
}

// GetTransactionByHashService retrieves an Ethereum transaction by its hash from the database.
func (e *ethTransactionService) GetTransactionByHashService(hash string) (*models.EthTransaction, error) {
	transaction, err := e.ethTransactionDAO.GetTransactionByHash(hash)
	if err != nil {
		logger.Error("eth-api::ERROR::GetTransactionByHashService", "error", err)
		return nil, err
	}

	return transaction, nil
}

// GetTransactionsByAddressService retrieves all Ethereum transactions associated with a specific address from the database.
func (e *ethTransactionService) GetTransactionsByAddressService(address string) ([]models.EthTransaction, error) {
	transactions, err := e.ethTransactionDAO.GetTransactionsByAddress(address)
	if err != nil {
		logger.Error("eth-api::ERROR::GetTransactionsByAddressService", "error", err)
		return nil, err
	}

	return transactions, nil
}

// FetchEthTransactionByHashFromAPI fetches an Ethereum transaction by its hash from an external API.
func (e *ethTransactionService) FetchEthTransactionByHashFromAPI(hash string) (*dto.EthTransactionDTO, error) {
	params := map[string]string{
		"module": "proxy",
		"action": "eth_getTransactionByHash",
		"txhash": hash,
	}
	url := helpers.GetUrl(e.config.ApiUrl, e.config.ApiKey, params)

	resp, err := http.Get(fmt.Sprintf("%s", url))
	if err != nil {
		logger.Error("eth-api::ERROR::FetchEthTransactionByHashFromAPI::http.Get", "error", err)
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
		logger.Error("eth-api::ERROR::FetchEthTransactionByHashFromAPI::io.ReadAll", "error", err)
		return nil, err
	}

	var response struct {
		Jsonrpc string                `json:"jsonrpc"`
		Id      int                   `json:"id"`
		Result  dto.EthTransactionDTO `json:"result"`
	}

	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		logger.Error("eth-api::ERROR::FetchEthTransactionByHashFromAPI::json.Unmarshal", "error", err)
		return nil, err
	}

	return &response.Result, nil
}

// SaveTransactionService saves an Ethereum transaction to the database.
func (e *ethTransactionService) SaveTransactionService(transaction *models.EthTransaction) (*models.EthTransaction, error) {
	block, err := e.ethBlockService.GetBlockByNumberService(transaction.BlockNumber)
	if err != nil {
		dtoBlock, err := e.ethBlockService.FetchEthBlockByNumberFromAPI(transaction.BlockNumber)
		if err != nil {
			logger.Error("eth-api::ERROR::SaveTransactionService::FetchEthBlockByNumberFromAPI", "error", err)
			return nil, err
		}
		block = dto.ConvertDTOToEthBlock(dtoBlock)
		_, err = e.ethBlockService.SaveBlockService(block)
		if err != nil {
			logger.Error("eth-api::ERROR::SaveTransactionService::SaveBlockService", "error", err)
			return nil, err
		}

		return transaction, err
	}

	transaction.EthBlockId = block.Id
	transaction, err = e.ethTransactionDAO.SaveTransaction(transaction)

	return transaction, err
}
