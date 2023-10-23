package dao

import (
	"eth-api/src/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"reflect"
	"testing"

	"gorm.io/driver/sqlite"
)

func TestGetBlockByNumber(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}

	db.AutoMigrate(&models.EthBlock{}, &models.Uncles{}, &models.Withdrawals{}, &models.EthTransaction{}, &models.AccessList{}, &models.StorageKey{})

	dao := NewEthBlockDAO(db)

	block := models.EthBlock{
		Number: "0x118ed01",
		Hash:   "0x1d674edee90409ccbcdcfac5e436f773daf3cf714d119d64ef405e263f1b4ccb",
	}
	db.Create(&block)

	testCases := []struct {
		name          string
		blockNumber   string
		expectedBlock *models.EthBlock
		expectedError error
	}{
		{
			name:          "Block exists",
			blockNumber:   "0x118ed01",
			expectedBlock: &block,
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			retrievedBlock, err := dao.GetBlockByNumber(tc.blockNumber)
			assert.Equal(t, tc.expectedError, err)

			if tc.expectedBlock != nil {
				if tc.expectedBlock.Transactions == nil {
					tc.expectedBlock.Transactions = []models.EthTransaction{}
				}
				assert.True(t, blocksEqual(tc.expectedBlock, retrievedBlock))
			} else {
				assert.Nil(t, retrievedBlock)
			}
		})

	}
}

func blocksEqual(b1, b2 *models.EthBlock) bool {
	b1.Timestamps = models.Timestamps{}
	b2.Timestamps = models.Timestamps{}
	return reflect.DeepEqual(b1, b2)
}
