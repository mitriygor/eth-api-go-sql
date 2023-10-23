package services

import (
	"encoding/json"
	"eth-api/src/config"
	"eth-api/src/helpers"
	"eth-api/src/helpers/logger"
	"eth-api/src/models/dao"
	"eth-api/src/models/dto"
	"fmt"
	"github.com/robfig/cron/v3"
	"io"
	"net/http"
	"sync"
	"time"
)

// CronService defines the interface for services that handle scheduled tasks.
type CronService interface {
	FetchLatestEthBlockNumberFromAPI() (string, error)
	GetLatestEthBlockNumberService() (string, error)
	FetchAndSaveEthBlocks(latestNum string, currNum string)
	RunCron()
}

// cronService is the concrete implementation of CronService.
type cronService struct {
	ethBlockDAO     dao.EthBlockDAO
	ethBlockService EthBlockService
	config          *config.Config
}

// NewCronService creates a new instance of CronService.
func NewCronService(ethBlockDAO dao.EthBlockDAO, ethBlockService EthBlockService, config *config.Config) CronService {
	return &cronService{
		ethBlockDAO:     ethBlockDAO,
		ethBlockService: ethBlockService,
		config:          config,
	}
}

// RunCron starts the cron job that fetches the latest Ethereum block number and logs the difference with the current block number in the database.
func (e *cronService) RunCron() {
	c := cron.New()

	c.AddFunc("@every 10s", func() {
		latestNum, err := e.FetchLatestEthBlockNumberFromAPI()
		if err == nil {
			currNum, err := e.GetLatestEthBlockNumberService()
			if err == nil {
				latestIntNum, _ := helpers.HexToInt(latestNum)
				currIntNum, _ := helpers.HexToInt(currNum)
				fmt.Printf("Latest Number: %v; Current Number: %v; Difference: %v;\n", latestIntNum, currIntNum, latestIntNum-currIntNum)
			}
		}
	})

	c.Start()
}

// FetchLatestEthBlockNumberFromAPI fetches the latest Ethereum block number from an external API.
func (e *cronService) FetchLatestEthBlockNumberFromAPI() (string, error) {
	params := map[string]string{
		"module":  "proxy",
		"action":  "eth_blockNumber",
		"boolean": "true",
	}

	url := helpers.GetUrl(e.config.ApiUrl, e.config.ApiKey, params)

	resp, err := http.Get(fmt.Sprintf("%s", url))
	if err != nil {
		logger.Error("eth-api::ERROR::FetchLatestEthBlockNumberFromAPI::http.Get", "error", err)
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("eth-api::ERROR::FetchLatestEthBlockNumberFromAPI::io.ReadAll", "error", err)
		return "", err
	}

	var response struct {
		Jsonrpc string `json:"jsonrpc"`
		Id      int    `json:"id"`
		Result  string `json:"result"`
	}

	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		logger.Error("eth-api::ERROR::FetchLatestEthBlockNumberFromAPI::json.Unmarshal", "error", err)
		return "", err
	}

	return response.Result, nil
}

// GetLatestEthBlockNumberService retrieves the latest Ethereum block number from the database.
func (e *cronService) GetLatestEthBlockNumberService() (string, error) {
	latestBlockNumber, err := e.ethBlockDAO.GetLatestBlockNumber()
	if err != nil {
		logger.Error("eth-api::ERROR::GetLatestEthBlockNumberService", "error", err)
		return "", err
	}
	return latestBlockNumber, nil
}

// FetchAndSaveEthBlocks fetches Ethereum blocks within a given range and saves them to the database.
func (e *cronService) FetchAndSaveEthBlocks(latestNum string, currNum string) {
	latestBlockNumber, err := helpers.HexToInt(latestNum)
	if err != nil {
		logger.Error("eth-api::ERROR::FetchAndSaveEthBlocks::helpers.HexToInt", "error", err)
		return
	}

	currentBlockNumber, err := helpers.HexToInt(currNum)
	if err != nil {
		logger.Error("eth-api::ERROR::FetchAndSaveEthBlocks::helpers.HexToInt", "error", err)
		return
	}

	var wg sync.WaitGroup

	for i := currentBlockNumber + 1; i <= latestBlockNumber; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			hex := helpers.IntToHex(i)

			dtoBlock, err := e.ethBlockService.FetchEthBlockByNumberFromAPI(hex)
			if err != nil {
				logger.Error("eth-api::ERROR::FetchAndSaveEthBlocks::FetchEthBlockByNumberFromAPI", "error", err)
			}

			ethBlock := dto.ConvertDTOToEthBlock(dtoBlock)
			_, err = e.ethBlockDAO.SaveBlock(ethBlock)
			if err != nil {
				logger.Error("eth-api::ERROR::FetchAndSaveEthBlocks::SaveBlock", "error", err)
			}

			time.Sleep(1 * time.Second)
		}(i)
	}

	wg.Wait()
}
