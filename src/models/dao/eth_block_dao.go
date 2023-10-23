package dao

import (
	"eth-api/src/models"
	"gorm.io/gorm"
)

// EthBlockDAO defines the interface for data access objects that interact with Ethereum blocks in the database.
type EthBlockDAO interface {
	GetBlockByNumber(number string) (*models.EthBlock, error)
	SaveBlock(block *models.EthBlock) (*models.EthBlock, error)
	GetLatestBlocks(page int, limit int) ([]models.EthBlock, error)
	GetLatestBlockNumber() (string, error)
}

// ethBlockDAO is the concrete implementation of EthBlockDAO that uses GORM to interact with the database.
type ethBlockDAO struct {
	DB *gorm.DB
}

// NewEthBlockDAO returns a new instance of EthBlockDAO with the given database connection.
func NewEthBlockDAO(db *gorm.DB) EthBlockDAO {
	return &ethBlockDAO{
		DB: db,
	}
}

// GetBlockByNumber retrieves a block by its number, including its transactions.
func (e *ethBlockDAO) GetBlockByNumber(number string) (*models.EthBlock, error) {
	var block models.EthBlock
	result := e.DB.Preload("Transactions").Where("number_ = ?", number).First(&block)
	return &block, result.Error
}

// SaveBlock saves a new block to the database.
func (e *ethBlockDAO) SaveBlock(block *models.EthBlock) (*models.EthBlock, error) {
	result := e.DB.Create(block)
	if result.Error != nil {
		return nil, result.Error
	}
	return block, nil
}

// GetLatestBlocks retrieves the latest blocks with pagination.
func (e *ethBlockDAO) GetLatestBlocks(page int, limit int) ([]models.EthBlock, error) {
	var blocks []models.EthBlock
	offset := (page - 1) * limit
	result := e.DB.Preload("Transactions").Order("number_ desc").Limit(limit).Offset(offset).Find(&blocks)
	return blocks, result.Error
}

// GetLatestBlockNumber retrieves the block number of the latest block.
func (e *ethBlockDAO) GetLatestBlockNumber() (string, error) {
	row := e.DB.Raw("SELECT id, number_ FROM`eth_blocks` ORDER BY CONVERT(number_, SIGNED) DESC LIMIT 1").Row()
	var id int
	var number string

	err := row.Scan(&id, &number)
	if err != nil {
		return "", err
	}

	return number, nil
}
