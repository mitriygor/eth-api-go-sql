package dao

import (
	"eth-api/src/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func TestGetTransactionByHash(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}

	db.AutoMigrate(&models.EthTransaction{})

	dao := NewEthTransactionDAO(db)

	transaction := models.EthTransaction{
		Hash: "0x6813b6cea801bc68080b7dc843a8b46b55aa34cdb8679e9efb2cfd098fe02202",
	}
	db.Create(&transaction)

	testCases := []struct {
		name            string
		transactionHash string
		expectedTx      *models.EthTransaction
		expectedError   error
	}{
		{
			name:            "Transaction exists",
			transactionHash: "0x6813b6cea801bc68080b7dc843a8b46b55aa34cdb8679e9efb2cfd098fe02202",
			expectedTx:      &transaction,
			expectedError:   nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			retrievedTx, err := dao.GetTransactionByHash(tc.transactionHash)
			assert.Equal(t, tc.expectedError, err)

			if tc.expectedTx != nil {
				assert.Equal(t, tc.expectedTx.Hash, retrievedTx.Hash)
			} else {
				assert.Nil(t, retrievedTx)
			}
		})
	}
}
