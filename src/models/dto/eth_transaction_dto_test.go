package dto

import (
	"eth-api/src/models"
	"reflect"
	"testing"
)

func TestConvertDTOToEthTransaction(t *testing.T) {
	tests := []struct {
		name string
		dto  *EthTransactionDTO
		want *models.EthTransaction
	}{
		{
			name: "basic transaction",
			dto: &EthTransactionDTO{
				BlockHash:            "0x1d674edee90409ccbcdcfac5e436f773daf3cf714d119d64ef405e263f1b4ccb",
				BlockNumber:          "0x118ed01",
				ChainId:              "0x1",
				From:                 "0xfbb1b73c4f0bda4f67dca266ce6ef42f520fbb98",
				Gas:                  "0x249f0",
				GasPrice:             "0x6fc23ac00",
				Hash:                 "0x6813b6cea801bc68080b7dc843a8b46b55aa34cdb8679e9efb2cfd098fe02202",
				Input:                "0xinput",
				MaxFeePerGas:         "1",
				MaxPriorityFeePerGas: "1",
				Nonce:                "0xae1a17",
				R:                    "0xf1391e146f0d95efdc2f570281bc4a203554f08ed73c1cb021acde37b8e24f97",
				S:                    "0x71820d2288a60dcd00daa689d6cca0b762379dcca793442372adc1f336bcd185",
				To:                   "0x628c6dfc8e1d9afe5e66bacb2483d55e9f912881",
				TransactionIndex:     "0x0",
				Type:                 "0x0",
				V:                    "0x25",
				Value:                "0x0",
				AccessList: []AccessListDTO{
					{
						Address:     "0xaddress",
						StorageKeys: []string{"0xstoragekey1", "0xstoragekey2"},
					},
				},
			},
			want: &models.EthTransaction{
				BlockHash:            "0x1d674edee90409ccbcdcfac5e436f773daf3cf714d119d64ef405e263f1b4ccb",
				BlockNumber:          "0x118ed01",
				ChainId:              "0x1",
				From:                 "0xfbb1b73c4f0bda4f67dca266ce6ef42f520fbb98",
				Gas:                  "0x249f0",
				GasPrice:             "0x6fc23ac00",
				Hash:                 "0x6813b6cea801bc68080b7dc843a8b46b55aa34cdb8679e9efb2cfd098fe02202",
				Input:                "0xinput",
				MaxFeePerGas:         "1",
				MaxPriorityFeePerGas: "1",
				Nonce:                "0xae1a17",
				R:                    "0xf1391e146f0d95efdc2f570281bc4a203554f08ed73c1cb021acde37b8e24f97",
				S:                    "0x71820d2288a60dcd00daa689d6cca0b762379dcca793442372adc1f336bcd185",
				To:                   "0x628c6dfc8e1d9afe5e66bacb2483d55e9f912881",
				TransactionIndex:     "0x0",
				Type:                 "0x0",
				V:                    "0x25",
				Value:                "0x0",
				AccessList: []models.AccessList{
					{
						Address: "0xaddress",
						StorageKeys: []models.StorageKey{
							{StorageKey: "0xstoragekey1"},
							{StorageKey: "0xstoragekey2"},
						},
					},
				},
			},
		},
		{
			name: "empty transaction",
			dto:  &EthTransactionDTO{},
			want: &models.EthTransaction{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConvertDTOToEthTransaction(tt.dto)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertDTOToEthTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertEthTransactionToDTO(t *testing.T) {
	tests := []struct {
		name        string
		transaction *models.EthTransaction
		want        *EthTransactionDTO
	}{
		{
			name: "basic transaction",
			transaction: &models.EthTransaction{
				BlockHash:            "0x1d674edee90409ccbcdcfac5e436f773daf3cf714d119d64ef405e263f1b4ccb",
				BlockNumber:          "0x118ed01",
				ChainId:              "0x1",
				From:                 "0xfbb1b73c4f0bda4f67dca266ce6ef42f520fbb98",
				Gas:                  "0x249f0",
				GasPrice:             "0x6fc23ac00",
				Hash:                 "0x6813b6cea801bc68080b7dc843a8b46b55aa34cdb8679e9efb2cfd098fe02202",
				Input:                "0xinput",
				MaxFeePerGas:         "1",
				MaxPriorityFeePerGas: "1",
				Nonce:                "0xae1a17",
				R:                    "0xf1391e146f0d95efdc2f570281bc4a203554f08ed73c1cb021acde37b8e24f97",
				S:                    "0x71820d2288a60dcd00daa689d6cca0b762379dcca793442372adc1f336bcd185",
				To:                   "0x628c6dfc8e1d9afe5e66bacb2483d55e9f912881",
				TransactionIndex:     "0x0",
				Type:                 "0x0",
				V:                    "0x25",
				Value:                "0x0",
				AccessList: []models.AccessList{
					{
						Address: "0xaddress",
						StorageKeys: []models.StorageKey{
							{StorageKey: "0xstoragekey1"},
							{StorageKey: "0xstoragekey2"},
						},
					},
				},
			},
			want: &EthTransactionDTO{
				BlockHash:            "0x1d674edee90409ccbcdcfac5e436f773daf3cf714d119d64ef405e263f1b4ccb",
				BlockNumber:          "0x118ed01",
				ChainId:              "0x1",
				From:                 "0xfbb1b73c4f0bda4f67dca266ce6ef42f520fbb98",
				Gas:                  "0x249f0",
				GasPrice:             "0x6fc23ac00",
				Hash:                 "0x6813b6cea801bc68080b7dc843a8b46b55aa34cdb8679e9efb2cfd098fe02202",
				Input:                "0xinput",
				MaxFeePerGas:         "1",
				MaxPriorityFeePerGas: "1",
				Nonce:                "0xae1a17",
				R:                    "0xf1391e146f0d95efdc2f570281bc4a203554f08ed73c1cb021acde37b8e24f97",
				S:                    "0x71820d2288a60dcd00daa689d6cca0b762379dcca793442372adc1f336bcd185",
				To:                   "0x628c6dfc8e1d9afe5e66bacb2483d55e9f912881",
				TransactionIndex:     "0x0",
				Type:                 "0x0",
				V:                    "0x25",
				Value:                "0x0",
				AccessList: []AccessListDTO{
					{
						Address:     "0xaddress",
						StorageKeys: []string{"0xstoragekey1", "0xstoragekey2"},
					},
				},
			},
		},
		{
			name:        "empty transaction",
			transaction: &models.EthTransaction{},
			want:        &EthTransactionDTO{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConvertEthTransactionToDTO(tt.transaction)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertEthTransactionToDTO() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertEthTransactionToLightDTO(t *testing.T) {
	tests := []struct {
		name        string
		transaction *models.EthTransaction
		want        *EthTransactionLightDTO
	}{
		{
			name: "basic transaction",
			transaction: &models.EthTransaction{
				BlockHash:   "0x1d674edee90409ccbcdcfac5e436f773daf3cf714d119d64ef405e263f1b4ccb",
				BlockNumber: "0x118ed01",
				ChainId:     "0x1",
				From:        "0xfbb1b73c4f0bda4f67dca266ce6ef42f520fbb98",
				Hash:        "0x6813b6cea801bc68080b7dc843a8b46b55aa34cdb8679e9efb2cfd098fe02202",
				To:          "0x628c6dfc8e1d9afe5e66bacb2483d55e9f912881",
				Value:       "0x0",
				AccessList: []models.AccessList{
					{
						Address: "0xaddress",
						StorageKeys: []models.StorageKey{
							{StorageKey: "0xstoragekey1"},
							{StorageKey: "0xstoragekey2"},
						},
					},
				},
			},
			want: &EthTransactionLightDTO{
				BlockHash:   "0x1d674edee90409ccbcdcfac5e436f773daf3cf714d119d64ef405e263f1b4ccb",
				BlockNumber: "0x118ed01",
				ChainId:     "0x1",
				From:        "0xfbb1b73c4f0bda4f67dca266ce6ef42f520fbb98",
				Hash:        "0x6813b6cea801bc68080b7dc843a8b46b55aa34cdb8679e9efb2cfd098fe02202",
				To:          "0x628c6dfc8e1d9afe5e66bacb2483d55e9f912881",
				Value:       "0x0",
			},
		},
		{
			name:        "empty transaction",
			transaction: &models.EthTransaction{},
			want:        &EthTransactionLightDTO{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConvertEthTransactionToLightDTO(tt.transaction)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertEthTransactionToLightDTO() = %v, want %v", got, tt.want)
			}
		})
	}
}
