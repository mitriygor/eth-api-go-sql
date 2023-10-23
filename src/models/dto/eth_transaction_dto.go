package dto

import "eth-api/src/models"

// EthTransactionDTO represents a transaction in the Ethereum blockchain as received from the external API.
type EthTransactionDTO struct {
	BlockHash            string          `json:"blockHash"`
	BlockNumber          string          `json:"blockNumber"`
	ChainId              string          `json:"chainId"`
	From                 string          `json:"from"`
	Gas                  string          `json:"gas"`
	GasPrice             string          `json:"gasPrice"`
	Hash                 string          `json:"hash"`
	Input                string          `json:"input"`
	MaxFeePerGas         string          `json:"maxFeePerGas"`
	MaxPriorityFeePerGas string          `json:"maxPriorityFeePerGas"`
	Nonce                string          `json:"nonce"`
	R                    string          `json:"r"`
	S                    string          `json:"s"`
	To                   string          `json:"to"`
	TransactionIndex     string          `json:"transactionIndex"`
	Type                 string          `json:"type"`
	V                    string          `json:"v"`
	Value                string          `json:"value"`
	AccessList           []AccessListDTO `json:"accessList"`
}

// AccessListDTO represents the access list in a transaction, detailing which storage keys are accessed by the transaction.
type AccessListDTO struct {
	Address     string   `json:"address"`
	StorageKeys []string `json:"storageKeys"`
}

// EthTransactionLightDTO represents a simplified version of a transaction, containing only essential information.
type EthTransactionLightDTO struct {
	BlockHash   string          `json:"blockHash"`
	BlockNumber string          `json:"blockNumber"`
	ChainId     string          `json:"chainId"`
	From        string          `json:"from"`
	Hash        string          `json:"hash"`
	To          string          `json:"to"`
	Value       string          `json:"value"`
	AccessList  []AccessListDTO `json:"accessList"`
}

// ConvertDTOToEthTransaction converts an EthTransactionDTO to a model EthTransaction.
func ConvertDTOToEthTransaction(dtoTransaction *EthTransactionDTO) *models.EthTransaction {
	transaction := &models.EthTransaction{
		BlockHash:            dtoTransaction.BlockHash,
		BlockNumber:          dtoTransaction.BlockNumber,
		ChainId:              dtoTransaction.ChainId,
		From:                 dtoTransaction.From,
		Gas:                  dtoTransaction.Gas,
		GasPrice:             dtoTransaction.GasPrice,
		Hash:                 dtoTransaction.Hash,
		Input:                dtoTransaction.Input,
		MaxFeePerGas:         dtoTransaction.MaxFeePerGas,
		MaxPriorityFeePerGas: dtoTransaction.MaxPriorityFeePerGas,
		Nonce:                dtoTransaction.Nonce,
		R:                    dtoTransaction.R,
		S:                    dtoTransaction.S,
		To:                   dtoTransaction.To,
		TransactionIndex:     dtoTransaction.TransactionIndex,
		Type:                 dtoTransaction.Type,
		V:                    dtoTransaction.V,
		Value:                dtoTransaction.Value,
	}

	for _, dtoAccessList := range dtoTransaction.AccessList {
		accessList := models.AccessList{
			Address: dtoAccessList.Address,
		}

		for _, storageKey := range dtoAccessList.StorageKeys {
			accessList.StorageKeys = append(accessList.StorageKeys, models.StorageKey{StorageKey: storageKey})
		}

		transaction.AccessList = append(transaction.AccessList, accessList)
	}

	return transaction
}

// ConvertEthTransactionToDTO converts a model EthTransaction to an EthTransactionDTO.
func ConvertEthTransactionToDTO(transaction *models.EthTransaction) *EthTransactionDTO {
	dtoTransaction := &EthTransactionDTO{
		BlockHash:            transaction.BlockHash,
		BlockNumber:          transaction.BlockNumber,
		ChainId:              transaction.ChainId,
		From:                 transaction.From,
		Gas:                  transaction.Gas,
		GasPrice:             transaction.GasPrice,
		Hash:                 transaction.Hash,
		Input:                transaction.Input,
		MaxFeePerGas:         transaction.MaxFeePerGas,
		MaxPriorityFeePerGas: transaction.MaxPriorityFeePerGas,
		Nonce:                transaction.Nonce,
		R:                    transaction.R,
		S:                    transaction.S,
		To:                   transaction.To,
		TransactionIndex:     transaction.TransactionIndex,
		Type:                 transaction.Type,
		V:                    transaction.V,
		Value:                transaction.Value,
	}

	for _, accessList := range transaction.AccessList {
		dtoAccessList := AccessListDTO{
			Address: accessList.Address,
		}

		for _, storageKey := range accessList.StorageKeys {
			dtoAccessList.StorageKeys = append(dtoAccessList.StorageKeys, storageKey.StorageKey)
		}

		dtoTransaction.AccessList = append(dtoTransaction.AccessList, dtoAccessList)
	}

	return dtoTransaction
}

// ConvertEthTransactionToLightDTO converts a model EthTransaction to an EthTransactionLightDTO.
func ConvertEthTransactionToLightDTO(transaction *models.EthTransaction) *EthTransactionLightDTO {
	dtoTransaction := &EthTransactionLightDTO{
		BlockHash:   transaction.BlockHash,
		BlockNumber: transaction.BlockNumber,
		ChainId:     transaction.ChainId,
		From:        transaction.From,
		Hash:        transaction.Hash,
		To:          transaction.To,
		Value:       transaction.Value,
	}

	return dtoTransaction
}
