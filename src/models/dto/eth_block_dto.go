package dto

import (
	"eth-api/src/models"
)

// EthBlockDTO represents a block in the Ethereum blockchain as received from the external API.
type EthBlockDTO struct {
	BaseFeePerGas    string              `json:"baseFeePerGas"`
	Difficulty       string              `json:"difficulty"`
	ExtraData        string              `json:"extraData"`
	GasLimit         string              `json:"gasLimit"`
	GasUsed          string              `json:"gasUsed"`
	Hash             string              `json:"hash"`
	LogsBloom        string              `json:"logsBloom"`
	Miner            string              `json:"miner"`
	MixHash          string              `json:"mixHash"`
	Nonce            string              `json:"nonce"`
	Number           string              `json:"number"`
	ParentHash       string              `json:"parentHash"`
	ReceiptsRoot     string              `json:"receiptsRoot"`
	Sha3Uncles       string              `json:"sha3Uncles"`
	Size             string              `json:"size"`
	StateRoot        string              `json:"stateRoot"`
	Timestamp        string              `json:"timestamp"`
	TotalDifficulty  string              `json:"totalDifficulty"`
	TransactionsRoot string              `json:"transactionsRoot"`
	WithdrawalsRoot  string              `json:"withdrawalsRoot"`
	Uncles           []string            `json:"uncles"`
	Transactions     []EthTransactionDTO `json:"transactions"`
	Withdrawals      []WithdrawalsDTO    `json:"withdrawals"`
}

// WithdrawalsDTO represents a withdrawal in the Ethereum blockchain as received from the external API.
type WithdrawalsDTO struct {
	Address        string `json:"address"`
	Amount         string `json:"amount"`
	Index          string `json:"index"`
	ValidatorIndex string `json:"validatorIndex"`
}

// EthBlockLightDTO represents a simplified version of a block, containing only essential information.
type EthBlockLightDTO struct {
	Hash         string   `json:"hash"`
	Number       string   `json:"number"`
	Transactions []string `json:"transactions"`
}

// ConvertEthBlockToLightDTO converts a model EthBlock to a lighter version, EthBlockLightDTO.
func ConvertEthBlockToLightDTO(block *models.EthBlock) *EthBlockLightDTO {
	dtoBlock := &EthBlockLightDTO{
		Hash:         block.Hash,
		Number:       block.Number,
		Transactions: []string{},
	}

	for _, transaction := range block.Transactions {
		dtoBlock.Transactions = append(dtoBlock.Transactions, transaction.Hash)
	}

	return dtoBlock
}

// ConvertDTOToEthBlock converts an EthBlockDTO to a model EthBlock.
func ConvertDTOToEthBlock(dtoBlock *EthBlockDTO) *models.EthBlock {
	block := &models.EthBlock{
		BaseFeePerGas:    dtoBlock.BaseFeePerGas,
		Difficulty:       dtoBlock.Difficulty,
		ExtraData:        dtoBlock.ExtraData,
		GasLimit:         dtoBlock.GasLimit,
		GasUsed:          dtoBlock.GasUsed,
		Hash:             dtoBlock.Hash,
		LogsBloom:        dtoBlock.LogsBloom,
		Miner:            dtoBlock.Miner,
		MixHash:          dtoBlock.MixHash,
		Nonce:            dtoBlock.Nonce,
		Number:           dtoBlock.Number,
		ParentHash:       dtoBlock.ParentHash,
		ReceiptsRoot:     dtoBlock.ReceiptsRoot,
		Sha3Uncles:       dtoBlock.Sha3Uncles,
		Size:             dtoBlock.Size,
		StateRoot:        dtoBlock.StateRoot,
		Timestamp:        dtoBlock.Timestamp,
		TotalDifficulty:  dtoBlock.TotalDifficulty,
		TransactionsRoot: dtoBlock.TransactionsRoot,
		WithdrawalsRoot:  dtoBlock.WithdrawalsRoot,
	}

	for _, dtoTransaction := range dtoBlock.Transactions {
		transaction := ConvertDTOToEthTransaction(&dtoTransaction)
		block.Transactions = append(block.Transactions, *transaction)
	}

	for _, dtoUncle := range dtoBlock.Uncles {
		uncle := models.Uncles{Uncles: dtoUncle}
		block.Uncles = append(block.Uncles, uncle)
	}

	for _, dtoWithdrawal := range dtoBlock.Withdrawals {
		withdrawal := models.Withdrawals{
			Address:        dtoWithdrawal.Address,
			Amount:         dtoWithdrawal.Amount,
			Index:          dtoWithdrawal.Index,
			ValidatorIndex: dtoWithdrawal.ValidatorIndex,
		}
		block.Withdrawals = append(block.Withdrawals, withdrawal)
	}

	return block
}

// ConvertEthBlockToDTO converts a model EthBlock to an EthBlockDTO.
func ConvertEthBlockToDTO(block *models.EthBlock) *EthBlockDTO {
	dtoBlock := &EthBlockDTO{
		BaseFeePerGas:    block.BaseFeePerGas,
		Difficulty:       block.Difficulty,
		ExtraData:        block.ExtraData,
		GasLimit:         block.GasLimit,
		GasUsed:          block.GasUsed,
		Hash:             block.Hash,
		LogsBloom:        block.LogsBloom,
		Miner:            block.Miner,
		MixHash:          block.MixHash,
		Nonce:            block.Nonce,
		Number:           block.Number,
		ParentHash:       block.ParentHash,
		ReceiptsRoot:     block.ReceiptsRoot,
		Sha3Uncles:       block.Sha3Uncles,
		Size:             block.Size,
		StateRoot:        block.StateRoot,
		Timestamp:        block.Timestamp,
		TotalDifficulty:  block.TotalDifficulty,
		TransactionsRoot: block.TransactionsRoot,
		WithdrawalsRoot:  block.WithdrawalsRoot,
	}

	for _, transaction := range block.Transactions {
		dtoTransaction := ConvertEthTransactionToDTO(&transaction)
		dtoBlock.Transactions = append(dtoBlock.Transactions, *dtoTransaction)
	}

	for _, uncle := range block.Uncles {
		dtoBlock.Uncles = append(dtoBlock.Uncles, uncle.Uncles)
	}

	for _, withdrawal := range block.Withdrawals {
		dtoWithdrawal := WithdrawalsDTO{
			Address:        withdrawal.Address,
			Amount:         withdrawal.Amount,
			Index:          withdrawal.Index,
			ValidatorIndex: withdrawal.ValidatorIndex,
		}
		dtoBlock.Withdrawals = append(dtoBlock.Withdrawals, dtoWithdrawal)
	}

	return dtoBlock
}
