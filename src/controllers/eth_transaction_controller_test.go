package controllers

import (
	"eth-api/src/models"
	"eth-api/src/models/dto"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type EthTransactionServiceMock struct {
	mock.Mock
}

func (m *EthTransactionServiceMock) GetTransactionByHashService(hash string) (*models.EthTransaction, error) {
	args := m.Called(hash)
	return args.Get(0).(*models.EthTransaction), args.Error(1)
}

func (m *EthTransactionServiceMock) GetTransactionsByAddressService(address string) ([]models.EthTransaction, error) {
	args := m.Called(address)
	return args.Get(0).([]models.EthTransaction), args.Error(1)
}

func (m *EthTransactionServiceMock) FetchEthTransactionByHashFromAPI(hash string) (*dto.EthTransactionDTO, error) {
	args := m.Called(hash)
	return args.Get(0).(*dto.EthTransactionDTO), args.Error(1)
}

func (m *EthTransactionServiceMock) SaveTransactionService(transaction *models.EthTransaction) (*models.EthTransaction, error) {
	args := m.Called(transaction)
	return args.Get(0).(*models.EthTransaction), args.Error(1)
}

func TestGetTransactionByHashController(t *testing.T) {
	app := fiber.New()
	ethTransactionServiceMock := new(EthTransactionServiceMock)
	controller := EthTransactionController{EthTransactionService: ethTransactionServiceMock}
	app.Get("/transaction/:hash", controller.GetTransactionByHashController)

	testCases := []struct {
		description  string
		hash         string
		expectedCode int
		expectedBody string
		mockSetup    func(*EthTransactionServiceMock)
	}{
		{
			description:  "should return a transaction",
			hash:         "0x6813b6cea801bc68080b7dc843a8b46b55aa34cdb8679e9efb2cfd098fe02202",
			expectedCode: http.StatusOK,
			expectedBody: `{"blockHash":"","blockNumber":"1","chainId":"","from":"0xfbb1b73c4f0bda4f67dca266ce6ef42f520fbb98","hash":"0x6813b6cea801bc68080b7dc843a8b46b55aa34cdb8679e9efb2cfd098fe02202","to":"0x628c6dfc8e1d9afe5e66bacb2483d55e9f912881","value":"0x0","accessList":null}`,
			mockSetup: func(m *EthTransactionServiceMock) {
				m.On("GetTransactionByHashService", "0x6813b6cea801bc68080b7dc843a8b46b55aa34cdb8679e9efb2cfd098fe02202").Return(&models.EthTransaction{
					Hash:        "0x6813b6cea801bc68080b7dc843a8b46b55aa34cdb8679e9efb2cfd098fe02202",
					BlockNumber: "1",
					From:        "0xfbb1b73c4f0bda4f67dca266ce6ef42f520fbb98",
					To:          "0x628c6dfc8e1d9afe5e66bacb2483d55e9f912881",
					Value:       "0x0",
					Gas:         "21000",
					GasPrice:    "1000000000",
				}, nil)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			tc.mockSetup(ethTransactionServiceMock)

			req := httptest.NewRequest("GET", "/transaction/"+tc.hash, nil)
			resp, err := app.Test(req)
			assert.NoError(t, err, "Test failed with error")

			assert.Equal(t, tc.expectedCode, resp.StatusCode, "Expected status code to match")

			body, _ := io.ReadAll(resp.Body)
			assert.Equal(t, tc.expectedBody, string(body), "Expected body to match")
		})
	}
}

func TestGetTransactionByAddressController(t *testing.T) {
	app := fiber.New()
	ethTransactionServiceMock := new(EthTransactionServiceMock)
	controller := EthTransactionController{EthTransactionService: ethTransactionServiceMock}
	app.Get("/transaction/address/:address", controller.GetTransactionByAddressController)

	testCases := []struct {
		description  string
		address      string
		expectedCode int
		expectedBody string
		mockSetup    func(*EthTransactionServiceMock)
	}{
		{
			description:  "should return transactions",
			address:      "0xfbb1b73c4f0bda4f67dca266ce6ef42f520fbb98",
			expectedCode: http.StatusOK,
			expectedBody: `[{"blockHash":"","blockNumber":"","chainId":"","from":"0xfbb1b73c4f0bda4f67dca266ce6ef42f520fbb98","hash":"0x6813b6cea801bc68080b7dc843a8b46b55aa34cdb8679e9efb2cfd098fe02202","to":"0x628c6dfc8e1d9afe5e66bacb2483d55e9f912881","value":"0x0","accessList":null}]`,
			mockSetup: func(m *EthTransactionServiceMock) {
				m.On("GetTransactionsByAddressService", "0xfbb1b73c4f0bda4f67dca266ce6ef42f520fbb98").Return([]models.EthTransaction{
					{
						Hash:        "0x6813b6cea801bc68080b7dc843a8b46b55aa34cdb8679e9efb2cfd098fe02202",
						BlockNumber: "",
						From:        "0xfbb1b73c4f0bda4f67dca266ce6ef42f520fbb98",
						To:          "0x628c6dfc8e1d9afe5e66bacb2483d55e9f912881",
						Value:       "0x0",
						Gas:         "21000",
						GasPrice:    "1000000000",
					},
				}, nil)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			tc.mockSetup(ethTransactionServiceMock)

			req := httptest.NewRequest("GET", "/transaction/address/"+tc.address, nil)
			resp, err := app.Test(req)
			assert.NoError(t, err, "Test failed with error")

			assert.Equal(t, tc.expectedCode, resp.StatusCode, "Expected status code to match")

			body, _ := io.ReadAll(resp.Body)
			assert.Equal(t, tc.expectedBody, string(body), "Expected body to match")
		})
	}
}
