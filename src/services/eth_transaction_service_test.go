package services

import (
	"eth-api/src/config"
	"eth-api/src/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type EthTransactionDAOMock struct {
	mock.Mock
}

func (m *EthTransactionDAOMock) GetTransactionByHash(hash string) (*models.EthTransaction, error) {
	args := m.Called(hash)
	result, _ := args.Get(0).(*models.EthTransaction)
	return result, args.Error(1)
}
func (m *EthTransactionDAOMock) SaveTransaction(transaction *models.EthTransaction) (*models.EthTransaction, error) {
	args := m.Called(transaction)
	return args.Get(0).(*models.EthTransaction), args.Error(1)
}

func (m *EthTransactionDAOMock) GetTransactionsByAddress(address string) ([]models.EthTransaction, error) {
	args := m.Called(address)
	return args.Get(0).([]models.EthTransaction), args.Error(1)
}

func TestGetTransactionByHashService(t *testing.T) {
	transactionDAO := new(EthTransactionDAOMock)
	c := &config.Config{}

	service := NewEthTransactionService(transactionDAO, c, nil)

	testCases := []struct {
		description   string
		hash          string
		expectedTx    *models.EthTransaction
		expectedError error
		mockSetup     func(*EthTransactionDAOMock)
	}{
		{
			description:   "should return a transaction",
			hash:          "0x6813b6cea801bc68080b7dc843a8b46b55aa34cdb8679e9efb2cfd098fe02202",
			expectedTx:    &models.EthTransaction{Hash: "0x6813b6cea801bc68080b7dc843a8b46b55aa34cdb8679e9efb2cfd098fe02202"},
			expectedError: nil,
			mockSetup: func(m *EthTransactionDAOMock) {
				m.On("GetTransactionByHash", "0x6813b6cea801bc68080b7dc843a8b46b55aa34cdb8679e9efb2cfd098fe02202").Return(&models.EthTransaction{Hash: "0x6813b6cea801bc68080b7dc843a8b46b55aa34cdb8679e9efb2cfd098fe02202"}, nil)
			},
		},
		{
			description:   "should return an error if not found",
			hash:          "0xc8985cecf7179ac3c2e0e26ee25c23f8343a920bd3387e1de9b46583d5c3049f",
			expectedTx:    nil,
			expectedError: gorm.ErrRecordNotFound,
			mockSetup: func(m *EthTransactionDAOMock) {
				m.On("GetTransactionByHash", "0xc8985cecf7179ac3c2e0e26ee25c23f8343a920bd3387e1de9b46583d5c3049f").Return(nil, gorm.ErrRecordNotFound)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			tc.mockSetup(transactionDAO)

			tx, err := service.GetTransactionByHashService(tc.hash)

			assert.Equal(t, tc.expectedTx, tx, "Expected transaction to match")
			assert.Equal(t, tc.expectedError, err, "Expected error to match")
		})
	}
}

func TestGetTransactionsByAddressService(t *testing.T) {
	ethTransactionDAO := new(EthTransactionDAOMock)
	service := ethTransactionService{ethTransactionDAO: ethTransactionDAO}

	testCases := []struct {
		description   string
		address       string
		expectedTxs   []models.EthTransaction
		expectedError error
		mockSetup     func(*EthTransactionDAOMock)
	}{
		{
			description:   "should return transactions",
			address:       "0x6813b6cea801bc68080b7dc843a8b46b55aa34cdb8679e9efb2cfd098fe02202",
			expectedTxs:   []models.EthTransaction{{Hash: "0x6813b6cea801bc68080b7dc843a8b46b55aa34cdb8679e9efb2cfd098fe02202"}},
			expectedError: nil,
			mockSetup: func(m *EthTransactionDAOMock) {
				m.On("GetTransactionsByAddress", "0x6813b6cea801bc68080b7dc843a8b46b55aa34cdb8679e9efb2cfd098fe02202").Return([]models.EthTransaction{{Hash: "0x6813b6cea801bc68080b7dc843a8b46b55aa34cdb8679e9efb2cfd098fe02202"}}, nil)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			tc.mockSetup(ethTransactionDAO)

			txs, err := service.GetTransactionsByAddressService(tc.address)

			assert.Equal(t, tc.expectedTxs, txs)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
