package dao

import (
	"eth-api/src/models"
	"gorm.io/gorm"
)

// EthTransactionDAO defines the interface for data access objects that interact with Ethereum transactions in the database.
type EthTransactionDAO interface {
	GetTransactionByHash(hash string) (*models.EthTransaction, error)
	GetTransactionsByAddress(address string) ([]models.EthTransaction, error)
	SaveTransaction(transaction *models.EthTransaction) (*models.EthTransaction, error)
}

// ethTransactionDAO is the concrete implementation of EthTransactionDAO that uses GORM to interact with the database.
type ethTransactionDAO struct {
	DB *gorm.DB
}

// NewEthTransactionDAO returns a new instance of EthTransactionDAO with the given database connection.
func NewEthTransactionDAO(db *gorm.DB) EthTransactionDAO {
	return &ethTransactionDAO{
		DB: db,
	}
}

// GetTransactionByHash retrieves a transaction by its hash.
func (e *ethTransactionDAO) GetTransactionByHash(hash string) (*models.EthTransaction, error) {
	var transaction models.EthTransaction
	result := e.DB.Where("hash = ?", hash).First(&transaction)
	return &transaction, result.Error
}

// SaveTransaction saves a new transaction to the database.
func (e *ethTransactionDAO) SaveTransaction(transaction *models.EthTransaction) (*models.EthTransaction, error) {
	result := e.DB.Create(&transaction)
	return transaction, result.Error
}

// GetTransactionsByAddress retrieves all transactions associated with an address.
func (e *ethTransactionDAO) GetTransactionsByAddress(address string) ([]models.EthTransaction, error) {
	var transactions []models.EthTransaction
	result := e.DB.Where("from_ = ? OR to_ = ?", address, address).Find(&transactions)
	return transactions, result.Error
}
