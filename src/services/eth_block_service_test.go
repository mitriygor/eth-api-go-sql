package services

import (
	"eth-api/src/config"
	"eth-api/src/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type EthBlockDAOMock struct {
	mock.Mock
}

func (m *EthBlockDAOMock) GetLatestBlockNumber() (string, error) {
	args := m.Called()
	result, _ := args.Get(0).(string)
	return result, args.Error(1)
}

func (m *EthBlockDAOMock) GetBlockByNumber(number string) (*models.EthBlock, error) {
	args := m.Called(number)
	result, _ := args.Get(0).(*models.EthBlock)
	return result, args.Error(1)
}

func (m *EthBlockDAOMock) SaveBlock(block *models.EthBlock) (*models.EthBlock, error) {
	args := m.Called(block)
	result, _ := args.Get(0).(*models.EthBlock)
	return result, args.Error(1)
}

func (m *EthBlockDAOMock) GetLatestBlocks(page int, limit int) ([]models.EthBlock, error) {
	args := m.Called(page, limit)
	result, _ := args.Get(0).([]models.EthBlock)
	return result, args.Error(1)
}

func TestGetBlockByNumberService(t *testing.T) {
	blockDAO := new(EthBlockDAOMock)
	c := &config.Config{}

	service := NewEthBlockService(blockDAO, c)

	testCases := []struct {
		description   string
		number        string
		expectedBlock *models.EthBlock
		expectedError error
		mockSetup     func(*EthBlockDAOMock)
	}{
		{
			description:   "should return a block",
			number:        "0x118ed01",
			expectedBlock: &models.EthBlock{Number: "0x118ed01"},
			expectedError: nil,
			mockSetup: func(m *EthBlockDAOMock) {
				m.On("GetBlockByNumber", "0x118ed01").Return(&models.EthBlock{Number: "0x118ed01"}, nil)
			},
		},
		{
			description:   "should return an error if not found",
			number:        "0x118cba5",
			expectedBlock: nil,
			expectedError: gorm.ErrRecordNotFound,
			mockSetup: func(m *EthBlockDAOMock) {
				m.On("GetBlockByNumber", "0x118cba5").Return(nil, gorm.ErrRecordNotFound)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			tc.mockSetup(blockDAO)

			block, err := service.GetBlockByNumberService(tc.number)

			assert.Equal(t, tc.expectedBlock, block, "Expected block to match")
			assert.Equal(t, tc.expectedError, err, "Expected error to match")
		})
	}
}
