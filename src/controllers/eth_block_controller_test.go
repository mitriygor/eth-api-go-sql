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

type EthBlockServiceMock struct {
	mock.Mock
}

func (m *EthBlockServiceMock) FetchEthBlockByNumberFromAPI(number string) (*dto.EthBlockDTO, error) {
	args := m.Called(number)
	return args.Get(0).(*dto.EthBlockDTO), args.Error(1)
}

func (m *EthBlockServiceMock) SaveBlockService(block *models.EthBlock) (*models.EthBlock, error) {
	args := m.Called(block)
	return args.Get(0).(*models.EthBlock), args.Error(1)
}

func (m *EthBlockServiceMock) GetBlockByNumberService(number string) (*models.EthBlock, error) {
	args := m.Called(number)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.EthBlock), args.Error(1)
}

func (m *EthBlockServiceMock) GetLatestBlocksService(page int) ([]models.EthBlock, error) {
	args := m.Called(page)
	return args.Get(0).([]models.EthBlock), args.Error(1)
}

func TestGetBlockByNumberController(t *testing.T) {
	app := fiber.New()
	ethBlockServiceMock := new(EthBlockServiceMock)
	controller := EthBlockController{EthBlockService: ethBlockServiceMock}
	app.Get("/block/:number", controller.GetBlockByNumberController)

	testCases := []struct {
		description  string
		number       string
		expectedCode int
		expectedBody string
		mockSetup    func(*EthBlockServiceMock)
	}{
		{
			description:  "should return a block",
			number:       "0x118ed01",
			expectedCode: http.StatusOK,
			expectedBody: `{"hash":"0x1d674edee90409ccbcdcfac5e436f773daf3cf714d119d64ef405e263f1b4ccb","number":"0x118ed01","transactions":[]}`,
			mockSetup: func(m *EthBlockServiceMock) {
				m.On("FetchEthBlockByNumberFromAPI", "0x118ed01").Return(&dto.EthBlockDTO{}, nil)
				m.On("GetBlockByNumberService", "0x118ed01").Return(&models.EthBlock{
					Number: "0x118ed01",
					Hash:   "0x1d674edee90409ccbcdcfac5e436f773daf3cf714d119d64ef405e263f1b4ccb",
				}, nil)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			tc.mockSetup(ethBlockServiceMock)

			req := httptest.NewRequest("GET", "/block/"+tc.number, nil)
			resp, err := app.Test(req)
			assert.NoError(t, err, "Test failed with error")

			assert.Equal(t, tc.expectedCode, resp.StatusCode, "Expected status code to match")

			body, _ := io.ReadAll(resp.Body)
			assert.Equal(t, tc.expectedBody, string(body), "Expected body to match")
		})
	}
}

func TestGetLatestBlocksController(t *testing.T) {
	app := fiber.New()
	ethBlockServiceMock := new(EthBlockServiceMock)
	controller := EthBlockController{EthBlockService: ethBlockServiceMock}
	app.Get("/blocks", controller.GetLatestBlocksController)

	testCases := []struct {
		description  string
		page         string
		expectedCode int
		expectedBody string
		mockSetup    func(*EthBlockServiceMock)
	}{
		{
			description:  "should return latest blocks",
			page:         "0x118ed01",
			expectedCode: http.StatusOK,
			expectedBody: `[{"hash":"0x1d674edee90409ccbcdcfac5e436f773daf3cf714d119d64ef405e263f1b4ccb","number":"0x118ed01","transactions":[]}]`,
			mockSetup: func(m *EthBlockServiceMock) {
				m.On("GetLatestBlocksService", 1).Return([]models.EthBlock{
					{
						Number: "0x118ed01",
						Hash:   "0x1d674edee90409ccbcdcfac5e436f773daf3cf714d119d64ef405e263f1b4ccb",
					},
				}, nil)
			},
		},
		{
			description:  "should handle invalid page query",
			page:         "invalid",
			expectedCode: http.StatusOK,
			expectedBody: `[{"hash":"0x1d674edee90409ccbcdcfac5e436f773daf3cf714d119d64ef405e263f1b4ccb","number":"0x118ed01","transactions":[]}]`,
			mockSetup: func(m *EthBlockServiceMock) {
				m.On("GetLatestBlocksService", 1).Return([]models.EthBlock{
					{
						Number: "0x118ed01",
						Hash:   "0x1d674edee90409ccbcdcfac5e436f773daf3cf714d119d64ef405e263f1b4ccb",
					},
				}, nil)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			tc.mockSetup(ethBlockServiceMock)

			req := httptest.NewRequest("GET", "/blocks?page="+tc.page, nil)
			resp, err := app.Test(req)
			assert.NoError(t, err, "Test failed with error")

			assert.Equal(t, tc.expectedCode, resp.StatusCode, "Expected status code to match")

			body, _ := io.ReadAll(resp.Body)
			assert.Equal(t, tc.expectedBody, string(body), "Expected body to match")
		})
	}
}
