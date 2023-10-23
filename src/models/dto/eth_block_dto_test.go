package dto

import (
	"eth-api/src/models"
	"reflect"
	"testing"
)

func TestConvertEthBlockToLightDTO(t *testing.T) {
	tests := []struct {
		name     string
		input    *models.EthBlock
		expected *EthBlockLightDTO
	}{
		{
			name: "simple test",
			input: &models.EthBlock{
				Hash:   "0x1d674edee90409ccbcdcfac5e436f773daf3cf714d119d64ef405e263f1b4ccb",
				Number: "0x118ed01",
				Transactions: []models.EthTransaction{
					{Hash: "0x6813b6cea801bc68080b7dc843a8b46b55aa34cdb8679e9efb2cfd098fe02202"},
					{Hash: "0xc8985cecf7179ac3c2e0e26ee25c23f8343a920bd3387e1de9b46583d5c3049f"},
				},
			},
			expected: &EthBlockLightDTO{
				Hash:         "0x1d674edee90409ccbcdcfac5e436f773daf3cf714d119d64ef405e263f1b4ccb",
				Number:       "0x118ed01",
				Transactions: []string{"0x6813b6cea801bc68080b7dc843a8b46b55aa34cdb8679e9efb2cfd098fe02202", "0xc8985cecf7179ac3c2e0e26ee25c23f8343a920bd3387e1de9b46583d5c3049f"},
			},
		},
		{
			name: "test with no transactions",
			input: &models.EthBlock{
				Hash:         "0x4",
				Number:       "4",
				Transactions: []models.EthTransaction{},
			},
			expected: &EthBlockLightDTO{
				Hash:         "0x4",
				Number:       "4",
				Transactions: []string{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dtoBlock := ConvertEthBlockToLightDTO(tt.input)
			if !reflect.DeepEqual(tt.expected, dtoBlock) {
				t.Errorf("Expected %+v, but got %+v", tt.expected, dtoBlock)
			}
		})
	}
}

func TestConvertDTOToEthBlock(t *testing.T) {
	tests := []struct {
		name     string
		input    *EthBlockDTO
		expected *models.EthBlock
	}{
		{
			name: "simple test",
			input: &EthBlockDTO{
				Hash:   "0x1d674edee90409ccbcdcfac5e436f773daf3cf714d119d64ef405e263f1b4ccb",
				Number: "0x118ed01",
				Transactions: []EthTransactionDTO{
					{Hash: "0x6813b6cea801bc68080b7dc843a8b46b55aa34cdb8679e9efb2cfd098fe02202"},
					{Hash: "0xc8985cecf7179ac3c2e0e26ee25c23f8343a920bd3387e1de9b46583d5c3049f"},
				},
			},
			expected: &models.EthBlock{
				Hash:   "0x1d674edee90409ccbcdcfac5e436f773daf3cf714d119d64ef405e263f1b4ccb",
				Number: "0x118ed01",
				Transactions: []models.EthTransaction{
					{Hash: "0x6813b6cea801bc68080b7dc843a8b46b55aa34cdb8679e9efb2cfd098fe02202"},
					{Hash: "0xc8985cecf7179ac3c2e0e26ee25c23f8343a920bd3387e1de9b46583d5c3049f"},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			block := ConvertDTOToEthBlock(tt.input)
			if !reflect.DeepEqual(tt.expected, block) {
				t.Errorf("Expected %+v, but got %+v", tt.expected, block)
			}
		})
	}
}

func TestConvertEthBlockToDTO(t *testing.T) {
	tests := []struct {
		name     string
		input    *models.EthBlock
		expected *EthBlockDTO
	}{
		{
			name: "simple test",
			input: &models.EthBlock{
				Hash:   "0x1d674edee90409ccbcdcfac5e436f773daf3cf714d119d64ef405e263f1b4ccb",
				Number: "0x118ed01",
				Transactions: []models.EthTransaction{
					{Hash: "0x6813b6cea801bc68080b7dc843a8b46b55aa34cdb8679e9efb2cfd098fe02202"},
					{Hash: "0xc8985cecf7179ac3c2e0e26ee25c23f8343a920bd3387e1de9b46583d5c3049f"},
				},
			},
			expected: &EthBlockDTO{
				Hash:   "0x1d674edee90409ccbcdcfac5e436f773daf3cf714d119d64ef405e263f1b4ccb",
				Number: "0x118ed01",
				Transactions: []EthTransactionDTO{
					{Hash: "0x6813b6cea801bc68080b7dc843a8b46b55aa34cdb8679e9efb2cfd098fe02202"},
					{Hash: "0xc8985cecf7179ac3c2e0e26ee25c23f8343a920bd3387e1de9b46583d5c3049f"},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dtoBlock := ConvertEthBlockToDTO(tt.input)
			if !reflect.DeepEqual(tt.expected, dtoBlock) {
				t.Errorf("Expected %+v, but got %+v", tt.expected, dtoBlock)
			}
		})
	}
}
